package routes

import (
	"authentication-service/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *handlers.AuthHandler) {
	mux.HandleFunc("POST /registerUser", handler.RegisterUser)
	mux.HandleFunc("POST /login", handler.LoginUser)
}
