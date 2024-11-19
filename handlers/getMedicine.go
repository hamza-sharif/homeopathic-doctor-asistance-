package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	genModel "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
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

	var limit, offset int
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	if limit == 0 {
		limit = 5000
	}

	count, meds, err := c.rt.Svc.GetMedicine(swag.StringValue(params.Name), limit, offset)
	if err != nil {
		log().Debugf("not able to get list of medicines: %v", err)
		return gen.NewGetMedicinesBadRequest().WithPayload("not able to get medicines")
	}

	return gen.NewGetMedicinesOK().WithPayload(&genModel.MedicineResponse{
		Size: swag.Int64(count),
		Data: convertMedicine(meds),
	})
}
