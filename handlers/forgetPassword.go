package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

func NewForgetPassword(rt *runtime.Runtime) gen.PostForgetPasswordHandler {
	return &forgetPass{
		rt: rt,
	}
}

type forgetPass struct {
	rt *runtime.Runtime
}

func (c *forgetPass) Handle(params gen.PostForgetPasswordParams) middleware.Responder {
	log().Debugf("Request: user Forget password")

	user, err := c.rt.Svc.GetUser(models.User{Username: params.Body.Username})
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPostForgetPasswordBadRequest().WithPayload("not able to update user password")
	}
	_, err = c.rt.Svc.UpdatePass(user, params.Body.Password)
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPostForgetPasswordBadRequest().WithPayload("not able to update user password")
	}
	return gen.NewPostForgetPasswordOK().WithPayload("Successfully updated user password")
}
