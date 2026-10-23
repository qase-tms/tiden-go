# ListIntentSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]IntentSession**](IntentSession.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListIntentSessionsResponse

`func NewListIntentSessionsResponse() *ListIntentSessionsResponse`

NewListIntentSessionsResponse instantiates a new ListIntentSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntentSessionsResponseWithDefaults

`func NewListIntentSessionsResponseWithDefaults() *ListIntentSessionsResponse`

NewListIntentSessionsResponseWithDefaults instantiates a new ListIntentSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ListIntentSessionsResponse) GetSessions() []IntentSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListIntentSessionsResponse) GetSessionsOk() (*[]IntentSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListIntentSessionsResponse) SetSessions(v []IntentSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ListIntentSessionsResponse) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetPagination

`func (o *ListIntentSessionsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListIntentSessionsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListIntentSessionsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListIntentSessionsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


