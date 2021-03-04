package handlers

/*
 * @Author: your name
 * @Date: 2020-12-28 21:46:29
 * @LastEditTime: 2021-01-26 17:44:53
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/update.go
 */

import (
	"net/http"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
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
	p.l.Info("Handling PUT Products")

	// What if the body data is too Huge for Reader
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	err := data.FromJSON(prod, r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	}
	p.l.Debug("Prod: ", prod)
	e := p.productDB.UpdateProduct(*prod)
	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
