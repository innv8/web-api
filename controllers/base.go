package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/innv8/web-api/utils"
)

// Base - this is a struct to store shared resources
// e.g db connections, the router.
type Base struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init - initializes the base requirements
func (b *Base) Init() (err error) {
	b.InitRouter()

	b.DB, err = utils.DBConnect()
	if err != nil {
		return err
	}

	return nil
}

func (b *Base) StartAPI() {
	port := os.Getenv("PORT")

	log.Printf("running on port %s", port)
	err := http.ListenAndServe(port, b.Router)
	if err != nil {
		log.Printf("unable to start API because %s", err.Error())
	}
}
