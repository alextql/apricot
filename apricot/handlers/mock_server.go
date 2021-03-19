package handlers

import (
	"fmt"
	"strings"

	"github.com/alex-techs/apricot/apricot/database"
	"github.com/alex-techs/apricot/apricot/helpers"
	"github.com/go-playground/validator"
	"github.com/kataras/iris/v12"
)

type MockServer struct {
}

func NewMockHandle() *MockServer {
	return &MockServer{}
}

type MockBody struct {
	Route string `json:"route" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func (m *MockServer) Make(c iris.Context) {
	var post MockBody
	var validate = validator.New()
	var response = helpers.NewResponse(c)
	if err := c.ReadJSON(&post); err != nil {
		response.Error(err.Error())
		return
	}

	if err := validate.Struct(post); err != nil {
		response.Error(err.Error())
		return
	}

	// 去掉路由开始的 /
	if strings.HasPrefix(post.Route, "/") {
		post.Route = post.Route[1:len(post.Route)]
	}

	// 去重
	if err := database.Db.QueryRow("select route from mock_server where route = ?", post.Route).Scan(post.Route); err == nil {
		stmt, _ := database.Db.Prepare("update mock_server set body = ? where route = ?")
		if _, err := stmt.Exec(post.Body, post.Route); err != nil {
			response.Error(err.Error())
			return
		}
		return
	}

	// 插入新的
	stmt, _ := database.Db.Prepare("insert into mock_server (route, body) values (? , ?)")
	_, err := stmt.Exec(post.Route, post.Body)
	if err != nil {
		response.Error(err.Error())
		return
	}

	response.Success(map[string]string{
		"route": post.Route,
		"path":  fmt.Sprintf("/mock/%s", post.Route),
	})
}

func (m *MockServer) Parse(c iris.Context) {
	var route, body string
	var response = helpers.NewResponse(c)

	if route = c.Params().GetTrim("route"); route == "" {
		response.Error("/mock/{route} route required")
		return
	}

	if err := database.Db.QueryRow("select body from mock_server where route = ?", route).Scan(&body); err == nil {
		response.Raw(body)
		return
	}

	c.StatusCode(404)
}
