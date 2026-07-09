package models

type UserRegistrationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
