// Code generated by go-swagger; DO NOT EDIT.

package private_key_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceCreatePrivateKeyTagReader is a Reader for the PublicAPIServiceCreatePrivateKeyTag structure.
type PublicAPIServiceCreatePrivateKeyTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceCreatePrivateKeyTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceCreatePrivateKeyTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceCreatePrivateKeyTagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceCreatePrivateKeyTagOK creates a PublicAPIServiceCreatePrivateKeyTagOK with default headers values
func NewPublicAPIServiceCreatePrivateKeyTagOK() *PublicAPIServiceCreatePrivateKeyTagOK {
	return &PublicAPIServiceCreatePrivateKeyTagOK{}
}

/*
PublicAPIServiceCreatePrivateKeyTagOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceCreatePrivateKeyTagOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service create private key tag o k response has a 2xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service create private key tag o k response has a 3xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service create private key tag o k response has a 4xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service create private key tag o k response has a 5xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service create private key tag o k response a status code equal to that given
func (o *PublicAPIServiceCreatePrivateKeyTagOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceCreatePrivateKeyTagOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_key_tag][%d] publicApiServiceCreatePrivateKeyTagOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeyTagOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_key_tag][%d] publicApiServiceCreatePrivateKeyTagOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeyTagOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceCreatePrivateKeyTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceCreatePrivateKeyTagDefault creates a PublicAPIServiceCreatePrivateKeyTagDefault with default headers values
func NewPublicAPIServiceCreatePrivateKeyTagDefault(code int) *PublicAPIServiceCreatePrivateKeyTagDefault {
	return &PublicAPIServiceCreatePrivateKeyTagDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceCreatePrivateKeyTagDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceCreatePrivateKeyTagDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service create private key tag default response
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service create private key tag default response has a 2xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service create private key tag default response has a 3xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service create private key tag default response has a 4xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service create private key tag default response has a 5xx status code
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service create private key tag default response a status code equal to that given
func (o *PublicAPIServiceCreatePrivateKeyTagDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceCreatePrivateKeyTagDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_key_tag][%d] PublicApiService_CreatePrivateKeyTag default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeyTagDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/create_private_key_tag][%d] PublicApiService_CreatePrivateKeyTag default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceCreatePrivateKeyTagDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceCreatePrivateKeyTagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
