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

// CancelAssumedStrainSimulationReader is a Reader for the CancelAssumedStrainSimulation structure.
type CancelAssumedStrainSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelAssumedStrainSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCancelAssumedStrainSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCancelAssumedStrainSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCancelAssumedStrainSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCancelAssumedStrainSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCancelAssumedStrainSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelAssumedStrainSimulationOK creates a CancelAssumedStrainSimulationOK with default headers values
func NewCancelAssumedStrainSimulationOK() *CancelAssumedStrainSimulationOK {
	return &CancelAssumedStrainSimulationOK{}
}

/*CancelAssumedStrainSimulationOK handles this case with default header values.

Simulation was successfully cancelled.
*/
type CancelAssumedStrainSimulationOK struct {
}

func (o *CancelAssumedStrainSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}/cancel][%d] cancelAssumedStrainSimulationOK ", 200)
}

func (o *CancelAssumedStrainSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelAssumedStrainSimulationUnauthorized creates a CancelAssumedStrainSimulationUnauthorized with default headers values
func NewCancelAssumedStrainSimulationUnauthorized() *CancelAssumedStrainSimulationUnauthorized {
	return &CancelAssumedStrainSimulationUnauthorized{}
}

/*CancelAssumedStrainSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type CancelAssumedStrainSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *CancelAssumedStrainSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}/cancel][%d] cancelAssumedStrainSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelAssumedStrainSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelAssumedStrainSimulationForbidden creates a CancelAssumedStrainSimulationForbidden with default headers values
func NewCancelAssumedStrainSimulationForbidden() *CancelAssumedStrainSimulationForbidden {
	return &CancelAssumedStrainSimulationForbidden{}
}

/*CancelAssumedStrainSimulationForbidden handles this case with default header values.

Forbidden
*/
type CancelAssumedStrainSimulationForbidden struct {
	Payload *models.Error
}

func (o *CancelAssumedStrainSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}/cancel][%d] cancelAssumedStrainSimulationForbidden  %+v", 403, o.Payload)
}

func (o *CancelAssumedStrainSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelAssumedStrainSimulationNotFound creates a CancelAssumedStrainSimulationNotFound with default headers values
func NewCancelAssumedStrainSimulationNotFound() *CancelAssumedStrainSimulationNotFound {
	return &CancelAssumedStrainSimulationNotFound{}
}

/*CancelAssumedStrainSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type CancelAssumedStrainSimulationNotFound struct {
	Payload *models.Error
}

func (o *CancelAssumedStrainSimulationNotFound) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}/cancel][%d] cancelAssumedStrainSimulationNotFound  %+v", 404, o.Payload)
}

func (o *CancelAssumedStrainSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelAssumedStrainSimulationDefault creates a CancelAssumedStrainSimulationDefault with default headers values
func NewCancelAssumedStrainSimulationDefault(code int) *CancelAssumedStrainSimulationDefault {
	return &CancelAssumedStrainSimulationDefault{
		_statusCode: code,
	}
}

/*CancelAssumedStrainSimulationDefault handles this case with default header values.

unexpected error
*/
type CancelAssumedStrainSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the cancel assumed strain simulation default response
func (o *CancelAssumedStrainSimulationDefault) Code() int {
	return o._statusCode
}

func (o *CancelAssumedStrainSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /assumedstrainsimulations/{id}/cancel][%d] cancelAssumedStrainSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *CancelAssumedStrainSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
