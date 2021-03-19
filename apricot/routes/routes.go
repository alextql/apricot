package routes

import (
	"github.com/alex-techs/apricot/apricot/handlers"
	"github.com/kataras/iris/v12"
)

type Route struct {
	apricot *iris.Application
}

func NewRoute(apricot *iris.Application) *Route {
	return &Route{
		apricot: apricot,
	}
}

func (r *Route) Register() {
	url := handlers.NewUrlHandle()
	urls := r.apricot.Party("/url")
	{
		urls.Post("/make", url.Make)
		urls.Get("/{code}", url.Parse)
	}

	m := handlers.NewMockHandle()
	mock := r.apricot.Party("/mock")
	{
		mock.Post("/make", m.Make)
		mock.Any("/{route:path}", m.Parse)
	}
}
