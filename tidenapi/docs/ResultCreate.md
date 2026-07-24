# ResultCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**TestopsIds** | Pointer to **[]int32** |  | [optional] 
**Execution** | Pointer to [**ResultExecution**](ResultExecution.md) |  | [optional] 
**Fields** | Pointer to **map[string]string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Steps** | Pointer to [**[]ResultStep**](ResultStep.md) |  | [optional] 
**StepsType** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**ParamGroups** | Pointer to [**[]ParamGroup**](ParamGroup.md) |  | [optional] 
**SuitePath** | Pointer to [**[]SuiteSegment**](SuiteSegment.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Defect** | Pointer to **bool** |  | [optional] 

## Methods

### NewResultCreate

`func NewResultCreate() *ResultCreate`

NewResultCreate instantiates a new ResultCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCreateWithDefaults

`func NewResultCreateWithDefaults() *ResultCreate`

NewResultCreateWithDefaults instantiates a new ResultCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResultCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ResultCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResultCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResultCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResultCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSignature

`func (o *ResultCreate) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ResultCreate) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ResultCreate) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ResultCreate) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetExternalId

`func (o *ResultCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ResultCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ResultCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ResultCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTestopsIds

`func (o *ResultCreate) GetTestopsIds() []int32`

GetTestopsIds returns the TestopsIds field if non-nil, zero value otherwise.

### GetTestopsIdsOk

`func (o *ResultCreate) GetTestopsIdsOk() (*[]int32, bool)`

GetTestopsIdsOk returns a tuple with the TestopsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsIds

`func (o *ResultCreate) SetTestopsIds(v []int32)`

SetTestopsIds sets TestopsIds field to given value.

### HasTestopsIds

`func (o *ResultCreate) HasTestopsIds() bool`

HasTestopsIds returns a boolean if a field has been set.

### GetExecution

`func (o *ResultCreate) GetExecution() ResultExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ResultCreate) GetExecutionOk() (*ResultExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ResultCreate) SetExecution(v ResultExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *ResultCreate) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetFields

`func (o *ResultCreate) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ResultCreate) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ResultCreate) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ResultCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetAttachments

`func (o *ResultCreate) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ResultCreate) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ResultCreate) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ResultCreate) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSteps

`func (o *ResultCreate) GetSteps() []ResultStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ResultCreate) GetStepsOk() (*[]ResultStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ResultCreate) SetSteps(v []ResultStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ResultCreate) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStepsType

`func (o *ResultCreate) GetStepsType() string`

GetStepsType returns the StepsType field if non-nil, zero value otherwise.

### GetStepsTypeOk

`func (o *ResultCreate) GetStepsTypeOk() (*string, bool)`

GetStepsTypeOk returns a tuple with the StepsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsType

`func (o *ResultCreate) SetStepsType(v string)`

SetStepsType sets StepsType field to given value.

### HasStepsType

`func (o *ResultCreate) HasStepsType() bool`

HasStepsType returns a boolean if a field has been set.

### GetParams

`func (o *ResultCreate) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ResultCreate) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ResultCreate) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ResultCreate) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetParamGroups

`func (o *ResultCreate) GetParamGroups() []ParamGroup`

GetParamGroups returns the ParamGroups field if non-nil, zero value otherwise.

### GetParamGroupsOk

`func (o *ResultCreate) GetParamGroupsOk() (*[]ParamGroup, bool)`

GetParamGroupsOk returns a tuple with the ParamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamGroups

`func (o *ResultCreate) SetParamGroups(v []ParamGroup)`

SetParamGroups sets ParamGroups field to given value.

### HasParamGroups

`func (o *ResultCreate) HasParamGroups() bool`

HasParamGroups returns a boolean if a field has been set.

### GetSuitePath

`func (o *ResultCreate) GetSuitePath() []SuiteSegment`

GetSuitePath returns the SuitePath field if non-nil, zero value otherwise.

### GetSuitePathOk

`func (o *ResultCreate) GetSuitePathOk() (*[]SuiteSegment, bool)`

GetSuitePathOk returns a tuple with the SuitePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitePath

`func (o *ResultCreate) SetSuitePath(v []SuiteSegment)`

SetSuitePath sets SuitePath field to given value.

### HasSuitePath

`func (o *ResultCreate) HasSuitePath() bool`

HasSuitePath returns a boolean if a field has been set.

### GetMessage

`func (o *ResultCreate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResultCreate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResultCreate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResultCreate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDefect

`func (o *ResultCreate) GetDefect() bool`

GetDefect returns the Defect field if non-nil, zero value otherwise.

### GetDefectOk

`func (o *ResultCreate) GetDefectOk() (*bool, bool)`

GetDefectOk returns a tuple with the Defect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefect

`func (o *ResultCreate) SetDefect(v bool)`

SetDefect sets Defect field to given value.

### HasDefect

`func (o *ResultCreate) HasDefect() bool`

HasDefect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


