// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddNewPoolHandlerFunc turns a function with the right signature into a add new pool handler
type AddNewPoolHandlerFunc func(AddNewPoolParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNewPoolHandlerFunc) Handle(params AddNewPoolParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddNewPoolHandler interface for that can handle valid add new pool params
type AddNewPoolHandler interface {
	Handle(AddNewPoolParams, interface{}) middleware.Responder
}

// NewAddNewPool creates a new http.Handler for the add new pool operation
func NewAddNewPool(ctx *middleware.Context, handler AddNewPoolHandler) *AddNewPool {
	return &AddNewPool{Context: ctx, Handler: handler}
}

/*AddNewPool swagger:route POST /pools pools addNewPool

Add a new pool

Add a pool to the poolstore, using details in body

*/
type AddNewPool struct {
	Context *middleware.Context
	Handler AddNewPoolHandler
}

func (o *AddNewPool) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddNewPoolParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}