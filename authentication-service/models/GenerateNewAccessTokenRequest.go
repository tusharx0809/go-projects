package models

type GenerateNewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
