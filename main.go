package main

import (
	"log"
	"net/http"

	"github.com/carlso70/go-todo/routing"
)

func main() {
	router := routing.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
