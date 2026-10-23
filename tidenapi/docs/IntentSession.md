# IntentSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**Objective** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ClosedReason** | Pointer to **string** |  | [optional] 
**ScopeBand** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Executor** | Pointer to **string** |  | [optional] 
**GitRepo** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**StartHead** | Pointer to **string** |  | [optional] 
**RetrievedIds** | Pointer to **[]string** |  | [optional] 
**Verdict** | Pointer to **string** |  | [optional] 
**VerdictAt** | Pointer to **time.Time** |  | [optional] 
**ProgressSnapshot** | Pointer to **map[string]interface{}** |  | [optional] 
**CloseMeta** | Pointer to **map[string]interface{}** |  | [optional] 
**Settlement** | Pointer to [**IntentSessionSettlement**](IntentSessionSettlement.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**LastActivityAt** | Pointer to **time.Time** |  | [optional] 
**EffectiveStatus** | Pointer to **string** | effective_status is computed at read time, never persisted and never written by a background job: equals status normally; \&quot;expired\&quot; when status&#x3D;provisional and now - last_activity_at exceeds the server&#39;s provisional TTL; \&quot;stale\&quot; when status&#x3D;materialized and it exceeds the materialized TTL. A closed session is always exactly \&quot;closed\&quot;. The INTENT_UNDISTILLED merge guard deliberately does not consult this field. | [optional] 

## Methods

### NewIntentSession

`func NewIntentSession() *IntentSession`

NewIntentSession instantiates a new IntentSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntentSessionWithDefaults

`func NewIntentSessionWithDefaults() *IntentSession`

NewIntentSessionWithDefaults instantiates a new IntentSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntentSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntentSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntentSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntentSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *IntentSession) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *IntentSession) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *IntentSession) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *IntentSession) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetBranchId

`func (o *IntentSession) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *IntentSession) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *IntentSession) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *IntentSession) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetObjective

`func (o *IntentSession) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *IntentSession) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *IntentSession) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *IntentSession) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetStatus

`func (o *IntentSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntentSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntentSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntentSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClosedReason

`func (o *IntentSession) GetClosedReason() string`

GetClosedReason returns the ClosedReason field if non-nil, zero value otherwise.

### GetClosedReasonOk

`func (o *IntentSession) GetClosedReasonOk() (*string, bool)`

GetClosedReasonOk returns a tuple with the ClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedReason

`func (o *IntentSession) SetClosedReason(v string)`

SetClosedReason sets ClosedReason field to given value.

### HasClosedReason

`func (o *IntentSession) HasClosedReason() bool`

HasClosedReason returns a boolean if a field has been set.

### GetScopeBand

`func (o *IntentSession) GetScopeBand() string`

GetScopeBand returns the ScopeBand field if non-nil, zero value otherwise.

### GetScopeBandOk

`func (o *IntentSession) GetScopeBandOk() (*string, bool)`

GetScopeBandOk returns a tuple with the ScopeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeBand

`func (o *IntentSession) SetScopeBand(v string)`

SetScopeBand sets ScopeBand field to given value.

### HasScopeBand

`func (o *IntentSession) HasScopeBand() bool`

HasScopeBand returns a boolean if a field has been set.

### GetAgent

`func (o *IntentSession) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *IntentSession) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *IntentSession) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *IntentSession) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExecutor

`func (o *IntentSession) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *IntentSession) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *IntentSession) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *IntentSession) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetGitRepo

`func (o *IntentSession) GetGitRepo() string`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *IntentSession) GetGitRepoOk() (*string, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *IntentSession) SetGitRepo(v string)`

SetGitRepo sets GitRepo field to given value.

### HasGitRepo

`func (o *IntentSession) HasGitRepo() bool`

HasGitRepo returns a boolean if a field has been set.

### GetGitBranch

