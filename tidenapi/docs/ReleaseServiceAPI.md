# \ReleaseServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseServiceCreateRelease**](ReleaseServiceAPI.md#ReleaseServiceCreateRelease) | **Post** /v1/products/{productId}/releases | Create a release from an external source (CI/SDK). Idempotent upsert on (product, version, environment). The environment is matched by slug and auto-created if unknown.
[**ReleaseServiceGetRelease**](ReleaseServiceAPI.md#ReleaseServiceGetRelease) | **Get** /v1/releases/{id} | 
[**ReleaseServiceListReleases**](ReleaseServiceAPI.md#ReleaseServiceListReleases) | **Get** /v1/products/{productId}/releases | 



## ReleaseServiceCreateRelease

> CreateReleaseResponse ReleaseServiceCreateRelease(ctx, productId).CreateReleaseBody(createReleaseBody).Execute()

Create a release from an external source (CI/SDK). Idempotent upsert on (product, version, environment). The environment is matched by slug and auto-created if unknown.

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
	createReleaseBody := *openapiclient.NewCreateReleaseBody() // CreateReleaseBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseServiceAPI.ReleaseServiceCreateRelease(context.Background(), productId).CreateReleaseBody(createReleaseBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseServiceAPI.ReleaseServiceCreateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseServiceCreateRelease`: CreateReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleaseServiceAPI.ReleaseServiceCreateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseServiceCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReleaseBody** | [**CreateReleaseBody**](CreateReleaseBody.md) |  | 

### Return type

[**CreateReleaseResponse**](CreateReleaseResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseServiceGetRelease

> GetReleaseResponse ReleaseServiceGetRelease(ctx, id).Execute()



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
	resp, r, err := apiClient.ReleaseServiceAPI.ReleaseServiceGetRelease(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseServiceAPI.ReleaseServiceGetRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseServiceGetRelease`: GetReleaseResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleaseServiceAPI.ReleaseServiceGetRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseServiceGetReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReleaseResponse**](GetReleaseResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseServiceListReleases

> ListReleasesResponse ReleaseServiceListReleases(ctx, productId).Environment(environment).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()



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
	environment := "environment_example" // string | optional environment slug filter (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseServiceAPI.ReleaseServiceListReleases(context.Background(), productId).Environment(environment).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseServiceAPI.ReleaseServiceListReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseServiceListReleases`: ListReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleaseServiceAPI.ReleaseServiceListReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseServiceListReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | **string** | optional environment slug filter | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListReleasesResponse**](ListReleasesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

