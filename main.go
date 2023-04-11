package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server")

	// Create an instance of a HTTP handler that will have the functions
	// to handle the HTTP requests
	handler := new(HTTPHandler)

	// Here we use a router to send the requests to the functions from
	// HTTP handler that will handle those requests
	router := mux.NewRouter()

	// Print information about requests
	router.Use(loggingMiddleware)

	// Register a simple /api/hello endpoint, which just says hello
	router.HandleFunc("/api/hello", handler.Get).Methods(http.MethodGet)

	// Notice that in the code of GetWithParams we can read the value of `name`
	// using mux.Vars(). If you call `/api/hello/esdras` it will use `esdras` in its response.
	router.HandleFunc("/api/hello/{name}", handler.GetWithParams).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
