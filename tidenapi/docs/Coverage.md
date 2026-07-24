# Coverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Covered** | Pointer to **[]string** |  | [optional] 
**Gaps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoverage

`func NewCoverage() *Coverage`

NewCoverage instantiates a new Coverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverageWithDefaults

`func NewCoverageWithDefaults() *Coverage`

NewCoverageWithDefaults instantiates a new Coverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCovered

`func (o *Coverage) GetCovered() []string`

GetCovered returns the Covered field if non-nil, zero value otherwise.

### GetCoveredOk

`func (o *Coverage) GetCoveredOk() (*[]string, bool)`

GetCoveredOk returns a tuple with the Covered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovered

`func (o *Coverage) SetCovered(v []string)`

SetCovered sets Covered field to given value.

### HasCovered

`func (o *Coverage) HasCovered() bool`

HasCovered returns a boolean if a field has been set.

### GetGaps

`func (o *Coverage) GetGaps() []string`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *Coverage) GetGapsOk() (*[]string, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *Coverage) SetGaps(v []string)`

SetGaps sets Gaps field to given value.

### HasGaps

`func (o *Coverage) HasGaps() bool`

HasGaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


