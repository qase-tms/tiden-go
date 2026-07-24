# \RequirementServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequirementServiceCreateRequirement**](RequirementServiceAPI.md#RequirementServiceCreateRequirement) | **Post** /v1/products/{productId}/requirements | 
[**RequirementServiceDeleteRequirement**](RequirementServiceAPI.md#RequirementServiceDeleteRequirement) | **Delete** /v1/requirements/{id} | 
[**RequirementServiceGetRequirement**](RequirementServiceAPI.md#RequirementServiceGetRequirement) | **Get** /v1/requirements/{id} | 
[**RequirementServiceListRequirements**](RequirementServiceAPI.md#RequirementServiceListRequirements) | **Get** /v1/products/{productId}/requirements | 
[**RequirementServiceUpdateRequirement**](RequirementServiceAPI.md#RequirementServiceUpdateRequirement) | **Put** /v1/requirements/{id} | 



## RequirementServiceCreateRequirement

> CreateRequirementResponse RequirementServiceCreateRequirement(ctx, productId).CreateRequirementBody(createRequirementBody).Execute()



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
	createRequirementBody := *openapiclient.NewCreateRequirementBody() // CreateRequirementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementServiceAPI.RequirementServiceCreateRequirement(context.Background(), productId).CreateRequirementBody(createRequirementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementServiceAPI.RequirementServiceCreateRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequirementServiceCreateRequirement`: CreateRequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementServiceAPI.RequirementServiceCreateRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequirementServiceCreateRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRequirementBody** | [**CreateRequirementBody**](CreateRequirementBody.md) |  | 

### Return type

[**CreateRequirementResponse**](CreateRequirementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequirementServiceDeleteRequirement

> DeleteRequirementResponse RequirementServiceDeleteRequirement(ctx, id).Branch(branch).Execute()



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
	branch := "branch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementServiceAPI.RequirementServiceDeleteRequirement(context.Background(), id).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementServiceAPI.RequirementServiceDeleteRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequirementServiceDeleteRequirement`: DeleteRequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementServiceAPI.RequirementServiceDeleteRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequirementServiceDeleteRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** |  | 

### Return type

[**DeleteRequirementResponse**](DeleteRequirementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequirementServiceGetRequirement

> GetRequirementResponse RequirementServiceGetRequirement(ctx, id).Execute()



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
	resp, r, err := apiClient.RequirementServiceAPI.RequirementServiceGetRequirement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementServiceAPI.RequirementServiceGetRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequirementServiceGetRequirement`: GetRequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementServiceAPI.RequirementServiceGetRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequirementServiceGetRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRequirementResponse**](GetRequirementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequirementServiceListRequirements

> ListRequirementsResponse RequirementServiceListRequirements(ctx, productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).IncludeSources(includeSources).Execute()



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
	includeSources := true // bool | When true, each returned requirement carries its full sources list (not just source_count). Agents need this for source-based identity matching (e.g. github_local_id / jira_issue anchors); the web UI leaves it off. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementServiceAPI.RequirementServiceListRequirements(context.Background(), productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).IncludeSources(includeSources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementServiceAPI.RequirementServiceListRequirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequirementServiceListRequirements`: ListRequirementsResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementServiceAPI.RequirementServiceListRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequirementServiceListRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **branch** | **string** |  | 
 **includeSources** | **bool** | When true, each returned requirement carries its full sources list (not just source_count). Agents need this for source-based identity matching (e.g. github_local_id / jira_issue anchors); the web UI leaves it off. | 

### Return type

[**ListRequirementsResponse**](ListRequirementsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequirementServiceUpdateRequirement

> UpdateRequirementResponse RequirementServiceUpdateRequirement(ctx, id).UpdateRequirementBody(updateRequirementBody).Execute()



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
	updateRequirementBody := *openapiclient.NewUpdateRequirementBody() // UpdateRequirementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementServiceAPI.RequirementServiceUpdateRequirement(context.Background(), id).UpdateRequirementBody(updateRequirementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementServiceAPI.RequirementServiceUpdateRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequirementServiceUpdateRequirement`: UpdateRequirementResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementServiceAPI.RequirementServiceUpdateRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequirementServiceUpdateRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequirementBody** | [**UpdateRequirementBody**](UpdateRequirementBody.md) |  | 

### Return type

[**UpdateRequirementResponse**](UpdateRequirementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

