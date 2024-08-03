// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name HomeopathicDoctorAssistant --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.HomeopathicDoctorAssistantAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HomeopathicDoctorAssistantAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GetDiseasesHandler == nil {
		api.GetDiseasesHandler = operations.GetDiseasesHandlerFunc(func(params operations.GetDiseasesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetDiseases has not yet been implemented")
		})
	}
	if api.GetMedicinesHandler == nil {
		api.GetMedicinesHandler = operations.GetMedicinesHandlerFunc(func(params operations.GetMedicinesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetMedicines has not yet been implemented")
		})
	}
	if api.GetPatientsHandler == nil {
		api.GetPatientsHandler = operations.GetPatientsHandlerFunc(func(params operations.GetPatientsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetPatients has not yet been implemented")
		})
	}
	if api.PostDiseasesHandler == nil {
		api.PostDiseasesHandler = operations.PostDiseasesHandlerFunc(func(params operations.PostDiseasesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostDiseases has not yet been implemented")
		})
	}
	if api.PostForgetPasswordHandler == nil {
		api.PostForgetPasswordHandler = operations.PostForgetPasswordHandlerFunc(func(params operations.PostForgetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostForgetPassword has not yet been implemented")
		})
	}
	if api.PostLoginHandler == nil {
		api.PostLoginHandler = operations.PostLoginHandlerFunc(func(params operations.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostLogin has not yet been implemented")
		})
	}
	if api.PostMedicinesHandler == nil {
		api.PostMedicinesHandler = operations.PostMedicinesHandlerFunc(func(params operations.PostMedicinesParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostMedicines has not yet been implemented")
		})
	}
	if api.PostPatientsHandler == nil {
		api.PostPatientsHandler = operations.PostPatientsHandlerFunc(func(params operations.PostPatientsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostPatients has not yet been implemented")
		})
	}
	if api.PutUpdatePasswordHandler == nil {
		api.PutUpdatePasswordHandler = operations.PutUpdatePasswordHandlerFunc(func(params operations.PutUpdatePasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutUpdatePassword has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
