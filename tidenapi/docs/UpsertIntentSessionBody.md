# UpsertIntentSessionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**Objective** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Verdict** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Executor** | Pointer to **string** |  | [optional] 
**GitRepo** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**StartHead** | Pointer to **string** |  | [optional] 
**RetrievedIds** | Pointer to [**RetrievedRequirementIds**](RetrievedRequirementIds.md) |  | [optional] 
**RadiusDrops** | Pointer to [**[]RadiusDrop**](RadiusDrop.md) |  | [optional] 
**ClosedReason** | Pointer to **string** | closed_reason is settable ONLY in a request that also sets (or the row already has) status &#x3D; closed — rejected otherwise. explicit | merged | never-materialized. | [optional] 
**ScopeBand** | Pointer to **string** |  | [optional] 
**ProgressSnapshot** | Pointer to **map[string]interface{}** | progress_snapshot / close_meta are opaque JSON passthrough (the caller&#39;s own shape; this service does not interpret them) — same transport as the IntentSession message&#39;s identically-named fields above. | [optional] 
**CloseMeta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpsertIntentSessionBody

`func NewUpsertIntentSessionBody() *UpsertIntentSessionBody`

NewUpsertIntentSessionBody instantiates a new UpsertIntentSessionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertIntentSessionBodyWithDefaults

`func NewUpsertIntentSessionBodyWithDefaults() *UpsertIntentSessionBody`

NewUpsertIntentSessionBodyWithDefaults instantiates a new UpsertIntentSessionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *UpsertIntentSessionBody) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UpsertIntentSessionBody) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UpsertIntentSessionBody) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *UpsertIntentSessionBody) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetObjective

`func (o *UpsertIntentSessionBody) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *UpsertIntentSessionBody) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *UpsertIntentSessionBody) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *UpsertIntentSessionBody) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetStatus

`func (o *UpsertIntentSessionBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpsertIntentSessionBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpsertIntentSessionBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpsertIntentSessionBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerdict

`func (o *UpsertIntentSessionBody) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *UpsertIntentSessionBody) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *UpsertIntentSessionBody) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *UpsertIntentSessionBody) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetBranchId

`func (o *UpsertIntentSessionBody) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *UpsertIntentSessionBody) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *UpsertIntentSessionBody) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *UpsertIntentSessionBody) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetAgent

`func (o *UpsertIntentSessionBody) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *UpsertIntentSessionBody) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *UpsertIntentSessionBody) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *UpsertIntentSessionBody) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExecutor

`func (o *UpsertIntentSessionBody) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *UpsertIntentSessionBody) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *UpsertIntentSessionBody) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *UpsertIntentSessionBody) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetGitRepo

`func (o *UpsertIntentSessionBody) GetGitRepo() string`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *UpsertIntentSessionBody) GetGitRepoOk() (*string, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *UpsertIntentSessionBody) SetGitRepo(v string)`

SetGitRepo sets GitRepo field to given value.

### HasGitRepo

`func (o *UpsertIntentSessionBody) HasGitRepo() bool`

HasGitRepo returns a boolean if a field has been set.

### GetGitBranch

`func (o *UpsertIntentSessionBody) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *UpsertIntentSessionBody) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *UpsertIntentSessionBody) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *UpsertIntentSessionBody) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetStartHead

`func (o *UpsertIntentSessionBody) GetStartHead() string`

GetStartHead returns the StartHead field if non-nil, zero value otherwise.

### GetStartHeadOk

`func (o *UpsertIntentSessionBody) GetStartHeadOk() (*string, bool)`

GetStartHeadOk returns a tuple with the StartHead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHead

`func (o *UpsertIntentSessionBody) SetStartHead(v string)`

SetStartHead sets StartHead field to given value.

### HasStartHead

`func (o *UpsertIntentSessionBody) HasStartHead() bool`

HasStartHead returns a boolean if a field has been set.

### GetRetrievedIds

`func (o *UpsertIntentSessionBody) GetRetrievedIds() RetrievedRequirementIds`

GetRetrievedIds returns the RetrievedIds field if non-nil, zero value otherwise.

### GetRetrievedIdsOk

