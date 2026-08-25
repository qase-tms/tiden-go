# RepoWatermark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** |  | [optional] 
**LastReconciledSha** | Pointer to **string** | Full git commit sha. | [optional] 
**Reason** | Pointer to **string** | Last advance reason: baseline | sync_merge | empty_sync | bootstrap | ingest. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRepoWatermark

`func NewRepoWatermark() *RepoWatermark`

NewRepoWatermark instantiates a new RepoWatermark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepoWatermarkWithDefaults

`func NewRepoWatermarkWithDefaults() *RepoWatermark`

NewRepoWatermarkWithDefaults instantiates a new RepoWatermark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *RepoWatermark) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepoWatermark) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepoWatermark) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *RepoWatermark) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetLastReconciledSha

`func (o *RepoWatermark) GetLastReconciledSha() string`

GetLastReconciledSha returns the LastReconciledSha field if non-nil, zero value otherwise.

### GetLastReconciledShaOk

`func (o *RepoWatermark) GetLastReconciledShaOk() (*string, bool)`

GetLastReconciledShaOk returns a tuple with the LastReconciledSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconciledSha

`func (o *RepoWatermark) SetLastReconciledSha(v string)`

SetLastReconciledSha sets LastReconciledSha field to given value.

### HasLastReconciledSha

`func (o *RepoWatermark) HasLastReconciledSha() bool`

HasLastReconciledSha returns a boolean if a field has been set.

### GetReason

`func (o *RepoWatermark) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RepoWatermark) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RepoWatermark) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RepoWatermark) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RepoWatermark) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RepoWatermark) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RepoWatermark) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RepoWatermark) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


