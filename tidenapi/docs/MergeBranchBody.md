# MergeBranchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolutions** | Pointer to **map[string]string** | Resolution map keys are prefixed: \&quot;req:&lt;uuid&gt;\&quot;, \&quot;test:&lt;uuid&gt;\&quot;, or \&quot;comp:&lt;uuid&gt;\&quot;. Server accepts un-prefixed keys as &#x60;req:&#x60; for v1 backwards compat (deprecated; logged with warning, removal scheduled for v2). | [optional] 

## Methods

### NewMergeBranchBody

`func NewMergeBranchBody() *MergeBranchBody`

NewMergeBranchBody instantiates a new MergeBranchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeBranchBodyWithDefaults

`func NewMergeBranchBodyWithDefaults() *MergeBranchBody`

NewMergeBranchBodyWithDefaults instantiates a new MergeBranchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolutions

`func (o *MergeBranchBody) GetResolutions() map[string]string`

GetResolutions returns the Resolutions field if non-nil, zero value otherwise.

### GetResolutionsOk

`func (o *MergeBranchBody) GetResolutionsOk() (*map[string]string, bool)`

GetResolutionsOk returns a tuple with the Resolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutions

`func (o *MergeBranchBody) SetResolutions(v map[string]string)`

SetResolutions sets Resolutions field to given value.

### HasResolutions

`func (o *MergeBranchBody) HasResolutions() bool`

HasResolutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


