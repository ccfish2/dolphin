// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteEndpointIDHandlerFunc turns a function with the right signature into a delete endpoint ID handler
type DeleteEndpointIDHandlerFunc func(DeleteEndpointIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEndpointIDHandlerFunc) Handle(params DeleteEndpointIDParams) middleware.Responder {
	return fn(params)
}

// DeleteEndpointIDHandler interface for that can handle valid delete endpoint ID params
type DeleteEndpointIDHandler interface {
	Handle(DeleteEndpointIDParams) middleware.Responder
}

// NewDeleteEndpointID creates a new http.Handler for the delete endpoint ID operation
func NewDeleteEndpointID(ctx *middleware.Context, handler DeleteEndpointIDHandler) *DeleteEndpointID {
	return &DeleteEndpointID{Context: ctx, Handler: handler}
}

/*
	DeleteEndpointID swagger:route DELETE /endpoint/{id} endpoint deleteEndpointId

# Delete endpoint

Deletes the endpoint specified by the ID. Deletion is imminent and
atomic, if the deletion request is valid and the endpoint exists,
deletion will occur even if errors are encountered in the process. If
errors have been encountered, the code 202 will be returned, otherwise
200 on success.

All resources associated with the endpoint will be freed and the
workload represented by the endpoint will be disconnected.It will no
longer be able to initiate or receive communications of any sort.
*/
type DeleteEndpointID struct {
	Context *middleware.Context
	Handler DeleteEndpointIDHandler
}

func (o *DeleteEndpointID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteEndpointIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
