# ListWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | Pointer to [**[]Workspace**](Workspace.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListWorkspacesResponse

`func NewListWorkspacesResponse() *ListWorkspacesResponse`

NewListWorkspacesResponse instantiates a new ListWorkspacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspacesResponseWithDefaults

`func NewListWorkspacesResponseWithDefaults() *ListWorkspacesResponse`

NewListWorkspacesResponseWithDefaults instantiates a new ListWorkspacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *ListWorkspacesResponse) GetWorkspaces() []Workspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *ListWorkspacesResponse) GetWorkspacesOk() (*[]Workspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *ListWorkspacesResponse) SetWorkspaces(v []Workspace)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *ListWorkspacesResponse) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### GetPagination

`func (o *ListWorkspacesResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWorkspacesResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWorkspacesResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListWorkspacesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


