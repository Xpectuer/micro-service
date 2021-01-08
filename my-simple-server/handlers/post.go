package handlers

import (
	"my-simple-server/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProducts is a method to create new product resource
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	// reflect the object
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println(prod)
	// pass a REF
	data.AddProduct(prod)
}
