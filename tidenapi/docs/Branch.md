# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByAgentRunId** | Pointer to **string** | Set when this branch was produced by an agent run; NULL for branches created by humans through the UI / CLI. | [optional] 
**Stats** | Pointer to [**BranchChangeStats**](BranchChangeStats.md) |  | [optional] 
**CreatedByAgent** | Pointer to **string** | Coding agent that created this branch (validated server-side against a fixed allowlist at write time); empty for human-created branches or an unrecognized value. | [optional] 
**CreatedByName** | Pointer to **string** | Display name/email of the user in created_by, resolved server-side by ListBranches/GetBranch in one batched lookup. Empty when created_by is unset or the user has since been deleted. | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**MergedAt** | Pointer to **time.Time** | Immutable merge-completion timestamp (column exists since migration 000076); unset for open branches and for branches merged before that column existed. Distinct from updated_at, which mutates on any later edit — this is the one trustworthy \&quot;when did this land\&quot; fact. | [optional] 
**Loop** | Pointer to [**BranchLoopStats**](BranchLoopStats.md) |  | [optional] 
**LatestRun** | Pointer to [**BranchLatestRun**](BranchLatestRun.md) |  | [optional] 
**CodeLinks** | Pointer to [**[]CodeLink**](CodeLink.md) |  | [optional] 
**Intent** | Pointer to [**BranchIntentState**](BranchIntentState.md) |  | [optional] 

## Methods

### NewBranch

`func NewBranch() *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Branch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Branch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Branch) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Branch) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Branch) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Branch) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Branch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Branch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Branch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Branch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Branch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Branch) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Branch) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Branch) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Branch) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetStatus

`func (o *Branch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Branch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Branch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Branch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Branch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Branch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Branch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Branch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Branch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Branch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Branch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Branch) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedByAgentRunId

`func (o *Branch) GetCreatedByAgentRunId() string`

GetCreatedByAgentRunId returns the CreatedByAgentRunId field if non-nil, zero value otherwise.

### GetCreatedByAgentRunIdOk

`func (o *Branch) GetCreatedByAgentRunIdOk() (*string, bool)`

GetCreatedByAgentRunIdOk returns a tuple with the CreatedByAgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentRunId

`func (o *Branch) SetCreatedByAgentRunId(v string)`

SetCreatedByAgentRunId sets CreatedByAgentRunId field to given value.

### HasCreatedByAgentRunId

`func (o *Branch) HasCreatedByAgentRunId() bool`

HasCreatedByAgentRunId returns a boolean if a field has been set.

### GetStats

`func (o *Branch) GetStats() BranchChangeStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Branch) GetStatsOk() (*BranchChangeStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Branch) SetStats(v BranchChangeStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Branch) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreatedByAgent

`func (o *Branch) GetCreatedByAgent() string`

GetCreatedByAgent returns the CreatedByAgent field if non-nil, zero value otherwise.

### GetCreatedByAgentOk

`func (o *Branch) GetCreatedByAgentOk() (*string, bool)`

GetCreatedByAgentOk returns a tuple with the CreatedByAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgent

`func (o *Branch) SetCreatedByAgent(v string)`

SetCreatedByAgent sets CreatedByAgent field to given value.

### HasCreatedByAgent

`func (o *Branch) HasCreatedByAgent() bool`

HasCreatedByAgent returns a boolean if a field has been set.

### GetCreatedByName

`func (o *Branch) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *Branch) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *Branch) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *Branch) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *Branch) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *Branch) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *Branch) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *Branch) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetMergedAt

`func (o *Branch) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *Branch) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *Branch) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.

### HasMergedAt

`func (o *Branch) HasMergedAt() bool`

HasMergedAt returns a boolean if a field has been set.

### GetLoop

`func (o *Branch) GetLoop() BranchLoopStats`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *Branch) GetLoopOk() (*BranchLoopStats, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *Branch) SetLoop(v BranchLoopStats)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *Branch) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetLatestRun

`func (o *Branch) GetLatestRun() BranchLatestRun`

GetLatestRun returns the LatestRun field if non-nil, zero value otherwise.

### GetLatestRunOk

`func (o *Branch) GetLatestRunOk() (*BranchLatestRun, bool)`

GetLatestRunOk returns a tuple with the LatestRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRun

`func (o *Branch) SetLatestRun(v BranchLatestRun)`

SetLatestRun sets LatestRun field to given value.

### HasLatestRun

`func (o *Branch) HasLatestRun() bool`

HasLatestRun returns a boolean if a field has been set.

### GetCodeLinks

`func (o *Branch) GetCodeLinks() []CodeLink`

GetCodeLinks returns the CodeLinks field if non-nil, zero value otherwise.

### GetCodeLinksOk

`func (o *Branch) GetCodeLinksOk() (*[]CodeLink, bool)`

GetCodeLinksOk returns a tuple with the CodeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLinks

`func (o *Branch) SetCodeLinks(v []CodeLink)`

SetCodeLinks sets CodeLinks field to given value.

### HasCodeLinks

`func (o *Branch) HasCodeLinks() bool`

HasCodeLinks returns a boolean if a field has been set.

### GetIntent

`func (o *Branch) GetIntent() BranchIntentState`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *Branch) GetIntentOk() (*BranchIntentState, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *Branch) SetIntent(v BranchIntentState)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *Branch) HasIntent() bool`

HasIntent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


