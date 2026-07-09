package models

type UserRegistrationRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
}
