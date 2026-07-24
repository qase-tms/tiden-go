# AcceptRiskBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**VerdictScope**](VerdictScope.md) |  | [optional] [default to VERDICT_SCOPE_UNSPECIFIED]
**ReleaseId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewAcceptRiskBody

`func NewAcceptRiskBody() *AcceptRiskBody`

NewAcceptRiskBody instantiates a new AcceptRiskBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptRiskBodyWithDefaults

`func NewAcceptRiskBodyWithDefaults() *AcceptRiskBody`

NewAcceptRiskBodyWithDefaults instantiates a new AcceptRiskBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *AcceptRiskBody) GetScope() VerdictScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AcceptRiskBody) GetScopeOk() (*VerdictScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AcceptRiskBody) SetScope(v VerdictScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AcceptRiskBody) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetReleaseId

`func (o *AcceptRiskBody) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *AcceptRiskBody) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *AcceptRiskBody) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *AcceptRiskBody) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetBranch

`func (o *AcceptRiskBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AcceptRiskBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AcceptRiskBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AcceptRiskBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetComponentId

`func (o *AcceptRiskBody) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *AcceptRiskBody) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *AcceptRiskBody) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *AcceptRiskBody) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetReason

`func (o *AcceptRiskBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AcceptRiskBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AcceptRiskBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AcceptRiskBody) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


