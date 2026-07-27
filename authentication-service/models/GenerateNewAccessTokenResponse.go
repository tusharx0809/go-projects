package models

type GenerateNewAccessTokenResponse struct {
	Success        bool   `json:"success"`
	Message        string `json:"message"`
	NewAccessToken string `json:"new_access_token"`
}
