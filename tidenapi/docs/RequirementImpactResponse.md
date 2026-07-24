# RequirementImpactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedRequirementIds** | Pointer to **[]string** | affected is the full set of requirement IDs reachable from the file seeds. | [optional] 
**CoveringTestIds** | Pointer to **[]string** | covering_test_ids is the set of test IDs covering at least one affected requirement. | [optional] 
**UncoveredRequirementIds** | Pointer to **[]string** | uncovered_requirement_ids is the subset of affected that have zero live test links. | [optional] 

## Methods

### NewRequirementImpactResponse

`func NewRequirementImpactResponse() *RequirementImpactResponse`

NewRequirementImpactResponse instantiates a new RequirementImpactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementImpactResponseWithDefaults

`func NewRequirementImpactResponseWithDefaults() *RequirementImpactResponse`

NewRequirementImpactResponseWithDefaults instantiates a new RequirementImpactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedRequirementIds

`func (o *RequirementImpactResponse) GetAffectedRequirementIds() []string`

GetAffectedRequirementIds returns the AffectedRequirementIds field if non-nil, zero value otherwise.

### GetAffectedRequirementIdsOk

`func (o *RequirementImpactResponse) GetAffectedRequirementIdsOk() (*[]string, bool)`

GetAffectedRequirementIdsOk returns a tuple with the AffectedRequirementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRequirementIds

`func (o *RequirementImpactResponse) SetAffectedRequirementIds(v []string)`

SetAffectedRequirementIds sets AffectedRequirementIds field to given value.

### HasAffectedRequirementIds

`func (o *RequirementImpactResponse) HasAffectedRequirementIds() bool`

HasAffectedRequirementIds returns a boolean if a field has been set.

### GetCoveringTestIds

`func (o *RequirementImpactResponse) GetCoveringTestIds() []string`

GetCoveringTestIds returns the CoveringTestIds field if non-nil, zero value otherwise.

### GetCoveringTestIdsOk

`func (o *RequirementImpactResponse) GetCoveringTestIdsOk() (*[]string, bool)`

GetCoveringTestIdsOk returns a tuple with the CoveringTestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveringTestIds

`func (o *RequirementImpactResponse) SetCoveringTestIds(v []string)`

SetCoveringTestIds sets CoveringTestIds field to given value.

### HasCoveringTestIds

`func (o *RequirementImpactResponse) HasCoveringTestIds() bool`

HasCoveringTestIds returns a boolean if a field has been set.

### GetUncoveredRequirementIds

`func (o *RequirementImpactResponse) GetUncoveredRequirementIds() []string`

GetUncoveredRequirementIds returns the UncoveredRequirementIds field if non-nil, zero value otherwise.

### GetUncoveredRequirementIdsOk

`func (o *RequirementImpactResponse) GetUncoveredRequirementIdsOk() (*[]string, bool)`

GetUncoveredRequirementIdsOk returns a tuple with the UncoveredRequirementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncoveredRequirementIds

`func (o *RequirementImpactResponse) SetUncoveredRequirementIds(v []string)`

SetUncoveredRequirementIds sets UncoveredRequirementIds field to given value.

### HasUncoveredRequirementIds

`func (o *RequirementImpactResponse) HasUncoveredRequirementIds() bool`

HasUncoveredRequirementIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


