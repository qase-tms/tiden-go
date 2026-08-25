# RequirementRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to [**RequirementTier**](RequirementTier.md) |  | [optional] [default to REQUIREMENT_TIER_UNSPECIFIED]
**Resources** | Pointer to **[]string** | resources are the root&#39;s branch-effective repo_file anchor paths, same meaning as TouchedNode.resources (a root can carry the session&#39;s evidence). | [optional] 

## Methods

### NewRequirementRef

`func NewRequirementRef() *RequirementRef`

NewRequirementRef instantiates a new RequirementRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementRefWithDefaults

`func NewRequirementRefWithDefaults() *RequirementRef`

NewRequirementRefWithDefaults instantiates a new RequirementRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequirementRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequirementRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequirementRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequirementRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *RequirementRef) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *RequirementRef) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *RequirementRef) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *RequirementRef) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *RequirementRef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequirementRef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequirementRef) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RequirementRef) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RequirementRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequirementRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequirementRef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequirementRef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTier

`func (o *RequirementRef) GetTier() RequirementTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RequirementRef) GetTierOk() (*RequirementTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RequirementRef) SetTier(v RequirementTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *RequirementRef) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetResources

`func (o *RequirementRef) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RequirementRef) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RequirementRef) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RequirementRef) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


