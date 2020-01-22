// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssetManagedDeviceStatus Asset:Managed Device Status
//
// Maintains the Managed Device Status.
//
// swagger:model assetManagedDeviceStatus
type AssetManagedDeviceStatus struct {
	MoBaseComplexType

	AssetManagedDeviceStatusAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetManagedDeviceStatus) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 AssetManagedDeviceStatusAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AssetManagedDeviceStatusAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetManagedDeviceStatus) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AssetManagedDeviceStatusAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset managed device status
func (m *AssetManagedDeviceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AssetManagedDeviceStatusAO1P1
	if err := m.AssetManagedDeviceStatusAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AssetManagedDeviceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetManagedDeviceStatus) UnmarshalBinary(b []byte) error {
	var res AssetManagedDeviceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AssetManagedDeviceStatusAO1P1 asset managed device status a o1 p1
// swagger:model AssetManagedDeviceStatusAO1P1
type AssetManagedDeviceStatusAO1P1 struct {

	// Port used for the connection to the Cloud by the Device Connector for the Managed Device.
	//
	CloudPort int64 `json:"CloudPort,omitempty"`

	// Maintains the reason for the failure of connection to the Device in case of connection failure.
	//
	ConnectionFailureReason string `json:"ConnectionFailureReason,omitempty"`

	// Maintains the status of the connection to the Device.
	//
	// Enum: [Unknown Success Failure]
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`

	// Maintains code related to error from Device Connector, if any.
	//
	ErrorCode int64 `json:"ErrorCode,omitempty"`

	// Maintains the reason for the error from Device Connector, if any.
	//
	ErrorReason string `json:"ErrorReason,omitempty"`

	// Maintains the process pid of the Device Connector for the Managed Device.
	//
	ProcessID int64 `json:"ProcessId,omitempty"`

	// Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.
	//
	ServerPort int64 `json:"ServerPort,omitempty"`

	// Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
	//
	// Enum: [New StartSent StartSentFailure StartSuccess StartFailure UpdateSentFailure UpdateSent DeleteSentFailure DeleteInProgress DeleteFailure DeleteSuccess]
	State *string `json:"State,omitempty"`

	// asset managed device status a o1 p1
	AssetManagedDeviceStatusAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *AssetManagedDeviceStatusAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Port used for the connection to the Cloud by the Device Connector for the Managed Device.
		//
		CloudPort int64 `json:"CloudPort,omitempty"`

		// Maintains the reason for the failure of connection to the Device in case of connection failure.
		//
		ConnectionFailureReason string `json:"ConnectionFailureReason,omitempty"`

		// Maintains the status of the connection to the Device.
		//
		// Enum: [Unknown Success Failure]
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`

		// Maintains code related to error from Device Connector, if any.
		//
		ErrorCode int64 `json:"ErrorCode,omitempty"`

		// Maintains the reason for the error from Device Connector, if any.
		//
		ErrorReason string `json:"ErrorReason,omitempty"`

		// Maintains the process pid of the Device Connector for the Managed Device.
		//
		ProcessID int64 `json:"ProcessId,omitempty"`

		// Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.
		//
		ServerPort int64 `json:"ServerPort,omitempty"`

		// Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
		//
		// Enum: [New StartSent StartSentFailure StartSuccess StartFailure UpdateSentFailure UpdateSent DeleteSentFailure DeleteInProgress DeleteFailure DeleteSuccess]
		State *string `json:"State,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AssetManagedDeviceStatusAO1P1

	rcv.CloudPort = stage1.CloudPort

	rcv.ConnectionFailureReason = stage1.ConnectionFailureReason

	rcv.ConnectionStatus = stage1.ConnectionStatus

	rcv.ErrorCode = stage1.ErrorCode

	rcv.ErrorReason = stage1.ErrorReason

	rcv.ProcessID = stage1.ProcessID

	rcv.ServerPort = stage1.ServerPort

	rcv.State = stage1.State

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "CloudPort")

	delete(stage2, "ConnectionFailureReason")

	delete(stage2, "ConnectionStatus")

	delete(stage2, "ErrorCode")

	delete(stage2, "ErrorReason")

	delete(stage2, "ProcessId")

	delete(stage2, "ServerPort")

	delete(stage2, "State")

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
		m.AssetManagedDeviceStatusAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m AssetManagedDeviceStatusAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Port used for the connection to the Cloud by the Device Connector for the Managed Device.
		//
		CloudPort int64 `json:"CloudPort,omitempty"`

		// Maintains the reason for the failure of connection to the Device in case of connection failure.
		//
		ConnectionFailureReason string `json:"ConnectionFailureReason,omitempty"`

		// Maintains the status of the connection to the Device.
		//
		// Enum: [Unknown Success Failure]
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`

		// Maintains code related to error from Device Connector, if any.
		//
		ErrorCode int64 `json:"ErrorCode,omitempty"`

		// Maintains the reason for the error from Device Connector, if any.
		//
		ErrorReason string `json:"ErrorReason,omitempty"`

		// Maintains the process pid of the Device Connector for the Managed Device.
		//
		ProcessID int64 `json:"ProcessId,omitempty"`

		// Port used for receiving requests from Intersight Assist by the Device Connector for the Managed Device.
		//
		ServerPort int64 `json:"ServerPort,omitempty"`

		// Maintains the state of the Managed Device, such as Start Success, Start Failure, etc. See ManagedDeviceState for device connection states.
		//
		// Enum: [New StartSent StartSentFailure StartSuccess StartFailure UpdateSentFailure UpdateSent DeleteSentFailure DeleteInProgress DeleteFailure DeleteSuccess]
		State *string `json:"State,omitempty"`
	}

	stage1.CloudPort = m.CloudPort

	stage1.ConnectionFailureReason = m.ConnectionFailureReason

	stage1.ConnectionStatus = m.ConnectionStatus

	stage1.ErrorCode = m.ErrorCode

	stage1.ErrorReason = m.ErrorReason

	stage1.ProcessID = m.ProcessID

	stage1.ServerPort = m.ServerPort

	stage1.State = m.State

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.AssetManagedDeviceStatusAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.AssetManagedDeviceStatusAO1P1)
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

