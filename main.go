package main

import (
	"log"

	"github.com/innv8/web-api/controllers"
)

func main() {
	var base controllers.Base
	var err error

	err = base.Init()
	if err != nil {
		log.Fatalf("unable to initialize app because %v", err)
	}

	// start the API
	base.StartAPI()
}
