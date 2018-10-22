// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterClusterIDStatusParams creates a new GetClusterClusterIDStatusParams object
// with the default values initialized.
func NewGetClusterClusterIDStatusParams() *GetClusterClusterIDStatusParams {
	var ()
	return &GetClusterClusterIDStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDStatusParamsWithTimeout creates a new GetClusterClusterIDStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDStatusParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDStatusParams {
	var ()
	return &GetClusterClusterIDStatusParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDStatusParamsWithContext creates a new GetClusterClusterIDStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDStatusParamsWithContext(ctx context.Context) *GetClusterClusterIDStatusParams {
	var ()
	return &GetClusterClusterIDStatusParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDStatusParamsWithHTTPClient creates a new GetClusterClusterIDStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDStatusParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDStatusParams {
	var ()
	return &GetClusterClusterIDStatusParams{
		HTTPClient: client,
	}
}

/*GetClusterClusterIDStatusParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID status operation typically these are written to a http.Request
*/
type GetClusterClusterIDStatusParams struct {

	/*ClusterID
	  cluster ID this API is performing on

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) WithContext(ctx context.Context) *GetClusterClusterIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) WithClusterID(clusterID string) *GetClusterClusterIDStatusParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID status params
func (o *GetClusterClusterIDStatusParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
