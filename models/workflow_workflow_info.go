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

// WorkflowWorkflowInfo Workflow:Workflow Info
//
// Contains information for a workflow execution which is a runtime instance of workflow.
//
// swagger:model workflowWorkflowInfo
type WorkflowWorkflowInfo struct {
	MoBaseMo

	// The Account to which the workflow is associated.
	Account *IamAccountRef `json:"Account,omitempty"`

	// The action of the workflow such as start, cancel, retry, pause.
	// Enum: [None Create Start Pause Resume Retry RetryFailed Cancel]
	Action *string `json:"Action,omitempty"`

	// The time when the workflow info will be removed from database.
	// Read Only: true
	// Format: date-time
	CleanupTime strfmt.DateTime `json:"CleanupTime,omitempty"`

	// The time when the workflow reached a final state.
	// Read Only: true
	// Format: date-time
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	// The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database.
	FailedWorkflowCleanupDuration int64 `json:"FailedWorkflowCleanupDuration,omitempty"`

	// All the given inputs for the workflow.
	Input interface{} `json:"Input,omitempty"`

	// A workflow instance Id which is the unique identified for the workflow execution.
	// Read Only: true
	InstID string `json:"InstId,omitempty"`

	// Denotes if this workflow is internal and should be hidden from user view of running workflows.
	Internal *bool `json:"Internal,omitempty"`

	// The last action that was issued on the workflow is saved in this field.
	// Read Only: true
	// Enum: [None Create Start Pause Resume Retry RetryFailed Cancel]
	LastAction string `json:"LastAction,omitempty"`

	// Collection of Workflow execution messages with severity.
	Message []*WorkflowMessage `json:"Message"`

	// Version of the workflow metadata for which this workflow execution was started.
	MetaVersion int64 `json:"MetaVersion,omitempty"`

	// A name of the workflow execution instance.
	Name string `json:"Name,omitempty"`

	// The Organization to which the workflow is associated.
	Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

	// All the generated outputs for the workflow.
	// Read Only: true
	Output interface{} `json:"Output,omitempty"`

	// Link to the task in the parent workflow which launched this workflow.
	// Read Only: true
	ParentTaskInfo *WorkflowTaskInfoRef `json:"ParentTaskInfo,omitempty"`

	// Reference to the PendingDynamicWorkflowInfo that was used to construct this workflow instance.
	// Read Only: true
	PendingDynamicWorkflowInfo *WorkflowPendingDynamicWorkflowInfoRef `json:"PendingDynamicWorkflowInfo,omitempty"`

	// Reference to the permission object for which user has access to start this workflow.
	Permission *IamPermissionRef `json:"Permission,omitempty"`

	// This field indicates percentage of workflow task execution.
	// Read Only: true
	Progress float32 `json:"Progress,omitempty"`

	// Type to capture all the properties for the workflow info passed on from workflow definition.
	// Read Only: true
	Properties *WorkflowWorkflowInfoProperties `json:"Properties,omitempty"`

	// This field is applicable when Retry action is issued for a workflow which is in a final state. When this field is not specified then the workflow will retry from the start of the workflow. When this field is specified then the workflow will be retried from the specified task. The field should carry the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.
	RetryFromTaskName string `json:"RetryFromTaskName,omitempty"`

	// The source microservice name which is the owner for this workflow.
	// Read Only: true
	Src string `json:"Src,omitempty"`

	// The time when the workflow was started for execution.
	// Read Only: true
	// Format: date-time
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	// A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED).
	// Read Only: true
	Status string `json:"Status,omitempty"`

	// The duration in hours after which the workflow info for successful workflow will be removed from database.
	SuccessWorkflowCleanupDuration int64 `json:"SuccessWorkflowCleanupDuration,omitempty"`

	// List of task instances that ran as part of this workflow execution.
	// Read Only: true
	TaskInfos []*WorkflowTaskInfoRef `json:"TaskInfos"`

	// The trace id to keep track of workflow execution.
	// Read Only: true
	TraceID string `json:"TraceId,omitempty"`

	// A type of the workflow (serverconfig, ansible_monitoring).
	// Read Only: true
	Type string `json:"Type,omitempty"`

	// The user identifier which indicates the user that started this workflow.
	// Read Only: true
	UserID string `json:"UserId,omitempty"`

	// Denotes the reason workflow is in waiting status.
	// Enum: [None GatherTasks Duplicate RateLimit WaitTask PendingRetryFailed]
	WaitReason *string `json:"WaitReason,omitempty"`

	// The workflow context which contains initiator and target information.
	WorkflowCtx interface{} `json:"WorkflowCtx,omitempty"`

	// The workflow definition that was used to launch this workflow execution instance.
	WorkflowDefinition *WorkflowWorkflowDefinitionRef `json:"WorkflowDefinition,omitempty"`

	// The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.
	// Enum: [SystemDefined UserDefined Dynamic]
	WorkflowMetaType *string `json:"WorkflowMetaType,omitempty"`

	// Total number of workflow tasks in this workflow.
	// Read Only: true
	WorkflowTaskCount int64 `json:"WorkflowTaskCount,omitempty"`

	// A collection of references to the [server.Profile](mo://server.Profile) Managed Object.
	// When this managed object is deleted, the referenced [server.Profile](mo://server.Profile) MO unsets its reference to this deleted MO.
	// Read Only: true
	Nr0Profile *ServerProfileRef `json:"_0_Profile,omitempty"`

	// A collection of references to the [hyperflex.ClusterProfile](mo://hyperflex.ClusterProfile) Managed Object.
	// When this managed object is deleted, the referenced [hyperflex.ClusterProfile](mo://hyperflex.ClusterProfile) MO unsets its reference to this deleted MO.
	// Read Only: true
	Nr1ClusterProfile *HyperflexClusterProfileRef `json:"_1_ClusterProfile,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowWorkflowInfo) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Action *string `json:"Action,omitempty"`

		CleanupTime strfmt.DateTime `json:"CleanupTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		FailedWorkflowCleanupDuration int64 `json:"FailedWorkflowCleanupDuration,omitempty"`

		Input interface{} `json:"Input,omitempty"`

		InstID string `json:"InstId,omitempty"`

		Internal *bool `json:"Internal,omitempty"`

		LastAction string `json:"LastAction,omitempty"`

		Message []*WorkflowMessage `json:"Message"`

		MetaVersion int64 `json:"MetaVersion,omitempty"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		ParentTaskInfo *WorkflowTaskInfoRef `json:"ParentTaskInfo,omitempty"`

		PendingDynamicWorkflowInfo *WorkflowPendingDynamicWorkflowInfoRef `json:"PendingDynamicWorkflowInfo,omitempty"`

		Permission *IamPermissionRef `json:"Permission,omitempty"`

		Progress float32 `json:"Progress,omitempty"`

		Properties *WorkflowWorkflowInfoProperties `json:"Properties,omitempty"`

		RetryFromTaskName string `json:"RetryFromTaskName,omitempty"`

		Src string `json:"Src,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		SuccessWorkflowCleanupDuration int64 `json:"SuccessWorkflowCleanupDuration,omitempty"`

		TaskInfos []*WorkflowTaskInfoRef `json:"TaskInfos"`

		TraceID string `json:"TraceId,omitempty"`

		Type string `json:"Type,omitempty"`

		UserID string `json:"UserId,omitempty"`

		WaitReason *string `json:"WaitReason,omitempty"`

		WorkflowCtx interface{} `json:"WorkflowCtx,omitempty"`

		WorkflowDefinition *WorkflowWorkflowDefinitionRef `json:"WorkflowDefinition,omitempty"`

		WorkflowMetaType *string `json:"WorkflowMetaType,omitempty"`

		WorkflowTaskCount int64 `json:"WorkflowTaskCount,omitempty"`

		Nr0Profile *ServerProfileRef `json:"_0_Profile,omitempty"`

		Nr1ClusterProfile *HyperflexClusterProfileRef `json:"_1_ClusterProfile,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Account = dataAO1.Account

	m.Action = dataAO1.Action

	m.CleanupTime = dataAO1.CleanupTime

	m.EndTime = dataAO1.EndTime

	m.FailedWorkflowCleanupDuration = dataAO1.FailedWorkflowCleanupDuration

	m.Input = dataAO1.Input

	m.InstID = dataAO1.InstID

	m.Internal = dataAO1.Internal

	m.LastAction = dataAO1.LastAction

	m.Message = dataAO1.Message

	m.MetaVersion = dataAO1.MetaVersion

	m.Name = dataAO1.Name

	m.Organization = dataAO1.Organization

	m.Output = dataAO1.Output

	m.ParentTaskInfo = dataAO1.ParentTaskInfo

	m.PendingDynamicWorkflowInfo = dataAO1.PendingDynamicWorkflowInfo

	m.Permission = dataAO1.Permission

	m.Progress = dataAO1.Progress

	m.Properties = dataAO1.Properties

	m.RetryFromTaskName = dataAO1.RetryFromTaskName

	m.Src = dataAO1.Src

	m.StartTime = dataAO1.StartTime

	m.Status = dataAO1.Status

	m.SuccessWorkflowCleanupDuration = dataAO1.SuccessWorkflowCleanupDuration

	m.TaskInfos = dataAO1.TaskInfos

	m.TraceID = dataAO1.TraceID

	m.Type = dataAO1.Type

	m.UserID = dataAO1.UserID

	m.WaitReason = dataAO1.WaitReason

	m.WorkflowCtx = dataAO1.WorkflowCtx

	m.WorkflowDefinition = dataAO1.WorkflowDefinition

	m.WorkflowMetaType = dataAO1.WorkflowMetaType

	m.WorkflowTaskCount = dataAO1.WorkflowTaskCount

	m.Nr0Profile = dataAO1.Nr0Profile

	m.Nr1ClusterProfile = dataAO1.Nr1ClusterProfile

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowWorkflowInfo) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Account *IamAccountRef `json:"Account,omitempty"`

		Action *string `json:"Action,omitempty"`

		CleanupTime strfmt.DateTime `json:"CleanupTime,omitempty"`

		EndTime strfmt.DateTime `json:"EndTime,omitempty"`

		FailedWorkflowCleanupDuration int64 `json:"FailedWorkflowCleanupDuration,omitempty"`

		Input interface{} `json:"Input,omitempty"`

		InstID string `json:"InstId,omitempty"`

		Internal *bool `json:"Internal,omitempty"`

		LastAction string `json:"LastAction,omitempty"`

		Message []*WorkflowMessage `json:"Message"`

		MetaVersion int64 `json:"MetaVersion,omitempty"`

		Name string `json:"Name,omitempty"`

		Organization *OrganizationOrganizationRef `json:"Organization,omitempty"`

		Output interface{} `json:"Output,omitempty"`

		ParentTaskInfo *WorkflowTaskInfoRef `json:"ParentTaskInfo,omitempty"`

		PendingDynamicWorkflowInfo *WorkflowPendingDynamicWorkflowInfoRef `json:"PendingDynamicWorkflowInfo,omitempty"`

		Permission *IamPermissionRef `json:"Permission,omitempty"`

		Progress float32 `json:"Progress,omitempty"`

		Properties *WorkflowWorkflowInfoProperties `json:"Properties,omitempty"`

		RetryFromTaskName string `json:"RetryFromTaskName,omitempty"`

		Src string `json:"Src,omitempty"`

		StartTime strfmt.DateTime `json:"StartTime,omitempty"`

		Status string `json:"Status,omitempty"`

		SuccessWorkflowCleanupDuration int64 `json:"SuccessWorkflowCleanupDuration,omitempty"`

		TaskInfos []*WorkflowTaskInfoRef `json:"TaskInfos"`

		TraceID string `json:"TraceId,omitempty"`

		Type string `json:"Type,omitempty"`

		UserID string `json:"UserId,omitempty"`

		WaitReason *string `json:"WaitReason,omitempty"`

		WorkflowCtx interface{} `json:"WorkflowCtx,omitempty"`

		WorkflowDefinition *WorkflowWorkflowDefinitionRef `json:"WorkflowDefinition,omitempty"`

		WorkflowMetaType *string `json:"WorkflowMetaType,omitempty"`

		WorkflowTaskCount int64 `json:"WorkflowTaskCount,omitempty"`

		Nr0Profile *ServerProfileRef `json:"_0_Profile,omitempty"`

		Nr1ClusterProfile *HyperflexClusterProfileRef `json:"_1_ClusterProfile,omitempty"`
	}

	dataAO1.Account = m.Account

	dataAO1.Action = m.Action

	dataAO1.CleanupTime = m.CleanupTime

	dataAO1.EndTime = m.EndTime

	dataAO1.FailedWorkflowCleanupDuration = m.FailedWorkflowCleanupDuration

	dataAO1.Input = m.Input

	dataAO1.InstID = m.InstID

	dataAO1.Internal = m.Internal

	dataAO1.LastAction = m.LastAction

	dataAO1.Message = m.Message

	dataAO1.MetaVersion = m.MetaVersion

	dataAO1.Name = m.Name

	dataAO1.Organization = m.Organization

	dataAO1.Output = m.Output

	dataAO1.ParentTaskInfo = m.ParentTaskInfo

	dataAO1.PendingDynamicWorkflowInfo = m.PendingDynamicWorkflowInfo

	dataAO1.Permission = m.Permission

	dataAO1.Progress = m.Progress

	dataAO1.Properties = m.Properties

	dataAO1.RetryFromTaskName = m.RetryFromTaskName

	dataAO1.Src = m.Src

	dataAO1.StartTime = m.StartTime

	dataAO1.Status = m.Status

	dataAO1.SuccessWorkflowCleanupDuration = m.SuccessWorkflowCleanupDuration

	dataAO1.TaskInfos = m.TaskInfos

	dataAO1.TraceID = m.TraceID

	dataAO1.Type = m.Type

	dataAO1.UserID = m.UserID

	dataAO1.WaitReason = m.WaitReason

	dataAO1.WorkflowCtx = m.WorkflowCtx

	dataAO1.WorkflowDefinition = m.WorkflowDefinition

	dataAO1.WorkflowMetaType = m.WorkflowMetaType

	dataAO1.WorkflowTaskCount = m.WorkflowTaskCount

	dataAO1.Nr0Profile = m.Nr0Profile

	dataAO1.Nr1ClusterProfile = m.Nr1ClusterProfile

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow workflow info
func (m *WorkflowWorkflowInfo) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCleanupTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentTaskInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePendingDynamicWorkflowInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskInfos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaitReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowMetaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNr0Profile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNr1ClusterProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowWorkflowInfo) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Account")
			}
			return err
		}
	}

	return nil
}

var workflowWorkflowInfoTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Create","Start","Pause","Resume","Retry","RetryFailed","Cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowInfoTypeActionPropEnum = append(workflowWorkflowInfoTypeActionPropEnum, v)
	}
}

// property enum
func (m *WorkflowWorkflowInfo) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowInfoTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowInfo) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("Action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateCleanupTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CleanupTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CleanupTime", "body", "date-time", m.CleanupTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var workflowWorkflowInfoTypeLastActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Create","Start","Pause","Resume","Retry","RetryFailed","Cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowInfoTypeLastActionPropEnum = append(workflowWorkflowInfoTypeLastActionPropEnum, v)
	}
}

// property enum
func (m *WorkflowWorkflowInfo) validateLastActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowInfoTypeLastActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowInfo) validateLastAction(formats strfmt.Registry) error {

	if swag.IsZero(m.LastAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastActionEnum("LastAction", "body", m.LastAction); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	for i := 0; i < len(m.Message); i++ {
		if swag.IsZero(m.Message[i]) { // not required
			continue
		}

		if m.Message[i] != nil {
			if err := m.Message[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Message" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateOrganization(formats strfmt.Registry) error {

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

func (m *WorkflowWorkflowInfo) validateParentTaskInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentTaskInfo) { // not required
		return nil
	}

	if m.ParentTaskInfo != nil {
		if err := m.ParentTaskInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ParentTaskInfo")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validatePendingDynamicWorkflowInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingDynamicWorkflowInfo) { // not required
		return nil
	}

	if m.PendingDynamicWorkflowInfo != nil {
		if err := m.PendingDynamicWorkflowInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PendingDynamicWorkflowInfo")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if m.Permission != nil {
		if err := m.Permission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Permission")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Properties")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateTaskInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.TaskInfos); i++ {
		if swag.IsZero(m.TaskInfos[i]) { // not required
			continue
		}

		if m.TaskInfos[i] != nil {
			if err := m.TaskInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TaskInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var workflowWorkflowInfoTypeWaitReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","GatherTasks","Duplicate","RateLimit","WaitTask","PendingRetryFailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowInfoTypeWaitReasonPropEnum = append(workflowWorkflowInfoTypeWaitReasonPropEnum, v)
	}
}

// property enum
func (m *WorkflowWorkflowInfo) validateWaitReasonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowInfoTypeWaitReasonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowInfo) validateWaitReason(formats strfmt.Registry) error {

	if swag.IsZero(m.WaitReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateWaitReasonEnum("WaitReason", "body", *m.WaitReason); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateWorkflowDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowDefinition) { // not required
		return nil
	}

	if m.WorkflowDefinition != nil {
		if err := m.WorkflowDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WorkflowDefinition")
			}
			return err
		}
	}

	return nil
}

var workflowWorkflowInfoTypeWorkflowMetaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SystemDefined","UserDefined","Dynamic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workflowWorkflowInfoTypeWorkflowMetaTypePropEnum = append(workflowWorkflowInfoTypeWorkflowMetaTypePropEnum, v)
	}
}

// property enum
func (m *WorkflowWorkflowInfo) validateWorkflowMetaTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, workflowWorkflowInfoTypeWorkflowMetaTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WorkflowWorkflowInfo) validateWorkflowMetaType(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowMetaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateWorkflowMetaTypeEnum("WorkflowMetaType", "body", *m.WorkflowMetaType); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateNr0Profile(formats strfmt.Registry) error {

	if swag.IsZero(m.Nr0Profile) { // not required
		return nil
	}

	if m.Nr0Profile != nil {
		if err := m.Nr0Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_0_Profile")
			}
			return err
		}
	}

	return nil
}

func (m *WorkflowWorkflowInfo) validateNr1ClusterProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Nr1ClusterProfile) { // not required
		return nil
	}

	if m.Nr1ClusterProfile != nil {
		if err := m.Nr1ClusterProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_1_ClusterProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowWorkflowInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowWorkflowInfo) UnmarshalBinary(b []byte) error {
	var res WorkflowWorkflowInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
