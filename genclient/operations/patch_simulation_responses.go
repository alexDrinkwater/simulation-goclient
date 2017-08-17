// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// PatchSimulationReader is a Reader for the PatchSimulation structure.
type PatchSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSimulationOK creates a PatchSimulationOK with default headers values
func NewPatchSimulationOK() *PatchSimulationOK {
	return &PatchSimulationOK{}
}

/*PatchSimulationOK handles this case with default header values.

Successfully patched simulation
*/
type PatchSimulationOK struct {
}

func (o *PatchSimulationOK) Error() string {
	return fmt.Sprintf("[PATCH /simulations/{id}][%d] patchSimulationOK ", 200)
}

func (o *PatchSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSimulationUnauthorized creates a PatchSimulationUnauthorized with default headers values
func NewPatchSimulationUnauthorized() *PatchSimulationUnauthorized {
	return &PatchSimulationUnauthorized{}
}

/*PatchSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PatchSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PatchSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /simulations/{id}][%d] patchSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationForbidden creates a PatchSimulationForbidden with default headers values
func NewPatchSimulationForbidden() *PatchSimulationForbidden {
	return &PatchSimulationForbidden{}
}

/*PatchSimulationForbidden handles this case with default header values.

Forbidden
*/
type PatchSimulationForbidden struct {
	Payload *models.Error
}

func (o *PatchSimulationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /simulations/{id}][%d] patchSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PatchSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationNotFound creates a PatchSimulationNotFound with default headers values
func NewPatchSimulationNotFound() *PatchSimulationNotFound {
	return &PatchSimulationNotFound{}
}

/*PatchSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type PatchSimulationNotFound struct {
	Payload *models.Error
}

func (o *PatchSimulationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /simulations/{id}][%d] patchSimulationNotFound  %+v", 404, o.Payload)
}

func (o *PatchSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSimulationDefault creates a PatchSimulationDefault with default headers values
func NewPatchSimulationDefault(code int) *PatchSimulationDefault {
	return &PatchSimulationDefault{
		_statusCode: code,
	}
}

/*PatchSimulationDefault handles this case with default header values.

unexpected error
*/
type PatchSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch simulation default response
func (o *PatchSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PatchSimulationDefault) Error() string {
	return fmt.Sprintf("[PATCH /simulations/{id}][%d] patchSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