// Validate validates this asset managed device status a o1 p1
func (m *AssetManagedDeviceStatusAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assetManagedDeviceStatusAO1P1TypeConnectionStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Success","Failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetManagedDeviceStatusAO1P1TypeConnectionStatusPropEnum = append(assetManagedDeviceStatusAO1P1TypeConnectionStatusPropEnum, v)
	}
}

const (

	// AssetManagedDeviceStatusAO1P1ConnectionStatusUnknown captures enum value "Unknown"
	AssetManagedDeviceStatusAO1P1ConnectionStatusUnknown string = "Unknown"

	// AssetManagedDeviceStatusAO1P1ConnectionStatusSuccess captures enum value "Success"
	AssetManagedDeviceStatusAO1P1ConnectionStatusSuccess string = "Success"

	// AssetManagedDeviceStatusAO1P1ConnectionStatusFailure captures enum value "Failure"
	AssetManagedDeviceStatusAO1P1ConnectionStatusFailure string = "Failure"
)

// prop value enum
func (m *AssetManagedDeviceStatusAO1P1) validateConnectionStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetManagedDeviceStatusAO1P1TypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetManagedDeviceStatusAO1P1) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("ConnectionStatus", "body", *m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

var assetManagedDeviceStatusAO1P1TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["New","StartSent","StartSentFailure","StartSuccess","StartFailure","UpdateSentFailure","UpdateSent","DeleteSentFailure","DeleteInProgress","DeleteFailure","DeleteSuccess"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetManagedDeviceStatusAO1P1TypeStatePropEnum = append(assetManagedDeviceStatusAO1P1TypeStatePropEnum, v)
	}
}

const (

	// AssetManagedDeviceStatusAO1P1StateNew captures enum value "New"
	AssetManagedDeviceStatusAO1P1StateNew string = "New"

	// AssetManagedDeviceStatusAO1P1StateStartSent captures enum value "StartSent"
	AssetManagedDeviceStatusAO1P1StateStartSent string = "StartSent"

	// AssetManagedDeviceStatusAO1P1StateStartSentFailure captures enum value "StartSentFailure"
	AssetManagedDeviceStatusAO1P1StateStartSentFailure string = "StartSentFailure"

	// AssetManagedDeviceStatusAO1P1StateStartSuccess captures enum value "StartSuccess"
	AssetManagedDeviceStatusAO1P1StateStartSuccess string = "StartSuccess"

	// AssetManagedDeviceStatusAO1P1StateStartFailure captures enum value "StartFailure"
	AssetManagedDeviceStatusAO1P1StateStartFailure string = "StartFailure"

	// AssetManagedDeviceStatusAO1P1StateUpdateSentFailure captures enum value "UpdateSentFailure"
	AssetManagedDeviceStatusAO1P1StateUpdateSentFailure string = "UpdateSentFailure"

	// AssetManagedDeviceStatusAO1P1StateUpdateSent captures enum value "UpdateSent"
	AssetManagedDeviceStatusAO1P1StateUpdateSent string = "UpdateSent"

	// AssetManagedDeviceStatusAO1P1StateDeleteSentFailure captures enum value "DeleteSentFailure"
	AssetManagedDeviceStatusAO1P1StateDeleteSentFailure string = "DeleteSentFailure"

	// AssetManagedDeviceStatusAO1P1StateDeleteInProgress captures enum value "DeleteInProgress"
	AssetManagedDeviceStatusAO1P1StateDeleteInProgress string = "DeleteInProgress"

	// AssetManagedDeviceStatusAO1P1StateDeleteFailure captures enum value "DeleteFailure"
	AssetManagedDeviceStatusAO1P1StateDeleteFailure string = "DeleteFailure"

	// AssetManagedDeviceStatusAO1P1StateDeleteSuccess captures enum value "DeleteSuccess"
	AssetManagedDeviceStatusAO1P1StateDeleteSuccess string = "DeleteSuccess"
)

// prop value enum
func (m *AssetManagedDeviceStatusAO1P1) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetManagedDeviceStatusAO1P1TypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetManagedDeviceStatusAO1P1) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetManagedDeviceStatusAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetManagedDeviceStatusAO1P1) UnmarshalBinary(b []byte) error {
	var res AssetManagedDeviceStatusAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
