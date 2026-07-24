# MergeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additions** | Pointer to **int32** |  | [optional] 
**Modifications** | Pointer to **int32** |  | [optional] 
**Deletions** | Pointer to **int32** |  | [optional] 
**Conflicts** | Pointer to **int32** |  | [optional] 
**TestAdditions** | Pointer to **int32** |  | [optional] 
**TestModifications** | Pointer to **int32** |  | [optional] 
**TestDeletions** | Pointer to **int32** |  | [optional] 
**ComponentAdditions** | Pointer to **int32** |  | [optional] 
**ComponentModifications** | Pointer to **int32** |  | [optional] 
**ComponentDeletions** | Pointer to **int32** |  | [optional] 

## Methods

### NewMergeStats

`func NewMergeStats() *MergeStats`

NewMergeStats instantiates a new MergeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeStatsWithDefaults

`func NewMergeStatsWithDefaults() *MergeStats`

NewMergeStatsWithDefaults instantiates a new MergeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditions

`func (o *MergeStats) GetAdditions() int32`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *MergeStats) GetAdditionsOk() (*int32, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *MergeStats) SetAdditions(v int32)`

SetAdditions sets Additions field to given value.

### HasAdditions

`func (o *MergeStats) HasAdditions() bool`

HasAdditions returns a boolean if a field has been set.

### GetModifications

`func (o *MergeStats) GetModifications() int32`

GetModifications returns the Modifications field if non-nil, zero value otherwise.

### GetModificationsOk

`func (o *MergeStats) GetModificationsOk() (*int32, bool)`

GetModificationsOk returns a tuple with the Modifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifications

`func (o *MergeStats) SetModifications(v int32)`

SetModifications sets Modifications field to given value.

### HasModifications

`func (o *MergeStats) HasModifications() bool`

HasModifications returns a boolean if a field has been set.

### GetDeletions

`func (o *MergeStats) GetDeletions() int32`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *MergeStats) GetDeletionsOk() (*int32, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *MergeStats) SetDeletions(v int32)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *MergeStats) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetConflicts

`func (o *MergeStats) GetConflicts() int32`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *MergeStats) GetConflictsOk() (*int32, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *MergeStats) SetConflicts(v int32)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *MergeStats) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetTestAdditions

`func (o *MergeStats) GetTestAdditions() int32`

GetTestAdditions returns the TestAdditions field if non-nil, zero value otherwise.

### GetTestAdditionsOk

`func (o *MergeStats) GetTestAdditionsOk() (*int32, bool)`

GetTestAdditionsOk returns a tuple with the TestAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdditions

`func (o *MergeStats) SetTestAdditions(v int32)`

SetTestAdditions sets TestAdditions field to given value.

### HasTestAdditions

`func (o *MergeStats) HasTestAdditions() bool`

HasTestAdditions returns a boolean if a field has been set.

### GetTestModifications

`func (o *MergeStats) GetTestModifications() int32`

GetTestModifications returns the TestModifications field if non-nil, zero value otherwise.

### GetTestModificationsOk

`func (o *MergeStats) GetTestModificationsOk() (*int32, bool)`

GetTestModificationsOk returns a tuple with the TestModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModifications

`func (o *MergeStats) SetTestModifications(v int32)`

SetTestModifications sets TestModifications field to given value.

### HasTestModifications

`func (o *MergeStats) HasTestModifications() bool`

HasTestModifications returns a boolean if a field has been set.

### GetTestDeletions

`func (o *MergeStats) GetTestDeletions() int32`

GetTestDeletions returns the TestDeletions field if non-nil, zero value otherwise.

### GetTestDeletionsOk

`func (o *MergeStats) GetTestDeletionsOk() (*int32, bool)`

GetTestDeletionsOk returns a tuple with the TestDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDeletions

`func (o *MergeStats) SetTestDeletions(v int32)`

SetTestDeletions sets TestDeletions field to given value.

### HasTestDeletions

`func (o *MergeStats) HasTestDeletions() bool`

HasTestDeletions returns a boolean if a field has been set.

### GetComponentAdditions

`func (o *MergeStats) GetComponentAdditions() int32`

GetComponentAdditions returns the ComponentAdditions field if non-nil, zero value otherwise.

### GetComponentAdditionsOk

`func (o *MergeStats) GetComponentAdditionsOk() (*int32, bool)`

GetComponentAdditionsOk returns a tuple with the ComponentAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAdditions

`func (o *MergeStats) SetComponentAdditions(v int32)`

SetComponentAdditions sets ComponentAdditions field to given value.

### HasComponentAdditions

`func (o *MergeStats) HasComponentAdditions() bool`

HasComponentAdditions returns a boolean if a field has been set.

### GetComponentModifications

`func (o *MergeStats) GetComponentModifications() int32`

GetComponentModifications returns the ComponentModifications field if non-nil, zero value otherwise.

### GetComponentModificationsOk

`func (o *MergeStats) GetComponentModificationsOk() (*int32, bool)`

GetComponentModificationsOk returns a tuple with the ComponentModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentModifications

`func (o *MergeStats) SetComponentModifications(v int32)`

SetComponentModifications sets ComponentModifications field to given value.

### HasComponentModifications

`func (o *MergeStats) HasComponentModifications() bool`

HasComponentModifications returns a boolean if a field has been set.

### GetComponentDeletions

`func (o *MergeStats) GetComponentDeletions() int32`

GetComponentDeletions returns the ComponentDeletions field if non-nil, zero value otherwise.

### GetComponentDeletionsOk

`func (o *MergeStats) GetComponentDeletionsOk() (*int32, bool)`

GetComponentDeletionsOk returns a tuple with the ComponentDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDeletions

`func (o *MergeStats) SetComponentDeletions(v int32)`

SetComponentDeletions sets ComponentDeletions field to given value.

### HasComponentDeletions

`func (o *MergeStats) HasComponentDeletions() bool`

HasComponentDeletions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


