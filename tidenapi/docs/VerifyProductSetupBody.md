# VerifyProductSetupBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoFingerprint** | Pointer to **string** |  | [optional] 
**RepoProductId** | Pointer to **string** |  | [optional] 
**RepoBound** | Pointer to **bool** |  | [optional] 
**GitHookWired** | Pointer to **bool** |  | [optional] 
**Agents** | Pointer to [**[]ProductSetupAgentStatus**](ProductSetupAgentStatus.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyProductSetupBody

`func NewVerifyProductSetupBody() *VerifyProductSetupBody`

NewVerifyProductSetupBody instantiates a new VerifyProductSetupBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyProductSetupBodyWithDefaults

`func NewVerifyProductSetupBodyWithDefaults() *VerifyProductSetupBody`

NewVerifyProductSetupBodyWithDefaults instantiates a new VerifyProductSetupBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoFingerprint

`func (o *VerifyProductSetupBody) GetRepoFingerprint() string`

GetRepoFingerprint returns the RepoFingerprint field if non-nil, zero value otherwise.

### GetRepoFingerprintOk

`func (o *VerifyProductSetupBody) GetRepoFingerprintOk() (*string, bool)`

GetRepoFingerprintOk returns a tuple with the RepoFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoFingerprint

`func (o *VerifyProductSetupBody) SetRepoFingerprint(v string)`

SetRepoFingerprint sets RepoFingerprint field to given value.

### HasRepoFingerprint

`func (o *VerifyProductSetupBody) HasRepoFingerprint() bool`

HasRepoFingerprint returns a boolean if a field has been set.

### GetRepoProductId

`func (o *VerifyProductSetupBody) GetRepoProductId() string`

GetRepoProductId returns the RepoProductId field if non-nil, zero value otherwise.

### GetRepoProductIdOk

`func (o *VerifyProductSetupBody) GetRepoProductIdOk() (*string, bool)`

GetRepoProductIdOk returns a tuple with the RepoProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoProductId

`func (o *VerifyProductSetupBody) SetRepoProductId(v string)`

SetRepoProductId sets RepoProductId field to given value.

### HasRepoProductId

`func (o *VerifyProductSetupBody) HasRepoProductId() bool`

HasRepoProductId returns a boolean if a field has been set.

### GetRepoBound

`func (o *VerifyProductSetupBody) GetRepoBound() bool`

GetRepoBound returns the RepoBound field if non-nil, zero value otherwise.

### GetRepoBoundOk

`func (o *VerifyProductSetupBody) GetRepoBoundOk() (*bool, bool)`

GetRepoBoundOk returns a tuple with the RepoBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoBound

`func (o *VerifyProductSetupBody) SetRepoBound(v bool)`

SetRepoBound sets RepoBound field to given value.

### HasRepoBound

`func (o *VerifyProductSetupBody) HasRepoBound() bool`

HasRepoBound returns a boolean if a field has been set.

### GetGitHookWired

`func (o *VerifyProductSetupBody) GetGitHookWired() bool`

GetGitHookWired returns the GitHookWired field if non-nil, zero value otherwise.

### GetGitHookWiredOk

`func (o *VerifyProductSetupBody) GetGitHookWiredOk() (*bool, bool)`

GetGitHookWiredOk returns a tuple with the GitHookWired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHookWired

`func (o *VerifyProductSetupBody) SetGitHookWired(v bool)`

SetGitHookWired sets GitHookWired field to given value.

### HasGitHookWired

`func (o *VerifyProductSetupBody) HasGitHookWired() bool`

HasGitHookWired returns a boolean if a field has been set.

### GetAgents

`func (o *VerifyProductSetupBody) GetAgents() []ProductSetupAgentStatus`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *VerifyProductSetupBody) GetAgentsOk() (*[]ProductSetupAgentStatus, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *VerifyProductSetupBody) SetAgents(v []ProductSetupAgentStatus)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *VerifyProductSetupBody) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetSource

`func (o *VerifyProductSetupBody) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VerifyProductSetupBody) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VerifyProductSetupBody) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VerifyProductSetupBody) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


