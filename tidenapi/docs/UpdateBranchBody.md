# UpdateBranchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Both fields are optional with absent-means-unchanged semantics: an unset field leaves the branch&#39;s current value alone; a present field (including an explicit empty string) overwrites it. There is no separate \&quot;clear\&quot; flag. | [optional] 
**CreatedByAgent** | Pointer to **string** | Validated server-side against the same fixed allowlist as CreateBranchRequest.created_by_agent. | [optional] 

## Methods

### NewUpdateBranchBody

`func NewUpdateBranchBody() *UpdateBranchBody`

NewUpdateBranchBody instantiates a new UpdateBranchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBranchBodyWithDefaults

`func NewUpdateBranchBodyWithDefaults() *UpdateBranchBody`

NewUpdateBranchBodyWithDefaults instantiates a new UpdateBranchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateBranchBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBranchBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBranchBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBranchBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedByAgent

`func (o *UpdateBranchBody) GetCreatedByAgent() string`

GetCreatedByAgent returns the CreatedByAgent field if non-nil, zero value otherwise.

### GetCreatedByAgentOk

`func (o *UpdateBranchBody) GetCreatedByAgentOk() (*string, bool)`

GetCreatedByAgentOk returns a tuple with the CreatedByAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgent

`func (o *UpdateBranchBody) SetCreatedByAgent(v string)`

SetCreatedByAgent sets CreatedByAgent field to given value.

### HasCreatedByAgent

`func (o *UpdateBranchBody) HasCreatedByAgent() bool`

HasCreatedByAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


