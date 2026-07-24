# DeriveTestLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLinked** | Pointer to [**[]FileAnchorCandidate**](FileAnchorCandidate.md) | Exact-file matches auto-linked durably (Quality Gate recomputes). | [optional] 
**Candidates** | Pointer to [**[]FileAnchorCandidate**](FileAnchorCandidate.md) | Directory-proximity matches for an agent to confirm + LinkRequirement. | [optional] 
**MultiRepoSkipped** | Pointer to **bool** | Set when the product spans &gt;1 repository: matching is skipped (tests carry no repo attribution, so a bare path match could cross-link across repos). | [optional] 

## Methods

### NewDeriveTestLinksResponse

`func NewDeriveTestLinksResponse() *DeriveTestLinksResponse`

NewDeriveTestLinksResponse instantiates a new DeriveTestLinksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeriveTestLinksResponseWithDefaults

`func NewDeriveTestLinksResponseWithDefaults() *DeriveTestLinksResponse`

NewDeriveTestLinksResponseWithDefaults instantiates a new DeriveTestLinksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLinked

`func (o *DeriveTestLinksResponse) GetAutoLinked() []FileAnchorCandidate`

GetAutoLinked returns the AutoLinked field if non-nil, zero value otherwise.

### GetAutoLinkedOk

`func (o *DeriveTestLinksResponse) GetAutoLinkedOk() (*[]FileAnchorCandidate, bool)`

GetAutoLinkedOk returns a tuple with the AutoLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLinked

`func (o *DeriveTestLinksResponse) SetAutoLinked(v []FileAnchorCandidate)`

SetAutoLinked sets AutoLinked field to given value.

### HasAutoLinked

`func (o *DeriveTestLinksResponse) HasAutoLinked() bool`

HasAutoLinked returns a boolean if a field has been set.

### GetCandidates

`func (o *DeriveTestLinksResponse) GetCandidates() []FileAnchorCandidate`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *DeriveTestLinksResponse) GetCandidatesOk() (*[]FileAnchorCandidate, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *DeriveTestLinksResponse) SetCandidates(v []FileAnchorCandidate)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *DeriveTestLinksResponse) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetMultiRepoSkipped

`func (o *DeriveTestLinksResponse) GetMultiRepoSkipped() bool`

GetMultiRepoSkipped returns the MultiRepoSkipped field if non-nil, zero value otherwise.

### GetMultiRepoSkippedOk

`func (o *DeriveTestLinksResponse) GetMultiRepoSkippedOk() (*bool, bool)`

GetMultiRepoSkippedOk returns a tuple with the MultiRepoSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRepoSkipped

`func (o *DeriveTestLinksResponse) SetMultiRepoSkipped(v bool)`

SetMultiRepoSkipped sets MultiRepoSkipped field to given value.

### HasMultiRepoSkipped

`func (o *DeriveTestLinksResponse) HasMultiRepoSkipped() bool`

HasMultiRepoSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


