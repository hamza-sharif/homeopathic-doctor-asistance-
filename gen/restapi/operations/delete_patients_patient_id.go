// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeletePatientsPatientIDHandlerFunc turns a function with the right signature into a delete patients patient ID handler
type DeletePatientsPatientIDHandlerFunc func(DeletePatientsPatientIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePatientsPatientIDHandlerFunc) Handle(params DeletePatientsPatientIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePatientsPatientIDHandler interface for that can handle valid delete patients patient ID params
type DeletePatientsPatientIDHandler interface {
	Handle(DeletePatientsPatientIDParams, interface{}) middleware.Responder
}

// NewDeletePatientsPatientID creates a new http.Handler for the delete patients patient ID operation
func NewDeletePatientsPatientID(ctx *middleware.Context, handler DeletePatientsPatientIDHandler) *DeletePatientsPatientID {
	return &DeletePatientsPatientID{Context: ctx, Handler: handler}
}

/*
	DeletePatientsPatientID swagger:route DELETE /patients/{patient_id} deletePatientsPatientId

Delete patient
*/
type DeletePatientsPatientID struct {
	Context *middleware.Context
	Handler DeletePatientsPatientIDHandler
}

func (o *DeletePatientsPatientID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeletePatientsPatientIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
