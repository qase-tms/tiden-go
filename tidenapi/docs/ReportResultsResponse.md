# ReportResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Accepted** | Pointer to **string** |  | [optional] 
**Duplicates** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ReportError**](ReportError.md) | Per-entry validation errors, mirroring IngestTestsResponse.errors: the same payload is also attached as google.rpc.Status details on the InvalidArgument error for gRPC clients. | [optional] 

## Methods

### NewReportResultsResponse

`func NewReportResultsResponse() *ReportResultsResponse`

NewReportResultsResponse instantiates a new ReportResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportResultsResponseWithDefaults

`func NewReportResultsResponseWithDefaults() *ReportResultsResponse`

NewReportResultsResponseWithDefaults instantiates a new ReportResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReportResultsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportResultsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportResultsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportResultsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccepted

`func (o *ReportResultsResponse) GetAccepted() string`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *ReportResultsResponse) GetAcceptedOk() (*string, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *ReportResultsResponse) SetAccepted(v string)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *ReportResultsResponse) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetDuplicates

`func (o *ReportResultsResponse) GetDuplicates() string`

GetDuplicates returns the Duplicates field if non-nil, zero value otherwise.

### GetDuplicatesOk

`func (o *ReportResultsResponse) GetDuplicatesOk() (*string, bool)`

GetDuplicatesOk returns a tuple with the Duplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicates

`func (o *ReportResultsResponse) SetDuplicates(v string)`

SetDuplicates sets Duplicates field to given value.

### HasDuplicates

`func (o *ReportResultsResponse) HasDuplicates() bool`

HasDuplicates returns a boolean if a field has been set.

### GetErrors

`func (o *ReportResultsResponse) GetErrors() []ReportError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ReportResultsResponse) GetErrorsOk() (*[]ReportError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ReportResultsResponse) SetErrors(v []ReportError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ReportResultsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


