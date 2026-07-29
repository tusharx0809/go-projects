package models

type UserLogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}
