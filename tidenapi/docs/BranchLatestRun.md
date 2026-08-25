# BranchLatestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeqNum** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Passed** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 

## Methods

### NewBranchLatestRun

`func NewBranchLatestRun() *BranchLatestRun`

NewBranchLatestRun instantiates a new BranchLatestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchLatestRunWithDefaults

`func NewBranchLatestRunWithDefaults() *BranchLatestRun`

NewBranchLatestRunWithDefaults instantiates a new BranchLatestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeqNum

`func (o *BranchLatestRun) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *BranchLatestRun) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *BranchLatestRun) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *BranchLatestRun) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetStatus

`func (o *BranchLatestRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BranchLatestRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BranchLatestRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BranchLatestRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *BranchLatestRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BranchLatestRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BranchLatestRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BranchLatestRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *BranchLatestRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BranchLatestRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BranchLatestRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BranchLatestRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetTotal

`func (o *BranchLatestRun) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BranchLatestRun) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BranchLatestRun) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BranchLatestRun) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPassed

`func (o *BranchLatestRun) GetPassed() int32`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *BranchLatestRun) GetPassedOk() (*int32, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *BranchLatestRun) SetPassed(v int32)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *BranchLatestRun) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetFailed

`func (o *BranchLatestRun) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BranchLatestRun) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BranchLatestRun) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BranchLatestRun) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


