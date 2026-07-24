# UpdateTestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Layer** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**TagsSet** | Pointer to **bool** | Replacement semantics: if a field is set on the request, it replaces the entire stored value. Use empty list/struct to clear. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFieldsSet** | Pointer to **bool** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**StepsSet** | Pointer to **bool** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **int32** |  | [optional] 
**ClearTestopsId** | Pointer to **bool** |  | [optional] 
**AuthorType** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**ParameterGroupsSet** | Pointer to **bool** |  | [optional] 
**ParameterGroups** | Pointer to [**[]TestParameterGroup**](TestParameterGroup.md) |  | [optional] 
**LatestExecutionSet** | Pointer to **bool** |  | [optional] 
**LatestExecution** | Pointer to [**TestExecution**](TestExecution.md) |  | [optional] 
**AttachmentsSet** | Pointer to **bool** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**RelationsSet** | Pointer to **bool** |  | [optional] 
**Relations** | Pointer to [**[]TestRelation**](TestRelation.md) |  | [optional] 

## Methods

### NewUpdateTestBody

`func NewUpdateTestBody() *UpdateTestBody`

NewUpdateTestBody instantiates a new UpdateTestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTestBodyWithDefaults

`func NewUpdateTestBodyWithDefaults() *UpdateTestBody`

NewUpdateTestBodyWithDefaults instantiates a new UpdateTestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *UpdateTestBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *UpdateTestBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *UpdateTestBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *UpdateTestBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateTestBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateTestBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateTestBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateTestBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *UpdateTestBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateTestBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateTestBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateTestBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateTestBody) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateTestBody) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateTestBody) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateTestBody) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateTestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateTestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateTestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateTestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateTestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateTestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateTestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateTestBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *UpdateTestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateTestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateTestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateTestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *UpdateTestBody) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *UpdateTestBody) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *UpdateTestBody) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *UpdateTestBody) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetMuted

`func (o *UpdateTestBody) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *UpdateTestBody) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *UpdateTestBody) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *UpdateTestBody) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetComponentId

`func (o *UpdateTestBody) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *UpdateTestBody) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *UpdateTestBody) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *UpdateTestBody) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetAssigneeId

`func (o *UpdateTestBody) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *UpdateTestBody) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *UpdateTestBody) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *UpdateTestBody) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetTagsSet

`func (o *UpdateTestBody) GetTagsSet() bool`

GetTagsSet returns the TagsSet field if non-nil, zero value otherwise.

### GetTagsSetOk

`func (o *UpdateTestBody) GetTagsSetOk() (*bool, bool)`

GetTagsSetOk returns a tuple with the TagsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsSet

`func (o *UpdateTestBody) SetTagsSet(v bool)`

SetTagsSet sets TagsSet field to given value.

### HasTagsSet

`func (o *UpdateTestBody) HasTagsSet() bool`

HasTagsSet returns a boolean if a field has been set.

### GetTags

`func (o *UpdateTestBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateTestBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateTestBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateTestBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFieldsSet

`func (o *UpdateTestBody) GetCustomFieldsSet() bool`

GetCustomFieldsSet returns the CustomFieldsSet field if non-nil, zero value otherwise.

### GetCustomFieldsSetOk

`func (o *UpdateTestBody) GetCustomFieldsSetOk() (*bool, bool)`

GetCustomFieldsSetOk returns a tuple with the CustomFieldsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsSet

`func (o *UpdateTestBody) SetCustomFieldsSet(v bool)`

SetCustomFieldsSet sets CustomFieldsSet field to given value.

### HasCustomFieldsSet

`func (o *UpdateTestBody) HasCustomFieldsSet() bool`

HasCustomFieldsSet returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateTestBody) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateTestBody) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateTestBody) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateTestBody) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStepsSet

`func (o *UpdateTestBody) GetStepsSet() bool`

GetStepsSet returns the StepsSet field if non-nil, zero value otherwise.

### GetStepsSetOk

`func (o *UpdateTestBody) GetStepsSetOk() (*bool, bool)`

GetStepsSetOk returns a tuple with the StepsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsSet

`func (o *UpdateTestBody) SetStepsSet(v bool)`

SetStepsSet sets StepsSet field to given value.

### HasStepsSet

`func (o *UpdateTestBody) HasStepsSet() bool`

HasStepsSet returns a boolean if a field has been set.

### GetSteps

`func (o *UpdateTestBody) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *UpdateTestBody) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *UpdateTestBody) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *UpdateTestBody) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetIsAutomated

`func (o *UpdateTestBody) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *UpdateTestBody) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *UpdateTestBody) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *UpdateTestBody) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetSignature

`func (o *UpdateTestBody) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UpdateTestBody) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UpdateTestBody) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UpdateTestBody) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTestopsId

`func (o *UpdateTestBody) GetTestopsId() int32`

GetTestopsId returns the TestopsId field if non-nil, zero value otherwise.

### GetTestopsIdOk

`func (o *UpdateTestBody) GetTestopsIdOk() (*int32, bool)`

GetTestopsIdOk returns a tuple with the TestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsId

`func (o *UpdateTestBody) SetTestopsId(v int32)`

SetTestopsId sets TestopsId field to given value.

### HasTestopsId

`func (o *UpdateTestBody) HasTestopsId() bool`

HasTestopsId returns a boolean if a field has been set.

### GetClearTestopsId

`func (o *UpdateTestBody) GetClearTestopsId() bool`

GetClearTestopsId returns the ClearTestopsId field if non-nil, zero value otherwise.

### GetClearTestopsIdOk

`func (o *UpdateTestBody) GetClearTestopsIdOk() (*bool, bool)`

GetClearTestopsIdOk returns a tuple with the ClearTestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearTestopsId

`func (o *UpdateTestBody) SetClearTestopsId(v bool)`

SetClearTestopsId sets ClearTestopsId field to given value.

### HasClearTestopsId

`func (o *UpdateTestBody) HasClearTestopsId() bool`

HasClearTestopsId returns a boolean if a field has been set.

### GetAuthorType

`func (o *UpdateTestBody) GetAuthorType() string`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *UpdateTestBody) GetAuthorTypeOk() (*string, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *UpdateTestBody) SetAuthorType(v string)`

