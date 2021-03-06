// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPoolsByGroupIDParams creates a new GetPoolsByGroupIDParams object
// with the default values initialized.
func NewGetPoolsByGroupIDParams() *GetPoolsByGroupIDParams {
	var ()
	return &GetPoolsByGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoolsByGroupIDParamsWithTimeout creates a new GetPoolsByGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoolsByGroupIDParamsWithTimeout(timeout time.Duration) *GetPoolsByGroupIDParams {
	var ()
	return &GetPoolsByGroupIDParams{

		timeout: timeout,
	}
}

// NewGetPoolsByGroupIDParamsWithContext creates a new GetPoolsByGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoolsByGroupIDParamsWithContext(ctx context.Context) *GetPoolsByGroupIDParams {
	var ()
	return &GetPoolsByGroupIDParams{

		Context: ctx,
	}
}

// NewGetPoolsByGroupIDParamsWithHTTPClient creates a new GetPoolsByGroupIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoolsByGroupIDParamsWithHTTPClient(client *http.Client) *GetPoolsByGroupIDParams {
	var ()
	return &GetPoolsByGroupIDParams{
		HTTPClient: client,
	}
}

/*GetPoolsByGroupIDParams contains all the parameters to send to the API endpoint
for the get pools by group ID operation typically these are written to a http.Request
*/
type GetPoolsByGroupIDParams struct {

	/*GroupID*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) WithTimeout(timeout time.Duration) *GetPoolsByGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) WithContext(ctx context.Context) *GetPoolsByGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) WithHTTPClient(client *http.Client) *GetPoolsByGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) WithGroupID(groupID string) *GetPoolsByGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get pools by group ID params
func (o *GetPoolsByGroupIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoolsByGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
