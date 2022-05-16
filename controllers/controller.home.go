package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (b *Base) HomeController(w http.ResponseWriter, r *http.Request) {
	log.Println("server hit by client")
	currentTime := time.Now().Format("2006 01 02 15:04:05")
	_, _ = fmt.Fprintln(w, "hello world ", currentTime)
}
