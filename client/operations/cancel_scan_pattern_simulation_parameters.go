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

// NewCancelScanPatternSimulationParams creates a new CancelScanPatternSimulationParams object
// with the default values initialized.
func NewCancelScanPatternSimulationParams() *CancelScanPatternSimulationParams {
	var ()
	return &CancelScanPatternSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelScanPatternSimulationParamsWithTimeout creates a new CancelScanPatternSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelScanPatternSimulationParamsWithTimeout(timeout time.Duration) *CancelScanPatternSimulationParams {
	var ()
	return &CancelScanPatternSimulationParams{

		timeout: timeout,
	}
}

// NewCancelScanPatternSimulationParamsWithContext creates a new CancelScanPatternSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelScanPatternSimulationParamsWithContext(ctx context.Context) *CancelScanPatternSimulationParams {
	var ()
	return &CancelScanPatternSimulationParams{

		Context: ctx,
	}
}

/*CancelScanPatternSimulationParams contains all the parameters to send to the API endpoint
for the cancel scan pattern simulation operation typically these are written to a http.Request
*/
type CancelScanPatternSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) WithTimeout(timeout time.Duration) *CancelScanPatternSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) WithContext(ctx context.Context) *CancelScanPatternSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) WithID(id int32) *CancelScanPatternSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the cancel scan pattern simulation params
func (o *CancelScanPatternSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CancelScanPatternSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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
