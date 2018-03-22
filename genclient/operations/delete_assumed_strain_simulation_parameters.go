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

// NewDeleteAssumedStrainSimulationParams creates a new DeleteAssumedStrainSimulationParams object
// with the default values initialized.
func NewDeleteAssumedStrainSimulationParams() *DeleteAssumedStrainSimulationParams {
	var ()
	return &DeleteAssumedStrainSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAssumedStrainSimulationParamsWithTimeout creates a new DeleteAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAssumedStrainSimulationParamsWithTimeout(timeout time.Duration) *DeleteAssumedStrainSimulationParams {
	var ()
	return &DeleteAssumedStrainSimulationParams{

		timeout: timeout,
	}
}

// NewDeleteAssumedStrainSimulationParamsWithContext creates a new DeleteAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAssumedStrainSimulationParamsWithContext(ctx context.Context) *DeleteAssumedStrainSimulationParams {
	var ()
	return &DeleteAssumedStrainSimulationParams{

		Context: ctx,
	}
}

// NewDeleteAssumedStrainSimulationParamsWithHTTPClient creates a new DeleteAssumedStrainSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAssumedStrainSimulationParamsWithHTTPClient(client *http.Client) *DeleteAssumedStrainSimulationParams {
	var ()
	return &DeleteAssumedStrainSimulationParams{
		HTTPClient: client,
	}
}

/*DeleteAssumedStrainSimulationParams contains all the parameters to send to the API endpoint
for the delete assumed strain simulation operation typically these are written to a http.Request
*/
type DeleteAssumedStrainSimulationParams struct {

	/*ID
	  ID of simulation

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) WithTimeout(timeout time.Duration) *DeleteAssumedStrainSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) WithContext(ctx context.Context) *DeleteAssumedStrainSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) WithHTTPClient(client *http.Client) *DeleteAssumedStrainSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) WithID(id int32) *DeleteAssumedStrainSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete assumed strain simulation params
func (o *DeleteAssumedStrainSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAssumedStrainSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
