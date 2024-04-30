// Code generated by go-swagger; DO NOT EDIT.

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBgpRoutesHandlerFunc turns a function with the right signature into a get bgp routes handler
type GetBgpRoutesHandlerFunc func(GetBgpRoutesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBgpRoutesHandlerFunc) Handle(params GetBgpRoutesParams) middleware.Responder {
	return fn(params)
}

// GetBgpRoutesHandler interface for that can handle valid get bgp routes params
type GetBgpRoutesHandler interface {
	Handle(GetBgpRoutesParams) middleware.Responder
}

// NewGetBgpRoutes creates a new http.Handler for the get bgp routes operation
func NewGetBgpRoutes(ctx *middleware.Context, handler GetBgpRoutesHandler) *GetBgpRoutes {
	return &GetBgpRoutes{Context: ctx, Handler: handler}
}

/*
	GetBgpRoutes swagger:route GET /bgp/routes bgp getBgpRoutes

Lists BGP routes from BGP Control Plane RIB.

Retrieves routes from BGP Control Plane RIB filtered by parameters you specify
*/
type GetBgpRoutes struct {
	Context *middleware.Context
	Handler GetBgpRoutesHandler
}

func (o *GetBgpRoutes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBgpRoutesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
