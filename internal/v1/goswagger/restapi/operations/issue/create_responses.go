// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
)

// CreateCreatedCode is the HTTP code returned for type CreateCreated
const CreateCreatedCode int = 201

/*
CreateCreated Created

swagger:response createCreated
*/
type CreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Issue `json:"body,omitempty"`
}

// NewCreateCreated creates CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {

	return &CreateCreated{}
}

// WithPayload adds the payload to the create created response
func (o *CreateCreated) WithPayload(payload *models.Issue) *CreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create created response
func (o *CreateCreated) SetPayload(payload *models.Issue) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateInternalServerErrorCode is the HTTP code returned for type CreateInternalServerError
const CreateInternalServerErrorCode int = 500

/*
CreateInternalServerError Internal Server Error

swagger:response createInternalServerError
*/
type CreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateInternalServerError creates CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {

	return &CreateInternalServerError{}
}

// WithPayload adds the payload to the create internal server error response
func (o *CreateInternalServerError) WithPayload(payload *models.Error) *CreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create internal server error response
func (o *CreateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
