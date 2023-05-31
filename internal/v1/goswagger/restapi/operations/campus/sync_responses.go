// Code generated by go-swagger; DO NOT EDIT.

package campus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
)

// SyncNoContentCode is the HTTP code returned for type SyncNoContent
const SyncNoContentCode int = 204

/*
SyncNoContent No content

swagger:response syncNoContent
*/
type SyncNoContent struct {
}

// NewSyncNoContent creates SyncNoContent with default headers values
func NewSyncNoContent() *SyncNoContent {

	return &SyncNoContent{}
}

// WriteResponse to the client
func (o *SyncNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// SyncInternalServerErrorCode is the HTTP code returned for type SyncInternalServerError
const SyncInternalServerErrorCode int = 500

/*
SyncInternalServerError Internal Server Error

swagger:response syncInternalServerError
*/
type SyncInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSyncInternalServerError creates SyncInternalServerError with default headers values
func NewSyncInternalServerError() *SyncInternalServerError {

	return &SyncInternalServerError{}
}

// WithPayload adds the payload to the sync internal server error response
func (o *SyncInternalServerError) WithPayload(payload *models.Error) *SyncInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sync internal server error response
func (o *SyncInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SyncInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
