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

// AssetConnectionControlMessageAllOf Definition of the list of properties defined in 'asset.ConnectionControlMessage', excluding properties defined in parent classes.
type AssetConnectionControlMessageAllOf struct {
	// The account id to which the device belongs.
	Account *string `json:"Account,omitempty"`
	// The version of the device connector currently running on the platform. Deprecated by newer connectors that will report this directly to the device connector gateway in a websocket header, but included to continue to support older versions which report any version change after connect.
	ConnectorVersion *string `json:"ConnectorVersion,omitempty"`
	// The Moid of the device under change. Used to route the message to a devices connection.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The domain group id to which the device belongs.
	DomainGroup *string `json:"DomainGroup,omitempty"`
	// Flag to force any open connections to be evicted. Used in case device has been deleted or blacklisted.
	Evict *bool `json:"Evict,omitempty"`
	// The current leadership of a device cluster member.
	Leadership *string `json:"Leadership,omitempty"`
	// The new identity assigned to a device on ownership change (claim/unclaim).
	NewIdentity *string `json:"NewIdentity,omitempty"`
	// The partition the device was last connected to, used to address the control message to the device connector gateway instance holding the devices connection.
	Partition *int64 `json:"Partition,omitempty"`
}

// NewAssetConnectionControlMessageAllOf instantiates a new AssetConnectionControlMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetConnectionControlMessageAllOf() *AssetConnectionControlMessageAllOf {
	this := AssetConnectionControlMessageAllOf{}
	var leadership string = "Unknown"
	this.Leadership = &leadership
	return &this
}

// NewAssetConnectionControlMessageAllOfWithDefaults instantiates a new AssetConnectionControlMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetConnectionControlMessageAllOfWithDefaults() *AssetConnectionControlMessageAllOf {
	this := AssetConnectionControlMessageAllOf{}
	var leadership string = "Unknown"
	this.Leadership = &leadership
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *AssetConnectionControlMessageAllOf) SetAccount(v string) {
	o.Account = &v
}

// GetConnectorVersion returns the ConnectorVersion field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetConnectorVersion() string {
	if o == nil || o.ConnectorVersion == nil {
		var ret string
		return ret
	}
	return *o.ConnectorVersion
}

// GetConnectorVersionOk returns a tuple with the ConnectorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetConnectorVersionOk() (*string, bool) {
	if o == nil || o.ConnectorVersion == nil {
		return nil, false
	}
	return o.ConnectorVersion, true
}

// HasConnectorVersion returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasConnectorVersion() bool {
	if o != nil && o.ConnectorVersion != nil {
		return true
	}

	return false
}

// SetConnectorVersion gets a reference to the given string and assigns it to the ConnectorVersion field.
func (o *AssetConnectionControlMessageAllOf) SetConnectorVersion(v string) {
	o.ConnectorVersion = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AssetConnectionControlMessageAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDomainGroup returns the DomainGroup field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetDomainGroup() string {
	if o == nil || o.DomainGroup == nil {
		var ret string
		return ret
	}
	return *o.DomainGroup
}

// GetDomainGroupOk returns a tuple with the DomainGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetDomainGroupOk() (*string, bool) {
	if o == nil || o.DomainGroup == nil {
		return nil, false
	}
	return o.DomainGroup, true
}

// HasDomainGroup returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasDomainGroup() bool {
	if o != nil && o.DomainGroup != nil {
		return true
	}

	return false
}

// SetDomainGroup gets a reference to the given string and assigns it to the DomainGroup field.
func (o *AssetConnectionControlMessageAllOf) SetDomainGroup(v string) {
	o.DomainGroup = &v
}

// GetEvict returns the Evict field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetEvict() bool {
	if o == nil || o.Evict == nil {
		var ret bool
		return ret
	}
	return *o.Evict
}

// GetEvictOk returns a tuple with the Evict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetEvictOk() (*bool, bool) {
	if o == nil || o.Evict == nil {
		return nil, false
	}
	return o.Evict, true
}

// HasEvict returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasEvict() bool {
	if o != nil && o.Evict != nil {
		return true
	}

	return false
}

// SetEvict gets a reference to the given bool and assigns it to the Evict field.
func (o *AssetConnectionControlMessageAllOf) SetEvict(v bool) {
	o.Evict = &v
}

// GetLeadership returns the Leadership field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetLeadership() string {
	if o == nil || o.Leadership == nil {
		var ret string
		return ret
	}
	return *o.Leadership
}

// GetLeadershipOk returns a tuple with the Leadership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetLeadershipOk() (*string, bool) {
	if o == nil || o.Leadership == nil {
		return nil, false
	}
	return o.Leadership, true
}

// HasLeadership returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasLeadership() bool {
	if o != nil && o.Leadership != nil {
		return true
	}

	return false
}

// SetLeadership gets a reference to the given string and assigns it to the Leadership field.
func (o *AssetConnectionControlMessageAllOf) SetLeadership(v string) {
	o.Leadership = &v
}

// GetNewIdentity returns the NewIdentity field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetNewIdentity() string {
	if o == nil || o.NewIdentity == nil {
		var ret string
		return ret
	}
	return *o.NewIdentity
}

// GetNewIdentityOk returns a tuple with the NewIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetNewIdentityOk() (*string, bool) {
	if o == nil || o.NewIdentity == nil {
		return nil, false
	}
	return o.NewIdentity, true
}

// HasNewIdentity returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasNewIdentity() bool {
	if o != nil && o.NewIdentity != nil {
		return true
	}

	return false
}

// SetNewIdentity gets a reference to the given string and assigns it to the NewIdentity field.
func (o *AssetConnectionControlMessageAllOf) SetNewIdentity(v string) {
	o.NewIdentity = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *AssetConnectionControlMessageAllOf) GetPartition() int64 {
	if o == nil || o.Partition == nil {
		var ret int64
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetConnectionControlMessageAllOf) GetPartitionOk() (*int64, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *AssetConnectionControlMessageAllOf) HasPartition() bool {
	if o != nil && o.Partition != nil {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int64 and assigns it to the Partition field.
func (o *AssetConnectionControlMessageAllOf) SetPartition(v int64) {
	o.Partition = &v
}

func (o AssetConnectionControlMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ConnectorVersion != nil {
		toSerialize["ConnectorVersion"] = o.ConnectorVersion
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.DomainGroup != nil {
		toSerialize["DomainGroup"] = o.DomainGroup
	}
	if o.Evict != nil {
		toSerialize["Evict"] = o.Evict
	}
	if o.Leadership != nil {
		toSerialize["Leadership"] = o.Leadership
	}
	if o.NewIdentity != nil {
		toSerialize["NewIdentity"] = o.NewIdentity
	}
	if o.Partition != nil {
		toSerialize["Partition"] = o.Partition
	}
	return json.Marshal(toSerialize)
}

type NullableAssetConnectionControlMessageAllOf struct {
	value *AssetConnectionControlMessageAllOf
	isSet bool
}

func (v NullableAssetConnectionControlMessageAllOf) Get() *AssetConnectionControlMessageAllOf {
	return v.value
}

func (v *NullableAssetConnectionControlMessageAllOf) Set(val *AssetConnectionControlMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetConnectionControlMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetConnectionControlMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetConnectionControlMessageAllOf(val *AssetConnectionControlMessageAllOf) *NullableAssetConnectionControlMessageAllOf {
	return &NullableAssetConnectionControlMessageAllOf{value: val, isSet: true}
}

func (v NullableAssetConnectionControlMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetConnectionControlMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
