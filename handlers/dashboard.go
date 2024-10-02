package handlers

import (
	"fmt"
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
		return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count %v" + err.Error())
	}
	monthlyPatients, err := c.rt.Svc.GetPatientsCount(startOfMonth, now)
	if err != nil {
		return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count %v" + err.Error())
	}
	var getPrice, getMonthlyPrice int
	fmt.Println("today ppppppppppppppppppppppppppppp", todayPatients)
	if todayPatients > 0 {
		getPrice, err = c.rt.Svc.GetPrice(startOfDay, now)
		fmt.Println("today cccccccccccccccccccccccccccccc", getPrice)
		if err != nil {
			return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count %v" + err.Error())
		}
	}

	fmt.Println("monthly ppppppppppppppppppppppppppppp", monthlyPatients)
	if monthlyPatients > 0 {
		getMonthlyPrice, err = c.rt.Svc.GetPrice(startOfMonth, now)
		fmt.Println("monthly ccccccccccccccccccccccccccccc", getMonthlyPrice)
		if err != nil {
			return gen.NewGetDashboardBadRequest().WithPayload("unable to get patients count %v" + err.Error())
		}
	}

	return gen.NewGetDashboardOK().WithPayload(&genModel.Dashboard{
		CostMonthly:     swag.Int32(int32(getMonthlyPrice)),
		CostToday:       swag.Int32(int32(getPrice)),
		PatientsMonthly: swag.Int32(int32(monthlyPatients)),
		PatientsPerDay:  swag.Int32(int32(todayPatients)),
	})
}
