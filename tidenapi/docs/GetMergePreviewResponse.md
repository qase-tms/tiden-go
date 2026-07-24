# GetMergePreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Additions** | Pointer to [**[]Requirement**](Requirement.md) |  | [optional] 
**Modifications** | Pointer to [**[]MergeModification**](MergeModification.md) |  | [optional] 
**Deletions** | Pointer to [**[]Requirement**](Requirement.md) |  | [optional] 
**Stats** | Pointer to [**MergeStats**](MergeStats.md) |  | [optional] 
**TestAdditions** | Pointer to [**[]Test**](Test.md) | Test-side merge effects (Phase 4 extension). Empty when the branch holds no test changes. | [optional] 
**TestModifications** | Pointer to [**[]TestMergeModification**](TestMergeModification.md) |  | [optional] 
**TestDeletions** | Pointer to [**[]Test**](Test.md) |  | [optional] 
**ComponentAdditions** | Pointer to [**[]Component**](Component.md) | Component-side merge effects. Empty when the branch holds no component changes. | [optional] 
**ComponentModifications** | Pointer to [**[]ComponentMergeModification**](ComponentMergeModification.md) |  | [optional] 
**ComponentDeletions** | Pointer to [**[]Component**](Component.md) |  | [optional] 

## Methods

### NewGetMergePreviewResponse

`func NewGetMergePreviewResponse() *GetMergePreviewResponse`

NewGetMergePreviewResponse instantiates a new GetMergePreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMergePreviewResponseWithDefaults

`func NewGetMergePreviewResponseWithDefaults() *GetMergePreviewResponse`

NewGetMergePreviewResponseWithDefaults instantiates a new GetMergePreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditions

`func (o *GetMergePreviewResponse) GetAdditions() []Requirement`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *GetMergePreviewResponse) GetAdditionsOk() (*[]Requirement, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *GetMergePreviewResponse) SetAdditions(v []Requirement)`

SetAdditions sets Additions field to given value.

### HasAdditions

`func (o *GetMergePreviewResponse) HasAdditions() bool`

HasAdditions returns a boolean if a field has been set.

### GetModifications

`func (o *GetMergePreviewResponse) GetModifications() []MergeModification`

GetModifications returns the Modifications field if non-nil, zero value otherwise.

### GetModificationsOk

`func (o *GetMergePreviewResponse) GetModificationsOk() (*[]MergeModification, bool)`

GetModificationsOk returns a tuple with the Modifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifications

`func (o *GetMergePreviewResponse) SetModifications(v []MergeModification)`

SetModifications sets Modifications field to given value.

### HasModifications

`func (o *GetMergePreviewResponse) HasModifications() bool`

HasModifications returns a boolean if a field has been set.

### GetDeletions

`func (o *GetMergePreviewResponse) GetDeletions() []Requirement`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *GetMergePreviewResponse) GetDeletionsOk() (*[]Requirement, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *GetMergePreviewResponse) SetDeletions(v []Requirement)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *GetMergePreviewResponse) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetStats

`func (o *GetMergePreviewResponse) GetStats() MergeStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetMergePreviewResponse) GetStatsOk() (*MergeStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetMergePreviewResponse) SetStats(v MergeStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetMergePreviewResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTestAdditions

`func (o *GetMergePreviewResponse) GetTestAdditions() []Test`

GetTestAdditions returns the TestAdditions field if non-nil, zero value otherwise.

### GetTestAdditionsOk

`func (o *GetMergePreviewResponse) GetTestAdditionsOk() (*[]Test, bool)`

GetTestAdditionsOk returns a tuple with the TestAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdditions

`func (o *GetMergePreviewResponse) SetTestAdditions(v []Test)`

SetTestAdditions sets TestAdditions field to given value.

### HasTestAdditions

`func (o *GetMergePreviewResponse) HasTestAdditions() bool`

HasTestAdditions returns a boolean if a field has been set.

### GetTestModifications

`func (o *GetMergePreviewResponse) GetTestModifications() []TestMergeModification`

GetTestModifications returns the TestModifications field if non-nil, zero value otherwise.

### GetTestModificationsOk

`func (o *GetMergePreviewResponse) GetTestModificationsOk() (*[]TestMergeModification, bool)`

GetTestModificationsOk returns a tuple with the TestModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestModifications

`func (o *GetMergePreviewResponse) SetTestModifications(v []TestMergeModification)`

SetTestModifications sets TestModifications field to given value.

### HasTestModifications

`func (o *GetMergePreviewResponse) HasTestModifications() bool`

HasTestModifications returns a boolean if a field has been set.

### GetTestDeletions

`func (o *GetMergePreviewResponse) GetTestDeletions() []Test`

GetTestDeletions returns the TestDeletions field if non-nil, zero value otherwise.

### GetTestDeletionsOk

`func (o *GetMergePreviewResponse) GetTestDeletionsOk() (*[]Test, bool)`

GetTestDeletionsOk returns a tuple with the TestDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDeletions

`func (o *GetMergePreviewResponse) SetTestDeletions(v []Test)`

SetTestDeletions sets TestDeletions field to given value.

### HasTestDeletions

`func (o *GetMergePreviewResponse) HasTestDeletions() bool`

HasTestDeletions returns a boolean if a field has been set.

### GetComponentAdditions

`func (o *GetMergePreviewResponse) GetComponentAdditions() []Component`

GetComponentAdditions returns the ComponentAdditions field if non-nil, zero value otherwise.

### GetComponentAdditionsOk

`func (o *GetMergePreviewResponse) GetComponentAdditionsOk() (*[]Component, bool)`

GetComponentAdditionsOk returns a tuple with the ComponentAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAdditions

`func (o *GetMergePreviewResponse) SetComponentAdditions(v []Component)`

SetComponentAdditions sets ComponentAdditions field to given value.

### HasComponentAdditions

`func (o *GetMergePreviewResponse) HasComponentAdditions() bool`

HasComponentAdditions returns a boolean if a field has been set.

### GetComponentModifications

`func (o *GetMergePreviewResponse) GetComponentModifications() []ComponentMergeModification`

GetComponentModifications returns the ComponentModifications field if non-nil, zero value otherwise.

### GetComponentModificationsOk

`func (o *GetMergePreviewResponse) GetComponentModificationsOk() (*[]ComponentMergeModification, bool)`

GetComponentModificationsOk returns a tuple with the ComponentModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentModifications

`func (o *GetMergePreviewResponse) SetComponentModifications(v []ComponentMergeModification)`

SetComponentModifications sets ComponentModifications field to given value.

### HasComponentModifications

`func (o *GetMergePreviewResponse) HasComponentModifications() bool`

HasComponentModifications returns a boolean if a field has been set.

### GetComponentDeletions

`func (o *GetMergePreviewResponse) GetComponentDeletions() []Component`

GetComponentDeletions returns the ComponentDeletions field if non-nil, zero value otherwise.

### GetComponentDeletionsOk

`func (o *GetMergePreviewResponse) GetComponentDeletionsOk() (*[]Component, bool)`

GetComponentDeletionsOk returns a tuple with the ComponentDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDeletions

`func (o *GetMergePreviewResponse) SetComponentDeletions(v []Component)`

SetComponentDeletions sets ComponentDeletions field to given value.

### HasComponentDeletions

`func (o *GetMergePreviewResponse) HasComponentDeletions() bool`

HasComponentDeletions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


