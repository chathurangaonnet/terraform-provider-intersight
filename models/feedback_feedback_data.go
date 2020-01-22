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

// FeedbackFeedbackData Feedback:Feedback Data
//
// Feedback data that collected on the UI from user.
//
// swagger:model feedbackFeedbackData
type FeedbackFeedbackData struct {
	MoBaseComplexType

	FeedbackFeedbackDataAO1P1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FeedbackFeedbackData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseComplexType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseComplexType = aO0

	// AO1
	var aO1 FeedbackFeedbackDataAO1P1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.FeedbackFeedbackDataAO1P1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FeedbackFeedbackData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseComplexType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.FeedbackFeedbackDataAO1P1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this feedback feedback data
func (m *FeedbackFeedbackData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseComplexType
	if err := m.MoBaseComplexType.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with FeedbackFeedbackDataAO1P1
	if err := m.FeedbackFeedbackDataAO1P1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *FeedbackFeedbackData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackData) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FeedbackFeedbackDataAO1P1 feedback feedback data a o1 p1
// swagger:model FeedbackFeedbackDataAO1P1
type FeedbackFeedbackDataAO1P1 struct {

	// Account name of the feedback sender. Copied in order to be persisted in case of account removal.
	//
	AccountName string `json:"AccountName,omitempty"`

	// Text of the feedback as provided by the user, if it is a bug or a comment.
	//
	Comment string `json:"Comment,omitempty"`

	// User's email address details.
	//
	Email string `json:"Email,omitempty"`

	// Evalation rating as provided by the user to capture user sentiment regarding the issue.
	//
	// Enum: [Excellent Poor Fair Good]
	Evaluation *string `json:"Evaluation,omitempty"`

	// If a user is open for follow-up or not.
	//
	FollowUp *bool `json:"FollowUp,omitempty"`

	// Bunch of last traceId for reproducing user last activity.
	//
	TraceIds interface{} `json:"TraceIds,omitempty"`

	// Type of the feedback from user.
	//
	// Enum: [Evaluation Bug]
	Type *string `json:"Type,omitempty"`

	// feedback feedback data a o1 p1
	FeedbackFeedbackDataAO1P1 map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *FeedbackFeedbackDataAO1P1) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Account name of the feedback sender. Copied in order to be persisted in case of account removal.
		//
		AccountName string `json:"AccountName,omitempty"`

		// Text of the feedback as provided by the user, if it is a bug or a comment.
		//
		Comment string `json:"Comment,omitempty"`

		// User's email address details.
		//
		Email string `json:"Email,omitempty"`

		// Evalation rating as provided by the user to capture user sentiment regarding the issue.
		//
		// Enum: [Excellent Poor Fair Good]
		Evaluation *string `json:"Evaluation,omitempty"`

		// If a user is open for follow-up or not.
		//
		FollowUp *bool `json:"FollowUp,omitempty"`

		// Bunch of last traceId for reproducing user last activity.
		//
		TraceIds interface{} `json:"TraceIds,omitempty"`

		// Type of the feedback from user.
		//
		// Enum: [Evaluation Bug]
		Type *string `json:"Type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv FeedbackFeedbackDataAO1P1

	rcv.AccountName = stage1.AccountName

	rcv.Comment = stage1.Comment

	rcv.Email = stage1.Email

	rcv.Evaluation = stage1.Evaluation

	rcv.FollowUp = stage1.FollowUp

	rcv.TraceIds = stage1.TraceIds

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "AccountName")

	delete(stage2, "Comment")

	delete(stage2, "Email")

	delete(stage2, "Evaluation")

	delete(stage2, "FollowUp")

	delete(stage2, "TraceIds")

	delete(stage2, "Type")

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
		m.FeedbackFeedbackDataAO1P1 = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m FeedbackFeedbackDataAO1P1) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Account name of the feedback sender. Copied in order to be persisted in case of account removal.
		//
		AccountName string `json:"AccountName,omitempty"`

		// Text of the feedback as provided by the user, if it is a bug or a comment.
		//
		Comment string `json:"Comment,omitempty"`

		// User's email address details.
		//
		Email string `json:"Email,omitempty"`

		// Evalation rating as provided by the user to capture user sentiment regarding the issue.
		//
		// Enum: [Excellent Poor Fair Good]
		Evaluation *string `json:"Evaluation,omitempty"`

		// If a user is open for follow-up or not.
		//
		FollowUp *bool `json:"FollowUp,omitempty"`

		// Bunch of last traceId for reproducing user last activity.
		//
		TraceIds interface{} `json:"TraceIds,omitempty"`

		// Type of the feedback from user.
		//
		// Enum: [Evaluation Bug]
		Type *string `json:"Type,omitempty"`
	}

	stage1.AccountName = m.AccountName

	stage1.Comment = m.Comment

	stage1.Email = m.Email

	stage1.Evaluation = m.Evaluation

	stage1.FollowUp = m.FollowUp

	stage1.TraceIds = m.TraceIds

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.FeedbackFeedbackDataAO1P1) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.FeedbackFeedbackDataAO1P1)
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

// Validate validates this feedback feedback data a o1 p1
func (m *FeedbackFeedbackDataAO1P1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var feedbackFeedbackDataAO1P1TypeEvaluationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Excellent","Poor","Fair","Good"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataAO1P1TypeEvaluationPropEnum = append(feedbackFeedbackDataAO1P1TypeEvaluationPropEnum, v)
	}
}

const (

	// FeedbackFeedbackDataAO1P1EvaluationExcellent captures enum value "Excellent"
	FeedbackFeedbackDataAO1P1EvaluationExcellent string = "Excellent"

	// FeedbackFeedbackDataAO1P1EvaluationPoor captures enum value "Poor"
	FeedbackFeedbackDataAO1P1EvaluationPoor string = "Poor"

	// FeedbackFeedbackDataAO1P1EvaluationFair captures enum value "Fair"
	FeedbackFeedbackDataAO1P1EvaluationFair string = "Fair"

	// FeedbackFeedbackDataAO1P1EvaluationGood captures enum value "Good"
	FeedbackFeedbackDataAO1P1EvaluationGood string = "Good"
)

// prop value enum
func (m *FeedbackFeedbackDataAO1P1) validateEvaluationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataAO1P1TypeEvaluationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackDataAO1P1) validateEvaluation(formats strfmt.Registry) error {

	if swag.IsZero(m.Evaluation) { // not required
		return nil
	}

	// value enum
	if err := m.validateEvaluationEnum("Evaluation", "body", *m.Evaluation); err != nil {
		return err
	}

	return nil
}

var feedbackFeedbackDataAO1P1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Evaluation","Bug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedbackFeedbackDataAO1P1TypeTypePropEnum = append(feedbackFeedbackDataAO1P1TypeTypePropEnum, v)
	}
}

const (

	// FeedbackFeedbackDataAO1P1TypeEvaluation captures enum value "Evaluation"
	FeedbackFeedbackDataAO1P1TypeEvaluation string = "Evaluation"

	// FeedbackFeedbackDataAO1P1TypeBug captures enum value "Bug"
	FeedbackFeedbackDataAO1P1TypeBug string = "Bug"
)

// prop value enum
func (m *FeedbackFeedbackDataAO1P1) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, feedbackFeedbackDataAO1P1TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FeedbackFeedbackDataAO1P1) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedbackFeedbackDataAO1P1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedbackFeedbackDataAO1P1) UnmarshalBinary(b []byte) error {
	var res FeedbackFeedbackDataAO1P1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
