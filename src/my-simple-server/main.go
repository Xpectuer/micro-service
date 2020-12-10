package main

import (
	"log"

	"net/http"
	"os"

	"github.com/Xpectuer/mircro-service/src/my-simple-server/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	/**
	Serve Mux  is a map spcifies
	the routers and handler funcs
	*/
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	log.Println("Starting Server")
	err := http.ListenAndServe(":9091", sm)
	log.Fatal(err)
}
