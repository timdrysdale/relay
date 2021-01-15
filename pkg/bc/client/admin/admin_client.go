// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePoolStore(params *DeletePoolStoreParams, authInfo runtime.ClientAuthInfoWriter) error

	ExportPoolStore(params *ExportPoolStoreParams, authInfo runtime.ClientAuthInfoWriter) (*ExportPoolStoreOK, error)

	GetStoreStatus(params *GetStoreStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetStoreStatusOK, error)

	ImportPoolStore(params *ImportPoolStoreParams, authInfo runtime.ClientAuthInfoWriter) (*ImportPoolStoreOK, error)

	SetLock(params *SetLockParams, authInfo runtime.ClientAuthInfoWriter) (*SetLockOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePoolStore deletes current state

  Delete the contents of the pool store including bookings
*/
func (a *Client) DeletePoolStore(params *DeletePoolStoreParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePoolStoreParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePoolStore",
		Method:             "DELETE",
		PathPattern:        "/admin/poolstore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePoolStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  ExportPoolStore exports current state

  Export the current pool store including bookings
*/
func (a *Client) ExportPoolStore(params *ExportPoolStoreParams, authInfo runtime.ClientAuthInfoWriter) (*ExportPoolStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportPoolStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportPoolStore",
		Method:             "GET",
		PathPattern:        "/admin/poolstore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportPoolStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExportPoolStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportPoolStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStoreStatus gets current status

  Get the current status (number of groups, pools, bookings, time til last booking finished etc)
*/
func (a *Client) GetStoreStatus(params *GetStoreStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetStoreStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoreStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStoreStatus",
		Method:             "GET",
		PathPattern:        "/admin/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoreStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStoreStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStoreStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportPoolStore imports new current state

  Import a new pool store including bookings
*/
func (a *Client) ImportPoolStore(params *ImportPoolStoreParams, authInfo runtime.ClientAuthInfoWriter) (*ImportPoolStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportPoolStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importPoolStore",
		Method:             "POST",
		PathPattern:        "/admin/poolstore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportPoolStoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportPoolStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importPoolStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetLock sets release booking lock

  Set whether the bookings are locked or not
*/
func (a *Client) SetLock(params *SetLockParams, authInfo runtime.ClientAuthInfoWriter) (*SetLockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetLockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setLock",
		Method:             "POST",
		PathPattern:        "/admin/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetLockReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetLockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setLock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
