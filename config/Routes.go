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
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
