package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	api "github.com/0sektor0/go-dtm/api"
)

type Server struct {
	_apiClient *api.ApiClient
}

func NewServer() (*Server, error) {
	server := &Server{}

	apiClient, err := api.NewApiClient()
	if err != nil {
		return nil, err
	}

	server._apiClient = apiClient
	return server, nil
}

func (this *Server) Start() {
	settings, _ := api.GetSettings()

	r := mux.NewRouter()
	r.HandleFunc("/auth", this.authorize).Methods("GET")

	log.Printf("starting on port %s", settings.Port)
	log.Fatal(http.ListenAndServe(settings.Port, r))
}

func (this *Server) authorize(response http.ResponseWriter, request *http.Request) {
	log.Println(request)
	this._apiClient.Authorize()
}
