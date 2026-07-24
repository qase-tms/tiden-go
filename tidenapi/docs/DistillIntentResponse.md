# DistillIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntentBranch** | Pointer to **string** | Name of the intent branch the changes were written to (empty when skipped). | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Updated** | Pointer to **int32** |  | [optional] 
**Dropped** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **bool** |  | [optional] 
**SkipReason** | Pointer to **string** |  | [optional] 

## Methods

### NewDistillIntentResponse

`func NewDistillIntentResponse() *DistillIntentResponse`

NewDistillIntentResponse instantiates a new DistillIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistillIntentResponseWithDefaults

`func NewDistillIntentResponseWithDefaults() *DistillIntentResponse`

NewDistillIntentResponseWithDefaults instantiates a new DistillIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntentBranch

`func (o *DistillIntentResponse) GetIntentBranch() string`

GetIntentBranch returns the IntentBranch field if non-nil, zero value otherwise.

### GetIntentBranchOk

`func (o *DistillIntentResponse) GetIntentBranchOk() (*string, bool)`

GetIntentBranchOk returns a tuple with the IntentBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentBranch

`func (o *DistillIntentResponse) SetIntentBranch(v string)`

SetIntentBranch sets IntentBranch field to given value.

### HasIntentBranch

`func (o *DistillIntentResponse) HasIntentBranch() bool`

HasIntentBranch returns a boolean if a field has been set.

### GetCreated

`func (o *DistillIntentResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DistillIntentResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DistillIntentResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DistillIntentResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DistillIntentResponse) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DistillIntentResponse) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DistillIntentResponse) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DistillIntentResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDropped

`func (o *DistillIntentResponse) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *DistillIntentResponse) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *DistillIntentResponse) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *DistillIntentResponse) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetSkipped

`func (o *DistillIntentResponse) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *DistillIntentResponse) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *DistillIntentResponse) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *DistillIntentResponse) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSkipReason

`func (o *DistillIntentResponse) GetSkipReason() string`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *DistillIntentResponse) GetSkipReasonOk() (*string, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *DistillIntentResponse) SetSkipReason(v string)`

SetSkipReason sets SkipReason field to given value.

### HasSkipReason

`func (o *DistillIntentResponse) HasSkipReason() bool`

HasSkipReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


