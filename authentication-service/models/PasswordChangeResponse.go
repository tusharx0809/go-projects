package models

type PasswordChangeResponse struct {
	Success  bool   `json:"success"`
	Messaage string `json:"message"`
}
