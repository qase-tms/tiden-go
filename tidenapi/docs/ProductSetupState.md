# ProductSetupState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**RepoFingerprint** | Pointer to **string** |  | [optional] 
**RepoProductId** | Pointer to **string** |  | [optional] 
**RepoBound** | Pointer to **bool** |  | [optional] 
**GitHookWired** | Pointer to **bool** |  | [optional] 
**Agents** | Pointer to [**[]ProductSetupAgentStatus**](ProductSetupAgentStatus.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProductSetupState

`func NewProductSetupState() *ProductSetupState`

NewProductSetupState instantiates a new ProductSetupState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSetupStateWithDefaults

`func NewProductSetupStateWithDefaults() *ProductSetupState`

NewProductSetupStateWithDefaults instantiates a new ProductSetupState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductSetupState) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductSetupState) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductSetupState) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductSetupState) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *ProductSetupState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProductSetupState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProductSetupState) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ProductSetupState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRepoFingerprint

`func (o *ProductSetupState) GetRepoFingerprint() string`

GetRepoFingerprint returns the RepoFingerprint field if non-nil, zero value otherwise.

### GetRepoFingerprintOk

`func (o *ProductSetupState) GetRepoFingerprintOk() (*string, bool)`

GetRepoFingerprintOk returns a tuple with the RepoFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoFingerprint

`func (o *ProductSetupState) SetRepoFingerprint(v string)`

SetRepoFingerprint sets RepoFingerprint field to given value.

### HasRepoFingerprint

`func (o *ProductSetupState) HasRepoFingerprint() bool`

HasRepoFingerprint returns a boolean if a field has been set.

### GetRepoProductId

`func (o *ProductSetupState) GetRepoProductId() string`

GetRepoProductId returns the RepoProductId field if non-nil, zero value otherwise.

### GetRepoProductIdOk

`func (o *ProductSetupState) GetRepoProductIdOk() (*string, bool)`

GetRepoProductIdOk returns a tuple with the RepoProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoProductId

`func (o *ProductSetupState) SetRepoProductId(v string)`

SetRepoProductId sets RepoProductId field to given value.

### HasRepoProductId

`func (o *ProductSetupState) HasRepoProductId() bool`

HasRepoProductId returns a boolean if a field has been set.

### GetRepoBound

`func (o *ProductSetupState) GetRepoBound() bool`

GetRepoBound returns the RepoBound field if non-nil, zero value otherwise.

### GetRepoBoundOk

`func (o *ProductSetupState) GetRepoBoundOk() (*bool, bool)`

GetRepoBoundOk returns a tuple with the RepoBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoBound

`func (o *ProductSetupState) SetRepoBound(v bool)`

SetRepoBound sets RepoBound field to given value.

### HasRepoBound

`func (o *ProductSetupState) HasRepoBound() bool`

HasRepoBound returns a boolean if a field has been set.

### GetGitHookWired

`func (o *ProductSetupState) GetGitHookWired() bool`

GetGitHookWired returns the GitHookWired field if non-nil, zero value otherwise.

### GetGitHookWiredOk

`func (o *ProductSetupState) GetGitHookWiredOk() (*bool, bool)`

GetGitHookWiredOk returns a tuple with the GitHookWired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHookWired

`func (o *ProductSetupState) SetGitHookWired(v bool)`

SetGitHookWired sets GitHookWired field to given value.

### HasGitHookWired

`func (o *ProductSetupState) HasGitHookWired() bool`

HasGitHookWired returns a boolean if a field has been set.

### GetAgents

`func (o *ProductSetupState) GetAgents() []ProductSetupAgentStatus`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *ProductSetupState) GetAgentsOk() (*[]ProductSetupAgentStatus, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *ProductSetupState) SetAgents(v []ProductSetupAgentStatus)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *ProductSetupState) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetSource

`func (o *ProductSetupState) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductSetupState) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductSetupState) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProductSetupState) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *ProductSetupState) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ProductSetupState) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ProductSetupState) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ProductSetupState) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductSetupState) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductSetupState) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductSetupState) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductSetupState) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProductSetupState) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductSetupState) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductSetupState) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductSetupState) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


