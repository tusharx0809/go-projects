package models

type ProfileResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	UserEmail    string `json:"user_email"`
	UserFullName string `json:"full_name"`
	Dob          string `json:"dob"`
}
