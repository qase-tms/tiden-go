# SubjectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectType** | Pointer to **string** |  | [optional] 
**SubjectId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**VerdictStatus**](VerdictStatus.md) |  | [optional] [default to VERDICT_STATUS_UNSPECIFIED]
**Criteria** | Pointer to [**[]CriterionResult**](CriterionResult.md) |  | [optional] 
**ResidualRisk** | Pointer to **float64** |  | [optional] 
**Ceiling** | Pointer to **int32** |  | [optional] 
**RiskSource** | Pointer to **string** | Provenance (feature fan-in): risk_source names the component whose risk profile drove a feature&#39;s risk fan-in; issue_sources names the components that contributed open issues. Empty for component/product subjects. | [optional] 
**IssueSources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSubjectResult

`func NewSubjectResult() *SubjectResult`

NewSubjectResult instantiates a new SubjectResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectResultWithDefaults

`func NewSubjectResultWithDefaults() *SubjectResult`

NewSubjectResultWithDefaults instantiates a new SubjectResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectType

`func (o *SubjectResult) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *SubjectResult) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *SubjectResult) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *SubjectResult) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetSubjectId

`func (o *SubjectResult) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SubjectResult) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SubjectResult) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *SubjectResult) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetName

`func (o *SubjectResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubjectResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubjectResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubjectResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *SubjectResult) GetStatus() VerdictStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubjectResult) GetStatusOk() (*VerdictStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubjectResult) SetStatus(v VerdictStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubjectResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCriteria

`func (o *SubjectResult) GetCriteria() []CriterionResult`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *SubjectResult) GetCriteriaOk() (*[]CriterionResult, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *SubjectResult) SetCriteria(v []CriterionResult)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *SubjectResult) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetResidualRisk

`func (o *SubjectResult) GetResidualRisk() float64`

GetResidualRisk returns the ResidualRisk field if non-nil, zero value otherwise.

### GetResidualRiskOk

`func (o *SubjectResult) GetResidualRiskOk() (*float64, bool)`

GetResidualRiskOk returns a tuple with the ResidualRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualRisk

`func (o *SubjectResult) SetResidualRisk(v float64)`

SetResidualRisk sets ResidualRisk field to given value.

### HasResidualRisk

`func (o *SubjectResult) HasResidualRisk() bool`

HasResidualRisk returns a boolean if a field has been set.

### GetCeiling

`func (o *SubjectResult) GetCeiling() int32`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *SubjectResult) GetCeilingOk() (*int32, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *SubjectResult) SetCeiling(v int32)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *SubjectResult) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetRiskSource

`func (o *SubjectResult) GetRiskSource() string`

GetRiskSource returns the RiskSource field if non-nil, zero value otherwise.

### GetRiskSourceOk

`func (o *SubjectResult) GetRiskSourceOk() (*string, bool)`

GetRiskSourceOk returns a tuple with the RiskSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskSource

`func (o *SubjectResult) SetRiskSource(v string)`

SetRiskSource sets RiskSource field to given value.

### HasRiskSource

`func (o *SubjectResult) HasRiskSource() bool`

HasRiskSource returns a boolean if a field has been set.

### GetIssueSources

`func (o *SubjectResult) GetIssueSources() []string`

GetIssueSources returns the IssueSources field if non-nil, zero value otherwise.

### GetIssueSourcesOk

`func (o *SubjectResult) GetIssueSourcesOk() (*[]string, bool)`

GetIssueSourcesOk returns a tuple with the IssueSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSources

`func (o *SubjectResult) SetIssueSources(v []string)`

SetIssueSources sets IssueSources field to given value.

### HasIssueSources

`func (o *SubjectResult) HasIssueSources() bool`

HasIssueSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


