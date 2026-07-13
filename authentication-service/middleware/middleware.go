package middleware

import (
	"authentication-service/models"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}
		// fmt.Println(authHeader)

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authotization header", http.StatusUnauthorized)
			return
		}

		claims := &models.JWTClaims{}
		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("SECRET_KEY")), nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "Your session has expired, please login again!", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			"userID",
			claims.UserID,
		)

		ctx = context.WithValue(
			ctx,
			"userUID",
			claims.UserUID,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
