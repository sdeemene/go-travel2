package api

import (
	"net/http"
)

func (api UserApi) UserRouters() {
	userRouter := api.router.PathPrefix("/api/v1/users").Subrouter()
	userRouter.HandleFunc("", api.GetUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", api.GetUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", api.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", api.DeleteUser).Methods(http.MethodDelete)
	userRouter.HandleFunc("/email/{email}", api.GetUserByEmail).Methods(http.MethodGet)
	userRouter.HandleFunc("/search", api.SearchUser).Methods(http.MethodGet)

}
