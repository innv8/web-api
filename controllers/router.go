package controllers

import "github.com/gorilla/mux"

func (b *Base) InitRouter() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/", b.HomeController)
}
