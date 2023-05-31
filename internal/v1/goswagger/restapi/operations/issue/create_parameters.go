// Code generated by go-swagger; DO NOT EDIT.

package issue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
)

// NewCreateParams creates a new CreateParams object
//
// There are no default values defined in the spec.
func NewCreateParams() CreateParams {

	return CreateParams{}
}

// CreateParams contains all the bound params for the create operation
// typically these are obtained from a http.Request
//
// swagger:parameters create
type CreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Issue content.
	  Required: true
	  In: body
	*/
	Issue *models.Issue
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateParams() beforehand.
func (o *CreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Issue
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("issue", "body", ""))
			} else {
				res = append(res, errors.NewParseError("issue", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Issue = &body
			}
		}
	} else {
		res = append(res, errors.Required("issue", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}