// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewCacheServiceMetricsRowHitRateGetParams creates a new CacheServiceMetricsRowHitRateGetParams object
// with the default values initialized.
func NewCacheServiceMetricsRowHitRateGetParams() *CacheServiceMetricsRowHitRateGetParams {

	return &CacheServiceMetricsRowHitRateGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowHitRateGetParamsWithTimeout creates a new CacheServiceMetricsRowHitRateGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsRowHitRateGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowHitRateGetParams {

	return &CacheServiceMetricsRowHitRateGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowHitRateGetParamsWithContext creates a new CacheServiceMetricsRowHitRateGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsRowHitRateGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowHitRateGetParams {

	return &CacheServiceMetricsRowHitRateGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsRowHitRateGetParamsWithHTTPClient creates a new CacheServiceMetricsRowHitRateGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsRowHitRateGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowHitRateGetParams {

	return &CacheServiceMetricsRowHitRateGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsRowHitRateGetParams contains all the parameters to send to the API endpoint
for the cache service metrics row hit rate get operation typically these are written to a http.Request
*/
type CacheServiceMetricsRowHitRateGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowHitRateGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowHitRateGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowHitRateGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row hit rate get params
func (o *CacheServiceMetricsRowHitRateGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowHitRateGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}