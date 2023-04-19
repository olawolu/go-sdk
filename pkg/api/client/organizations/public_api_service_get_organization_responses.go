// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceGetOrganizationReader is a Reader for the PublicAPIServiceGetOrganization structure.
type PublicAPIServiceGetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceGetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicAPIServiceGetOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPublicAPIServiceGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPublicAPIServiceGetOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceGetOrganizationOK creates a PublicAPIServiceGetOrganizationOK with default headers values
func NewPublicAPIServiceGetOrganizationOK() *PublicAPIServiceGetOrganizationOK {
	return &PublicAPIServiceGetOrganizationOK{}
}

/*
PublicAPIServiceGetOrganizationOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceGetOrganizationOK struct {
	Payload *models.V1GetOrganizationResponse
}

// IsSuccess returns true when this public Api service get organization o k response has a 2xx status code
func (o *PublicAPIServiceGetOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service get organization o k response has a 3xx status code
func (o *PublicAPIServiceGetOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get organization o k response has a 4xx status code
func (o *PublicAPIServiceGetOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service get organization o k response has a 5xx status code
func (o *PublicAPIServiceGetOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get organization o k response a status code equal to that given
func (o *PublicAPIServiceGetOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceGetOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationOK) GetPayload() *models.V1GetOrganizationResponse {
	return o.Payload
}

func (o *PublicAPIServiceGetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetOrganizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetOrganizationForbidden creates a PublicAPIServiceGetOrganizationForbidden with default headers values
func NewPublicAPIServiceGetOrganizationForbidden() *PublicAPIServiceGetOrganizationForbidden {
	return &PublicAPIServiceGetOrganizationForbidden{}
}

/*
PublicAPIServiceGetOrganizationForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PublicAPIServiceGetOrganizationForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this public Api service get organization forbidden response has a 2xx status code
func (o *PublicAPIServiceGetOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get organization forbidden response has a 3xx status code
func (o *PublicAPIServiceGetOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get organization forbidden response has a 4xx status code
func (o *PublicAPIServiceGetOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get organization forbidden response has a 5xx status code
func (o *PublicAPIServiceGetOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get organization forbidden response a status code equal to that given
func (o *PublicAPIServiceGetOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PublicAPIServiceGetOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationForbidden) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PublicAPIServiceGetOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetOrganizationNotFound creates a PublicAPIServiceGetOrganizationNotFound with default headers values
func NewPublicAPIServiceGetOrganizationNotFound() *PublicAPIServiceGetOrganizationNotFound {
	return &PublicAPIServiceGetOrganizationNotFound{}
}

/*
PublicAPIServiceGetOrganizationNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PublicAPIServiceGetOrganizationNotFound struct {
	Payload string
}

// IsSuccess returns true when this public Api service get organization not found response has a 2xx status code
func (o *PublicAPIServiceGetOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get organization not found response has a 3xx status code
func (o *PublicAPIServiceGetOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get organization not found response has a 4xx status code
func (o *PublicAPIServiceGetOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get organization not found response has a 5xx status code
func (o *PublicAPIServiceGetOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get organization not found response a status code equal to that given
func (o *PublicAPIServiceGetOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PublicAPIServiceGetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationNotFound) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] publicApiServiceGetOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationNotFound) GetPayload() string {
	return o.Payload
}

func (o *PublicAPIServiceGetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetOrganizationDefault creates a PublicAPIServiceGetOrganizationDefault with default headers values
func NewPublicAPIServiceGetOrganizationDefault(code int) *PublicAPIServiceGetOrganizationDefault {
	return &PublicAPIServiceGetOrganizationDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceGetOrganizationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceGetOrganizationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service get organization default response
func (o *PublicAPIServiceGetOrganizationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service get organization default response has a 2xx status code
func (o *PublicAPIServiceGetOrganizationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service get organization default response has a 3xx status code
func (o *PublicAPIServiceGetOrganizationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service get organization default response has a 4xx status code
func (o *PublicAPIServiceGetOrganizationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service get organization default response has a 5xx status code
func (o *PublicAPIServiceGetOrganizationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service get organization default response a status code equal to that given
func (o *PublicAPIServiceGetOrganizationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceGetOrganizationDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] PublicApiService_GetOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_organization][%d] PublicApiService_GetOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetOrganizationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceGetOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
