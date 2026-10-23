# MergeIntentSessionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Settled** | Pointer to **bool** | settled reports whether the session&#39;s settlement has been recorded (settlement IS NOT NULL) — NOT whether the session is closed. A closed session with no settlement is exactly the state the merge guard refuses. | [optional] 

## Methods

### NewMergeIntentSessionState

`func NewMergeIntentSessionState() *MergeIntentSessionState`

NewMergeIntentSessionState instantiates a new MergeIntentSessionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeIntentSessionStateWithDefaults

`func NewMergeIntentSessionStateWithDefaults() *MergeIntentSessionState`

NewMergeIntentSessionStateWithDefaults instantiates a new MergeIntentSessionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *MergeIntentSessionState) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MergeIntentSessionState) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MergeIntentSessionState) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *MergeIntentSessionState) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStatus

`func (o *MergeIntentSessionState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MergeIntentSessionState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MergeIntentSessionState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MergeIntentSessionState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettled

`func (o *MergeIntentSessionState) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *MergeIntentSessionState) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *MergeIntentSessionState) SetSettled(v bool)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *MergeIntentSessionState) HasSettled() bool`

HasSettled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


