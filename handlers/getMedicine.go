package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewGetMedicine(rt *runtime.Runtime) gen.GetMedicinesHandler {
	return &getMedicine{
		rt: rt,
	}
}

type getMedicine struct {
	rt *runtime.Runtime
}

func (c *getMedicine) Handle(params gen.GetMedicinesParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: get Medicines")

	meds, err := c.rt.Svc.GetMedicine(swag.StringValue(params.Name))
	if err != nil {
		log().Debugf("not able to get list of medicines: %v", err)
		return gen.NewGetMedicinesBadRequest().WithPayload("not able to get medicines")
	}
	return gen.NewGetMedicinesOK().WithPayload(convertMedicine(meds))
}
