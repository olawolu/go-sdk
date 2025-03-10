// Code generated by go-swagger; DO NOT EDIT.

package private_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceGetPrivateKeysReader is a Reader for the PublicAPIServiceGetPrivateKeys structure.
type PublicAPIServiceGetPrivateKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceGetPrivateKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceGetPrivateKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceGetPrivateKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceGetPrivateKeysOK creates a PublicAPIServiceGetPrivateKeysOK with default headers values
func NewPublicAPIServiceGetPrivateKeysOK() *PublicAPIServiceGetPrivateKeysOK {
	return &PublicAPIServiceGetPrivateKeysOK{}
}

/*
PublicAPIServiceGetPrivateKeysOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceGetPrivateKeysOK struct {
	Payload *models.V1GetPrivateKeysResponse
}

// IsSuccess returns true when this public Api service get private keys o k response has a 2xx status code
func (o *PublicAPIServiceGetPrivateKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service get private keys o k response has a 3xx status code
func (o *PublicAPIServiceGetPrivateKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get private keys o k response has a 4xx status code
func (o *PublicAPIServiceGetPrivateKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service get private keys o k response has a 5xx status code
func (o *PublicAPIServiceGetPrivateKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get private keys o k response a status code equal to that given
func (o *PublicAPIServiceGetPrivateKeysOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceGetPrivateKeysOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_private_keys][%d] publicApiServiceGetPrivateKeysOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetPrivateKeysOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_private_keys][%d] publicApiServiceGetPrivateKeysOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetPrivateKeysOK) GetPayload() *models.V1GetPrivateKeysResponse {
	return o.Payload
}

func (o *PublicAPIServiceGetPrivateKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetPrivateKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetPrivateKeysDefault creates a PublicAPIServiceGetPrivateKeysDefault with default headers values
func NewPublicAPIServiceGetPrivateKeysDefault(code int) *PublicAPIServiceGetPrivateKeysDefault {
	return &PublicAPIServiceGetPrivateKeysDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceGetPrivateKeysDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceGetPrivateKeysDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service get private keys default response
func (o *PublicAPIServiceGetPrivateKeysDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service get private keys default response has a 2xx status code
func (o *PublicAPIServiceGetPrivateKeysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service get private keys default response has a 3xx status code
func (o *PublicAPIServiceGetPrivateKeysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service get private keys default response has a 4xx status code
func (o *PublicAPIServiceGetPrivateKeysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service get private keys default response has a 5xx status code
func (o *PublicAPIServiceGetPrivateKeysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service get private keys default response a status code equal to that given
func (o *PublicAPIServiceGetPrivateKeysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceGetPrivateKeysDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_private_keys][%d] PublicApiService_GetPrivateKeys default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetPrivateKeysDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_private_keys][%d] PublicApiService_GetPrivateKeys default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetPrivateKeysDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceGetPrivateKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
