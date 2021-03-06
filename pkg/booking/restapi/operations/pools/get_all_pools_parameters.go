// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAllPoolsParams creates a new GetAllPoolsParams object
// no default values defined in spec.
func NewGetAllPoolsParams() GetAllPoolsParams {

	return GetAllPoolsParams{}
}

// GetAllPoolsParams contains all the bound params for the get all pools operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAllPools
type GetAllPoolsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Limit search to pools with exact match to the name (meaningless on own)
	  In: query
	*/
	Exact *bool
	/*Limit search to pools with name containing this string (case sensitive)
	  In: query
	*/
	Name *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllPoolsParams() beforehand.
func (o *GetAllPoolsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qExact, qhkExact, _ := qs.GetOK("exact")
	if err := o.bindExact(qExact, qhkExact, route.Formats); err != nil {
		res = append(res, err)
	}

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExact binds and validates parameter Exact from query.
func (o *GetAllPoolsParams) bindExact(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("exact", "query", "bool", raw)
	}
	o.Exact = &value

	return nil
}

// bindName binds and validates parameter Name from query.
func (o *GetAllPoolsParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}
