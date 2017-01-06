package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Triangle triangle
// swagger:model Triangle
type Triangle struct {

	// normal
	Normal *TriangleNormal `json:"normal,omitempty"`

	// vertices
	Vertices []*Vertex `json:"vertices"`
}

// Validate validates this triangle
func (m *Triangle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNormal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVertices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Triangle) validateNormal(formats strfmt.Registry) error {

	if swag.IsZero(m.Normal) { // not required
		return nil
	}

	if m.Normal != nil {

		if err := m.Normal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("normal")
			}
			return err
		}
	}

	return nil
}

func (m *Triangle) validateVertices(formats strfmt.Registry) error {

	if swag.IsZero(m.Vertices) { // not required
		return nil
	}

	for i := 0; i < len(m.Vertices); i++ {

		if swag.IsZero(m.Vertices[i]) { // not required
			continue
		}

		if m.Vertices[i] != nil {

			if err := m.Vertices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vertices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// TriangleNormal triangle normal
// swagger:model TriangleNormal
type TriangleNormal struct {

	// angle
	Angle float64 `json:"angle,omitempty"`

	// x
	X float64 `json:"x,omitempty"`

	// y
	Y float64 `json:"y,omitempty"`

	// z
	Z float64 `json:"z,omitempty"`
}

// Validate validates this triangle normal
func (m *TriangleNormal) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}