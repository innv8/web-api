package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/innv8/web-api/entities"
	"github.com/innv8/web-api/middleware"
	"github.com/innv8/web-api/models"
)

func (b *Base) CreateUserController(w http.ResponseWriter, r *http.Request) {
	// read the data from the client
	var data entities.User
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("an error occured : %v", err)
		middleware.JSONResponse(w, http.StatusBadRequest, "no data was sent")
		return
	}

	err = models.CreateUser(data, b.DB)
	if err != nil {
		middleware.JSONResponse(w, http.StatusBadRequest, "creating user failed")
		return
	}

	middleware.JSONResponse(w, http.StatusCreated, "user created")
}

func (b *Base) ListUsers(w http.ResponseWriter, r *http.Request) {
	// check if there is data in the first_name url key
	urlValues := r.URL.Query()
	var firstName = urlValues.Get("first_name")
	users, err := models.ListUsers(firstName, b.DB)
	if err != nil {
		middleware.JSONResponse(w, http.StatusBadRequest, "unable to list users")
		return
	}
	middleware.JSONResponse(w, http.StatusOK, users)
}
