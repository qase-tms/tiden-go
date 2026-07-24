# CriterionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criterion** | Pointer to **string** | \&quot;pass_rate\&quot; | \&quot;coverage\&quot; | \&quot;risk\&quot; | \&quot;freshness\&quot; | ... | [optional] 
**Status** | Pointer to [**CriterionStatus**](CriterionStatus.md) |  | [optional] [default to CRITERION_STATUS_UNSPECIFIED]
**Score** | Pointer to **float64** |  | [optional] 
**Detail** | Pointer to [**CriterionDetail**](CriterionDetail.md) |  | [optional] 
**Soft** | Pointer to **bool** |  | [optional] 
**Accepted** | Pointer to **bool** |  | [optional] 

## Methods

### NewCriterionResult

`func NewCriterionResult() *CriterionResult`

NewCriterionResult instantiates a new CriterionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionResultWithDefaults

`func NewCriterionResultWithDefaults() *CriterionResult`

NewCriterionResultWithDefaults instantiates a new CriterionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriterion

`func (o *CriterionResult) GetCriterion() string`

GetCriterion returns the Criterion field if non-nil, zero value otherwise.

### GetCriterionOk

`func (o *CriterionResult) GetCriterionOk() (*string, bool)`

GetCriterionOk returns a tuple with the Criterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterion

`func (o *CriterionResult) SetCriterion(v string)`

SetCriterion sets Criterion field to given value.

### HasCriterion

`func (o *CriterionResult) HasCriterion() bool`

HasCriterion returns a boolean if a field has been set.

### GetStatus

`func (o *CriterionResult) GetStatus() CriterionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CriterionResult) GetStatusOk() (*CriterionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CriterionResult) SetStatus(v CriterionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CriterionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScore

`func (o *CriterionResult) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CriterionResult) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CriterionResult) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *CriterionResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDetail

`func (o *CriterionResult) GetDetail() CriterionDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CriterionResult) GetDetailOk() (*CriterionDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CriterionResult) SetDetail(v CriterionDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CriterionResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSoft

`func (o *CriterionResult) GetSoft() bool`

GetSoft returns the Soft field if non-nil, zero value otherwise.

### GetSoftOk

`func (o *CriterionResult) GetSoftOk() (*bool, bool)`

GetSoftOk returns a tuple with the Soft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoft

`func (o *CriterionResult) SetSoft(v bool)`

SetSoft sets Soft field to given value.

### HasSoft

`func (o *CriterionResult) HasSoft() bool`

HasSoft returns a boolean if a field has been set.

### GetAccepted

`func (o *CriterionResult) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *CriterionResult) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *CriterionResult) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *CriterionResult) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


