package api

import (
	"net/http"
)

func (api UserApi) UserAuthRouters() {
	subRouter := api.router.PathPrefix("/api/v1/auth").Subrouter()
	subRouter.HandleFunc("/login", api.AuthenticateUser).Methods(http.MethodPost)
	subRouter.HandleFunc("/register", api.CreateUser).Methods(http.MethodPost)
	// subRouter.HandleFunc("/logout", "notImplemented").Methods(http.MethodPut)
	// subRouter.HandleFunc("/refresh", "notImplemented").Methods(http.MethodPut)
	// subRouter.HandleFunc("/password_request/{email}", "notImplemented").Methods(http.MethodDelete)
	// subRouter.HandleFunc("/validate_otp/{code}", "notImplemented").Methods(http.MethodDelete)
	// subRouter.HandleFunc("/check_email/{email}", "notImplemented").Methods(http.MethodDelete)
}
