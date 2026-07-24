# AgentRunResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fetched** | Pointer to **int32** |  | [optional] 
**AlreadyLinked** | Pointer to **int32** |  | [optional] 
**Added** | Pointer to **int32** |  | [optional] 
**Updated** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **int32** |  | [optional] 
**Rejected** | Pointer to **int32** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**BranchName** | Pointer to **string** |  | [optional] 
**Passthrough** | Pointer to **bool** |  | [optional] 
**Changes** | Pointer to [**[]Change**](Change.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Observations** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 

## Methods

### NewAgentRunResult

`func NewAgentRunResult() *AgentRunResult`

NewAgentRunResult instantiates a new AgentRunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunResultWithDefaults

`func NewAgentRunResultWithDefaults() *AgentRunResult`

NewAgentRunResultWithDefaults instantiates a new AgentRunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetched

`func (o *AgentRunResult) GetFetched() int32`

GetFetched returns the Fetched field if non-nil, zero value otherwise.

### GetFetchedOk

`func (o *AgentRunResult) GetFetchedOk() (*int32, bool)`

GetFetchedOk returns a tuple with the Fetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetched

`func (o *AgentRunResult) SetFetched(v int32)`

SetFetched sets Fetched field to given value.

### HasFetched

`func (o *AgentRunResult) HasFetched() bool`

HasFetched returns a boolean if a field has been set.

### GetAlreadyLinked

`func (o *AgentRunResult) GetAlreadyLinked() int32`

GetAlreadyLinked returns the AlreadyLinked field if non-nil, zero value otherwise.

### GetAlreadyLinkedOk

`func (o *AgentRunResult) GetAlreadyLinkedOk() (*int32, bool)`

GetAlreadyLinkedOk returns a tuple with the AlreadyLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyLinked

`func (o *AgentRunResult) SetAlreadyLinked(v int32)`

SetAlreadyLinked sets AlreadyLinked field to given value.

### HasAlreadyLinked

`func (o *AgentRunResult) HasAlreadyLinked() bool`

HasAlreadyLinked returns a boolean if a field has been set.

### GetAdded

`func (o *AgentRunResult) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *AgentRunResult) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *AgentRunResult) SetAdded(v int32)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *AgentRunResult) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetUpdated

`func (o *AgentRunResult) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AgentRunResult) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AgentRunResult) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AgentRunResult) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetSkipped

`func (o *AgentRunResult) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *AgentRunResult) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *AgentRunResult) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *AgentRunResult) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetRejected

`func (o *AgentRunResult) GetRejected() int32`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *AgentRunResult) GetRejectedOk() (*int32, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *AgentRunResult) SetRejected(v int32)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *AgentRunResult) HasRejected() bool`

HasRejected returns a boolean if a field has been set.

### GetBranchId

`func (o *AgentRunResult) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *AgentRunResult) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *AgentRunResult) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *AgentRunResult) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetBranchName

`func (o *AgentRunResult) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *AgentRunResult) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *AgentRunResult) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *AgentRunResult) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetPassthrough

`func (o *AgentRunResult) GetPassthrough() bool`

GetPassthrough returns the Passthrough field if non-nil, zero value otherwise.

### GetPassthroughOk

`func (o *AgentRunResult) GetPassthroughOk() (*bool, bool)`

GetPassthroughOk returns a tuple with the Passthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthrough

`func (o *AgentRunResult) SetPassthrough(v bool)`

SetPassthrough sets Passthrough field to given value.

### HasPassthrough

`func (o *AgentRunResult) HasPassthrough() bool`

HasPassthrough returns a boolean if a field has been set.

### GetChanges

`func (o *AgentRunResult) GetChanges() []Change`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *AgentRunResult) GetChangesOk() (*[]Change, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *AgentRunResult) SetChanges(v []Change)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *AgentRunResult) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetKind

`func (o *AgentRunResult) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentRunResult) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentRunResult) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentRunResult) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetObservations

`func (o *AgentRunResult) GetObservations() int32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *AgentRunResult) GetObservationsOk() (*int32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *AgentRunResult) SetObservations(v int32)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *AgentRunResult) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetDeleted

`func (o *AgentRunResult) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *AgentRunResult) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *AgentRunResult) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *AgentRunResult) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


