// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/timdrysdale/relay/pkg/booking/models"
)

// AddNewPoolOKCode is the HTTP code returned for type AddNewPoolOK
const AddNewPoolOKCode int = 200

/*AddNewPoolOK add new pool o k

swagger:response addNewPoolOK
*/
type AddNewPoolOK struct {

	/*
	  In: Body
	*/
	Payload *models.ID `json:"body,omitempty"`
}

// NewAddNewPoolOK creates AddNewPoolOK with default headers values
func NewAddNewPoolOK() *AddNewPoolOK {

	return &AddNewPoolOK{}
}

// WithPayload adds the payload to the add new pool o k response
func (o *AddNewPoolOK) WithPayload(payload *models.ID) *AddNewPoolOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new pool o k response
func (o *AddNewPoolOK) SetPayload(payload *models.ID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPoolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddNewPoolUnauthorizedCode is the HTTP code returned for type AddNewPoolUnauthorized
const AddNewPoolUnauthorizedCode int = 401

/*AddNewPoolUnauthorized Unauthorized

swagger:response addNewPoolUnauthorized
*/
type AddNewPoolUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewAddNewPoolUnauthorized creates AddNewPoolUnauthorized with default headers values
func NewAddNewPoolUnauthorized() *AddNewPoolUnauthorized {

	return &AddNewPoolUnauthorized{}
}

// WithPayload adds the payload to the add new pool unauthorized response
func (o *AddNewPoolUnauthorized) WithPayload(payload interface{}) *AddNewPoolUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new pool unauthorized response
func (o *AddNewPoolUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPoolUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddNewPoolNotFoundCode is the HTTP code returned for type AddNewPoolNotFound
const AddNewPoolNotFoundCode int = 404

/*AddNewPoolNotFound Unauthorized

swagger:response addNewPoolNotFound
*/
type AddNewPoolNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewAddNewPoolNotFound creates AddNewPoolNotFound with default headers values
func NewAddNewPoolNotFound() *AddNewPoolNotFound {

	return &AddNewPoolNotFound{}
}

// WithPayload adds the payload to the add new pool not found response
func (o *AddNewPoolNotFound) WithPayload(payload interface{}) *AddNewPoolNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new pool not found response
func (o *AddNewPoolNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPoolNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AddNewPoolInternalServerErrorCode is the HTTP code returned for type AddNewPoolInternalServerError
const AddNewPoolInternalServerErrorCode int = 500

/*AddNewPoolInternalServerError add new pool internal server error

swagger:response addNewPoolInternalServerError
*/
type AddNewPoolInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewAddNewPoolInternalServerError creates AddNewPoolInternalServerError with default headers values
func NewAddNewPoolInternalServerError() *AddNewPoolInternalServerError {

	return &AddNewPoolInternalServerError{}
}

// WithPayload adds the payload to the add new pool internal server error response
func (o *AddNewPoolInternalServerError) WithPayload(payload interface{}) *AddNewPoolInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new pool internal server error response
func (o *AddNewPoolInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewPoolInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
