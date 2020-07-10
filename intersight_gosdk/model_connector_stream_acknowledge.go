/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-08T07:46:03Z.
 *
 * API version: 0.0.1-15983
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ConnectorStreamAcknowledge Acknowledge a set of messages read from the device, on receipt device will drop the acknowledged messages from its cache.
type ConnectorStreamAcknowledge struct {
	ConnectorStreamMessage
	// The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache.
	AckSequence *int64 `json:"AckSequence,omitempty"`
}

// NewConnectorStreamAcknowledge instantiates a new ConnectorStreamAcknowledge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamAcknowledge() *ConnectorStreamAcknowledge {
	this := ConnectorStreamAcknowledge{}
	return &this
}

// NewConnectorStreamAcknowledgeWithDefaults instantiates a new ConnectorStreamAcknowledge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamAcknowledgeWithDefaults() *ConnectorStreamAcknowledge {
	this := ConnectorStreamAcknowledge{}
	return &this
}

// GetAckSequence returns the AckSequence field value if set, zero value otherwise.
func (o *ConnectorStreamAcknowledge) GetAckSequence() int64 {
	if o == nil || o.AckSequence == nil {
		var ret int64
		return ret
	}
	return *o.AckSequence
}

// GetAckSequenceOk returns a tuple with the AckSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamAcknowledge) GetAckSequenceOk() (*int64, bool) {
	if o == nil || o.AckSequence == nil {
		return nil, false
	}
	return o.AckSequence, true
}

// HasAckSequence returns a boolean if a field has been set.
func (o *ConnectorStreamAcknowledge) HasAckSequence() bool {
	if o != nil && o.AckSequence != nil {
		return true
	}

	return false
}

// SetAckSequence gets a reference to the given int64 and assigns it to the AckSequence field.
func (o *ConnectorStreamAcknowledge) SetAckSequence(v int64) {
	o.AckSequence = &v
}

func (o ConnectorStreamAcknowledge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	if o.AckSequence != nil {
		toSerialize["AckSequence"] = o.AckSequence
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorStreamAcknowledge struct {
	value *ConnectorStreamAcknowledge
	isSet bool
}

func (v NullableConnectorStreamAcknowledge) Get() *ConnectorStreamAcknowledge {
	return v.value
}

func (v *NullableConnectorStreamAcknowledge) Set(val *ConnectorStreamAcknowledge) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamAcknowledge) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamAcknowledge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamAcknowledge(val *ConnectorStreamAcknowledge) *NullableConnectorStreamAcknowledge {
	return &NullableConnectorStreamAcknowledge{value: val, isSet: true}
}

func (v NullableConnectorStreamAcknowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamAcknowledge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
