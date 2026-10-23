# CreateTestRunBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to **map[string]string** |  | [optional] 
**BuildSha** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**ClientMeta** | Pointer to **map[string]string** |  | [optional] 
**IntentSessionId** | Pointer to **string** |  | [optional] 
**IntentBranch** | Pointer to **string** | The session&#39;s Tiden intent branch: live-documentation sync lands there instead of resolving the free-text git &#x60;branch&#x60; name (which never creates a Tiden branch). Set by an in-session &#x60;tiden run exec&#x60;; empty otherwise. | [optional] 

## Methods

### NewCreateTestRunBody

`func NewCreateTestRunBody() *CreateTestRunBody`

NewCreateTestRunBody instantiates a new CreateTestRunBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunBodyWithDefaults

`func NewCreateTestRunBodyWithDefaults() *CreateTestRunBody`

NewCreateTestRunBodyWithDefaults instantiates a new CreateTestRunBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateTestRunBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTestRunBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTestRunBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateTestRunBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTestRunBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateTestRunBody) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateTestRunBody) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateTestRunBody) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateTestRunBody) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetBranch

`func (o *CreateTestRunBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CreateTestRunBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CreateTestRunBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CreateTestRunBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConfigurations

`func (o *CreateTestRunBody) GetConfigurations() map[string]string`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *CreateTestRunBody) GetConfigurationsOk() (*map[string]string, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *CreateTestRunBody) SetConfigurations(v map[string]string)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *CreateTestRunBody) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetBuildSha

`func (o *CreateTestRunBody) GetBuildSha() string`

GetBuildSha returns the BuildSha field if non-nil, zero value otherwise.

### GetBuildShaOk

`func (o *CreateTestRunBody) GetBuildShaOk() (*string, bool)`

GetBuildShaOk returns a tuple with the BuildSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSha

`func (o *CreateTestRunBody) SetBuildSha(v string)`

SetBuildSha sets BuildSha field to given value.

### HasBuildSha

`func (o *CreateTestRunBody) HasBuildSha() bool`

HasBuildSha returns a boolean if a field has been set.

### GetStartedAt

`func (o *CreateTestRunBody) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CreateTestRunBody) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CreateTestRunBody) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CreateTestRunBody) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetClientMeta

`func (o *CreateTestRunBody) GetClientMeta() map[string]string`

GetClientMeta returns the ClientMeta field if non-nil, zero value otherwise.

### GetClientMetaOk

`func (o *CreateTestRunBody) GetClientMetaOk() (*map[string]string, bool)`

GetClientMetaOk returns a tuple with the ClientMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMeta

`func (o *CreateTestRunBody) SetClientMeta(v map[string]string)`

SetClientMeta sets ClientMeta field to given value.

### HasClientMeta

`func (o *CreateTestRunBody) HasClientMeta() bool`

HasClientMeta returns a boolean if a field has been set.

### GetIntentSessionId

`func (o *CreateTestRunBody) GetIntentSessionId() string`

GetIntentSessionId returns the IntentSessionId field if non-nil, zero value otherwise.

### GetIntentSessionIdOk

`func (o *CreateTestRunBody) GetIntentSessionIdOk() (*string, bool)`

GetIntentSessionIdOk returns a tuple with the IntentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentSessionId

`func (o *CreateTestRunBody) SetIntentSessionId(v string)`

SetIntentSessionId sets IntentSessionId field to given value.

### HasIntentSessionId

`func (o *CreateTestRunBody) HasIntentSessionId() bool`

HasIntentSessionId returns a boolean if a field has been set.

### GetIntentBranch

`func (o *CreateTestRunBody) GetIntentBranch() string`

GetIntentBranch returns the IntentBranch field if non-nil, zero value otherwise.

### GetIntentBranchOk

`func (o *CreateTestRunBody) GetIntentBranchOk() (*string, bool)`

GetIntentBranchOk returns a tuple with the IntentBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentBranch

`func (o *CreateTestRunBody) SetIntentBranch(v string)`

SetIntentBranch sets IntentBranch field to given value.

### HasIntentBranch

`func (o *CreateTestRunBody) HasIntentBranch() bool`

HasIntentBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


