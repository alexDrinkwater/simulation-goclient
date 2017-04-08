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

// GetPorositySimulationReader is a Reader for the GetPorositySimulation structure.
type GetPorositySimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPorositySimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPorositySimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPorositySimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPorositySimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPorositySimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPorositySimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPorositySimulationOK creates a GetPorositySimulationOK with default headers values
func NewGetPorositySimulationOK() *GetPorositySimulationOK {
	return &GetPorositySimulationOK{}
}

/*GetPorositySimulationOK handles this case with default header values.

Successfully retrieved simulation
*/
type GetPorositySimulationOK struct {
	Payload *models.PorositySimulation
}

func (o *GetPorositySimulationOK) Error() string {
	return fmt.Sprintf("[GET /porositysimulations/{id}][%d] getPorositySimulationOK  %+v", 200, o.Payload)
}

func (o *GetPorositySimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PorositySimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPorositySimulationUnauthorized creates a GetPorositySimulationUnauthorized with default headers values
func NewGetPorositySimulationUnauthorized() *GetPorositySimulationUnauthorized {
	return &GetPorositySimulationUnauthorized{}
}

/*GetPorositySimulationUnauthorized handles this case with default header values.

Not authorized
*/
type GetPorositySimulationUnauthorized struct {
	Payload *models.Error
}

func (o *GetPorositySimulationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /porositysimulations/{id}][%d] getPorositySimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPorositySimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPorositySimulationForbidden creates a GetPorositySimulationForbidden with default headers values
func NewGetPorositySimulationForbidden() *GetPorositySimulationForbidden {
	return &GetPorositySimulationForbidden{}
}

/*GetPorositySimulationForbidden handles this case with default header values.

Forbidden
*/
type GetPorositySimulationForbidden struct {
	Payload *models.Error
}

func (o *GetPorositySimulationForbidden) Error() string {
	return fmt.Sprintf("[GET /porositysimulations/{id}][%d] getPorositySimulationForbidden  %+v", 403, o.Payload)
}

func (o *GetPorositySimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPorositySimulationNotFound creates a GetPorositySimulationNotFound with default headers values
func NewGetPorositySimulationNotFound() *GetPorositySimulationNotFound {
	return &GetPorositySimulationNotFound{}
}

/*GetPorositySimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type GetPorositySimulationNotFound struct {
	Payload *models.Error
}

func (o *GetPorositySimulationNotFound) Error() string {
	return fmt.Sprintf("[GET /porositysimulations/{id}][%d] getPorositySimulationNotFound  %+v", 404, o.Payload)
}

func (o *GetPorositySimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPorositySimulationDefault creates a GetPorositySimulationDefault with default headers values
func NewGetPorositySimulationDefault(code int) *GetPorositySimulationDefault {
	return &GetPorositySimulationDefault{
		_statusCode: code,
	}
}

/*GetPorositySimulationDefault handles this case with default header values.

unexpected error
*/
type GetPorositySimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get porosity simulation default response
func (o *GetPorositySimulationDefault) Code() int {
	return o._statusCode
}

func (o *GetPorositySimulationDefault) Error() string {
	return fmt.Sprintf("[GET /porositysimulations/{id}][%d] getPorositySimulation default  %+v", o._statusCode, o.Payload)
}

func (o *GetPorositySimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}