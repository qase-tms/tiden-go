# MatrixComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**VerdictStatus**](VerdictStatus.md) |  | [optional] [default to VERDICT_STATUS_UNSPECIFIED]
**ResidualRisk** | Pointer to **float64** |  | [optional] 
**Ceiling** | Pointer to **int32** |  | [optional] 
**Requirements** | Pointer to [**[]MatrixRequirement**](MatrixRequirement.md) |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 

## Methods

### NewMatrixComponent

`func NewMatrixComponent() *MatrixComponent`

NewMatrixComponent instantiates a new MatrixComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixComponentWithDefaults

`func NewMatrixComponentWithDefaults() *MatrixComponent`

NewMatrixComponentWithDefaults instantiates a new MatrixComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *MatrixComponent) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *MatrixComponent) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *MatrixComponent) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *MatrixComponent) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetName

`func (o *MatrixComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MatrixComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MatrixComponent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MatrixComponent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *MatrixComponent) GetStatus() VerdictStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MatrixComponent) GetStatusOk() (*VerdictStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MatrixComponent) SetStatus(v VerdictStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MatrixComponent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResidualRisk

`func (o *MatrixComponent) GetResidualRisk() float64`

GetResidualRisk returns the ResidualRisk field if non-nil, zero value otherwise.

### GetResidualRiskOk

`func (o *MatrixComponent) GetResidualRiskOk() (*float64, bool)`

GetResidualRiskOk returns a tuple with the ResidualRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRisk

`func (o *MatrixComponent) SetResidualRisk(v float64)`

SetResidualRisk sets ResidualRisk field to given value.

### HasResidualRisk

`func (o *MatrixComponent) HasResidualRisk() bool`

HasResidualRisk returns a boolean if a field has been set.

### GetCeiling

`func (o *MatrixComponent) GetCeiling() int32`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *MatrixComponent) GetCeilingOk() (*int32, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *MatrixComponent) SetCeiling(v int32)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *MatrixComponent) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetRequirements

`func (o *MatrixComponent) GetRequirements() []MatrixRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *MatrixComponent) GetRequirementsOk() (*[]MatrixRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *MatrixComponent) SetRequirements(v []MatrixRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *MatrixComponent) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetRepository

`func (o *MatrixComponent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MatrixComponent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MatrixComponent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MatrixComponent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


