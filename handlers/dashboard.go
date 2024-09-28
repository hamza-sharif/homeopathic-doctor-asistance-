package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	genModel "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewGetDashboard(rt *runtime.Runtime) gen.GetDashboardHandler {
	return &getDashboard{
		rt: rt,
	}
}

type getDashboard struct {
	rt *runtime.Runtime
}

func (c *getDashboard) Handle(params gen.GetDashboardParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: get dashboard")

	return gen.NewGetDashboardOK().WithPayload(&genModel.Dashboard{
		CostMonthly:     1,
		CostToday:       2,
		PatientsMonthly: 3,
		PatientsPerDay:  4,
	})
}
