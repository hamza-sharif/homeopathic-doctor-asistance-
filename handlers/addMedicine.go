package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func NewAddMedicine(rt *runtime.Runtime) gen.PostMedicinesHandler {
	return &addMedicine{
		rt: rt,
	}
}

type addMedicine struct {
	rt *runtime.Runtime
}

func (c *addMedicine) Handle(params gen.PostMedicinesParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: add new patient")

	err := c.rt.Svc.AddMedicine(&models.Medicine{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		log().Debugf("not able to get list of patients: %v", err)
		return gen.NewPostMedicinesBadRequest().WithPayload("not able to update user")
	}
	return gen.NewPostMedicinesCreated().WithPayload("medicine added successfully")
}
