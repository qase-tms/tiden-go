# CriterionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailingTests** | Pointer to [**[]FailingTest**](FailingTest.md) |  | [optional] 
**OpenDefects** | Pointer to [**[]OpenDefect**](OpenDefect.md) |  | [optional] 
**CoverageGap** | Pointer to **string** |  | [optional] 

## Methods

### NewCriterionDetail

`func NewCriterionDetail() *CriterionDetail`

NewCriterionDetail instantiates a new CriterionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionDetailWithDefaults

`func NewCriterionDetailWithDefaults() *CriterionDetail`

NewCriterionDetailWithDefaults instantiates a new CriterionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailingTests

`func (o *CriterionDetail) GetFailingTests() []FailingTest`

GetFailingTests returns the FailingTests field if non-nil, zero value otherwise.

### GetFailingTestsOk

`func (o *CriterionDetail) GetFailingTestsOk() (*[]FailingTest, bool)`

GetFailingTestsOk returns a tuple with the FailingTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingTests

`func (o *CriterionDetail) SetFailingTests(v []FailingTest)`

SetFailingTests sets FailingTests field to given value.

### HasFailingTests

`func (o *CriterionDetail) HasFailingTests() bool`

HasFailingTests returns a boolean if a field has been set.

### GetOpenDefects

`func (o *CriterionDetail) GetOpenDefects() []OpenDefect`

GetOpenDefects returns the OpenDefects field if non-nil, zero value otherwise.

### GetOpenDefectsOk

`func (o *CriterionDetail) GetOpenDefectsOk() (*[]OpenDefect, bool)`

GetOpenDefectsOk returns a tuple with the OpenDefects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDefects

`func (o *CriterionDetail) SetOpenDefects(v []OpenDefect)`

SetOpenDefects sets OpenDefects field to given value.

### HasOpenDefects

`func (o *CriterionDetail) HasOpenDefects() bool`

HasOpenDefects returns a boolean if a field has been set.

### GetCoverageGap

`func (o *CriterionDetail) GetCoverageGap() string`

GetCoverageGap returns the CoverageGap field if non-nil, zero value otherwise.

### GetCoverageGapOk

`func (o *CriterionDetail) GetCoverageGapOk() (*string, bool)`

GetCoverageGapOk returns a tuple with the CoverageGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageGap

`func (o *CriterionDetail) SetCoverageGap(v string)`

SetCoverageGap sets CoverageGap field to given value.

### HasCoverageGap

`func (o *CriterionDetail) HasCoverageGap() bool`

HasCoverageGap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


