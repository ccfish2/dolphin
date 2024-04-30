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

// DaemonConfigurationStatus Response to a daemon configuration request. Contains the addressing
// information, k8s, node monitor and immutable and mutable configuration
// settings.
//
// swagger:model DaemonConfigurationStatus
type DaemonConfigurationStatus struct {

	// Maximum IPv4 GRO size on workload facing devices
	GROIPV4MaxSize int64 `json:"GROIPv4MaxSize,omitempty"`

	// Maximum IPv6 GRO size on workload facing devices
	GROMaxSize int64 `json:"GROMaxSize,omitempty"`

	// Maximum IPv4 GSO size on workload facing devices
	GSOIPV4MaxSize int64 `json:"GSOIPv4MaxSize,omitempty"`

	// Maximum IPv6 GSO size on workload facing devices
	GSOMaxSize int64 `json:"GSOMaxSize,omitempty"`

	// addressing
	Addressing *NodeAddressing `json:"addressing,omitempty"`

	// Config map which contains all the active daemon configurations
	DaemonConfigurationMap map[string]interface{} `json:"daemonConfigurationMap,omitempty"`

	// datapath mode
	DatapathMode DatapathMode `json:"datapathMode,omitempty"`

	// MTU on workload facing devices
	DeviceMTU int64 `json:"deviceMTU,omitempty"`

	// Configured compatibility mode for --egress-multi-home-ip-rule-compat
	EgressMultiHomeIPRuleCompat bool `json:"egress-multi-home-ip-rule-compat,omitempty"`

	// Immutable configuration (read-only)
	Immutable ConfigurationMap `json:"immutable,omitempty"`

	// Configured IPAM mode
	IpamMode string `json:"ipam-mode,omitempty"`

	// k8s configuration
	K8sConfiguration string `json:"k8s-configuration,omitempty"`

	// k8s endpoint
	K8sEndpoint string `json:"k8s-endpoint,omitempty"`

	// kvstore configuration
	KvstoreConfiguration *KVstoreConfiguration `json:"kvstoreConfiguration,omitempty"`

	// masquerade
	Masquerade bool `json:"masquerade,omitempty"`

	// masquerade protocols
	MasqueradeProtocols *DaemonConfigurationStatusMasqueradeProtocols `json:"masqueradeProtocols,omitempty"`

	// Status of the node monitor
	NodeMonitor *MonitorStatus `json:"nodeMonitor,omitempty"`

	// Currently applied configuration
	Realized *DaemonConfigurationSpec `json:"realized,omitempty"`

	// MTU for network facing routes
	RouteMTU int64 `json:"routeMTU,omitempty"`
}

// Validate validates this daemon configuration status
func (m *DaemonConfigurationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatapathMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImmutable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKvstoreConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasqueradeProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DaemonConfigurationStatus) validateAddressing(formats strfmt.Registry) error {
	if swag.IsZero(m.Addressing) { // not required
		return nil
	}

	if m.Addressing != nil {
		if err := m.Addressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateDatapathMode(formats strfmt.Registry) error {
	if swag.IsZero(m.DatapathMode) { // not required
		return nil
	}

	if err := m.DatapathMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datapathMode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datapathMode")
		}
		return err
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateImmutable(formats strfmt.Registry) error {
	if swag.IsZero(m.Immutable) { // not required
		return nil
	}

	if m.Immutable != nil {
		if err := m.Immutable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("immutable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("immutable")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateKvstoreConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.KvstoreConfiguration) { // not required
		return nil
	}

	if m.KvstoreConfiguration != nil {
		if err := m.KvstoreConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstoreConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kvstoreConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateMasqueradeProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.MasqueradeProtocols) { // not required
		return nil
	}

	if m.MasqueradeProtocols != nil {
		if err := m.MasqueradeProtocols.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("masqueradeProtocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("masqueradeProtocols")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateNodeMonitor(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeMonitor) { // not required
		return nil
	}

	if m.NodeMonitor != nil {
		if err := m.NodeMonitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) validateRealized(formats strfmt.Registry) error {
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

// ContextValidate validate this daemon configuration status based on the context it is used
func (m *DaemonConfigurationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddressing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatapathMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImmutable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKvstoreConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMasqueradeProtocols(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeMonitor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRealized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DaemonConfigurationStatus) contextValidateAddressing(ctx context.Context, formats strfmt.Registry) error {

	if m.Addressing != nil {

		if swag.IsZero(m.Addressing) { // not required
			return nil
		}

		if err := m.Addressing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateDatapathMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DatapathMode) { // not required
		return nil
	}

	if err := m.DatapathMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("datapathMode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("datapathMode")
		}
		return err
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateImmutable(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Immutable) { // not required
		return nil
	}

	if err := m.Immutable.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("immutable")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("immutable")
		}
		return err
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateKvstoreConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.KvstoreConfiguration != nil {

		if swag.IsZero(m.KvstoreConfiguration) { // not required
			return nil
		}

		if err := m.KvstoreConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstoreConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kvstoreConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateMasqueradeProtocols(ctx context.Context, formats strfmt.Registry) error {

	if m.MasqueradeProtocols != nil {

		if swag.IsZero(m.MasqueradeProtocols) { // not required
			return nil
		}

		if err := m.MasqueradeProtocols.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("masqueradeProtocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("masqueradeProtocols")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateNodeMonitor(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeMonitor != nil {

		if swag.IsZero(m.NodeMonitor) { // not required
			return nil
		}

		if err := m.NodeMonitor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *DaemonConfigurationStatus) contextValidateRealized(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DaemonConfigurationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DaemonConfigurationStatus) UnmarshalBinary(b []byte) error {
	var res DaemonConfigurationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DaemonConfigurationStatusMasqueradeProtocols Status of masquerading feature
//
// swagger:model DaemonConfigurationStatusMasqueradeProtocols
type DaemonConfigurationStatusMasqueradeProtocols struct {

	// Status of masquerading for IPv4 traffic
	IPV4 bool `json:"ipv4,omitempty"`

	// Status of masquerading for IPv6 traffic
	IPV6 bool `json:"ipv6,omitempty"`
}

// Validate validates this daemon configuration status masquerade protocols
func (m *DaemonConfigurationStatusMasqueradeProtocols) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this daemon configuration status masquerade protocols based on context it is used
func (m *DaemonConfigurationStatusMasqueradeProtocols) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DaemonConfigurationStatusMasqueradeProtocols) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DaemonConfigurationStatusMasqueradeProtocols) UnmarshalBinary(b []byte) error {
	var res DaemonConfigurationStatusMasqueradeProtocols
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
