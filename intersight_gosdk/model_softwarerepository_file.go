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
	"time"
)

// SoftwarerepositoryFile A file that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository. Instances of this MO will be implicitly created as part of the file import operation.
type SoftwarerepositoryFile struct {
	MoBaseMo
	// User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.
	Description *string `json:"Description,omitempty"`
	// The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.
	DownloadCount *int64 `json:"DownloadCount,omitempty"`
	// The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source.
	ImportAction *string `json:"ImportAction,omitempty"`
	// The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process.
	ImportState *string `json:"ImportState,omitempty"`
	// The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached.
	ImportedTime *time.Time `json:"ImportedTime,omitempty"`
	// The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Md5sum *string `json:"Md5sum,omitempty"`
	// The name of the file. It is populated as part of the image import operation.
	Name *string `json:"Name,omitempty"`
	// The date on which the file was released or distributed by its vendor.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
	// The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Sha512sum *string `json:"Sha512sum,omitempty"`
	// The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Size *int64 `json:"Size,omitempty"`
	// The software advisory, if any, provided by the vendor for this file.
	SoftwareAdvisoryUrl *string                       `json:"SoftwareAdvisoryUrl,omitempty"`
	Source              *SoftwarerepositoryFileServer `json:"Source,omitempty"`
	// Vendor provided version for the file.
	Version *string `json:"Version,omitempty"`
}

// NewSoftwarerepositoryFile instantiates a new SoftwarerepositoryFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryFile() *SoftwarerepositoryFile {
	this := SoftwarerepositoryFile{}
	var importAction string = "None"
	this.ImportAction = &importAction
	var importState string = "ReadyForImport"
	this.ImportState = &importState
	return &this
}

// NewSoftwarerepositoryFileWithDefaults instantiates a new SoftwarerepositoryFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryFileWithDefaults() *SoftwarerepositoryFile {
	this := SoftwarerepositoryFile{}
	var importAction string = "None"
	this.ImportAction = &importAction
	var importState string = "ReadyForImport"
	this.ImportState = &importState
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SoftwarerepositoryFile) SetDescription(v string) {
	o.Description = &v
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetDownloadCount() int64 {
	if o == nil || o.DownloadCount == nil {
		var ret int64
		return ret
	}
	return *o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetDownloadCountOk() (*int64, bool) {
	if o == nil || o.DownloadCount == nil {
		return nil, false
	}
	return o.DownloadCount, true
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasDownloadCount() bool {
	if o != nil && o.DownloadCount != nil {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given int64 and assigns it to the DownloadCount field.
func (o *SoftwarerepositoryFile) SetDownloadCount(v int64) {
	o.DownloadCount = &v
}

// GetImportAction returns the ImportAction field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetImportAction() string {
	if o == nil || o.ImportAction == nil {
		var ret string
		return ret
	}
	return *o.ImportAction
}

// GetImportActionOk returns a tuple with the ImportAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetImportActionOk() (*string, bool) {
	if o == nil || o.ImportAction == nil {
		return nil, false
	}
	return o.ImportAction, true
}

// HasImportAction returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasImportAction() bool {
	if o != nil && o.ImportAction != nil {
		return true
	}

	return false
}

// SetImportAction gets a reference to the given string and assigns it to the ImportAction field.
func (o *SoftwarerepositoryFile) SetImportAction(v string) {
	o.ImportAction = &v
}

// GetImportState returns the ImportState field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetImportState() string {
	if o == nil || o.ImportState == nil {
		var ret string
		return ret
	}
	return *o.ImportState
}

// GetImportStateOk returns a tuple with the ImportState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetImportStateOk() (*string, bool) {
	if o == nil || o.ImportState == nil {
		return nil, false
	}
	return o.ImportState, true
}

// HasImportState returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasImportState() bool {
	if o != nil && o.ImportState != nil {
		return true
	}

	return false
}

// SetImportState gets a reference to the given string and assigns it to the ImportState field.
func (o *SoftwarerepositoryFile) SetImportState(v string) {
	o.ImportState = &v
}

// GetImportedTime returns the ImportedTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetImportedTime() time.Time {
	if o == nil || o.ImportedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ImportedTime
}

// GetImportedTimeOk returns a tuple with the ImportedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetImportedTimeOk() (*time.Time, bool) {
	if o == nil || o.ImportedTime == nil {
		return nil, false
	}
	return o.ImportedTime, true
}

// HasImportedTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasImportedTime() bool {
	if o != nil && o.ImportedTime != nil {
		return true
	}

	return false
}

