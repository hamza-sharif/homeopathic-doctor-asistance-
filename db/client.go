package db

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
	"time"
)

type Client interface {
	GetUser(user *models.User) (*models.User, error)
	LoginUser(username, password string) (*models.User, error)
	UpdateUser(user *models.User) error
	UpdateToken(userID uint, newToken string) error
	Register(user *models.User) error
	GetUserByToken(token string) (*models.User, error)

	AddPatient(patient *models.Patient) error
	GetPatient(patient *models.Patient, limit, offset int) ([]*models.Patient, error)
	GetPatientWithFilterCount(patient *models.Patient) (int, error)
	GetPatientCount(startDate, endDate time.Time) (int, error)
	GetPrice(startDate, endDate time.Time) (int, error)
	SetPrice(price int) error

	GetAllMedicine() ([]*models.Medicine, error)
	GetMedicineByName(filter string) ([]*models.Medicine, error)
	AddMedicine(*models.Medicine) error
}
