/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PolicyAbstractConfigProfile AbstractConfigProfile is an abstract base type for all config actions on a profile.
type PolicyAbstractConfigProfile struct {
	PolicyAbstractProfile
	// User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
	Action               *string              `json:"Action,omitempty"`
	ConfigContext        *PolicyConfigContext `json:"ConfigContext,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigProfile PolicyAbstractConfigProfile

// NewPolicyAbstractConfigProfile instantiates a new PolicyAbstractConfigProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigProfile() *PolicyAbstractConfigProfile {
	this := PolicyAbstractConfigProfile{}
	return &this
}

// NewPolicyAbstractConfigProfileWithDefaults instantiates a new PolicyAbstractConfigProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigProfileWithDefaults() *PolicyAbstractConfigProfile {
	this := PolicyAbstractConfigProfile{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PolicyAbstractConfigProfile) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigProfile) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PolicyAbstractConfigProfile) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PolicyAbstractConfigProfile) SetAction(v string) {
	o.Action = &v
}

// GetConfigContext returns the ConfigContext field value if set, zero value otherwise.
func (o *PolicyAbstractConfigProfile) GetConfigContext() PolicyConfigContext {
	if o == nil || o.ConfigContext == nil {
		var ret PolicyConfigContext
		return ret
	}
	return *o.ConfigContext
}

// GetConfigContextOk returns a tuple with the ConfigContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigProfile) GetConfigContextOk() (*PolicyConfigContext, bool) {
	if o == nil || o.ConfigContext == nil {
		return nil, false
	}
	return o.ConfigContext, true
}

// HasConfigContext returns a boolean if a field has been set.
func (o *PolicyAbstractConfigProfile) HasConfigContext() bool {
	if o != nil && o.ConfigContext != nil {
		return true
	}

	return false
}

// SetConfigContext gets a reference to the given PolicyConfigContext and assigns it to the ConfigContext field.
func (o *PolicyAbstractConfigProfile) SetConfigContext(v PolicyConfigContext) {
	o.ConfigContext = &v
}

func (o PolicyAbstractConfigProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.ConfigContext != nil {
		toSerialize["ConfigContext"] = o.ConfigContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractConfigProfile) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyAbstractConfigProfileWithoutEmbeddedStruct struct {
		// User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign.
		Action        *string              `json:"Action,omitempty"`
		ConfigContext *PolicyConfigContext `json:"ConfigContext,omitempty"`
	}

	varPolicyAbstractConfigProfileWithoutEmbeddedStruct := PolicyAbstractConfigProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigProfileWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractConfigProfile := _PolicyAbstractConfigProfile{}
		varPolicyAbstractConfigProfile.Action = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.Action
		varPolicyAbstractConfigProfile.ConfigContext = varPolicyAbstractConfigProfileWithoutEmbeddedStruct.ConfigContext
		*o = PolicyAbstractConfigProfile(varPolicyAbstractConfigProfile)
	} else {
		return err
	}

	varPolicyAbstractConfigProfile := _PolicyAbstractConfigProfile{}

	err = json.Unmarshal(bytes, &varPolicyAbstractConfigProfile)
	if err == nil {
		o.PolicyAbstractProfile = varPolicyAbstractConfigProfile.PolicyAbstractProfile
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Action")
		delete(additionalProperties, "ConfigContext")

		// remove fields from embedded structs
		reflectPolicyAbstractProfile := reflect.ValueOf(o.PolicyAbstractProfile)
		for i := 0; i < reflectPolicyAbstractProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractProfile.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAbstractConfigProfile struct {
	value *PolicyAbstractConfigProfile
	isSet bool
}

func (v NullablePolicyAbstractConfigProfile) Get() *PolicyAbstractConfigProfile {
	return v.value
}

func (v *NullablePolicyAbstractConfigProfile) Set(val *PolicyAbstractConfigProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigProfile(val *PolicyAbstractConfigProfile) *NullablePolicyAbstractConfigProfile {
	return &NullablePolicyAbstractConfigProfile{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
