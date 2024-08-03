package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/spf13/viper"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/config"
	gen "github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

func NewUpdatePassword(rt *runtime.Runtime) gen.PutUpdatePasswordHandler {
	return &updatePass{
		rt: rt,
	}
}

type updatePass struct {
	rt *runtime.Runtime
}

func (c *updatePass) Handle(params gen.PutUpdatePasswordParams) middleware.Responder {
	log().Debugf("Request: user log in")

	user, err := ValidateHeader(params.HTTPRequest.Header.Get("Authorization"), *c.rt)
	if err != nil {
		log().Errorf("failed to verify token [status 401] [error :%+v]", err)
		return gen.NewputUnauthorized()
		return genHydra.NewAcceptConsentChallengeUnauthorized().
			WithPayload(&genModel.Error{
				Code:    swag.String(strconv.Itoa(http.StatusUnauthorized)),
				Message: swag.String(viper.GetString(config.ErrorMessageToken401)),
			})
	}
	user, err := c.rt.Svc.UpdatePass(user)
}
