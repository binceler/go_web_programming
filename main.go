package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
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
	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("file")
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
}
