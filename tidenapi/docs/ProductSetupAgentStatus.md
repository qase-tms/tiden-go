# ProductSetupAgentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** |  | [optional] 
**Detected** | Pointer to **bool** |  | [optional] 
**Wired** | Pointer to **bool** |  | [optional] 

## Methods

### NewProductSetupAgentStatus

`func NewProductSetupAgentStatus() *ProductSetupAgentStatus`

NewProductSetupAgentStatus instantiates a new ProductSetupAgentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSetupAgentStatusWithDefaults

`func NewProductSetupAgentStatusWithDefaults() *ProductSetupAgentStatus`

NewProductSetupAgentStatusWithDefaults instantiates a new ProductSetupAgentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *ProductSetupAgentStatus) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ProductSetupAgentStatus) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ProductSetupAgentStatus) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ProductSetupAgentStatus) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetDetected

`func (o *ProductSetupAgentStatus) GetDetected() bool`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *ProductSetupAgentStatus) GetDetectedOk() (*bool, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *ProductSetupAgentStatus) SetDetected(v bool)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *ProductSetupAgentStatus) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### GetWired

`func (o *ProductSetupAgentStatus) GetWired() bool`

GetWired returns the Wired field if non-nil, zero value otherwise.

### GetWiredOk

`func (o *ProductSetupAgentStatus) GetWiredOk() (*bool, bool)`

GetWiredOk returns a tuple with the Wired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWired

`func (o *ProductSetupAgentStatus) SetWired(v bool)`

SetWired sets Wired field to given value.

### HasWired

`func (o *ProductSetupAgentStatus) HasWired() bool`

HasWired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


