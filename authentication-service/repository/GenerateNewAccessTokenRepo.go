package repository

import (
	"context"
	"errors"
)

func (r *AuthRepo) GetUserUID(userID int) (string, error) {
	var userUID string

	var isRevoked bool
	checkQuery := "SELECT revoked FROM refresh_tokens WHERE user_id = $1"

	err := r.Authdb.QueryRow(context.Background(), checkQuery, userID).Scan(&isRevoked)

	if isRevoked {
		return "", errors.New("Please login Again!")
	}

	query := "SELECT user_uid FROM users WHERE user_id = $1"
	err = r.Authdb.QueryRow(context.Background(), query, userID).Scan(&userUID)

	if err != nil {
		return "", err
	}

	return userUID, nil
}
