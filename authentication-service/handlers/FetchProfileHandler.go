package handlers

import (
	"authentication-service/models"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) FetchProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	userUID := r.Context().Value("userUID").(string)

	userEmail, userFullName, userDob, err := h.Service.FetchProfileService(userID, userUID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			models.ProfileResponse{
				UserEmail:    "",
				UserFullName: "",
				Dob:          "",
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		models.ProfileResponse{
			UserEmail:    userEmail,
			UserFullName: userFullName,
			Dob:          userDob,
		},
	)
}
