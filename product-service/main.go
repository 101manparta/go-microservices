package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	products = make(map[string]Product)
	mu       sync.RWMutex
)

func main() {
	http.HandleFunc("/products", productsHandler)
	log.Println("Product service running at :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()

		var list []Product
		for _, p := range products {
			list = append(list, p)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(list); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodPost:
		var p Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		mu.Lock()
		products[p.ID] = p
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
