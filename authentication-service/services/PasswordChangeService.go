package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) ChangePasswordService(oldPassword string, newPassword string, userID int, userUID string) (bool, error) {
	isFetched, oldPasswordHash := s.Repo.GetPasswordHash(userID, userUID)

	if !isFetched {
		return false, errors.New("Something went wrong!")
	}

	err := bcrypt.CompareHashAndPassword([]byte(oldPasswordHash), []byte(oldPassword))

	if err != nil {
		return false, errors.New("Old Password must be correct")
	}

	isPasswordValidated, err := validatePassword(newPassword)

	if !isPasswordValidated {
		return false, err
	}

	newPasswordHash, err := HashPassword(newPassword)

	if err != nil {
		return false, err
	}

	_, err = s.Repo.ChangePasswordRepo(newPasswordHash, userID, userUID)

	if err != nil {
		return false, err
	}

	return true, nil

}
