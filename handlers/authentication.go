package handlers

import (
	"net/http"

	"github.com/go-openapi/errors"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
)

// BearerAuth applies when the "Authorization" header is set.
func BearerAuth(token string, rt *runtime.Runtime) (interface{}, error) {
	log().WithField("Authorization", token).Info("authenticating request")

	user, err := rt.Svc.ValidateToken(token)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, "authentication failed: invalid Authorization token")
	}

	return user, nil
}
