# CreateBranchBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedByAgentRunId** | Pointer to **string** | Set when an agent run is creating the branch on behalf of a user. Surfaced on the Branch message + UI banner. | [optional] 
**SyncRepository** | Pointer to **string** | Sync metadata (drift-sync v1): a sync/&lt;...&gt; branch records the git range its requirement delta covered — repository is the canonical repo id, base &#x3D; the repo watermark the delta was computed from, target &#x3D; repo main HEAD at sync time. On merge, the watermark advances base→target (CAS). Set all three or none; leave unset on every non-sync branch. | [optional] 
**SyncBaseSha** | Pointer to **string** |  | [optional] 
**SyncTargetSha** | Pointer to **string** |  | [optional] 
**CreatedByAgent** | Pointer to **string** | Coding agent creating this branch on the caller&#39;s behalf (e.g. \&quot;claude-code\&quot;, \&quot;codex\&quot;), so name + description + agent can be set in one explicit create. Validated server-side against a fixed allowlist; an unrecognized value is stored as empty string, never as free text. | [optional] 

## Methods

### NewCreateBranchBody

`func NewCreateBranchBody() *CreateBranchBody`

NewCreateBranchBody instantiates a new CreateBranchBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBranchBodyWithDefaults

`func NewCreateBranchBodyWithDefaults() *CreateBranchBody`

NewCreateBranchBodyWithDefaults instantiates a new CreateBranchBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBranchBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBranchBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBranchBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBranchBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateBranchBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBranchBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBranchBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBranchBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedByAgentRunId

`func (o *CreateBranchBody) GetCreatedByAgentRunId() string`

GetCreatedByAgentRunId returns the CreatedByAgentRunId field if non-nil, zero value otherwise.

### GetCreatedByAgentRunIdOk

`func (o *CreateBranchBody) GetCreatedByAgentRunIdOk() (*string, bool)`

GetCreatedByAgentRunIdOk returns a tuple with the CreatedByAgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentRunId

`func (o *CreateBranchBody) SetCreatedByAgentRunId(v string)`

SetCreatedByAgentRunId sets CreatedByAgentRunId field to given value.

### HasCreatedByAgentRunId

`func (o *CreateBranchBody) HasCreatedByAgentRunId() bool`

HasCreatedByAgentRunId returns a boolean if a field has been set.

### GetSyncRepository

`func (o *CreateBranchBody) GetSyncRepository() string`

GetSyncRepository returns the SyncRepository field if non-nil, zero value otherwise.

### GetSyncRepositoryOk

`func (o *CreateBranchBody) GetSyncRepositoryOk() (*string, bool)`

GetSyncRepositoryOk returns a tuple with the SyncRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncRepository

`func (o *CreateBranchBody) SetSyncRepository(v string)`

SetSyncRepository sets SyncRepository field to given value.

### HasSyncRepository

`func (o *CreateBranchBody) HasSyncRepository() bool`

HasSyncRepository returns a boolean if a field has been set.

### GetSyncBaseSha

`func (o *CreateBranchBody) GetSyncBaseSha() string`

GetSyncBaseSha returns the SyncBaseSha field if non-nil, zero value otherwise.

### GetSyncBaseShaOk

`func (o *CreateBranchBody) GetSyncBaseShaOk() (*string, bool)`

GetSyncBaseShaOk returns a tuple with the SyncBaseSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncBaseSha

`func (o *CreateBranchBody) SetSyncBaseSha(v string)`

SetSyncBaseSha sets SyncBaseSha field to given value.

### HasSyncBaseSha

`func (o *CreateBranchBody) HasSyncBaseSha() bool`

HasSyncBaseSha returns a boolean if a field has been set.

### GetSyncTargetSha

`func (o *CreateBranchBody) GetSyncTargetSha() string`

GetSyncTargetSha returns the SyncTargetSha field if non-nil, zero value otherwise.

### GetSyncTargetShaOk

`func (o *CreateBranchBody) GetSyncTargetShaOk() (*string, bool)`

GetSyncTargetShaOk returns a tuple with the SyncTargetSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTargetSha

`func (o *CreateBranchBody) SetSyncTargetSha(v string)`

SetSyncTargetSha sets SyncTargetSha field to given value.

### HasSyncTargetSha

`func (o *CreateBranchBody) HasSyncTargetSha() bool`

HasSyncTargetSha returns a boolean if a field has been set.

### GetCreatedByAgent

`func (o *CreateBranchBody) GetCreatedByAgent() string`

GetCreatedByAgent returns the CreatedByAgent field if non-nil, zero value otherwise.

### GetCreatedByAgentOk

`func (o *CreateBranchBody) GetCreatedByAgentOk() (*string, bool)`

GetCreatedByAgentOk returns a tuple with the CreatedByAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgent

`func (o *CreateBranchBody) SetCreatedByAgent(v string)`

SetCreatedByAgent sets CreatedByAgent field to given value.

### HasCreatedByAgent

`func (o *CreateBranchBody) HasCreatedByAgent() bool`

HasCreatedByAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


