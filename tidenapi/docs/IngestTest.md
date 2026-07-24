# IngestTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** |  | [optional] 
**SuitePath** | Pointer to [**[]IngestSuiteSegment**](IngestSuiteSegment.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Layer** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**RequirementSeqNums** | Pointer to **[]int32** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]TestParameter**](TestParameter.md) |  | [optional] 
**Execution** | Pointer to [**TestExecution**](TestExecution.md) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Relations** | Pointer to [**[]TestRelation**](TestRelation.md) |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 

## Methods

### NewIngestTest

`func NewIngestTest() *IngestTest`

NewIngestTest instantiates a new IngestTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestTestWithDefaults

`func NewIngestTestWithDefaults() *IngestTest`

NewIngestTestWithDefaults instantiates a new IngestTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *IngestTest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IngestTest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IngestTest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IngestTest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSuitePath

`func (o *IngestTest) GetSuitePath() []IngestSuiteSegment`

GetSuitePath returns the SuitePath field if non-nil, zero value otherwise.

### GetSuitePathOk

`func (o *IngestTest) GetSuitePathOk() (*[]IngestSuiteSegment, bool)`

GetSuitePathOk returns a tuple with the SuitePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitePath

`func (o *IngestTest) SetSuitePath(v []IngestSuiteSegment)`

SetSuitePath sets SuitePath field to given value.

### HasSuitePath

`func (o *IngestTest) HasSuitePath() bool`

HasSuitePath returns a boolean if a field has been set.

### GetTitle

`func (o *IngestTest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IngestTest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IngestTest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IngestTest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *IngestTest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestTest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestTest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestTest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLayer

`func (o *IngestTest) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *IngestTest) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *IngestTest) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *IngestTest) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetType

`func (o *IngestTest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestTest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestTest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IngestTest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPriority

`func (o *IngestTest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IngestTest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IngestTest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IngestTest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMuted

`func (o *IngestTest) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *IngestTest) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *IngestTest) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *IngestTest) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetTags

`func (o *IngestTest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IngestTest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IngestTest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IngestTest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSteps

`func (o *IngestTest) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *IngestTest) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *IngestTest) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *IngestTest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetCustomFields

`func (o *IngestTest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IngestTest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IngestTest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IngestTest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFilePath

`func (o *IngestTest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *IngestTest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *IngestTest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *IngestTest) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRequirementSeqNums

`func (o *IngestTest) GetRequirementSeqNums() []int32`

GetRequirementSeqNums returns the RequirementSeqNums field if non-nil, zero value otherwise.

### GetRequirementSeqNumsOk

`func (o *IngestTest) GetRequirementSeqNumsOk() (*[]int32, bool)`

GetRequirementSeqNumsOk returns a tuple with the RequirementSeqNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementSeqNums

`func (o *IngestTest) SetRequirementSeqNums(v []int32)`

SetRequirementSeqNums sets RequirementSeqNums field to given value.

### HasRequirementSeqNums

`func (o *IngestTest) HasRequirementSeqNums() bool`

HasRequirementSeqNums returns a boolean if a field has been set.

### GetSignature

`func (o *IngestTest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *IngestTest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *IngestTest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *IngestTest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTestopsId

`func (o *IngestTest) GetTestopsId() int32`

GetTestopsId returns the TestopsId field if non-nil, zero value otherwise.

### GetTestopsIdOk

`func (o *IngestTest) GetTestopsIdOk() (*int32, bool)`

GetTestopsIdOk returns a tuple with the TestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsId

`func (o *IngestTest) SetTestopsId(v int32)`

SetTestopsId sets TestopsId field to given value.

### HasTestopsId

`func (o *IngestTest) HasTestopsId() bool`

HasTestopsId returns a boolean if a field has been set.

### GetParameters

`func (o *IngestTest) GetParameters() []TestParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IngestTest) GetParametersOk() (*[]TestParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IngestTest) SetParameters(v []TestParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IngestTest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetExecution

`func (o *IngestTest) GetExecution() TestExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *IngestTest) GetExecutionOk() (*TestExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *IngestTest) SetExecution(v TestExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *IngestTest) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetAttachments

`func (o *IngestTest) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *IngestTest) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *IngestTest) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *IngestTest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRelations

`func (o *IngestTest) GetRelations() []TestRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *IngestTest) GetRelationsOk() (*[]TestRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *IngestTest) SetRelations(v []TestRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *IngestTest) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetIsAutomated

`func (o *IngestTest) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *IngestTest) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *IngestTest) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *IngestTest) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


