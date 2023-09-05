// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceCreateUsersReader is a Reader for the PublicAPIServiceCreateUsers structure.
type PublicAPIServiceCreateUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceCreateUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceCreateUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceCreateUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceCreateUsersOK creates a PublicAPIServiceCreateUsersOK with default headers values
func NewPublicAPIServiceCreateUsersOK() *PublicAPIServiceCreateUsersOK {
	return &PublicAPIServiceCreateUsersOK{}
}

/*
PublicAPIServiceCreateUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceCreateUsersOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service create users o k response has a 2xx status code
func (o *PublicAPIServiceCreateUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service create users o k response has a 3xx status code
func (o *PublicAPIServiceCreateUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service create users o k response has a 4xx status code
func (o *PublicAPIServiceCreateUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service create users o k response has a 5xx status code
func (o *PublicAPIServiceCreateUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service create users o k response a status code equal to that given
func (o *PublicAPIServiceCreateUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceCreateUsersOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_users][%d] publicApiServiceCreateUsersOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreateUsersOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_users][%d] publicApiServiceCreateUsersOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreateUsersOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceCreateUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceCreateUsersDefault creates a PublicAPIServiceCreateUsersDefault with default headers values
func NewPublicAPIServiceCreateUsersDefault(code int) *PublicAPIServiceCreateUsersDefault {
	return &PublicAPIServiceCreateUsersDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceCreateUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceCreateUsersDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service create users default response
func (o *PublicAPIServiceCreateUsersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service create users default response has a 2xx status code
func (o *PublicAPIServiceCreateUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service create users default response has a 3xx status code
func (o *PublicAPIServiceCreateUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service create users default response has a 4xx status code
func (o *PublicAPIServiceCreateUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service create users default response has a 5xx status code
func (o *PublicAPIServiceCreateUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service create users default response a status code equal to that given
func (o *PublicAPIServiceCreateUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceCreateUsersDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_users][%d] PublicApiService_CreateUsers default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreateUsersDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_users][%d] PublicApiService_CreateUsers default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreateUsersDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceCreateUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
