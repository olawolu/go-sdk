// Code generated by go-swagger; DO NOT EDIT.

package who_am_i

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new who am i API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for who am i API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PublicAPIServiceGetWhoami(params *PublicAPIServiceGetWhoamiParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetWhoamiOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PublicAPIServiceGetWhoami whos am i

Get basic information about your current API or WebAuthN user and their organization. Affords Sub-Organization look ups via Parent Organization for WebAuthN users.
*/
func (a *Client) PublicAPIServiceGetWhoami(params *PublicAPIServiceGetWhoamiParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetWhoamiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceGetWhoamiParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_GetWhoami",
		Method:             "POST",
		PathPattern:        "/public/v1/query/whoami",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceGetWhoamiReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PublicAPIServiceGetWhoamiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceGetWhoamiDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
