package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewLoginUser(rt *runtime.Runtime) gen.PostLoginHandler {
	return &loginUser{
		rt: rt,
	}
}

type loginUser struct {
	rt *runtime.Runtime
}

func (c *loginUser) Handle(params gen.PostLoginParams) middleware.Responder {
	log().Debugf("Request: user log in")

	user, err := c.rt.Svc.LoginUser(params.Body.Username, params.Body.Password)
	if err != nil {
		if err.Error() == config.ErrorMessage404 {
			return gen.NewPostLoginNotFound().WithPayload("User not found")
		}
		if err.Error() == config.ErrorMessage500 {
			return gen.NewPostLoginBadRequest().WithPayload("Internal Server Error")
		}
		if err.Error() == config.ErrorMessageToken401 {
			return gen.NewPostLoginUnauthorized().WithPayload("Token Not Authorized")
		}
		return gen.NewPostLoginBadRequest().WithPayload("Bad request")
	}

	return gen.NewPostLoginOK().WithPayload(convertUser(user))
}
