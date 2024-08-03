package handlers

import (
	"net/http"

	"github.com/go-openapi/errors"
	wraperrors "github.com/pkg/errors"

	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/models"
)

// BearerAuth applies when the "Authorization" header is set.
func BearerAuth(token string) (interface{}, error) {
	log().WithField("Authorization", token).Info("authenticating request")
	var err error
	// call a function to validate the toke, but it should be an external function.
	if err != nil {
		log().Error("authentication failed: invalid Authorization token")

		return nil, errors.New(http.StatusUnauthorized, "authentication failed: invalid Authorization token")
	}

	return nil, nil
}

// ValidateHeader check the validation of the JWT.
func ValidateHeader(encodedToken string, rt runtime.Runtime) (*models.User, error) {
	user, err := rt.Svc.ValidateToken(encodedToken)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to verify the signature of the token")
	}

	return user, nil
}
