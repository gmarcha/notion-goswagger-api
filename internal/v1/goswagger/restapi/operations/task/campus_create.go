// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CampusCreateHandlerFunc turns a function with the right signature into a campus create handler
type CampusCreateHandlerFunc func(CampusCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CampusCreateHandlerFunc) Handle(params CampusCreateParams) middleware.Responder {
	return fn(params)
}

// CampusCreateHandler interface for that can handle valid campus create params
type CampusCreateHandler interface {
	Handle(CampusCreateParams) middleware.Responder
}

// NewCampusCreate creates a new http.Handler for the campus create operation
func NewCampusCreate(ctx *middleware.Context, handler CampusCreateHandler) *CampusCreate {
	return &CampusCreate{Context: ctx, Handler: handler}
}

/*
	CampusCreate swagger:route POST /tasks/campuses/{id} Task campusCreate

# Create campus tasks

Create campus tasks like onboarding steps, etc..
*/
type CampusCreate struct {
	Context *middleware.Context
	Handler CampusCreateHandler
}

func (o *CampusCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCampusCreateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
