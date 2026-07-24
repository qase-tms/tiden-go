# CreateRequirementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]RequirementSourceInput**](RequirementSourceInput.md) |  | [optional] 

## Methods

### NewCreateRequirementBody

`func NewCreateRequirementBody() *CreateRequirementBody`

NewCreateRequirementBody instantiates a new CreateRequirementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequirementBodyWithDefaults

`func NewCreateRequirementBodyWithDefaults() *CreateRequirementBody`

NewCreateRequirementBodyWithDefaults instantiates a new CreateRequirementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *CreateRequirementBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateRequirementBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateRequirementBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateRequirementBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetComponentId

`func (o *CreateRequirementBody) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *CreateRequirementBody) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *CreateRequirementBody) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *CreateRequirementBody) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateRequirementBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateRequirementBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateRequirementBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateRequirementBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *CreateRequirementBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateRequirementBody) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateRequirementBody) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateRequirementBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatus

`func (o *CreateRequirementBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRequirementBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRequirementBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRequirementBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *CreateRequirementBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateRequirementBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateRequirementBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateRequirementBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssigneeId

`func (o *CreateRequirementBody) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CreateRequirementBody) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CreateRequirementBody) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *CreateRequirementBody) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetBranch

`func (o *CreateRequirementBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CreateRequirementBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CreateRequirementBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CreateRequirementBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetType

`func (o *CreateRequirementBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRequirementBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRequirementBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRequirementBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSources

`func (o *CreateRequirementBody) GetSources() []RequirementSourceInput`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateRequirementBody) GetSourcesOk() (*[]RequirementSourceInput, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateRequirementBody) SetSources(v []RequirementSourceInput)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CreateRequirementBody) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


