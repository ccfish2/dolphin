// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMapNameEventsHandlerFunc turns a function with the right signature into a get map name events handler
type GetMapNameEventsHandlerFunc func(GetMapNameEventsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMapNameEventsHandlerFunc) Handle(params GetMapNameEventsParams) middleware.Responder {
	return fn(params)
}

// GetMapNameEventsHandler interface for that can handle valid get map name events params
type GetMapNameEventsHandler interface {
	Handle(GetMapNameEventsParams) middleware.Responder
}

// NewGetMapNameEvents creates a new http.Handler for the get map name events operation
func NewGetMapNameEvents(ctx *middleware.Context, handler GetMapNameEventsHandler) *GetMapNameEvents {
	return &GetMapNameEvents{Context: ctx, Handler: handler}
}

/*
	GetMapNameEvents swagger:route GET /map/{name}/events daemon getMapNameEvents

Retrieves the recent event logs associated with this endpoint.
*/
type GetMapNameEvents struct {
	Context *middleware.Context
	Handler GetMapNameEventsHandler
}

func (o *GetMapNameEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMapNameEventsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
