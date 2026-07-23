package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *AuthRepo) GetPasswordHash(userID int, userUID string) (bool, string) {
	parsedUID, err := uuid.Parse(userUID)

	if err != nil {
		return false, ""
	}
	var hashed_password string

	query := "SELECT hashed_password FROM users WHERE user_id = $1 AND user_uid = $2"

	err = r.Authdb.QueryRow(context.Background(), query, userID, parsedUID).Scan(&hashed_password)

	if err != nil {
		return false, ""
	}

	return true, hashed_password
}

func (r *AuthRepo) ChangePasswordRepo(newPasswordHash string, userID int, userUID string) (bool, error) {
	parsedUID, err := uuid.Parse(userUID)

	if err != nil {
		return false, err
	}

	query := "SELECT func_change_password($1,$2,$3)"

	_, err = r.Authdb.Exec(context.Background(), query, newPasswordHash, userID, parsedUID)

	if err != nil {
		return false, err
	}

	return true, nil
}
