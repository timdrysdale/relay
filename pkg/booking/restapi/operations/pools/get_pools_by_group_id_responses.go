// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPoolsByGroupIDOKCode is the HTTP code returned for type GetPoolsByGroupIDOK
const GetPoolsByGroupIDOKCode int = 200

/*GetPoolsByGroupIDOK get pools by group Id o k

swagger:response getPoolsByGroupIdOK
*/
type GetPoolsByGroupIDOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetPoolsByGroupIDOK creates GetPoolsByGroupIDOK with default headers values
func NewGetPoolsByGroupIDOK() *GetPoolsByGroupIDOK {

	return &GetPoolsByGroupIDOK{}
}

// WithPayload adds the payload to the get pools by group Id o k response
func (o *GetPoolsByGroupIDOK) WithPayload(payload []string) *GetPoolsByGroupIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pools by group Id o k response
func (o *GetPoolsByGroupIDOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPoolsByGroupIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
