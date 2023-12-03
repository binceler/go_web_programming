package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/", Anasayfa)
	//http.HandleFunc("/detay", Detay)
	r := httprouter.New()
	r.GET("/yazilar/:slug", Anasayfa)
	http.ListenAndServe(":9090", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	data := params.ByName("slug")
	view.Execute(w, data)
}
