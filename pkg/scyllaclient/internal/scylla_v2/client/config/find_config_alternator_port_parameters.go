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

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigAlternatorPortParams creates a new FindConfigAlternatorPortParams object
// with the default values initialized.
func NewFindConfigAlternatorPortParams() *FindConfigAlternatorPortParams {

	return &FindConfigAlternatorPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAlternatorPortParamsWithTimeout creates a new FindConfigAlternatorPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigAlternatorPortParamsWithTimeout(timeout time.Duration) *FindConfigAlternatorPortParams {

	return &FindConfigAlternatorPortParams{

		timeout: timeout,
	}
}

// NewFindConfigAlternatorPortParamsWithContext creates a new FindConfigAlternatorPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigAlternatorPortParamsWithContext(ctx context.Context) *FindConfigAlternatorPortParams {

	return &FindConfigAlternatorPortParams{

		Context: ctx,
	}
}

// NewFindConfigAlternatorPortParamsWithHTTPClient creates a new FindConfigAlternatorPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigAlternatorPortParamsWithHTTPClient(client *http.Client) *FindConfigAlternatorPortParams {

	return &FindConfigAlternatorPortParams{
		HTTPClient: client,
	}
}

/*FindConfigAlternatorPortParams contains all the parameters to send to the API endpoint
for the find config alternator port operation typically these are written to a http.Request
*/
type FindConfigAlternatorPortParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config alternator port params
func (o *FindConfigAlternatorPortParams) WithTimeout(timeout time.Duration) *FindConfigAlternatorPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config alternator port params
func (o *FindConfigAlternatorPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config alternator port params
func (o *FindConfigAlternatorPortParams) WithContext(ctx context.Context) *FindConfigAlternatorPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config alternator port params
func (o *FindConfigAlternatorPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config alternator port params
func (o *FindConfigAlternatorPortParams) WithHTTPClient(client *http.Client) *FindConfigAlternatorPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config alternator port params
func (o *FindConfigAlternatorPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAlternatorPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
