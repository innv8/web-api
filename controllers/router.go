package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (b *Base) InitRouter() {
	b.Router = mux.NewRouter()

	b.Router.HandleFunc("/", b.HomeController)
	b.Router.HandleFunc("/user", b.CreateUserController).Methods(http.MethodPost)
	b.Router.HandleFunc("/users", b.ListUsers).Methods(http.MethodGet)
}
