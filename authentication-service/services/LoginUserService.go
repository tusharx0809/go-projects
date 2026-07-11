package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) LoginUser(emailOrUsername string, password string) (bool, string, error) {

	var passwordHash string
	var err error

	_, passwordHash, err = s.Repo.LoginUser(emailOrUsername)

	if err != nil {
		return false, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return false, "", errors.New("Incorrect credentials!")
	}

	return true, passwordHash, nil
}
