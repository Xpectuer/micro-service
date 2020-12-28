package handlers

import (
	"context"
	"fmt"
	"log"
	"my-simple-server/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is
type Products struct {
	l *log.Logger
}

// NewProducts is
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts gets the products from list
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()
	/**
	* Why not use encoder	?
	* The func Encode() writes the json string directly into
	* the stream
	* Encode() is faster than Marshal
	 */
	//d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json ", http.StatusInternalServerError)
	}
	//w.Write(d)

}

// AddProducts is a method to create new product resource
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// What if the body data is too Huge for Reader
	// prod := &data.Product{}
	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	// 	return
	// }
	// p.l.Printf("Prod: %#v", prod)

	// reflect the object
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println(prod)
	// pass a REF
	data.AddProduct(prod)
}

// UpdateProducts is a method let user to update the product with specified
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")
	vars := mux.Vars(r)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		http.Error(rw, "Unable to Convert Id", http.StatusBadRequest)
	}

	// What if the body data is too Huge for Reader
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	e := data.UpdateProduct(id, prod)
	if e == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if e != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

// KeyProduct Used as a key in context
type KeyProduct struct{}

// MiddlewareProductValidation will execute
// before the acutual handler called
// Just like Spring-AOP
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// What if the body data is too Huge for Reader
		p.l.Printf("Validating the Product Data")
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] Deserializing product", err)
			http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
			return
		}

		p.l.Println(prod)

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Validating Product", err)
			http.Error(
				rw,
				fmt.Sprintf("Invalid JSON Object: %s", err),
				http.StatusBadRequest)
			return
		}
		// reflect
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		// update the request
		r = r.WithContext(ctx)
		// Call the next hanlder, which can be another middleware the chain
		next.ServeHTTP(rw, r)

	})
}
