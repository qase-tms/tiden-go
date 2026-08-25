# \ProductServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductServiceCreateProduct**](ProductServiceAPI.md#ProductServiceCreateProduct) | **Post** /v1/workspaces/{workspaceId}/products | Creates a product in a workspace.
[**ProductServiceGetProduct**](ProductServiceAPI.md#ProductServiceGetProduct) | **Get** /v1/products/{id} | Fetches one product by id.
[**ProductServiceListProducts**](ProductServiceAPI.md#ProductServiceListProducts) | **Get** /v1/workspaces/{workspaceId}/products | Lists a workspace&#39;s products.
[**ProductServiceVerifyProductSetup**](ProductServiceAPI.md#ProductServiceVerifyProductSetup) | **Post** /v1/products/{productId}/setup:verify | Records a CLI setup verification snapshot for the product.



## ProductServiceCreateProduct

> CreateProductResponse ProductServiceCreateProduct(ctx, workspaceId).CreateProductBody(createProductBody).Execute()

Creates a product in a workspace.



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
	workspaceId := "workspaceId_example" // string | 
	createProductBody := *openapiclient.NewCreateProductBody() // CreateProductBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductServiceAPI.ProductServiceCreateProduct(context.Background(), workspaceId).CreateProductBody(createProductBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductServiceAPI.ProductServiceCreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductServiceCreateProduct`: CreateProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductServiceAPI.ProductServiceCreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductServiceCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProductBody** | [**CreateProductBody**](CreateProductBody.md) |  | 

### Return type

[**CreateProductResponse**](CreateProductResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductServiceGetProduct

> GetProductResponse ProductServiceGetProduct(ctx, id).Execute()

Fetches one product by id.



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
	resp, r, err := apiClient.ProductServiceAPI.ProductServiceGetProduct(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductServiceAPI.ProductServiceGetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductServiceGetProduct`: GetProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductServiceAPI.ProductServiceGetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductServiceGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProductResponse**](GetProductResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductServiceListProducts

> ListProductsResponse ProductServiceListProducts(ctx, workspaceId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()

Lists a workspace's products.



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
	workspaceId := "workspaceId_example" // string | 
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductServiceAPI.ProductServiceListProducts(context.Background(), workspaceId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductServiceAPI.ProductServiceListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductServiceListProducts`: ListProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductServiceAPI.ProductServiceListProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductServiceListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListProductsResponse**](ListProductsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductServiceVerifyProductSetup

> VerifyProductSetupResponse ProductServiceVerifyProductSetup(ctx, productId).VerifyProductSetupBody(verifyProductSetupBody).Execute()

Records a CLI setup verification snapshot for the product.



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
	verifyProductSetupBody := *openapiclient.NewVerifyProductSetupBody() // VerifyProductSetupBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductServiceAPI.ProductServiceVerifyProductSetup(context.Background(), productId).VerifyProductSetupBody(verifyProductSetupBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductServiceAPI.ProductServiceVerifyProductSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductServiceVerifyProductSetup`: VerifyProductSetupResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductServiceAPI.ProductServiceVerifyProductSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductServiceVerifyProductSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyProductSetupBody** | [**VerifyProductSetupBody**](VerifyProductSetupBody.md) |  | 

### Return type

[**VerifyProductSetupResponse**](VerifyProductSetupResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

