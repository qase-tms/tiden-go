# Requirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ChildrenCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**BranchStatus** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceCount** | Pointer to **int32** |  | [optional] 
**Sources** | Pointer to [**[]RequirementSource**](RequirementSource.md) |  | [optional] 

## Methods

### NewRequirement

`func NewRequirement() *Requirement`

NewRequirement instantiates a new Requirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementWithDefaults

`func NewRequirementWithDefaults() *Requirement`

NewRequirementWithDefaults instantiates a new Requirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Requirement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Requirement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Requirement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Requirement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Requirement) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Requirement) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Requirement) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Requirement) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetParentId

`func (o *Requirement) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Requirement) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Requirement) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Requirement) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetComponentId

`func (o *Requirement) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *Requirement) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *Requirement) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *Requirement) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetTitle

`func (o *Requirement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Requirement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Requirement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Requirement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *Requirement) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Requirement) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Requirement) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Requirement) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPosition

`func (o *Requirement) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Requirement) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Requirement) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Requirement) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Requirement) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Requirement) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Requirement) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Requirement) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetChildrenCount

`func (o *Requirement) GetChildrenCount() int32`

GetChildrenCount returns the ChildrenCount field if non-nil, zero value otherwise.

### GetChildrenCountOk

`func (o *Requirement) GetChildrenCountOk() (*int32, bool)`

GetChildrenCountOk returns a tuple with the ChildrenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenCount

`func (o *Requirement) SetChildrenCount(v int32)`

SetChildrenCount sets ChildrenCount field to given value.

### HasChildrenCount

`func (o *Requirement) HasChildrenCount() bool`

HasChildrenCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Requirement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Requirement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Requirement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Requirement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Requirement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Requirement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Requirement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Requirement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSeqNum

`func (o *Requirement) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *Requirement) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *Requirement) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *Requirement) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetStatus

`func (o *Requirement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Requirement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Requirement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Requirement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *Requirement) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Requirement) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Requirement) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Requirement) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssigneeId

`func (o *Requirement) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *Requirement) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *Requirement) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *Requirement) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetBranchId

`func (o *Requirement) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *Requirement) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *Requirement) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *Requirement) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetSourceId

`func (o *Requirement) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Requirement) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Requirement) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Requirement) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetBranchStatus

`func (o *Requirement) GetBranchStatus() string`

GetBranchStatus returns the BranchStatus field if non-nil, zero value otherwise.

### GetBranchStatusOk

`func (o *Requirement) GetBranchStatusOk() (*string, bool)`

GetBranchStatusOk returns a tuple with the BranchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchStatus

`func (o *Requirement) SetBranchStatus(v string)`

SetBranchStatus sets BranchStatus field to given value.

### HasBranchStatus

`func (o *Requirement) HasBranchStatus() bool`

HasBranchStatus returns a boolean if a field has been set.

### GetType

`func (o *Requirement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Requirement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Requirement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Requirement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceCount

`func (o *Requirement) GetSourceCount() int32`

GetSourceCount returns the SourceCount field if non-nil, zero value otherwise.

### GetSourceCountOk

`func (o *Requirement) GetSourceCountOk() (*int32, bool)`

GetSourceCountOk returns a tuple with the SourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCount

`func (o *Requirement) SetSourceCount(v int32)`

SetSourceCount sets SourceCount field to given value.

### HasSourceCount

`func (o *Requirement) HasSourceCount() bool`

HasSourceCount returns a boolean if a field has been set.

### GetSources

`func (o *Requirement) GetSources() []RequirementSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Requirement) GetSourcesOk() (*[]RequirementSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Requirement) SetSources(v []RequirementSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Requirement) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


