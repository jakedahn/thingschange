package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/context"
	// "github.com/gorilla/mux"
	"github.com/gorilla/pat"

	"github.com/jakedahn/thingschange/api"
)

func SetupRoutes() *pat.Router {
	r := pat.New()

	r.Get("/api/v0/checks", api.V0GetLists)
	http.Handle("/", r)

	return r
}

func main() {
	var (
		bind_address = flag.String("bind_address", "0.0.0.0", "Specify ip to bind to")
		bind_port    = flag.String("bind_port", "3000", "Specify the port to bind to")
	)
	flag.Parse()

	SetupRoutes()

	var bind_str = fmt.Sprintf("%s:%s", *bind_address, *bind_port)
	log.Printf("Starting server: http://%s", bind_str)
	err := http.ListenAndServe(bind_str, nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
