# AgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**AgentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**InputsJson** | Pointer to **string** |  | [optional] 
**LlmCredentialId** | Pointer to **string** |  | [optional] 
**DataCredentialId** | Pointer to **string** |  | [optional] 
**ScheduleCron** | Pointer to **string** |  | [optional] 
**ScheduleTimezone** | Pointer to **string** |  | [optional] 
**NextRunAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAgentConfig

`func NewAgentConfig() *AgentConfig`

NewAgentConfig instantiates a new AgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigWithDefaults

`func NewAgentConfigWithDefaults() *AgentConfig`

NewAgentConfigWithDefaults instantiates a new AgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *AgentConfig) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AgentConfig) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AgentConfig) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AgentConfig) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetAgentType

`func (o *AgentConfig) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentConfig) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentConfig) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *AgentConfig) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetName

`func (o *AgentConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *AgentConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInputsJson

`func (o *AgentConfig) GetInputsJson() string`

GetInputsJson returns the InputsJson field if non-nil, zero value otherwise.

### GetInputsJsonOk

`func (o *AgentConfig) GetInputsJsonOk() (*string, bool)`

GetInputsJsonOk returns a tuple with the InputsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsJson

`func (o *AgentConfig) SetInputsJson(v string)`

SetInputsJson sets InputsJson field to given value.

### HasInputsJson

`func (o *AgentConfig) HasInputsJson() bool`

HasInputsJson returns a boolean if a field has been set.

### GetLlmCredentialId

`func (o *AgentConfig) GetLlmCredentialId() string`

GetLlmCredentialId returns the LlmCredentialId field if non-nil, zero value otherwise.

### GetLlmCredentialIdOk

`func (o *AgentConfig) GetLlmCredentialIdOk() (*string, bool)`

GetLlmCredentialIdOk returns a tuple with the LlmCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmCredentialId

`func (o *AgentConfig) SetLlmCredentialId(v string)`

SetLlmCredentialId sets LlmCredentialId field to given value.

### HasLlmCredentialId

`func (o *AgentConfig) HasLlmCredentialId() bool`

HasLlmCredentialId returns a boolean if a field has been set.

### GetDataCredentialId

`func (o *AgentConfig) GetDataCredentialId() string`

GetDataCredentialId returns the DataCredentialId field if non-nil, zero value otherwise.

### GetDataCredentialIdOk

`func (o *AgentConfig) GetDataCredentialIdOk() (*string, bool)`

GetDataCredentialIdOk returns a tuple with the DataCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCredentialId

`func (o *AgentConfig) SetDataCredentialId(v string)`

SetDataCredentialId sets DataCredentialId field to given value.

### HasDataCredentialId

`func (o *AgentConfig) HasDataCredentialId() bool`

HasDataCredentialId returns a boolean if a field has been set.

### GetScheduleCron

`func (o *AgentConfig) GetScheduleCron() string`

GetScheduleCron returns the ScheduleCron field if non-nil, zero value otherwise.

### GetScheduleCronOk

`func (o *AgentConfig) GetScheduleCronOk() (*string, bool)`

GetScheduleCronOk returns a tuple with the ScheduleCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCron

`func (o *AgentConfig) SetScheduleCron(v string)`

SetScheduleCron sets ScheduleCron field to given value.

### HasScheduleCron

`func (o *AgentConfig) HasScheduleCron() bool`

HasScheduleCron returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *AgentConfig) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *AgentConfig) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *AgentConfig) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *AgentConfig) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetNextRunAt

`func (o *AgentConfig) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *AgentConfig) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *AgentConfig) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *AgentConfig) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


