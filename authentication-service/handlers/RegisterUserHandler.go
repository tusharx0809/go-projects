package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
	"time"
)

func (h *AuthHandler) RegisterUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req models.UserRegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserRegistrationResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}
	parsedDob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserRegistrationResponse{
				Success: false,
				Message: "Invalid date_of_birth format. Use YYYY-MM-DD",
			},
		)
		return
	}
	_, err = h.Service.RegisterUser(req.FirstName, req.LastName, req.Email, req.Password, req.Username, parsedDob)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserRegistrationResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		models.UserRegistrationResponse{
			Success: true,
			Message: "User Registered successfully!",
		},
	)

}
