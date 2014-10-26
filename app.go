package main

import (
	"flag"
	"log"
	"net/http"

	// "github.com/gorilla/context"
	"github.com/gorilla/pat"

	"api"
	"models"
)

func SetupRoutes() *mux.Router {
	r := pat.New()
	r.Get("/api/v0/checks", V0GetLists)
	http.Handle("/", r)
	return r
}

func main() {
	var (
		bind_address = flag.String("bind_address", "0.0.0.0", "Specify ip to bind to")
		bind_port    = flag.String("bind_port", "3000", "Specify the port to bind to")
		rabbit_pass  = flag.String("rabbit_pass", "guest", "Specify the rabbit password")
	)

	flag.Parse()

	router := SetupRoutes()

	var bind_str = fmt.Sprintf("http://%s:%s/", *bind_address, *bind_port)
	log.Printf("Starting server: %s", bind_str)
	err := http.ListenAndServe(bind_str, nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
