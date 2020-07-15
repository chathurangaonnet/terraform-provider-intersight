# MemoryPersistentMemoryLogicalNamespaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** | Capacity of this Namespace that is created or modified. | [optional] 
**Mode** | Pointer to **string** | Mode of this Namespace that is created or modified. | [optional] [default to "raw"]
**Name** | Pointer to **string** | Name of this Namespace to be created on the server. | [optional] 
**SocketId** | Pointer to **int32** | Socket ID of the region on which this Namespace has to be created or modified. | [optional] [default to 1]
**SocketMemoryId** | Pointer to **string** | Socket Memory ID of the region on which this Namespace has to be created or modified. | [optional] [default to "Not Applicable"]

## Methods

### NewMemoryPersistentMemoryLogicalNamespaceAllOf

`func NewMemoryPersistentMemoryLogicalNamespaceAllOf() *MemoryPersistentMemoryLogicalNamespaceAllOf`

NewMemoryPersistentMemoryLogicalNamespaceAllOf instantiates a new MemoryPersistentMemoryLogicalNamespaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPersistentMemoryLogicalNamespaceAllOfWithDefaults

`func NewMemoryPersistentMemoryLogicalNamespaceAllOfWithDefaults() *MemoryPersistentMemoryLogicalNamespaceAllOf`

NewMemoryPersistentMemoryLogicalNamespaceAllOfWithDefaults instantiates a new MemoryPersistentMemoryLogicalNamespaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSocketId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketId() int32`

GetSocketId returns the SocketId field if non-nil, zero value otherwise.

### GetSocketIdOk

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketIdOk() (*int32, bool)`

GetSocketIdOk returns a tuple with the SocketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetSocketId(v int32)`

SetSocketId sets SocketId field to given value.

### HasSocketId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasSocketId() bool`

HasSocketId returns a boolean if a field has been set.

### GetSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketMemoryId() string`

GetSocketMemoryId returns the SocketMemoryId field if non-nil, zero value otherwise.

### GetSocketMemoryIdOk

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) GetSocketMemoryIdOk() (*string, bool)`

GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) SetSocketMemoryId(v string)`

SetSocketMemoryId sets SocketMemoryId field to given value.

### HasSocketMemoryId

`func (o *MemoryPersistentMemoryLogicalNamespaceAllOf) HasSocketMemoryId() bool`

HasSocketMemoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


