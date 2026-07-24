# ListAgentConfigsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configs** | Pointer to [**[]AgentConfig**](AgentConfig.md) |  | [optional] 

## Methods

### NewListAgentConfigsResponse

`func NewListAgentConfigsResponse() *ListAgentConfigsResponse`

NewListAgentConfigsResponse instantiates a new ListAgentConfigsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAgentConfigsResponseWithDefaults

`func NewListAgentConfigsResponseWithDefaults() *ListAgentConfigsResponse`

NewListAgentConfigsResponseWithDefaults instantiates a new ListAgentConfigsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigs

`func (o *ListAgentConfigsResponse) GetConfigs() []AgentConfig`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *ListAgentConfigsResponse) GetConfigsOk() (*[]AgentConfig, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *ListAgentConfigsResponse) SetConfigs(v []AgentConfig)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *ListAgentConfigsResponse) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


