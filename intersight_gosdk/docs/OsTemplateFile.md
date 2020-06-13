# OsTemplateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the OS Template File that user uploads for unattended installation. | [optional] 
**Placeholders** | Pointer to **[]string** |  | [optional] 
**TemplateContent** | Pointer to **string** | The content of the entire template file is stored as value. The content can either be a static file content or a template content. The template is expected to conform to the golang template syntax.  The placeholders, if any, would be populated and the values provided would be  used to populate this template. | [optional] 

## Methods

### NewOsTemplateFile

`func NewOsTemplateFile() *OsTemplateFile`

NewOsTemplateFile instantiates a new OsTemplateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsTemplateFileWithDefaults

`func NewOsTemplateFileWithDefaults() *OsTemplateFile`

NewOsTemplateFileWithDefaults instantiates a new OsTemplateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OsTemplateFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsTemplateFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsTemplateFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsTemplateFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlaceholders

`func (o *OsTemplateFile) GetPlaceholders() []string`

GetPlaceholders returns the Placeholders field if non-nil, zero value otherwise.

### GetPlaceholdersOk

`func (o *OsTemplateFile) GetPlaceholdersOk() (*[]string, bool)`

GetPlaceholdersOk returns a tuple with the Placeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholders

`func (o *OsTemplateFile) SetPlaceholders(v []string)`

SetPlaceholders sets Placeholders field to given value.

### HasPlaceholders

`func (o *OsTemplateFile) HasPlaceholders() bool`

HasPlaceholders returns a boolean if a field has been set.

### GetTemplateContent

`func (o *OsTemplateFile) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *OsTemplateFile) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *OsTemplateFile) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *OsTemplateFile) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


