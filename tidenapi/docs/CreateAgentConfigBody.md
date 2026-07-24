# CreateAgentConfigBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**InputsJson** | Pointer to **string** |  | [optional] 
**LlmCredentialId** | Pointer to **string** |  | [optional] 
**DataCredentialId** | Pointer to **string** |  | [optional] 
**ScheduleCron** | Pointer to **string** |  | [optional] 
**ScheduleTimezone** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateAgentConfigBody

`func NewCreateAgentConfigBody() *CreateAgentConfigBody`

NewCreateAgentConfigBody instantiates a new CreateAgentConfigBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentConfigBodyWithDefaults

`func NewCreateAgentConfigBodyWithDefaults() *CreateAgentConfigBody`

NewCreateAgentConfigBodyWithDefaults instantiates a new CreateAgentConfigBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *CreateAgentConfigBody) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *CreateAgentConfigBody) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *CreateAgentConfigBody) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *CreateAgentConfigBody) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetName

`func (o *CreateAgentConfigBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgentConfigBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgentConfigBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAgentConfigBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateAgentConfigBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAgentConfigBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAgentConfigBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAgentConfigBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInputsJson

`func (o *CreateAgentConfigBody) GetInputsJson() string`

GetInputsJson returns the InputsJson field if non-nil, zero value otherwise.

### GetInputsJsonOk

`func (o *CreateAgentConfigBody) GetInputsJsonOk() (*string, bool)`

GetInputsJsonOk returns a tuple with the InputsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsJson

`func (o *CreateAgentConfigBody) SetInputsJson(v string)`

SetInputsJson sets InputsJson field to given value.

### HasInputsJson

`func (o *CreateAgentConfigBody) HasInputsJson() bool`

HasInputsJson returns a boolean if a field has been set.

### GetLlmCredentialId

`func (o *CreateAgentConfigBody) GetLlmCredentialId() string`

GetLlmCredentialId returns the LlmCredentialId field if non-nil, zero value otherwise.

### GetLlmCredentialIdOk

`func (o *CreateAgentConfigBody) GetLlmCredentialIdOk() (*string, bool)`

GetLlmCredentialIdOk returns a tuple with the LlmCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmCredentialId

`func (o *CreateAgentConfigBody) SetLlmCredentialId(v string)`

SetLlmCredentialId sets LlmCredentialId field to given value.

### HasLlmCredentialId

`func (o *CreateAgentConfigBody) HasLlmCredentialId() bool`

HasLlmCredentialId returns a boolean if a field has been set.

### GetDataCredentialId

`func (o *CreateAgentConfigBody) GetDataCredentialId() string`

GetDataCredentialId returns the DataCredentialId field if non-nil, zero value otherwise.

### GetDataCredentialIdOk

`func (o *CreateAgentConfigBody) GetDataCredentialIdOk() (*string, bool)`

GetDataCredentialIdOk returns a tuple with the DataCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCredentialId

`func (o *CreateAgentConfigBody) SetDataCredentialId(v string)`

SetDataCredentialId sets DataCredentialId field to given value.

### HasDataCredentialId

`func (o *CreateAgentConfigBody) HasDataCredentialId() bool`

HasDataCredentialId returns a boolean if a field has been set.

### GetScheduleCron

`func (o *CreateAgentConfigBody) GetScheduleCron() string`

GetScheduleCron returns the ScheduleCron field if non-nil, zero value otherwise.

### GetScheduleCronOk

`func (o *CreateAgentConfigBody) GetScheduleCronOk() (*string, bool)`

GetScheduleCronOk returns a tuple with the ScheduleCron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCron

`func (o *CreateAgentConfigBody) SetScheduleCron(v string)`

SetScheduleCron sets ScheduleCron field to given value.

### HasScheduleCron

`func (o *CreateAgentConfigBody) HasScheduleCron() bool`

HasScheduleCron returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *CreateAgentConfigBody) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *CreateAgentConfigBody) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *CreateAgentConfigBody) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *CreateAgentConfigBody) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


