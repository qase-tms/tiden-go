# BranchRequirementLinkProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**RequirementId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ReviewedBy** | Pointer to **string** |  | [optional] 
**ReviewedAt** | Pointer to **time.Time** |  | [optional] 
**ReviewNote** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBranchRequirementLinkProposal

`func NewBranchRequirementLinkProposal() *BranchRequirementLinkProposal`

NewBranchRequirementLinkProposal instantiates a new BranchRequirementLinkProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchRequirementLinkProposalWithDefaults

`func NewBranchRequirementLinkProposalWithDefaults() *BranchRequirementLinkProposal`

NewBranchRequirementLinkProposalWithDefaults instantiates a new BranchRequirementLinkProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BranchRequirementLinkProposal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BranchRequirementLinkProposal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BranchRequirementLinkProposal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BranchRequirementLinkProposal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBranchId

`func (o *BranchRequirementLinkProposal) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *BranchRequirementLinkProposal) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *BranchRequirementLinkProposal) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *BranchRequirementLinkProposal) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetTestId

`func (o *BranchRequirementLinkProposal) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *BranchRequirementLinkProposal) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *BranchRequirementLinkProposal) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *BranchRequirementLinkProposal) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetRequirementId

`func (o *BranchRequirementLinkProposal) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *BranchRequirementLinkProposal) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *BranchRequirementLinkProposal) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *BranchRequirementLinkProposal) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetStatus

`func (o *BranchRequirementLinkProposal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BranchRequirementLinkProposal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BranchRequirementLinkProposal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BranchRequirementLinkProposal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BranchRequirementLinkProposal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BranchRequirementLinkProposal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BranchRequirementLinkProposal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BranchRequirementLinkProposal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetReviewedBy

`func (o *BranchRequirementLinkProposal) GetReviewedBy() string`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *BranchRequirementLinkProposal) GetReviewedByOk() (*string, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *BranchRequirementLinkProposal) SetReviewedBy(v string)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *BranchRequirementLinkProposal) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### GetReviewedAt

`func (o *BranchRequirementLinkProposal) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *BranchRequirementLinkProposal) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *BranchRequirementLinkProposal) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.

### HasReviewedAt

`func (o *BranchRequirementLinkProposal) HasReviewedAt() bool`

HasReviewedAt returns a boolean if a field has been set.

### GetReviewNote

`func (o *BranchRequirementLinkProposal) GetReviewNote() string`

GetReviewNote returns the ReviewNote field if non-nil, zero value otherwise.

### GetReviewNoteOk

`func (o *BranchRequirementLinkProposal) GetReviewNoteOk() (*string, bool)`

GetReviewNoteOk returns a tuple with the ReviewNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewNote

`func (o *BranchRequirementLinkProposal) SetReviewNote(v string)`

SetReviewNote sets ReviewNote field to given value.

### HasReviewNote

`func (o *BranchRequirementLinkProposal) HasReviewNote() bool`

HasReviewNote returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BranchRequirementLinkProposal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BranchRequirementLinkProposal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BranchRequirementLinkProposal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BranchRequirementLinkProposal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BranchRequirementLinkProposal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BranchRequirementLinkProposal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BranchRequirementLinkProposal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BranchRequirementLinkProposal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


