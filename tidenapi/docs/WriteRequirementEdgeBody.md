# WriteRequirementEdgeBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcRequirementId** | Pointer to **string** |  | [optional] 
**DstRequirementId** | Pointer to **string** | Exactly one dst endpoint: dst_requirement_id (req→req, edge_type depends_on/traces_to) OR dst_component_id (req→component, edge_type impacts_component). dst_requirement_id may be empty when dst_component_id is set. | [optional] 
**EdgeType** | Pointer to **string** | edge_type must be \&quot;depends_on\&quot;/\&quot;traces_to\&quot; (req→req) or \&quot;impacts_component\&quot; (req→component). | [optional] 
**Confidence** | Pointer to **float64** | confidence must be provided; explicit 0.0 is valid, omitted is not. | [optional] 
**AgentRunId** | Pointer to **string** | agent_run_id is attributed to the written edge. | [optional] 
**DstComponentId** | Pointer to **string** | dst_component_id sets a req→component endpoint (shift-left v3). Mutually exclusive with dst_requirement_id. | [optional] 

## Methods

### NewWriteRequirementEdgeBody

`func NewWriteRequirementEdgeBody() *WriteRequirementEdgeBody`

NewWriteRequirementEdgeBody instantiates a new WriteRequirementEdgeBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequirementEdgeBodyWithDefaults

`func NewWriteRequirementEdgeBodyWithDefaults() *WriteRequirementEdgeBody`

NewWriteRequirementEdgeBodyWithDefaults instantiates a new WriteRequirementEdgeBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrcRequirementId

`func (o *WriteRequirementEdgeBody) GetSrcRequirementId() string`

GetSrcRequirementId returns the SrcRequirementId field if non-nil, zero value otherwise.

### GetSrcRequirementIdOk

`func (o *WriteRequirementEdgeBody) GetSrcRequirementIdOk() (*string, bool)`

GetSrcRequirementIdOk returns a tuple with the SrcRequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcRequirementId

`func (o *WriteRequirementEdgeBody) SetSrcRequirementId(v string)`

SetSrcRequirementId sets SrcRequirementId field to given value.

### HasSrcRequirementId

`func (o *WriteRequirementEdgeBody) HasSrcRequirementId() bool`

HasSrcRequirementId returns a boolean if a field has been set.

### GetDstRequirementId

`func (o *WriteRequirementEdgeBody) GetDstRequirementId() string`

GetDstRequirementId returns the DstRequirementId field if non-nil, zero value otherwise.

### GetDstRequirementIdOk

`func (o *WriteRequirementEdgeBody) GetDstRequirementIdOk() (*string, bool)`

GetDstRequirementIdOk returns a tuple with the DstRequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstRequirementId

`func (o *WriteRequirementEdgeBody) SetDstRequirementId(v string)`

SetDstRequirementId sets DstRequirementId field to given value.

### HasDstRequirementId

`func (o *WriteRequirementEdgeBody) HasDstRequirementId() bool`

HasDstRequirementId returns a boolean if a field has been set.

### GetEdgeType

`func (o *WriteRequirementEdgeBody) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *WriteRequirementEdgeBody) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *WriteRequirementEdgeBody) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *WriteRequirementEdgeBody) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### GetConfidence

`func (o *WriteRequirementEdgeBody) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *WriteRequirementEdgeBody) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *WriteRequirementEdgeBody) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *WriteRequirementEdgeBody) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAgentRunId

`func (o *WriteRequirementEdgeBody) GetAgentRunId() string`

GetAgentRunId returns the AgentRunId field if non-nil, zero value otherwise.

### GetAgentRunIdOk

`func (o *WriteRequirementEdgeBody) GetAgentRunIdOk() (*string, bool)`

GetAgentRunIdOk returns a tuple with the AgentRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRunId

`func (o *WriteRequirementEdgeBody) SetAgentRunId(v string)`

SetAgentRunId sets AgentRunId field to given value.

### HasAgentRunId

`func (o *WriteRequirementEdgeBody) HasAgentRunId() bool`

HasAgentRunId returns a boolean if a field has been set.

### GetDstComponentId

`func (o *WriteRequirementEdgeBody) GetDstComponentId() string`

GetDstComponentId returns the DstComponentId field if non-nil, zero value otherwise.

### GetDstComponentIdOk

`func (o *WriteRequirementEdgeBody) GetDstComponentIdOk() (*string, bool)`

GetDstComponentIdOk returns a tuple with the DstComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstComponentId

`func (o *WriteRequirementEdgeBody) SetDstComponentId(v string)`

SetDstComponentId sets DstComponentId field to given value.

### HasDstComponentId

`func (o *WriteRequirementEdgeBody) HasDstComponentId() bool`

HasDstComponentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


