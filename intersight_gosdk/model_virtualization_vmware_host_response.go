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

// VirtualizationVmwareHostResponse The response body of a HTTP GET request for the 'virtualization.VmwareHost' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'virtualization.VmwareHost' resources.
type VirtualizationVmwareHostResponse struct {
	VirtualizationVmwareHostResponseInterface interface{ GetObjectType() string }
}

func (s VirtualizationVmwareHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.VirtualizationVmwareHostResponseInterface)
}

func (s *VirtualizationVmwareHostResponse) UnmarshalJSON(src []byte) error {
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

type NullableVirtualizationVmwareHostResponse struct {
	value *VirtualizationVmwareHostResponse
	isSet bool
}

func (v NullableVirtualizationVmwareHostResponse) Get() *VirtualizationVmwareHostResponse {
	return v.value
}

func (v *NullableVirtualizationVmwareHostResponse) Set(val *VirtualizationVmwareHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareHostResponse(val *VirtualizationVmwareHostResponse) *NullableVirtualizationVmwareHostResponse {
	return &NullableVirtualizationVmwareHostResponse{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
