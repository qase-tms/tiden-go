# BranchChangeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementAdditions** | Pointer to **int32** |  | [optional] 
**RequirementModifications** | Pointer to **int32** |  | [optional] 
**RequirementDeletions** | Pointer to **int32** |  | [optional] 
**TestAdditions** | Pointer to **int32** |  | [optional] 
**TestModifications** | Pointer to **int32** |  | [optional] 
**TestDeletions** | Pointer to **int32** |  | [optional] 
**ComponentAdditions** | Pointer to **int32** |  | [optional] 
**ComponentModifications** | Pointer to **int32** |  | [optional] 
**ComponentDeletions** | Pointer to **int32** |  | [optional] 
**Conflicts** | Pointer to **int32** |  | [optional] 

## Methods

### NewBranchChangeStats

`func NewBranchChangeStats() *BranchChangeStats`

NewBranchChangeStats instantiates a new BranchChangeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchChangeStatsWithDefaults

`func NewBranchChangeStatsWithDefaults() *BranchChangeStats`

NewBranchChangeStatsWithDefaults instantiates a new BranchChangeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementAdditions

`func (o *BranchChangeStats) GetRequirementAdditions() int32`

GetRequirementAdditions returns the RequirementAdditions field if non-nil, zero value otherwise.

### GetRequirementAdditionsOk

`func (o *BranchChangeStats) GetRequirementAdditionsOk() (*int32, bool)`

GetRequirementAdditionsOk returns a tuple with the RequirementAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementAdditions

`func (o *BranchChangeStats) SetRequirementAdditions(v int32)`

SetRequirementAdditions sets RequirementAdditions field to given value.

### HasRequirementAdditions

`func (o *BranchChangeStats) HasRequirementAdditions() bool`

HasRequirementAdditions returns a boolean if a field has been set.

### GetRequirementModifications

`func (o *BranchChangeStats) GetRequirementModifications() int32`

GetRequirementModifications returns the RequirementModifications field if non-nil, zero value otherwise.

### GetRequirementModificationsOk

`func (o *BranchChangeStats) GetRequirementModificationsOk() (*int32, bool)`

GetRequirementModificationsOk returns a tuple with the RequirementModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementModifications

`func (o *BranchChangeStats) SetRequirementModifications(v int32)`

SetRequirementModifications sets RequirementModifications field to given value.

### HasRequirementModifications

`func (o *BranchChangeStats) HasRequirementModifications() bool`

HasRequirementModifications returns a boolean if a field has been set.

### GetRequirementDeletions

`func (o *BranchChangeStats) GetRequirementDeletions() int32`

GetRequirementDeletions returns the RequirementDeletions field if non-nil, zero value otherwise.

### GetRequirementDeletionsOk

`func (o *BranchChangeStats) GetRequirementDeletionsOk() (*int32, bool)`

GetRequirementDeletionsOk returns a tuple with the RequirementDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementDeletions

`func (o *BranchChangeStats) SetRequirementDeletions(v int32)`

SetRequirementDeletions sets RequirementDeletions field to given value.

### HasRequirementDeletions

`func (o *BranchChangeStats) HasRequirementDeletions() bool`

HasRequirementDeletions returns a boolean if a field has been set.

### GetTestAdditions

`func (o *BranchChangeStats) GetTestAdditions() int32`

GetTestAdditions returns the TestAdditions field if non-nil, zero value otherwise.

### GetTestAdditionsOk

`func (o *BranchChangeStats) GetTestAdditionsOk() (*int32, bool)`

GetTestAdditionsOk returns a tuple with the TestAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdditions

`func (o *BranchChangeStats) SetTestAdditions(v int32)`

SetTestAdditions sets TestAdditions field to given value.

### HasTestAdditions

`func (o *BranchChangeStats) HasTestAdditions() bool`

HasTestAdditions returns a boolean if a field has been set.

### GetTestModifications

`func (o *BranchChangeStats) GetTestModifications() int32`

GetTestModifications returns the TestModifications field if non-nil, zero value otherwise.

### GetTestModificationsOk

`func (o *BranchChangeStats) GetTestModificationsOk() (*int32, bool)`

GetTestModificationsOk returns a tuple with the TestModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModifications

`func (o *BranchChangeStats) SetTestModifications(v int32)`

SetTestModifications sets TestModifications field to given value.

### HasTestModifications

`func (o *BranchChangeStats) HasTestModifications() bool`

HasTestModifications returns a boolean if a field has been set.

### GetTestDeletions

`func (o *BranchChangeStats) GetTestDeletions() int32`

GetTestDeletions returns the TestDeletions field if non-nil, zero value otherwise.

### GetTestDeletionsOk

`func (o *BranchChangeStats) GetTestDeletionsOk() (*int32, bool)`

GetTestDeletionsOk returns a tuple with the TestDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDeletions

`func (o *BranchChangeStats) SetTestDeletions(v int32)`

SetTestDeletions sets TestDeletions field to given value.

### HasTestDeletions

`func (o *BranchChangeStats) HasTestDeletions() bool`

HasTestDeletions returns a boolean if a field has been set.

### GetComponentAdditions

`func (o *BranchChangeStats) GetComponentAdditions() int32`

GetComponentAdditions returns the ComponentAdditions field if non-nil, zero value otherwise.

### GetComponentAdditionsOk

`func (o *BranchChangeStats) GetComponentAdditionsOk() (*int32, bool)`

GetComponentAdditionsOk returns a tuple with the ComponentAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAdditions

`func (o *BranchChangeStats) SetComponentAdditions(v int32)`

SetComponentAdditions sets ComponentAdditions field to given value.

### HasComponentAdditions

`func (o *BranchChangeStats) HasComponentAdditions() bool`

HasComponentAdditions returns a boolean if a field has been set.

### GetComponentModifications

`func (o *BranchChangeStats) GetComponentModifications() int32`

GetComponentModifications returns the ComponentModifications field if non-nil, zero value otherwise.

### GetComponentModificationsOk

`func (o *BranchChangeStats) GetComponentModificationsOk() (*int32, bool)`

GetComponentModificationsOk returns a tuple with the ComponentModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentModifications

`func (o *BranchChangeStats) SetComponentModifications(v int32)`

SetComponentModifications sets ComponentModifications field to given value.

### HasComponentModifications

`func (o *BranchChangeStats) HasComponentModifications() bool`

HasComponentModifications returns a boolean if a field has been set.

### GetComponentDeletions

`func (o *BranchChangeStats) GetComponentDeletions() int32`

GetComponentDeletions returns the ComponentDeletions field if non-nil, zero value otherwise.

### GetComponentDeletionsOk

`func (o *BranchChangeStats) GetComponentDeletionsOk() (*int32, bool)`

GetComponentDeletionsOk returns a tuple with the ComponentDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDeletions

`func (o *BranchChangeStats) SetComponentDeletions(v int32)`

SetComponentDeletions sets ComponentDeletions field to given value.

### HasComponentDeletions

`func (o *BranchChangeStats) HasComponentDeletions() bool`

HasComponentDeletions returns a boolean if a field has been set.

### GetConflicts

`func (o *BranchChangeStats) GetConflicts() int32`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *BranchChangeStats) GetConflictsOk() (*int32, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *BranchChangeStats) SetConflicts(v int32)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *BranchChangeStats) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


