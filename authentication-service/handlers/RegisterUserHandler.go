package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
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

	var errResp string = ""
	switch {
	case len(req.FirstName) <= 0:
		errResp = "First name cannot be empty!"
	case len(req.LastName) <= 0:
		errResp = "Last name cannot be empty!"
	case len(req.Email) <= 0:
		errResp = "Email cannot be empty!"
	case len(req.Password) <= 10:
		errResp = "Password must be atleast 10 characters long!"
	case len(req.Username) <= 0:
		errResp = "Username cannot be empty"
	case len(req.DateOfBirth.GoString()) <= 0:
		errResp = "Date Of Birth cannot be empty!"
	}

	if len(errResp) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.UserRegistrationResponse{
				Success: false,
				Message: errResp,
			},
		)
		return
	}

	//to be continued

}
