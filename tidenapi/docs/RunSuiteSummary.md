# RunSuiteSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **[]string** |  | [optional] 
**Stats** | Pointer to [**RunStats**](RunStats.md) |  | [optional] 

## Methods

### NewRunSuiteSummary

`func NewRunSuiteSummary() *RunSuiteSummary`

NewRunSuiteSummary instantiates a new RunSuiteSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSuiteSummaryWithDefaults

`func NewRunSuiteSummaryWithDefaults() *RunSuiteSummary`

NewRunSuiteSummaryWithDefaults instantiates a new RunSuiteSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RunSuiteSummary) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RunSuiteSummary) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RunSuiteSummary) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RunSuiteSummary) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStats

`func (o *RunSuiteSummary) GetStats() RunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *RunSuiteSummary) GetStatsOk() (*RunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *RunSuiteSummary) SetStats(v RunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *RunSuiteSummary) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


