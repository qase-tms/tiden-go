# \IssueServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueServiceConfirmSourceMapUpload**](IssueServiceAPI.md#IssueServiceConfirmSourceMapUpload) | **Post** /v1/sourcemaps/{id}:confirm | Phase 3: server validates the staged object + atomically promotes to live. Exempt: addressed by source-map id (the product-gated entry point is CreateSourceMapUpload); source maps are observability infra, not a billed cap.
[**IssueServiceCreateSourceMapUpload**](IssueServiceAPI.md#IssueServiceCreateSourceMapUpload) | **Post** /v1/products/{productId}/sourcemaps | Phase 1: create a pending row, return a presigned PUT to a staging key.



## IssueServiceConfirmSourceMapUpload

> ConfirmSourceMapUploadResponse IssueServiceConfirmSourceMapUpload(ctx, id).Body(body).Execute()

Phase 3: server validates the staged object + atomically promotes to live. Exempt: addressed by source-map id (the product-gated entry point is CreateSourceMapUpload); source maps are observability infra, not a billed cap.

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceConfirmSourceMapUpload(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceConfirmSourceMapUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceConfirmSourceMapUpload`: ConfirmSourceMapUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceConfirmSourceMapUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceConfirmSourceMapUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**ConfirmSourceMapUploadResponse**](ConfirmSourceMapUploadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceCreateSourceMapUpload

> CreateSourceMapUploadResponse IssueServiceCreateSourceMapUpload(ctx, productId).CreateSourceMapUploadBody(createSourceMapUploadBody).Execute()

Phase 1: create a pending row, return a presigned PUT to a staging key.

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
	createSourceMapUploadBody := *openapiclient.NewCreateSourceMapUploadBody() // CreateSourceMapUploadBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceCreateSourceMapUpload(context.Background(), productId).CreateSourceMapUploadBody(createSourceMapUploadBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceCreateSourceMapUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceCreateSourceMapUpload`: CreateSourceMapUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceCreateSourceMapUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceCreateSourceMapUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSourceMapUploadBody** | [**CreateSourceMapUploadBody**](CreateSourceMapUploadBody.md) |  | 

### Return type

[**CreateSourceMapUploadResponse**](CreateSourceMapUploadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

