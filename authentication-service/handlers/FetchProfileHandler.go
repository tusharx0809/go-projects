package handlers

import (
	"net/http"
)

func (h *AuthHandler) FetchProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
	userUID := r.Context().Value("userUID").(string)

	// fmt.Println(userID)
	// fmt.Println(userUID)
}
