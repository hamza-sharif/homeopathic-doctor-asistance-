package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
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

	filter := map[string]interface{}{
		"username": params.Body.Username,
	}
	user, err := c.rt.Svc.GetUser(filter)
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPostForgetPasswordBadRequest().WithPayload("not able to update user")
	}
	_, err = c.rt.Svc.UpdatePass(user, params.Body.Password)
	if err != nil {
		log().Debugf("not able to update password")
		return gen.NewPostForgetPasswordBadRequest().WithPayload("not able to update user")
	}
	return gen.NewPostForgetPasswordOK()
}
