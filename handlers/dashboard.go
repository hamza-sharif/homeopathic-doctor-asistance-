package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	genModel "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
	"time"
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

	now := time.Now()
	// Start of the current day (00:00:00)
	startOfDay := now.Truncate(24 * time.Hour)

	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	todayPatients, err := c.rt.Svc.GetPatientsCount(startOfDay, now)
	if err != nil {
		return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count per Day %v" + err.Error())
	}
	monthlyPatients, err := c.rt.Svc.GetPatientsCount(startOfMonth, now)
	if err != nil {
		return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count per Month%v" + err.Error())
	}
	var getPrice, getMonthlyPrice int
	if todayPatients > 0 {
		getPrice, err = c.rt.Svc.GetBill(startOfDay, now)
		if err != nil {
			return gen.NewGetDashboardBadRequest().WithPayload("unable to get today revenue %v" + err.Error())
		}
	}

	if monthlyPatients > 0 {
		getMonthlyPrice, err = c.rt.Svc.GetBill(startOfMonth, now)
		if err != nil {
			return gen.NewGetDashboardBadRequest().WithPayload("unable to get total revenue  per Month %v" + err.Error())
		}
	}

	return gen.NewGetDashboardOK().WithPayload(&genModel.Dashboard{
		CostMonthly:     swag.Int32(int32(getMonthlyPrice)),
		CostToday:       swag.Int32(int32(getPrice)),
		PatientsMonthly: swag.Int32(int32(monthlyPatients)),
		PatientsPerDay:  swag.Int32(int32(todayPatients)),
		Fee:             swag.Int32(int32(c.rt.Svc.GetPrice())),
	})
}
