package handlers

import (
	"github.com/Xpectuer/mircro-service/my-simple-server/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products DeleteProduct
// Delete a product
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

//DeleteProduct handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	//id := getProductID(r)
	id := 0
	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		//data.ToJSON(rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		//data.ToJSON(rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
