package routes

import (
	"authentication-service/handlers"
	"authentication-service/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *handlers.AuthHandler) {
	mux.HandleFunc("POST /registerUser", handler.RegisterUserHandler)
	mux.HandleFunc("POST /login", handler.LoginUserHandler)

	mux.Handle("GET /profile",
		middleware.JWTMiddleware(http.HandlerFunc(handler.FetchProfileHandler)),
	)
	mux.Handle("POST /changePassword",
		middleware.JWTMiddleware(http.HandlerFunc(handler.ChangePasswordHandler)),
	)
}
