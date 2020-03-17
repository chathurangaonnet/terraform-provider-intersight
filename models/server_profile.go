// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerProfile Server Profile
//
// A profile specifying configuration settings for a physical server.
//
// swagger:model serverProfile
type ServerProfile struct {
	PolicyAbstractConfigProfile

	// The assigned physical server to the profile. A physical server can be assigned to more than one server profiles as long as it is only associated with one.
	AssignedServer *ComputeRackUnitRef `json:"AssignedServer,omitempty"`

	// The associated physical server to the profile. A physical server can only be associated to one server profiles.
	// Read Only: true
	AssociatedServer *ComputeRackUnitRef `json:"AssociatedServer,omitempty"`

	// The configuration change details are captured here.
	// Read Only: true
	ConfigChangeDetails []*ServerConfigChangeDetailRef `json:"ConfigChangeDetails"`

	// Pending configuration changes at the summary level. Detail changes are saved in configChangeDetails as a separate object.
	ConfigChanges *PolicyConfigChange `json:"ConfigChanges,omitempty"`

	// The configuration results including deploy, undeploy and compliance-check results. The errors usually are detected and reported during the apply/deploy phases.
	// Read Only: true
	ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`

	// is pmc deployed secure passphrase set
	IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy.
	PmcDeployedSecurePassphrase string `json:"PmcDeployedSecurePassphrase,omitempty"`

	// The WorkflowInfos in the workflow engine that are running for this server Profile.
	// Read Only: true
	RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ServerProfile) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PolicyAbstractConfigProfile
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PolicyAbstractConfigProfile = aO0

	// AO1
	var dataAO1 struct {
		AssignedServer *ComputeRackUnitRef `json:"AssignedServer,omitempty"`

		AssociatedServer *ComputeRackUnitRef `json:"AssociatedServer,omitempty"`

		ConfigChangeDetails []*ServerConfigChangeDetailRef `json:"ConfigChangeDetails"`

		ConfigChanges *PolicyConfigChange `json:"ConfigChanges,omitempty"`

		ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`

		IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PmcDeployedSecurePassphrase string `json:"PmcDeployedSecurePassphrase,omitempty"`

		RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AssignedServer = dataAO1.AssignedServer

	m.AssociatedServer = dataAO1.AssociatedServer

	m.ConfigChangeDetails = dataAO1.ConfigChangeDetails

	m.ConfigChanges = dataAO1.ConfigChanges

	m.ConfigResult = dataAO1.ConfigResult

	m.IsPmcDeployedSecurePassphraseSet = dataAO1.IsPmcDeployedSecurePassphraseSet

	m.Organization = dataAO1.Organization

	m.PmcDeployedSecurePassphrase = dataAO1.PmcDeployedSecurePassphrase

	m.RunningWorkflows = dataAO1.RunningWorkflows

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ServerProfile) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PolicyAbstractConfigProfile)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AssignedServer *ComputeRackUnitRef `json:"AssignedServer,omitempty"`

		AssociatedServer *ComputeRackUnitRef `json:"AssociatedServer,omitempty"`

		ConfigChangeDetails []*ServerConfigChangeDetailRef `json:"ConfigChangeDetails"`

		ConfigChanges *PolicyConfigChange `json:"ConfigChanges,omitempty"`

		ConfigResult *ServerConfigResultRef `json:"ConfigResult,omitempty"`

		IsPmcDeployedSecurePassphraseSet *bool `json:"IsPmcDeployedSecurePassphraseSet,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		PmcDeployedSecurePassphrase string `json:"PmcDeployedSecurePassphrase,omitempty"`

		RunningWorkflows []*WorkflowWorkflowInfoRef `json:"RunningWorkflows"`
	}

	dataAO1.AssignedServer = m.AssignedServer

	dataAO1.AssociatedServer = m.AssociatedServer

	dataAO1.ConfigChangeDetails = m.ConfigChangeDetails

	dataAO1.ConfigChanges = m.ConfigChanges

	dataAO1.ConfigResult = m.ConfigResult

	dataAO1.IsPmcDeployedSecurePassphraseSet = m.IsPmcDeployedSecurePassphraseSet

	dataAO1.Organization = m.Organization

	dataAO1.PmcDeployedSecurePassphrase = m.PmcDeployedSecurePassphrase

	dataAO1.RunningWorkflows = m.RunningWorkflows

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this server profile
func (m *ServerProfile) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PolicyAbstractConfigProfile
	if err := m.PolicyAbstractConfigProfile.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigChangeDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunningWorkflows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerProfile) validateAssignedServer(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedServer) { // not required
		return nil
	}

	if m.AssignedServer != nil {
		if err := m.AssignedServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AssignedServer")
			}
			return err
		}
	}

	return nil
}

func (m *ServerProfile) validateAssociatedServer(formats strfmt.Registry) error {

	if swag.IsZero(m.AssociatedServer) { // not required
		return nil
	}

	if m.AssociatedServer != nil {
		if err := m.AssociatedServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AssociatedServer")
			}
			return err
		}
	}

	return nil
}

func (m *ServerProfile) validateConfigChangeDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigChangeDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigChangeDetails); i++ {
		if swag.IsZero(m.ConfigChangeDetails[i]) { // not required
			continue
		}

		if m.ConfigChangeDetails[i] != nil {
			if err := m.ConfigChangeDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ConfigChangeDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerProfile) validateConfigChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigChanges) { // not required
		return nil
	}

	if m.ConfigChanges != nil {
		if err := m.ConfigChanges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigChanges")
			}
			return err
		}
	}

	return nil
}

func (m *ServerProfile) validateConfigResult(formats strfmt.Registry) error {

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

func (m *ServerProfile) validateOrganization(formats strfmt.Registry) error {

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

func (m *ServerProfile) validateRunningWorkflows(formats strfmt.Registry) error {

	if swag.IsZero(m.RunningWorkflows) { // not required
		return nil
	}

	for i := 0; i < len(m.RunningWorkflows); i++ {
		if swag.IsZero(m.RunningWorkflows[i]) { // not required
			continue
		}

		if m.RunningWorkflows[i] != nil {
			if err := m.RunningWorkflows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RunningWorkflows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerProfile) UnmarshalBinary(b []byte) error {
	var res ServerProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
