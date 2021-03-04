/*
 * @Author: your name
 * @Date: 2020-12-28 21:46:12
 * @LastEditTime: 2021-01-20 20:16:53
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/get.go
 */
/*
 * @Author: your name
 * @Date: 2020-12-28 21:46:12
 * @LastEditTime: 2021-01-20 19:36:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /micro-service/my-simple-server/handlers/get.go
 */
package handlers

/*package is really important for swagger */
// GetProducts gets the products from list
import (
	"net/http"

	"github.com/Xpectuer/micro-service/my-simple-server/data"
)

// swagger:route GET /products products listProducts
// Return a list of Products
// responses:
//	200: productsResponse

// GetProducts return a list of products
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Debug("Handle GET Products")
	cur := h.URL.Query().Get("currency")
	rw.Header().Add("Content-Type", "application/json")

	prods, err := p.productDB.GetProducts(cur)
	if err != nil {
		p.l.Error("Unable to getProducts")
	}
	/**
	* Why not use encoder	?
	* The func Encode() writes the json string directly into
	* the stream
	* Encode() is faster than Marshal
	 */
	//d, err := json.Marshal(lp)
	err = data.ToJSON(prods, rw)
	if err != nil {
		p.l.Error("Uable to serialize product", "error", err)

	}
	//w.Write(d)

}

// swagger:route GET /products/{id} products ListSingleProduct
// Return a single product by specified ID
// responses:
//  200: productsResponse
//  404: errorResponse

// ListSingleProduct handles GET requests
func (p *Products) ListSingleProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Debug("get record id", id)

	cur := r.URL.Query().Get("currency")
	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Debug(" fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Debug(" fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Debug("erializing product", err)
	}
}
