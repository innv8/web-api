package main

import (
	"log"

	"github.com/innv8/web-api/controllers"
	"github.com/joho/godotenv"
)

func main() {
	var base controllers.Base
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("unable to open .env because %v", err)
	}

	err = base.Init()
	if err != nil {
		log.Fatalf("unable to initialize app because %v", err)
	}

	// start the API
	base.StartAPI()
}
