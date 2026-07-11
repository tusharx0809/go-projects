package repository

import (
	"context"
	"errors"
)

func (r *AuthRepo) LoginUser(username string) (bool, string, error) {
	var passwordHash string

	query := "SELECT hashed_password FROM users where user_name = $1 OR email = $1"
	err := r.Authdb.QueryRow(context.Background(), query, username).Scan(&passwordHash)

	if err != nil {
		return false, "", errors.New("Incorrect credentials!")
	}

	return true, passwordHash, nil
}
