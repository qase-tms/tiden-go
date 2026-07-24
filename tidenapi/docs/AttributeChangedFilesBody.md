# AttributeChangedFilesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | branch the requirement lives on (empty &#x3D; main). | [optional] 
**ChangedFiles** | Pointer to [**[]ChangedFile**](ChangedFile.md) |  | [optional] 

## Methods

### NewAttributeChangedFilesBody

`func NewAttributeChangedFilesBody() *AttributeChangedFilesBody`

NewAttributeChangedFilesBody instantiates a new AttributeChangedFilesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeChangedFilesBodyWithDefaults

`func NewAttributeChangedFilesBodyWithDefaults() *AttributeChangedFilesBody`

NewAttributeChangedFilesBodyWithDefaults instantiates a new AttributeChangedFilesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *AttributeChangedFilesBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AttributeChangedFilesBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AttributeChangedFilesBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AttributeChangedFilesBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetChangedFiles

`func (o *AttributeChangedFilesBody) GetChangedFiles() []ChangedFile`

GetChangedFiles returns the ChangedFiles field if non-nil, zero value otherwise.

### GetChangedFilesOk

`func (o *AttributeChangedFilesBody) GetChangedFilesOk() (*[]ChangedFile, bool)`

GetChangedFilesOk returns a tuple with the ChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedFiles

`func (o *AttributeChangedFilesBody) SetChangedFiles(v []ChangedFile)`

SetChangedFiles sets ChangedFiles field to given value.

### HasChangedFiles

`func (o *AttributeChangedFilesBody) HasChangedFiles() bool`

HasChangedFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


