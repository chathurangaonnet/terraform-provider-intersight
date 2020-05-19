/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-12T21:47:47-07:00.
 *
 * API version: 1.0.9-1617
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"fmt"
)

// HyperflexSysConfigPolicyRelationship A relationship to the 'hyperflex.SysConfigPolicy' resource, or the expanded 'hyperflex.SysConfigPolicy' resource, or the 'null' value.
type HyperflexSysConfigPolicyRelationship struct {
	HyperflexSysConfigPolicyRelationshipInterface interface{}
}

func (s HyperflexSysConfigPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.HyperflexSysConfigPolicyRelationshipInterface)
}

func (s *HyperflexSysConfigPolicyRelationship) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaledHyperflexSysConfigPolicy *HyperflexSysConfigPolicy = &HyperflexSysConfigPolicy{}
	err = json.Unmarshal(src, unmarshaledHyperflexSysConfigPolicy)
	if err == nil {
		s.HyperflexSysConfigPolicyRelationshipInterface = unmarshaledHyperflexSysConfigPolicy
		return nil
	}
	var unmarshaledMoMoRef *MoMoRef = &MoMoRef{}
	err = json.Unmarshal(src, unmarshaledMoMoRef)
	if err == nil {
		s.HyperflexSysConfigPolicyRelationshipInterface = unmarshaledMoMoRef
		return nil
	}
	return fmt.Errorf("No oneOf model could be deserialized from payload: %s", string(src))
}

type NullableHyperflexSysConfigPolicyRelationship struct {
	value *HyperflexSysConfigPolicyRelationship
	isSet bool
}

func (v NullableHyperflexSysConfigPolicyRelationship) Get() *HyperflexSysConfigPolicyRelationship {
	return v.value
}

func (v *NullableHyperflexSysConfigPolicyRelationship) Set(val *HyperflexSysConfigPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSysConfigPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSysConfigPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSysConfigPolicyRelationship(val *HyperflexSysConfigPolicyRelationship) *NullableHyperflexSysConfigPolicyRelationship {
	return &NullableHyperflexSysConfigPolicyRelationship{value: val, isSet: true}
}

func (v NullableHyperflexSysConfigPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSysConfigPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
