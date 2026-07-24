# StartAgentRunBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputsOverrideJson** | Pointer to **string** | inputs_override may override inputs_json on the config for this run only. Useful for \&quot;Run with these URLs\&quot; without persisting them on the config. | [optional] 

## Methods

### NewStartAgentRunBody

`func NewStartAgentRunBody() *StartAgentRunBody`

NewStartAgentRunBody instantiates a new StartAgentRunBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartAgentRunBodyWithDefaults

`func NewStartAgentRunBodyWithDefaults() *StartAgentRunBody`

NewStartAgentRunBodyWithDefaults instantiates a new StartAgentRunBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputsOverrideJson

`func (o *StartAgentRunBody) GetInputsOverrideJson() string`

GetInputsOverrideJson returns the InputsOverrideJson field if non-nil, zero value otherwise.

### GetInputsOverrideJsonOk

`func (o *StartAgentRunBody) GetInputsOverrideJsonOk() (*string, bool)`

GetInputsOverrideJsonOk returns a tuple with the InputsOverrideJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsOverrideJson

`func (o *StartAgentRunBody) SetInputsOverrideJson(v string)`

SetInputsOverrideJson sets InputsOverrideJson field to given value.

### HasInputsOverrideJson

`func (o *StartAgentRunBody) HasInputsOverrideJson() bool`

HasInputsOverrideJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


