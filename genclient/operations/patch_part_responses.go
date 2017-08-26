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

// PatchPartReader is a Reader for the PatchPart structure.
type PatchPartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchPartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchPartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchPartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchPartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchPartOK creates a PatchPartOK with default headers values
func NewPatchPartOK() *PatchPartOK {
	return &PatchPartOK{}
}

/*PatchPartOK handles this case with default header values.

Patched part
*/
type PatchPartOK struct {
	Payload *models.Part
}

func (o *PatchPartOK) Error() string {
	return fmt.Sprintf("[PATCH /parts/{id}][%d] patchPartOK  %+v", 200, o.Payload)
}

func (o *PatchPartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Part)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPartUnauthorized creates a PatchPartUnauthorized with default headers values
func NewPatchPartUnauthorized() *PatchPartUnauthorized {
	return &PatchPartUnauthorized{}
}

/*PatchPartUnauthorized handles this case with default header values.

Not authorized
*/
type PatchPartUnauthorized struct {
	Payload *models.Error
}

func (o *PatchPartUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /parts/{id}][%d] patchPartUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchPartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPartForbidden creates a PatchPartForbidden with default headers values
func NewPatchPartForbidden() *PatchPartForbidden {
	return &PatchPartForbidden{}
}

/*PatchPartForbidden handles this case with default header values.

Forbidden
*/
type PatchPartForbidden struct {
	Payload *models.Error
}

func (o *PatchPartForbidden) Error() string {
	return fmt.Sprintf("[PATCH /parts/{id}][%d] patchPartForbidden  %+v", 403, o.Payload)
}

func (o *PatchPartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPartNotFound creates a PatchPartNotFound with default headers values
func NewPatchPartNotFound() *PatchPartNotFound {
	return &PatchPartNotFound{}
}

/*PatchPartNotFound handles this case with default header values.

Part not found (id invalid)
*/
type PatchPartNotFound struct {
	Payload *models.Error
}

func (o *PatchPartNotFound) Error() string {
	return fmt.Sprintf("[PATCH /parts/{id}][%d] patchPartNotFound  %+v", 404, o.Payload)
}

func (o *PatchPartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPartDefault creates a PatchPartDefault with default headers values
func NewPatchPartDefault(code int) *PatchPartDefault {
	return &PatchPartDefault{
		_statusCode: code,
	}
}

/*PatchPartDefault handles this case with default header values.

unexpected error
*/
type PatchPartDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch part default response
func (o *PatchPartDefault) Code() int {
	return o._statusCode
}

func (o *PatchPartDefault) Error() string {
	return fmt.Sprintf("[PATCH /parts/{id}][%d] patchPart default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
