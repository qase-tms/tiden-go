# MatrixRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Coverage** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]MatrixCell**](MatrixCell.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**BranchStatus** | Pointer to **string** |  | [optional] 
**CanonicalId** | Pointer to **string** | Canonical (main) id of this row: source_id for a branch COW copy, empty otherwise. Branch scope keys rows by their branch-local id while parent_id and Verdict.subjects carry main ids, so a client needs this to rebuild the feature tree the way the server resolver does. | [optional] 

## Methods

### NewMatrixRequirement

`func NewMatrixRequirement() *MatrixRequirement`

NewMatrixRequirement instantiates a new MatrixRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixRequirementWithDefaults

`func NewMatrixRequirementWithDefaults() *MatrixRequirement`

NewMatrixRequirementWithDefaults instantiates a new MatrixRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *MatrixRequirement) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *MatrixRequirement) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *MatrixRequirement) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *MatrixRequirement) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetDisplay

`func (o *MatrixRequirement) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *MatrixRequirement) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *MatrixRequirement) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *MatrixRequirement) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetCoverage

`func (o *MatrixRequirement) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *MatrixRequirement) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *MatrixRequirement) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *MatrixRequirement) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetCells

`func (o *MatrixRequirement) GetCells() []MatrixCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *MatrixRequirement) GetCellsOk() (*[]MatrixCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *MatrixRequirement) SetCells(v []MatrixCell)`

SetCells sets Cells field to given value.

### HasCells

`func (o *MatrixRequirement) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetTitle

`func (o *MatrixRequirement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MatrixRequirement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MatrixRequirement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MatrixRequirement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetParentId

`func (o *MatrixRequirement) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MatrixRequirement) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MatrixRequirement) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MatrixRequirement) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetBranchStatus

`func (o *MatrixRequirement) GetBranchStatus() string`

GetBranchStatus returns the BranchStatus field if non-nil, zero value otherwise.

### GetBranchStatusOk

`func (o *MatrixRequirement) GetBranchStatusOk() (*string, bool)`

GetBranchStatusOk returns a tuple with the BranchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchStatus

`func (o *MatrixRequirement) SetBranchStatus(v string)`

SetBranchStatus sets BranchStatus field to given value.

### HasBranchStatus

`func (o *MatrixRequirement) HasBranchStatus() bool`

HasBranchStatus returns a boolean if a field has been set.

### GetCanonicalId

`func (o *MatrixRequirement) GetCanonicalId() string`

GetCanonicalId returns the CanonicalId field if non-nil, zero value otherwise.

### GetCanonicalIdOk

`func (o *MatrixRequirement) GetCanonicalIdOk() (*string, bool)`

GetCanonicalIdOk returns a tuple with the CanonicalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalId

`func (o *MatrixRequirement) SetCanonicalId(v string)`

SetCanonicalId sets CanonicalId field to given value.

### HasCanonicalId

`func (o *MatrixRequirement) HasCanonicalId() bool`

HasCanonicalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


