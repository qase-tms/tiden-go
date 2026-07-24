# RequirementSourcesUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]RequirementSourceInput**](RequirementSourceInput.md) |  | [optional] 
**Merge** | Pointer to **bool** | merge, when true, unions incoming sources with existing ones using anchor-key dedup instead of replacing the whole set. Agent writes set merge&#x3D;true; UI edits leave it false (default) for explicit replace semantics. | [optional] 

## Methods

### NewRequirementSourcesUpdate

`func NewRequirementSourcesUpdate() *RequirementSourcesUpdate`

NewRequirementSourcesUpdate instantiates a new RequirementSourcesUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementSourcesUpdateWithDefaults

`func NewRequirementSourcesUpdateWithDefaults() *RequirementSourcesUpdate`

NewRequirementSourcesUpdateWithDefaults instantiates a new RequirementSourcesUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *RequirementSourcesUpdate) GetSources() []RequirementSourceInput`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RequirementSourcesUpdate) GetSourcesOk() (*[]RequirementSourceInput, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RequirementSourcesUpdate) SetSources(v []RequirementSourceInput)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RequirementSourcesUpdate) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetMerge

`func (o *RequirementSourcesUpdate) GetMerge() bool`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *RequirementSourcesUpdate) GetMergeOk() (*bool, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *RequirementSourcesUpdate) SetMerge(v bool)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *RequirementSourcesUpdate) HasMerge() bool`

HasMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


