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

	"github.com/3dsim/simulation-goclient/models"
)

// NewPutThermalSimulationParams creates a new PutThermalSimulationParams object
// with the default values initialized.
func NewPutThermalSimulationParams() *PutThermalSimulationParams {
	var ()
	return &PutThermalSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutThermalSimulationParamsWithTimeout creates a new PutThermalSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutThermalSimulationParamsWithTimeout(timeout time.Duration) *PutThermalSimulationParams {
	var ()
	return &PutThermalSimulationParams{

		timeout: timeout,
	}
}

// NewPutThermalSimulationParamsWithContext creates a new PutThermalSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutThermalSimulationParamsWithContext(ctx context.Context) *PutThermalSimulationParams {
	var ()
	return &PutThermalSimulationParams{

		Context: ctx,
	}
}

// NewPutThermalSimulationParamsWithHTTPClient creates a new PutThermalSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutThermalSimulationParamsWithHTTPClient(client *http.Client) *PutThermalSimulationParams {
	var ()
	return &PutThermalSimulationParams{
		HTTPClient: client,
	}
}

/*PutThermalSimulationParams contains all the parameters to send to the API endpoint
for the put thermal simulation operation typically these are written to a http.Request
*/
type PutThermalSimulationParams struct {

	/*ThermalSimulation
	  ThermalSimulation fields required to update a simulation

	*/
	ThermalSimulation *models.ThermalSimulation
	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put thermal simulation params
func (o *PutThermalSimulationParams) WithTimeout(timeout time.Duration) *PutThermalSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put thermal simulation params
func (o *PutThermalSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put thermal simulation params
func (o *PutThermalSimulationParams) WithContext(ctx context.Context) *PutThermalSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put thermal simulation params
func (o *PutThermalSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put thermal simulation params
func (o *PutThermalSimulationParams) WithHTTPClient(client *http.Client) *PutThermalSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put thermal simulation params
func (o *PutThermalSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThermalSimulation adds the thermalSimulation to the put thermal simulation params
func (o *PutThermalSimulationParams) WithThermalSimulation(thermalSimulation *models.ThermalSimulation) *PutThermalSimulationParams {
	o.SetThermalSimulation(thermalSimulation)
	return o
}

// SetThermalSimulation adds the thermalSimulation to the put thermal simulation params
func (o *PutThermalSimulationParams) SetThermalSimulation(thermalSimulation *models.ThermalSimulation) {
	o.ThermalSimulation = thermalSimulation
}

// WithID adds the id to the put thermal simulation params
func (o *PutThermalSimulationParams) WithID(id int32) *PutThermalSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put thermal simulation params
func (o *PutThermalSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutThermalSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ThermalSimulation == nil {
		o.ThermalSimulation = new(models.ThermalSimulation)
	}

	if err := r.SetBodyParam(o.ThermalSimulation); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}