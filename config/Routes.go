package config

import (
	"github.com/julienschmidt/httprouter"
	admin "httrouter/admin/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/admin", admin.Dashboard{}.Index)

	r.ServeFiles("/admin/*filepath", http.Dir("admin/"))
	return r
}
