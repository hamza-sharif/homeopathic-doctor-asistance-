package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"time"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	genmode "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func NewGetPatient(rt *runtime.Runtime) gen.GetPatientsHandler {
	return &getPatient{
		rt: rt,
	}
}

type getPatient struct {
	rt *runtime.Runtime
}

func (c *getPatient) Handle(params gen.GetPatientsParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: get patients patient")
	log().Debugf("Request: 111")

	filter, limit, offset := getPatientFilter(params)
	log().Debugf("Request: get patients patient")
	pts, size, err := c.rt.Svc.GetPatients(filter, limit, offset)
	if err != nil {
		log().Debugf("not able to get list of patients: %v", err)
		return gen.NewGetPatientsBadRequest().WithPayload("not able to update user")
	}
	return gen.NewGetPatientsOK().WithPayload(&genmode.PatientResponse{
		Size: int32(size),
		Data: convertPatients(pts),
	})
}

func getPatientFilter(p gen.GetPatientsParams) (models.Patient, int, int) {
	patient := models.Patient{}

	if p.Name != nil {
		patient.Name = swag.StringValue(p.Name)
	}

	if p.MobileNo != nil {
		patient.MobileNo = swag.StringValue(p.MobileNo)
	}

	if p.Cnic != nil {
		patient.Cnic = swag.StringValue(p.Cnic)
	}

	if p.FatherOrHusbandName != nil {
		patient.FatherOrHusbandName = swag.StringValue(p.FatherOrHusbandName)
	}

	if p.DateTime != nil {
		patient.CreatedAt = time.Time(*p.DateTime)
	}

	var limit, offset int
	if p.Limit != nil {
		limit = int(*p.Limit)
	}
	if p.Offset != nil {
		offset = int(*p.Offset)
	}

	return patient, limit, offset
}
