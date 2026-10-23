# \AgentServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentServiceCreateAgentConfig**](AgentServiceAPI.md#AgentServiceCreateAgentConfig) | **Post** /v1/products/{productId}/agent-configs | Creates an agent configuration in a product.
[**AgentServiceDeleteAgentConfig**](AgentServiceAPI.md#AgentServiceDeleteAgentConfig) | **Delete** /v1/agent-configs/{id} | Deletes an agent configuration.
[**AgentServiceGetAgentConfig**](AgentServiceAPI.md#AgentServiceGetAgentConfig) | **Get** /v1/agent-configs/{id} | Fetches one agent configuration by id.
[**AgentServiceListAgentConfigs**](AgentServiceAPI.md#AgentServiceListAgentConfigs) | **Get** /v1/products/{productId}/agent-configs | Lists a product&#39;s agent configurations.
[**AgentServiceListAgentTypes**](AgentServiceAPI.md#AgentServiceListAgentTypes) | **Get** /v1/agent-types | Lists the catalog of available agent types.



## AgentServiceCreateAgentConfig

> CreateAgentConfigResponse AgentServiceCreateAgentConfig(ctx, productId).CreateAgentConfigBody(createAgentConfigBody).Execute()

Creates an agent configuration in a product.



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
	createAgentConfigBody := *openapiclient.NewCreateAgentConfigBody() // CreateAgentConfigBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.AgentServiceCreateAgentConfig(context.Background(), productId).CreateAgentConfigBody(createAgentConfigBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.AgentServiceCreateAgentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentServiceCreateAgentConfig`: CreateAgentConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.AgentServiceCreateAgentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentServiceCreateAgentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAgentConfigBody** | [**CreateAgentConfigBody**](CreateAgentConfigBody.md) |  | 

### Return type

[**CreateAgentConfigResponse**](CreateAgentConfigResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentServiceDeleteAgentConfig

> map[string]interface{} AgentServiceDeleteAgentConfig(ctx, id).Execute()

Deletes an agent configuration.



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
	resp, r, err := apiClient.AgentServiceAPI.AgentServiceDeleteAgentConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.AgentServiceDeleteAgentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentServiceDeleteAgentConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.AgentServiceDeleteAgentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentServiceDeleteAgentConfigRequest struct via the builder pattern


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


## AgentServiceGetAgentConfig

> GetAgentConfigResponse AgentServiceGetAgentConfig(ctx, id).Execute()

Fetches one agent configuration by id.



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
	resp, r, err := apiClient.AgentServiceAPI.AgentServiceGetAgentConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.AgentServiceGetAgentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentServiceGetAgentConfig`: GetAgentConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.AgentServiceGetAgentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentServiceGetAgentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentConfigResponse**](GetAgentConfigResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentServiceListAgentConfigs

> ListAgentConfigsResponse AgentServiceListAgentConfigs(ctx, productId).Execute()

Lists a product's agent configurations.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.AgentServiceListAgentConfigs(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.AgentServiceListAgentConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentServiceListAgentConfigs`: ListAgentConfigsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.AgentServiceListAgentConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentServiceListAgentConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAgentConfigsResponse**](ListAgentConfigsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentServiceListAgentTypes

> ListAgentTypesResponse AgentServiceListAgentTypes(ctx).Execute()

Lists the catalog of available agent types.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.AgentServiceListAgentTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.AgentServiceListAgentTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentServiceListAgentTypes`: ListAgentTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.AgentServiceListAgentTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAgentServiceListAgentTypesRequest struct via the builder pattern


### Return type

[**ListAgentTypesResponse**](ListAgentTypesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

