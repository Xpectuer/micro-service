package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// make things happen
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	//
	// })
	// log.New( where to out put(file or stdout), prefix string, flags )
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	log.Println("Starting Server")
	err := http.ListenAndServe(":9091", nil)
	log.Fatal(err)
}
