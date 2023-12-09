package config

import (
	"github.com/julienschmidt/httprouter"
	admin "httrouter/admin/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMIN
	r.GET("/admin", admin.Dashboard{}.Index)

	//SERVE FILES
	r.ServeFiles("/admin/*filepath", http.Dir("admin/"))
	return r
}
