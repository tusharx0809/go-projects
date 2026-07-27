package repository

import "context"

func (r *AuthRepo) GetUserUID(userID int) (string, error) {
	var userUID string
	query := "SELECT user_uid FROM users WHERE user_id = $1"
	err := r.Authdb.QueryRow(context.Background(), query, userID).Scan(&userUID)

	if err != nil {
		return "", err
	}

	return userUID, nil
}
