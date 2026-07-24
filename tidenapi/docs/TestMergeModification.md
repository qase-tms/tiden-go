# TestMergeModification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchVersion** | Pointer to [**Test**](Test.md) |  | [optional] 
**MainVersion** | Pointer to [**Test**](Test.md) |  | [optional] 
**HasConflict** | Pointer to **bool** |  | [optional] 
**ConflictingFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTestMergeModification

`func NewTestMergeModification() *TestMergeModification`

NewTestMergeModification instantiates a new TestMergeModification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestMergeModificationWithDefaults

`func NewTestMergeModificationWithDefaults() *TestMergeModification`

NewTestMergeModificationWithDefaults instantiates a new TestMergeModification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchVersion

`func (o *TestMergeModification) GetBranchVersion() Test`

GetBranchVersion returns the BranchVersion field if non-nil, zero value otherwise.

### GetBranchVersionOk

`func (o *TestMergeModification) GetBranchVersionOk() (*Test, bool)`

GetBranchVersionOk returns a tuple with the BranchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchVersion

`func (o *TestMergeModification) SetBranchVersion(v Test)`

SetBranchVersion sets BranchVersion field to given value.

### HasBranchVersion

`func (o *TestMergeModification) HasBranchVersion() bool`

HasBranchVersion returns a boolean if a field has been set.

### GetMainVersion

`func (o *TestMergeModification) GetMainVersion() Test`

GetMainVersion returns the MainVersion field if non-nil, zero value otherwise.

### GetMainVersionOk

`func (o *TestMergeModification) GetMainVersionOk() (*Test, bool)`

GetMainVersionOk returns a tuple with the MainVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainVersion

`func (o *TestMergeModification) SetMainVersion(v Test)`

SetMainVersion sets MainVersion field to given value.

### HasMainVersion

`func (o *TestMergeModification) HasMainVersion() bool`

HasMainVersion returns a boolean if a field has been set.

### GetHasConflict

`func (o *TestMergeModification) GetHasConflict() bool`

GetHasConflict returns the HasConflict field if non-nil, zero value otherwise.

### GetHasConflictOk

`func (o *TestMergeModification) GetHasConflictOk() (*bool, bool)`

GetHasConflictOk returns a tuple with the HasConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConflict

`func (o *TestMergeModification) SetHasConflict(v bool)`

SetHasConflict sets HasConflict field to given value.

### HasHasConflict

`func (o *TestMergeModification) HasHasConflict() bool`

HasHasConflict returns a boolean if a field has been set.

### GetConflictingFields

`func (o *TestMergeModification) GetConflictingFields() []string`

GetConflictingFields returns the ConflictingFields field if non-nil, zero value otherwise.

### GetConflictingFieldsOk

`func (o *TestMergeModification) GetConflictingFieldsOk() (*[]string, bool)`

GetConflictingFieldsOk returns a tuple with the ConflictingFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingFields

`func (o *TestMergeModification) SetConflictingFields(v []string)`

SetConflictingFields sets ConflictingFields field to given value.

### HasConflictingFields

`func (o *TestMergeModification) HasConflictingFields() bool`

HasConflictingFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


