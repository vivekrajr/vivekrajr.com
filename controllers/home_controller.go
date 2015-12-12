package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ic HomeController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello Universe!\n")
}
