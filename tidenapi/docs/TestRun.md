# TestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**EnvironmentSlug** | Pointer to **string** |  | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 
**BranchName** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to **map[string]string** |  | [optional] 
**BuildSha** | Pointer to **string** |  | [optional] 
**ClientMeta** | Pointer to **map[string]string** |  | [optional] 
**Stats** | Pointer to [**RunStats**](RunStats.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**LiveDocStatus** | Pointer to **string** |  | [optional] 
**LiveDocOperationId** | Pointer to **string** |  | [optional] 
**LiveDocStats** | Pointer to [**LiveDocStats**](LiveDocStats.md) |  | [optional] 
**LiveDocError** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTestRun

`func NewTestRun() *TestRun`

NewTestRun instantiates a new TestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunWithDefaults

`func NewTestRunWithDefaults() *TestRun`

NewTestRunWithDefaults instantiates a new TestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *TestRun) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *TestRun) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *TestRun) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *TestRun) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSeqNum

`func (o *TestRun) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *TestRun) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *TestRun) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *TestRun) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *TestRun) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestRun) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestRun) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestRun) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TestRun) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRun) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRun) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRun) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TestRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *TestRun) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *TestRun) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *TestRun) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *TestRun) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentSlug

`func (o *TestRun) GetEnvironmentSlug() string`

GetEnvironmentSlug returns the EnvironmentSlug field if non-nil, zero value otherwise.

### GetEnvironmentSlugOk

`func (o *TestRun) GetEnvironmentSlugOk() (*string, bool)`

GetEnvironmentSlugOk returns a tuple with the EnvironmentSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSlug

`func (o *TestRun) SetEnvironmentSlug(v string)`

SetEnvironmentSlug sets EnvironmentSlug field to given value.

### HasEnvironmentSlug

`func (o *TestRun) HasEnvironmentSlug() bool`

HasEnvironmentSlug returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *TestRun) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *TestRun) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *TestRun) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *TestRun) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetBranchName

`func (o *TestRun) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *TestRun) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *TestRun) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *TestRun) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetBranchId

`func (o *TestRun) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *TestRun) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *TestRun) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *TestRun) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetConfigurations

`func (o *TestRun) GetConfigurations() map[string]string`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *TestRun) GetConfigurationsOk() (*map[string]string, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *TestRun) SetConfigurations(v map[string]string)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *TestRun) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetBuildSha

`func (o *TestRun) GetBuildSha() string`

GetBuildSha returns the BuildSha field if non-nil, zero value otherwise.

### GetBuildShaOk

`func (o *TestRun) GetBuildShaOk() (*string, bool)`

GetBuildShaOk returns a tuple with the BuildSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSha

`func (o *TestRun) SetBuildSha(v string)`

SetBuildSha sets BuildSha field to given value.

### HasBuildSha

`func (o *TestRun) HasBuildSha() bool`

HasBuildSha returns a boolean if a field has been set.

### GetClientMeta

`func (o *TestRun) GetClientMeta() map[string]string`

GetClientMeta returns the ClientMeta field if non-nil, zero value otherwise.

### GetClientMetaOk

`func (o *TestRun) GetClientMetaOk() (*map[string]string, bool)`

GetClientMetaOk returns a tuple with the ClientMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMeta

`func (o *TestRun) SetClientMeta(v map[string]string)`

SetClientMeta sets ClientMeta field to given value.

### HasClientMeta

`func (o *TestRun) HasClientMeta() bool`

HasClientMeta returns a boolean if a field has been set.

### GetStats

`func (o *TestRun) GetStats() RunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *TestRun) GetStatsOk() (*RunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *TestRun) SetStats(v RunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *TestRun) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStartedAt

`func (o *TestRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TestRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TestRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TestRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *TestRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TestRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TestRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *TestRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetLiveDocStatus

`func (o *TestRun) GetLiveDocStatus() string`

GetLiveDocStatus returns the LiveDocStatus field if non-nil, zero value otherwise.

### GetLiveDocStatusOk

`func (o *TestRun) GetLiveDocStatusOk() (*string, bool)`

GetLiveDocStatusOk returns a tuple with the LiveDocStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDocStatus

`func (o *TestRun) SetLiveDocStatus(v string)`

SetLiveDocStatus sets LiveDocStatus field to given value.

### HasLiveDocStatus

`func (o *TestRun) HasLiveDocStatus() bool`

HasLiveDocStatus returns a boolean if a field has been set.

### GetLiveDocOperationId

`func (o *TestRun) GetLiveDocOperationId() string`

GetLiveDocOperationId returns the LiveDocOperationId field if non-nil, zero value otherwise.

### GetLiveDocOperationIdOk

`func (o *TestRun) GetLiveDocOperationIdOk() (*string, bool)`

GetLiveDocOperationIdOk returns a tuple with the LiveDocOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDocOperationId

`func (o *TestRun) SetLiveDocOperationId(v string)`

SetLiveDocOperationId sets LiveDocOperationId field to given value.

### HasLiveDocOperationId

`func (o *TestRun) HasLiveDocOperationId() bool`

HasLiveDocOperationId returns a boolean if a field has been set.

### GetLiveDocStats

`func (o *TestRun) GetLiveDocStats() LiveDocStats`

GetLiveDocStats returns the LiveDocStats field if non-nil, zero value otherwise.

### GetLiveDocStatsOk

`func (o *TestRun) GetLiveDocStatsOk() (*LiveDocStats, bool)`

GetLiveDocStatsOk returns a tuple with the LiveDocStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDocStats

`func (o *TestRun) SetLiveDocStats(v LiveDocStats)`

SetLiveDocStats sets LiveDocStats field to given value.

### HasLiveDocStats

`func (o *TestRun) HasLiveDocStats() bool`

HasLiveDocStats returns a boolean if a field has been set.

### GetLiveDocError

`func (o *TestRun) GetLiveDocError() string`

GetLiveDocError returns the LiveDocError field if non-nil, zero value otherwise.

### GetLiveDocErrorOk

`func (o *TestRun) GetLiveDocErrorOk() (*string, bool)`

GetLiveDocErrorOk returns a tuple with the LiveDocError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDocError

`func (o *TestRun) SetLiveDocError(v string)`

SetLiveDocError sets LiveDocError field to given value.

### HasLiveDocError

`func (o *TestRun) HasLiveDocError() bool`

HasLiveDocError returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TestRun) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TestRun) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TestRun) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TestRun) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TestRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TestRun) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TestRun) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TestRun) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TestRun) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


