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

// RecoveryOnDemandBackupRelationship A relationship to the 'recovery.OnDemandBackup' resource, or the expanded 'recovery.OnDemandBackup' resource, or the 'null' value.
type RecoveryOnDemandBackupRelationship struct {
	RecoveryOnDemandBackupRelationshipInterface interface{}
}

func (s RecoveryOnDemandBackupRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.RecoveryOnDemandBackupRelationshipInterface)
}

func (s *RecoveryOnDemandBackupRelationship) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaledMoMoRef *MoMoRef = &MoMoRef{}
	err = json.Unmarshal(src, unmarshaledMoMoRef)
	if err == nil {
		s.RecoveryOnDemandBackupRelationshipInterface = unmarshaledMoMoRef
		return nil
	}
	var unmarshaledRecoveryOnDemandBackup *RecoveryOnDemandBackup = &RecoveryOnDemandBackup{}
	err = json.Unmarshal(src, unmarshaledRecoveryOnDemandBackup)
	if err == nil {
		s.RecoveryOnDemandBackupRelationshipInterface = unmarshaledRecoveryOnDemandBackup
		return nil
	}
	return fmt.Errorf("No oneOf model could be deserialized from payload: %s", string(src))
}

type NullableRecoveryOnDemandBackupRelationship struct {
	value *RecoveryOnDemandBackupRelationship
	isSet bool
}

func (v NullableRecoveryOnDemandBackupRelationship) Get() *RecoveryOnDemandBackupRelationship {
	return v.value
}

func (v *NullableRecoveryOnDemandBackupRelationship) Set(val *RecoveryOnDemandBackupRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryOnDemandBackupRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryOnDemandBackupRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryOnDemandBackupRelationship(val *RecoveryOnDemandBackupRelationship) *NullableRecoveryOnDemandBackupRelationship {
	return &NullableRecoveryOnDemandBackupRelationship{value: val, isSet: true}
}

func (v NullableRecoveryOnDemandBackupRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryOnDemandBackupRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
