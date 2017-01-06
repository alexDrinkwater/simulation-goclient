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

// NewGetPorositySimulationsParams creates a new GetPorositySimulationsParams object
// with the default values initialized.
func NewGetPorositySimulationsParams() *GetPorositySimulationsParams {
	var ()
	return &GetPorositySimulationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPorositySimulationsParamsWithTimeout creates a new GetPorositySimulationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPorositySimulationsParamsWithTimeout(timeout time.Duration) *GetPorositySimulationsParams {
	var ()
	return &GetPorositySimulationsParams{

		timeout: timeout,
	}
}

// NewGetPorositySimulationsParamsWithContext creates a new GetPorositySimulationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPorositySimulationsParamsWithContext(ctx context.Context) *GetPorositySimulationsParams {
	var ()
	return &GetPorositySimulationsParams{

		Context: ctx,
	}
}

/*GetPorositySimulationsParams contains all the parameters to send to the API endpoint
for the get porosity simulations operation typically these are written to a http.Request
*/
type GetPorositySimulationsParams struct {

	/*Filter
	  Filter objects based on fields in the object.  E.g. filter=name:my-simulation,state:running

	*/
	Filter *string
	/*Limit
	  number of materials to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. offset of 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithTimeout(timeout time.Duration) *GetPorositySimulationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithContext(ctx context.Context) *GetPorositySimulationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithFilter adds the filter to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithFilter(filter *string) *GetPorositySimulationsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithLimit(limit *int32) *GetPorositySimulationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithOffset(offset *int32) *GetPorositySimulationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get porosity simulations params
func (o *GetPorositySimulationsParams) WithSort(sort []string) *GetPorositySimulationsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get porosity simulations params
func (o *GetPorositySimulationsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetPorositySimulationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}