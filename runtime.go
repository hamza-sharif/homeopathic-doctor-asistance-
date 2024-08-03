package homeopathic_doctor_assistant

import (
	wraperrors "github.com/pkg/errors"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/db/postgres"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/services"
)

// Runtime initializes values for entry point to our application.
type Runtime struct {
	Svc *services.Service
}

func NewRuntime() (*Runtime, error) {
	db, err := postgres.NewClient()
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect with database")
	}

	return &Runtime{Svc: services.NewService(db)}, nil
}
