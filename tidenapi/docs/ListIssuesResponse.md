# ListIssuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | Pointer to [**[]Issue**](Issue.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 
**Processing** | Pointer to **bool** | true when ungrouped events exist for this product (worker hasn&#39;t caught up); drives the \&quot;Receiving events, grouping…\&quot; banner. | [optional] 

## Methods

### NewListIssuesResponse

`func NewListIssuesResponse() *ListIssuesResponse`

NewListIssuesResponse instantiates a new ListIssuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIssuesResponseWithDefaults

`func NewListIssuesResponseWithDefaults() *ListIssuesResponse`

NewListIssuesResponseWithDefaults instantiates a new ListIssuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *ListIssuesResponse) GetIssues() []Issue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ListIssuesResponse) GetIssuesOk() (*[]Issue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ListIssuesResponse) SetIssues(v []Issue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ListIssuesResponse) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetPagination

`func (o *ListIssuesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListIssuesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListIssuesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListIssuesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetProcessing

`func (o *ListIssuesResponse) GetProcessing() bool`

GetProcessing returns the Processing field if non-nil, zero value otherwise.

### GetProcessingOk

`func (o *ListIssuesResponse) GetProcessingOk() (*bool, bool)`

GetProcessingOk returns a tuple with the Processing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessing

`func (o *ListIssuesResponse) SetProcessing(v bool)`

SetProcessing sets Processing field to given value.

### HasProcessing

`func (o *ListIssuesResponse) HasProcessing() bool`

HasProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


