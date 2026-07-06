package services

import (
	"authentication-service/repository"
)

type AuthService struct {
	Repo *repository.AuthRepo
}

func NewAuthService(
	repo *repository.AuthRepo,
) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}
