package postgres

import (
	"time"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
	wraperrors "github.com/pkg/errors"
)

func (cli *client) AddPatient(patient *models.Patient) error {

	patient.Price = cli.Fee

	return cli.conn.Create(&patient).Error
}

func (cli *client) GetPatient(patient *models.Patient, limit, offset int) ([]*models.Patient, error) {
	var patients []*models.Patient
	query := cli.conn.Model(&models.Patient{}).Order("created_at desc")

	if patient != nil {
		name := patient.Name
		if patient.Name != "" {
			query = query.Where("name ILIKE ?", "%"+patient.Name+"%")
			patient.Name = ""
		}
		query = query.Where(patient)
		patient.Name = name
	}
	//result := cli.conn.Model(&models.Patient{}).Where(patient).Order("created_at desc").
	//	Limit(limit).Offset(offset).Find(&patients)

	result := query.Limit(limit).Offset(offset).Find(&patients)

	if result.Error != nil {
		return nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}
	return patients, nil
}

func (cli *client) GetPatientWithFilterCount(patient *models.Patient) (int, error) {
	var count int64
	query := cli.conn.Model(&models.Patient{})

	if patient != nil {
		if patient.Name != "" {
			query = query.Where("name ILIKE ?", "%"+patient.Name+"%")
			patient.Name = ""
		}
		query = query.Where(patient)
	}
	result := query.Count(&count)

	if result.Error != nil {
		return 0, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}

	return int(count), nil
}
func (cli *client) DeletePatient(patientID string) error {
	result := cli.conn.Delete(&models.Patient{ID: patientID})
	return result.Error
}

func (cli *client) GetPatientCount(startDate, endDate time.Time) (int, error) {
	var count int64
	result := cli.conn.Model(&models.Patient{}).Where("created_at BETWEEN ? AND ?", startDate, endDate).Count(&count)

	if result.Error != nil {
		return 0, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}

	return int(count), nil
}

func (cli *client) GetBill(startDate, endDate time.Time) (int, error) {
	var count int64
	err := cli.conn.Model(&models.Patient{}).Select("SUM(price)").Where("created_at BETWEEN ? AND ?", startDate, endDate).Scan(&count).Error
	if err != nil {
		return 0, wraperrors.Wrap(err, config.ErrorMessage500)
	}

	return int(count), nil
}

func (cli *client) SetPrice(price int) error {
	cli.Fee = price
	err := cli.conn.Model(&models.Price{}).Where("id", 1).Updates(&models.Price{Fee: price}).Error
	if err != nil {
		wraperrors.Wrap(err, config.ErrorMessage500)
	}

	return nil
}

func (cli *client) GetPrice() int {

	return cli.Fee
}

func (cli *client) GetMedicine(filter string, limit, offset int) (int64, []*models.Medicine, error) {
	var meds []*models.Medicine
	var count int64
	query := cli.conn.Model(&models.Medicine{})
	if filter != "" {
		filter = "%" + filter + "%"
		query = query.Where("name ILIKE ?", filter)
	}

	resultCount := query.Count(&count)

	if resultCount.Error != nil {
		return 0, nil, wraperrors.Wrap(resultCount.Error, config.ErrorMessage500)
	}

	result := query.Limit(limit).Offset(offset).Find(&meds)
	if result.Error != nil {
		return 0, nil, wraperrors.Wrap(result.Error, config.ErrorMessage500)
	}

	return count, meds, nil
}

func (cli *client) AddMedicine(medicine *models.Medicine) error {
	return cli.conn.Create(&medicine).Error
}

func (cli *client) DeleteMedicine(medicineID int32) error {
	md := &models.Medicine{}
	md.ID = uint(medicineID)

	result := cli.conn.Delete(md)
	return result.Error
}
