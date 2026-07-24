# \ComponentServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentServiceCreateComponent**](ComponentServiceAPI.md#ComponentServiceCreateComponent) | **Post** /v1/products/{productId}/components | 
[**ComponentServiceListComponents**](ComponentServiceAPI.md#ComponentServiceListComponents) | **Get** /v1/products/{productId}/components | 
[**ComponentServiceUpdateComponent**](ComponentServiceAPI.md#ComponentServiceUpdateComponent) | **Put** /v1/components/{id} | 



## ComponentServiceCreateComponent

> CreateComponentResponse ComponentServiceCreateComponent(ctx, productId).CreateComponentBody(createComponentBody).Execute()



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
	createComponentBody := *openapiclient.NewCreateComponentBody() // CreateComponentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentServiceAPI.ComponentServiceCreateComponent(context.Background(), productId).CreateComponentBody(createComponentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentServiceAPI.ComponentServiceCreateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentServiceCreateComponent`: CreateComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentServiceAPI.ComponentServiceCreateComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentServiceCreateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createComponentBody** | [**CreateComponentBody**](CreateComponentBody.md) |  | 

### Return type

[**CreateComponentResponse**](CreateComponentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentServiceListComponents

> ListComponentsResponse ComponentServiceListComponents(ctx, productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).Execute()



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
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)
	branch := "branch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentServiceAPI.ComponentServiceListComponents(context.Background(), productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentServiceAPI.ComponentServiceListComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentServiceListComponents`: ListComponentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentServiceAPI.ComponentServiceListComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentServiceListComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **branch** | **string** |  | 

### Return type

[**ListComponentsResponse**](ListComponentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComponentServiceUpdateComponent

> UpdateComponentResponse ComponentServiceUpdateComponent(ctx, id).UpdateComponentBody(updateComponentBody).Execute()



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
	updateComponentBody := *openapiclient.NewUpdateComponentBody() // UpdateComponentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentServiceAPI.ComponentServiceUpdateComponent(context.Background(), id).UpdateComponentBody(updateComponentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentServiceAPI.ComponentServiceUpdateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComponentServiceUpdateComponent`: UpdateComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentServiceAPI.ComponentServiceUpdateComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiComponentServiceUpdateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateComponentBody** | [**UpdateComponentBody**](UpdateComponentBody.md) |  | 

### Return type

[**UpdateComponentResponse**](UpdateComponentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

