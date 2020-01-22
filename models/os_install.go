// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OsInstall OS Install
//
// This MO models the target server, answers and other properties needed for
// OS installation. The OS installation can be started in the target server by doing
// a POST on this MO.
// The requests to this MO starts a OS installation workflow that can be tracked
// using workflow engine MO workflow.WorkflowInfo.
//
// swagger:model osInstall
type OsInstall struct {
	OsBaseInstallConfig

	// If the answers source is selected as 'Template' in 'Answers' property, this relation provides the os.ConfigurationFile instance to be used for this OS install.
	//
	ConfigurationFile *OsConfigurationFileRef `json:"ConfigurationFile,omitempty"`

	// OS Image to be installed as part of this OS installation.
	//
	Image *SoftwarerepositoryOperatingSystemFileRef `json:"Image,omitempty"`

	// The name of the OS install configuration.
	//
	//
	Name string `json:"Name,omitempty"`

	// Relationship to the Organization that owns the Managed Object.
	//
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// Location of the Intersight OS Deployment Utilityimage, if the user has downloaded and available locally, to be used for this OS install configuration. This image is applicable for vMedia install method.
	// Cisco publishes a OS Deployment Utility image that bootstraps and installs the user provided operating system images along with answers for unattended instllation.
	// If this property is empty for vMedia install type, the image hosted in Intersight image repository will be used. Note that in this case, the image will be downloaded from Intersight in every target server every time.
	//
	OsduImage *FirmwareServerConfigurationUtilityDistributableRef `json:"OsduImage,omitempty"`

	// Post Install Scripts to be executed specified in order.
	//
	PostInstallScripts []*OsPostInstallScriptRef `json:"PostInstallScripts"`

	// This relation provides the target server in which the OS is to be
	// installed.
	//
	Server *ComputePhysicalRef `json:"Server,omitempty"`

	// This relation is populated with the reference of OS install workflow
	// started for this request. This workflow info MO shall be used for
	// tracking further status and completion.
	//
	WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OsInstall) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OsBaseInstallConfig
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OsBaseInstallConfig = aO0

	// AO1
	var dataAO1 struct {
		ConfigurationFile *OsConfigurationFileRef `json:"ConfigurationFile,omitempty"`

		Image *SoftwarerepositoryOperatingSystemFileRef `json:"Image,omitempty"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		OsduImage *FirmwareServerConfigurationUtilityDistributableRef `json:"OsduImage,omitempty"`

		PostInstallScripts []*OsPostInstallScriptRef `json:"PostInstallScripts"`

		Server *ComputePhysicalRef `json:"Server,omitempty"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ConfigurationFile = dataAO1.ConfigurationFile

	m.Image = dataAO1.Image

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	m.OsduImage = dataAO1.OsduImage

	m.PostInstallScripts = dataAO1.PostInstallScripts

	m.Server = dataAO1.Server

	m.WorkflowInfo = dataAO1.WorkflowInfo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OsInstall) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.OsBaseInstallConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ConfigurationFile *OsConfigurationFileRef `json:"ConfigurationFile,omitempty"`

		Image *SoftwarerepositoryOperatingSystemFileRef `json:"Image,omitempty"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		OsduImage *FirmwareServerConfigurationUtilityDistributableRef `json:"OsduImage,omitempty"`

		PostInstallScripts []*OsPostInstallScriptRef `json:"PostInstallScripts"`

		Server *ComputePhysicalRef `json:"Server,omitempty"`

		WorkflowInfo *WorkflowWorkflowInfoRef `json:"WorkflowInfo,omitempty"`
	}

	dataAO1.ConfigurationFile = m.ConfigurationFile

	dataAO1.Image = m.Image

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	dataAO1.OsduImage = m.OsduImage

	dataAO1.PostInstallScripts = m.PostInstallScripts

	dataAO1.Server = m.Server

	dataAO1.WorkflowInfo = m.WorkflowInfo

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this os install
func (m *OsInstall) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OsBaseInstallConfig
	if err := m.OsBaseInstallConfig.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsduImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostInstallScripts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsInstall) validateConfigurationFile(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigurationFile) { // not required
		return nil
	}

	if m.ConfigurationFile != nil {
		if err := m.ConfigurationFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ConfigurationFile")
			}
			return err
		}
	}

	return nil
}

func (m *OsInstall) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Image")
			}
			return err
		}
	}

	return nil
}

func (m *OsInstall) validateOrganization(formats strfmt.Registry) error {

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

func (m *OsInstall) validateOsduImage(formats strfmt.Registry) error {

	if swag.IsZero(m.OsduImage) { // not required
		return nil
	}

	if m.OsduImage != nil {
		if err := m.OsduImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OsduImage")
			}
			return err
		}
	}

	return nil
}

func (m *OsInstall) validatePostInstallScripts(formats strfmt.Registry) error {

	if swag.IsZero(m.PostInstallScripts) { // not required
		return nil
	}

	for i := 0; i < len(m.PostInstallScripts); i++ {
		if swag.IsZero(m.PostInstallScripts[i]) { // not required
			continue
		}

		if m.PostInstallScripts[i] != nil {
			if err := m.PostInstallScripts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PostInstallScripts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OsInstall) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Server")
			}
			return err
		}
	}

	return nil
}

func (m *OsInstall) validateWorkflowInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowInfo) { // not required
		return nil
	}

	if m.WorkflowInfo != nil {
		if err := m.WorkflowInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkflowInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OsInstall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsInstall) UnmarshalBinary(b []byte) error {
	var res OsInstall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
