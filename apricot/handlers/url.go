package handlers

import (
	"fmt"
	"github.com/alex-techs/apricot/apricot/database"
	"github.com/alex-techs/apricot/apricot/helpers"
	"github.com/kataras/iris/v12"
)

type Url struct {
}

func NewUrlHandle() *Url {
	return &Url{}
}

func (u *Url) Make(c iris.Context) {
	var code, uri string
	var response = helpers.NewResponse(c)

	if uri = c.PostValue("url"); uri == "" {
		response.Error("url required")
		return
	}

	// 去重
	hash := helpers.Sha1Encode(helpers.UrlEscape(uri))
	if err := database.Db.QueryRow("select code from short_url where hash = ?", hash).Scan(&code); err == nil {
		response.Success(map[string]string{
			"code": code,
			"path": fmt.Sprintf("/url/%s", code),
		})
		return
	}

	// 插入新的
	stmt, _ := database.Db.Prepare("insert into short_url (hash, code, url) values (? , ? , ?)")
	added, err := stmt.Exec(hash, "", uri)
	if err != nil {
		response.Error(err.Error())
		return
	}

	id, _ := added.LastInsertId()
	code = helpers.Base64Encode(id)
	stmt, _ = database.Db.Prepare("update short_url set code = ? where id = ?")
	if _, err := stmt.Exec(code, id); err != nil {
		response.Error(err.Error())
		return
	}

	response.Success(map[string]string{
		"code": code,
		"path": fmt.Sprintf("/url/%s", code),
	})
}

func (u *Url) Parse(c iris.Context) {
	var code, uri string
	var response = helpers.NewResponse(c)

	if code = c.Params().Get("code"); code == "" {
		response.Error("/url/{code} code required")
		return
	}

	if err := database.Db.QueryRow("select url from short_url where code = ?", code).Scan(&uri); err == nil {
		c.Redirect(uri, iris.StatusMovedPermanently)
		return
	}

	c.StatusCode(404)
}
