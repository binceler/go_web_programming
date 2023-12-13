package controllers

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"httrouter/admin/helpers"
	admin_models "httrouter/admin/models"
	"net/http"
	"text/template"
)

type Categories struct{}

func (categories Categories) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("categories/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["categories"] = admin_models.Category{}.GetAll()
	data["alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (categories Categories) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	CategoryTitle := r.FormValue("category-name")
	Slug := slug.Make(CategoryTitle)

	admin_models.Category{Title: CategoryTitle, Slug: Slug}.Add()
	helpers.SetAlert(w, r, "Kategori Başarıyla Kaydedildi!")

	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}

func (categories Categories) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	category := admin_models.Category{}.Get(params.ByName("id"))
	category.Delete()

	helpers.SetAlert(w, r, "Kayıt Başarıyla Silindi!")
	http.Redirect(w, r, "/admin/kategoriler", http.StatusSeeOther)
}
