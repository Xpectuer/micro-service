/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/Xpectuer/micro-service/currency/data"
	protos "github.com/Xpectuer/micro-service/currency/protos/currency"
	"github.com/Xpectuer/micro-service/currency/server"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// var bindingAdress = env.String("BIND_ADDRESS", false, ":9092", "Bind adress for the server")

func main() {

	log := hclog.Default()
	printFlag(log)
	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}
	// create a new gRPC server
	gs := grpc.NewServer()
	// create a new Currency Server
	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9093))
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	log.Info("Serving gRPC on", "port", 9093)
	gs.Serve(l)
}

// PrintFlag Function output a Custom Banner from ./banner.txt
func printFlag(l hclog.Logger) {
	content, err := ioutil.ReadFile("./banner.txt")
	if err != nil {
		l.Error("Unable to read file", err)
		return
	}

	fmt.Println("---------Starting Server------------")
	fmt.Printf("%s\n", content)
}
