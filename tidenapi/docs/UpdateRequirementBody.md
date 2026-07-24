# UpdateRequirementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourcesUpdate** | Pointer to [**RequirementSourcesUpdate**](RequirementSourcesUpdate.md) |  | [optional] 

## Methods

### NewUpdateRequirementBody

`func NewUpdateRequirementBody() *UpdateRequirementBody`

NewUpdateRequirementBody instantiates a new UpdateRequirementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequirementBodyWithDefaults

`func NewUpdateRequirementBodyWithDefaults() *UpdateRequirementBody`

NewUpdateRequirementBodyWithDefaults instantiates a new UpdateRequirementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateRequirementBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateRequirementBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateRequirementBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateRequirementBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *UpdateRequirementBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateRequirementBody) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateRequirementBody) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateRequirementBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetParentId

`func (o *UpdateRequirementBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateRequirementBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateRequirementBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateRequirementBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetComponentId

`func (o *UpdateRequirementBody) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *UpdateRequirementBody) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *UpdateRequirementBody) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *UpdateRequirementBody) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateRequirementBody) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateRequirementBody) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateRequirementBody) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateRequirementBody) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateRequirementBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateRequirementBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateRequirementBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateRequirementBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateRequirementBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateRequirementBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateRequirementBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateRequirementBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssigneeId

`func (o *UpdateRequirementBody) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *UpdateRequirementBody) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *UpdateRequirementBody) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *UpdateRequirementBody) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetBranch

`func (o *UpdateRequirementBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *UpdateRequirementBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *UpdateRequirementBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *UpdateRequirementBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetType

`func (o *UpdateRequirementBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateRequirementBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateRequirementBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateRequirementBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourcesUpdate

`func (o *UpdateRequirementBody) GetSourcesUpdate() RequirementSourcesUpdate`

GetSourcesUpdate returns the SourcesUpdate field if non-nil, zero value otherwise.

### GetSourcesUpdateOk

`func (o *UpdateRequirementBody) GetSourcesUpdateOk() (*RequirementSourcesUpdate, bool)`

GetSourcesUpdateOk returns a tuple with the SourcesUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesUpdate

`func (o *UpdateRequirementBody) SetSourcesUpdate(v RequirementSourcesUpdate)`

SetSourcesUpdate sets SourcesUpdate field to given value.

### HasSourcesUpdate

`func (o *UpdateRequirementBody) HasSourcesUpdate() bool`

HasSourcesUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


