package api

import "net/http"

func (api PlaceApi) PlaceRouters() {
	placeRouter := api.router.PathPrefix("/api/v1/places").Subrouter()
	placeRouter.HandleFunc("", api.GetPlaces).Methods(http.MethodGet)
	placeRouter.HandleFunc("", api.CreatePlace).Methods(http.MethodPost)
	placeRouter.HandleFunc("/search", api.SearchPlace).Methods(http.MethodGet)
	placeRouter.HandleFunc("/{id}", api.GetPlace).Methods(http.MethodGet)
	placeRouter.HandleFunc("/{id}", api.UpdatePlace).Methods(http.MethodPut)
	placeRouter.HandleFunc("/{id}", api.DeletePlace).Methods(http.MethodDelete)
}
