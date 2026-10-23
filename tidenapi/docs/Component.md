# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**BranchStatus** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** | Repository-aware component scope.  canonical repo id (e.g. \&quot;github.com/acme/backend\&quot;); unset &#x3D; unscoped/main repo | [optional] 
**ComponentPaths** | Pointer to **[]string** |  | [optional] 
**RepositoryAliases** | Pointer to **[]string** |  | [optional] 

## Methods

### NewComponent

`func NewComponent() *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Component) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Component) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Component) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Component) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Component) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Component) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Component) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Component) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *Component) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Component) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Component) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Component) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Component) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Component) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Component) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Component) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Component) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Component) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Component) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Component) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Component) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Component) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Component) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Component) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBranchId

`func (o *Component) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *Component) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *Component) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *Component) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetSourceId

`func (o *Component) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Component) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Component) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Component) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetBranchStatus

`func (o *Component) GetBranchStatus() string`

GetBranchStatus returns the BranchStatus field if non-nil, zero value otherwise.

### GetBranchStatusOk

`func (o *Component) GetBranchStatusOk() (*string, bool)`

GetBranchStatusOk returns a tuple with the BranchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchStatus

`func (o *Component) SetBranchStatus(v string)`

SetBranchStatus sets BranchStatus field to given value.

### HasBranchStatus

`func (o *Component) HasBranchStatus() bool`

HasBranchStatus returns a boolean if a field has been set.

### GetRepository

`func (o *Component) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Component) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Component) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Component) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetComponentPaths

`func (o *Component) GetComponentPaths() []string`

GetComponentPaths returns the ComponentPaths field if non-nil, zero value otherwise.

### GetComponentPathsOk

`func (o *Component) GetComponentPathsOk() (*[]string, bool)`

GetComponentPathsOk returns a tuple with the ComponentPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPaths

`func (o *Component) SetComponentPaths(v []string)`

SetComponentPaths sets ComponentPaths field to given value.

### HasComponentPaths

`func (o *Component) HasComponentPaths() bool`

HasComponentPaths returns a boolean if a field has been set.

### GetRepositoryAliases

`func (o *Component) GetRepositoryAliases() []string`

GetRepositoryAliases returns the RepositoryAliases field if non-nil, zero value otherwise.

### GetRepositoryAliasesOk

`func (o *Component) GetRepositoryAliasesOk() (*[]string, bool)`

GetRepositoryAliasesOk returns a tuple with the RepositoryAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryAliases

`func (o *Component) SetRepositoryAliases(v []string)`

SetRepositoryAliases sets RepositoryAliases field to given value.

### HasRepositoryAliases

`func (o *Component) HasRepositoryAliases() bool`

HasRepositoryAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


