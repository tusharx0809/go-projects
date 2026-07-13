package services

func (s *AuthService) FetchProfileService(userID int, userUID string) (string, string, string, error) {
	userEmail, userFullName, userDob, err := s.Repo.FetchProfileRepo(userID, userUID)

	if err != nil {
		return "", "", "", err
	}

	return userEmail, userFullName, userDob, nil
}
