# \BranchServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BranchServiceCreateBranch**](BranchServiceAPI.md#BranchServiceCreateBranch) | **Post** /v1/products/{productId}/branches | Creates a copy-on-write branch of a product&#39;s main line.
[**BranchServiceDeleteBranch**](BranchServiceAPI.md#BranchServiceDeleteBranch) | **Delete** /v1/branches/{id} | Deletes a branch and discards its copy-on-write changes.
[**BranchServiceGetBranch**](BranchServiceAPI.md#BranchServiceGetBranch) | **Get** /v1/branches/{id} | Fetches one branch by id.
[**BranchServiceGetMergePreview**](BranchServiceAPI.md#BranchServiceGetMergePreview) | **Get** /v1/branches/{id}/merge-preview | Previews the effect of merging a branch into main.
[**BranchServiceListBranchCodeLinks**](BranchServiceAPI.md#BranchServiceListBranchCodeLinks) | **Get** /v1/branches/{branchId}/code-links | Lists a branch&#39;s durable code links (git branches, pull requests).
[**BranchServiceListBranches**](BranchServiceAPI.md#BranchServiceListBranches) | **Get** /v1/products/{productId}/branches | Lists a product&#39;s branches.
[**BranchServiceMergeBranch**](BranchServiceAPI.md#BranchServiceMergeBranch) | **Post** /v1/branches/{id}/merge | Merges a branch&#39;s changes into main and closes the branch.
[**BranchServiceUpdateBranch**](BranchServiceAPI.md#BranchServiceUpdateBranch) | **Patch** /v1/branches/{id} | Updates a branch&#39;s description and/or created_by_agent.
[**BranchServiceUpsertBranchCodeLinks**](BranchServiceAPI.md#BranchServiceUpsertBranchCodeLinks) | **Post** /v1/branches/{branchId}/code-links | Upserts a batch of code links onto a branch.



## BranchServiceCreateBranch

> CreateBranchResponse BranchServiceCreateBranch(ctx, productId).CreateBranchBody(createBranchBody).Execute()

Creates a copy-on-write branch of a product's main line.



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
	productId := "productId_example" // string | 
	createBranchBody := *openapiclient.NewCreateBranchBody() // CreateBranchBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceCreateBranch(context.Background(), productId).CreateBranchBody(createBranchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceCreateBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceCreateBranch`: CreateBranchResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceCreateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceCreateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBranchBody** | [**CreateBranchBody**](CreateBranchBody.md) |  | 

### Return type

[**CreateBranchResponse**](CreateBranchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceDeleteBranch

> map[string]interface{} BranchServiceDeleteBranch(ctx, id).Execute()

Deletes a branch and discards its copy-on-write changes.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceDeleteBranch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceDeleteBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceDeleteBranch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceDeleteBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceDeleteBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceGetBranch

> GetBranchResponse BranchServiceGetBranch(ctx, id).Execute()

Fetches one branch by id.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceGetBranch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceGetBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceGetBranch`: GetBranchResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceGetBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceGetBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBranchResponse**](GetBranchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceGetMergePreview

> GetMergePreviewResponse BranchServiceGetMergePreview(ctx, id).Execute()

Previews the effect of merging a branch into main.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceGetMergePreview(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceGetMergePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceGetMergePreview`: GetMergePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceGetMergePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceGetMergePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMergePreviewResponse**](GetMergePreviewResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceListBranchCodeLinks

> ListBranchCodeLinksResponse BranchServiceListBranchCodeLinks(ctx, branchId).Execute()

Lists a branch's durable code links (git branches, pull requests).



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
	branchId := "branchId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceListBranchCodeLinks(context.Background(), branchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceListBranchCodeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceListBranchCodeLinks`: ListBranchCodeLinksResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceListBranchCodeLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceListBranchCodeLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBranchCodeLinksResponse**](ListBranchCodeLinksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceListBranches

> ListBranchesResponse BranchServiceListBranches(ctx, productId).IncludeStats(includeStats).IncludeStatus(includeStatus).Execute()

Lists a product's branches.



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
	productId := "productId_example" // string | 
	includeStats := true // bool | When true, each returned Branch carries BranchChangeStats (per-branch change counts vs main). (optional)
	includeStatus := true // bool | When true, each returned Branch carries loop/latest-run/code-link/intent status signals (Branch.loop, .latest_run, .code_links, .intent). Kept separate from include_stats: the sidebar branch dropdown calls List without stats and must not pay for this extra work either. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceListBranches(context.Background(), productId).IncludeStats(includeStats).IncludeStatus(includeStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceListBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceListBranches`: ListBranchesResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceListBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceListBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeStats** | **bool** | When true, each returned Branch carries BranchChangeStats (per-branch change counts vs main). | 
 **includeStatus** | **bool** | When true, each returned Branch carries loop/latest-run/code-link/intent status signals (Branch.loop, .latest_run, .code_links, .intent). Kept separate from include_stats: the sidebar branch dropdown calls List without stats and must not pay for this extra work either. | 

### Return type

[**ListBranchesResponse**](ListBranchesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceMergeBranch

> MergeBranchResponse BranchServiceMergeBranch(ctx, id).MergeBranchBody(mergeBranchBody).Execute()

Merges a branch's changes into main and closes the branch.



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
	id := "id_example" // string | 
	mergeBranchBody := *openapiclient.NewMergeBranchBody() // MergeBranchBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceMergeBranch(context.Background(), id).MergeBranchBody(mergeBranchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceMergeBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceMergeBranch`: MergeBranchResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceMergeBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceMergeBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mergeBranchBody** | [**MergeBranchBody**](MergeBranchBody.md) |  | 

### Return type

[**MergeBranchResponse**](MergeBranchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceUpdateBranch

> UpdateBranchResponse BranchServiceUpdateBranch(ctx, id).UpdateBranchBody(updateBranchBody).Execute()

Updates a branch's description and/or created_by_agent.



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
	id := "id_example" // string | 
	updateBranchBody := *openapiclient.NewUpdateBranchBody() // UpdateBranchBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceUpdateBranch(context.Background(), id).UpdateBranchBody(updateBranchBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceUpdateBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceUpdateBranch`: UpdateBranchResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceUpdateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceUpdateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBranchBody** | [**UpdateBranchBody**](UpdateBranchBody.md) |  | 

### Return type

[**UpdateBranchResponse**](UpdateBranchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BranchServiceUpsertBranchCodeLinks

> UpsertBranchCodeLinksResponse BranchServiceUpsertBranchCodeLinks(ctx, branchId).UpsertBranchCodeLinksBody(upsertBranchCodeLinksBody).Execute()

Upserts a batch of code links onto a branch.



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
	branchId := "branchId_example" // string | 
	upsertBranchCodeLinksBody := *openapiclient.NewUpsertBranchCodeLinksBody() // UpsertBranchCodeLinksBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceUpsertBranchCodeLinks(context.Background(), branchId).UpsertBranchCodeLinksBody(upsertBranchCodeLinksBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchServiceAPI.BranchServiceUpsertBranchCodeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BranchServiceUpsertBranchCodeLinks`: UpsertBranchCodeLinksResponse
	fmt.Fprintf(os.Stdout, "Response from `BranchServiceAPI.BranchServiceUpsertBranchCodeLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBranchServiceUpsertBranchCodeLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertBranchCodeLinksBody** | [**UpsertBranchCodeLinksBody**](UpsertBranchCodeLinksBody.md) |  | 

### Return type

[**UpsertBranchCodeLinksResponse**](UpsertBranchCodeLinksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

