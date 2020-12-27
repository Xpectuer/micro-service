package handlers

import (
	"log"
	"my-simple-server/data"
	"net/http"
	"regexp"
	"strconv"
)

// Products is
type Products struct {
	l *log.Logger
}

// NewProducts is
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Convert list of Product into JSON
func (p *Products) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	if h.Method == http.MethodGet {
		p.getProducts(w, h)
		return
	}

	if h.Method == http.MethodPost {
		p.AddProduct(w, h)
		return
	}

	if h.Method == http.MethodPut {

		// expect ID from URI
		// use regex to extract path variables
		p.l.Println(h.URL.Path)
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(h.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Result group captured more than one!")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Unable to convert id into integer", http.StatusBadRequest)
			return
		}
		p.l.Println("got id", id)
		p.UpdateProduct(id, w, h)
		return

	}
	// handle and update
	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
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

// AddProduct is a method to create new product resource
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	// What if the body data is too Huge for Reader
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// UpdateProduct is a method let user to update the product with specified
func (p *Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

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
