package repository

import (
	"context"
	"time"
)

func (r *AuthRepo) RegisterUser(firstName string, lastName string, email string, hashed_password string, username string, dob time.Time) (bool, error) {

	query := "SELECT func_register_user($1,$2,$3,$4,$5,$6)"

	_, err := r.Authdb.Exec(context.Background(), query, firstName, lastName, email, hashed_password, username, dob)

	if err != nil {
		return false, err
	}

	return true, nil

}
