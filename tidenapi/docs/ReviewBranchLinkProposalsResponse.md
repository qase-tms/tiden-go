# ReviewBranchLinkProposalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proposals** | Pointer to [**[]BranchRequirementLinkProposal**](BranchRequirementLinkProposal.md) |  | [optional] 

## Methods

### NewReviewBranchLinkProposalsResponse

`func NewReviewBranchLinkProposalsResponse() *ReviewBranchLinkProposalsResponse`

NewReviewBranchLinkProposalsResponse instantiates a new ReviewBranchLinkProposalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewBranchLinkProposalsResponseWithDefaults

`func NewReviewBranchLinkProposalsResponseWithDefaults() *ReviewBranchLinkProposalsResponse`

NewReviewBranchLinkProposalsResponseWithDefaults instantiates a new ReviewBranchLinkProposalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposals

`func (o *ReviewBranchLinkProposalsResponse) GetProposals() []BranchRequirementLinkProposal`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *ReviewBranchLinkProposalsResponse) GetProposalsOk() (*[]BranchRequirementLinkProposal, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *ReviewBranchLinkProposalsResponse) SetProposals(v []BranchRequirementLinkProposal)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *ReviewBranchLinkProposalsResponse) HasProposals() bool`

HasProposals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


