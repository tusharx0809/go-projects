package server

import (
	"authentication-service/handlers"
	"authentication-service/routes"
	"net/http"
)

func StartServer(handler *handlers.AuthHandler) *http.Server {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, handler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
