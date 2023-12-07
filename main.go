package main

import (
	"httrouter/config"
	"net/http"
)

func main() {
	// Model view controller
	http.ListenAndServe(":9090", config.Routes())
}
