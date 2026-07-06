package models

import "time"

type UserRegistrationRequest struct {
	Fname    string    `json:"first_name"`
	Lname    string    `json:"last_name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Username string    `json:"username"`
	Dob      time.Time `json:"date_of_birth"`
}
