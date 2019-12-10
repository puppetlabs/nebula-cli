// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProfilePreferencesParams creates a new GetProfilePreferencesParams object
// with the default values initialized.
func NewGetProfilePreferencesParams() *GetProfilePreferencesParams {

	return &GetProfilePreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProfilePreferencesParamsWithTimeout creates a new GetProfilePreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProfilePreferencesParamsWithTimeout(timeout time.Duration) *GetProfilePreferencesParams {

	return &GetProfilePreferencesParams{

		timeout: timeout,
	}
}

// NewGetProfilePreferencesParamsWithContext creates a new GetProfilePreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProfilePreferencesParamsWithContext(ctx context.Context) *GetProfilePreferencesParams {

	return &GetProfilePreferencesParams{

		Context: ctx,
	}
}

// NewGetProfilePreferencesParamsWithHTTPClient creates a new GetProfilePreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProfilePreferencesParamsWithHTTPClient(client *http.Client) *GetProfilePreferencesParams {

	return &GetProfilePreferencesParams{
		HTTPClient: client,
	}
}

/*GetProfilePreferencesParams contains all the parameters to send to the API endpoint
for the get profile preferences operation typically these are written to a http.Request
*/
type GetProfilePreferencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get profile preferences params
func (o *GetProfilePreferencesParams) WithTimeout(timeout time.Duration) *GetProfilePreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get profile preferences params
func (o *GetProfilePreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get profile preferences params
func (o *GetProfilePreferencesParams) WithContext(ctx context.Context) *GetProfilePreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get profile preferences params
func (o *GetProfilePreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get profile preferences params
func (o *GetProfilePreferencesParams) WithHTTPClient(client *http.Client) *GetProfilePreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get profile preferences params
func (o *GetProfilePreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilePreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
