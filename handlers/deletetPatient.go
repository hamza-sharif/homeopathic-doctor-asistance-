package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewDeletePatient(rt *runtime.Runtime) gen.DeletePatientsPatientIDHandler {
	return &deletePatient{
		rt: rt,
	}
}

type deletePatient struct {
	rt *runtime.Runtime
}

func (c *deletePatient) Handle(params gen.DeletePatientsPatientIDParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: add new patient")

	//patient := createPatient(params.Body)
	err := c.rt.Svc.DeletePatient(params.PatientID)

	if err != nil {
		log().Debugf("not able to add patient: %v", err)
		return gen.NewPostPatientsBadRequest().WithPayload("not able to add patient")
	}
	return gen.NewPostPatientsCreated().WithPayload("patient added successfully")
}
