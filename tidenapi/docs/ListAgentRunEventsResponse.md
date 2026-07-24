# ListAgentRunEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]AgentRunEvent**](AgentRunEvent.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListAgentRunEventsResponse

`func NewListAgentRunEventsResponse() *ListAgentRunEventsResponse`

NewListAgentRunEventsResponse instantiates a new ListAgentRunEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAgentRunEventsResponseWithDefaults

`func NewListAgentRunEventsResponseWithDefaults() *ListAgentRunEventsResponse`

NewListAgentRunEventsResponseWithDefaults instantiates a new ListAgentRunEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListAgentRunEventsResponse) GetEvents() []AgentRunEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListAgentRunEventsResponse) GetEventsOk() (*[]AgentRunEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListAgentRunEventsResponse) SetEvents(v []AgentRunEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListAgentRunEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListAgentRunEventsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAgentRunEventsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAgentRunEventsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAgentRunEventsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


