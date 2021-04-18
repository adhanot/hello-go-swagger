// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"hello/models"
)

// NewRegisterParams creates a new RegisterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterParams() *RegisterParams {
	return &RegisterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterParamsWithTimeout creates a new RegisterParams object
// with the ability to set a timeout on a request.
func NewRegisterParamsWithTimeout(timeout time.Duration) *RegisterParams {
	return &RegisterParams{
		timeout: timeout,
	}
}

// NewRegisterParamsWithContext creates a new RegisterParams object
// with the ability to set a context for a request.
func NewRegisterParamsWithContext(ctx context.Context) *RegisterParams {
	return &RegisterParams{
		Context: ctx,
	}
}

// NewRegisterParamsWithHTTPClient creates a new RegisterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterParamsWithHTTPClient(client *http.Client) *RegisterParams {
	return &RegisterParams{
		HTTPClient: client,
	}
}

/* RegisterParams contains all the parameters to send to the API endpoint
   for the register operation.

   Typically these are written to a http.Request.
*/
type RegisterParams struct {

	/* Signup.

	   Registeration Payload
	*/
	Signup *models.RegisterUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterParams) WithDefaults() *RegisterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register params
func (o *RegisterParams) WithTimeout(timeout time.Duration) *RegisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register params
func (o *RegisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register params
func (o *RegisterParams) WithContext(ctx context.Context) *RegisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register params
func (o *RegisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register params
func (o *RegisterParams) WithHTTPClient(client *http.Client) *RegisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register params
func (o *RegisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSignup adds the signup to the register params
func (o *RegisterParams) WithSignup(signup *models.RegisterUser) *RegisterParams {
	o.SetSignup(signup)
	return o
}

// SetSignup adds the signup to the register params
func (o *RegisterParams) SetSignup(signup *models.RegisterUser) {
	o.Signup = signup
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Signup != nil {
		if err := r.SetBodyParam(o.Signup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
