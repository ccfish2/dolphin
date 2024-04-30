// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecorderStatus Configuration of a recorder
//
// swagger:model RecorderStatus
type RecorderStatus struct {

	// realized
	Realized *RecorderSpec `json:"realized,omitempty"`
}

// Validate validates this recorder status
func (m *RecorderStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecorderStatus) validateRealized(formats strfmt.Registry) error {
	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {
		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recorder status based on the context it is used
func (m *RecorderStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRealized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecorderStatus) contextValidateRealized(ctx context.Context, formats strfmt.Registry) error {

	if m.Realized != nil {

		if swag.IsZero(m.Realized) { // not required
			return nil
		}

		if err := m.Realized.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecorderStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecorderStatus) UnmarshalBinary(b []byte) error {
	var res RecorderStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
