// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VnicEthAdapterPolicy Ethernet Adapter
//
// An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface you can configure various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings.
//
// swagger:model vnicEthAdapterPolicy
type VnicEthAdapterPolicy struct {
	PolicyAbstractPolicy

	// Enables advanced filtering on the interface.
	//
	AdvancedFilter *bool `json:"AdvancedFilter,omitempty"`

	// Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency.
	//
	ArfsSettings *VnicArfsSettings `json:"ArfsSettings,omitempty"`

	// Completion Queue resource settings.
	//
	CompletionQueueSettings *VnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`

	// Interrupt Settings for the virtual ethernet interface.
	//
	InterruptSettings *VnicEthInterruptSettings `json:"InterruptSettings,omitempty"`

	// Network Virtualization using Generic Routing Encapsulation Settings.
	//
	NvgreSettings *VnicNvgreSettings `json:"NvgreSettings,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Settings for RDMA over Converged Ethernet.
	//
	RoceSettings *VnicRoceSettings `json:"RoceSettings,omitempty"`

	// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
	//
	RssSettings *bool `json:"RssSettings,omitempty"`

	// Receive Queue resouce settings.
	//
	RxQueueSettings *VnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`

	// The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not.
	//
	TCPOffloadSettings *VnicTCPOffloadSettings `json:"TcpOffloadSettings,omitempty"`

	// Transmit Queue resource settings.
	//
	TxQueueSettings *VnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`

	// Virtual Extensible LAN Protocol Settings.
	//
	VxlanSettings *VnicVxlanSettings `json:"VxlanSettings,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VnicEthAdapterPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractPolicy = aO0

	// AO1
	var dataAO1 struct {
		AdvancedFilter *bool `json:"AdvancedFilter,omitempty"`

		ArfsSettings *VnicArfsSettings `json:"ArfsSettings,omitempty"`

		CompletionQueueSettings *VnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`

		InterruptSettings *VnicEthInterruptSettings `json:"InterruptSettings,omitempty"`

		NvgreSettings *VnicNvgreSettings `json:"NvgreSettings,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		RoceSettings *VnicRoceSettings `json:"RoceSettings,omitempty"`

		RssSettings *bool `json:"RssSettings,omitempty"`

		RxQueueSettings *VnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`

		TCPOffloadSettings *VnicTCPOffloadSettings `json:"TcpOffloadSettings,omitempty"`

		TxQueueSettings *VnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`

		VxlanSettings *VnicVxlanSettings `json:"VxlanSettings,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdvancedFilter = dataAO1.AdvancedFilter

	m.ArfsSettings = dataAO1.ArfsSettings

	m.CompletionQueueSettings = dataAO1.CompletionQueueSettings

	m.InterruptSettings = dataAO1.InterruptSettings

	m.NvgreSettings = dataAO1.NvgreSettings

	m.Organization = dataAO1.Organization

	m.RoceSettings = dataAO1.RoceSettings

	m.RssSettings = dataAO1.RssSettings

	m.RxQueueSettings = dataAO1.RxQueueSettings

	m.TCPOffloadSettings = dataAO1.TCPOffloadSettings

	m.TxQueueSettings = dataAO1.TxQueueSettings

	m.VxlanSettings = dataAO1.VxlanSettings

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VnicEthAdapterPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AdvancedFilter *bool `json:"AdvancedFilter,omitempty"`

		ArfsSettings *VnicArfsSettings `json:"ArfsSettings,omitempty"`

		CompletionQueueSettings *VnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`

		InterruptSettings *VnicEthInterruptSettings `json:"InterruptSettings,omitempty"`

		NvgreSettings *VnicNvgreSettings `json:"NvgreSettings,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		RoceSettings *VnicRoceSettings `json:"RoceSettings,omitempty"`

		RssSettings *bool `json:"RssSettings,omitempty"`

		RxQueueSettings *VnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`

		TCPOffloadSettings *VnicTCPOffloadSettings `json:"TcpOffloadSettings,omitempty"`

		TxQueueSettings *VnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`

		VxlanSettings *VnicVxlanSettings `json:"VxlanSettings,omitempty"`
	}

	dataAO1.AdvancedFilter = m.AdvancedFilter

	dataAO1.ArfsSettings = m.ArfsSettings

	dataAO1.CompletionQueueSettings = m.CompletionQueueSettings

	dataAO1.InterruptSettings = m.InterruptSettings

	dataAO1.NvgreSettings = m.NvgreSettings

	dataAO1.Organization = m.Organization

	dataAO1.RoceSettings = m.RoceSettings

	dataAO1.RssSettings = m.RssSettings

	dataAO1.RxQueueSettings = m.RxQueueSettings

	dataAO1.TCPOffloadSettings = m.TCPOffloadSettings

	dataAO1.TxQueueSettings = m.TxQueueSettings

	dataAO1.VxlanSettings = m.VxlanSettings

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vnic eth adapter policy
func (m *VnicEthAdapterPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractPolicy
	if err := m.PolicyAbstractPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArfsSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterruptSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvgreSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoceSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRxQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPOffloadSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxQueueSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVxlanSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VnicEthAdapterPolicy) validateArfsSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ArfsSettings) { // not required
		return nil
	}

	if m.ArfsSettings != nil {
		if err := m.ArfsSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ArfsSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateCompletionQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionQueueSettings) { // not required
		return nil
	}

	if m.CompletionQueueSettings != nil {
		if err := m.CompletionQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CompletionQueueSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateInterruptSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.InterruptSettings) { // not required
		return nil
	}

	if m.InterruptSettings != nil {
		if err := m.InterruptSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InterruptSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateNvgreSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NvgreSettings) { // not required
		return nil
	}

	if m.NvgreSettings != nil {
		if err := m.NvgreSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NvgreSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Organization")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateRoceSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.RoceSettings) { // not required
		return nil
	}

	if m.RoceSettings != nil {
		if err := m.RoceSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RoceSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateRxQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.RxQueueSettings) { // not required
		return nil
	}

	if m.RxQueueSettings != nil {
		if err := m.RxQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RxQueueSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateTCPOffloadSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.TCPOffloadSettings) { // not required
		return nil
	}

	if m.TCPOffloadSettings != nil {
		if err := m.TCPOffloadSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TcpOffloadSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateTxQueueSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.TxQueueSettings) { // not required
		return nil
	}

	if m.TxQueueSettings != nil {
		if err := m.TxQueueSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TxQueueSettings")
			}
			return err
		}
	}

	return nil
}

func (m *VnicEthAdapterPolicy) validateVxlanSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.VxlanSettings) { // not required
		return nil
	}

	if m.VxlanSettings != nil {
		if err := m.VxlanSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("VxlanSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VnicEthAdapterPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VnicEthAdapterPolicy) UnmarshalBinary(b []byte) error {
	var res VnicEthAdapterPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
