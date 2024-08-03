package services

import "github.com/hamza-sharif/homeopathic-doctor-assistant/models"

func (s *Service) ValidateToken(token string) (*models.User, error) {
	return s.db.GetUserByToken(token)
}
