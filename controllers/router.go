package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (b *Base) InitRouter() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/", b.HomeController)
	b.Router.HandleFunc("/users", b.CreateUserController).Methods(http.MethodPost)
}
