package repository

import (
	"context"
)

func (r *AuthRepo) LogoutUserRepo(userID int) (bool, string, error) {
	query := "UPDATE refresh_tokens SET revoked = true, modified_at = NOW() WHERE user_id = $1"

	_, err := r.Authdb.Exec(context.Background(), query, userID)

	if err != nil {
		return false, "", err
	}

	return true, "User Logged out", nil
}
