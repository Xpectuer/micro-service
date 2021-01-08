package main

import (
	"net"
	"os"

	protos "currency/protos/currency"
	"currency/server"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// var bindingAdress = env.String("BIND_ADDRESS", false, ":9092", "Bind adress for the server")

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	log.Info("Serving gRPC on port 9093...")
	gs.Serve(l)
}
