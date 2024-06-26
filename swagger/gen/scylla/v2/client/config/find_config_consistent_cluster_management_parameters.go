// Code generated by go-swagger; DO NOT EDIT.

package config

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
)

// NewFindConfigConsistentClusterManagementParams creates a new FindConfigConsistentClusterManagementParams object
// with the default values initialized.
func NewFindConfigConsistentClusterManagementParams() *FindConfigConsistentClusterManagementParams {

	return &FindConfigConsistentClusterManagementParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigConsistentClusterManagementParamsWithTimeout creates a new FindConfigConsistentClusterManagementParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigConsistentClusterManagementParamsWithTimeout(timeout time.Duration) *FindConfigConsistentClusterManagementParams {

	return &FindConfigConsistentClusterManagementParams{

		timeout: timeout,
	}
}

// NewFindConfigConsistentClusterManagementParamsWithContext creates a new FindConfigConsistentClusterManagementParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigConsistentClusterManagementParamsWithContext(ctx context.Context) *FindConfigConsistentClusterManagementParams {

	return &FindConfigConsistentClusterManagementParams{

		Context: ctx,
	}
}

// NewFindConfigConsistentClusterManagementParamsWithHTTPClient creates a new FindConfigConsistentClusterManagementParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigConsistentClusterManagementParamsWithHTTPClient(client *http.Client) *FindConfigConsistentClusterManagementParams {

	return &FindConfigConsistentClusterManagementParams{
		HTTPClient: client,
	}
}

/*
FindConfigConsistentClusterManagementParams contains all the parameters to send to the API endpoint
for the find config consistent cluster management operation typically these are written to a http.Request
*/
type FindConfigConsistentClusterManagementParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) WithTimeout(timeout time.Duration) *FindConfigConsistentClusterManagementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) WithContext(ctx context.Context) *FindConfigConsistentClusterManagementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) WithHTTPClient(client *http.Client) *FindConfigConsistentClusterManagementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config consistent cluster management params
func (o *FindConfigConsistentClusterManagementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigConsistentClusterManagementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
