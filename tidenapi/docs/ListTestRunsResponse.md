# ListTestRunsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runs** | Pointer to [**[]TestRun**](TestRun.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListTestRunsResponse

`func NewListTestRunsResponse() *ListTestRunsResponse`

NewListTestRunsResponse instantiates a new ListTestRunsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTestRunsResponseWithDefaults

`func NewListTestRunsResponseWithDefaults() *ListTestRunsResponse`

NewListTestRunsResponseWithDefaults instantiates a new ListTestRunsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuns

`func (o *ListTestRunsResponse) GetRuns() []TestRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ListTestRunsResponse) GetRunsOk() (*[]TestRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ListTestRunsResponse) SetRuns(v []TestRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ListTestRunsResponse) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetPagination

`func (o *ListTestRunsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTestRunsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTestRunsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListTestRunsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


