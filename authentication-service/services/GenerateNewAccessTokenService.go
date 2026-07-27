package services

func (s *AuthService) GenerateNewAccessToken(userID int) (string, error) {

	userUID, err := s.Repo.GetUserUID(userID)

	if err != nil {
		return "", err
	}

	newAccessToken, err := GenerateJWTAccessToken(userID, userUID)

	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
