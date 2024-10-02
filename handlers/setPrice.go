package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewSetPrice(rt *runtime.Runtime) gen.PutUpdatePriceHandler {
	return &putPrice{
		rt: rt,
	}
}

type putPrice struct {
	rt *runtime.Runtime
}

func (c *putPrice) Handle(params gen.PutUpdatePriceParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: user log in")

	err := c.rt.Svc.UpdatePrice(int(params.Body.Price))
	if err != nil {
		return gen.NewPostLoginBadRequest()
	}

	return gen.NewPutUpdatePriceOK()
}
