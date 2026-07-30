package services

func (s *AuthService) GenerateNewAccessToken(userID int, refreshToken string) (string, error) {

	userUID, err := s.Repo.GetUserUID(userID, refreshToken)

	if err != nil {
		return "", err
	}

	newAccessToken, err := GenerateJWTAccessToken(userID, userUID)

	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