// SetImportedTime gets a reference to the given time.Time and assigns it to the ImportedTime field.
func (o *SoftwarerepositoryFile) SetImportedTime(v time.Time) {
	o.ImportedTime = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetLastAccessTime() time.Time {
	if o == nil || o.LastAccessTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *SoftwarerepositoryFile) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetMd5sum() string {
	if o == nil || o.Md5sum == nil {
		var ret string
		return ret
	}
	return *o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetMd5sumOk() (*string, bool) {
	if o == nil || o.Md5sum == nil {
		return nil, false
	}
	return o.Md5sum, true
}

// HasMd5sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasMd5sum() bool {
	if o != nil && o.Md5sum != nil {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given string and assigns it to the Md5sum field.
func (o *SoftwarerepositoryFile) SetMd5sum(v string) {
	o.Md5sum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwarerepositoryFile) SetName(v string) {
	o.Name = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *SoftwarerepositoryFile) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetSha512sum returns the Sha512sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetSha512sum() string {
	if o == nil || o.Sha512sum == nil {
		var ret string
		return ret
	}
	return *o.Sha512sum
}

// GetSha512sumOk returns a tuple with the Sha512sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetSha512sumOk() (*string, bool) {
	if o == nil || o.Sha512sum == nil {
		return nil, false
	}
	return o.Sha512sum, true
}

// HasSha512sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasSha512sum() bool {
	if o != nil && o.Sha512sum != nil {
		return true
	}

	return false
}

// SetSha512sum gets a reference to the given string and assigns it to the Sha512sum field.
func (o *SoftwarerepositoryFile) SetSha512sum(v string) {
	o.Sha512sum = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SoftwarerepositoryFile) SetSize(v int64) {
	o.Size = &v
}

// GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetSoftwareAdvisoryUrl() string {
	if o == nil || o.SoftwareAdvisoryUrl == nil {
		var ret string
		return ret
	}
	return *o.SoftwareAdvisoryUrl
}

// GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetSoftwareAdvisoryUrlOk() (*string, bool) {
	if o == nil || o.SoftwareAdvisoryUrl == nil {
		return nil, false
	}
	return o.SoftwareAdvisoryUrl, true
}

// HasSoftwareAdvisoryUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasSoftwareAdvisoryUrl() bool {
	if o != nil && o.SoftwareAdvisoryUrl != nil {
		return true
	}

	return false
}

// SetSoftwareAdvisoryUrl gets a reference to the given string and assigns it to the SoftwareAdvisoryUrl field.
func (o *SoftwarerepositoryFile) SetSoftwareAdvisoryUrl(v string) {
	o.SoftwareAdvisoryUrl = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetSource() SoftwarerepositoryFileServer {
	if o == nil || o.Source == nil {
		var ret SoftwarerepositoryFileServer
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetSourceOk() (*SoftwarerepositoryFileServer, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given SoftwarerepositoryFileServer and assigns it to the Source field.
func (o *SoftwarerepositoryFile) SetSource(v SoftwarerepositoryFileServer) {
	o.Source = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryFile) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFile) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryFile) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryFile) SetVersion(v string) {
	o.Version = &v
}

func (o SoftwarerepositoryFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DownloadCount != nil {
		toSerialize["DownloadCount"] = o.DownloadCount
	}
	if o.ImportAction != nil {
		toSerialize["ImportAction"] = o.ImportAction
	}
	if o.ImportState != nil {
		toSerialize["ImportState"] = o.ImportState
	}
	if o.ImportedTime != nil {
		toSerialize["ImportedTime"] = o.ImportedTime
	}
	if o.LastAccessTime != nil {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if o.Md5sum != nil {
		toSerialize["Md5sum"] = o.Md5sum
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReleaseDate != nil {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if o.Sha512sum != nil {
		toSerialize["Sha512sum"] = o.Sha512sum
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.SoftwareAdvisoryUrl != nil {
		toSerialize["SoftwareAdvisoryUrl"] = o.SoftwareAdvisoryUrl
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryFile struct {
	value *SoftwarerepositoryFile
	isSet bool
}

func (v NullableSoftwarerepositoryFile) Get() *SoftwarerepositoryFile {
	return v.value
}

func (v *NullableSoftwarerepositoryFile) Set(val *SoftwarerepositoryFile) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryFile) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryFile(val *SoftwarerepositoryFile) *NullableSoftwarerepositoryFile {
	return &NullableSoftwarerepositoryFile{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
