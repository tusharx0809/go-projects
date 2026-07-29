package models

type UserLogoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
