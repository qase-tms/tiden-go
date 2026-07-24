# ComponentMergeModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchVersion** | Pointer to [**Component**](Component.md) |  | [optional] 
**MainVersion** | Pointer to [**Component**](Component.md) |  | [optional] 
**HasConflict** | Pointer to **bool** |  | [optional] 
**ConflictingFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewComponentMergeModification

`func NewComponentMergeModification() *ComponentMergeModification`

NewComponentMergeModification instantiates a new ComponentMergeModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentMergeModificationWithDefaults

`func NewComponentMergeModificationWithDefaults() *ComponentMergeModification`

NewComponentMergeModificationWithDefaults instantiates a new ComponentMergeModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchVersion

`func (o *ComponentMergeModification) GetBranchVersion() Component`

GetBranchVersion returns the BranchVersion field if non-nil, zero value otherwise.

### GetBranchVersionOk

`func (o *ComponentMergeModification) GetBranchVersionOk() (*Component, bool)`

GetBranchVersionOk returns a tuple with the BranchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchVersion

`func (o *ComponentMergeModification) SetBranchVersion(v Component)`

SetBranchVersion sets BranchVersion field to given value.

### HasBranchVersion

`func (o *ComponentMergeModification) HasBranchVersion() bool`

HasBranchVersion returns a boolean if a field has been set.

### GetMainVersion

`func (o *ComponentMergeModification) GetMainVersion() Component`

GetMainVersion returns the MainVersion field if non-nil, zero value otherwise.

### GetMainVersionOk

`func (o *ComponentMergeModification) GetMainVersionOk() (*Component, bool)`

GetMainVersionOk returns a tuple with the MainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainVersion

`func (o *ComponentMergeModification) SetMainVersion(v Component)`

SetMainVersion sets MainVersion field to given value.

### HasMainVersion

`func (o *ComponentMergeModification) HasMainVersion() bool`

HasMainVersion returns a boolean if a field has been set.

### GetHasConflict

`func (o *ComponentMergeModification) GetHasConflict() bool`

GetHasConflict returns the HasConflict field if non-nil, zero value otherwise.

### GetHasConflictOk

`func (o *ComponentMergeModification) GetHasConflictOk() (*bool, bool)`

GetHasConflictOk returns a tuple with the HasConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConflict

`func (o *ComponentMergeModification) SetHasConflict(v bool)`

SetHasConflict sets HasConflict field to given value.

### HasHasConflict

`func (o *ComponentMergeModification) HasHasConflict() bool`

HasHasConflict returns a boolean if a field has been set.

### GetConflictingFields

`func (o *ComponentMergeModification) GetConflictingFields() []string`

GetConflictingFields returns the ConflictingFields field if non-nil, zero value otherwise.

### GetConflictingFieldsOk

`func (o *ComponentMergeModification) GetConflictingFieldsOk() (*[]string, bool)`

GetConflictingFieldsOk returns a tuple with the ConflictingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingFields

`func (o *ComponentMergeModification) SetConflictingFields(v []string)`

SetConflictingFields sets ConflictingFields field to given value.

### HasConflictingFields

`func (o *ComponentMergeModification) HasConflictingFields() bool`

HasConflictingFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


