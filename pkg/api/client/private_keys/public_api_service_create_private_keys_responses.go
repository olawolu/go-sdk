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

// PublicAPIServiceCreatePrivateKeysReader is a Reader for the PublicAPIServiceCreatePrivateKeys structure.
type PublicAPIServiceCreatePrivateKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceCreatePrivateKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceCreatePrivateKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceCreatePrivateKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceCreatePrivateKeysOK creates a PublicAPIServiceCreatePrivateKeysOK with default headers values
func NewPublicAPIServiceCreatePrivateKeysOK() *PublicAPIServiceCreatePrivateKeysOK {
	return &PublicAPIServiceCreatePrivateKeysOK{}
}

/*
PublicAPIServiceCreatePrivateKeysOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceCreatePrivateKeysOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service create private keys o k response has a 2xx status code
func (o *PublicAPIServiceCreatePrivateKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service create private keys o k response has a 3xx status code
func (o *PublicAPIServiceCreatePrivateKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service create private keys o k response has a 4xx status code
func (o *PublicAPIServiceCreatePrivateKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service create private keys o k response has a 5xx status code
func (o *PublicAPIServiceCreatePrivateKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service create private keys o k response a status code equal to that given
func (o *PublicAPIServiceCreatePrivateKeysOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceCreatePrivateKeysOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_keys][%d] publicApiServiceCreatePrivateKeysOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeysOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_keys][%d] publicApiServiceCreatePrivateKeysOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeysOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceCreatePrivateKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceCreatePrivateKeysDefault creates a PublicAPIServiceCreatePrivateKeysDefault with default headers values
func NewPublicAPIServiceCreatePrivateKeysDefault(code int) *PublicAPIServiceCreatePrivateKeysDefault {
	return &PublicAPIServiceCreatePrivateKeysDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceCreatePrivateKeysDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceCreatePrivateKeysDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service create private keys default response
func (o *PublicAPIServiceCreatePrivateKeysDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service create private keys default response has a 2xx status code
func (o *PublicAPIServiceCreatePrivateKeysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service create private keys default response has a 3xx status code
func (o *PublicAPIServiceCreatePrivateKeysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service create private keys default response has a 4xx status code
func (o *PublicAPIServiceCreatePrivateKeysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service create private keys default response has a 5xx status code
func (o *PublicAPIServiceCreatePrivateKeysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service create private keys default response a status code equal to that given
func (o *PublicAPIServiceCreatePrivateKeysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceCreatePrivateKeysDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_keys][%d] PublicApiService_CreatePrivateKeys default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeysDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_keys][%d] PublicApiService_CreatePrivateKeys default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeysDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceCreatePrivateKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
