# ListRunResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]TestRunResult**](TestRunResult.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListRunResultsResponse

`func NewListRunResultsResponse() *ListRunResultsResponse`

NewListRunResultsResponse instantiates a new ListRunResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRunResultsResponseWithDefaults

`func NewListRunResultsResponseWithDefaults() *ListRunResultsResponse`

NewListRunResultsResponseWithDefaults instantiates a new ListRunResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListRunResultsResponse) GetResults() []TestRunResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListRunResultsResponse) GetResultsOk() (*[]TestRunResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListRunResultsResponse) SetResults(v []TestRunResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListRunResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPagination

`func (o *ListRunResultsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRunResultsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRunResultsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListRunResultsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


