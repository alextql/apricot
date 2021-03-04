package database

import (
	"database/sql"
	"github.com/alex-techs/apricot/apricot/config"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

func InitSqlite3() error {
	if Db, err = sql.Open("sqlite3", config.C.Database.Path); err != nil {
		return err
	}

	return nil
}
