# MergeBranchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**Branch**](Branch.md) |  | [optional] 
**Stats** | Pointer to [**MergeStats**](MergeStats.md) |  | [optional] 

## Methods

### NewMergeBranchResponse

`func NewMergeBranchResponse() *MergeBranchResponse`

NewMergeBranchResponse instantiates a new MergeBranchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeBranchResponseWithDefaults

`func NewMergeBranchResponseWithDefaults() *MergeBranchResponse`

NewMergeBranchResponseWithDefaults instantiates a new MergeBranchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *MergeBranchResponse) GetBranch() Branch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *MergeBranchResponse) GetBranchOk() (*Branch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *MergeBranchResponse) SetBranch(v Branch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *MergeBranchResponse) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetStats

`func (o *MergeBranchResponse) GetStats() MergeStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *MergeBranchResponse) GetStatsOk() (*MergeStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *MergeBranchResponse) SetStats(v MergeStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *MergeBranchResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


