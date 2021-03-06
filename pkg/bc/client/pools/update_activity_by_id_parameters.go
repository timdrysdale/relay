// Code generated by go-swagger; DO NOT EDIT.

package pools

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

	"github.com/timdrysdale/relay/pkg/bc/models"
)

// NewUpdateActivityByIDParams creates a new UpdateActivityByIDParams object
// with the default values initialized.
func NewUpdateActivityByIDParams() *UpdateActivityByIDParams {
	var ()
	return &UpdateActivityByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateActivityByIDParamsWithTimeout creates a new UpdateActivityByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateActivityByIDParamsWithTimeout(timeout time.Duration) *UpdateActivityByIDParams {
	var ()
	return &UpdateActivityByIDParams{

		timeout: timeout,
	}
}

// NewUpdateActivityByIDParamsWithContext creates a new UpdateActivityByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateActivityByIDParamsWithContext(ctx context.Context) *UpdateActivityByIDParams {
	var ()
	return &UpdateActivityByIDParams{

		Context: ctx,
	}
}

// NewUpdateActivityByIDParamsWithHTTPClient creates a new UpdateActivityByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateActivityByIDParamsWithHTTPClient(client *http.Client) *UpdateActivityByIDParams {
	var ()
	return &UpdateActivityByIDParams{
		HTTPClient: client,
	}
}

/*UpdateActivityByIDParams contains all the parameters to send to the API endpoint
for the update activity by ID operation typically these are written to a http.Request
*/
type UpdateActivityByIDParams struct {

	/*Activity*/
	Activity *models.Activity
	/*ActivityID*/
	ActivityID string
	/*PoolID*/
	PoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update activity by ID params
func (o *UpdateActivityByIDParams) WithTimeout(timeout time.Duration) *UpdateActivityByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update activity by ID params
func (o *UpdateActivityByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update activity by ID params
func (o *UpdateActivityByIDParams) WithContext(ctx context.Context) *UpdateActivityByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update activity by ID params
func (o *UpdateActivityByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update activity by ID params
func (o *UpdateActivityByIDParams) WithHTTPClient(client *http.Client) *UpdateActivityByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update activity by ID params
func (o *UpdateActivityByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivity adds the activity to the update activity by ID params
func (o *UpdateActivityByIDParams) WithActivity(activity *models.Activity) *UpdateActivityByIDParams {
	o.SetActivity(activity)
	return o
}

// SetActivity adds the activity to the update activity by ID params
func (o *UpdateActivityByIDParams) SetActivity(activity *models.Activity) {
	o.Activity = activity
}

// WithActivityID adds the activityID to the update activity by ID params
func (o *UpdateActivityByIDParams) WithActivityID(activityID string) *UpdateActivityByIDParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the update activity by ID params
func (o *UpdateActivityByIDParams) SetActivityID(activityID string) {
	o.ActivityID = activityID
}

// WithPoolID adds the poolID to the update activity by ID params
func (o *UpdateActivityByIDParams) WithPoolID(poolID string) *UpdateActivityByIDParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the update activity by ID params
func (o *UpdateActivityByIDParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateActivityByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Activity != nil {
		if err := r.SetBodyParam(o.Activity); err != nil {
			return err
		}
	}

	// path param activity_id
	if err := r.SetPathParam("activity_id", o.ActivityID); err != nil {
		return err
	}

	// path param pool_id
	if err := r.SetPathParam("pool_id", o.PoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
