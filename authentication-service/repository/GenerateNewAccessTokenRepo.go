package repository

import (
	"context"
	"errors"
)

func (r *AuthRepo) GetUserUID(userID int, refreshToken string) (string, error) {
	var userUID string
	var refreshTokenDB string

	refreshTokenQuery := "SELECT refresh_token FROM refresh_tokens WHERE refresh_token = $1"

	err := r.Authdb.QueryRow(context.Background(), refreshTokenQuery, refreshToken).Scan(&refreshTokenDB)

	if err != nil {
		return "", err
	}

	var isRevoked bool
	checkQuery := "SELECT revoked FROM refresh_tokens WHERE user_id = $1"

	err = r.Authdb.QueryRow(context.Background(), checkQuery, userID).Scan(&isRevoked)

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
