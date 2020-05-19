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
)

// SyslogRemoteClientBaseAllOf Definition of the list of properties defined in 'syslog.RemoteClientBase', excluding properties defined in parent classes.
type SyslogRemoteClientBaseAllOf struct {
	// Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field.
	Enabled *bool `json:"Enabled,omitempty"`
	// Hostname or IP Address of the syslog server where log should be stored.
	Hostname *string `json:"Hostname,omitempty"`
	// Lowest level of messages to be included in the remote log.
	MinSeverity *string `json:"MinSeverity,omitempty"`
	// Port number used for logging on syslog server.
	Port *int64 `json:"Port,omitempty"`
	// Transport layer protocol for transmission of log messages to syslog server.
	Protocol *string `json:"Protocol,omitempty"`
}

// NewSyslogRemoteClientBaseAllOf instantiates a new SyslogRemoteClientBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogRemoteClientBaseAllOf() *SyslogRemoteClientBaseAllOf {
	this := SyslogRemoteClientBaseAllOf{}
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	var protocol string = "udp"
	this.Protocol = &protocol
	return &this
}

// NewSyslogRemoteClientBaseAllOfWithDefaults instantiates a new SyslogRemoteClientBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogRemoteClientBaseAllOfWithDefaults() *SyslogRemoteClientBaseAllOf {
	this := SyslogRemoteClientBaseAllOf{}
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	var protocol string = "udp"
	this.Protocol = &protocol
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SyslogRemoteClientBaseAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBaseAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyslogRemoteClientBaseAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SyslogRemoteClientBaseAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyslogRemoteClientBaseAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBaseAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyslogRemoteClientBaseAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyslogRemoteClientBaseAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SyslogRemoteClientBaseAllOf) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBaseAllOf) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SyslogRemoteClientBaseAllOf) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SyslogRemoteClientBaseAllOf) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyslogRemoteClientBaseAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBaseAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyslogRemoteClientBaseAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SyslogRemoteClientBaseAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyslogRemoteClientBaseAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBaseAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyslogRemoteClientBaseAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyslogRemoteClientBaseAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

func (o SyslogRemoteClientBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.MinSeverity != nil {
		toSerialize["MinSeverity"] = o.MinSeverity
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableSyslogRemoteClientBaseAllOf struct {
	value *SyslogRemoteClientBaseAllOf
	isSet bool
}

func (v NullableSyslogRemoteClientBaseAllOf) Get() *SyslogRemoteClientBaseAllOf {
	return v.value
}

func (v *NullableSyslogRemoteClientBaseAllOf) Set(val *SyslogRemoteClientBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogRemoteClientBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogRemoteClientBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogRemoteClientBaseAllOf(val *SyslogRemoteClientBaseAllOf) *NullableSyslogRemoteClientBaseAllOf {
	return &NullableSyslogRemoteClientBaseAllOf{value: val, isSet: true}
}

func (v NullableSyslogRemoteClientBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogRemoteClientBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
