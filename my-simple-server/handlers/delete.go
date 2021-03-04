package handlers

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

import (
	"net/http"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
)

// swagger:route DELETE /products/{id} products DeleteProduct
// Delete a product
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

//DeleteProduct handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Add("Content-Type", "application/json")
	id := getProductID(r)
	p.l.Debug("deleting record", "id", id)

	err := p.productDB.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Debug("deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Error("deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
