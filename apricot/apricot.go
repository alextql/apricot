package apricot

import (
	"fmt"

	"github.com/alex-techs/apricot/apricot/config"
	"github.com/alex-techs/apricot/apricot/database"
	"github.com/alex-techs/apricot/apricot/logger"
	"github.com/alex-techs/apricot/apricot/options"
	"github.com/alex-techs/apricot/apricot/routes"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Apricot struct {
	apricot *iris.Application
	logger  *logrus.Logger
}

func NewApricotEngine(opts *options.Options) (*Apricot, error) {
	if err := config.LoadConfig(opts.ConfigFile); err != nil {
		return nil, err
	}

	logger.InitLog()

	if err := database.InitSqlite3(opts.DatabaseFile); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create database: %s", err))
	}

	return &Apricot{
		apricot: iris.New(),
	}, nil
}

func (a *Apricot) Listen() error {
	a.apricot.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	route := routes.NewRoute(a.apricot)
	route.Register()

	return a.apricot.Listen(config.Get().Apricot.Listen)
}
