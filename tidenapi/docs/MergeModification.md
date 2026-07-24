# MergeModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchVersion** | Pointer to [**Requirement**](Requirement.md) |  | [optional] 
**MainVersion** | Pointer to [**Requirement**](Requirement.md) |  | [optional] 
**HasConflict** | Pointer to **bool** |  | [optional] 
**ConflictingFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMergeModification

`func NewMergeModification() *MergeModification`

NewMergeModification instantiates a new MergeModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeModificationWithDefaults

`func NewMergeModificationWithDefaults() *MergeModification`

NewMergeModificationWithDefaults instantiates a new MergeModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchVersion

`func (o *MergeModification) GetBranchVersion() Requirement`

GetBranchVersion returns the BranchVersion field if non-nil, zero value otherwise.

### GetBranchVersionOk

`func (o *MergeModification) GetBranchVersionOk() (*Requirement, bool)`

GetBranchVersionOk returns a tuple with the BranchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchVersion

`func (o *MergeModification) SetBranchVersion(v Requirement)`

SetBranchVersion sets BranchVersion field to given value.

### HasBranchVersion

`func (o *MergeModification) HasBranchVersion() bool`

HasBranchVersion returns a boolean if a field has been set.

### GetMainVersion

`func (o *MergeModification) GetMainVersion() Requirement`

GetMainVersion returns the MainVersion field if non-nil, zero value otherwise.

### GetMainVersionOk

`func (o *MergeModification) GetMainVersionOk() (*Requirement, bool)`

GetMainVersionOk returns a tuple with the MainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainVersion

`func (o *MergeModification) SetMainVersion(v Requirement)`

SetMainVersion sets MainVersion field to given value.

### HasMainVersion

`func (o *MergeModification) HasMainVersion() bool`

HasMainVersion returns a boolean if a field has been set.

### GetHasConflict

`func (o *MergeModification) GetHasConflict() bool`

GetHasConflict returns the HasConflict field if non-nil, zero value otherwise.

### GetHasConflictOk

`func (o *MergeModification) GetHasConflictOk() (*bool, bool)`

GetHasConflictOk returns a tuple with the HasConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConflict

`func (o *MergeModification) SetHasConflict(v bool)`

SetHasConflict sets HasConflict field to given value.

### HasHasConflict

`func (o *MergeModification) HasHasConflict() bool`

HasHasConflict returns a boolean if a field has been set.

### GetConflictingFields

`func (o *MergeModification) GetConflictingFields() []string`

GetConflictingFields returns the ConflictingFields field if non-nil, zero value otherwise.

### GetConflictingFieldsOk

`func (o *MergeModification) GetConflictingFieldsOk() (*[]string, bool)`

GetConflictingFieldsOk returns a tuple with the ConflictingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingFields

`func (o *MergeModification) SetConflictingFields(v []string)`

SetConflictingFields sets ConflictingFields field to given value.

### HasConflictingFields

`func (o *MergeModification) HasConflictingFields() bool`

HasConflictingFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


