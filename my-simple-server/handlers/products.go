package handlers

/*
 * @Author: your name
 * @Date: 2020-12-16 00:13:09
 * @LastEditTime: 2021-01-26 17:44:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/products.go
 */

import (
	"net/http"
	"strconv"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

// Products is just products
type Products struct {
	l         hclog.Logger
	v         *data.Validation
	productDB *data.ProductsDB
}

// KeyProduct Used as a key in context
type KeyProduct struct{}

// NewProducts is
func NewProducts(l hclog.Logger, v *data.Validation, productDB *data.ProductsDB) *Products {
	return &Products{l, v, productDB}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
