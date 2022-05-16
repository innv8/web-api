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
	b.Router.HandleFunc("/user/{id}", b.GetUser).Methods(http.MethodGet)
	b.Router.HandleFunc("/user/{id}", b.UpdateUser).Methods(http.MethodPut)
	b.Router.HandleFunc("/user/{id}", b.DeleteUser).Methods(http.MethodDelete)
}
