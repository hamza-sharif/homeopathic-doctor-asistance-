package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
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
	log().Debugf("Request: add new patient")

	filter := getPatientFilter(params)
	pts, err := c.rt.Svc.GetPatients(filter)
	if err != nil {
		log().Debugf("not able to get list of patients: %v", err)
		return gen.NewGetPatientsBadRequest().WithPayload("not able to update user")
	}
	return gen.NewGetPatientsOK().WithPayload(convertPatients(pts))
}

func getPatientFilter(p gen.GetPatientsParams) map[string]interface{} {
	filter := map[string]interface{}{}

	if p.Name != nil {
		filter["name"] = swag.StringValue(p.Name)
	}

	if p.MobileNo != nil {
		filter["mobile_no"] = swag.StringValue(p.MobileNo)
	}

	if p.FatherOrHusbandName != nil {
		filter["father_or_husband_name"] = swag.StringValue(p.FatherOrHusbandName)
	}

	if p.DateTime != nil {
		filter["created_at"] = p.DateTime
	}

	return filter
}
