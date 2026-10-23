# MatrixCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestCase** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**TestId** | Pointer to **string** | test reference; pinned by TIDEN-135 so the grid stays diffable across runs — do not repurpose it.  join key: lets a client attach gate-accurate status to | [optional] 

## Methods

### NewMatrixCell

`func NewMatrixCell() *MatrixCell`

NewMatrixCell instantiates a new MatrixCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatrixCellWithDefaults

`func NewMatrixCellWithDefaults() *MatrixCell`

NewMatrixCellWithDefaults instantiates a new MatrixCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestCase

`func (o *MatrixCell) GetTestCase() string`

GetTestCase returns the TestCase field if non-nil, zero value otherwise.

### GetTestCaseOk

`func (o *MatrixCell) GetTestCaseOk() (*string, bool)`

GetTestCaseOk returns a tuple with the TestCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCase

`func (o *MatrixCell) SetTestCase(v string)`

SetTestCase sets TestCase field to given value.

### HasTestCase

`func (o *MatrixCell) HasTestCase() bool`

HasTestCase returns a boolean if a field has been set.

### GetStatus

`func (o *MatrixCell) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MatrixCell) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MatrixCell) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MatrixCell) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisplay

`func (o *MatrixCell) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *MatrixCell) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *MatrixCell) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *MatrixCell) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTestId

`func (o *MatrixCell) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *MatrixCell) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *MatrixCell) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *MatrixCell) HasTestId() bool`

HasTestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


