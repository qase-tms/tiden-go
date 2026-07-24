# RequirementSourceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Locator** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**RepoPath** | Pointer to **string** |  | [optional] 
**LineStart** | Pointer to **int32** |  | [optional] 
**LineEnd** | Pointer to **int32** |  | [optional] 
**Excerpt** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedByAgentRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewRequirementSourceInput

`func NewRequirementSourceInput() *RequirementSourceInput`

NewRequirementSourceInput instantiates a new RequirementSourceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementSourceInputWithDefaults

`func NewRequirementSourceInputWithDefaults() *RequirementSourceInput`

NewRequirementSourceInputWithDefaults instantiates a new RequirementSourceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *RequirementSourceInput) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RequirementSourceInput) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RequirementSourceInput) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *RequirementSourceInput) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *RequirementSourceInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequirementSourceInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequirementSourceInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequirementSourceInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLocator

`func (o *RequirementSourceInput) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *RequirementSourceInput) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *RequirementSourceInput) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *RequirementSourceInput) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetUrl

`func (o *RequirementSourceInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequirementSourceInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequirementSourceInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequirementSourceInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRepoPath

`func (o *RequirementSourceInput) GetRepoPath() string`

GetRepoPath returns the RepoPath field if non-nil, zero value otherwise.

### GetRepoPathOk

`func (o *RequirementSourceInput) GetRepoPathOk() (*string, bool)`

GetRepoPathOk returns a tuple with the RepoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoPath

`func (o *RequirementSourceInput) SetRepoPath(v string)`

SetRepoPath sets RepoPath field to given value.

### HasRepoPath

`func (o *RequirementSourceInput) HasRepoPath() bool`

HasRepoPath returns a boolean if a field has been set.

### GetLineStart

`func (o *RequirementSourceInput) GetLineStart() int32`

GetLineStart returns the LineStart field if non-nil, zero value otherwise.

### GetLineStartOk

`func (o *RequirementSourceInput) GetLineStartOk() (*int32, bool)`

GetLineStartOk returns a tuple with the LineStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStart

`func (o *RequirementSourceInput) SetLineStart(v int32)`

SetLineStart sets LineStart field to given value.

### HasLineStart

`func (o *RequirementSourceInput) HasLineStart() bool`

HasLineStart returns a boolean if a field has been set.

### GetLineEnd

`func (o *RequirementSourceInput) GetLineEnd() int32`

GetLineEnd returns the LineEnd field if non-nil, zero value otherwise.

### GetLineEndOk

`func (o *RequirementSourceInput) GetLineEndOk() (*int32, bool)`

GetLineEndOk returns a tuple with the LineEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnd

`func (o *RequirementSourceInput) SetLineEnd(v int32)`

SetLineEnd sets LineEnd field to given value.

### HasLineEnd

`func (o *RequirementSourceInput) HasLineEnd() bool`

HasLineEnd returns a boolean if a field has been set.

### GetExcerpt

`func (o *RequirementSourceInput) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *RequirementSourceInput) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *RequirementSourceInput) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *RequirementSourceInput) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetMetadata

`func (o *RequirementSourceInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequirementSourceInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequirementSourceInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequirementSourceInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedByAgentRunId

`func (o *RequirementSourceInput) GetCreatedByAgentRunId() string`

GetCreatedByAgentRunId returns the CreatedByAgentRunId field if non-nil, zero value otherwise.

### GetCreatedByAgentRunIdOk

`func (o *RequirementSourceInput) GetCreatedByAgentRunIdOk() (*string, bool)`

GetCreatedByAgentRunIdOk returns a tuple with the CreatedByAgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentRunId

`func (o *RequirementSourceInput) SetCreatedByAgentRunId(v string)`

SetCreatedByAgentRunId sets CreatedByAgentRunId field to given value.

### HasCreatedByAgentRunId

`func (o *RequirementSourceInput) HasCreatedByAgentRunId() bool`

HasCreatedByAgentRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


