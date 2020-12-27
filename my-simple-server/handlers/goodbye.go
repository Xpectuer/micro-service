package handlers

import (
	"log"
	"net/http"
)

// Goodbye is
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye return Goodbye Instance
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Byeee!"))
}
