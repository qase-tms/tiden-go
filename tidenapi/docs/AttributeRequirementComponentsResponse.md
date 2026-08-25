# AttributeRequirementComponentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributed** | Pointer to [**[]AttributedRequirementComponent**](AttributedRequirementComponent.md) | Requirements newly attributed this run. Bounded (see maxReturnedAttributedRequirements in the service implementation) so a large product&#39;s backfill can&#39;t exceed the gRPC max message size; the DB write itself is unbounded — attributed_count is the true total. | [optional] 
**AttributedCount** | Pointer to **int32** |  | [optional] 
**SkippedMultiComponentCount** | Pointer to **int32** | Requirements whose repo_file anchors resolved to more than one distinct component — left untouched. | [optional] 
**SkippedUnownedCount** | Pointer to **int32** | Requirements whose repo_file anchors resolved to no component at all (including a repository-ambiguous anchor path) — left untouched. | [optional] 
**ConsideredCount** | Pointer to **int32** | Requirements evaluated: main-branch, non-deleted, component_id NULL, with at least one repo_file anchor. | [optional] 

## Methods

### NewAttributeRequirementComponentsResponse

`func NewAttributeRequirementComponentsResponse() *AttributeRequirementComponentsResponse`

NewAttributeRequirementComponentsResponse instantiates a new AttributeRequirementComponentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRequirementComponentsResponseWithDefaults

`func NewAttributeRequirementComponentsResponseWithDefaults() *AttributeRequirementComponentsResponse`

NewAttributeRequirementComponentsResponseWithDefaults instantiates a new AttributeRequirementComponentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributed

`func (o *AttributeRequirementComponentsResponse) GetAttributed() []AttributedRequirementComponent`

GetAttributed returns the Attributed field if non-nil, zero value otherwise.

### GetAttributedOk

`func (o *AttributeRequirementComponentsResponse) GetAttributedOk() (*[]AttributedRequirementComponent, bool)`

GetAttributedOk returns a tuple with the Attributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributed

`func (o *AttributeRequirementComponentsResponse) SetAttributed(v []AttributedRequirementComponent)`

SetAttributed sets Attributed field to given value.

### HasAttributed

`func (o *AttributeRequirementComponentsResponse) HasAttributed() bool`

HasAttributed returns a boolean if a field has been set.

### GetAttributedCount

`func (o *AttributeRequirementComponentsResponse) GetAttributedCount() int32`

GetAttributedCount returns the AttributedCount field if non-nil, zero value otherwise.

### GetAttributedCountOk

`func (o *AttributeRequirementComponentsResponse) GetAttributedCountOk() (*int32, bool)`

GetAttributedCountOk returns a tuple with the AttributedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributedCount

`func (o *AttributeRequirementComponentsResponse) SetAttributedCount(v int32)`

SetAttributedCount sets AttributedCount field to given value.

### HasAttributedCount

`func (o *AttributeRequirementComponentsResponse) HasAttributedCount() bool`

HasAttributedCount returns a boolean if a field has been set.

### GetSkippedMultiComponentCount

`func (o *AttributeRequirementComponentsResponse) GetSkippedMultiComponentCount() int32`

GetSkippedMultiComponentCount returns the SkippedMultiComponentCount field if non-nil, zero value otherwise.

### GetSkippedMultiComponentCountOk

`func (o *AttributeRequirementComponentsResponse) GetSkippedMultiComponentCountOk() (*int32, bool)`

GetSkippedMultiComponentCountOk returns a tuple with the SkippedMultiComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedMultiComponentCount

`func (o *AttributeRequirementComponentsResponse) SetSkippedMultiComponentCount(v int32)`

SetSkippedMultiComponentCount sets SkippedMultiComponentCount field to given value.

### HasSkippedMultiComponentCount

`func (o *AttributeRequirementComponentsResponse) HasSkippedMultiComponentCount() bool`

HasSkippedMultiComponentCount returns a boolean if a field has been set.

### GetSkippedUnownedCount

`func (o *AttributeRequirementComponentsResponse) GetSkippedUnownedCount() int32`

GetSkippedUnownedCount returns the SkippedUnownedCount field if non-nil, zero value otherwise.

### GetSkippedUnownedCountOk

`func (o *AttributeRequirementComponentsResponse) GetSkippedUnownedCountOk() (*int32, bool)`

GetSkippedUnownedCountOk returns a tuple with the SkippedUnownedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedUnownedCount

`func (o *AttributeRequirementComponentsResponse) SetSkippedUnownedCount(v int32)`

SetSkippedUnownedCount sets SkippedUnownedCount field to given value.

### HasSkippedUnownedCount

`func (o *AttributeRequirementComponentsResponse) HasSkippedUnownedCount() bool`

HasSkippedUnownedCount returns a boolean if a field has been set.

### GetConsideredCount

`func (o *AttributeRequirementComponentsResponse) GetConsideredCount() int32`

GetConsideredCount returns the ConsideredCount field if non-nil, zero value otherwise.

### GetConsideredCountOk

`func (o *AttributeRequirementComponentsResponse) GetConsideredCountOk() (*int32, bool)`

GetConsideredCountOk returns a tuple with the ConsideredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsideredCount

`func (o *AttributeRequirementComponentsResponse) SetConsideredCount(v int32)`

SetConsideredCount sets ConsideredCount field to given value.

### HasConsideredCount

`func (o *AttributeRequirementComponentsResponse) HasConsideredCount() bool`

HasConsideredCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


