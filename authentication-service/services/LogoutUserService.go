package services

func (s *AuthService) LogoutUserService(userID int) (bool, string, error) {
	_, message, err := s.Repo.LogoutUserRepo(userID)

	if err != nil {
		return false, message, err
	}

	return true, message, nil
}
