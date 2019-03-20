package main

import (
	"log"
	"net/http"
  	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
    r.HandleFunc("/authorize", authorize).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}

func authorize(response http.ResponseWriter, request *http.Request) {
	log.Println(request)
}