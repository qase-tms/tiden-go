# TestRunResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**TestSeqNum** | Pointer to **int32** |  | [optional] 
**EventSeq** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IdentityKey** | Pointer to **string** |  | [optional] 
**ExecutionKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Thread** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Stacktrace** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**ParamGroups** | Pointer to [**[]ParamGroup**](ParamGroup.md) |  | [optional] 
**Fields** | Pointer to **map[string]string** |  | [optional] 
**Steps** | Pointer to [**[]ResultStep**](ResultStep.md) |  | [optional] 
**SuitePath** | Pointer to [**[]SuiteSegment**](SuiteSegment.md) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**Defect** | Pointer to **bool** |  | [optional] 
**IsLatestAttempt** | Pointer to **bool** |  | [optional] 
**Attempt** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTestRunResult

`func NewTestRunResult() *TestRunResult`

NewTestRunResult instantiates a new TestRunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunResultWithDefaults

`func NewTestRunResultWithDefaults() *TestRunResult`

NewTestRunResultWithDefaults instantiates a new TestRunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRunResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRunResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRunResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestRunResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunId

`func (o *TestRunResult) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *TestRunResult) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *TestRunResult) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *TestRunResult) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetTestId

`func (o *TestRunResult) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestRunResult) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestRunResult) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *TestRunResult) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSeqNum

`func (o *TestRunResult) GetTestSeqNum() int32`

GetTestSeqNum returns the TestSeqNum field if non-nil, zero value otherwise.

### GetTestSeqNumOk

`func (o *TestRunResult) GetTestSeqNumOk() (*int32, bool)`

GetTestSeqNumOk returns a tuple with the TestSeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSeqNum

`func (o *TestRunResult) SetTestSeqNum(v int32)`

SetTestSeqNum sets TestSeqNum field to given value.

### HasTestSeqNum

`func (o *TestRunResult) HasTestSeqNum() bool`

HasTestSeqNum returns a boolean if a field has been set.

### GetEventSeq

`func (o *TestRunResult) GetEventSeq() string`

GetEventSeq returns the EventSeq field if non-nil, zero value otherwise.

### GetEventSeqOk

`func (o *TestRunResult) GetEventSeqOk() (*string, bool)`

GetEventSeqOk returns a tuple with the EventSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSeq

`func (o *TestRunResult) SetEventSeq(v string)`

SetEventSeq sets EventSeq field to given value.

### HasEventSeq

`func (o *TestRunResult) HasEventSeq() bool`

HasEventSeq returns a boolean if a field has been set.

### GetTitle

`func (o *TestRunResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestRunResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestRunResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestRunResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSignature

`func (o *TestRunResult) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TestRunResult) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TestRunResult) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *TestRunResult) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetExternalId

`func (o *TestRunResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TestRunResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TestRunResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TestRunResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdentityKey

`func (o *TestRunResult) GetIdentityKey() string`

GetIdentityKey returns the IdentityKey field if non-nil, zero value otherwise.

### GetIdentityKeyOk

`func (o *TestRunResult) GetIdentityKeyOk() (*string, bool)`

GetIdentityKeyOk returns a tuple with the IdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityKey

`func (o *TestRunResult) SetIdentityKey(v string)`

SetIdentityKey sets IdentityKey field to given value.

### HasIdentityKey

`func (o *TestRunResult) HasIdentityKey() bool`

HasIdentityKey returns a boolean if a field has been set.

### GetExecutionKey

`func (o *TestRunResult) GetExecutionKey() string`

GetExecutionKey returns the ExecutionKey field if non-nil, zero value otherwise.

### GetExecutionKeyOk

`func (o *TestRunResult) GetExecutionKeyOk() (*string, bool)`

GetExecutionKeyOk returns a tuple with the ExecutionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionKey

`func (o *TestRunResult) SetExecutionKey(v string)`

SetExecutionKey sets ExecutionKey field to given value.

### HasExecutionKey

`func (o *TestRunResult) HasExecutionKey() bool`

HasExecutionKey returns a boolean if a field has been set.

### GetStatus

`func (o *TestRunResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRunResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRunResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestRunResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDurationMs

`func (o *TestRunResult) GetDurationMs() string`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TestRunResult) GetDurationMsOk() (*string, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TestRunResult) SetDurationMs(v string)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TestRunResult) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetStartedAt

