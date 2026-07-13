package models

type ProfileResponse struct {
	UserEmail    string `json:"user_email"`
	UserFullName string `json:"full_name"`
	Dob          string `json:"dob"`
}
