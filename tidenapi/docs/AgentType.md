# AgentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InputSchemaJson** | Pointer to **string** |  | [optional] 
**SupportedProviders** | Pointer to **[]string** |  | [optional] 
**DefaultModels** | Pointer to **map[string]string** |  | [optional] 
**NeedsDataCredential** | Pointer to **bool** | true for jira/asana/etc. | [optional] 
**ProducesBranch** | Pointer to **bool** |  | [optional] 

## Methods

### NewAgentType

`func NewAgentType() *AgentType`

NewAgentType instantiates a new AgentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTypeWithDefaults

`func NewAgentTypeWithDefaults() *AgentType`

NewAgentTypeWithDefaults instantiates a new AgentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AgentType) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentType) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentType) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AgentType) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTitle

`func (o *AgentType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentType) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *AgentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputSchemaJson

`func (o *AgentType) GetInputSchemaJson() string`

GetInputSchemaJson returns the InputSchemaJson field if non-nil, zero value otherwise.

### GetInputSchemaJsonOk

`func (o *AgentType) GetInputSchemaJsonOk() (*string, bool)`

GetInputSchemaJsonOk returns a tuple with the InputSchemaJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchemaJson

`func (o *AgentType) SetInputSchemaJson(v string)`

SetInputSchemaJson sets InputSchemaJson field to given value.

### HasInputSchemaJson

`func (o *AgentType) HasInputSchemaJson() bool`

HasInputSchemaJson returns a boolean if a field has been set.

### GetSupportedProviders

`func (o *AgentType) GetSupportedProviders() []string`

GetSupportedProviders returns the SupportedProviders field if non-nil, zero value otherwise.

### GetSupportedProvidersOk

`func (o *AgentType) GetSupportedProvidersOk() (*[]string, bool)`

GetSupportedProvidersOk returns a tuple with the SupportedProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProviders

`func (o *AgentType) SetSupportedProviders(v []string)`

SetSupportedProviders sets SupportedProviders field to given value.

### HasSupportedProviders

`func (o *AgentType) HasSupportedProviders() bool`

HasSupportedProviders returns a boolean if a field has been set.

### GetDefaultModels

`func (o *AgentType) GetDefaultModels() map[string]string`

GetDefaultModels returns the DefaultModels field if non-nil, zero value otherwise.

### GetDefaultModelsOk

`func (o *AgentType) GetDefaultModelsOk() (*map[string]string, bool)`

GetDefaultModelsOk returns a tuple with the DefaultModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModels

`func (o *AgentType) SetDefaultModels(v map[string]string)`

SetDefaultModels sets DefaultModels field to given value.

### HasDefaultModels

`func (o *AgentType) HasDefaultModels() bool`

HasDefaultModels returns a boolean if a field has been set.

### GetNeedsDataCredential

`func (o *AgentType) GetNeedsDataCredential() bool`

GetNeedsDataCredential returns the NeedsDataCredential field if non-nil, zero value otherwise.

### GetNeedsDataCredentialOk

`func (o *AgentType) GetNeedsDataCredentialOk() (*bool, bool)`

GetNeedsDataCredentialOk returns a tuple with the NeedsDataCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsDataCredential

`func (o *AgentType) SetNeedsDataCredential(v bool)`

SetNeedsDataCredential sets NeedsDataCredential field to given value.

### HasNeedsDataCredential

`func (o *AgentType) HasNeedsDataCredential() bool`

HasNeedsDataCredential returns a boolean if a field has been set.

### GetProducesBranch

`func (o *AgentType) GetProducesBranch() bool`

GetProducesBranch returns the ProducesBranch field if non-nil, zero value otherwise.

### GetProducesBranchOk

`func (o *AgentType) GetProducesBranchOk() (*bool, bool)`

GetProducesBranchOk returns a tuple with the ProducesBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducesBranch

`func (o *AgentType) SetProducesBranch(v bool)`

SetProducesBranch sets ProducesBranch field to given value.

### HasProducesBranch

`func (o *AgentType) HasProducesBranch() bool`

HasProducesBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


