# CoverageGap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirement** | Pointer to [**Requirement**](Requirement.md) |  | [optional] 
**CoverageStatus** | Pointer to **string** |  | [optional] 
**LinkedTestCount** | Pointer to **int32** |  | [optional] 
**ProposedTestCount** | Pointer to **int32** |  | [optional] 
**StaleTestCount** | Pointer to **int32** |  | [optional] 
**CoverageItemCount** | Pointer to **int32** |  | [optional] 
**LastTestUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CoverageStatusReason** | Pointer to **string** |  | [optional] 
**RankingSignals** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCoverageGap

`func NewCoverageGap() *CoverageGap`

NewCoverageGap instantiates a new CoverageGap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverageGapWithDefaults

`func NewCoverageGapWithDefaults() *CoverageGap`

NewCoverageGapWithDefaults instantiates a new CoverageGap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirement

`func (o *CoverageGap) GetRequirement() Requirement`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *CoverageGap) GetRequirementOk() (*Requirement, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *CoverageGap) SetRequirement(v Requirement)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *CoverageGap) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### GetCoverageStatus

`func (o *CoverageGap) GetCoverageStatus() string`

GetCoverageStatus returns the CoverageStatus field if non-nil, zero value otherwise.

### GetCoverageStatusOk

`func (o *CoverageGap) GetCoverageStatusOk() (*string, bool)`

GetCoverageStatusOk returns a tuple with the CoverageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatus

`func (o *CoverageGap) SetCoverageStatus(v string)`

SetCoverageStatus sets CoverageStatus field to given value.

### HasCoverageStatus

`func (o *CoverageGap) HasCoverageStatus() bool`

HasCoverageStatus returns a boolean if a field has been set.

### GetLinkedTestCount

`func (o *CoverageGap) GetLinkedTestCount() int32`

GetLinkedTestCount returns the LinkedTestCount field if non-nil, zero value otherwise.

### GetLinkedTestCountOk

`func (o *CoverageGap) GetLinkedTestCountOk() (*int32, bool)`

GetLinkedTestCountOk returns a tuple with the LinkedTestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTestCount

`func (o *CoverageGap) SetLinkedTestCount(v int32)`

SetLinkedTestCount sets LinkedTestCount field to given value.

### HasLinkedTestCount

`func (o *CoverageGap) HasLinkedTestCount() bool`

HasLinkedTestCount returns a boolean if a field has been set.

### GetProposedTestCount

`func (o *CoverageGap) GetProposedTestCount() int32`

GetProposedTestCount returns the ProposedTestCount field if non-nil, zero value otherwise.

### GetProposedTestCountOk

`func (o *CoverageGap) GetProposedTestCountOk() (*int32, bool)`

GetProposedTestCountOk returns a tuple with the ProposedTestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTestCount

`func (o *CoverageGap) SetProposedTestCount(v int32)`

SetProposedTestCount sets ProposedTestCount field to given value.

### HasProposedTestCount

`func (o *CoverageGap) HasProposedTestCount() bool`

HasProposedTestCount returns a boolean if a field has been set.

### GetStaleTestCount

`func (o *CoverageGap) GetStaleTestCount() int32`

GetStaleTestCount returns the StaleTestCount field if non-nil, zero value otherwise.

### GetStaleTestCountOk

`func (o *CoverageGap) GetStaleTestCountOk() (*int32, bool)`

GetStaleTestCountOk returns a tuple with the StaleTestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleTestCount

`func (o *CoverageGap) SetStaleTestCount(v int32)`

SetStaleTestCount sets StaleTestCount field to given value.

### HasStaleTestCount

`func (o *CoverageGap) HasStaleTestCount() bool`

HasStaleTestCount returns a boolean if a field has been set.

### GetCoverageItemCount

`func (o *CoverageGap) GetCoverageItemCount() int32`

GetCoverageItemCount returns the CoverageItemCount field if non-nil, zero value otherwise.

### GetCoverageItemCountOk

`func (o *CoverageGap) GetCoverageItemCountOk() (*int32, bool)`

GetCoverageItemCountOk returns a tuple with the CoverageItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageItemCount

`func (o *CoverageGap) SetCoverageItemCount(v int32)`

SetCoverageItemCount sets CoverageItemCount field to given value.

### HasCoverageItemCount

`func (o *CoverageGap) HasCoverageItemCount() bool`

HasCoverageItemCount returns a boolean if a field has been set.

### GetLastTestUpdatedAt

`func (o *CoverageGap) GetLastTestUpdatedAt() time.Time`

GetLastTestUpdatedAt returns the LastTestUpdatedAt field if non-nil, zero value otherwise.

### GetLastTestUpdatedAtOk

`func (o *CoverageGap) GetLastTestUpdatedAtOk() (*time.Time, bool)`

GetLastTestUpdatedAtOk returns a tuple with the LastTestUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestUpdatedAt

`func (o *CoverageGap) SetLastTestUpdatedAt(v time.Time)`

SetLastTestUpdatedAt sets LastTestUpdatedAt field to given value.

### HasLastTestUpdatedAt

`func (o *CoverageGap) HasLastTestUpdatedAt() bool`

HasLastTestUpdatedAt returns a boolean if a field has been set.

### GetCoverageStatusReason

`func (o *CoverageGap) GetCoverageStatusReason() string`

GetCoverageStatusReason returns the CoverageStatusReason field if non-nil, zero value otherwise.

### GetCoverageStatusReasonOk

`func (o *CoverageGap) GetCoverageStatusReasonOk() (*string, bool)`

GetCoverageStatusReasonOk returns a tuple with the CoverageStatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatusReason

`func (o *CoverageGap) SetCoverageStatusReason(v string)`

SetCoverageStatusReason sets CoverageStatusReason field to given value.

### HasCoverageStatusReason

`func (o *CoverageGap) HasCoverageStatusReason() bool`

HasCoverageStatusReason returns a boolean if a field has been set.

### GetRankingSignals

`func (o *CoverageGap) GetRankingSignals() []string`

GetRankingSignals returns the RankingSignals field if non-nil, zero value otherwise.

### GetRankingSignalsOk

`func (o *CoverageGap) GetRankingSignalsOk() (*[]string, bool)`

GetRankingSignalsOk returns a tuple with the RankingSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankingSignals

`func (o *CoverageGap) SetRankingSignals(v []string)`

SetRankingSignals sets RankingSignals field to given value.

### HasRankingSignals

`func (o *CoverageGap) HasRankingSignals() bool`

HasRankingSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


