package repository

import (
	"context"
	"errors"
	"time"
)

func (r *AuthRepo) CreateDBStructure() error {
	query :=
		`CREATE TABLE IF NOT EXISTS users (
		user_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_uid UUID DEFAULT gen_random_uuid() NOT NULL UNIQUE,
		email VARCHAR(200) NOT NULL UNIQUE,
		user_name VARCHAR(100) UNIQUE NOT NULL,
		hashed_password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		modified_at TIMESTAMP WITH TIME ZONE
	);
	`

	_, err := r.Authdb.Exec(context.Background(), query)

	if err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS profiles (
		profile_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_ID INT REFERENCES users(user_id),
		f_name VARCHAR(100),
		l_name VARCHAR(100),
		date_of_birth DATE,
		is_active BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		modified_at TIMESTAMP WITH TIME ZONE
	);
	`

	_, err = r.Authdb.Exec(context.Background(), query)

	if err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS refresh_tokens(
		refresh_token_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_ID INT REFERENCES users(user_id),
		refresh_token TEXT NOT NULL,
		expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
		revoked BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		modified_at TIMESTAMP WITH TIME ZONE
	);
	`

	_, err = r.Authdb.Exec(context.Background(), query)

	if err != nil {
		return err
	}

	query = `
			-- FUNCTION: public.func_add_refresh_token(integer, text)

			-- DROP FUNCTION IF EXISTS public.func_add_refresh_token(integer, text);

			CREATE OR REPLACE FUNCTION public.func_add_refresh_token(
				in_user_id integer,
				in_token text)
				RETURNS void
				LANGUAGE 'plpgsql'
				COST 100
				VOLATILE PARALLEL UNSAFE
			AS $BODY$
			DECLARE
				prev_refresh_token_id INT;
			BEGIN
				SELECT refresh_token_id INTO prev_refresh_token_id FROM refresh_tokens WHERE user_id = in_user_id;
				IF prev_refresh_token_id IS NOT NULL THEN
					UPDATE refresh_tokens
					SET refresh_token = in_token,
						expires_at = NOW() + INTERVAL '7 days',
						revoked = false
					WHERE refresh_token_id = prev_refresh_token_id;
				ELSE
					INSERT INTO refresh_tokens(user_id, refresh_token, expires_at, revoked, created_at)
					VALUES(in_user_id, in_token, NOW()+INTERVAL '7 days', false, NOW());

				END IF;
			END;
			$BODY$;

			ALTER FUNCTION public.func_add_refresh_token(integer, text)
				OWNER TO postgres;

		CREATE OR REPLACE FUNCTION public.func_change_password(
			new_password_hash text,
			in_user_id integer,
			in_user_uid text)
			RETURNS void
			LANGUAGE 'plpgsql'
			COST 100
			VOLATILE PARALLEL UNSAFE
		AS $BODY$
		BEGIN
			UPDATE users
			SET hashed_password = new_password_hash,
			modified_at = NOW()
			WHERE user_id = in_user_id AND user_uid = in_user_uid::uuid;
		END;
		$BODY$;

		ALTER FUNCTION public.func_change_password(text, integer, text)
			OWNER TO postgres;	

		CREATE OR REPLACE FUNCTION public.func_fetch_profile(
		in_user_id integer,
		in_user_uid text)
		RETURNS TABLE(user_email character varying, full_name text, dob text) 
		LANGUAGE 'plpgsql'
		COST 100
		VOLATILE PARALLEL UNSAFE
		ROWS 1000

	AS $BODY$
	BEGIN
		RETURN QUERY
		SELECT
			u.email AS user_email,
			CONCAT_WS(' ',p.f_name, p.l_name) AS full_name,
			date_of_birth::text AS dob
		FROM users u
		JOIN profiles p ON u.user_id = p.user_id
		WHERE u.user_uid = in_user_uid::uuid AND p.user_id = in_user_id;
	END;
	$BODY$;

	ALTER FUNCTION public.func_fetch_profile(integer, text)
		OWNER TO postgres;

		CREATE OR REPLACE FUNCTION public.func_register_user(
		in_f_name character varying,
		in_l_name character varying,
		in_email character varying,
		in_hashed_password character varying,
		in_username character varying,
		in_dob date)
		RETURNS void
		LANGUAGE 'plpgsql'
		COST 100
		VOLATILE PARALLEL UNSAFE
	AS $BODY$
		DECLARE
				new_user_id INT;
		BEGIN
			INSERT INTO users(email,user_name,hashed_password)
			VALUES(in_email,in_username,in_hashed_password)
			RETURNING user_id INTO new_user_id;

			INSERT INTO profiles(user_id,f_name,l_name,date_of_birth)
			VALUES(new_user_id,in_f_name,in_l_name,in_dob);
		END;
	$BODY$;

	ALTER FUNCTION public.func_register_user(character varying, character varying, character varying, character varying, character varying, date)
		OWNER TO postgres;
	`

	_, err = r.Authdb.Exec(context.Background(), query)

	if err != nil {
		return err
	}

	return nil
}

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

func (r *AuthRepo) RegisterUserRepo(firstName string, lastName string, email string, hashed_password string, username string, dob time.Time) (bool, error) {

	err := r.CreateDBStructure()

	if err != nil {
		return false, err
	}

	query := "SELECT func_register_user($1,$2,$3,$4,$5,$6)"

	isEmailUnique := r.CheckEmail(email)

	if isEmailUnique == 1 {
		return false, errors.New("Email already registered! Please try another!")
	}

	isUsernameUnique := r.CheckUsername(username)

	if isUsernameUnique == 1 {
		return false, errors.New("Username not available, try another!")
	}

	_, err = r.Authdb.Exec(context.Background(), query, firstName, lastName, email, hashed_password, username, dob)

	if err != nil {
		return false, err
	}

	return true, nil

}
