# CreateTestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Layer** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **int32** |  | [optional] 
**AuthorType** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**ParameterGroups** | Pointer to [**[]TestParameterGroup**](TestParameterGroup.md) |  | [optional] 
**LatestExecution** | Pointer to [**TestExecution**](TestExecution.md) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Relations** | Pointer to [**[]TestRelation**](TestRelation.md) |  | [optional] 

## Methods

### NewCreateTestBody

`func NewCreateTestBody() *CreateTestBody`

NewCreateTestBody instantiates a new CreateTestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestBodyWithDefaults

`func NewCreateTestBodyWithDefaults() *CreateTestBody`

NewCreateTestBodyWithDefaults instantiates a new CreateTestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateTestBody) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateTestBody) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateTestBody) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateTestBody) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTitle

`func (o *CreateTestBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTestBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTestBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateTestBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *CreateTestBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateTestBody) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateTestBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateTestBody) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetBranch

`func (o *CreateTestBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CreateTestBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CreateTestBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CreateTestBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *CreateTestBody) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateTestBody) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateTestBody) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateTestBody) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *CreateTestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *CreateTestBody) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *CreateTestBody) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *CreateTestBody) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *CreateTestBody) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetMuted

`func (o *CreateTestBody) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *CreateTestBody) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *CreateTestBody) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *CreateTestBody) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetComponentId

`func (o *CreateTestBody) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *CreateTestBody) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *CreateTestBody) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *CreateTestBody) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetAssigneeId

`func (o *CreateTestBody) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CreateTestBody) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CreateTestBody) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *CreateTestBody) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetTags

`func (o *CreateTestBody) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTestBody) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTestBody) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTestBody) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateTestBody) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateTestBody) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateTestBody) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateTestBody) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSteps

`func (o *CreateTestBody) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateTestBody) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateTestBody) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CreateTestBody) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetFramework

`func (o *CreateTestBody) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CreateTestBody) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CreateTestBody) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CreateTestBody) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetFilePath

`func (o *CreateTestBody) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CreateTestBody) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CreateTestBody) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *CreateTestBody) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetIsAutomated

`func (o *CreateTestBody) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *CreateTestBody) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *CreateTestBody) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *CreateTestBody) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetSignature

`func (o *CreateTestBody) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CreateTestBody) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CreateTestBody) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CreateTestBody) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTestopsId

`func (o *CreateTestBody) GetTestopsId() int32`

GetTestopsId returns the TestopsId field if non-nil, zero value otherwise.

### GetTestopsIdOk

`func (o *CreateTestBody) GetTestopsIdOk() (*int32, bool)`

GetTestopsIdOk returns a tuple with the TestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsId

`func (o *CreateTestBody) SetTestopsId(v int32)`

SetTestopsId sets TestopsId field to given value.

### HasTestopsId

`func (o *CreateTestBody) HasTestopsId() bool`

HasTestopsId returns a boolean if a field has been set.

### GetAuthorType

`func (o *CreateTestBody) GetAuthorType() string`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *CreateTestBody) GetAuthorTypeOk() (*string, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *CreateTestBody) SetAuthorType(v string)`

SetAuthorType sets AuthorType field to given value.

### HasAuthorType

`func (o *CreateTestBody) HasAuthorType() bool`

HasAuthorType returns a boolean if a field has been set.

### GetAuthorId

`func (o *CreateTestBody) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateTestBody) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateTestBody) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateTestBody) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *CreateTestBody) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *CreateTestBody) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *CreateTestBody) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *CreateTestBody) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetParameterGroups

`func (o *CreateTestBody) GetParameterGroups() []TestParameterGroup`

GetParameterGroups returns the ParameterGroups field if non-nil, zero value otherwise.

### GetParameterGroupsOk

`func (o *CreateTestBody) GetParameterGroupsOk() (*[]TestParameterGroup, bool)`

GetParameterGroupsOk returns a tuple with the ParameterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroups

`func (o *CreateTestBody) SetParameterGroups(v []TestParameterGroup)`

SetParameterGroups sets ParameterGroups field to given value.

### HasParameterGroups

`func (o *CreateTestBody) HasParameterGroups() bool`

HasParameterGroups returns a boolean if a field has been set.

### GetLatestExecution

`func (o *CreateTestBody) GetLatestExecution() TestExecution`

GetLatestExecution returns the LatestExecution field if non-nil, zero value otherwise.

### GetLatestExecutionOk

`func (o *CreateTestBody) GetLatestExecutionOk() (*TestExecution, bool)`

GetLatestExecutionOk returns a tuple with the LatestExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecution

`func (o *CreateTestBody) SetLatestExecution(v TestExecution)`

SetLatestExecution sets LatestExecution field to given value.

### HasLatestExecution

`func (o *CreateTestBody) HasLatestExecution() bool`

HasLatestExecution returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateTestBody) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateTestBody) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateTestBody) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateTestBody) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRelations

`func (o *CreateTestBody) GetRelations() []TestRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *CreateTestBody) GetRelationsOk() (*[]TestRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *CreateTestBody) SetRelations(v []TestRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *CreateTestBody) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


