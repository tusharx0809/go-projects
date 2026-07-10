package services

import (
	"errors"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func validatePassword(password string) (bool, error) {
	if len(password) < 10 {
		return false, errors.New("Password must be at least 10 characters long!")
	}

	var (
		hasUpper   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasDigit || !hasSpecial {
		return false, errors.New("Password must contain at least one uppercase letter, one digit, and one special character!")
	}

	return true, nil

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
	case len(username) <= 0:
		errResp = "Username cannot be empty"
	}

	if len(errResp) > 0 {
		return false, errors.New(errResp)
	}

	isPasswordValidated, err := validatePassword(password)

	if !isPasswordValidated {
		return false, err
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
