# \EnvironmentServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentServiceCreateEnvironment**](EnvironmentServiceAPI.md#EnvironmentServiceCreateEnvironment) | **Post** /v1/products/{productId}/environments | Creates an environment in a product.
[**EnvironmentServiceDeleteEnvironment**](EnvironmentServiceAPI.md#EnvironmentServiceDeleteEnvironment) | **Delete** /v1/environments/{id} | Deletes an environment.
[**EnvironmentServiceGetEnvironment**](EnvironmentServiceAPI.md#EnvironmentServiceGetEnvironment) | **Get** /v1/environments/{id} | Fetches one environment by id.
[**EnvironmentServiceListEnvironments**](EnvironmentServiceAPI.md#EnvironmentServiceListEnvironments) | **Get** /v1/products/{productId}/environments | Lists a product&#39;s environments.



## EnvironmentServiceCreateEnvironment

> CreateEnvironmentResponse EnvironmentServiceCreateEnvironment(ctx, productId).CreateEnvironmentBody(createEnvironmentBody).Execute()

Creates an environment in a product.



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
	createEnvironmentBody := *openapiclient.NewCreateEnvironmentBody() // CreateEnvironmentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.EnvironmentServiceCreateEnvironment(context.Background(), productId).CreateEnvironmentBody(createEnvironmentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.EnvironmentServiceCreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentServiceCreateEnvironment`: CreateEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.EnvironmentServiceCreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentServiceCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEnvironmentBody** | [**CreateEnvironmentBody**](CreateEnvironmentBody.md) |  | 

### Return type

[**CreateEnvironmentResponse**](CreateEnvironmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentServiceDeleteEnvironment

> map[string]interface{} EnvironmentServiceDeleteEnvironment(ctx, id).Execute()

Deletes an environment.



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
	resp, r, err := apiClient.EnvironmentServiceAPI.EnvironmentServiceDeleteEnvironment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.EnvironmentServiceDeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentServiceDeleteEnvironment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.EnvironmentServiceDeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentServiceDeleteEnvironmentRequest struct via the builder pattern


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


## EnvironmentServiceGetEnvironment

> GetEnvironmentResponse EnvironmentServiceGetEnvironment(ctx, id).Execute()

Fetches one environment by id.



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
	resp, r, err := apiClient.EnvironmentServiceAPI.EnvironmentServiceGetEnvironment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.EnvironmentServiceGetEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentServiceGetEnvironment`: GetEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.EnvironmentServiceGetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentServiceGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEnvironmentResponse**](GetEnvironmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentServiceListEnvironments

> ListEnvironmentsResponse EnvironmentServiceListEnvironments(ctx, productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()

Lists a product's environments.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.EnvironmentServiceListEnvironments(context.Background(), productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.EnvironmentServiceListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentServiceListEnvironments`: ListEnvironmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.EnvironmentServiceListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentServiceListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListEnvironmentsResponse**](ListEnvironmentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

