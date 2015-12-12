package controllers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ic HomeController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)

	data := struct {
		Year string
	}{
		Year: strconv.Itoa(time.Now().Year()),
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
