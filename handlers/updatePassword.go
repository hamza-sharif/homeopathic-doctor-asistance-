package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
	model "github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func NewUpdatePassword(rt *runtime.Runtime) gen.PutUpdatePasswordHandler {
	return &updatePass{
		rt: rt,
	}
}

type updatePass struct {
	rt *runtime.Runtime
}

func (c *updatePass) Handle(params gen.PutUpdatePasswordParams, principal interface{}) middleware.Responder {
	log().Debugf("Request: user update password")
	user, ok := principal.(*model.User)
	if !ok {
		log().Debugf("not able to convert %v to User", principal)
		return gen.NewPostLoginBadRequest().WithPayload("not able to get user  from token")
	}
	_, err := c.rt.Svc.UpdatePass(user, params.Body.NewPassword)
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPutUpdatePasswordBadRequest().WithPayload("not able to update user")
	}
	return gen.NewPutUpdatePasswordOK().WithPayload("password updated successfully")
}
