package repository

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepo struct {
	Authdb *pgxpool.Pool
}

func NewAuthRepo(
	authdb *pgxpool.Pool,
) *AuthRepo {
	return &AuthRepo{
		Authdb: authdb,
	}
}
