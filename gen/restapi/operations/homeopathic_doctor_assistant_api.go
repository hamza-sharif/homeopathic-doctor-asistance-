// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewHomeopathicDoctorAssistantAPI creates a new HomeopathicDoctorAssistant instance
func NewHomeopathicDoctorAssistantAPI(spec *loads.Document) *HomeopathicDoctorAssistantAPI {
	return &HomeopathicDoctorAssistantAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		GetDashboardHandler: GetDashboardHandlerFunc(func(params GetDashboardParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetDashboard has not yet been implemented")
		}),
		GetDiseasesHandler: GetDiseasesHandlerFunc(func(params GetDiseasesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetDiseases has not yet been implemented")
		}),
		GetMedicinesHandler: GetMedicinesHandlerFunc(func(params GetMedicinesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetMedicines has not yet been implemented")
		}),
		GetPatientsHandler: GetPatientsHandlerFunc(func(params GetPatientsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetPatients has not yet been implemented")
		}),
		PostDiseasesHandler: PostDiseasesHandlerFunc(func(params PostDiseasesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostDiseases has not yet been implemented")
		}),
		PostForgetPasswordHandler: PostForgetPasswordHandlerFunc(func(params PostForgetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation PostForgetPassword has not yet been implemented")
		}),
		PostLoginHandler: PostLoginHandlerFunc(func(params PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation PostLogin has not yet been implemented")
		}),
		PostMedicinesHandler: PostMedicinesHandlerFunc(func(params PostMedicinesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostMedicines has not yet been implemented")
		}),
		PostPatientsHandler: PostPatientsHandlerFunc(func(params PostPatientsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostPatients has not yet been implemented")
		}),
		PutUpdatePasswordHandler: PutUpdatePasswordHandlerFunc(func(params PutUpdatePasswordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PutUpdatePassword has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*HomeopathicDoctorAssistantAPI API for managing users, patients, medicines, and diseases */
type HomeopathicDoctorAssistantAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GetDashboardHandler sets the operation handler for the get dashboard operation
	GetDashboardHandler GetDashboardHandler
	// GetDiseasesHandler sets the operation handler for the get diseases operation
	GetDiseasesHandler GetDiseasesHandler
	// GetMedicinesHandler sets the operation handler for the get medicines operation
	GetMedicinesHandler GetMedicinesHandler
	// GetPatientsHandler sets the operation handler for the get patients operation
	GetPatientsHandler GetPatientsHandler
	// PostDiseasesHandler sets the operation handler for the post diseases operation
	PostDiseasesHandler PostDiseasesHandler
	// PostForgetPasswordHandler sets the operation handler for the post forget password operation
	PostForgetPasswordHandler PostForgetPasswordHandler
	// PostLoginHandler sets the operation handler for the post login operation
	PostLoginHandler PostLoginHandler
	// PostMedicinesHandler sets the operation handler for the post medicines operation
	PostMedicinesHandler PostMedicinesHandler
	// PostPatientsHandler sets the operation handler for the post patients operation
	PostPatientsHandler PostPatientsHandler
	// PutUpdatePasswordHandler sets the operation handler for the put update password operation
	PutUpdatePasswordHandler PutUpdatePasswordHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *HomeopathicDoctorAssistantAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *HomeopathicDoctorAssistantAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *HomeopathicDoctorAssistantAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *HomeopathicDoctorAssistantAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *HomeopathicDoctorAssistantAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *HomeopathicDoctorAssistantAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *HomeopathicDoctorAssistantAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *HomeopathicDoctorAssistantAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *HomeopathicDoctorAssistantAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the HomeopathicDoctorAssistantAPI
func (o *HomeopathicDoctorAssistantAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.GetDashboardHandler == nil {
		unregistered = append(unregistered, "GetDashboardHandler")
	}
	if o.GetDiseasesHandler == nil {
		unregistered = append(unregistered, "GetDiseasesHandler")
	}
	if o.GetMedicinesHandler == nil {
		unregistered = append(unregistered, "GetMedicinesHandler")
	}
	if o.GetPatientsHandler == nil {
		unregistered = append(unregistered, "GetPatientsHandler")
	}
	if o.PostDiseasesHandler == nil {
		unregistered = append(unregistered, "PostDiseasesHandler")
	}
	if o.PostForgetPasswordHandler == nil {
		unregistered = append(unregistered, "PostForgetPasswordHandler")
	}
	if o.PostLoginHandler == nil {
		unregistered = append(unregistered, "PostLoginHandler")
	}
	if o.PostMedicinesHandler == nil {
		unregistered = append(unregistered, "PostMedicinesHandler")
	}
	if o.PostPatientsHandler == nil {
		unregistered = append(unregistered, "PostPatientsHandler")
	}
	if o.PutUpdatePasswordHandler == nil {
		unregistered = append(unregistered, "PutUpdatePasswordHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *HomeopathicDoctorAssistantAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *HomeopathicDoctorAssistantAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *HomeopathicDoctorAssistantAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *HomeopathicDoctorAssistantAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *HomeopathicDoctorAssistantAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *HomeopathicDoctorAssistantAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the homeopathic doctor assistant API
func (o *HomeopathicDoctorAssistantAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *HomeopathicDoctorAssistantAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/dashboard"] = NewGetDashboard(o.context, o.GetDashboardHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/diseases"] = NewGetDiseases(o.context, o.GetDiseasesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/medicines"] = NewGetMedicines(o.context, o.GetMedicinesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/patients"] = NewGetPatients(o.context, o.GetPatientsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/diseases"] = NewPostDiseases(o.context, o.PostDiseasesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/forget-password"] = NewPostForgetPassword(o.context, o.PostForgetPasswordHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewPostLogin(o.context, o.PostLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/medicines"] = NewPostMedicines(o.context, o.PostMedicinesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/patients"] = NewPostPatients(o.context, o.PostPatientsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/update-password"] = NewPutUpdatePassword(o.context, o.PutUpdatePasswordHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *HomeopathicDoctorAssistantAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *HomeopathicDoctorAssistantAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *HomeopathicDoctorAssistantAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *HomeopathicDoctorAssistantAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *HomeopathicDoctorAssistantAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
