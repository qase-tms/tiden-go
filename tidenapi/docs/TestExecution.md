# TestExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Thread** | Pointer to **string** |  | [optional] 
**RunSeq** | Pointer to **int32** | Test-run seq this execution came from (live-doc sync stamps it; 0 &#x3D; unknown — run seqs start at 1). Read-only: server-populated, never accepted from clients. | [optional] 

## Methods

### NewTestExecution

`func NewTestExecution() *TestExecution`

NewTestExecution instantiates a new TestExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestExecutionWithDefaults

`func NewTestExecutionWithDefaults() *TestExecution`

NewTestExecutionWithDefaults instantiates a new TestExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *TestExecution) GetDurationMs() string`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TestExecution) GetDurationMsOk() (*string, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TestExecution) SetDurationMs(v string)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TestExecution) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetStatus

`func (o *TestExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *TestExecution) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TestExecution) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TestExecution) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TestExecution) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TestExecution) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TestExecution) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TestExecution) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TestExecution) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetThread

`func (o *TestExecution) GetThread() string`

GetThread returns the Thread field if non-nil, zero value otherwise.

### GetThreadOk

`func (o *TestExecution) GetThreadOk() (*string, bool)`

GetThreadOk returns a tuple with the Thread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThread

`func (o *TestExecution) SetThread(v string)`

SetThread sets Thread field to given value.

### HasThread

`func (o *TestExecution) HasThread() bool`

HasThread returns a boolean if a field has been set.

### GetRunSeq

`func (o *TestExecution) GetRunSeq() int32`

GetRunSeq returns the RunSeq field if non-nil, zero value otherwise.

### GetRunSeqOk

`func (o *TestExecution) GetRunSeqOk() (*int32, bool)`

GetRunSeqOk returns a tuple with the RunSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunSeq

`func (o *TestExecution) SetRunSeq(v int32)`

SetRunSeq sets RunSeq field to given value.

### HasRunSeq

`func (o *TestExecution) HasRunSeq() bool`

HasRunSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


