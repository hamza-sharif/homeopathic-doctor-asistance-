package db

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

type Client interface {
	LoginUser(username, password string) (*models.User, error)
	UpdatePass(user *models.User) error
	UpdateToken(userID uint, newToken string) error
	Register(user models.User) error
	GetUserByToken(token string) (*models.User, error)
}