`func (o *IntentSession) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *IntentSession) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *IntentSession) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *IntentSession) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetStartHead

`func (o *IntentSession) GetStartHead() string`

GetStartHead returns the StartHead field if non-nil, zero value otherwise.

### GetStartHeadOk

`func (o *IntentSession) GetStartHeadOk() (*string, bool)`

GetStartHeadOk returns a tuple with the StartHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHead

`func (o *IntentSession) SetStartHead(v string)`

SetStartHead sets StartHead field to given value.

### HasStartHead

`func (o *IntentSession) HasStartHead() bool`

HasStartHead returns a boolean if a field has been set.

### GetRetrievedIds

`func (o *IntentSession) GetRetrievedIds() []string`

GetRetrievedIds returns the RetrievedIds field if non-nil, zero value otherwise.

### GetRetrievedIdsOk

`func (o *IntentSession) GetRetrievedIdsOk() (*[]string, bool)`

GetRetrievedIdsOk returns a tuple with the RetrievedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedIds

`func (o *IntentSession) SetRetrievedIds(v []string)`

SetRetrievedIds sets RetrievedIds field to given value.

### HasRetrievedIds

`func (o *IntentSession) HasRetrievedIds() bool`

HasRetrievedIds returns a boolean if a field has been set.

### GetVerdict

`func (o *IntentSession) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *IntentSession) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *IntentSession) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *IntentSession) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetVerdictAt

`func (o *IntentSession) GetVerdictAt() time.Time`

GetVerdictAt returns the VerdictAt field if non-nil, zero value otherwise.

### GetVerdictAtOk

`func (o *IntentSession) GetVerdictAtOk() (*time.Time, bool)`

GetVerdictAtOk returns a tuple with the VerdictAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdictAt

`func (o *IntentSession) SetVerdictAt(v time.Time)`

SetVerdictAt sets VerdictAt field to given value.

### HasVerdictAt

`func (o *IntentSession) HasVerdictAt() bool`

HasVerdictAt returns a boolean if a field has been set.

### GetProgressSnapshot

`func (o *IntentSession) GetProgressSnapshot() map[string]interface{}`

GetProgressSnapshot returns the ProgressSnapshot field if non-nil, zero value otherwise.

### GetProgressSnapshotOk

`func (o *IntentSession) GetProgressSnapshotOk() (*map[string]interface{}, bool)`

GetProgressSnapshotOk returns a tuple with the ProgressSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSnapshot

`func (o *IntentSession) SetProgressSnapshot(v map[string]interface{})`

SetProgressSnapshot sets ProgressSnapshot field to given value.

### HasProgressSnapshot

`func (o *IntentSession) HasProgressSnapshot() bool`

HasProgressSnapshot returns a boolean if a field has been set.

### GetCloseMeta

`func (o *IntentSession) GetCloseMeta() map[string]interface{}`

GetCloseMeta returns the CloseMeta field if non-nil, zero value otherwise.

### GetCloseMetaOk

`func (o *IntentSession) GetCloseMetaOk() (*map[string]interface{}, bool)`

GetCloseMetaOk returns a tuple with the CloseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseMeta

`func (o *IntentSession) SetCloseMeta(v map[string]interface{})`

SetCloseMeta sets CloseMeta field to given value.

### HasCloseMeta

`func (o *IntentSession) HasCloseMeta() bool`

HasCloseMeta returns a boolean if a field has been set.

### GetSettlement

`func (o *IntentSession) GetSettlement() IntentSessionSettlement`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *IntentSession) GetSettlementOk() (*IntentSessionSettlement, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *IntentSession) SetSettlement(v IntentSessionSettlement)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *IntentSession) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### GetStartedAt

`func (o *IntentSession) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IntentSession) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IntentSession) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IntentSession) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetLastActivityAt

`func (o *IntentSession) GetLastActivityAt() time.Time`

GetLastActivityAt returns the LastActivityAt field if non-nil, zero value otherwise.

### GetLastActivityAtOk

`func (o *IntentSession) GetLastActivityAtOk() (*time.Time, bool)`

GetLastActivityAtOk returns a tuple with the LastActivityAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityAt

`func (o *IntentSession) SetLastActivityAt(v time.Time)`

SetLastActivityAt sets LastActivityAt field to given value.

### HasLastActivityAt

`func (o *IntentSession) HasLastActivityAt() bool`

HasLastActivityAt returns a boolean if a field has been set.

### GetEffectiveStatus

`func (o *IntentSession) GetEffectiveStatus() string`

GetEffectiveStatus returns the EffectiveStatus field if non-nil, zero value otherwise.

### GetEffectiveStatusOk

`func (o *IntentSession) GetEffectiveStatusOk() (*string, bool)`

GetEffectiveStatusOk returns a tuple with the EffectiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStatus

`func (o *IntentSession) SetEffectiveStatus(v string)`

SetEffectiveStatus sets EffectiveStatus field to given value.

### HasEffectiveStatus

`func (o *IntentSession) HasEffectiveStatus() bool`

HasEffectiveStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


