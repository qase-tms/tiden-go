# ListReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]Release**](Release.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListReleasesResponse

`func NewListReleasesResponse() *ListReleasesResponse`

NewListReleasesResponse instantiates a new ListReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReleasesResponseWithDefaults

`func NewListReleasesResponseWithDefaults() *ListReleasesResponse`

NewListReleasesResponseWithDefaults instantiates a new ListReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *ListReleasesResponse) GetReleases() []Release`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ListReleasesResponse) GetReleasesOk() (*[]Release, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ListReleasesResponse) SetReleases(v []Release)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *ListReleasesResponse) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetPagination

`func (o *ListReleasesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListReleasesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListReleasesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListReleasesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


