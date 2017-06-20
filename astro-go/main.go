package main

import (
	astro "github.com/astromechio/astro-go"

	"log"

	"github.com/gorilla/mux"
)

func main() {
	server, err := astro.DefaultEnvServer("add")
	if err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()

	mux.Methods("POST").Path("/").HandlerFunc(AddHandler())

	server.ListenAndServe(mux)
}
