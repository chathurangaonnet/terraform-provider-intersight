# HyperflexVcenterConfigPolicyListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;hyperflex.VcenterConfigPolicy&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]HyperflexVcenterConfigPolicy**](hyperflex.VcenterConfigPolicy.md) | The array of &#39;hyperflex.VcenterConfigPolicy&#39; resources matching the request. | [optional] 

## Methods

### NewHyperflexVcenterConfigPolicyListAllOf

`func NewHyperflexVcenterConfigPolicyListAllOf() *HyperflexVcenterConfigPolicyListAllOf`

NewHyperflexVcenterConfigPolicyListAllOf instantiates a new HyperflexVcenterConfigPolicyListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVcenterConfigPolicyListAllOfWithDefaults

`func NewHyperflexVcenterConfigPolicyListAllOfWithDefaults() *HyperflexVcenterConfigPolicyListAllOf`

NewHyperflexVcenterConfigPolicyListAllOfWithDefaults instantiates a new HyperflexVcenterConfigPolicyListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HyperflexVcenterConfigPolicyListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HyperflexVcenterConfigPolicyListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HyperflexVcenterConfigPolicyListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HyperflexVcenterConfigPolicyListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *HyperflexVcenterConfigPolicyListAllOf) GetResults() []HyperflexVcenterConfigPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HyperflexVcenterConfigPolicyListAllOf) GetResultsOk() (*[]HyperflexVcenterConfigPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HyperflexVcenterConfigPolicyListAllOf) SetResults(v []HyperflexVcenterConfigPolicy)`

SetResults sets Results field to given value.

### HasResults

`func (o *HyperflexVcenterConfigPolicyListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *HyperflexVcenterConfigPolicyListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *HyperflexVcenterConfigPolicyListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


