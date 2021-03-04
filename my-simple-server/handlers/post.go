/*
 * @Author: your name
 * @Date: 2020-12-28 21:46:23
 * @LastEditTime: 2021-01-26 17:45:15
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/post.go
 */
package handlers

import (
	"net/http"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
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
	p.l.Info("Handling POST Products")
	// reflect the object
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Debug("product accepted is", prod)
	// pass a REF
	p.productDB.AddProduct(*prod)
}
