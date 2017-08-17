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

// UnarchiveMachineReader is a Reader for the UnarchiveMachine structure.
type UnarchiveMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnarchiveMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUnarchiveMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUnarchiveMachineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUnarchiveMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUnarchiveMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUnarchiveMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnarchiveMachineOK creates a UnarchiveMachineOK with default headers values
func NewUnarchiveMachineOK() *UnarchiveMachineOK {
	return &UnarchiveMachineOK{}
}

/*UnarchiveMachineOK handles this case with default header values.

Machine was successfully unarchived.
*/
type UnarchiveMachineOK struct {
	Payload *models.Machine
}

func (o *UnarchiveMachineOK) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}/unarchive][%d] unarchiveMachineOK  %+v", 200, o.Payload)
}

func (o *UnarchiveMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Machine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnarchiveMachineUnauthorized creates a UnarchiveMachineUnauthorized with default headers values
func NewUnarchiveMachineUnauthorized() *UnarchiveMachineUnauthorized {
	return &UnarchiveMachineUnauthorized{}
}

/*UnarchiveMachineUnauthorized handles this case with default header values.

Not authorized
*/
type UnarchiveMachineUnauthorized struct {
	Payload *models.Error
}

func (o *UnarchiveMachineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}/unarchive][%d] unarchiveMachineUnauthorized  %+v", 401, o.Payload)
}

func (o *UnarchiveMachineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnarchiveMachineForbidden creates a UnarchiveMachineForbidden with default headers values
func NewUnarchiveMachineForbidden() *UnarchiveMachineForbidden {
	return &UnarchiveMachineForbidden{}
}

/*UnarchiveMachineForbidden handles this case with default header values.

Forbidden
*/
type UnarchiveMachineForbidden struct {
	Payload *models.Error
}

func (o *UnarchiveMachineForbidden) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}/unarchive][%d] unarchiveMachineForbidden  %+v", 403, o.Payload)
}

func (o *UnarchiveMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnarchiveMachineNotFound creates a UnarchiveMachineNotFound with default headers values
func NewUnarchiveMachineNotFound() *UnarchiveMachineNotFound {
	return &UnarchiveMachineNotFound{}
}

/*UnarchiveMachineNotFound handles this case with default header values.

Machine not found (id invalid)
*/
type UnarchiveMachineNotFound struct {
}

func (o *UnarchiveMachineNotFound) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}/unarchive][%d] unarchiveMachineNotFound ", 404)
}

func (o *UnarchiveMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnarchiveMachineDefault creates a UnarchiveMachineDefault with default headers values
func NewUnarchiveMachineDefault(code int) *UnarchiveMachineDefault {
	return &UnarchiveMachineDefault{
		_statusCode: code,
	}
}

/*UnarchiveMachineDefault handles this case with default header values.

unexpected error
*/
type UnarchiveMachineDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the unarchive machine default response
func (o *UnarchiveMachineDefault) Code() int {
	return o._statusCode
}

func (o *UnarchiveMachineDefault) Error() string {
	return fmt.Sprintf("[PUT /machines/{id}/unarchive][%d] unarchiveMachine default  %+v", o._statusCode, o.Payload)
}

func (o *UnarchiveMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
