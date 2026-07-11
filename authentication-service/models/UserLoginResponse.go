package models

type UserLoginResponse struct {
	Success bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
