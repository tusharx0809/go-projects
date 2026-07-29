package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func (h *AuthHandler) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserLogoutRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserLogoutResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}
	claims := &models.JWTRefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(
		req.RefreshToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_SECRET_KEY")), nil
		},
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.UserLogoutResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.UserLogoutResponse{
				Success: false,
				Message: "Already Logged out",
			},
		)
		return
	}

	userID := claims.UserID

	_, message, err := h.Service.LogoutUserService(userID)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(
			models.UserLogoutResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		models.UserLogoutResponse{
			Success: true,
			Message: message,
		},
	)

}
