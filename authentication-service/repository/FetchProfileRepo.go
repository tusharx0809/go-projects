package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *AuthRepo) FetchProfileRepo(userID int, userUID string) (string, string, string, error) {
	parsedUserUID, err := uuid.Parse(userUID)

	if err != nil {
		return "", "", "", err
	}

	query := "SELECT * FROM func_fetch_profile($1,$2)"
	var (
		userEmail    string
		userFullName string
		userDob      string
	)

	err = r.Authdb.QueryRow(context.Background(), query, userID, parsedUserUID).Scan(&userEmail, &userFullName, &userDob)

	if err != nil {
		return "", "", "", err
	}

	return userEmail, userFullName, userDob, nil

}
