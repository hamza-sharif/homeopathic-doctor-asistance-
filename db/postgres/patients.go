package postgres

import (
	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
	wraperrors "github.com/pkg/errors"
)

func (cli *client) AddPatient(patient *models.Patient) error {
	return cli.conn.Create(&patient).Error
}

func (cli *client) GetPatient(filter map[string]interface{}) ([]*models.Patient, error) {
	var patients []*models.Patient
	result := cli.conn.Model(&models.Patient{}).Where(filter).Find(patients)

	if result.Error != nil {
		return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}

	return patients, nil
}

func (cli *client) GetMedicineByName(filter string) ([]*models.Medicine, error) {
	var meds []*models.Medicine
	filter = "%" + filter + "%"
	result := cli.conn.Model(&models.Medicine{}).Where("name LIKE ?", filter).Limit(5).Find(meds)
	if result.Error != nil {
		return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}
	return meds, nil
}

func (cli *client) GetAllMedicine() ([]*models.Medicine, error) {
	var meds []*models.Medicine
	result := cli.conn.Model(&models.Medicine{}).Find(meds)
	if result.Error != nil {
		return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}
	return meds, nil
}

func (cli *client) AddMedicine(medicine *models.Medicine) error {
	return cli.conn.Create(medicine).Error
}
