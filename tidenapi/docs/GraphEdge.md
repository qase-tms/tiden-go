# GraphEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | Pointer to **string** |  | [optional] 
**Dst** | Pointer to **string** |  | [optional] 
**EdgeType** | Pointer to **string** |  | [optional] 
**SourceKind** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float64** |  | [optional] 

## Methods

### NewGraphEdge

`func NewGraphEdge() *GraphEdge`

NewGraphEdge instantiates a new GraphEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphEdgeWithDefaults

`func NewGraphEdgeWithDefaults() *GraphEdge`

NewGraphEdgeWithDefaults instantiates a new GraphEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *GraphEdge) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *GraphEdge) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *GraphEdge) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *GraphEdge) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetDst

`func (o *GraphEdge) GetDst() string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *GraphEdge) GetDstOk() (*string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *GraphEdge) SetDst(v string)`

SetDst sets Dst field to given value.

### HasDst

`func (o *GraphEdge) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetEdgeType

`func (o *GraphEdge) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *GraphEdge) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *GraphEdge) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *GraphEdge) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### GetSourceKind

`func (o *GraphEdge) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *GraphEdge) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *GraphEdge) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *GraphEdge) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.

### GetConfidence

`func (o *GraphEdge) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GraphEdge) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GraphEdge) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GraphEdge) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


