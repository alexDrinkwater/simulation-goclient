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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSimulationParams creates a new GetSimulationParams object
// with the default values initialized.
func NewGetSimulationParams() *GetSimulationParams {
	var ()
	return &GetSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSimulationParamsWithTimeout creates a new GetSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSimulationParamsWithTimeout(timeout time.Duration) *GetSimulationParams {
	var ()
	return &GetSimulationParams{

		timeout: timeout,
	}
}

// NewGetSimulationParamsWithContext creates a new GetSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSimulationParamsWithContext(ctx context.Context) *GetSimulationParams {
	var ()
	return &GetSimulationParams{

		Context: ctx,
	}
}

// NewGetSimulationParamsWithHTTPClient creates a new GetSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSimulationParamsWithHTTPClient(client *http.Client) *GetSimulationParams {
	var ()
	return &GetSimulationParams{
		HTTPClient: client,
	}
}

/*GetSimulationParams contains all the parameters to send to the API endpoint
for the get simulation operation typically these are written to a http.Request
*/
type GetSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get simulation params
func (o *GetSimulationParams) WithTimeout(timeout time.Duration) *GetSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get simulation params
func (o *GetSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get simulation params
func (o *GetSimulationParams) WithContext(ctx context.Context) *GetSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get simulation params
func (o *GetSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get simulation params
func (o *GetSimulationParams) WithHTTPClient(client *http.Client) *GetSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get simulation params
func (o *GetSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get simulation params
func (o *GetSimulationParams) WithID(id int32) *GetSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get simulation params
func (o *GetSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
