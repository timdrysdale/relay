// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteGroupHandlerFunc turns a function with the right signature into a delete group handler
type DeleteGroupHandlerFunc func(DeleteGroupParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteGroupHandlerFunc) Handle(params DeleteGroupParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteGroupHandler interface for that can handle valid delete group params
type DeleteGroupHandler interface {
	Handle(DeleteGroupParams, interface{}) middleware.Responder
}

// NewDeleteGroup creates a new http.Handler for the delete group operation
func NewDeleteGroup(ctx *middleware.Context, handler DeleteGroupHandler) *DeleteGroup {
	return &DeleteGroup{Context: ctx, Handler: handler}
}

/*DeleteGroup swagger:route DELETE /groups/{group_id} groups deleteGroup

groups

Delete this group, but not the pools associated with it.

*/
type DeleteGroup struct {
	Context *middleware.Context
	Handler DeleteGroupHandler
}

func (o *DeleteGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteGroupParams()

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
