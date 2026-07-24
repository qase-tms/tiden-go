# IngestTestsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** | pytest | jest | junit | ... | [optional] 
**Tests** | Pointer to [**[]IngestTest**](IngestTest.md) |  | [optional] 

## Methods

### NewIngestTestsBody

`func NewIngestTestsBody() *IngestTestsBody`

NewIngestTestsBody instantiates a new IngestTestsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestTestsBodyWithDefaults

`func NewIngestTestsBodyWithDefaults() *IngestTestsBody`

NewIngestTestsBodyWithDefaults instantiates a new IngestTestsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *IngestTestsBody) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *IngestTestsBody) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *IngestTestsBody) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *IngestTestsBody) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFramework

`func (o *IngestTestsBody) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *IngestTestsBody) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *IngestTestsBody) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *IngestTestsBody) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetTests

`func (o *IngestTestsBody) GetTests() []IngestTest`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *IngestTestsBody) GetTestsOk() (*[]IngestTest, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *IngestTestsBody) SetTests(v []IngestTest)`

SetTests sets Tests field to given value.

### HasTests

`func (o *IngestTestsBody) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


