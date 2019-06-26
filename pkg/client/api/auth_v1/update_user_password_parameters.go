// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

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

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// NewUpdateUserPasswordParams creates a new UpdateUserPasswordParams object
// with the default values initialized.
func NewUpdateUserPasswordParams() *UpdateUserPasswordParams {
	var ()
	return &UpdateUserPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserPasswordParamsWithTimeout creates a new UpdateUserPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserPasswordParamsWithTimeout(timeout time.Duration) *UpdateUserPasswordParams {
	var ()
	return &UpdateUserPasswordParams{

		timeout: timeout,
	}
}

// NewUpdateUserPasswordParamsWithContext creates a new UpdateUserPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserPasswordParamsWithContext(ctx context.Context) *UpdateUserPasswordParams {
	var ()
	return &UpdateUserPasswordParams{

		Context: ctx,
	}
}

// NewUpdateUserPasswordParamsWithHTTPClient creates a new UpdateUserPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserPasswordParamsWithHTTPClient(client *http.Client) *UpdateUserPasswordParams {
	var ()
	return &UpdateUserPasswordParams{
		HTTPClient: client,
	}
}

/*UpdateUserPasswordParams contains all the parameters to send to the API endpoint
for the update user password operation typically these are written to a http.Request
*/
type UpdateUserPasswordParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*Body
	  User password to update

	*/
	Body *models.UpdatePasswordSubmission
	/*ID
	  ID of the user to update

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user password params
func (o *UpdateUserPasswordParams) WithTimeout(timeout time.Duration) *UpdateUserPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user password params
func (o *UpdateUserPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user password params
func (o *UpdateUserPasswordParams) WithContext(ctx context.Context) *UpdateUserPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user password params
func (o *UpdateUserPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user password params
func (o *UpdateUserPasswordParams) WithHTTPClient(client *http.Client) *UpdateUserPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user password params
func (o *UpdateUserPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the update user password params
func (o *UpdateUserPasswordParams) WithAccept(accept string) *UpdateUserPasswordParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the update user password params
func (o *UpdateUserPasswordParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithBody adds the body to the update user password params
func (o *UpdateUserPasswordParams) WithBody(body *models.UpdatePasswordSubmission) *UpdateUserPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user password params
func (o *UpdateUserPasswordParams) SetBody(body *models.UpdatePasswordSubmission) {
	o.Body = body
}

// WithID adds the id to the update user password params
func (o *UpdateUserPasswordParams) WithID(id string) *UpdateUserPasswordParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update user password params
func (o *UpdateUserPasswordParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