SetAuthorType sets AuthorType field to given value.

### HasAuthorType

`func (o *UpdateTestBody) HasAuthorType() bool`

HasAuthorType returns a boolean if a field has been set.

### GetAuthorId

`func (o *UpdateTestBody) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *UpdateTestBody) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *UpdateTestBody) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *UpdateTestBody) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *UpdateTestBody) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *UpdateTestBody) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *UpdateTestBody) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *UpdateTestBody) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetParameterGroupsSet

`func (o *UpdateTestBody) GetParameterGroupsSet() bool`

GetParameterGroupsSet returns the ParameterGroupsSet field if non-nil, zero value otherwise.

### GetParameterGroupsSetOk

`func (o *UpdateTestBody) GetParameterGroupsSetOk() (*bool, bool)`

GetParameterGroupsSetOk returns a tuple with the ParameterGroupsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroupsSet

`func (o *UpdateTestBody) SetParameterGroupsSet(v bool)`

SetParameterGroupsSet sets ParameterGroupsSet field to given value.

### HasParameterGroupsSet

`func (o *UpdateTestBody) HasParameterGroupsSet() bool`

HasParameterGroupsSet returns a boolean if a field has been set.

### GetParameterGroups

`func (o *UpdateTestBody) GetParameterGroups() []TestParameterGroup`

GetParameterGroups returns the ParameterGroups field if non-nil, zero value otherwise.

### GetParameterGroupsOk

`func (o *UpdateTestBody) GetParameterGroupsOk() (*[]TestParameterGroup, bool)`

GetParameterGroupsOk returns a tuple with the ParameterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroups

`func (o *UpdateTestBody) SetParameterGroups(v []TestParameterGroup)`

SetParameterGroups sets ParameterGroups field to given value.

### HasParameterGroups

`func (o *UpdateTestBody) HasParameterGroups() bool`

HasParameterGroups returns a boolean if a field has been set.

### GetLatestExecutionSet

`func (o *UpdateTestBody) GetLatestExecutionSet() bool`

GetLatestExecutionSet returns the LatestExecutionSet field if non-nil, zero value otherwise.

### GetLatestExecutionSetOk

`func (o *UpdateTestBody) GetLatestExecutionSetOk() (*bool, bool)`

GetLatestExecutionSetOk returns a tuple with the LatestExecutionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecutionSet

`func (o *UpdateTestBody) SetLatestExecutionSet(v bool)`

SetLatestExecutionSet sets LatestExecutionSet field to given value.

### HasLatestExecutionSet

`func (o *UpdateTestBody) HasLatestExecutionSet() bool`

HasLatestExecutionSet returns a boolean if a field has been set.

### GetLatestExecution

`func (o *UpdateTestBody) GetLatestExecution() TestExecution`

GetLatestExecution returns the LatestExecution field if non-nil, zero value otherwise.

### GetLatestExecutionOk

`func (o *UpdateTestBody) GetLatestExecutionOk() (*TestExecution, bool)`

GetLatestExecutionOk returns a tuple with the LatestExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecution

`func (o *UpdateTestBody) SetLatestExecution(v TestExecution)`

SetLatestExecution sets LatestExecution field to given value.

### HasLatestExecution

`func (o *UpdateTestBody) HasLatestExecution() bool`

HasLatestExecution returns a boolean if a field has been set.

### GetAttachmentsSet

`func (o *UpdateTestBody) GetAttachmentsSet() bool`

GetAttachmentsSet returns the AttachmentsSet field if non-nil, zero value otherwise.

### GetAttachmentsSetOk

`func (o *UpdateTestBody) GetAttachmentsSetOk() (*bool, bool)`

GetAttachmentsSetOk returns a tuple with the AttachmentsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentsSet

`func (o *UpdateTestBody) SetAttachmentsSet(v bool)`

SetAttachmentsSet sets AttachmentsSet field to given value.

### HasAttachmentsSet

`func (o *UpdateTestBody) HasAttachmentsSet() bool`

HasAttachmentsSet returns a boolean if a field has been set.

### GetAttachments

`func (o *UpdateTestBody) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *UpdateTestBody) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *UpdateTestBody) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *UpdateTestBody) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRelationsSet

`func (o *UpdateTestBody) GetRelationsSet() bool`

GetRelationsSet returns the RelationsSet field if non-nil, zero value otherwise.

### GetRelationsSetOk

`func (o *UpdateTestBody) GetRelationsSetOk() (*bool, bool)`

GetRelationsSetOk returns a tuple with the RelationsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationsSet

`func (o *UpdateTestBody) SetRelationsSet(v bool)`

SetRelationsSet sets RelationsSet field to given value.

### HasRelationsSet

`func (o *UpdateTestBody) HasRelationsSet() bool`

HasRelationsSet returns a boolean if a field has been set.

### GetRelations

`func (o *UpdateTestBody) GetRelations() []TestRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *UpdateTestBody) GetRelationsOk() (*[]TestRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *UpdateTestBody) SetRelations(v []TestRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *UpdateTestBody) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


