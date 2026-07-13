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
				Success: false,
				Message: err.Error(),
				Token:   "",
			},
		)
		return
	}

	var passwordHash string
	_, passwordHash, err = h.Service.LoginUserService(req.EmailOrUsername, req.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.UserLoginResponse{
				Success: false,
				Message: err.Error(),
				Token:   "",
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		models.UserLoginResponse{
			Success: true,
			Message: "Login Successful!",
			Token:   passwordHash,
		},
	)

}
