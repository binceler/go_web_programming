package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/", Anasayfa)
	//http.HandleFunc("/detay", Detay)
	r := httprouter.New()
	r.GET("/", Anasayfa)
	r.POST("/deneme", Deneme)
	http.ListenAndServe(":9090", r)
}

func Anasayfa(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

func Deneme(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	check := r.FormValue("check")
	formSelect := r.FormValue("select")
	fmt.Println(check)
	fmt.Println(formSelect)
}