`func (o *TestRunResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TestRunResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TestRunResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TestRunResult) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *TestRunResult) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *TestRunResult) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *TestRunResult) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *TestRunResult) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetThread

`func (o *TestRunResult) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *TestRunResult) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *TestRunResult) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *TestRunResult) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetMessage

`func (o *TestRunResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestRunResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestRunResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestRunResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStacktrace

`func (o *TestRunResult) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *TestRunResult) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *TestRunResult) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *TestRunResult) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### GetParams

`func (o *TestRunResult) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TestRunResult) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TestRunResult) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *TestRunResult) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetParamGroups

`func (o *TestRunResult) GetParamGroups() []ParamGroup`

GetParamGroups returns the ParamGroups field if non-nil, zero value otherwise.

### GetParamGroupsOk

`func (o *TestRunResult) GetParamGroupsOk() (*[]ParamGroup, bool)`

GetParamGroupsOk returns a tuple with the ParamGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamGroups

`func (o *TestRunResult) SetParamGroups(v []ParamGroup)`

SetParamGroups sets ParamGroups field to given value.

### HasParamGroups

`func (o *TestRunResult) HasParamGroups() bool`

HasParamGroups returns a boolean if a field has been set.

### GetFields

`func (o *TestRunResult) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TestRunResult) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TestRunResult) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TestRunResult) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSteps

`func (o *TestRunResult) GetSteps() []ResultStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *TestRunResult) GetStepsOk() (*[]ResultStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *TestRunResult) SetSteps(v []ResultStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *TestRunResult) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSuitePath

`func (o *TestRunResult) GetSuitePath() []SuiteSegment`

GetSuitePath returns the SuitePath field if non-nil, zero value otherwise.

### GetSuitePathOk

`func (o *TestRunResult) GetSuitePathOk() (*[]SuiteSegment, bool)`

GetSuitePathOk returns a tuple with the SuitePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitePath

`func (o *TestRunResult) SetSuitePath(v []SuiteSegment)`

SetSuitePath sets SuitePath field to given value.

### HasSuitePath

`func (o *TestRunResult) HasSuitePath() bool`

HasSuitePath returns a boolean if a field has been set.

### GetAttachments

`func (o *TestRunResult) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestRunResult) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestRunResult) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestRunResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetMuted

`func (o *TestRunResult) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *TestRunResult) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *TestRunResult) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *TestRunResult) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetDefect

`func (o *TestRunResult) GetDefect() bool`

GetDefect returns the Defect field if non-nil, zero value otherwise.

### GetDefectOk

`func (o *TestRunResult) GetDefectOk() (*bool, bool)`

GetDefectOk returns a tuple with the Defect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefect

`func (o *TestRunResult) SetDefect(v bool)`

SetDefect sets Defect field to given value.

### HasDefect

`func (o *TestRunResult) HasDefect() bool`

HasDefect returns a boolean if a field has been set.

### GetIsLatestAttempt

`func (o *TestRunResult) GetIsLatestAttempt() bool`

GetIsLatestAttempt returns the IsLatestAttempt field if non-nil, zero value otherwise.

### GetIsLatestAttemptOk

`func (o *TestRunResult) GetIsLatestAttemptOk() (*bool, bool)`

GetIsLatestAttemptOk returns a tuple with the IsLatestAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestAttempt

`func (o *TestRunResult) SetIsLatestAttempt(v bool)`

SetIsLatestAttempt sets IsLatestAttempt field to given value.

### HasIsLatestAttempt

`func (o *TestRunResult) HasIsLatestAttempt() bool`

HasIsLatestAttempt returns a boolean if a field has been set.

### GetAttempt

`func (o *TestRunResult) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *TestRunResult) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *TestRunResult) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *TestRunResult) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestRunResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestRunResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestRunResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestRunResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


