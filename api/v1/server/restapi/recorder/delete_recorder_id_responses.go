// Code generated by go-swagger; DO NOT EDIT.

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ccfish2/dolphin/api/v1/models"
)

// DeleteRecorderIDOKCode is the HTTP code returned for type DeleteRecorderIDOK
const DeleteRecorderIDOKCode int = 200

/*
DeleteRecorderIDOK Success

swagger:response deleteRecorderIdOK
*/
type DeleteRecorderIDOK struct {
}

// NewDeleteRecorderIDOK creates DeleteRecorderIDOK with default headers values
func NewDeleteRecorderIDOK() *DeleteRecorderIDOK {

	return &DeleteRecorderIDOK{}
}

// WriteResponse to the client
func (o *DeleteRecorderIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteRecorderIDForbiddenCode is the HTTP code returned for type DeleteRecorderIDForbidden
const DeleteRecorderIDForbiddenCode int = 403

/*
DeleteRecorderIDForbidden Forbidden

swagger:response deleteRecorderIdForbidden
*/
type DeleteRecorderIDForbidden struct {
}

// NewDeleteRecorderIDForbidden creates DeleteRecorderIDForbidden with default headers values
func NewDeleteRecorderIDForbidden() *DeleteRecorderIDForbidden {

	return &DeleteRecorderIDForbidden{}
}

// WriteResponse to the client
func (o *DeleteRecorderIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteRecorderIDNotFoundCode is the HTTP code returned for type DeleteRecorderIDNotFound
const DeleteRecorderIDNotFoundCode int = 404

/*
DeleteRecorderIDNotFound Recorder not found

swagger:response deleteRecorderIdNotFound
*/
type DeleteRecorderIDNotFound struct {
}

// NewDeleteRecorderIDNotFound creates DeleteRecorderIDNotFound with default headers values
func NewDeleteRecorderIDNotFound() *DeleteRecorderIDNotFound {

	return &DeleteRecorderIDNotFound{}
}

// WriteResponse to the client
func (o *DeleteRecorderIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteRecorderIDFailureCode is the HTTP code returned for type DeleteRecorderIDFailure
const DeleteRecorderIDFailureCode int = 500

/*
DeleteRecorderIDFailure Recorder deletion failed

swagger:response deleteRecorderIdFailure
*/
type DeleteRecorderIDFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteRecorderIDFailure creates DeleteRecorderIDFailure with default headers values
func NewDeleteRecorderIDFailure() *DeleteRecorderIDFailure {

	return &DeleteRecorderIDFailure{}
}

// WithPayload adds the payload to the delete recorder Id failure response
func (o *DeleteRecorderIDFailure) WithPayload(payload models.Error) *DeleteRecorderIDFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete recorder Id failure response
func (o *DeleteRecorderIDFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRecorderIDFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
