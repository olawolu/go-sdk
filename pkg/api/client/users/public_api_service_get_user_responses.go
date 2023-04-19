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

// PublicAPIServiceGetUserReader is a Reader for the PublicAPIServiceGetUser structure.
type PublicAPIServiceGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicAPIServiceGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPublicAPIServiceGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPublicAPIServiceGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceGetUserOK creates a PublicAPIServiceGetUserOK with default headers values
func NewPublicAPIServiceGetUserOK() *PublicAPIServiceGetUserOK {
	return &PublicAPIServiceGetUserOK{}
}

/*
PublicAPIServiceGetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceGetUserOK struct {
	Payload *models.V1GetUserResponse
}

// IsSuccess returns true when this public Api service get user o k response has a 2xx status code
func (o *PublicAPIServiceGetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service get user o k response has a 3xx status code
func (o *PublicAPIServiceGetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get user o k response has a 4xx status code
func (o *PublicAPIServiceGetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service get user o k response has a 5xx status code
func (o *PublicAPIServiceGetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get user o k response a status code equal to that given
func (o *PublicAPIServiceGetUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceGetUserOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetUserOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetUserOK) GetPayload() *models.V1GetUserResponse {
	return o.Payload
}

func (o *PublicAPIServiceGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetUserForbidden creates a PublicAPIServiceGetUserForbidden with default headers values
func NewPublicAPIServiceGetUserForbidden() *PublicAPIServiceGetUserForbidden {
	return &PublicAPIServiceGetUserForbidden{}
}

/*
PublicAPIServiceGetUserForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PublicAPIServiceGetUserForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this public Api service get user forbidden response has a 2xx status code
func (o *PublicAPIServiceGetUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get user forbidden response has a 3xx status code
func (o *PublicAPIServiceGetUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get user forbidden response has a 4xx status code
func (o *PublicAPIServiceGetUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get user forbidden response has a 5xx status code
func (o *PublicAPIServiceGetUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get user forbidden response a status code equal to that given
func (o *PublicAPIServiceGetUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PublicAPIServiceGetUserForbidden) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetUserForbidden) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetUserForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PublicAPIServiceGetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetUserNotFound creates a PublicAPIServiceGetUserNotFound with default headers values
func NewPublicAPIServiceGetUserNotFound() *PublicAPIServiceGetUserNotFound {
	return &PublicAPIServiceGetUserNotFound{}
}

/*
PublicAPIServiceGetUserNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PublicAPIServiceGetUserNotFound struct {
	Payload string
}

// IsSuccess returns true when this public Api service get user not found response has a 2xx status code
func (o *PublicAPIServiceGetUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get user not found response has a 3xx status code
func (o *PublicAPIServiceGetUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get user not found response has a 4xx status code
func (o *PublicAPIServiceGetUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get user not found response has a 5xx status code
func (o *PublicAPIServiceGetUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get user not found response a status code equal to that given
func (o *PublicAPIServiceGetUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PublicAPIServiceGetUserNotFound) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetUserNotFound) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] publicApiServiceGetUserNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetUserNotFound) GetPayload() string {
	return o.Payload
}

func (o *PublicAPIServiceGetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetUserDefault creates a PublicAPIServiceGetUserDefault with default headers values
func NewPublicAPIServiceGetUserDefault(code int) *PublicAPIServiceGetUserDefault {
	return &PublicAPIServiceGetUserDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceGetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceGetUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service get user default response
func (o *PublicAPIServiceGetUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service get user default response has a 2xx status code
func (o *PublicAPIServiceGetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service get user default response has a 3xx status code
func (o *PublicAPIServiceGetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service get user default response has a 4xx status code
func (o *PublicAPIServiceGetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service get user default response has a 5xx status code
func (o *PublicAPIServiceGetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service get user default response a status code equal to that given
func (o *PublicAPIServiceGetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceGetUserDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] PublicApiService_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetUserDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/query/get_user][%d] PublicApiService_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
