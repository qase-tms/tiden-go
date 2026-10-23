# EventBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewEventBucket

`func NewEventBucket() *EventBucket`

NewEventBucket instantiates a new EventBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBucketWithDefaults

`func NewEventBucketWithDefaults() *EventBucket`

NewEventBucketWithDefaults instantiates a new EventBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *EventBucket) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventBucket) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventBucket) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventBucket) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetCount

`func (o *EventBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventBucket) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


