package handlers

import (
	"github.com/go-openapi/loads"
	runtime "github.com/hamza-sharif/homeopathic-doctor-assistant"
	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

type Handler *operations.HomeopathicDoctorAssistantAPI

func NewHandler(runtime *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewHomeopathicDoctorAssistantAPI(spec)

	//user apis
	handler.PostLoginHandler = NewLoginUser(runtime)
	//handler.BearerAuth = BearerAuth   it`s not required in this application

	return handler
}
