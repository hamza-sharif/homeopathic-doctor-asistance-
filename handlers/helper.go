package handlers

import (
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func convertUser(user models.User) *gen.User {
	return &gen.User{
		Username: user.Username,
		Password: user.Password,
		Token:    user.Token,
	}
}
