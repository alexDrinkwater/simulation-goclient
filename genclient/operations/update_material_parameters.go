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

	"github.com/3dsim/simulation-goclient/models"
)

// NewUpdateMaterialParams creates a new UpdateMaterialParams object
// with the default values initialized.
func NewUpdateMaterialParams() *UpdateMaterialParams {
	var ()
	return &UpdateMaterialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMaterialParamsWithTimeout creates a new UpdateMaterialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMaterialParamsWithTimeout(timeout time.Duration) *UpdateMaterialParams {
	var ()
	return &UpdateMaterialParams{

		timeout: timeout,
	}
}

// NewUpdateMaterialParamsWithContext creates a new UpdateMaterialParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMaterialParamsWithContext(ctx context.Context) *UpdateMaterialParams {
	var ()
	return &UpdateMaterialParams{

		Context: ctx,
	}
}

// NewUpdateMaterialParamsWithHTTPClient creates a new UpdateMaterialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMaterialParamsWithHTTPClient(client *http.Client) *UpdateMaterialParams {
	var ()
	return &UpdateMaterialParams{
		HTTPClient: client,
	}
}

/*UpdateMaterialParams contains all the parameters to send to the API endpoint
for the update material operation typically these are written to a http.Request
*/
type UpdateMaterialParams struct {

	/*ID
	  material identifier

	*/
	ID int32
	/*Material*/
	Material *models.Material

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update material params
func (o *UpdateMaterialParams) WithTimeout(timeout time.Duration) *UpdateMaterialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update material params
func (o *UpdateMaterialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update material params
func (o *UpdateMaterialParams) WithContext(ctx context.Context) *UpdateMaterialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update material params
func (o *UpdateMaterialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update material params
func (o *UpdateMaterialParams) WithHTTPClient(client *http.Client) *UpdateMaterialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update material params
func (o *UpdateMaterialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update material params
func (o *UpdateMaterialParams) WithID(id int32) *UpdateMaterialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update material params
func (o *UpdateMaterialParams) SetID(id int32) {
	o.ID = id
}

// WithMaterial adds the material to the update material params
func (o *UpdateMaterialParams) WithMaterial(material *models.Material) *UpdateMaterialParams {
	o.SetMaterial(material)
	return o
}

// SetMaterial adds the material to the update material params
func (o *UpdateMaterialParams) SetMaterial(material *models.Material) {
	o.Material = material
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMaterialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Material == nil {
		o.Material = new(models.Material)
	}

	if err := r.SetBodyParam(o.Material); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
