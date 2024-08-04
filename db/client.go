package db

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

type Client interface {
	GetUser(filter map[string]interface{}) (*models.User, error)
	LoginUser(username, password string) (*models.User, error)
	UpdateUser(user *models.User) error
	UpdateToken(userID uint, newToken string) error
	Register(user models.User) error
	GetUserByToken(token string) (*models.User, error)

	AddPatient(patient *models.Patient) error
	GetPatient(filter map[string]interface{}) ([]*models.Patient, error)

	GetAllMedicine() ([]*models.Medicine, error)
	GetMedicineByName(filter string) ([]*models.Medicine, error)
	AddMedicine(*models.Medicine) error
}
