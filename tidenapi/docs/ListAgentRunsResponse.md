# ListAgentRunsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | Pointer to [**[]AgentRun**](AgentRun.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListAgentRunsResponse

`func NewListAgentRunsResponse() *ListAgentRunsResponse`

NewListAgentRunsResponse instantiates a new ListAgentRunsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAgentRunsResponseWithDefaults

`func NewListAgentRunsResponseWithDefaults() *ListAgentRunsResponse`

NewListAgentRunsResponseWithDefaults instantiates a new ListAgentRunsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *ListAgentRunsResponse) GetRuns() []AgentRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ListAgentRunsResponse) GetRunsOk() (*[]AgentRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ListAgentRunsResponse) SetRuns(v []AgentRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ListAgentRunsResponse) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListAgentRunsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListAgentRunsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListAgentRunsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListAgentRunsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


