# \IntentServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntentServiceDistillIntent**](IntentServiceAPI.md#IntentServiceDistillIntent) | **Post** /v1/products/{productId}/intent:distill | 



## IntentServiceDistillIntent

> DistillIntentResponse IntentServiceDistillIntent(ctx, productId).DistillIntentBody(distillIntentBody).Execute()



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
	productId := "productId_example" // string | Product whose requirement tree the intent is reconciled against.
	distillIntentBody := *openapiclient.NewDistillIntentBody() // DistillIntentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntentServiceAPI.IntentServiceDistillIntent(context.Background(), productId).DistillIntentBody(distillIntentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntentServiceAPI.IntentServiceDistillIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntentServiceDistillIntent`: DistillIntentResponse
	fmt.Fprintf(os.Stdout, "Response from `IntentServiceAPI.IntentServiceDistillIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product whose requirement tree the intent is reconciled against. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntentServiceDistillIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distillIntentBody** | [**DistillIntentBody**](DistillIntentBody.md) |  | 

### Return type

[**DistillIntentResponse**](DistillIntentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

