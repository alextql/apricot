package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alex-techs/apricot/apricot"
	"github.com/alex-techs/apricot/apricot/config"
	"github.com/alex-techs/apricot/apricot/database"
	"github.com/alex-techs/apricot/apricot/logger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Apricot"
	app.Flags = []cli.Flag{}

	app.Action = func(cmdline *cli.Context) error {
		if err := config.LoadConfig(); err != nil {
			return err
		}

		logger.InitLog()

		if err := database.InitSqlite3(); err != nil {
			return errors.New(fmt.Sprintf("failed to create database: %s", err))
		}

		app := apricot.NewApricot()
		if err := app.Listen(); err != nil {
			return errors.New(fmt.Sprintf("apricot listen crashed , error: %s", err))
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
