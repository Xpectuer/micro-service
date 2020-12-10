package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Goodbye is ...
type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// log in Hello Scope
	g.l.Println("Goodbye")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Goodbye %d", d)
}
