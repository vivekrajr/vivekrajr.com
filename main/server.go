package main

import (
	"github.com/vivekrajr/vivekrajr.com/routes"
	"net/http"
)

func main() {
	r := routes.AllRoutes()

	http.ListenAndServe(":8000", r)
}
