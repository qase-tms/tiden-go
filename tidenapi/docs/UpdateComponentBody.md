# UpdateComponentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name/description use proto3 presence: an OMITTED field is left unchanged; a present field (even empty) is applied. Prevents a partial update that only sets repository scope from clobbering name/description to \&quot;\&quot;. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**RepositorySet** | Pointer to **bool** |  | [optional] 
**ComponentPaths** | Pointer to **[]string** |  | [optional] 
**ComponentPathsSet** | Pointer to **bool** |  | [optional] 
**RepositoryAliases** | Pointer to **[]string** |  | [optional] 
**RepositoryAliasesSet** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateComponentBody

`func NewUpdateComponentBody() *UpdateComponentBody`

NewUpdateComponentBody instantiates a new UpdateComponentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateComponentBodyWithDefaults

`func NewUpdateComponentBodyWithDefaults() *UpdateComponentBody`

NewUpdateComponentBodyWithDefaults instantiates a new UpdateComponentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateComponentBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateComponentBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateComponentBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateComponentBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateComponentBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateComponentBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateComponentBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateComponentBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBranch

`func (o *UpdateComponentBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *UpdateComponentBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *UpdateComponentBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *UpdateComponentBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRepository

`func (o *UpdateComponentBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *UpdateComponentBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *UpdateComponentBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *UpdateComponentBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositorySet

`func (o *UpdateComponentBody) GetRepositorySet() bool`

GetRepositorySet returns the RepositorySet field if non-nil, zero value otherwise.

### GetRepositorySetOk

`func (o *UpdateComponentBody) GetRepositorySetOk() (*bool, bool)`

GetRepositorySetOk returns a tuple with the RepositorySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySet

`func (o *UpdateComponentBody) SetRepositorySet(v bool)`

SetRepositorySet sets RepositorySet field to given value.

### HasRepositorySet

`func (o *UpdateComponentBody) HasRepositorySet() bool`

HasRepositorySet returns a boolean if a field has been set.

### GetComponentPaths

`func (o *UpdateComponentBody) GetComponentPaths() []string`

GetComponentPaths returns the ComponentPaths field if non-nil, zero value otherwise.

### GetComponentPathsOk

`func (o *UpdateComponentBody) GetComponentPathsOk() (*[]string, bool)`

GetComponentPathsOk returns a tuple with the ComponentPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPaths

`func (o *UpdateComponentBody) SetComponentPaths(v []string)`

SetComponentPaths sets ComponentPaths field to given value.

### HasComponentPaths

`func (o *UpdateComponentBody) HasComponentPaths() bool`

HasComponentPaths returns a boolean if a field has been set.

### GetComponentPathsSet

`func (o *UpdateComponentBody) GetComponentPathsSet() bool`

GetComponentPathsSet returns the ComponentPathsSet field if non-nil, zero value otherwise.

### GetComponentPathsSetOk

`func (o *UpdateComponentBody) GetComponentPathsSetOk() (*bool, bool)`

GetComponentPathsSetOk returns a tuple with the ComponentPathsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPathsSet

`func (o *UpdateComponentBody) SetComponentPathsSet(v bool)`

SetComponentPathsSet sets ComponentPathsSet field to given value.

### HasComponentPathsSet

`func (o *UpdateComponentBody) HasComponentPathsSet() bool`

HasComponentPathsSet returns a boolean if a field has been set.

### GetRepositoryAliases

`func (o *UpdateComponentBody) GetRepositoryAliases() []string`

GetRepositoryAliases returns the RepositoryAliases field if non-nil, zero value otherwise.

### GetRepositoryAliasesOk

`func (o *UpdateComponentBody) GetRepositoryAliasesOk() (*[]string, bool)`

GetRepositoryAliasesOk returns a tuple with the RepositoryAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryAliases

`func (o *UpdateComponentBody) SetRepositoryAliases(v []string)`

SetRepositoryAliases sets RepositoryAliases field to given value.

### HasRepositoryAliases

`func (o *UpdateComponentBody) HasRepositoryAliases() bool`

HasRepositoryAliases returns a boolean if a field has been set.

### GetRepositoryAliasesSet

`func (o *UpdateComponentBody) GetRepositoryAliasesSet() bool`

GetRepositoryAliasesSet returns the RepositoryAliasesSet field if non-nil, zero value otherwise.

### GetRepositoryAliasesSetOk

`func (o *UpdateComponentBody) GetRepositoryAliasesSetOk() (*bool, bool)`

GetRepositoryAliasesSetOk returns a tuple with the RepositoryAliasesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryAliasesSet

`func (o *UpdateComponentBody) SetRepositoryAliasesSet(v bool)`

SetRepositoryAliasesSet sets RepositoryAliasesSet field to given value.

### HasRepositoryAliasesSet

`func (o *UpdateComponentBody) HasRepositoryAliasesSet() bool`

HasRepositoryAliasesSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


