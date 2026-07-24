# ReportResultsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ResultCreate**](ResultCreate.md) |  | [optional] 

## Methods

### NewReportResultsBody

`func NewReportResultsBody() *ReportResultsBody`

NewReportResultsBody instantiates a new ReportResultsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportResultsBodyWithDefaults

`func NewReportResultsBodyWithDefaults() *ReportResultsBody`

NewReportResultsBodyWithDefaults instantiates a new ReportResultsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ReportResultsBody) GetResults() []ResultCreate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ReportResultsBody) GetResultsOk() (*[]ResultCreate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ReportResultsBody) SetResults(v []ResultCreate)`

SetResults sets Results field to given value.

### HasResults

`func (o *ReportResultsBody) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


