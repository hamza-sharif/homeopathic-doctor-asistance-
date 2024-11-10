package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewDeleteMedicine(rt *runtime.Runtime) gen.DeleteMedicinesMedicineIDHandler {
	return &deleteMedicine{
		rt: rt,
	}
}

type deleteMedicine struct {
	rt *runtime.Runtime
}

func (c *deleteMedicine) Handle(params gen.DeleteMedicinesMedicineIDParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: add new Medicine")

	//Medicine := createMedicine(params.Body)
	err := c.rt.Svc.DeleteMedicine(params.MedicineID)

	if err != nil {
		log().Debugf("not able to add Medicine: %v", err)
		return gen.NewDeleteMedicinesMedicineIDBadRequest().WithPayload("not able to add Medicine")
	}
	return gen.NewDeleteMedicinesMedicineIDOK().WithPayload("Medicine Deleted successfully")
}
