# ImpactCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedPaths** | Pointer to **int32** | requested_paths is how many repo paths the caller sent. | [optional] 
**MatchedPaths** | Pointer to **int32** | matched_paths is how many of them are anchored to at least one requirement. 0 with product_has_anchors &#x3D; true means \&quot;we cannot tell\&quot;, not \&quot;nothing\&quot;. | [optional] 
**ProductHasAnchors** | Pointer to **bool** | product_has_anchors is false when the product has no repo_file anchors at all — an empty answer then says nothing about the change. | [optional] 
**UnmatchedPaths** | Pointer to **[]string** | unmatched_paths is the requested paths that matched no anchor (capped). | [optional] 
**BlindDirectories** | Pointer to **[]string** | blind_directories is the directories among unmatched_paths under which the product has no anchor at all — i.e. tiers we are structurally blind to. | [optional] 
**UnverifiedRepositorySeeds** | Pointer to **int32** | unverified_repository_seeds counts seeds kept without a repository check because their requirement has no component. Only set when request.repository is non-empty; it should fall to 0 as components get assigned. | [optional] 

## Methods

### NewImpactCoverage

`func NewImpactCoverage() *ImpactCoverage`

NewImpactCoverage instantiates a new ImpactCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactCoverageWithDefaults

`func NewImpactCoverageWithDefaults() *ImpactCoverage`

NewImpactCoverageWithDefaults instantiates a new ImpactCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedPaths

`func (o *ImpactCoverage) GetRequestedPaths() int32`

GetRequestedPaths returns the RequestedPaths field if non-nil, zero value otherwise.

### GetRequestedPathsOk

`func (o *ImpactCoverage) GetRequestedPathsOk() (*int32, bool)`

GetRequestedPathsOk returns a tuple with the RequestedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPaths

`func (o *ImpactCoverage) SetRequestedPaths(v int32)`

SetRequestedPaths sets RequestedPaths field to given value.

### HasRequestedPaths

`func (o *ImpactCoverage) HasRequestedPaths() bool`

HasRequestedPaths returns a boolean if a field has been set.

### GetMatchedPaths

`func (o *ImpactCoverage) GetMatchedPaths() int32`

GetMatchedPaths returns the MatchedPaths field if non-nil, zero value otherwise.

### GetMatchedPathsOk

`func (o *ImpactCoverage) GetMatchedPathsOk() (*int32, bool)`

GetMatchedPathsOk returns a tuple with the MatchedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPaths

`func (o *ImpactCoverage) SetMatchedPaths(v int32)`

SetMatchedPaths sets MatchedPaths field to given value.

### HasMatchedPaths

`func (o *ImpactCoverage) HasMatchedPaths() bool`

HasMatchedPaths returns a boolean if a field has been set.

### GetProductHasAnchors

`func (o *ImpactCoverage) GetProductHasAnchors() bool`

GetProductHasAnchors returns the ProductHasAnchors field if non-nil, zero value otherwise.

### GetProductHasAnchorsOk

`func (o *ImpactCoverage) GetProductHasAnchorsOk() (*bool, bool)`

GetProductHasAnchorsOk returns a tuple with the ProductHasAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductHasAnchors

`func (o *ImpactCoverage) SetProductHasAnchors(v bool)`

SetProductHasAnchors sets ProductHasAnchors field to given value.

### HasProductHasAnchors

`func (o *ImpactCoverage) HasProductHasAnchors() bool`

HasProductHasAnchors returns a boolean if a field has been set.

### GetUnmatchedPaths

`func (o *ImpactCoverage) GetUnmatchedPaths() []string`

GetUnmatchedPaths returns the UnmatchedPaths field if non-nil, zero value otherwise.

### GetUnmatchedPathsOk

`func (o *ImpactCoverage) GetUnmatchedPathsOk() (*[]string, bool)`

GetUnmatchedPathsOk returns a tuple with the UnmatchedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedPaths

`func (o *ImpactCoverage) SetUnmatchedPaths(v []string)`

SetUnmatchedPaths sets UnmatchedPaths field to given value.

### HasUnmatchedPaths

`func (o *ImpactCoverage) HasUnmatchedPaths() bool`

HasUnmatchedPaths returns a boolean if a field has been set.

### GetBlindDirectories

`func (o *ImpactCoverage) GetBlindDirectories() []string`

GetBlindDirectories returns the BlindDirectories field if non-nil, zero value otherwise.

### GetBlindDirectoriesOk

`func (o *ImpactCoverage) GetBlindDirectoriesOk() (*[]string, bool)`

GetBlindDirectoriesOk returns a tuple with the BlindDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlindDirectories

`func (o *ImpactCoverage) SetBlindDirectories(v []string)`

SetBlindDirectories sets BlindDirectories field to given value.

### HasBlindDirectories

`func (o *ImpactCoverage) HasBlindDirectories() bool`

HasBlindDirectories returns a boolean if a field has been set.

### GetUnverifiedRepositorySeeds

`func (o *ImpactCoverage) GetUnverifiedRepositorySeeds() int32`

GetUnverifiedRepositorySeeds returns the UnverifiedRepositorySeeds field if non-nil, zero value otherwise.

### GetUnverifiedRepositorySeedsOk

`func (o *ImpactCoverage) GetUnverifiedRepositorySeedsOk() (*int32, bool)`

GetUnverifiedRepositorySeedsOk returns a tuple with the UnverifiedRepositorySeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverifiedRepositorySeeds

`func (o *ImpactCoverage) SetUnverifiedRepositorySeeds(v int32)`

SetUnverifiedRepositorySeeds sets UnverifiedRepositorySeeds field to given value.

### HasUnverifiedRepositorySeeds

`func (o *ImpactCoverage) HasUnverifiedRepositorySeeds() bool`

HasUnverifiedRepositorySeeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


