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

// NewGetScanPatternSimulationParams creates a new GetScanPatternSimulationParams object
// with the default values initialized.
func NewGetScanPatternSimulationParams() *GetScanPatternSimulationParams {
	var ()
	return &GetScanPatternSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanPatternSimulationParamsWithTimeout creates a new GetScanPatternSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScanPatternSimulationParamsWithTimeout(timeout time.Duration) *GetScanPatternSimulationParams {
	var ()
	return &GetScanPatternSimulationParams{

		timeout: timeout,
	}
}

// NewGetScanPatternSimulationParamsWithContext creates a new GetScanPatternSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScanPatternSimulationParamsWithContext(ctx context.Context) *GetScanPatternSimulationParams {
	var ()
	return &GetScanPatternSimulationParams{

		Context: ctx,
	}
}

/*GetScanPatternSimulationParams contains all the parameters to send to the API endpoint
for the get scan pattern simulation operation typically these are written to a http.Request
*/
type GetScanPatternSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) WithTimeout(timeout time.Duration) *GetScanPatternSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) WithContext(ctx context.Context) *GetScanPatternSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) WithID(id int32) *GetScanPatternSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scan pattern simulation params
func (o *GetScanPatternSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanPatternSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
