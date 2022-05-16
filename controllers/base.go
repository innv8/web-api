package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Base - this is a struct to store shared resources
// e.g db connections, the router.
type Base struct {
	Router *mux.Router
}

// Init - initializes the base requirements
func (b *Base) Init() error {
	b.InitRouter()

	return nil
}

func (b *Base) StartAPI() {
	port := ":5000"

	log.Printf("running on port %s", port)
	err := http.ListenAndServe(port, b.Router)
	if err != nil {
		log.Printf("unable to start API because %s", err.Error())
	}
}
