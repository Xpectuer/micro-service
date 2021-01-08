package main

import (
	"fmt"
	"my-simple-server/client/client"
	"my-simple-server/client/client/products"
	"testing"
)

func TestMain(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9091")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()

	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	//c.Products.ListProducts(params)
}
