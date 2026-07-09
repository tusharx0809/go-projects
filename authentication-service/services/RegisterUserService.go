package services

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *AuthService) RegisterUser(firstName string, lastName string, email string, password string, username string, dob time.Time) (bool, error) {
	var errResp string = ""
	switch {
	case len(firstName) <= 0:
		errResp = "First name cannot be empty!"
	case len(lastName) <= 0:
		errResp = "Last name cannot be empty!"
	case len(email) <= 0:
		errResp = "Email cannot be empty!"
	case len(password) <= 10:
		errResp = "Password must be atleast 10 characters long!"
	case len(username) <= 0:
		errResp = "Username cannot be empty"
	}

	if len(errResp) > 0 {
		return false, errors.New(errResp)
	}

	hashed_password, err := HashPassword(password)

	if err != nil {
		return false, err
	}
	_, err = s.Repo.RegisterUser(firstName, lastName, email, hashed_password, username, dob)

	if err != nil {
		return false, err
	}

	return true, nil

}
