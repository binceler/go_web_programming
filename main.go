package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", Anasayfa)
	http.ListenAndServe(":8080", nil)
}

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Merhaba Dünya")
	view, _ := template.ParseFiles("index.html")
	//data := "Go'dan gelen veri..."

	data := make(map[string]interface{})
	data["sayilar"] = []int{1, 2, 3, 4, 5}
	data["is_admin"] = false
	data["sayi"] = 10
	/* data := Data{
		Veri:    "Büşra",
		Sayilar: []int{1, 2, 3, 4},
	} */
	view.Execute(w, data)
}

/* type Data struct {
	Veri    string
	Sayilar []int
} */
