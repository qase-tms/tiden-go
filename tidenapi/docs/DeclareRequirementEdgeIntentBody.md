# DeclareRequirementEdgeIntentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | branch is the branch the intent is declared on (empty &#x3D; main). | [optional] 
**SrcRequirementId** | Pointer to **string** |  | [optional] 
**DstRequirementId** | Pointer to **string** | Exactly one dst endpoint (see WriteRequirementEdgeRequest). dst_requirement_id may be empty when dst_component_id is set. | [optional] 
**EdgeType** | Pointer to **string** | edge_type must be \&quot;depends_on\&quot;/\&quot;traces_to\&quot; (req→req) or \&quot;impacts_component\&quot; (req→component). | [optional] 
**Confidence** | Pointer to **float64** | confidence must be provided; explicit 0.0 is valid, omitted is not. | [optional] 
**Rationale** | Pointer to **string** |  | [optional] 
**DstComponentId** | Pointer to **string** | dst_component_id sets a req→component endpoint (shift-left v3). Mutually exclusive with dst_requirement_id. | [optional] 

## Methods

### NewDeclareRequirementEdgeIntentBody

`func NewDeclareRequirementEdgeIntentBody() *DeclareRequirementEdgeIntentBody`

NewDeclareRequirementEdgeIntentBody instantiates a new DeclareRequirementEdgeIntentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclareRequirementEdgeIntentBodyWithDefaults

`func NewDeclareRequirementEdgeIntentBodyWithDefaults() *DeclareRequirementEdgeIntentBody`

NewDeclareRequirementEdgeIntentBodyWithDefaults instantiates a new DeclareRequirementEdgeIntentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *DeclareRequirementEdgeIntentBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DeclareRequirementEdgeIntentBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DeclareRequirementEdgeIntentBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DeclareRequirementEdgeIntentBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetSrcRequirementId

`func (o *DeclareRequirementEdgeIntentBody) GetSrcRequirementId() string`

GetSrcRequirementId returns the SrcRequirementId field if non-nil, zero value otherwise.

### GetSrcRequirementIdOk

`func (o *DeclareRequirementEdgeIntentBody) GetSrcRequirementIdOk() (*string, bool)`

GetSrcRequirementIdOk returns a tuple with the SrcRequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcRequirementId

`func (o *DeclareRequirementEdgeIntentBody) SetSrcRequirementId(v string)`

SetSrcRequirementId sets SrcRequirementId field to given value.

### HasSrcRequirementId

`func (o *DeclareRequirementEdgeIntentBody) HasSrcRequirementId() bool`

HasSrcRequirementId returns a boolean if a field has been set.

### GetDstRequirementId

`func (o *DeclareRequirementEdgeIntentBody) GetDstRequirementId() string`

GetDstRequirementId returns the DstRequirementId field if non-nil, zero value otherwise.

### GetDstRequirementIdOk

`func (o *DeclareRequirementEdgeIntentBody) GetDstRequirementIdOk() (*string, bool)`

GetDstRequirementIdOk returns a tuple with the DstRequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstRequirementId

`func (o *DeclareRequirementEdgeIntentBody) SetDstRequirementId(v string)`

SetDstRequirementId sets DstRequirementId field to given value.

### HasDstRequirementId

`func (o *DeclareRequirementEdgeIntentBody) HasDstRequirementId() bool`

HasDstRequirementId returns a boolean if a field has been set.

### GetEdgeType

`func (o *DeclareRequirementEdgeIntentBody) GetEdgeType() string`

GetEdgeType returns the EdgeType field if non-nil, zero value otherwise.

### GetEdgeTypeOk

`func (o *DeclareRequirementEdgeIntentBody) GetEdgeTypeOk() (*string, bool)`

GetEdgeTypeOk returns a tuple with the EdgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeType

`func (o *DeclareRequirementEdgeIntentBody) SetEdgeType(v string)`

SetEdgeType sets EdgeType field to given value.

### HasEdgeType

`func (o *DeclareRequirementEdgeIntentBody) HasEdgeType() bool`

HasEdgeType returns a boolean if a field has been set.

### GetConfidence

`func (o *DeclareRequirementEdgeIntentBody) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DeclareRequirementEdgeIntentBody) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DeclareRequirementEdgeIntentBody) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *DeclareRequirementEdgeIntentBody) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetRationale

`func (o *DeclareRequirementEdgeIntentBody) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *DeclareRequirementEdgeIntentBody) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *DeclareRequirementEdgeIntentBody) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *DeclareRequirementEdgeIntentBody) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetDstComponentId

`func (o *DeclareRequirementEdgeIntentBody) GetDstComponentId() string`

GetDstComponentId returns the DstComponentId field if non-nil, zero value otherwise.

### GetDstComponentIdOk

`func (o *DeclareRequirementEdgeIntentBody) GetDstComponentIdOk() (*string, bool)`

GetDstComponentIdOk returns a tuple with the DstComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstComponentId

`func (o *DeclareRequirementEdgeIntentBody) SetDstComponentId(v string)`

SetDstComponentId sets DstComponentId field to given value.

### HasDstComponentId

`func (o *DeclareRequirementEdgeIntentBody) HasDstComponentId() bool`

HasDstComponentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


