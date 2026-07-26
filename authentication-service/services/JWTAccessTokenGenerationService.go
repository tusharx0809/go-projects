package services

import (
	"authentication-service/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTAccessToken(userID int, userUID string) (string, error) {

	var ACCESS_JWT_KEY string = os.Getenv("ACCESS_SECRET_KEY")

	claims := models.JWTAccessTokenClaims{
		UserID:  userID,
		UserUID: userUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ACCESS_JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
