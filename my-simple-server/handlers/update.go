package handlers

import (
	"github.com/Xpectuer/micro-service/my-simple-server/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateProducts is a method let user to update the product with specified
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	// What if the body data is too Huge for Reader
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err := data.FromJSON(prod, r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	e := data.UpdateProduct(prod)
	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
