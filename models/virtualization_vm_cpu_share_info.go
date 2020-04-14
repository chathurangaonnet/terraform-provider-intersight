// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualizationVMCPUShareInfo Virtualization:Vm Cpu Share Info
//
// Information about the virtual machine's CPU sharing and limits.
//
// swagger:model virtualizationVmCpuShareInfo
type VirtualizationVMCPUShareInfo struct {
	MoBaseComplexType

	VirtualizationVMCPUShareInfoAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VirtualizationVMCPUShareInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 VirtualizationVMCPUShareInfoAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VirtualizationVMCPUShareInfoAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VirtualizationVMCPUShareInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VirtualizationVMCPUShareInfoAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this virtualization Vm Cpu share info
func (m *VirtualizationVMCPUShareInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VirtualizationVMCPUShareInfoAO1P1
	if err := m.VirtualizationVMCPUShareInfoAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationVMCPUShareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationVMCPUShareInfo) UnmarshalBinary(b []byte) error {
	var res VirtualizationVMCPUShareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualizationVMCPUShareInfoAO1P1 virtualization VM CPU share info a o1 p1
//
// swagger:model VirtualizationVMCPUShareInfoAO1P1
type VirtualizationVMCPUShareInfoAO1P1 struct {

	// Upper limit on CPU allocation (MHz).
	CPULimit int64 `json:"CpuLimit,omitempty"`

	// Amount of CPU for virtualization overhead.
	CPUOverheadLimit int64 `json:"CpuOverheadLimit,omitempty"`

	// Guaranteed minimum allocation of CPU resource (MHz).
	CPUReservation int64 `json:"CpuReservation,omitempty"`

	// Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.
	CPUShares int64 `json:"CpuShares,omitempty"`

	// virtualization VM CPU share info a o1 p1
	VirtualizationVMCPUShareInfoAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *VirtualizationVMCPUShareInfoAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Upper limit on CPU allocation (MHz).
		CPULimit int64 `json:"CpuLimit,omitempty"`

		// Amount of CPU for virtualization overhead.
		CPUOverheadLimit int64 `json:"CpuOverheadLimit,omitempty"`

		// Guaranteed minimum allocation of CPU resource (MHz).
		CPUReservation int64 `json:"CpuReservation,omitempty"`

		// Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.
		CPUShares int64 `json:"CpuShares,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VirtualizationVMCPUShareInfoAO1P1

	rcv.CPULimit = stage1.CPULimit
	rcv.CPUOverheadLimit = stage1.CPUOverheadLimit
	rcv.CPUReservation = stage1.CPUReservation
	rcv.CPUShares = stage1.CPUShares
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "CpuLimit")
	delete(stage2, "CpuOverheadLimit")
	delete(stage2, "CpuReservation")
	delete(stage2, "CpuShares")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.VirtualizationVMCPUShareInfoAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m VirtualizationVMCPUShareInfoAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Upper limit on CPU allocation (MHz).
		CPULimit int64 `json:"CpuLimit,omitempty"`

		// Amount of CPU for virtualization overhead.
		CPUOverheadLimit int64 `json:"CpuOverheadLimit,omitempty"`

		// Guaranteed minimum allocation of CPU resource (MHz).
		CPUReservation int64 `json:"CpuReservation,omitempty"`

		// Shows the relative importance of a VM. There is no unit for this value. It is a relative measure based on the settings for other resource pools. For more information, see VMware documentation.
		CPUShares int64 `json:"CpuShares,omitempty"`
	}

	stage1.CPULimit = m.CPULimit
	stage1.CPUOverheadLimit = m.CPUOverheadLimit
	stage1.CPUReservation = m.CPUReservation
	stage1.CPUShares = m.CPUShares

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.VirtualizationVMCPUShareInfoAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.VirtualizationVMCPUShareInfoAO1P1)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this virtualization VM CPU share info a o1 p1
func (m *VirtualizationVMCPUShareInfoAO1P1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualizationVMCPUShareInfoAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualizationVMCPUShareInfoAO1P1) UnmarshalBinary(b []byte) error {
	var res VirtualizationVMCPUShareInfoAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}