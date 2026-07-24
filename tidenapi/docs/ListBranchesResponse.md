# ListBranchesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]Branch**](Branch.md) |  | [optional] 

## Methods

### NewListBranchesResponse

`func NewListBranchesResponse() *ListBranchesResponse`

NewListBranchesResponse instantiates a new ListBranchesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBranchesResponseWithDefaults

`func NewListBranchesResponseWithDefaults() *ListBranchesResponse`

NewListBranchesResponseWithDefaults instantiates a new ListBranchesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *ListBranchesResponse) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *ListBranchesResponse) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *ListBranchesResponse) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *ListBranchesResponse) HasBranches() bool`

HasBranches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


