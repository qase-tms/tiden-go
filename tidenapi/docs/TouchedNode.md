# TouchedNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Repository** | Pointer to **string** | repository is the canonical repo id of this node&#39;s component, or \&quot;\&quot; when unscoped. Lets the CLI verdict compare changed files vs resources per-repo (avoids cross-repo path collisions). resources are repo-relative. | [optional] 
**Tier** | Pointer to [**RequirementTier**](RequirementTier.md) |  | [optional] [default to REQUIREMENT_TIER_UNSPECIFIED]

## Methods

### NewTouchedNode

`func NewTouchedNode() *TouchedNode`

NewTouchedNode instantiates a new TouchedNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTouchedNodeWithDefaults

`func NewTouchedNodeWithDefaults() *TouchedNode`

NewTouchedNodeWithDefaults instantiates a new TouchedNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TouchedNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TouchedNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TouchedNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TouchedNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *TouchedNode) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *TouchedNode) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *TouchedNode) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *TouchedNode) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *TouchedNode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TouchedNode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TouchedNode) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TouchedNode) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *TouchedNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TouchedNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TouchedNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TouchedNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *TouchedNode) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TouchedNode) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TouchedNode) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TouchedNode) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRepository

`func (o *TouchedNode) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TouchedNode) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TouchedNode) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TouchedNode) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTier

`func (o *TouchedNode) GetTier() RequirementTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TouchedNode) GetTierOk() (*RequirementTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TouchedNode) SetTier(v RequirementTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TouchedNode) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


