package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewPostPatient(rt *runtime.Runtime) gen.PostPatientsHandler {
	return &postPatient{
		rt: rt,
	}
}

type postPatient struct {
	rt *runtime.Runtime
}

func (c *postPatient) Handle(params gen.PostPatientsParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: add new patient")

	patient := createPatient(params.Body)
	err := c.rt.Svc.AddPatients(patient)
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPostPatientsBadRequest().WithPayload("not able to update user")
	}
	return gen.NewPostPatientsCreated()
}
