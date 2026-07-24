# RunCaseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityKey** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**SuitePath** | Pointer to **[]string** |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**TestSeqNum** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **string** |  | [optional] 
**Attempts** | Pointer to **int32** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**Combos** | Pointer to [**[]RunParamCombo**](RunParamCombo.md) |  | [optional] 

## Methods

### NewRunCaseSummary

`func NewRunCaseSummary() *RunCaseSummary`

NewRunCaseSummary instantiates a new RunCaseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCaseSummaryWithDefaults

`func NewRunCaseSummaryWithDefaults() *RunCaseSummary`

NewRunCaseSummaryWithDefaults instantiates a new RunCaseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityKey

`func (o *RunCaseSummary) GetIdentityKey() string`

GetIdentityKey returns the IdentityKey field if non-nil, zero value otherwise.

### GetIdentityKeyOk

`func (o *RunCaseSummary) GetIdentityKeyOk() (*string, bool)`

GetIdentityKeyOk returns a tuple with the IdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityKey

`func (o *RunCaseSummary) SetIdentityKey(v string)`

SetIdentityKey sets IdentityKey field to given value.

### HasIdentityKey

`func (o *RunCaseSummary) HasIdentityKey() bool`

HasIdentityKey returns a boolean if a field has been set.

### GetTitle

`func (o *RunCaseSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RunCaseSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RunCaseSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RunCaseSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSuitePath

`func (o *RunCaseSummary) GetSuitePath() []string`

GetSuitePath returns the SuitePath field if non-nil, zero value otherwise.

### GetSuitePathOk

`func (o *RunCaseSummary) GetSuitePathOk() (*[]string, bool)`

GetSuitePathOk returns a tuple with the SuitePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitePath

`func (o *RunCaseSummary) SetSuitePath(v []string)`

SetSuitePath sets SuitePath field to given value.

### HasSuitePath

`func (o *RunCaseSummary) HasSuitePath() bool`

HasSuitePath returns a boolean if a field has been set.

### GetTestId

`func (o *RunCaseSummary) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *RunCaseSummary) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *RunCaseSummary) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *RunCaseSummary) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSeqNum

`func (o *RunCaseSummary) GetTestSeqNum() int32`

GetTestSeqNum returns the TestSeqNum field if non-nil, zero value otherwise.

### GetTestSeqNumOk

`func (o *RunCaseSummary) GetTestSeqNumOk() (*int32, bool)`

GetTestSeqNumOk returns a tuple with the TestSeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSeqNum

`func (o *RunCaseSummary) SetTestSeqNum(v int32)`

SetTestSeqNum sets TestSeqNum field to given value.

### HasTestSeqNum

`func (o *RunCaseSummary) HasTestSeqNum() bool`

HasTestSeqNum returns a boolean if a field has been set.

### GetStatus

`func (o *RunCaseSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunCaseSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunCaseSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunCaseSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDurationMs

`func (o *RunCaseSummary) GetDurationMs() string`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *RunCaseSummary) GetDurationMsOk() (*string, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *RunCaseSummary) SetDurationMs(v string)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *RunCaseSummary) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetAttempts

`func (o *RunCaseSummary) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *RunCaseSummary) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *RunCaseSummary) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *RunCaseSummary) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetMuted

`func (o *RunCaseSummary) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *RunCaseSummary) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *RunCaseSummary) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *RunCaseSummary) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetCombos

`func (o *RunCaseSummary) GetCombos() []RunParamCombo`

GetCombos returns the Combos field if non-nil, zero value otherwise.

### GetCombosOk

`func (o *RunCaseSummary) GetCombosOk() (*[]RunParamCombo, bool)`

GetCombosOk returns a tuple with the Combos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombos

`func (o *RunCaseSummary) SetCombos(v []RunParamCombo)`

SetCombos sets Combos field to given value.

### HasCombos

`func (o *RunCaseSummary) HasCombos() bool`

HasCombos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


