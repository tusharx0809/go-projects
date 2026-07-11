package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
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
}
