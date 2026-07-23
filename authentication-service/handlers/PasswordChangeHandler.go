package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req models.PasswordChangeRequest
	userID := r.Context().Value("userID").(int)
	userUID := r.Context().Value("userUID").(string)

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.PasswordChangeResponse{
				Success:  false,
				Messaage: err.Error(),
			},
		)
		return
	}

	_, err = h.Service.ChangePasswordService(req.OldPassword, req.NewPassword, userID, userUID)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(
			models.PasswordChangeResponse{
				Success:  false,
				Messaage: err.Error(),
			},
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		models.PasswordChangeResponse{
			Success:  false,
			Messaage: "Password Changed successfully!",
		},
	)

}
