// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddActivityByPoolID(params *AddActivityByPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*AddActivityByPoolIDOK, error)

	AddNewPool(params *AddNewPoolParams, authInfo runtime.ClientAuthInfoWriter) (*AddNewPoolOK, error)

	DeleteActivityByID(params *DeleteActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) error

	DeletePool(params *DeletePoolParams, authInfo runtime.ClientAuthInfoWriter) error

	GetActivityByID(params *GetActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivityByIDOK, error)

	GetAllPools(params *GetAllPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPoolsOK, error)

	GetPoolDescriptionByID(params *GetPoolDescriptionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolDescriptionByIDOK, error)

	GetPoolStatusByID(params *GetPoolStatusByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolStatusByIDOK, error)

	RequestSessionByPoolID(params *RequestSessionByPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*RequestSessionByPoolIDOK, error)

	UpdateActivityByID(params *UpdateActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateActivityByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddActivityByPoolID adds an activity to a pool

  Adds an activity to a pool (activty must include valid exp field)
*/
func (a *Client) AddActivityByPoolID(params *AddActivityByPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*AddActivityByPoolIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddActivityByPoolIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addActivityByPoolID",
		Method:             "POST",
		PathPattern:        "/pools/{pool_id}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddActivityByPoolIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddActivityByPoolIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addActivityByPoolID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddNewPool adds a new pool

  Add a pool to the poolstore, using details in body
*/
func (a *Client) AddNewPool(params *AddNewPoolParams, authInfo runtime.ClientAuthInfoWriter) (*AddNewPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNewPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addNewPool",
		Method:             "POST",
		PathPattern:        "/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddNewPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddNewPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addNewPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteActivityByID deletes activity

  Delete activity by activity ID
*/
func (a *Client) DeleteActivityByID(params *DeleteActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteActivityByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteActivityByID",
		Method:             "DELETE",
		PathPattern:        "/pools/{pool_id}/activities/{activity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteActivityByIDReader{formats: a.formats},
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
  DeletePool deletes this pool

  Delete this pool and all its data (including activities)
*/
func (a *Client) DeletePool(params *DeletePoolParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePoolParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePool",
		Method:             "DELETE",
		PathPattern:        "/pools/{pool_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePoolReader{formats: a.formats},
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
  GetActivityByID gets activity description

  Get activity description by activity ID
*/
func (a *Client) GetActivityByID(params *GetActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetActivityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivityByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getActivityByID",
		Method:             "GET",
		PathPattern:        "/pools/{pool_id}/activities/{activity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivityByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetActivityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getActivityByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllPools pools

  Gets a list of all pool_ids
*/
func (a *Client) GetAllPools(params *GetAllPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllPools",
		Method:             "GET",
		PathPattern:        "/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllPoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllPools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPoolDescriptionByID pools

  Gets a description of the pool
*/
func (a *Client) GetPoolDescriptionByID(params *GetPoolDescriptionByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolDescriptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolDescriptionByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPoolDescriptionByID",
		Method:             "GET",
		PathPattern:        "/pools/{pool_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolDescriptionByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPoolDescriptionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoolDescriptionByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPoolStatusByID gets the status of the pool s activities

  Gets the status of pool's activities
*/
func (a *Client) GetPoolStatusByID(params *GetPoolStatusByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolStatusByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolStatusByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPoolStatusByID",
		Method:             "GET",
		PathPattern:        "/pools/{pool_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolStatusByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPoolStatusByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoolStatusByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RequestSessionByPoolID requests a session from a pool

  Request a session on an activity from the pool
*/
func (a *Client) RequestSessionByPoolID(params *RequestSessionByPoolIDParams, authInfo runtime.ClientAuthInfoWriter) (*RequestSessionByPoolIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestSessionByPoolIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestSessionByPoolID",
		Method:             "POST",
		PathPattern:        "/pools/{pool_id}/sessions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestSessionByPoolIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RequestSessionByPoolIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for requestSessionByPoolID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateActivityByID updates an activity in a pool

  Updates an activity in a pool (or adds one with a specific ID if does not exist)
*/
func (a *Client) UpdateActivityByID(params *UpdateActivityByIDParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateActivityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActivityByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateActivityByID",
		Method:             "PUT",
		PathPattern:        "/pools/{pool_id}/activities/{activity_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateActivityByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateActivityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateActivityByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
