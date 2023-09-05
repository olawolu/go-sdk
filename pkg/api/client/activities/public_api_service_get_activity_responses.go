// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceGetActivityReader is a Reader for the PublicAPIServiceGetActivity structure.
type PublicAPIServiceGetActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceGetActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceGetActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceGetActivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceGetActivityOK creates a PublicAPIServiceGetActivityOK with default headers values
func NewPublicAPIServiceGetActivityOK() *PublicAPIServiceGetActivityOK {
	return &PublicAPIServiceGetActivityOK{}
}

/*
PublicAPIServiceGetActivityOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceGetActivityOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service get activity o k response has a 2xx status code
func (o *PublicAPIServiceGetActivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service get activity o k response has a 3xx status code
func (o *PublicAPIServiceGetActivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get activity o k response has a 4xx status code
func (o *PublicAPIServiceGetActivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service get activity o k response has a 5xx status code
func (o *PublicAPIServiceGetActivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get activity o k response a status code equal to that given
func (o *PublicAPIServiceGetActivityOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceGetActivityOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_activity][%d] publicApiServiceGetActivityOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetActivityOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_activity][%d] publicApiServiceGetActivityOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetActivityOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceGetActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetActivityDefault creates a PublicAPIServiceGetActivityDefault with default headers values
func NewPublicAPIServiceGetActivityDefault(code int) *PublicAPIServiceGetActivityDefault {
	return &PublicAPIServiceGetActivityDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceGetActivityDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceGetActivityDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service get activity default response
func (o *PublicAPIServiceGetActivityDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service get activity default response has a 2xx status code
func (o *PublicAPIServiceGetActivityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service get activity default response has a 3xx status code
func (o *PublicAPIServiceGetActivityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service get activity default response has a 4xx status code
func (o *PublicAPIServiceGetActivityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service get activity default response has a 5xx status code
func (o *PublicAPIServiceGetActivityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service get activity default response a status code equal to that given
func (o *PublicAPIServiceGetActivityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceGetActivityDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_activity][%d] PublicApiService_GetActivity default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetActivityDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_activity][%d] PublicApiService_GetActivity default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetActivityDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceGetActivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
