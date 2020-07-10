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

// ConnectorDownloadStatusAllOf Definition of the list of properties defined in 'connector.DownloadStatus', excluding properties defined in parent classes.
type ConnectorDownloadStatusAllOf struct {
	Checksum *ConnectorFileChecksum `json:"Checksum,omitempty"`
	// Any error encountered. Set to empty when download is in progress or completed.} type: string
	DownloadError interface{} `json:"DownloadError,omitempty"`
	// The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.
	DownloadProgress *int64 `json:"DownloadProgress,omitempty"`
	// The number of retries the plugin attempted before succeeding or failing the download.
	DownloadRetries *int64 `json:"DownloadRetries,omitempty"`
	// The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file.
	Sha256checksum *string `json:"Sha256checksum,omitempty"`
}

// NewConnectorDownloadStatusAllOf instantiates a new ConnectorDownloadStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorDownloadStatusAllOf() *ConnectorDownloadStatusAllOf {
	this := ConnectorDownloadStatusAllOf{}
	return &this
}

// NewConnectorDownloadStatusAllOfWithDefaults instantiates a new ConnectorDownloadStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorDownloadStatusAllOfWithDefaults() *ConnectorDownloadStatusAllOf {
	this := ConnectorDownloadStatusAllOf{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *ConnectorDownloadStatusAllOf) GetChecksum() ConnectorFileChecksum {
	if o == nil || o.Checksum == nil {
		var ret ConnectorFileChecksum
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatusAllOf) GetChecksumOk() (*ConnectorFileChecksum, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *ConnectorDownloadStatusAllOf) HasChecksum() bool {
	if o != nil && o.Checksum != nil {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given ConnectorFileChecksum and assigns it to the Checksum field.
func (o *ConnectorDownloadStatusAllOf) SetChecksum(v ConnectorFileChecksum) {
	o.Checksum = &v
}

// GetDownloadError returns the DownloadError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorDownloadStatusAllOf) GetDownloadError() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DownloadError
}

// GetDownloadErrorOk returns a tuple with the DownloadError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorDownloadStatusAllOf) GetDownloadErrorOk() (*interface{}, bool) {
	if o == nil || o.DownloadError == nil {
		return nil, false
	}
	return &o.DownloadError, true
}

// HasDownloadError returns a boolean if a field has been set.
func (o *ConnectorDownloadStatusAllOf) HasDownloadError() bool {
	if o != nil && o.DownloadError != nil {
		return true
	}

	return false
}

// SetDownloadError gets a reference to the given interface{} and assigns it to the DownloadError field.
func (o *ConnectorDownloadStatusAllOf) SetDownloadError(v interface{}) {
	o.DownloadError = v
}

// GetDownloadProgress returns the DownloadProgress field value if set, zero value otherwise.
func (o *ConnectorDownloadStatusAllOf) GetDownloadProgress() int64 {
	if o == nil || o.DownloadProgress == nil {
		var ret int64
		return ret
	}
	return *o.DownloadProgress
}

// GetDownloadProgressOk returns a tuple with the DownloadProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatusAllOf) GetDownloadProgressOk() (*int64, bool) {
	if o == nil || o.DownloadProgress == nil {
		return nil, false
	}
	return o.DownloadProgress, true
}

// HasDownloadProgress returns a boolean if a field has been set.
func (o *ConnectorDownloadStatusAllOf) HasDownloadProgress() bool {
	if o != nil && o.DownloadProgress != nil {
		return true
	}

	return false
}

// SetDownloadProgress gets a reference to the given int64 and assigns it to the DownloadProgress field.
func (o *ConnectorDownloadStatusAllOf) SetDownloadProgress(v int64) {
	o.DownloadProgress = &v
}

// GetDownloadRetries returns the DownloadRetries field value if set, zero value otherwise.
func (o *ConnectorDownloadStatusAllOf) GetDownloadRetries() int64 {
	if o == nil || o.DownloadRetries == nil {
		var ret int64
		return ret
	}
	return *o.DownloadRetries
}

// GetDownloadRetriesOk returns a tuple with the DownloadRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatusAllOf) GetDownloadRetriesOk() (*int64, bool) {
	if o == nil || o.DownloadRetries == nil {
		return nil, false
	}
	return o.DownloadRetries, true
}

// HasDownloadRetries returns a boolean if a field has been set.
func (o *ConnectorDownloadStatusAllOf) HasDownloadRetries() bool {
	if o != nil && o.DownloadRetries != nil {
		return true
	}

	return false
}

// SetDownloadRetries gets a reference to the given int64 and assigns it to the DownloadRetries field.
func (o *ConnectorDownloadStatusAllOf) SetDownloadRetries(v int64) {
	o.DownloadRetries = &v
}

// GetSha256checksum returns the Sha256checksum field value if set, zero value otherwise.
func (o *ConnectorDownloadStatusAllOf) GetSha256checksum() string {
	if o == nil || o.Sha256checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha256checksum
}

// GetSha256checksumOk returns a tuple with the Sha256checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorDownloadStatusAllOf) GetSha256checksumOk() (*string, bool) {
	if o == nil || o.Sha256checksum == nil {
		return nil, false
	}
	return o.Sha256checksum, true
}

// HasSha256checksum returns a boolean if a field has been set.
func (o *ConnectorDownloadStatusAllOf) HasSha256checksum() bool {
	if o != nil && o.Sha256checksum != nil {
		return true
	}

	return false
}

// SetSha256checksum gets a reference to the given string and assigns it to the Sha256checksum field.
func (o *ConnectorDownloadStatusAllOf) SetSha256checksum(v string) {
	o.Sha256checksum = &v
}

func (o ConnectorDownloadStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Checksum != nil {
		toSerialize["Checksum"] = o.Checksum
	}
	if o.DownloadError != nil {
		toSerialize["DownloadError"] = o.DownloadError
	}
	if o.DownloadProgress != nil {
		toSerialize["DownloadProgress"] = o.DownloadProgress
	}
	if o.DownloadRetries != nil {
		toSerialize["DownloadRetries"] = o.DownloadRetries
	}
	if o.Sha256checksum != nil {
		toSerialize["Sha256checksum"] = o.Sha256checksum
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorDownloadStatusAllOf struct {
	value *ConnectorDownloadStatusAllOf
	isSet bool
}

func (v NullableConnectorDownloadStatusAllOf) Get() *ConnectorDownloadStatusAllOf {
	return v.value
}

func (v *NullableConnectorDownloadStatusAllOf) Set(val *ConnectorDownloadStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorDownloadStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorDownloadStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorDownloadStatusAllOf(val *ConnectorDownloadStatusAllOf) *NullableConnectorDownloadStatusAllOf {
	return &NullableConnectorDownloadStatusAllOf{value: val, isSet: true}
}

func (v NullableConnectorDownloadStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorDownloadStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
