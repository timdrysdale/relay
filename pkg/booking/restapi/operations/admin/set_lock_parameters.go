// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewSetLockParams creates a new SetLockParams object
// no default values defined in spec.
func NewSetLockParams() SetLockParams {

	return SetLockParams{}
}

// SetLockParams contains all the bound params for the set lock operation
// typically these are obtained from a http.Request
//
// swagger:parameters setLock
type SetLockParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*set booking lock
	  Required: true
	  In: query
	*/
	Lock bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSetLockParams() beforehand.
func (o *SetLockParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLock, qhkLock, _ := qs.GetOK("lock")
	if err := o.bindLock(qLock, qhkLock, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLock binds and validates parameter Lock from query.
func (o *SetLockParams) bindLock(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("lock", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("lock", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("lock", "query", "bool", raw)
	}
	o.Lock = value

	return nil
}