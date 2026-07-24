# GetRequirementGraphResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]GraphNode**](GraphNode.md) |  | [optional] 
**Edges** | Pointer to [**[]GraphEdge**](GraphEdge.md) |  | [optional] 

## Methods

### NewGetRequirementGraphResponse

`func NewGetRequirementGraphResponse() *GetRequirementGraphResponse`

NewGetRequirementGraphResponse instantiates a new GetRequirementGraphResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRequirementGraphResponseWithDefaults

`func NewGetRequirementGraphResponseWithDefaults() *GetRequirementGraphResponse`

NewGetRequirementGraphResponseWithDefaults instantiates a new GetRequirementGraphResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *GetRequirementGraphResponse) GetNodes() []GraphNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetRequirementGraphResponse) GetNodesOk() (*[]GraphNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetRequirementGraphResponse) SetNodes(v []GraphNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GetRequirementGraphResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *GetRequirementGraphResponse) GetEdges() []GraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GetRequirementGraphResponse) GetEdgesOk() (*[]GraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GetRequirementGraphResponse) SetEdges(v []GraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GetRequirementGraphResponse) HasEdges() bool`

HasEdges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


