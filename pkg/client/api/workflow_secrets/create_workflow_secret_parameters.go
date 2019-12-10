// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets

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

// NewCreateWorkflowSecretParams creates a new CreateWorkflowSecretParams object
// with the default values initialized.
func NewCreateWorkflowSecretParams() *CreateWorkflowSecretParams {
	var ()
	return &CreateWorkflowSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkflowSecretParamsWithTimeout creates a new CreateWorkflowSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkflowSecretParamsWithTimeout(timeout time.Duration) *CreateWorkflowSecretParams {
	var ()
	return &CreateWorkflowSecretParams{

		timeout: timeout,
	}
}

// NewCreateWorkflowSecretParamsWithContext creates a new CreateWorkflowSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkflowSecretParamsWithContext(ctx context.Context) *CreateWorkflowSecretParams {
	var ()
	return &CreateWorkflowSecretParams{

		Context: ctx,
	}
}

// NewCreateWorkflowSecretParamsWithHTTPClient creates a new CreateWorkflowSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkflowSecretParamsWithHTTPClient(client *http.Client) *CreateWorkflowSecretParams {
	var ()
	return &CreateWorkflowSecretParams{
		HTTPClient: client,
	}
}

/*CreateWorkflowSecretParams contains all the parameters to send to the API endpoint
for the create workflow secret operation typically these are written to a http.Request
*/
type CreateWorkflowSecretParams struct {

	/*Body
	  Secret to name value pair to create

	*/
	Body CreateWorkflowSecretBody
	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create workflow secret params
func (o *CreateWorkflowSecretParams) WithTimeout(timeout time.Duration) *CreateWorkflowSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workflow secret params
func (o *CreateWorkflowSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workflow secret params
func (o *CreateWorkflowSecretParams) WithContext(ctx context.Context) *CreateWorkflowSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workflow secret params
func (o *CreateWorkflowSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workflow secret params
func (o *CreateWorkflowSecretParams) WithHTTPClient(client *http.Client) *CreateWorkflowSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workflow secret params
func (o *CreateWorkflowSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create workflow secret params
func (o *CreateWorkflowSecretParams) WithBody(body CreateWorkflowSecretBody) *CreateWorkflowSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create workflow secret params
func (o *CreateWorkflowSecretParams) SetBody(body CreateWorkflowSecretBody) {
	o.Body = body
}

// WithWorkflowName adds the workflowName to the create workflow secret params
func (o *CreateWorkflowSecretParams) WithWorkflowName(workflowName string) *CreateWorkflowSecretParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the create workflow secret params
func (o *CreateWorkflowSecretParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkflowSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
