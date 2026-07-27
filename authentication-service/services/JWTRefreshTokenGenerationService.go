package services

import (
	"authentication-service/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTRefreshToken(userID int) (string, error) {
	var REFRESH_JWT_KEY string = os.Getenv("REFRESH_SECRET_KEY")

	claims := models.JWTRefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(REFRESH_JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
