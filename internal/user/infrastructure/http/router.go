package http

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// endpoints
	mux.Handle("/users", &userHandler{})
	// mux.Handle("/products", &productHandler{})

	return mux
}
