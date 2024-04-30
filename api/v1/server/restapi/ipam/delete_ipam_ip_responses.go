// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ccfish2/dolphin/api/v1/models"
)

// DeleteIpamIPOKCode is the HTTP code returned for type DeleteIpamIPOK
const DeleteIpamIPOKCode int = 200

/*
DeleteIpamIPOK Success

swagger:response deleteIpamIpOK
*/
type DeleteIpamIPOK struct {
}

// NewDeleteIpamIPOK creates DeleteIpamIPOK with default headers values
func NewDeleteIpamIPOK() *DeleteIpamIPOK {

	return &DeleteIpamIPOK{}
}

// WriteResponse to the client
func (o *DeleteIpamIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteIpamIPInvalidCode is the HTTP code returned for type DeleteIpamIPInvalid
const DeleteIpamIPInvalidCode int = 400

/*
DeleteIpamIPInvalid Invalid IP address

swagger:response deleteIpamIpInvalid
*/
type DeleteIpamIPInvalid struct {
}

// NewDeleteIpamIPInvalid creates DeleteIpamIPInvalid with default headers values
func NewDeleteIpamIPInvalid() *DeleteIpamIPInvalid {

	return &DeleteIpamIPInvalid{}
}

// WriteResponse to the client
func (o *DeleteIpamIPInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteIpamIPForbiddenCode is the HTTP code returned for type DeleteIpamIPForbidden
const DeleteIpamIPForbiddenCode int = 403

/*
DeleteIpamIPForbidden Forbidden

swagger:response deleteIpamIpForbidden
*/
type DeleteIpamIPForbidden struct {
}

// NewDeleteIpamIPForbidden creates DeleteIpamIPForbidden with default headers values
func NewDeleteIpamIPForbidden() *DeleteIpamIPForbidden {

	return &DeleteIpamIPForbidden{}
}

// WriteResponse to the client
func (o *DeleteIpamIPForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteIpamIPNotFoundCode is the HTTP code returned for type DeleteIpamIPNotFound
const DeleteIpamIPNotFoundCode int = 404

/*
DeleteIpamIPNotFound IP address not found

swagger:response deleteIpamIpNotFound
*/
type DeleteIpamIPNotFound struct {
}

// NewDeleteIpamIPNotFound creates DeleteIpamIPNotFound with default headers values
func NewDeleteIpamIPNotFound() *DeleteIpamIPNotFound {

	return &DeleteIpamIPNotFound{}
}

// WriteResponse to the client
func (o *DeleteIpamIPNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteIpamIPFailureCode is the HTTP code returned for type DeleteIpamIPFailure
const DeleteIpamIPFailureCode int = 500

/*
DeleteIpamIPFailure Address release failure

swagger:response deleteIpamIpFailure
*/
type DeleteIpamIPFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteIpamIPFailure creates DeleteIpamIPFailure with default headers values
func NewDeleteIpamIPFailure() *DeleteIpamIPFailure {

	return &DeleteIpamIPFailure{}
}

// WithPayload adds the payload to the delete ipam Ip failure response
func (o *DeleteIpamIPFailure) WithPayload(payload models.Error) *DeleteIpamIPFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete ipam Ip failure response
func (o *DeleteIpamIPFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIpamIPFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteIpamIPDisabledCode is the HTTP code returned for type DeleteIpamIPDisabled
const DeleteIpamIPDisabledCode int = 501

/*
DeleteIpamIPDisabled Allocation for address family disabled

swagger:response deleteIpamIpDisabled
*/
type DeleteIpamIPDisabled struct {
}

// NewDeleteIpamIPDisabled creates DeleteIpamIPDisabled with default headers values
func NewDeleteIpamIPDisabled() *DeleteIpamIPDisabled {

	return &DeleteIpamIPDisabled{}
}

// WriteResponse to the client
func (o *DeleteIpamIPDisabled) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(501)
}
