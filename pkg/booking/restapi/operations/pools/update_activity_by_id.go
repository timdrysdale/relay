// Code generated by go-swagger; DO NOT EDIT.

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateActivityByIDHandlerFunc turns a function with the right signature into a update activity by ID handler
type UpdateActivityByIDHandlerFunc func(UpdateActivityByIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateActivityByIDHandlerFunc) Handle(params UpdateActivityByIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateActivityByIDHandler interface for that can handle valid update activity by ID params
type UpdateActivityByIDHandler interface {
	Handle(UpdateActivityByIDParams, interface{}) middleware.Responder
}

// NewUpdateActivityByID creates a new http.Handler for the update activity by ID operation
func NewUpdateActivityByID(ctx *middleware.Context, handler UpdateActivityByIDHandler) *UpdateActivityByID {
	return &UpdateActivityByID{Context: ctx, Handler: handler}
}

/*UpdateActivityByID swagger:route PUT /pools/{pool_id}/activities/{activity_id} pools updateActivityById

Update an activity in a pool

Updates an activity in a pool (or adds one with a specific ID if does not exist)

*/
type UpdateActivityByID struct {
	Context *middleware.Context
	Handler UpdateActivityByIDHandler
}

func (o *UpdateActivityByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateActivityByIDParams()

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
