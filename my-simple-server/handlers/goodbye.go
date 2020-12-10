package handlers

import (
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
	rw.Write([]byte("Byeee!"))
}
