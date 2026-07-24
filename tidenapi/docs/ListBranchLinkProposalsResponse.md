# ListBranchLinkProposalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proposals** | Pointer to [**[]BranchRequirementLinkProposal**](BranchRequirementLinkProposal.md) |  | [optional] 

## Methods

### NewListBranchLinkProposalsResponse

`func NewListBranchLinkProposalsResponse() *ListBranchLinkProposalsResponse`

NewListBranchLinkProposalsResponse instantiates a new ListBranchLinkProposalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBranchLinkProposalsResponseWithDefaults

`func NewListBranchLinkProposalsResponseWithDefaults() *ListBranchLinkProposalsResponse`

NewListBranchLinkProposalsResponseWithDefaults instantiates a new ListBranchLinkProposalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposals

`func (o *ListBranchLinkProposalsResponse) GetProposals() []BranchRequirementLinkProposal`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *ListBranchLinkProposalsResponse) GetProposalsOk() (*[]BranchRequirementLinkProposal, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *ListBranchLinkProposalsResponse) SetProposals(v []BranchRequirementLinkProposal)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *ListBranchLinkProposalsResponse) HasProposals() bool`

HasProposals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


