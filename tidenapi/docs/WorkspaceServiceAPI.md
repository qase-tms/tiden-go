# \WorkspaceServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceServiceListWorkspaces**](WorkspaceServiceAPI.md#WorkspaceServiceListWorkspaces) | **Get** /v1/workspaces | 



## WorkspaceServiceListWorkspaces

> ListWorkspacesResponse WorkspaceServiceListWorkspaces(ctx).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qase-tms/tiden-go/tidenapi"
)

func main() {
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceServiceAPI.WorkspaceServiceListWorkspaces(context.Background()).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceServiceAPI.WorkspaceServiceListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceServiceListWorkspaces`: ListWorkspacesResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceServiceAPI.WorkspaceServiceListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceServiceListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListWorkspacesResponse**](ListWorkspacesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

