package services

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func (s *Service) GetUserByUsername(username, password string) (models.User, error) {
	return s.db.LoginUser(username, password)
}

func (s *Service) UpdatePass(user models.User, pass string) (models.User, error) {
	user.Password = pass
	err := s.db.UpdateUser(user)
	return user, err
}

func (s *Service) RegisterUser(user models.User) error {
	return s.db.Register(user)
}
