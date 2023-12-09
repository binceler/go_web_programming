package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"httrouter/admin/helpers"
	"net/http"
	"text/template"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var view, err = template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
