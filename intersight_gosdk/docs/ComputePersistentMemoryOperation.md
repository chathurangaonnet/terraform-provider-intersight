# ComputePersistentMemoryOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminAction** | Pointer to **string** | Administrative actions that can be performed on the Persistent Memory Modules. | [optional] [default to "None"]
**IsSecurePassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;securePassphrase&#39; property has been set. | [optional] [readonly] 
**Modules** | Pointer to [**[]ComputePersistentMemoryModule**](compute.PersistentMemoryModule.md) |  | [optional] 
**SecurePassphrase** | Pointer to **string** | Secure passphrase of the Persistent Memory Modules of the server. | [optional] 

## Methods

### NewComputePersistentMemoryOperation

`func NewComputePersistentMemoryOperation() *ComputePersistentMemoryOperation`

NewComputePersistentMemoryOperation instantiates a new ComputePersistentMemoryOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePersistentMemoryOperationWithDefaults

`func NewComputePersistentMemoryOperationWithDefaults() *ComputePersistentMemoryOperation`

NewComputePersistentMemoryOperationWithDefaults instantiates a new ComputePersistentMemoryOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminAction

`func (o *ComputePersistentMemoryOperation) GetAdminAction() string`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *ComputePersistentMemoryOperation) GetAdminActionOk() (*string, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *ComputePersistentMemoryOperation) SetAdminAction(v string)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *ComputePersistentMemoryOperation) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperation) GetIsSecurePassphraseSet() bool`

GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field if non-nil, zero value otherwise.

### GetIsSecurePassphraseSetOk

`func (o *ComputePersistentMemoryOperation) GetIsSecurePassphraseSetOk() (*bool, bool)`

GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperation) SetIsSecurePassphraseSet(v bool)`

SetIsSecurePassphraseSet sets IsSecurePassphraseSet field to given value.

### HasIsSecurePassphraseSet

`func (o *ComputePersistentMemoryOperation) HasIsSecurePassphraseSet() bool`

HasIsSecurePassphraseSet returns a boolean if a field has been set.

### GetModules

`func (o *ComputePersistentMemoryOperation) GetModules() []ComputePersistentMemoryModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ComputePersistentMemoryOperation) GetModulesOk() (*[]ComputePersistentMemoryModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ComputePersistentMemoryOperation) SetModules(v []ComputePersistentMemoryModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ComputePersistentMemoryOperation) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetSecurePassphrase

`func (o *ComputePersistentMemoryOperation) GetSecurePassphrase() string`

GetSecurePassphrase returns the SecurePassphrase field if non-nil, zero value otherwise.

### GetSecurePassphraseOk

`func (o *ComputePersistentMemoryOperation) GetSecurePassphraseOk() (*string, bool)`

GetSecurePassphraseOk returns a tuple with the SecurePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePassphrase

`func (o *ComputePersistentMemoryOperation) SetSecurePassphrase(v string)`

SetSecurePassphrase sets SecurePassphrase field to given value.

### HasSecurePassphrase

`func (o *ComputePersistentMemoryOperation) HasSecurePassphrase() bool`

HasSecurePassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


