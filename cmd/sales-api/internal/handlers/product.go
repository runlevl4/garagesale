package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/runlevl4/garagesale/internal/product"
)

// Product has handler methods for dealing with products
type Product struct {
	DB  *sqlx.DB
	Log *log.Logger
}

// ListProducts is a basic HTTP Handler.
func (p *Product) List(w http.ResponseWriter, r *http.Request) {

	list, err := product.List(p.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Printf("List : error retrieving products [%s]", err)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Printf("listProducts : error marshaling json : %s\n", err)
		return
	}

	// Note that header needs to be set before status is set.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.Log.Printf("listProducts : error writing response : %s\n", err)
	}
	p.Log.Printf("listProducts | %s | %d | %s\n", http.StatusText(http.StatusOK), http.StatusOK, string(data))
}
