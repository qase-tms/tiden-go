# RequirementSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**RequirementId** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Locator** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**RepoPath** | Pointer to **string** |  | [optional] 
**LineStart** | Pointer to **int32** |  | [optional] 
**LineEnd** | Pointer to **int32** |  | [optional] 
**Excerpt** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByAgentRunId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRequirementSource

`func NewRequirementSource() *RequirementSource`

NewRequirementSource instantiates a new RequirementSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementSourceWithDefaults

`func NewRequirementSourceWithDefaults() *RequirementSource`

NewRequirementSourceWithDefaults instantiates a new RequirementSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequirementSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequirementSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequirementSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequirementSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *RequirementSource) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *RequirementSource) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *RequirementSource) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *RequirementSource) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetBranchId

`func (o *RequirementSource) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *RequirementSource) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *RequirementSource) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *RequirementSource) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetRequirementId

`func (o *RequirementSource) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *RequirementSource) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *RequirementSource) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *RequirementSource) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetSourceType

`func (o *RequirementSource) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RequirementSource) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RequirementSource) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *RequirementSource) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *RequirementSource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequirementSource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequirementSource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequirementSource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLocator

`func (o *RequirementSource) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *RequirementSource) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *RequirementSource) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *RequirementSource) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetUrl

`func (o *RequirementSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequirementSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequirementSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequirementSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRepoPath

`func (o *RequirementSource) GetRepoPath() string`

GetRepoPath returns the RepoPath field if non-nil, zero value otherwise.

### GetRepoPathOk

`func (o *RequirementSource) GetRepoPathOk() (*string, bool)`

GetRepoPathOk returns a tuple with the RepoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoPath

`func (o *RequirementSource) SetRepoPath(v string)`

SetRepoPath sets RepoPath field to given value.

### HasRepoPath

`func (o *RequirementSource) HasRepoPath() bool`

HasRepoPath returns a boolean if a field has been set.

### GetLineStart

`func (o *RequirementSource) GetLineStart() int32`

GetLineStart returns the LineStart field if non-nil, zero value otherwise.

### GetLineStartOk

`func (o *RequirementSource) GetLineStartOk() (*int32, bool)`

GetLineStartOk returns a tuple with the LineStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStart

`func (o *RequirementSource) SetLineStart(v int32)`

SetLineStart sets LineStart field to given value.

### HasLineStart

`func (o *RequirementSource) HasLineStart() bool`

HasLineStart returns a boolean if a field has been set.

### GetLineEnd

`func (o *RequirementSource) GetLineEnd() int32`

GetLineEnd returns the LineEnd field if non-nil, zero value otherwise.

### GetLineEndOk

`func (o *RequirementSource) GetLineEndOk() (*int32, bool)`

GetLineEndOk returns a tuple with the LineEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnd

`func (o *RequirementSource) SetLineEnd(v int32)`

SetLineEnd sets LineEnd field to given value.

### HasLineEnd

`func (o *RequirementSource) HasLineEnd() bool`

HasLineEnd returns a boolean if a field has been set.

### GetExcerpt

`func (o *RequirementSource) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *RequirementSource) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *RequirementSource) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *RequirementSource) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetMetadata

`func (o *RequirementSource) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequirementSource) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequirementSource) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequirementSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RequirementSource) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequirementSource) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequirementSource) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RequirementSource) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByAgentRunId

`func (o *RequirementSource) GetCreatedByAgentRunId() string`

GetCreatedByAgentRunId returns the CreatedByAgentRunId field if non-nil, zero value otherwise.

### GetCreatedByAgentRunIdOk

`func (o *RequirementSource) GetCreatedByAgentRunIdOk() (*string, bool)`

GetCreatedByAgentRunIdOk returns a tuple with the CreatedByAgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentRunId

`func (o *RequirementSource) SetCreatedByAgentRunId(v string)`

SetCreatedByAgentRunId sets CreatedByAgentRunId field to given value.

### HasCreatedByAgentRunId

`func (o *RequirementSource) HasCreatedByAgentRunId() bool`

HasCreatedByAgentRunId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequirementSource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequirementSource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequirementSource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequirementSource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RequirementSource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RequirementSource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RequirementSource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RequirementSource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


