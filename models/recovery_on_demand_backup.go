// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecoveryOnDemandBackup On Demand Backup
//
// Handles requests for on demand backup for a given endpoint.
//
// swagger:model recoveryOnDemandBackup
type RecoveryOnDemandBackup struct {
	RecoveryAbstractBackupConfig

	// The status of ondemand backup.
	// Read Only: true
	ConfigResult *RecoveryConfigResultRef `json:"ConfigResult,omitempty"`

	// Relationship to the end point on which back up is configured.
	DeviceID *AssetDeviceRegistrationRef `json:"DeviceId,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoveryOnDemandBackup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RecoveryAbstractBackupConfig
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RecoveryAbstractBackupConfig = aO0

	// AO1
	var dataAO1 struct {
		ConfigResult *RecoveryConfigResultRef `json:"ConfigResult,omitempty"`

		DeviceID *AssetDeviceRegistrationRef `json:"DeviceId,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConfigResult = dataAO1.ConfigResult

	m.DeviceID = dataAO1.DeviceID

	m.Organization = dataAO1.Organization

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoveryOnDemandBackup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.RecoveryAbstractBackupConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ConfigResult *RecoveryConfigResultRef `json:"ConfigResult,omitempty"`

		DeviceID *AssetDeviceRegistrationRef `json:"DeviceId,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`
	}

	dataAO1.ConfigResult = m.ConfigResult

	dataAO1.DeviceID = m.DeviceID

	dataAO1.Organization = m.Organization

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recovery on demand backup
func (m *RecoveryOnDemandBackup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RecoveryAbstractBackupConfig
	if err := m.RecoveryAbstractBackupConfig.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryOnDemandBackup) validateConfigResult(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigResult) { // not required
		return nil
	}

	if m.ConfigResult != nil {
		if err := m.ConfigResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigResult")
			}
			return err
		}
	}

	return nil
}

func (m *RecoveryOnDemandBackup) validateDeviceID(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceID) { // not required
		return nil
	}

	if m.DeviceID != nil {
		if err := m.DeviceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeviceId")
			}
			return err
		}
	}

	return nil
}

func (m *RecoveryOnDemandBackup) validateOrganization(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RecoveryOnDemandBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryOnDemandBackup) UnmarshalBinary(b []byte) error {
	var res RecoveryOnDemandBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
