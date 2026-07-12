package routes

import (
	"authentication-service/handlers"
	"authentication-service/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *handlers.AuthHandler) {
	mux.HandleFunc("POST /registerUser", handler.RegisterUser)
	mux.HandleFunc("POST /login", handler.LoginUser)

	mux.Handle("GET /profile",
		middleware.JWTMiddleware(http.HandlerFunc(handler.FetchProfile)),
	)
}
