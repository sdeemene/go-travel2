package api

import (
	"net/http"
)

func (api AddressApi) AddressRouters() {
	addressRoute := api.router.PathPrefix("/api/v1/addresses").Subrouter()
	addressRoute.HandleFunc("", api.CreateAddress).Methods(http.MethodPost)
	addressRoute.HandleFunc("", api.GetAddresses).Methods(http.MethodGet)
	addressRoute.HandleFunc("/search", api.SearchAddress).Methods(http.MethodGet)
	addressRoute.HandleFunc("/{id}", api.GetAddress).Methods(http.MethodGet)
	addressRoute.HandleFunc("/{id}", api.UpdateAddress).Methods(http.MethodPut)
	addressRoute.HandleFunc("/{id}", api.DeleteAddress).Methods(http.MethodDelete)
}
