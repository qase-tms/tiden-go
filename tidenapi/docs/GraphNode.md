# GraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CoverageStatus** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | kind distinguishes node types in the graph (shift-left v3). \&quot;requirement\&quot; (default) or \&quot;component\&quot; — a component node reached via an impacts_component edge. Component nodes carry their name in &#x60;title&#x60;; seq_num/status/parent_id are empty. | [optional] 

## Methods

### NewGraphNode

`func NewGraphNode() *GraphNode`

NewGraphNode instantiates a new GraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeWithDefaults

`func NewGraphNodeWithDefaults() *GraphNode`

NewGraphNodeWithDefaults instantiates a new GraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *GraphNode) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *GraphNode) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *GraphNode) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *GraphNode) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *GraphNode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GraphNode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GraphNode) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GraphNode) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *GraphNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GraphNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GraphNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GraphNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCoverageStatus

`func (o *GraphNode) GetCoverageStatus() string`

GetCoverageStatus returns the CoverageStatus field if non-nil, zero value otherwise.

### GetCoverageStatusOk

`func (o *GraphNode) GetCoverageStatusOk() (*string, bool)`

GetCoverageStatusOk returns a tuple with the CoverageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageStatus

`func (o *GraphNode) SetCoverageStatus(v string)`

SetCoverageStatus sets CoverageStatus field to given value.

### HasCoverageStatus

`func (o *GraphNode) HasCoverageStatus() bool`

HasCoverageStatus returns a boolean if a field has been set.

### GetParentId

`func (o *GraphNode) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GraphNode) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GraphNode) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GraphNode) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetKind

`func (o *GraphNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GraphNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GraphNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GraphNode) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


