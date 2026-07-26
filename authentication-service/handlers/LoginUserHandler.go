package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserLoginResponse{
				Success:      false,
				Message:      err.Error(),
				AccessToken:  "",
				RefreshToken: "",
			},
		)
		return
	}

	var accessToken string
	var refreshToken string

	_, accessToken, refreshToken, err = h.Service.LoginUserService(req.EmailOrUsername, req.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.UserLoginResponse{
				Success:      false,
				Message:      err.Error(),
				AccessToken:  "",
				RefreshToken: "",
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		models.UserLoginResponse{
			Success:      true,
			Message:      "Login Successful!",
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)

}
