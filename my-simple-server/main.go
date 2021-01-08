package main

import (
	"context"
	"log"

	"net/http"
	"os"
	"os/signal"
	"time"
	// alias for handlers
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"my-simple-server/data"
	"my-simple-server/handlers"

	"github.com/nicholasjackson/env"
	"google.golang.org/grpc/protos"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9092", "Bind adress for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l, v)

	conn, err := grpc.Dial("localhost:9093")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create client
	cc := protos.NewCurrencyClient(conn)

	//create Handlers (constructor styled injection)
	ph := handlers.NewHandlers(l, v, cc)
	/**
	Serve Mux  is a map spcifies
	the routers and handler funcs
	*/
	sm := mux.NewRouter()
	log.Println("Registering Handlers... ")
	//sm.Handle("/products", ph)
	// In go public function started with Capital letter

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingleProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProducts)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	//----
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:        *bindAddress,
		Handler:     ch(sm),
		IdleTimeout: 120 * time.Second,
		// timeout for request reader
		ReadTimeout: 1 * time.Second,
		// timeout for response writer
		WriteTimeout: 1 * time.Second,
	}
	// graceful shutdown

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// specifies the channel
	sigChan := make(chan os.Signal)

	// relay the incoming signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// consume
	sig := <-sigChan
	l.Println("Received terminate, gracefully shutdown", sig)

	// allow handlers to gracefully shutdown in 30s
	// after 30s , forcefully close it
	// In golang 1.15 Version, cancel func cannot be swallowed
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)

}
