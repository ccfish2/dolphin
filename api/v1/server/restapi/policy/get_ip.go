// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIPHandlerFunc turns a function with the right signature into a get IP handler
type GetIPHandlerFunc func(GetIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIPHandlerFunc) Handle(params GetIPParams) middleware.Responder {
	return fn(params)
}

// GetIPHandler interface for that can handle valid get IP params
type GetIPHandler interface {
	Handle(GetIPParams) middleware.Responder
}

// NewGetIP creates a new http.Handler for the get IP operation
func NewGetIP(ctx *middleware.Context, handler GetIPHandler) *GetIP {
	return &GetIP{Context: ctx, Handler: handler}
}

/*
	GetIP swagger:route GET /ip policy getIp

# Lists information about known IP addresses

Retrieves a list of IPs with known associated information such as
their identities, host addresses, Kubernetes pod names, etc.
The list can optionally filtered by a CIDR IP range.
*/
type GetIP struct {
	Context *middleware.Context
	Handler GetIPHandler
}

func (o *GetIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetIPParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
