# CreateComponentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**ComponentPaths** | Pointer to **[]string** |  | [optional] 
**RepositoryAliases** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateComponentBody

`func NewCreateComponentBody() *CreateComponentBody`

NewCreateComponentBody instantiates a new CreateComponentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComponentBodyWithDefaults

`func NewCreateComponentBodyWithDefaults() *CreateComponentBody`

NewCreateComponentBodyWithDefaults instantiates a new CreateComponentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateComponentBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateComponentBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateComponentBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateComponentBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateComponentBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateComponentBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateComponentBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateComponentBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBranch

`func (o *CreateComponentBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CreateComponentBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CreateComponentBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CreateComponentBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRepository

`func (o *CreateComponentBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CreateComponentBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CreateComponentBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CreateComponentBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetComponentPaths

`func (o *CreateComponentBody) GetComponentPaths() []string`

GetComponentPaths returns the ComponentPaths field if non-nil, zero value otherwise.

### GetComponentPathsOk

`func (o *CreateComponentBody) GetComponentPathsOk() (*[]string, bool)`

GetComponentPathsOk returns a tuple with the ComponentPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPaths

`func (o *CreateComponentBody) SetComponentPaths(v []string)`

SetComponentPaths sets ComponentPaths field to given value.

### HasComponentPaths

`func (o *CreateComponentBody) HasComponentPaths() bool`

HasComponentPaths returns a boolean if a field has been set.

### GetRepositoryAliases

`func (o *CreateComponentBody) GetRepositoryAliases() []string`

GetRepositoryAliases returns the RepositoryAliases field if non-nil, zero value otherwise.

### GetRepositoryAliasesOk

`func (o *CreateComponentBody) GetRepositoryAliasesOk() (*[]string, bool)`

GetRepositoryAliasesOk returns a tuple with the RepositoryAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryAliases

`func (o *CreateComponentBody) SetRepositoryAliases(v []string)`

SetRepositoryAliases sets RepositoryAliases field to given value.

### HasRepositoryAliases

`func (o *CreateComponentBody) HasRepositoryAliases() bool`

HasRepositoryAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


