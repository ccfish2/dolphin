// Code generated by go-swagger; DO NOT EDIT.

package statedb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetStatedbDumpOKCode is the HTTP code returned for type GetStatedbDumpOK
const GetStatedbDumpOKCode int = 200

/*
GetStatedbDumpOK Success

swagger:response getStatedbDumpOK
*/
type GetStatedbDumpOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetStatedbDumpOK creates GetStatedbDumpOK with default headers values
func NewGetStatedbDumpOK() *GetStatedbDumpOK {

	return &GetStatedbDumpOK{}
}

// WithPayload adds the payload to the get statedb dump o k response
func (o *GetStatedbDumpOK) WithPayload(payload io.ReadCloser) *GetStatedbDumpOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get statedb dump o k response
func (o *GetStatedbDumpOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStatedbDumpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
