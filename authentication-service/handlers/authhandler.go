package handlers

import (
	"authentication-service/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(
	service *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}
