# AdvanceRepoWatermarkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** | Full git commit sha to advance to. | [optional] 
**Reason** | Pointer to **string** | baseline (insert-only, first contact), empty_sync (CAS on expected_current_sha), or bootstrap (unconditional — explicit full re-generation). sync_merge and ingest are internal writers (the branch merge hook and the codebase agent) and are rejected here. | [optional] 
**ExpectedCurrentSha** | Pointer to **string** | Required for empty_sync: the watermark value the caller&#39;s detection saw. The advance is a compare-and-set on it — a lost ordering race returns advanced&#x3D;false instead of overwriting. | [optional] 

## Methods

### NewAdvanceRepoWatermarkBody

`func NewAdvanceRepoWatermarkBody() *AdvanceRepoWatermarkBody`

NewAdvanceRepoWatermarkBody instantiates a new AdvanceRepoWatermarkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvanceRepoWatermarkBodyWithDefaults

`func NewAdvanceRepoWatermarkBodyWithDefaults() *AdvanceRepoWatermarkBody`

NewAdvanceRepoWatermarkBodyWithDefaults instantiates a new AdvanceRepoWatermarkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *AdvanceRepoWatermarkBody) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AdvanceRepoWatermarkBody) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AdvanceRepoWatermarkBody) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AdvanceRepoWatermarkBody) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSha

`func (o *AdvanceRepoWatermarkBody) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *AdvanceRepoWatermarkBody) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *AdvanceRepoWatermarkBody) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *AdvanceRepoWatermarkBody) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetReason

`func (o *AdvanceRepoWatermarkBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdvanceRepoWatermarkBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdvanceRepoWatermarkBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AdvanceRepoWatermarkBody) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExpectedCurrentSha

`func (o *AdvanceRepoWatermarkBody) GetExpectedCurrentSha() string`

GetExpectedCurrentSha returns the ExpectedCurrentSha field if non-nil, zero value otherwise.

### GetExpectedCurrentShaOk

`func (o *AdvanceRepoWatermarkBody) GetExpectedCurrentShaOk() (*string, bool)`

GetExpectedCurrentShaOk returns a tuple with the ExpectedCurrentSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCurrentSha

`func (o *AdvanceRepoWatermarkBody) SetExpectedCurrentSha(v string)`

SetExpectedCurrentSha sets ExpectedCurrentSha field to given value.

### HasExpectedCurrentSha

`func (o *AdvanceRepoWatermarkBody) HasExpectedCurrentSha() bool`

HasExpectedCurrentSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


