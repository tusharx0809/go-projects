package repository

import "context"

func (r *AuthRepo) LoginUserByEmail(email string) (bool, string) {

	var emailInt int
	emailInt = r.CheckEmail(email)

	if emailInt == -1 {
		return false, "Incorrect credentials!"
	}

	var passwordHash string
	query := "SELECT hashed_password FROM users where email = $1"
	err := r.Authdb.QueryRow(context.Background(), query, email).Scan(&passwordHash)

	if err != nil {
		return false, ""
	}

	return true, passwordHash
}

func (r *AuthRepo) LoginUserByUsername(username string) (bool, string) {
	var usernameInt int
	usernameInt = r.CheckUsername(username)

	if usernameInt == -1 {
		return false, "Incorrect credentials!"
	}

	var passwordHash string

	query := "SELECT hashed_password FROM users where user_name = $1"
	err := r.Authdb.QueryRow(context.Background(), query, username).Scan(&passwordHash)

	if err != nil {
		return false, ""
	}

	return true, passwordHash
}
