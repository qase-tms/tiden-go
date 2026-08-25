# BranchLoopStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerdictCycles** | Pointer to **int32** |  | [optional] 
**BlockedBeforePass** | Pointer to **int32** |  | [optional] 
**WentGreen** | Pointer to **bool** |  | [optional] 
**RunCount** | Pointer to **int32** |  | [optional] 
**LoopCycles** | Pointer to **int32** | loop_cycles is the requirement-coverage-rung judgment-cycle count (branches.loop_cycle_count) — how many times a branch-scope compute&#39;s per-requirement coverage rung actually moved, distinct from verdict_cycles above (a count of gate STATE CHANGES, which rises on its own once the gate goes per-stage). Absent when the branch was never counted (loop_cycle_progress_key IS NULL: pre-cutover, or no branch-scope verdict yet) — never the same as an explicit zero. | [optional] 

## Methods

### NewBranchLoopStats

`func NewBranchLoopStats() *BranchLoopStats`

NewBranchLoopStats instantiates a new BranchLoopStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchLoopStatsWithDefaults

`func NewBranchLoopStatsWithDefaults() *BranchLoopStats`

NewBranchLoopStatsWithDefaults instantiates a new BranchLoopStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdictCycles

`func (o *BranchLoopStats) GetVerdictCycles() int32`

GetVerdictCycles returns the VerdictCycles field if non-nil, zero value otherwise.

### GetVerdictCyclesOk

`func (o *BranchLoopStats) GetVerdictCyclesOk() (*int32, bool)`

GetVerdictCyclesOk returns a tuple with the VerdictCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdictCycles

`func (o *BranchLoopStats) SetVerdictCycles(v int32)`

SetVerdictCycles sets VerdictCycles field to given value.

### HasVerdictCycles

`func (o *BranchLoopStats) HasVerdictCycles() bool`

HasVerdictCycles returns a boolean if a field has been set.

### GetBlockedBeforePass

`func (o *BranchLoopStats) GetBlockedBeforePass() int32`

GetBlockedBeforePass returns the BlockedBeforePass field if non-nil, zero value otherwise.

### GetBlockedBeforePassOk

`func (o *BranchLoopStats) GetBlockedBeforePassOk() (*int32, bool)`

GetBlockedBeforePassOk returns a tuple with the BlockedBeforePass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBeforePass

`func (o *BranchLoopStats) SetBlockedBeforePass(v int32)`

SetBlockedBeforePass sets BlockedBeforePass field to given value.

### HasBlockedBeforePass

`func (o *BranchLoopStats) HasBlockedBeforePass() bool`

HasBlockedBeforePass returns a boolean if a field has been set.

### GetWentGreen

`func (o *BranchLoopStats) GetWentGreen() bool`

GetWentGreen returns the WentGreen field if non-nil, zero value otherwise.

### GetWentGreenOk

`func (o *BranchLoopStats) GetWentGreenOk() (*bool, bool)`

GetWentGreenOk returns a tuple with the WentGreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWentGreen

`func (o *BranchLoopStats) SetWentGreen(v bool)`

SetWentGreen sets WentGreen field to given value.

### HasWentGreen

`func (o *BranchLoopStats) HasWentGreen() bool`

HasWentGreen returns a boolean if a field has been set.

### GetRunCount

`func (o *BranchLoopStats) GetRunCount() int32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *BranchLoopStats) GetRunCountOk() (*int32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *BranchLoopStats) SetRunCount(v int32)`

SetRunCount sets RunCount field to given value.

### HasRunCount

`func (o *BranchLoopStats) HasRunCount() bool`

HasRunCount returns a boolean if a field has been set.

### GetLoopCycles

`func (o *BranchLoopStats) GetLoopCycles() int32`

GetLoopCycles returns the LoopCycles field if non-nil, zero value otherwise.

### GetLoopCyclesOk

`func (o *BranchLoopStats) GetLoopCyclesOk() (*int32, bool)`

GetLoopCyclesOk returns a tuple with the LoopCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopCycles

`func (o *BranchLoopStats) SetLoopCycles(v int32)`

SetLoopCycles sets LoopCycles field to given value.

### HasLoopCycles

`func (o *BranchLoopStats) HasLoopCycles() bool`

HasLoopCycles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


