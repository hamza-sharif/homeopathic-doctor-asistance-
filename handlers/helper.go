package handlers

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	genModel "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func convertUser(user *models.User) *genModel.User {
	return &genModel.User{
		Username: user.Username,
		Password: user.Password,
		Token:    user.Token,
	}
}

func createPatient(params *genModel.Patient) *models.Patient {
	return &models.Patient{
		ID:                  uuid.NewString(),
		Name:                params.Name,
		Address:             params.Address,
		Age:                 int(params.Age),
		Cnic:                params.Cnic,
		FatherOrHusbandName: params.FatherOrHusbandName,
		Gender:              params.Gender,
		MobileNo:            params.MobileNo,
		Weight:              int(params.Weight),
		Medicine:            params.Medicine,
		Disease:             params.Disease,
	}
}

func convertPatients(patients []*models.Patient) []*genModel.Patient {
	var ptsList []*genModel.Patient
	for _, pts := range patients {
		ptsList = append(ptsList, &genModel.Patient{
			ID:                  pts.ID,
			Name:                pts.Name,
			Address:             pts.Address,
			Age:                 int64(pts.Age),
			Cnic:                pts.Cnic,
			FatherOrHusbandName: pts.FatherOrHusbandName,
			Gender:              pts.Gender,
			MobileNo:            pts.MobileNo,
			Weight:              int64(pts.Weight),
			DateTime:            strfmt.DateTime(pts.CreatedAt),
			Medicine:            pts.Medicine,
			Disease:             pts.Disease,
		})
	}

	return ptsList
}

func convertMedicine(patients []*models.Medicine) []*genModel.Medicine {
	var medList []*genModel.Medicine
	for _, meds := range patients {
		medList = append(medList, &genModel.Medicine{
			Name:        meds.Name,
			Description: meds.Description,
			ID:          int32(meds.ID),
		})
	}

	return medList
}
