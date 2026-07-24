# GetRunSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suites** | Pointer to [**[]RunSuiteSummary**](RunSuiteSummary.md) |  | [optional] 
**Cases** | Pointer to [**[]RunCaseSummary**](RunCaseSummary.md) |  | [optional] 

## Methods

### NewGetRunSummaryResponse

`func NewGetRunSummaryResponse() *GetRunSummaryResponse`

NewGetRunSummaryResponse instantiates a new GetRunSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRunSummaryResponseWithDefaults

`func NewGetRunSummaryResponseWithDefaults() *GetRunSummaryResponse`

NewGetRunSummaryResponseWithDefaults instantiates a new GetRunSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuites

`func (o *GetRunSummaryResponse) GetSuites() []RunSuiteSummary`

GetSuites returns the Suites field if non-nil, zero value otherwise.

### GetSuitesOk

`func (o *GetRunSummaryResponse) GetSuitesOk() (*[]RunSuiteSummary, bool)`

GetSuitesOk returns a tuple with the Suites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuites

`func (o *GetRunSummaryResponse) SetSuites(v []RunSuiteSummary)`

SetSuites sets Suites field to given value.

### HasSuites

`func (o *GetRunSummaryResponse) HasSuites() bool`

HasSuites returns a boolean if a field has been set.

### GetCases

`func (o *GetRunSummaryResponse) GetCases() []RunCaseSummary`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *GetRunSummaryResponse) GetCasesOk() (*[]RunCaseSummary, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *GetRunSummaryResponse) SetCases(v []RunCaseSummary)`

SetCases sets Cases field to given value.

### HasCases

`func (o *GetRunSummaryResponse) HasCases() bool`

HasCases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


