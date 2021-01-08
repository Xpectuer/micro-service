package handlers

/*package is really important for swagger */
// GetProducts gets the products from list
import (
	"github.com/Xpectuer/micro-service/my-simple-server/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of Products
// responses:
//	200: productsResponse

// GetProducts return a list of products
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("[DEBUG]Handle GET Products")

	rw.Header().Add("Content-Type", "application/json")

	lp := data.GetProducts()
	/**
	* Why not use encoder	?
	* The func Encode() writes the json string directly into
	* the stream
	* Encode() is faster than Marshal
	 */
	//d, err := json.Marshal(lp)
	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json ", http.StatusInternalServerError)
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

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
