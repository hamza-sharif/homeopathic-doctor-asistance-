package handlers

import (
	"github.com/go-openapi/loads"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

type Handler *operations.HomeopathicDoctorAssistantAPI

func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewHomeopathicDoctorAssistantAPI(spec)

	//user Authentication.
	handler.BearerAuth = func(token string) (interface{}, error) {
		// This function will be called with the object available in the context
		return BearerAuth(token, rt)
	}

	// user apis
	handler.PostLoginHandler = NewLoginUser(rt)
	handler.PutUpdatePasswordHandler = NewUpdatePassword(rt)
	handler.PostPatientsHandler = NewPostPatient(rt)
	handler.GetPatientsHandler = NewGetPatient(rt)
	handler.DeletePatientsPatientIDHandler = NewDeletePatient(rt)
	handler.GetMedicinesHandler = NewGetMedicine(rt)
	handler.GetDashboardHandler = NewGetDashboard(rt)
	handler.PutUpdatePriceHandler = NewSetPrice(rt)

	return handler
}
