package services

import (
	"authentication-service/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(userID int, userUID string) (string, error) {

	var JWT_KEY string = os.Getenv("SECRET_KEY")

	claims := models.JWTClaims{
		UserID:  userID,
		UserUID: userUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
