// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

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

// NewListWorkflowRunsParams creates a new ListWorkflowRunsParams object
// with the default values initialized.
func NewListWorkflowRunsParams() *ListWorkflowRunsParams {
	var ()
	return &ListWorkflowRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListWorkflowRunsParamsWithTimeout creates a new ListWorkflowRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListWorkflowRunsParamsWithTimeout(timeout time.Duration) *ListWorkflowRunsParams {
	var ()
	return &ListWorkflowRunsParams{

		timeout: timeout,
	}
}

// NewListWorkflowRunsParamsWithContext creates a new ListWorkflowRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListWorkflowRunsParamsWithContext(ctx context.Context) *ListWorkflowRunsParams {
	var ()
	return &ListWorkflowRunsParams{

		Context: ctx,
	}
}

// NewListWorkflowRunsParamsWithHTTPClient creates a new ListWorkflowRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListWorkflowRunsParamsWithHTTPClient(client *http.Client) *ListWorkflowRunsParams {
	var ()
	return &ListWorkflowRunsParams{
		HTTPClient: client,
	}
}

/*ListWorkflowRunsParams contains all the parameters to send to the API endpoint
for the list workflow runs operation typically these are written to a http.Request
*/
type ListWorkflowRunsParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*WorkflowName
	  Workflow name. Must be unique within a user account

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list workflow runs params
func (o *ListWorkflowRunsParams) WithTimeout(timeout time.Duration) *ListWorkflowRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list workflow runs params
func (o *ListWorkflowRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list workflow runs params
func (o *ListWorkflowRunsParams) WithContext(ctx context.Context) *ListWorkflowRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list workflow runs params
func (o *ListWorkflowRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list workflow runs params
func (o *ListWorkflowRunsParams) WithHTTPClient(client *http.Client) *ListWorkflowRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list workflow runs params
func (o *ListWorkflowRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the list workflow runs params
func (o *ListWorkflowRunsParams) WithAccept(accept string) *ListWorkflowRunsParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the list workflow runs params
func (o *ListWorkflowRunsParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithWorkflowName adds the workflowName to the list workflow runs params
func (o *ListWorkflowRunsParams) WithWorkflowName(workflowName string) *ListWorkflowRunsParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the list workflow runs params
func (o *ListWorkflowRunsParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *ListWorkflowRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	// path param workflow_name
	if err := r.SetPathParam("workflow_name", o.WorkflowName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
