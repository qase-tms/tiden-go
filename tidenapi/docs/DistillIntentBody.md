# DistillIntentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transcript** | Pointer to **string** | Rendered conversation (normalized, slim — user/assistant text only). The backend never sees an agent-specific transcript format. | [optional] 
**CredentialId** | Pointer to **string** | Optional LLM credential to use. When empty, the backend picks the first usable llm.* credential in the product&#39;s workspace. | [optional] 
**Model** | Pointer to **string** | Optional model override; when empty, falls back to the credential&#39;s metadata \&quot;model\&quot; field. | [optional] 
**Slug** | Pointer to **string** | Optional branch slug override; when empty, defaults to \&quot;session\&quot;. | [optional] 
**ChangedFiles** | Pointer to **[]string** | Repo-relative paths of files changed in the coding session (from git post-commit hook). The backend validates LLM-emitted repo_file anchors against this allowlist, dropping any path not in the set. Empty &#x3D; no anchors written (backward-compatible: older CLI sends nothing). | [optional] 
**SessionId** | Pointer to **string** | Coding-session identity. When either is set, every created/updated requirement gets a manual_input \&quot;Intent capture\&quot; provenance source carrying {sessionId, agent}, so a reviewer can trace it back to the session that produced it. Both empty (older CLI) writes no provenance source. | [optional] 
**Agent** | Pointer to **string** |  | [optional] 

## Methods

### NewDistillIntentBody

`func NewDistillIntentBody() *DistillIntentBody`

NewDistillIntentBody instantiates a new DistillIntentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistillIntentBodyWithDefaults

`func NewDistillIntentBodyWithDefaults() *DistillIntentBody`

NewDistillIntentBodyWithDefaults instantiates a new DistillIntentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscript

`func (o *DistillIntentBody) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *DistillIntentBody) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *DistillIntentBody) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *DistillIntentBody) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetCredentialId

`func (o *DistillIntentBody) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *DistillIntentBody) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *DistillIntentBody) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *DistillIntentBody) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetModel

`func (o *DistillIntentBody) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DistillIntentBody) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DistillIntentBody) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DistillIntentBody) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSlug

`func (o *DistillIntentBody) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DistillIntentBody) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DistillIntentBody) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DistillIntentBody) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetChangedFiles

`func (o *DistillIntentBody) GetChangedFiles() []string`

GetChangedFiles returns the ChangedFiles field if non-nil, zero value otherwise.

### GetChangedFilesOk

`func (o *DistillIntentBody) GetChangedFilesOk() (*[]string, bool)`

GetChangedFilesOk returns a tuple with the ChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedFiles

`func (o *DistillIntentBody) SetChangedFiles(v []string)`

SetChangedFiles sets ChangedFiles field to given value.

### HasChangedFiles

`func (o *DistillIntentBody) HasChangedFiles() bool`

HasChangedFiles returns a boolean if a field has been set.

### GetSessionId

`func (o *DistillIntentBody) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DistillIntentBody) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DistillIntentBody) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DistillIntentBody) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetAgent

`func (o *DistillIntentBody) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *DistillIntentBody) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *DistillIntentBody) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *DistillIntentBody) HasAgent() bool`

HasAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


