package models

type UserLoginRequest struct {
	EmailOrUsername string `json:"email_or_username"`
	Password        string `json:"password"`
}
