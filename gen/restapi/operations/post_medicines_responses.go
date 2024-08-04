// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMedicinesCreatedCode is the HTTP code returned for type PostMedicinesCreated
const PostMedicinesCreatedCode int = 201

/*
PostMedicinesCreated Medicine added successfully

swagger:response postMedicinesCreated
*/
type PostMedicinesCreated struct {
}

// NewPostMedicinesCreated creates PostMedicinesCreated with default headers values
func NewPostMedicinesCreated() *PostMedicinesCreated {

	return &PostMedicinesCreated{}
}

// WriteResponse to the client
func (o *PostMedicinesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostMedicinesBadRequestCode is the HTTP code returned for type PostMedicinesBadRequest
const PostMedicinesBadRequestCode int = 400

/*
PostMedicinesBadRequest bad request

swagger:response postMedicinesBadRequest
*/
type PostMedicinesBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostMedicinesBadRequest creates PostMedicinesBadRequest with default headers values
func NewPostMedicinesBadRequest() *PostMedicinesBadRequest {

	return &PostMedicinesBadRequest{}
}

// WithPayload adds the payload to the post medicines bad request response
func (o *PostMedicinesBadRequest) WithPayload(payload interface{}) *PostMedicinesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post medicines bad request response
func (o *PostMedicinesBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMedicinesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostMedicinesUnauthorizedCode is the HTTP code returned for type PostMedicinesUnauthorized
const PostMedicinesUnauthorizedCode int = 401

/*
PostMedicinesUnauthorized internal server error

swagger:response postMedicinesUnauthorized
*/
type PostMedicinesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostMedicinesUnauthorized creates PostMedicinesUnauthorized with default headers values
func NewPostMedicinesUnauthorized() *PostMedicinesUnauthorized {

	return &PostMedicinesUnauthorized{}
}

// WithPayload adds the payload to the post medicines unauthorized response
func (o *PostMedicinesUnauthorized) WithPayload(payload interface{}) *PostMedicinesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post medicines unauthorized response
func (o *PostMedicinesUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMedicinesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}