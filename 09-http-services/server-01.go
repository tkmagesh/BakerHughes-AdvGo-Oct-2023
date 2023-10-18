package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"price"`
}

var products []Product = []Product{
	{Id: 100, Name: "Pen", Cost: 10},
	{Id: 101, Name: "Pencil", Cost: 5},
}

type AppServer struct {
}

func (server *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		// fmt.Fprintln(w, "All the products will be served")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
