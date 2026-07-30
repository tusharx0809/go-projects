package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func (h *AuthHandler) GenerateNewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req models.GenerateNewAccessTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.GenerateNewAccessTokenResponse{
				Success:        false,
				Message:        err.Error(),
				NewAccessToken: "",
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

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.GenerateNewAccessTokenResponse{
				Success:        false,
				Message:        "Please login again!",
				NewAccessToken: "",
			},
		)
		return
	}

	userID := claims.UserID
	newAccesstoken, err := h.Service.GenerateNewAccessToken(userID, req.RefreshToken)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.GenerateNewAccessTokenResponse{
				Success:        false,
				Message:        err.Error(),
				NewAccessToken: "",
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		models.GenerateNewAccessTokenResponse{
			Success:        true,
			Message:        "token generated",
			NewAccessToken: newAccesstoken,
		},
	)
}