`func (o *UpsertIntentSessionBody) GetRetrievedIdsOk() (*RetrievedRequirementIds, bool)`

GetRetrievedIdsOk returns a tuple with the RetrievedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedIds

`func (o *UpsertIntentSessionBody) SetRetrievedIds(v RetrievedRequirementIds)`

SetRetrievedIds sets RetrievedIds field to given value.

### HasRetrievedIds

`func (o *UpsertIntentSessionBody) HasRetrievedIds() bool`

HasRetrievedIds returns a boolean if a field has been set.

### GetRadiusDrops

`func (o *UpsertIntentSessionBody) GetRadiusDrops() []RadiusDrop`

GetRadiusDrops returns the RadiusDrops field if non-nil, zero value otherwise.

### GetRadiusDropsOk

`func (o *UpsertIntentSessionBody) GetRadiusDropsOk() (*[]RadiusDrop, bool)`

GetRadiusDropsOk returns a tuple with the RadiusDrops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusDrops

`func (o *UpsertIntentSessionBody) SetRadiusDrops(v []RadiusDrop)`

SetRadiusDrops sets RadiusDrops field to given value.

### HasRadiusDrops

`func (o *UpsertIntentSessionBody) HasRadiusDrops() bool`

HasRadiusDrops returns a boolean if a field has been set.

### GetClosedReason

`func (o *UpsertIntentSessionBody) GetClosedReason() string`

GetClosedReason returns the ClosedReason field if non-nil, zero value otherwise.

### GetClosedReasonOk

`func (o *UpsertIntentSessionBody) GetClosedReasonOk() (*string, bool)`

GetClosedReasonOk returns a tuple with the ClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedReason

`func (o *UpsertIntentSessionBody) SetClosedReason(v string)`

SetClosedReason sets ClosedReason field to given value.

### HasClosedReason

`func (o *UpsertIntentSessionBody) HasClosedReason() bool`

HasClosedReason returns a boolean if a field has been set.

### GetScopeBand

`func (o *UpsertIntentSessionBody) GetScopeBand() string`

GetScopeBand returns the ScopeBand field if non-nil, zero value otherwise.

### GetScopeBandOk

`func (o *UpsertIntentSessionBody) GetScopeBandOk() (*string, bool)`

GetScopeBandOk returns a tuple with the ScopeBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeBand

`func (o *UpsertIntentSessionBody) SetScopeBand(v string)`

SetScopeBand sets ScopeBand field to given value.

### HasScopeBand

`func (o *UpsertIntentSessionBody) HasScopeBand() bool`

HasScopeBand returns a boolean if a field has been set.

### GetProgressSnapshot

`func (o *UpsertIntentSessionBody) GetProgressSnapshot() map[string]interface{}`

GetProgressSnapshot returns the ProgressSnapshot field if non-nil, zero value otherwise.

### GetProgressSnapshotOk

`func (o *UpsertIntentSessionBody) GetProgressSnapshotOk() (*map[string]interface{}, bool)`

GetProgressSnapshotOk returns a tuple with the ProgressSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressSnapshot

`func (o *UpsertIntentSessionBody) SetProgressSnapshot(v map[string]interface{})`

SetProgressSnapshot sets ProgressSnapshot field to given value.

### HasProgressSnapshot

`func (o *UpsertIntentSessionBody) HasProgressSnapshot() bool`

HasProgressSnapshot returns a boolean if a field has been set.

### GetCloseMeta

`func (o *UpsertIntentSessionBody) GetCloseMeta() map[string]interface{}`

GetCloseMeta returns the CloseMeta field if non-nil, zero value otherwise.

### GetCloseMetaOk

`func (o *UpsertIntentSessionBody) GetCloseMetaOk() (*map[string]interface{}, bool)`

GetCloseMetaOk returns a tuple with the CloseMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseMeta

`func (o *UpsertIntentSessionBody) SetCloseMeta(v map[string]interface{})`

SetCloseMeta sets CloseMeta field to given value.

### HasCloseMeta

`func (o *UpsertIntentSessionBody) HasCloseMeta() bool`

HasCloseMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


