# AgentRunEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**DataJson** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentRunEvent

`func NewAgentRunEvent() *AgentRunEvent`

NewAgentRunEvent instantiates a new AgentRunEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunEventWithDefaults

`func NewAgentRunEventWithDefaults() *AgentRunEvent`

NewAgentRunEventWithDefaults instantiates a new AgentRunEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentRunEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRunEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRunEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRunEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunId

`func (o *AgentRunEvent) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AgentRunEvent) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AgentRunEvent) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *AgentRunEvent) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetTs

`func (o *AgentRunEvent) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *AgentRunEvent) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *AgentRunEvent) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *AgentRunEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetLevel

`func (o *AgentRunEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AgentRunEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AgentRunEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AgentRunEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetKind

`func (o *AgentRunEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentRunEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentRunEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentRunEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *AgentRunEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AgentRunEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AgentRunEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AgentRunEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDataJson

`func (o *AgentRunEvent) GetDataJson() string`

GetDataJson returns the DataJson field if non-nil, zero value otherwise.

### GetDataJsonOk

`func (o *AgentRunEvent) GetDataJsonOk() (*string, bool)`

GetDataJsonOk returns a tuple with the DataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataJson

`func (o *AgentRunEvent) SetDataJson(v string)`

SetDataJson sets DataJson field to given value.

### HasDataJson

`func (o *AgentRunEvent) HasDataJson() bool`

HasDataJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


