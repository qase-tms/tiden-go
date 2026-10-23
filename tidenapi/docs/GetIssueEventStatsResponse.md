# GetIssueEventStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **string** |  | [optional] 
**Buckets** | Pointer to [**[]EventBucket**](EventBucket.md) |  | [optional] 
**Last24h** | Pointer to **int32** |  | [optional] 
**Environments** | Pointer to [**[]IssueEnvironmentCount**](IssueEnvironmentCount.md) |  | [optional] 

## Methods

### NewGetIssueEventStatsResponse

`func NewGetIssueEventStatsResponse() *GetIssueEventStatsResponse`

NewGetIssueEventStatsResponse instantiates a new GetIssueEventStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueEventStatsResponseWithDefaults

`func NewGetIssueEventStatsResponseWithDefaults() *GetIssueEventStatsResponse`

NewGetIssueEventStatsResponseWithDefaults instantiates a new GetIssueEventStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *GetIssueEventStatsResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetIssueEventStatsResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetIssueEventStatsResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetIssueEventStatsResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetBuckets

`func (o *GetIssueEventStatsResponse) GetBuckets() []EventBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetIssueEventStatsResponse) GetBucketsOk() (*[]EventBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetIssueEventStatsResponse) SetBuckets(v []EventBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *GetIssueEventStatsResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetLast24h

`func (o *GetIssueEventStatsResponse) GetLast24h() int32`

GetLast24h returns the Last24h field if non-nil, zero value otherwise.

### GetLast24hOk

`func (o *GetIssueEventStatsResponse) GetLast24hOk() (*int32, bool)`

GetLast24hOk returns a tuple with the Last24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast24h

`func (o *GetIssueEventStatsResponse) SetLast24h(v int32)`

SetLast24h sets Last24h field to given value.

### HasLast24h

`func (o *GetIssueEventStatsResponse) HasLast24h() bool`

HasLast24h returns a boolean if a field has been set.

### GetEnvironments

`func (o *GetIssueEventStatsResponse) GetEnvironments() []IssueEnvironmentCount`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *GetIssueEventStatsResponse) GetEnvironmentsOk() (*[]IssueEnvironmentCount, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *GetIssueEventStatsResponse) SetEnvironments(v []IssueEnvironmentCount)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *GetIssueEventStatsResponse) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


