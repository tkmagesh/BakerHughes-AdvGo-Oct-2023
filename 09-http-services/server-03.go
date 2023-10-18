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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func newProductHandler(w http.ResponseWriter, r *http.Request) {
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

func getCustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

type AppServer struct {
	routes map[string]http.HandlerFunc
}

func (server *AppServer) Register(method, path string, handler func(http.ResponseWriter, *http.Request)) {
	slug := fmt.Sprintf("%s-%s", method, path)
	server.routes[slug] = handler
}

func (server *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := fmt.Sprintf("%s-%s", r.Method, r.URL.Path)
	if handler, exists := server.routes[slug]; exists {
		handler(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]http.HandlerFunc),
	}
}

// logger middleware
func loggerMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	appServer := NewAppServer()
	appServer.Register(http.MethodGet, "/", loggerMiddleware(indexHandler))
	appServer.Register(http.MethodGet, "/products", loggerMiddleware(getProductsHandler))
	appServer.Register(http.MethodPost, "/products", loggerMiddleware(newProductHandler))
	appServer.Register(http.MethodGet, "/customers", loggerMiddleware(getCustomersHandler))
	http.ListenAndServe(":8080", appServer)
}
