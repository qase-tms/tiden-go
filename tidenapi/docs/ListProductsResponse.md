# ListProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]Product**](Product.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 
**Items** | Pointer to [**[]ProductWithSummary**](ProductWithSummary.md) | Same page as products, each with a per-product rollup for the list page. | [optional] 

## Methods

### NewListProductsResponse

`func NewListProductsResponse() *ListProductsResponse`

NewListProductsResponse instantiates a new ListProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductsResponseWithDefaults

`func NewListProductsResponseWithDefaults() *ListProductsResponse`

NewListProductsResponseWithDefaults instantiates a new ListProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *ListProductsResponse) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ListProductsResponse) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ListProductsResponse) SetProducts(v []Product)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ListProductsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPagination

`func (o *ListProductsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListProductsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListProductsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListProductsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetItems

`func (o *ListProductsResponse) GetItems() []ProductWithSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListProductsResponse) GetItemsOk() (*[]ProductWithSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListProductsResponse) SetItems(v []ProductWithSummary)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListProductsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


