package api

import "net/http"

func (api TravelApi) TravelRouters() {
	travelRouter := api.router.PathPrefix("/api/v1/travels").Subrouter()
	travelRouter.HandleFunc("", api.GetTravels).Methods(http.MethodGet)
	travelRouter.HandleFunc("", api.CreateTravel).Methods(http.MethodPost)
	travelRouter.HandleFunc("/search", api.SearchTravel).Methods(http.MethodGet)
	travelRouter.HandleFunc("/{id}", api.GetTravel).Methods(http.MethodGet)
	travelRouter.HandleFunc("/{id}", api.UpdateTravel).Methods(http.MethodPut)
	travelRouter.HandleFunc("/{id}", api.DeleteTravel).Methods(http.MethodDelete)
	travelRouter.HandleFunc("/user/{id}", api.GetUserTravels).Methods(http.MethodGet)
	travelRouter.HandleFunc("/rate/{id}", api.RateTravel).Methods(http.MethodPatch)
}
