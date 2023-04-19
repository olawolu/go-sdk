// Code generated by go-swagger; DO NOT EDIT.

package private_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new private keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for private keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PublicAPIServiceCreatePrivateKeys(params *PublicAPIServiceCreatePrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceCreatePrivateKeysOK, error)

	PublicAPIServiceGetPrivateKey(params *PublicAPIServiceGetPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetPrivateKeyOK, error)

	PublicAPIServiceGetPrivateKeys(params *PublicAPIServiceGetPrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetPrivateKeysOK, error)

	PublicAPIServiceSignRawPayload(params *PublicAPIServiceSignRawPayloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceSignRawPayloadOK, error)

	PublicAPIServiceSignTransaction(params *PublicAPIServiceSignTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceSignTransactionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PublicAPIServiceCreatePrivateKeys creates private keys

Create new Private Keys
*/
func (a *Client) PublicAPIServiceCreatePrivateKeys(params *PublicAPIServiceCreatePrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceCreatePrivateKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceCreatePrivateKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_CreatePrivateKeys",
		Method:             "POST",
		PathPattern:        "/public/v1/submit/create_private_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceCreatePrivateKeysReader{formats: a.formats},
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
	success, ok := result.(*PublicAPIServiceCreatePrivateKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceCreatePrivateKeysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PublicAPIServiceGetPrivateKey gets private key

Get details about a Private Key
*/
func (a *Client) PublicAPIServiceGetPrivateKey(params *PublicAPIServiceGetPrivateKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetPrivateKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceGetPrivateKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_GetPrivateKey",
		Method:             "POST",
		PathPattern:        "/tkhq/public/v1/query/get_private_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceGetPrivateKeyReader{formats: a.formats},
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
	success, ok := result.(*PublicAPIServiceGetPrivateKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceGetPrivateKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PublicAPIServiceGetPrivateKeys lists private keys

List all Private Keys within an Organization
*/
func (a *Client) PublicAPIServiceGetPrivateKeys(params *PublicAPIServiceGetPrivateKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceGetPrivateKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceGetPrivateKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_GetPrivateKeys",
		Method:             "POST",
		PathPattern:        "/public/v1/query/list_private_keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceGetPrivateKeysReader{formats: a.formats},
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
	success, ok := result.(*PublicAPIServiceGetPrivateKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceGetPrivateKeysDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PublicAPIServiceSignRawPayload signs raw payload

Sign a raw payload with a Private Key
*/
func (a *Client) PublicAPIServiceSignRawPayload(params *PublicAPIServiceSignRawPayloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceSignRawPayloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceSignRawPayloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_SignRawPayload",
		Method:             "POST",
		PathPattern:        "/public/v1/submit/sign_raw_payload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceSignRawPayloadReader{formats: a.formats},
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
	success, ok := result.(*PublicAPIServiceSignRawPayloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceSignRawPayloadDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PublicAPIServiceSignTransaction signs transaction

Sign a transaction with a Private Key
*/
func (a *Client) PublicAPIServiceSignTransaction(params *PublicAPIServiceSignTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PublicAPIServiceSignTransactionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPublicAPIServiceSignTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PublicApiService_SignTransaction",
		Method:             "POST",
		PathPattern:        "/public/v1/submit/sign_transaction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PublicAPIServiceSignTransactionReader{formats: a.formats},
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
	success, ok := result.(*PublicAPIServiceSignTransactionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PublicAPIServiceSignTransactionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
