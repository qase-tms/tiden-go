# ListComponentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]Component**](Component.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListComponentsResponse

`func NewListComponentsResponse() *ListComponentsResponse`

NewListComponentsResponse instantiates a new ListComponentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListComponentsResponseWithDefaults

`func NewListComponentsResponseWithDefaults() *ListComponentsResponse`

NewListComponentsResponseWithDefaults instantiates a new ListComponentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ListComponentsResponse) GetComponents() []Component`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ListComponentsResponse) GetComponentsOk() (*[]Component, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ListComponentsResponse) SetComponents(v []Component)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ListComponentsResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetPagination

`func (o *ListComponentsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListComponentsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListComponentsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListComponentsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


