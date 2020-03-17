// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComputeServerSetting Server Settings
//
// Models the configurable properties of a server in Intersight.
//
// swagger:model computeServerSetting
type ComputeServerSetting struct {
	InventoryBase

	// User configured state of the locator LED for the server.
	// Enum: [None On Off]
	AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`

	// User configured power state of the server.
	// Enum: [Policy PowerOn PowerOff PowerCycle HardReset Shutdown Reboot]
	AdminPowerState *string `json:"AdminPowerState,omitempty"`

	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server.
	// Read Only: true
	// Enum: [Applied Applying Failed]
	ConfigState string `json:"ConfigState,omitempty"`

	// Relates to locator LED MO of this server.
	// Read Only: true
	LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

	// The name of the device chosen by user for configuring One-Time Boot device.
	OneTimeBootDevice string `json:"OneTimeBootDevice,omitempty"`

	// The Persistent Memory Modules operation properties.
	PersistentMemoryOperation *ComputePersistentMemoryOperation `json:"PersistentMemoryOperation,omitempty"`

	// Relates to the device end point from which this server was discovered.
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// Compute RackUnit for which the settings are related to.
	// Read Only: true
	Server *ComputeRackUnitRef `json:"Server,omitempty"`

	// The common server configurable properties between a server and a server profile.
	ServerConfig *ComputeServerConfig `json:"ServerConfig,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ComputeServerSetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 InventoryBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.InventoryBase = aO0

	// AO1
	var dataAO1 struct {
		AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`

		AdminPowerState *string `json:"AdminPowerState,omitempty"`

		ConfigState string `json:"ConfigState,omitempty"`

		LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

		OneTimeBootDevice string `json:"OneTimeBootDevice,omitempty"`

		PersistentMemoryOperation *ComputePersistentMemoryOperation `json:"PersistentMemoryOperation,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Server *ComputeRackUnitRef `json:"Server,omitempty"`

		ServerConfig *ComputeServerConfig `json:"ServerConfig,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AdminLocatorLedState = dataAO1.AdminLocatorLedState

	m.AdminPowerState = dataAO1.AdminPowerState

	m.ConfigState = dataAO1.ConfigState

	m.LocatorLed = dataAO1.LocatorLed

	m.OneTimeBootDevice = dataAO1.OneTimeBootDevice

	m.PersistentMemoryOperation = dataAO1.PersistentMemoryOperation

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.Server = dataAO1.Server

	m.ServerConfig = dataAO1.ServerConfig

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ComputeServerSetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.InventoryBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`

		AdminPowerState *string `json:"AdminPowerState,omitempty"`

		ConfigState string `json:"ConfigState,omitempty"`

		LocatorLed *EquipmentLocatorLedRef `json:"LocatorLed,omitempty"`

		OneTimeBootDevice string `json:"OneTimeBootDevice,omitempty"`

		PersistentMemoryOperation *ComputePersistentMemoryOperation `json:"PersistentMemoryOperation,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		Server *ComputeRackUnitRef `json:"Server,omitempty"`

		ServerConfig *ComputeServerConfig `json:"ServerConfig,omitempty"`
	}

	dataAO1.AdminLocatorLedState = m.AdminLocatorLedState

	dataAO1.AdminPowerState = m.AdminPowerState

	dataAO1.ConfigState = m.ConfigState

	dataAO1.LocatorLed = m.LocatorLed

	dataAO1.OneTimeBootDevice = m.OneTimeBootDevice

	dataAO1.PersistentMemoryOperation = m.PersistentMemoryOperation

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.Server = m.Server

	dataAO1.ServerConfig = m.ServerConfig

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this compute server setting
func (m *ComputeServerSetting) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with InventoryBase
	if err := m.InventoryBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdminLocatorLedState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdminPowerState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocatorLed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistentMemoryOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computeServerSettingTypeAdminLocatorLedStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","On","Off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computeServerSettingTypeAdminLocatorLedStatePropEnum = append(computeServerSettingTypeAdminLocatorLedStatePropEnum, v)
	}
}

// property enum
func (m *ComputeServerSetting) validateAdminLocatorLedStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computeServerSettingTypeAdminLocatorLedStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputeServerSetting) validateAdminLocatorLedState(formats strfmt.Registry) error {

	if swag.IsZero(m.AdminLocatorLedState) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdminLocatorLedStateEnum("AdminLocatorLedState", "body", *m.AdminLocatorLedState); err != nil {
		return err
	}

	return nil
}

var computeServerSettingTypeAdminPowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Policy","PowerOn","PowerOff","PowerCycle","HardReset","Shutdown","Reboot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computeServerSettingTypeAdminPowerStatePropEnum = append(computeServerSettingTypeAdminPowerStatePropEnum, v)
	}
}

// property enum
func (m *ComputeServerSetting) validateAdminPowerStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computeServerSettingTypeAdminPowerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputeServerSetting) validateAdminPowerState(formats strfmt.Registry) error {

	if swag.IsZero(m.AdminPowerState) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdminPowerStateEnum("AdminPowerState", "body", *m.AdminPowerState); err != nil {
		return err
	}

	return nil
}

var computeServerSettingTypeConfigStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Applied","Applying","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computeServerSettingTypeConfigStatePropEnum = append(computeServerSettingTypeConfigStatePropEnum, v)
	}
}

// property enum
func (m *ComputeServerSetting) validateConfigStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, computeServerSettingTypeConfigStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ComputeServerSetting) validateConfigState(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigState) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigStateEnum("ConfigState", "body", m.ConfigState); err != nil {
		return err
	}

	return nil
}

func (m *ComputeServerSetting) validateLocatorLed(formats strfmt.Registry) error {

	if swag.IsZero(m.LocatorLed) { // not required
		return nil
	}

	if m.LocatorLed != nil {
		if err := m.LocatorLed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LocatorLed")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeServerSetting) validatePersistentMemoryOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.PersistentMemoryOperation) { // not required
		return nil
	}

	if m.PersistentMemoryOperation != nil {
		if err := m.PersistentMemoryOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PersistentMemoryOperation")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeServerSetting) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

func (m *ComputeServerSetting) validateServer(formats strfmt.Registry) error {

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

func (m *ComputeServerSetting) validateServerConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ServerConfig) { // not required
		return nil
	}

	if m.ServerConfig != nil {
		if err := m.ServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ServerConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeServerSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeServerSetting) UnmarshalBinary(b []byte) error {
	var res ComputeServerSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
