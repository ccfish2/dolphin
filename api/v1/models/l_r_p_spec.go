// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LRPSpec Configuration of an LRP
//
// swagger:model LRPSpec
type LRPSpec struct {

	// mapping of frontends to pod backends
	FrontendMappings []*FrontendMapping `json:"frontend-mappings"`

	// LRP frontend type
	FrontendType string `json:"frontend-type,omitempty"`

	// LRP config type
	LrpType string `json:"lrp-type,omitempty"`

	// LRP service name
	Name string `json:"name,omitempty"`

	// LRP service namespace
	Namespace string `json:"namespace,omitempty"`

	// matching k8s service namespace and name
	ServiceID string `json:"service-id,omitempty"`

	// Unique identification
	UID string `json:"uid,omitempty"`
}

// Validate validates this l r p spec
func (m *LRPSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrontendMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LRPSpec) validateFrontendMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.FrontendMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.FrontendMappings); i++ {
		if swag.IsZero(m.FrontendMappings[i]) { // not required
			continue
		}

		if m.FrontendMappings[i] != nil {
			if err := m.FrontendMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frontend-mappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("frontend-mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this l r p spec based on the context it is used
func (m *LRPSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrontendMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LRPSpec) contextValidateFrontendMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FrontendMappings); i++ {

		if m.FrontendMappings[i] != nil {

			if swag.IsZero(m.FrontendMappings[i]) { // not required
				return nil
			}

			if err := m.FrontendMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frontend-mappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("frontend-mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LRPSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LRPSpec) UnmarshalBinary(b []byte) error {
	var res LRPSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
