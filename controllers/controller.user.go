package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (b *Base) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id = vars["id"]
	idInt, _ := strconv.Atoi(id)
	if idInt == 0 {
		middleware.JSONResponse(w, http.StatusBadGateway, "wrong id")
		return
	}

	user, err := models.GetUser(idInt, b.DB)
	if err != nil {
		middleware.JSONResponse(w, http.StatusNotFound, "user not found")
		return
	}

	middleware.JSONResponse(w, http.StatusOK, user)
}

func (b *Base) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// read the data from the client
	var data entities.User
	var err error

	vars := mux.Vars(r)
	var id = vars["id"]
	idInt, _ := strconv.Atoi(id)
	if idInt == 0 {
		middleware.JSONResponse(w, http.StatusBadGateway, "wrong id")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("an error occured : %v", err)
		middleware.JSONResponse(w, http.StatusBadRequest, "no data was sent")
		return
	}

	err = models.UpdateUser(idInt, data, b.DB)
	if err != nil {
		middleware.JSONResponse(w, http.StatusBadRequest, "an error occurred")
		return
	}
	middleware.JSONResponse(w, http.StatusOK, "ok")

}

func (b *Base) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id = vars["id"]
	idInt, _ := strconv.Atoi(id)
	if idInt == 0 {
		middleware.JSONResponse(w, http.StatusBadGateway, "wrong id")
		return
	}

	err := models.DeleteUser(idInt, b.DB)
	if err != nil {
		middleware.JSONResponse(w, http.StatusBadRequest, "an error occurred")
		return
	}
	middleware.JSONResponse(w, http.StatusOK, "ok")

}
