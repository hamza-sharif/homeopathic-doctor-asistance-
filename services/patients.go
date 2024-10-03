package services

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
	"time"
)

func (s *Service) AddPatients(patient *models.Patient) error {
	return s.db.AddPatient(patient)
}

func (s *Service) GetPatients(patient models.Patient, limit, offset int) ([]*models.Patient, int, error) {
	patients, err := s.db.GetPatient(&patient, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	pCount, err := s.db.GetPatientWithFilterCount(&patient)
	if err != nil {
		return nil, 0, err
	}
	return patients, pCount, nil
}

func (s *Service) GetPatientsCount(startDate, endDate time.Time) (int, error) {
	return s.db.GetPatientCount(startDate, endDate)
}
func (s *Service) GetBill(startDate, endDate time.Time) (int, error) {
	return s.db.GetBill(startDate, endDate)
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

func (s *Service) UpdatePrice(price int) error {
	return s.db.SetPrice(price)
}

func (s *Service) GetPrice() int {
	return s.db.GetPrice()
}
