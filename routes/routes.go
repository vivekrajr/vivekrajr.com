package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vivekrajr/vivekrajr.com/controllers"
	"net/http"
)

type Routes struct {
	Route *httprouter.Router
}

func new() *Routes {
	return &Routes{
		Route: httprouter.New(),
	}
}

func AllRoutes() *httprouter.Router {
	routes := new()

	IndexRoutes(routes)

	return routes.Route
}

func IndexRoutes(r *Routes) {
	ic := controllers.NewHomeController()

	r.Route.GET("/", ic.Index)
	r.Route.GET("/home/", ic.Index)
	r.Route.GET("/home/index", ic.Index)
	r.Route.ServeFiles("/static/*filepath", http.Dir("templates/static"))
}
