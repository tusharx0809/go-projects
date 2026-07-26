package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) LoginUserService(emailOrUsername string, password string) (bool, string, string, error) {

	var passwordHash string
	var err error

	_, passwordHash, err = s.Repo.LoginUserRepo(emailOrUsername)

	if err != nil {
		return false, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return false, "", "", errors.New("Incorrect credentials!")
	}

	userID, userUID, err := s.Repo.FetchClaims(emailOrUsername)

	if err != nil {
		return false, "", "", err
	}

	accessTokenString, err := GenerateJWTAccessToken(userID, userUID)
	refreshTokenString, err := GenerateJWTRefreshToken(userID)

	if err != nil {
		return false, "", "", err
	}

	err = s.Repo.AddRefreshTokenInDB(userID, refreshTokenString)

	if err != nil {
		return false, "", "", err
	}

	return true, accessTokenString, refreshTokenString, nil
}
