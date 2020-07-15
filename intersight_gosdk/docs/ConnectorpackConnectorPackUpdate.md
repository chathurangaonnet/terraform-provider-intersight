# ConnectorpackConnectorPackUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | Version of connector pack currently running in UCS Director. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the connector pack. | [optional] [readonly] 
**NewVersion** | Pointer to **string** | Version of connector pack to be installed in the next upgrade cycle. | [optional] [readonly] 

## Methods

### NewConnectorpackConnectorPackUpdate

`func NewConnectorpackConnectorPackUpdate() *ConnectorpackConnectorPackUpdate`

NewConnectorpackConnectorPackUpdate instantiates a new ConnectorpackConnectorPackUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorpackConnectorPackUpdateWithDefaults

`func NewConnectorpackConnectorPackUpdateWithDefaults() *ConnectorpackConnectorPackUpdate`

NewConnectorpackConnectorPackUpdateWithDefaults instantiates a new ConnectorpackConnectorPackUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ConnectorpackConnectorPackUpdate) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ConnectorpackConnectorPackUpdate) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetName

`func (o *ConnectorpackConnectorPackUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorpackConnectorPackUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorpackConnectorPackUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorpackConnectorPackUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewVersion

`func (o *ConnectorpackConnectorPackUpdate) GetNewVersion() string`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *ConnectorpackConnectorPackUpdate) GetNewVersionOk() (*string, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *ConnectorpackConnectorPackUpdate) SetNewVersion(v string)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *ConnectorpackConnectorPackUpdate) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


