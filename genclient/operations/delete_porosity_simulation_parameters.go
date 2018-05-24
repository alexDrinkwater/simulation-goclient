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

// NewDeletePorositySimulationParams creates a new DeletePorositySimulationParams object
// with the default values initialized.
func NewDeletePorositySimulationParams() *DeletePorositySimulationParams {
	var ()
	return &DeletePorositySimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePorositySimulationParamsWithTimeout creates a new DeletePorositySimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePorositySimulationParamsWithTimeout(timeout time.Duration) *DeletePorositySimulationParams {
	var ()
	return &DeletePorositySimulationParams{

		timeout: timeout,
	}
}

// NewDeletePorositySimulationParamsWithContext creates a new DeletePorositySimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePorositySimulationParamsWithContext(ctx context.Context) *DeletePorositySimulationParams {
	var ()
	return &DeletePorositySimulationParams{

		Context: ctx,
	}
}

// NewDeletePorositySimulationParamsWithHTTPClient creates a new DeletePorositySimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePorositySimulationParamsWithHTTPClient(client *http.Client) *DeletePorositySimulationParams {
	var ()
	return &DeletePorositySimulationParams{
		HTTPClient: client,
	}
}

/*DeletePorositySimulationParams contains all the parameters to send to the API endpoint
for the delete porosity simulation operation typically these are written to a http.Request
*/
type DeletePorositySimulationParams struct {

	/*ID
	  ID of simulation

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete porosity simulation params
func (o *DeletePorositySimulationParams) WithTimeout(timeout time.Duration) *DeletePorositySimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete porosity simulation params
func (o *DeletePorositySimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete porosity simulation params
func (o *DeletePorositySimulationParams) WithContext(ctx context.Context) *DeletePorositySimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete porosity simulation params
func (o *DeletePorositySimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete porosity simulation params
func (o *DeletePorositySimulationParams) WithHTTPClient(client *http.Client) *DeletePorositySimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete porosity simulation params
func (o *DeletePorositySimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete porosity simulation params
func (o *DeletePorositySimulationParams) WithID(id int32) *DeletePorositySimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete porosity simulation params
func (o *DeletePorositySimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePorositySimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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