# ComputeVerdictBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**VerdictScope**](VerdictScope.md) |  | [optional] [default to VERDICT_SCOPE_UNSPECIFIED]
**ReleaseId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**SubjectType** | Pointer to **string** |  | [optional] 
**SubjectId** | Pointer to **string** |  | [optional] 

## Methods

### NewComputeVerdictBody

`func NewComputeVerdictBody() *ComputeVerdictBody`

NewComputeVerdictBody instantiates a new ComputeVerdictBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeVerdictBodyWithDefaults

`func NewComputeVerdictBodyWithDefaults() *ComputeVerdictBody`

NewComputeVerdictBodyWithDefaults instantiates a new ComputeVerdictBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ComputeVerdictBody) GetScope() VerdictScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ComputeVerdictBody) GetScopeOk() (*VerdictScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ComputeVerdictBody) SetScope(v VerdictScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ComputeVerdictBody) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetReleaseId

`func (o *ComputeVerdictBody) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *ComputeVerdictBody) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *ComputeVerdictBody) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *ComputeVerdictBody) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetBranch

`func (o *ComputeVerdictBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ComputeVerdictBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ComputeVerdictBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ComputeVerdictBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSubjectType

`func (o *ComputeVerdictBody) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *ComputeVerdictBody) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *ComputeVerdictBody) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *ComputeVerdictBody) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetSubjectId

`func (o *ComputeVerdictBody) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ComputeVerdictBody) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ComputeVerdictBody) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *ComputeVerdictBody) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


