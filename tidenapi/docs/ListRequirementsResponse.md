# ListRequirementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | Pointer to [**[]Requirement**](Requirement.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListRequirementsResponse

`func NewListRequirementsResponse() *ListRequirementsResponse`

NewListRequirementsResponse instantiates a new ListRequirementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRequirementsResponseWithDefaults

`func NewListRequirementsResponseWithDefaults() *ListRequirementsResponse`

NewListRequirementsResponseWithDefaults instantiates a new ListRequirementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *ListRequirementsResponse) GetRequirements() []Requirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ListRequirementsResponse) GetRequirementsOk() (*[]Requirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ListRequirementsResponse) SetRequirements(v []Requirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ListRequirementsResponse) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetPagination

`func (o *ListRequirementsResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRequirementsResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRequirementsResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListRequirementsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


