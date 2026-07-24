# StreamResultOfAgentRunEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**AgentRunEvent**](AgentRunEvent.md) |  | [optional] 
**Error** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewStreamResultOfAgentRunEvent

`func NewStreamResultOfAgentRunEvent() *StreamResultOfAgentRunEvent`

NewStreamResultOfAgentRunEvent instantiates a new StreamResultOfAgentRunEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfAgentRunEventWithDefaults

`func NewStreamResultOfAgentRunEventWithDefaults() *StreamResultOfAgentRunEvent`

NewStreamResultOfAgentRunEventWithDefaults instantiates a new StreamResultOfAgentRunEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfAgentRunEvent) GetResult() AgentRunEvent`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfAgentRunEvent) GetResultOk() (*AgentRunEvent, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfAgentRunEvent) SetResult(v AgentRunEvent)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfAgentRunEvent) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfAgentRunEvent) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfAgentRunEvent) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfAgentRunEvent) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfAgentRunEvent) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


