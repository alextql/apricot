package apricot

import (
	"github.com/alex-techs/apricot/apricot/config"
	"github.com/alex-techs/apricot/apricot/handlers"
	"github.com/kataras/iris/v12"
)

type Apricot struct {
	apricot *iris.Application
}

func NewApricot() *Apricot {
	return &Apricot{
		apricot: iris.New(),
	}
}

func (a *Apricot) Listen() error {
	a.registerRoutes()
	a.apricot.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))

	return a.apricot.Listen(config.C.Apricot.Listen)
}

func (a *Apricot) registerRoutes() {
	url := handlers.NewUrlHandle()
	urls := a.apricot.Party("/url")
	{
		urls.Post("/make", url.Make)
		urls.Get("/{code}", url.Parse)
	}
}
