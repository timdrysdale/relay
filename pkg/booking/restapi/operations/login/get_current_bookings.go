// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCurrentBookingsHandlerFunc turns a function with the right signature into a get current bookings handler
type GetCurrentBookingsHandlerFunc func(GetCurrentBookingsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCurrentBookingsHandlerFunc) Handle(params GetCurrentBookingsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCurrentBookingsHandler interface for that can handle valid get current bookings params
type GetCurrentBookingsHandler interface {
	Handle(GetCurrentBookingsParams, interface{}) middleware.Responder
}

// NewGetCurrentBookings creates a new http.Handler for the get current bookings operation
func NewGetCurrentBookings(ctx *middleware.Context, handler GetCurrentBookingsHandler) *GetCurrentBookings {
	return &GetCurrentBookings{Context: ctx, Handler: handler}
}

/*GetCurrentBookings swagger:route GET /login login getCurrentBookings

Get current bookings

Get details of currently held bookings and max number of bookings that can be held

*/
type GetCurrentBookings struct {
	Context *middleware.Context
	Handler GetCurrentBookingsHandler
}

func (o *GetCurrentBookings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCurrentBookingsParams()

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
