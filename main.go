package main

import (
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", Anasayfa)
	http.HandleFunc("/detay", Detay)
	http.ListenAndServe(":9090", nil)
}

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "navbar.html")
	data := "Go'dan gelen veri..."

	/*data := make(map[string]interface{})
	data["sayilar"] = []int{1, 2, 3, 4, 5}
	data["is_admin"] = false
	data["sayi"] = 10
	data := Data{
		Veri:    "Büşra",
		Sayilar: []int{1, 2, 3, 4},
	} */
	view.ExecuteTemplate(w, "anasayfa", data)
}

func Detay(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detay.html", "navbar.html")
	view.ExecuteTemplate(w, "detay", nil)
}

/* type Data struct {
	Veri    string
	Sayilar []int
} */
