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

// IamResourceRolesResponse The response body of a HTTP GET request for the 'iam.ResourceRoles' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'iam.ResourceRoles' resources.
type IamResourceRolesResponse struct {
	IamResourceRolesResponseInterface interface{ GetObjectType() string }
}

func (s IamResourceRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.IamResourceRolesResponseInterface)
}

func (s *IamResourceRolesResponse) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["ObjectType"]; ok {
		switch v {
		default:
			return fmt.Errorf("No oneOf model has 'ObjectType' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'ObjectType' not found in unmarshaled payload: %+v", unmarshaled)
	}
}

type NullableIamResourceRolesResponse struct {
	value *IamResourceRolesResponse
	isSet bool
}

func (v NullableIamResourceRolesResponse) Get() *IamResourceRolesResponse {
	return v.value
}

func (v *NullableIamResourceRolesResponse) Set(val *IamResourceRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourceRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourceRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourceRolesResponse(val *IamResourceRolesResponse) *NullableIamResourceRolesResponse {
	return &NullableIamResourceRolesResponse{value: val, isSet: true}
}

func (v NullableIamResourceRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourceRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
