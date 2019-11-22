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

// NiaapiRevisionInfo Niaapi:Revision Info
//
// The Revision info including date comment and revision number.
//
// swagger:model niaapiRevisionInfo
type NiaapiRevisionInfo struct {
	NiaapiRevisionInfoAO0P0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NiaapiRevisionInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NiaapiRevisionInfoAO0P0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NiaapiRevisionInfoAO0P0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NiaapiRevisionInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.NiaapiRevisionInfoAO0P0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this niaapi revision info
func (m *NiaapiRevisionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NiaapiRevisionInfoAO0P0
	if err := m.NiaapiRevisionInfoAO0P0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiRevisionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiRevisionInfo) UnmarshalBinary(b []byte) error {
	var res NiaapiRevisionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NiaapiRevisionInfoAO0P0 niaapi revision info a o0 p0
// swagger:model NiaapiRevisionInfoAO0P0
type NiaapiRevisionInfoAO0P0 struct {

	// The date the revision is made.
	//
	// Format: date-time
	DatePublished strfmt.DateTime `json:"DatePublished,omitempty"`

	// The concrete type of this complex type.
	//
	// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
	// ObjectType is optional.
	// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
	// are heterogeneous, i.e. the array can contain nested documents of different types.
	//
	//
	ObjectType string `json:"ObjectType,omitempty"`

	// The changes made in this revision.
	//
	RevisionComment string `json:"RevisionComment,omitempty"`

	// The Revision No. of this revision.
	//
	RevisionNo string `json:"RevisionNo,omitempty"`

	// niaapi revision info a o0 p0
	NiaapiRevisionInfoAO0P0 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *NiaapiRevisionInfoAO0P0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// The date the revision is made.
		//
		// Format: date-time
		DatePublished strfmt.DateTime `json:"DatePublished,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The changes made in this revision.
		//
		RevisionComment string `json:"RevisionComment,omitempty"`

		// The Revision No. of this revision.
		//
		RevisionNo string `json:"RevisionNo,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv NiaapiRevisionInfoAO0P0

	rcv.DatePublished = stage1.DatePublished

	rcv.ObjectType = stage1.ObjectType

	rcv.RevisionComment = stage1.RevisionComment

	rcv.RevisionNo = stage1.RevisionNo

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "DatePublished")

	delete(stage2, "ObjectType")

	delete(stage2, "RevisionComment")

	delete(stage2, "RevisionNo")

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
		m.NiaapiRevisionInfoAO0P0 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m NiaapiRevisionInfoAO0P0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// The date the revision is made.
		//
		// Format: date-time
		DatePublished strfmt.DateTime `json:"DatePublished,omitempty"`

		// The concrete type of this complex type.
		//
		// The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the
		// ObjectType is optional.
		// The type is ambiguous when a managed object contains an array of nested documents, and the documents in the array
		// are heterogeneous, i.e. the array can contain nested documents of different types.
		//
		//
		ObjectType string `json:"ObjectType,omitempty"`

		// The changes made in this revision.
		//
		RevisionComment string `json:"RevisionComment,omitempty"`

		// The Revision No. of this revision.
		//
		RevisionNo string `json:"RevisionNo,omitempty"`
	}

	stage1.DatePublished = m.DatePublished

	stage1.ObjectType = m.ObjectType

	stage1.RevisionComment = m.RevisionComment

	stage1.RevisionNo = m.RevisionNo

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.NiaapiRevisionInfoAO0P0) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.NiaapiRevisionInfoAO0P0)
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

// Validate validates this niaapi revision info a o0 p0
func (m *NiaapiRevisionInfoAO0P0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatePublished(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NiaapiRevisionInfoAO0P0) validateDatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.DatePublished) { // not required
		return nil
	}

	if err := validate.FormatOf("DatePublished", "body", "date-time", m.DatePublished.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NiaapiRevisionInfoAO0P0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NiaapiRevisionInfoAO0P0) UnmarshalBinary(b []byte) error {
	var res NiaapiRevisionInfoAO0P0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}