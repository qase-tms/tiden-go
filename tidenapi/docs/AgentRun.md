# AgentRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AgentConfigId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**TriggerKind** | Pointer to **string** |  | [optional] 
**TriggerMetaJson** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorSummary** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**SummaryMd** | Pointer to **string** |  | [optional] 
**LlmInputTokens** | Pointer to **string** |  | [optional] 
**LlmOutputTokens** | Pointer to **string** |  | [optional] 
**LlmCostUsdCents** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Result** | Pointer to [**AgentRunResult**](AgentRunResult.md) |  | [optional] 

## Methods

### NewAgentRun

`func NewAgentRun() *AgentRun`

NewAgentRun instantiates a new AgentRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunWithDefaults

`func NewAgentRunWithDefaults() *AgentRun`

NewAgentRunWithDefaults instantiates a new AgentRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAgentConfigId

`func (o *AgentRun) GetAgentConfigId() string`

GetAgentConfigId returns the AgentConfigId field if non-nil, zero value otherwise.

### GetAgentConfigIdOk

`func (o *AgentRun) GetAgentConfigIdOk() (*string, bool)`

GetAgentConfigIdOk returns a tuple with the AgentConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfigId

`func (o *AgentRun) SetAgentConfigId(v string)`

SetAgentConfigId sets AgentConfigId field to given value.

### HasAgentConfigId

`func (o *AgentRun) HasAgentConfigId() bool`

HasAgentConfigId returns a boolean if a field has been set.

### GetProductId

`func (o *AgentRun) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AgentRun) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AgentRun) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AgentRun) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetTriggerKind

`func (o *AgentRun) GetTriggerKind() string`

GetTriggerKind returns the TriggerKind field if non-nil, zero value otherwise.

### GetTriggerKindOk

`func (o *AgentRun) GetTriggerKindOk() (*string, bool)`

GetTriggerKindOk returns a tuple with the TriggerKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerKind

`func (o *AgentRun) SetTriggerKind(v string)`

SetTriggerKind sets TriggerKind field to given value.

### HasTriggerKind

`func (o *AgentRun) HasTriggerKind() bool`

HasTriggerKind returns a boolean if a field has been set.

### GetTriggerMetaJson

`func (o *AgentRun) GetTriggerMetaJson() string`

GetTriggerMetaJson returns the TriggerMetaJson field if non-nil, zero value otherwise.

### GetTriggerMetaJsonOk

`func (o *AgentRun) GetTriggerMetaJsonOk() (*string, bool)`

GetTriggerMetaJsonOk returns a tuple with the TriggerMetaJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMetaJson

`func (o *AgentRun) SetTriggerMetaJson(v string)`

SetTriggerMetaJson sets TriggerMetaJson field to given value.

### HasTriggerMetaJson

`func (o *AgentRun) HasTriggerMetaJson() bool`

HasTriggerMetaJson returns a boolean if a field has been set.

### GetStatus

`func (o *AgentRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStep

`func (o *AgentRun) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *AgentRun) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *AgentRun) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *AgentRun) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetStartedAt

`func (o *AgentRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *AgentRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *AgentRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *AgentRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *AgentRun) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *AgentRun) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *AgentRun) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *AgentRun) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetErrorSummary

`func (o *AgentRun) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *AgentRun) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *AgentRun) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *AgentRun) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetBranchId

`func (o *AgentRun) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *AgentRun) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *AgentRun) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *AgentRun) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetSummaryMd

`func (o *AgentRun) GetSummaryMd() string`

GetSummaryMd returns the SummaryMd field if non-nil, zero value otherwise.

### GetSummaryMdOk

`func (o *AgentRun) GetSummaryMdOk() (*string, bool)`

GetSummaryMdOk returns a tuple with the SummaryMd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryMd

`func (o *AgentRun) SetSummaryMd(v string)`

SetSummaryMd sets SummaryMd field to given value.

### HasSummaryMd

`func (o *AgentRun) HasSummaryMd() bool`

HasSummaryMd returns a boolean if a field has been set.

### GetLlmInputTokens

`func (o *AgentRun) GetLlmInputTokens() string`

GetLlmInputTokens returns the LlmInputTokens field if non-nil, zero value otherwise.

### GetLlmInputTokensOk

`func (o *AgentRun) GetLlmInputTokensOk() (*string, bool)`

GetLlmInputTokensOk returns a tuple with the LlmInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmInputTokens

`func (o *AgentRun) SetLlmInputTokens(v string)`

SetLlmInputTokens sets LlmInputTokens field to given value.

### HasLlmInputTokens

`func (o *AgentRun) HasLlmInputTokens() bool`

HasLlmInputTokens returns a boolean if a field has been set.

### GetLlmOutputTokens

`func (o *AgentRun) GetLlmOutputTokens() string`

GetLlmOutputTokens returns the LlmOutputTokens field if non-nil, zero value otherwise.

### GetLlmOutputTokensOk

`func (o *AgentRun) GetLlmOutputTokensOk() (*string, bool)`

GetLlmOutputTokensOk returns a tuple with the LlmOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmOutputTokens

`func (o *AgentRun) SetLlmOutputTokens(v string)`

SetLlmOutputTokens sets LlmOutputTokens field to given value.

### HasLlmOutputTokens

`func (o *AgentRun) HasLlmOutputTokens() bool`

HasLlmOutputTokens returns a boolean if a field has been set.

### GetLlmCostUsdCents

`func (o *AgentRun) GetLlmCostUsdCents() string`

GetLlmCostUsdCents returns the LlmCostUsdCents field if non-nil, zero value otherwise.

### GetLlmCostUsdCentsOk

`func (o *AgentRun) GetLlmCostUsdCentsOk() (*string, bool)`

GetLlmCostUsdCentsOk returns a tuple with the LlmCostUsdCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmCostUsdCents

`func (o *AgentRun) SetLlmCostUsdCents(v string)`

SetLlmCostUsdCents sets LlmCostUsdCents field to given value.

### HasLlmCostUsdCents

`func (o *AgentRun) HasLlmCostUsdCents() bool`

HasLlmCostUsdCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetResult

`func (o *AgentRun) GetResult() AgentRunResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AgentRun) GetResultOk() (*AgentRunResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AgentRun) SetResult(v AgentRunResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AgentRun) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


