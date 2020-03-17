// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplianceImageBundle Appliance:Image Bundle
//
// ImageBundle keeps track of all the software bundles installed in the Intersight Appliances.
// Each ImageBundle managed object is derived from a software upgrade manifest. ImageBundle has
// additional properties computed during the manifest processing. Additional properties are the
// dynamic attributes of the software packages declared in the software manifest. For example,
// SHA256 values of the software packages are computed during the software manifest processing.
// An ImageBundle managed object named 'current' is always present in the Intersight Appliance.
// The software upgrade service creates another ImageBundle managed object named 'pending'
// when there is a pending software upgrade. The upgrade service renames the 'pending' bundle
// to the 'current' bundle after the software upgrade is successful.
//
// swagger:model applianceImageBundle
type ApplianceImageBundle struct {
	MoBaseMo

	// Collection of the Intersight Appliance's system installation packages.
	// Read Only: true
	AnsiblePackages []*OnpremImagePackage `json:"AnsiblePackages"`

	// Indicates that the software upgrade was automatically initiated by the Intersight Appliance.
	// Read Only: true
	AutoUpgrade *bool `json:"AutoUpgrade,omitempty"`

	// Collection of the Intersight Appliance's device connector packages.
	// Read Only: true
	DcPackages []*OnpremImagePackage `json:"DcPackages"`

	// Collection of the Intersight Appliance's developer debug packages. Optional, and not installed by default.
	// Read Only: true
	DebugPackages []*OnpremImagePackage `json:"DebugPackages"`

	// Short description of the software upgrade bundle.
	// Read Only: true
	Description string `json:"Description,omitempty"`

	// Collection of the Intersight Appliance's endpoint UI packages such as Cisco UCSM, Cisco HyperFlex etc.
	// Read Only: true
	EndpointPackages []*OnpremImagePackage `json:"EndpointPackages"`

	// Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm.
	// Read Only: true
	Fingerprint string `json:"Fingerprint,omitempty"`

	// Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing.
	// Read Only: true
	HasError *bool `json:"HasError,omitempty"`

	// Collection of the Intersight Appliance's infrastructure service packages such as database.
	// Read Only: true
	InfraPackages []*OnpremImagePackage `json:"InfraPackages"`

	// Collection of the Intersight Appliance's initialization service packages.
	// Read Only: true
	InitPackages []*OnpremImagePackage `json:"InitPackages"`

	// Name of the software upgrade bundle.
	// Read Only: true
	Name string `json:"Name,omitempty"`

	// Detailed description of the software upgrade bundle.
	// Read Only: true
	Notes string `json:"Notes,omitempty"`

	// Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.
	// Read Only: true
	// Enum: [Normal Critical]
	Priority string `json:"Priority,omitempty"`

	// Software upgrade manifest's release date and time.
	// Read Only: true
	// Format: date-time
	ReleaseTime strfmt.DateTime `json:"ReleaseTime,omitempty"`

	// Collection of the Intersight Appliance's micro-services pakages.
	// Read Only: true
	ServicePackages []*OnpremImagePackage `json:"ServicePackages"`

	// Status message set during the manifest processing.
	// Read Only: true
	StatusMessage string `json:"StatusMessage,omitempty"`

	// Collection of the Intersight Appliance's system packages such as DNS etc.
	// Read Only: true
	SystemPackages []*OnpremImagePackage `json:"SystemPackages"`

	// Collection of the Intersight Appliance's UI packages of the micro-services.
	// Read Only: true
	UIPackages []*OnpremImagePackage `json:"UiPackages"`

	// End date of the software upgrade process.
	// Read Only: true
	// Format: date-time
	UpgradeEndTime strfmt.DateTime `json:"UpgradeEndTime,omitempty"`

	// Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period.
	// Read Only: true
	UpgradeGracePeriod int64 `json:"UpgradeGracePeriod,omitempty"`

	// Duration (in minutes) for which services will be disrupted.
	// Read Only: true
	UpgradeImpactDuration int64 `json:"UpgradeImpactDuration,omitempty"`

	// UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance.
	// Read Only: true
	// Enum: [None Disruptive Disruptive-reboot]
	UpgradeImpactEnum string `json:"UpgradeImpactEnum,omitempty"`

	// Start date of the software upgrade process.
	// Read Only: true
	// Format: date-time
	UpgradeStartTime strfmt.DateTime `json:"UpgradeStartTime,omitempty"`

	// Software upgrade manifest's version.
	// Read Only: true
	Version string `json:"Version,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApplianceImageBundle) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		AnsiblePackages []*OnpremImagePackage `json:"AnsiblePackages"`

		AutoUpgrade *bool `json:"AutoUpgrade,omitempty"`

		DcPackages []*OnpremImagePackage `json:"DcPackages"`

		DebugPackages []*OnpremImagePackage `json:"DebugPackages"`

		Description string `json:"Description,omitempty"`

		EndpointPackages []*OnpremImagePackage `json:"EndpointPackages"`

		Fingerprint string `json:"Fingerprint,omitempty"`

		HasError *bool `json:"HasError,omitempty"`

		InfraPackages []*OnpremImagePackage `json:"InfraPackages"`

		InitPackages []*OnpremImagePackage `json:"InitPackages"`

		Name string `json:"Name,omitempty"`

		Notes string `json:"Notes,omitempty"`

		Priority string `json:"Priority,omitempty"`

		ReleaseTime strfmt.DateTime `json:"ReleaseTime,omitempty"`

		ServicePackages []*OnpremImagePackage `json:"ServicePackages"`

		StatusMessage string `json:"StatusMessage,omitempty"`

		SystemPackages []*OnpremImagePackage `json:"SystemPackages"`

		UIPackages []*OnpremImagePackage `json:"UiPackages"`

		UpgradeEndTime strfmt.DateTime `json:"UpgradeEndTime,omitempty"`

		UpgradeGracePeriod int64 `json:"UpgradeGracePeriod,omitempty"`

		UpgradeImpactDuration int64 `json:"UpgradeImpactDuration,omitempty"`

		UpgradeImpactEnum string `json:"UpgradeImpactEnum,omitempty"`

		UpgradeStartTime strfmt.DateTime `json:"UpgradeStartTime,omitempty"`

		Version string `json:"Version,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AnsiblePackages = dataAO1.AnsiblePackages

	m.AutoUpgrade = dataAO1.AutoUpgrade

	m.DcPackages = dataAO1.DcPackages

	m.DebugPackages = dataAO1.DebugPackages

	m.Description = dataAO1.Description

	m.EndpointPackages = dataAO1.EndpointPackages

	m.Fingerprint = dataAO1.Fingerprint

	m.HasError = dataAO1.HasError

	m.InfraPackages = dataAO1.InfraPackages

	m.InitPackages = dataAO1.InitPackages

	m.Name = dataAO1.Name

	m.Notes = dataAO1.Notes

	m.Priority = dataAO1.Priority

	m.ReleaseTime = dataAO1.ReleaseTime

	m.ServicePackages = dataAO1.ServicePackages

	m.StatusMessage = dataAO1.StatusMessage

	m.SystemPackages = dataAO1.SystemPackages

	m.UIPackages = dataAO1.UIPackages

	m.UpgradeEndTime = dataAO1.UpgradeEndTime

	m.UpgradeGracePeriod = dataAO1.UpgradeGracePeriod

	m.UpgradeImpactDuration = dataAO1.UpgradeImpactDuration

	m.UpgradeImpactEnum = dataAO1.UpgradeImpactEnum

	m.UpgradeStartTime = dataAO1.UpgradeStartTime

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApplianceImageBundle) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AnsiblePackages []*OnpremImagePackage `json:"AnsiblePackages"`

		AutoUpgrade *bool `json:"AutoUpgrade,omitempty"`

		DcPackages []*OnpremImagePackage `json:"DcPackages"`

		DebugPackages []*OnpremImagePackage `json:"DebugPackages"`

		Description string `json:"Description,omitempty"`

		EndpointPackages []*OnpremImagePackage `json:"EndpointPackages"`

		Fingerprint string `json:"Fingerprint,omitempty"`

		HasError *bool `json:"HasError,omitempty"`

		InfraPackages []*OnpremImagePackage `json:"InfraPackages"`

		InitPackages []*OnpremImagePackage `json:"InitPackages"`

		Name string `json:"Name,omitempty"`

		Notes string `json:"Notes,omitempty"`

		Priority string `json:"Priority,omitempty"`

		ReleaseTime strfmt.DateTime `json:"ReleaseTime,omitempty"`

		ServicePackages []*OnpremImagePackage `json:"ServicePackages"`

		StatusMessage string `json:"StatusMessage,omitempty"`

		SystemPackages []*OnpremImagePackage `json:"SystemPackages"`

		UIPackages []*OnpremImagePackage `json:"UiPackages"`

		UpgradeEndTime strfmt.DateTime `json:"UpgradeEndTime,omitempty"`

		UpgradeGracePeriod int64 `json:"UpgradeGracePeriod,omitempty"`

		UpgradeImpactDuration int64 `json:"UpgradeImpactDuration,omitempty"`

		UpgradeImpactEnum string `json:"UpgradeImpactEnum,omitempty"`

		UpgradeStartTime strfmt.DateTime `json:"UpgradeStartTime,omitempty"`

		Version string `json:"Version,omitempty"`
	}

	dataAO1.AnsiblePackages = m.AnsiblePackages

	dataAO1.AutoUpgrade = m.AutoUpgrade

	dataAO1.DcPackages = m.DcPackages

	dataAO1.DebugPackages = m.DebugPackages

	dataAO1.Description = m.Description

	dataAO1.EndpointPackages = m.EndpointPackages

	dataAO1.Fingerprint = m.Fingerprint

	dataAO1.HasError = m.HasError

	dataAO1.InfraPackages = m.InfraPackages

	dataAO1.InitPackages = m.InitPackages

	dataAO1.Name = m.Name

	dataAO1.Notes = m.Notes

	dataAO1.Priority = m.Priority

	dataAO1.ReleaseTime = m.ReleaseTime

	dataAO1.ServicePackages = m.ServicePackages

	dataAO1.StatusMessage = m.StatusMessage

	dataAO1.SystemPackages = m.SystemPackages

	dataAO1.UIPackages = m.UIPackages

	dataAO1.UpgradeEndTime = m.UpgradeEndTime

	dataAO1.UpgradeGracePeriod = m.UpgradeGracePeriod

	dataAO1.UpgradeImpactDuration = m.UpgradeImpactDuration

	dataAO1.UpgradeImpactEnum = m.UpgradeImpactEnum

	dataAO1.UpgradeStartTime = m.UpgradeStartTime

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this appliance image bundle
func (m *ApplianceImageBundle) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnsiblePackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDcPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebugPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfraPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUIPackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeImpactEnum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplianceImageBundle) validateAnsiblePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.AnsiblePackages) { // not required
		return nil
	}

	for i := 0; i < len(m.AnsiblePackages); i++ {
		if swag.IsZero(m.AnsiblePackages[i]) { // not required
			continue
		}

		if m.AnsiblePackages[i] != nil {
			if err := m.AnsiblePackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AnsiblePackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateDcPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.DcPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.DcPackages); i++ {
		if swag.IsZero(m.DcPackages[i]) { // not required
			continue
		}

		if m.DcPackages[i] != nil {
			if err := m.DcPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DcPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateDebugPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.DebugPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.DebugPackages); i++ {
		if swag.IsZero(m.DebugPackages[i]) { // not required
			continue
		}

		if m.DebugPackages[i] != nil {
			if err := m.DebugPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DebugPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateEndpointPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.EndpointPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.EndpointPackages); i++ {
		if swag.IsZero(m.EndpointPackages[i]) { // not required
			continue
		}

		if m.EndpointPackages[i] != nil {
			if err := m.EndpointPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndpointPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateInfraPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.InfraPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.InfraPackages); i++ {
		if swag.IsZero(m.InfraPackages[i]) { // not required
			continue
		}

		if m.InfraPackages[i] != nil {
			if err := m.InfraPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InfraPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateInitPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.InitPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.InitPackages); i++ {
		if swag.IsZero(m.InitPackages[i]) { // not required
			continue
		}

		if m.InitPackages[i] != nil {
			if err := m.InitPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InitPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var applianceImageBundleTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Normal","Critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceImageBundleTypePriorityPropEnum = append(applianceImageBundleTypePriorityPropEnum, v)
	}
}

// property enum
func (m *ApplianceImageBundle) validatePriorityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceImageBundleTypePriorityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceImageBundle) validatePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriorityEnum("Priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *ApplianceImageBundle) validateReleaseTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ReleaseTime", "body", "date-time", m.ReleaseTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApplianceImageBundle) validateServicePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePackages) { // not required
		return nil
	}

	for i := 0; i < len(m.ServicePackages); i++ {
		if swag.IsZero(m.ServicePackages[i]) { // not required
			continue
		}

		if m.ServicePackages[i] != nil {
			if err := m.ServicePackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ServicePackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateSystemPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.SystemPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemPackages); i++ {
		if swag.IsZero(m.SystemPackages[i]) { // not required
			continue
		}

		if m.SystemPackages[i] != nil {
			if err := m.SystemPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SystemPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateUIPackages(formats strfmt.Registry) error {

	if swag.IsZero(m.UIPackages) { // not required
		return nil
	}

	for i := 0; i < len(m.UIPackages); i++ {
		if swag.IsZero(m.UIPackages[i]) { // not required
			continue
		}

		if m.UIPackages[i] != nil {
			if err := m.UIPackages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UiPackages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplianceImageBundle) validateUpgradeEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpgradeEndTime", "body", "date-time", m.UpgradeEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var applianceImageBundleTypeUpgradeImpactEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Disruptive","Disruptive-reboot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applianceImageBundleTypeUpgradeImpactEnumPropEnum = append(applianceImageBundleTypeUpgradeImpactEnumPropEnum, v)
	}
}

// property enum
func (m *ApplianceImageBundle) validateUpgradeImpactEnumEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, applianceImageBundleTypeUpgradeImpactEnumPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApplianceImageBundle) validateUpgradeImpactEnum(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeImpactEnum) { // not required
		return nil
	}

	// value enum
	if err := m.validateUpgradeImpactEnumEnum("UpgradeImpactEnum", "body", m.UpgradeImpactEnum); err != nil {
		return err
	}

	return nil
}

func (m *ApplianceImageBundle) validateUpgradeStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradeStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("UpgradeStartTime", "body", "date-time", m.UpgradeStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplianceImageBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplianceImageBundle) UnmarshalBinary(b []byte) error {
	var res ApplianceImageBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
