# ComponentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**VerdictStatus**](VerdictStatus.md) |  | [optional] [default to VERDICT_STATUS_UNSPECIFIED]
**ResidualRisk** | Pointer to **float64** |  | [optional] 
**Ceiling** | Pointer to **int32** |  | [optional] 
**Criteria** | Pointer to [**[]CriterionResult**](CriterionResult.md) |  | [optional] 

## Methods

### NewComponentResult

`func NewComponentResult() *ComponentResult`

NewComponentResult instantiates a new ComponentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentResultWithDefaults

`func NewComponentResultWithDefaults() *ComponentResult`

NewComponentResultWithDefaults instantiates a new ComponentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *ComponentResult) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentResult) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentResult) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentResult) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetName

`func (o *ComponentResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ComponentResult) GetStatus() VerdictStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComponentResult) GetStatusOk() (*VerdictStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComponentResult) SetStatus(v VerdictStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComponentResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResidualRisk

`func (o *ComponentResult) GetResidualRisk() float64`

GetResidualRisk returns the ResidualRisk field if non-nil, zero value otherwise.

### GetResidualRiskOk

`func (o *ComponentResult) GetResidualRiskOk() (*float64, bool)`

GetResidualRiskOk returns a tuple with the ResidualRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRisk

`func (o *ComponentResult) SetResidualRisk(v float64)`

SetResidualRisk sets ResidualRisk field to given value.

### HasResidualRisk

`func (o *ComponentResult) HasResidualRisk() bool`

HasResidualRisk returns a boolean if a field has been set.

### GetCeiling

`func (o *ComponentResult) GetCeiling() int32`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *ComponentResult) GetCeilingOk() (*int32, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *ComponentResult) SetCeiling(v int32)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *ComponentResult) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetCriteria

`func (o *ComponentResult) GetCriteria() []CriterionResult`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ComponentResult) GetCriteriaOk() (*[]CriterionResult, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ComponentResult) SetCriteria(v []CriterionResult)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ComponentResult) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


