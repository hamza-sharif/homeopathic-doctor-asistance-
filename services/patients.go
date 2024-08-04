package services

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func (s *Service) AddPatients(patient *models.Patient) error {
	return s.db.AddPatient(patient)
}

func (s *Service) GetPatients(filter map[string]interface{}) ([]*models.Patient, error) {
	return s.db.GetPatient(filter)
}

func (s *Service) GetMedicine(name string) ([]*models.Medicine, error) {
	if name == "" {
		return s.db.GetAllMedicine()
	}
	return s.db.GetMedicineByName(name)
}

func (s *Service) AddMedicine(medicine *models.Medicine) error {
	return s.db.AddMedicine(medicine)
}
