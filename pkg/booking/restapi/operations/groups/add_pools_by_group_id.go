// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddPoolsByGroupIDHandlerFunc turns a function with the right signature into a add pools by group ID handler
type AddPoolsByGroupIDHandlerFunc func(AddPoolsByGroupIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddPoolsByGroupIDHandlerFunc) Handle(params AddPoolsByGroupIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddPoolsByGroupIDHandler interface for that can handle valid add pools by group ID params
type AddPoolsByGroupIDHandler interface {
	Handle(AddPoolsByGroupIDParams, interface{}) middleware.Responder
}

// NewAddPoolsByGroupID creates a new http.Handler for the add pools by group ID operation
func NewAddPoolsByGroupID(ctx *middleware.Context, handler AddPoolsByGroupIDHandler) *AddPoolsByGroupID {
	return &AddPoolsByGroupID{Context: ctx, Handler: handler}
}

/*AddPoolsByGroupID swagger:route POST /groups/{group_id}/pools groups addPoolsByGroupId

groups

Add a list of pool_ids in the group (keep existing). Return new complete list.

*/
type AddPoolsByGroupID struct {
	Context *middleware.Context
	Handler AddPoolsByGroupIDHandler
}

func (o *AddPoolsByGroupID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddPoolsByGroupIDParams()

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
