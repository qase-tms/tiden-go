# ListIssueEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]IssueEvent**](IssueEvent.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListIssueEventsResponse

`func NewListIssueEventsResponse() *ListIssueEventsResponse`

NewListIssueEventsResponse instantiates a new ListIssueEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIssueEventsResponseWithDefaults

`func NewListIssueEventsResponseWithDefaults() *ListIssueEventsResponse`

NewListIssueEventsResponseWithDefaults instantiates a new ListIssueEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListIssueEventsResponse) GetEvents() []IssueEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListIssueEventsResponse) GetEventsOk() (*[]IssueEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListIssueEventsResponse) SetEvents(v []IssueEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListIssueEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetPagination

`func (o *ListIssueEventsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListIssueEventsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListIssueEventsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListIssueEventsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


