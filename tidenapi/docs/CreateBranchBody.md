# CreateBranchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedByAgentRunId** | Pointer to **string** | Set by the agent worker when an agent run is creating the branch on behalf of a user. Surfaced on the Branch message + UI banner. | [optional] 

## Methods

### NewCreateBranchBody

`func NewCreateBranchBody() *CreateBranchBody`

NewCreateBranchBody instantiates a new CreateBranchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBranchBodyWithDefaults

`func NewCreateBranchBodyWithDefaults() *CreateBranchBody`

NewCreateBranchBodyWithDefaults instantiates a new CreateBranchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBranchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBranchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBranchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBranchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateBranchBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBranchBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBranchBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBranchBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedByAgentRunId

`func (o *CreateBranchBody) GetCreatedByAgentRunId() string`

GetCreatedByAgentRunId returns the CreatedByAgentRunId field if non-nil, zero value otherwise.

### GetCreatedByAgentRunIdOk

`func (o *CreateBranchBody) GetCreatedByAgentRunIdOk() (*string, bool)`

GetCreatedByAgentRunIdOk returns a tuple with the CreatedByAgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentRunId

`func (o *CreateBranchBody) SetCreatedByAgentRunId(v string)`

SetCreatedByAgentRunId sets CreatedByAgentRunId field to given value.

### HasCreatedByAgentRunId

`func (o *CreateBranchBody) HasCreatedByAgentRunId() bool`

HasCreatedByAgentRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


