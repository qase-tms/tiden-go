# FeatureContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Root** | Pointer to [**RequirementRef**](RequirementRef.md) |  | [optional] 
**TouchedNodes** | Pointer to [**[]TouchedNode**](TouchedNode.md) |  | [optional] 
**Coverage** | Pointer to [**Coverage**](Coverage.md) |  | [optional] 
**Via** | Pointer to **[]string** |  | [optional] 
**Tier** | Pointer to [**RequirementTier**](RequirementTier.md) |  | [optional] [default to REQUIREMENT_TIER_UNSPECIFIED]

## Methods

### NewFeatureContext

`func NewFeatureContext() *FeatureContext`

NewFeatureContext instantiates a new FeatureContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureContextWithDefaults

`func NewFeatureContextWithDefaults() *FeatureContext`

NewFeatureContextWithDefaults instantiates a new FeatureContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoot

`func (o *FeatureContext) GetRoot() RequirementRef`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *FeatureContext) GetRootOk() (*RequirementRef, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *FeatureContext) SetRoot(v RequirementRef)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *FeatureContext) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetTouchedNodes

`func (o *FeatureContext) GetTouchedNodes() []TouchedNode`

GetTouchedNodes returns the TouchedNodes field if non-nil, zero value otherwise.

### GetTouchedNodesOk

`func (o *FeatureContext) GetTouchedNodesOk() (*[]TouchedNode, bool)`

GetTouchedNodesOk returns a tuple with the TouchedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouchedNodes

`func (o *FeatureContext) SetTouchedNodes(v []TouchedNode)`

SetTouchedNodes sets TouchedNodes field to given value.

### HasTouchedNodes

`func (o *FeatureContext) HasTouchedNodes() bool`

HasTouchedNodes returns a boolean if a field has been set.

### GetCoverage

`func (o *FeatureContext) GetCoverage() Coverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *FeatureContext) GetCoverageOk() (*Coverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *FeatureContext) SetCoverage(v Coverage)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *FeatureContext) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetVia

`func (o *FeatureContext) GetVia() []string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *FeatureContext) GetViaOk() (*[]string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *FeatureContext) SetVia(v []string)`

SetVia sets Via field to given value.

### HasVia

`func (o *FeatureContext) HasVia() bool`

HasVia returns a boolean if a field has been set.

### GetTier

`func (o *FeatureContext) GetTier() RequirementTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *FeatureContext) GetTierOk() (*RequirementTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *FeatureContext) SetTier(v RequirementTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *FeatureContext) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


