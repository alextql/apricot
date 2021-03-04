package database

import (
	"database/sql"

	"github.com/alex-techs/apricot/apricot/helpers"
	"github.com/alex-techs/apricot/apricot/logger"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

func InitSqlite3(dbfile string) error {
	dbfile, err = helpers.FileAbsPath(dbfile)
	if err != nil {
		return err
	}

	if Db, err = sql.Open("sqlite3", dbfile); err != nil {
		return err
	}

	logger.Logger.WithField("file", dbfile).Info("sqlite database started")
	return nil
}
