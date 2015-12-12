package main

import (
	"github.com/vivekrajr/vivekrajr.com/routes"
	"net/http"
)

func main() {
	r := routes.AllRoutes()

	http.ListenAndServe("127.0.0.1:8000", r)
}
