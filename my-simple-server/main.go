package main

import (
	"log"
	"my-simple-server/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	/**
	Serve Mux  is a map spcifies
	the routers and handler funcs
	*/
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	log.Println("Starting Server")
	err := http.ListenAndServe(":9091", sm)
	log.Fatal(err)
}
