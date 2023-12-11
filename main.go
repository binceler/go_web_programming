package main

import (
	"httrouter/config"
	"net/http"
)

func main() {
	// Model view controller
	//admin_models.Post{}.Migrate()
	/*admin_models.Post{
		Title: "Büşra İnceler - Gopher",
		Slug:  "busra-inceler-gopher",
	}.Add()*/
	//getAll := admin_models.Post{}.GetAll("description = ?", "")
	//fmt.Println(getAll)
	//post := admin_models.Post{}.Get("description = ?", "123")
	//post.Update("description", "123")
	//post.Updates(admin_models.Post{Title: "Anksiyete nasıl önlenir", Slug: "anksiyete-nasıl-önlenir"})
	//post.Delete()
	http.ListenAndServe(":6070", config.Routes())
}
