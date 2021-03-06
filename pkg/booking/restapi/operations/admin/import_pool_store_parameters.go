// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/timdrysdale/relay/pkg/booking/models"
)

// NewImportPoolStoreParams creates a new ImportPoolStoreParams object
// no default values defined in spec.
func NewImportPoolStoreParams() ImportPoolStoreParams {

	return ImportPoolStoreParams{}
}

// ImportPoolStoreParams contains all the bound params for the import pool store operation
// typically these are obtained from a http.Request
//
// swagger:parameters importPoolStore
type ImportPoolStoreParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Poolstore *models.Poolstore
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewImportPoolStoreParams() beforehand.
func (o *ImportPoolStoreParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Poolstore
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("poolstore", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Poolstore = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
