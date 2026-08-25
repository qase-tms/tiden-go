# ImpactedRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** |  | [optional] 
**Hops** | Pointer to **int32** | hops is the shortest number of graph edges from any seed. 0 &#x3D; direct anchor. | [optional] 
**ViaEdgeType** | Pointer to **string** | via_edge_type is the edge type that first reached this requirement at &#x60;hops&#x60; (\&quot;parent\&quot;, \&quot;covers\&quot;, \&quot;co_anchored\&quot;, …). Empty when hops &#x3D; 0. | [optional] 
**ViaConfidence** | Pointer to **float64** | via_confidence is that edge&#39;s confidence, when the edge carries one. Derived co_anchored/covers confidence is 1/fan-out; parent edges have none. | [optional] 
**AnchorPaths** | Pointer to **[]string** | anchor_paths lists the requested repo paths anchored to this requirement. Populated only when hops &#x3D; 0. | [optional] 

## Methods

### NewImpactedRequirement

`func NewImpactedRequirement() *ImpactedRequirement`

NewImpactedRequirement instantiates a new ImpactedRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpactedRequirementWithDefaults

`func NewImpactedRequirementWithDefaults() *ImpactedRequirement`

NewImpactedRequirementWithDefaults instantiates a new ImpactedRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *ImpactedRequirement) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *ImpactedRequirement) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *ImpactedRequirement) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *ImpactedRequirement) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetHops

`func (o *ImpactedRequirement) GetHops() int32`

GetHops returns the Hops field if non-nil, zero value otherwise.

### GetHopsOk

`func (o *ImpactedRequirement) GetHopsOk() (*int32, bool)`

GetHopsOk returns a tuple with the Hops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHops

`func (o *ImpactedRequirement) SetHops(v int32)`

SetHops sets Hops field to given value.

### HasHops

`func (o *ImpactedRequirement) HasHops() bool`

HasHops returns a boolean if a field has been set.

### GetViaEdgeType

`func (o *ImpactedRequirement) GetViaEdgeType() string`

GetViaEdgeType returns the ViaEdgeType field if non-nil, zero value otherwise.

### GetViaEdgeTypeOk

`func (o *ImpactedRequirement) GetViaEdgeTypeOk() (*string, bool)`

GetViaEdgeTypeOk returns a tuple with the ViaEdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaEdgeType

`func (o *ImpactedRequirement) SetViaEdgeType(v string)`

SetViaEdgeType sets ViaEdgeType field to given value.

### HasViaEdgeType

`func (o *ImpactedRequirement) HasViaEdgeType() bool`

HasViaEdgeType returns a boolean if a field has been set.

### GetViaConfidence

`func (o *ImpactedRequirement) GetViaConfidence() float64`

GetViaConfidence returns the ViaConfidence field if non-nil, zero value otherwise.

### GetViaConfidenceOk

`func (o *ImpactedRequirement) GetViaConfidenceOk() (*float64, bool)`

GetViaConfidenceOk returns a tuple with the ViaConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaConfidence

`func (o *ImpactedRequirement) SetViaConfidence(v float64)`

SetViaConfidence sets ViaConfidence field to given value.

### HasViaConfidence

`func (o *ImpactedRequirement) HasViaConfidence() bool`

HasViaConfidence returns a boolean if a field has been set.

### GetAnchorPaths

`func (o *ImpactedRequirement) GetAnchorPaths() []string`

GetAnchorPaths returns the AnchorPaths field if non-nil, zero value otherwise.

### GetAnchorPathsOk

`func (o *ImpactedRequirement) GetAnchorPathsOk() (*[]string, bool)`

GetAnchorPathsOk returns a tuple with the AnchorPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorPaths

`func (o *ImpactedRequirement) SetAnchorPaths(v []string)`

SetAnchorPaths sets AnchorPaths field to given value.

### HasAnchorPaths

`func (o *ImpactedRequirement) HasAnchorPaths() bool`

HasAnchorPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


