// Code generated by go-swagger; DO NOT EDIT.

package consensus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/go-sdk/pkg/api/models"
)

// PublicAPIServiceApproveActivityReader is a Reader for the PublicAPIServiceApproveActivity structure.
type PublicAPIServiceApproveActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceApproveActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceApproveActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPublicAPIServiceApproveActivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceApproveActivityOK creates a PublicAPIServiceApproveActivityOK with default headers values
func NewPublicAPIServiceApproveActivityOK() *PublicAPIServiceApproveActivityOK {
	return &PublicAPIServiceApproveActivityOK{}
}

/*
PublicAPIServiceApproveActivityOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceApproveActivityOK struct {
	Payload *models.V1ActivityResponse
}

// IsSuccess returns true when this public Api service approve activity o k response has a 2xx status code
func (o *PublicAPIServiceApproveActivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service approve activity o k response has a 3xx status code
func (o *PublicAPIServiceApproveActivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service approve activity o k response has a 4xx status code
func (o *PublicAPIServiceApproveActivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service approve activity o k response has a 5xx status code
func (o *PublicAPIServiceApproveActivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service approve activity o k response a status code equal to that given
func (o *PublicAPIServiceApproveActivityOK) IsCode(code int) bool {
	return code == 200
}

func (o *PublicAPIServiceApproveActivityOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/approve_activity][%d] publicApiServiceApproveActivityOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceApproveActivityOK) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/approve_activity][%d] publicApiServiceApproveActivityOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceApproveActivityOK) GetPayload() *models.V1ActivityResponse {
	return o.Payload
}

func (o *PublicAPIServiceApproveActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceApproveActivityDefault creates a PublicAPIServiceApproveActivityDefault with default headers values
func NewPublicAPIServiceApproveActivityDefault(code int) *PublicAPIServiceApproveActivityDefault {
	return &PublicAPIServiceApproveActivityDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceApproveActivityDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceApproveActivityDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the public Api service approve activity default response
func (o *PublicAPIServiceApproveActivityDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this public Api service approve activity default response has a 2xx status code
func (o *PublicAPIServiceApproveActivityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service approve activity default response has a 3xx status code
func (o *PublicAPIServiceApproveActivityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service approve activity default response has a 4xx status code
func (o *PublicAPIServiceApproveActivityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service approve activity default response has a 5xx status code
func (o *PublicAPIServiceApproveActivityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service approve activity default response a status code equal to that given
func (o *PublicAPIServiceApproveActivityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PublicAPIServiceApproveActivityDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/submit/approve_activity][%d] PublicApiService_ApproveActivity default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceApproveActivityDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/submit/approve_activity][%d] PublicApiService_ApproveActivity default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceApproveActivityDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceApproveActivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
