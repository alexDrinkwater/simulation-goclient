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

// GetMaterialsReader is a Reader for the GetMaterials structure.
type GetMaterialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMaterialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMaterialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetMaterialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMaterialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetMaterialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMaterialsOK creates a GetMaterialsOK with default headers values
func NewGetMaterialsOK() *GetMaterialsOK {
	return &GetMaterialsOK{}
}

/*GetMaterialsOK handles this case with default header values.

Successfully retrieved list
*/
type GetMaterialsOK struct {
	Payload []*models.Material
}

func (o *GetMaterialsOK) Error() string {
	return fmt.Sprintf("[GET /materials][%d] getMaterialsOK  %+v", 200, o.Payload)
}

func (o *GetMaterialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMaterialsUnauthorized creates a GetMaterialsUnauthorized with default headers values
func NewGetMaterialsUnauthorized() *GetMaterialsUnauthorized {
	return &GetMaterialsUnauthorized{}
}

/*GetMaterialsUnauthorized handles this case with default header values.

Not authorized
*/
type GetMaterialsUnauthorized struct {
	Payload *models.Error
}

func (o *GetMaterialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /materials][%d] getMaterialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMaterialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMaterialsForbidden creates a GetMaterialsForbidden with default headers values
func NewGetMaterialsForbidden() *GetMaterialsForbidden {
	return &GetMaterialsForbidden{}
}

/*GetMaterialsForbidden handles this case with default header values.

Forbidden
*/
type GetMaterialsForbidden struct {
	Payload *models.Error
}

func (o *GetMaterialsForbidden) Error() string {
	return fmt.Sprintf("[GET /materials][%d] getMaterialsForbidden  %+v", 403, o.Payload)
}

func (o *GetMaterialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMaterialsDefault creates a GetMaterialsDefault with default headers values
func NewGetMaterialsDefault(code int) *GetMaterialsDefault {
	return &GetMaterialsDefault{
		_statusCode: code,
	}
}

/*GetMaterialsDefault handles this case with default header values.

unexpected error
*/
type GetMaterialsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get materials default response
func (o *GetMaterialsDefault) Code() int {
	return o._statusCode
}

func (o *GetMaterialsDefault) Error() string {
	return fmt.Sprintf("[GET /materials][%d] getMaterials default  %+v", o._statusCode, o.Payload)
}

func (o *GetMaterialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
