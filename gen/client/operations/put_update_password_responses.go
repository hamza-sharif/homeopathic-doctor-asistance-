// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutUpdatePasswordReader is a Reader for the PutUpdatePassword structure.
type PutUpdatePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUpdatePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUpdatePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUpdatePasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUpdatePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /update-password] PutUpdatePassword", response, response.Code())
	}
}

// NewPutUpdatePasswordOK creates a PutUpdatePasswordOK with default headers values
func NewPutUpdatePasswordOK() *PutUpdatePasswordOK {
	return &PutUpdatePasswordOK{}
}

/*
PutUpdatePasswordOK describes a response with status code 200, with default header values.

Password updated successfully
*/
type PutUpdatePasswordOK struct {
	Payload string
}

// IsSuccess returns true when this put update password o k response has a 2xx status code
func (o *PutUpdatePasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put update password o k response has a 3xx status code
func (o *PutUpdatePasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put update password o k response has a 4xx status code
func (o *PutUpdatePasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put update password o k response has a 5xx status code
func (o *PutUpdatePasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put update password o k response a status code equal to that given
func (o *PutUpdatePasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put update password o k response
func (o *PutUpdatePasswordOK) Code() int {
	return 200
}

func (o *PutUpdatePasswordOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordOK %s", 200, payload)
}

func (o *PutUpdatePasswordOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordOK %s", 200, payload)
}

func (o *PutUpdatePasswordOK) GetPayload() string {
	return o.Payload
}

func (o *PutUpdatePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUpdatePasswordBadRequest creates a PutUpdatePasswordBadRequest with default headers values
func NewPutUpdatePasswordBadRequest() *PutUpdatePasswordBadRequest {
	return &PutUpdatePasswordBadRequest{}
}

/*
PutUpdatePasswordBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PutUpdatePasswordBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this put update password bad request response has a 2xx status code
func (o *PutUpdatePasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put update password bad request response has a 3xx status code
func (o *PutUpdatePasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put update password bad request response has a 4xx status code
func (o *PutUpdatePasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put update password bad request response has a 5xx status code
func (o *PutUpdatePasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put update password bad request response a status code equal to that given
func (o *PutUpdatePasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put update password bad request response
func (o *PutUpdatePasswordBadRequest) Code() int {
	return 400
}

func (o *PutUpdatePasswordBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordBadRequest %s", 400, payload)
}

func (o *PutUpdatePasswordBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordBadRequest %s", 400, payload)
}

func (o *PutUpdatePasswordBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PutUpdatePasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUpdatePasswordUnauthorized creates a PutUpdatePasswordUnauthorized with default headers values
func NewPutUpdatePasswordUnauthorized() *PutUpdatePasswordUnauthorized {
	return &PutUpdatePasswordUnauthorized{}
}

/*
PutUpdatePasswordUnauthorized describes a response with status code 401, with default header values.

internal server error
*/
type PutUpdatePasswordUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this put update password unauthorized response has a 2xx status code
func (o *PutUpdatePasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put update password unauthorized response has a 3xx status code
func (o *PutUpdatePasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put update password unauthorized response has a 4xx status code
func (o *PutUpdatePasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put update password unauthorized response has a 5xx status code
func (o *PutUpdatePasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put update password unauthorized response a status code equal to that given
func (o *PutUpdatePasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put update password unauthorized response
func (o *PutUpdatePasswordUnauthorized) Code() int {
	return 401
}

func (o *PutUpdatePasswordUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordUnauthorized %s", 401, payload)
}

func (o *PutUpdatePasswordUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /update-password][%d] putUpdatePasswordUnauthorized %s", 401, payload)
}

func (o *PutUpdatePasswordUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PutUpdatePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutUpdatePasswordBody put update password body
swagger:model PutUpdatePasswordBody
*/
type PutUpdatePasswordBody struct {

	// new password
	NewPassword string `json:"newPassword,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this put update password body
func (o *PutUpdatePasswordBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put update password body based on context it is used
func (o *PutUpdatePasswordBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUpdatePasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUpdatePasswordBody) UnmarshalBinary(b []byte) error {
	var res PutUpdatePasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
