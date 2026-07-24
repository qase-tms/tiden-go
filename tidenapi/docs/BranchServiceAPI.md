# \BranchServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BranchServiceCreateBranch**](BranchServiceAPI.md#BranchServiceCreateBranch) | **Post** /v1/products/{productId}/branches | 
[**BranchServiceDeleteBranch**](BranchServiceAPI.md#BranchServiceDeleteBranch) | **Delete** /v1/branches/{id} | 
[**BranchServiceGetBranch**](BranchServiceAPI.md#BranchServiceGetBranch) | **Get** /v1/branches/{id} | 
[**BranchServiceGetMergePreview**](BranchServiceAPI.md#BranchServiceGetMergePreview) | **Get** /v1/branches/{id}/merge-preview | 
[**BranchServiceListBranches**](BranchServiceAPI.md#BranchServiceListBranches) | **Get** /v1/products/{productId}/branches | 
[**BranchServiceMergeBranch**](BranchServiceAPI.md#BranchServiceMergeBranch) | **Post** /v1/branches/{id}/merge | 



## BranchServiceCreateBranch

> CreateBranchResponse BranchServiceCreateBranch(ctx, productId).CreateBranchBody(createBranchBody).Execute()



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


## BranchServiceListBranches

> ListBranchesResponse BranchServiceListBranches(ctx, productId).IncludeStats(includeStats).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BranchServiceAPI.BranchServiceListBranches(context.Background(), productId).IncludeStats(includeStats).Execute()
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

