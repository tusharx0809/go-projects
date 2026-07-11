package repository

import (
	"context"
	"errors"
)

func (r *AuthRepo) LoginUser(usernameOrEmail string) (bool, string, error) {
	var passwordHash string

	query := "SELECT hashed_password FROM users where user_name = $1 OR email = $1"
	err := r.Authdb.QueryRow(context.Background(), query, usernameOrEmail).Scan(&passwordHash)

	if err != nil {
		return false, "", errors.New("Incorrect credentials!")
	}

	return true, passwordHash, nil
}

func (r *AuthRepo) FetchClaims(usernameOrEmail string) (int, string, error) {
	var userID int
	var userUID string

	query := "SELECT user_id, user_uid FROM users WHERE user_name = $1 OR email = $1"
	err := r.Authdb.QueryRow(context.Background(), query, usernameOrEmail).Scan(&userID, &userUID)

	if err != nil {
		return -1, "", err
	}

	return userID, userUID, nil
}
