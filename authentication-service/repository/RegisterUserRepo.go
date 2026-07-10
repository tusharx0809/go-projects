package repository

import (
	"context"
	"errors"
	"time"
)

func (r *AuthRepo) CheckEmail(email string) int {
	query := "SELECT 1 FROM users WHERE email=$1"
	var emailInt int
	err := r.Authdb.QueryRow(context.Background(), query, email).Scan(&emailInt)

	if err != nil {
		return -1
	}
	return emailInt
}

func (r *AuthRepo) CheckUsername(username string) int {
	query := "SELECT 1 FROM users WHERE user_name=$1"
	var usernameInt int
	err := r.Authdb.QueryRow(context.Background(), query, username).Scan(&usernameInt)

	if err != nil {
		return -1
	}
	return usernameInt
}

func (r *AuthRepo) RegisterUser(firstName string, lastName string, email string, hashed_password string, username string, dob time.Time) (bool, error) {
	query := "SELECT func_register_user($1,$2,$3,$4,$5,$6)"

	isEmailUnique := r.CheckEmail(email)

	if isEmailUnique == 1 {
		return false, errors.New("Email already registered! Please try another!")
	}

	isUsernameUnique := r.CheckUsername(username)

	if isUsernameUnique == 1 {
		return false, errors.New("Username not available, try another!")
	}

	_, err := r.Authdb.Exec(context.Background(), query, firstName, lastName, email, hashed_password, username, dob)

	if err != nil {
		return false, err
	}

	return true, nil

}
