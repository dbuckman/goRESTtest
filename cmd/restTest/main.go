package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type JSONResponse struct {
	Method  string
	Headers http.Header
}

func main() {
	// Init the mux router
	router := mux.NewRouter()
	// Route handles & endpoints
	router.HandleFunc("/delete", Delete).Methods("DELETE")
	router.HandleFunc("/get", Get).Methods("GET")
	// fmt.Println("Hello, playground")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Get(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	method := r.Method

	response := JSONResponse{Method: method, Headers: headers}
	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	method := r.Method

	response := JSONResponse{Method: method, Headers: headers}
	json.NewEncoder(w).Encode(response)
}
