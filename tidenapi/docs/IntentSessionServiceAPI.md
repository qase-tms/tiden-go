# \IntentSessionServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntentSessionServiceGetIntentSession**](IntentSessionServiceAPI.md#IntentSessionServiceGetIntentSession) | **Get** /v1/products/{productId}/intent-sessions/{sessionId} | Fetches one session by id.
[**IntentSessionServiceListIntentSessions**](IntentSessionServiceAPI.md#IntentSessionServiceListIntentSessions) | **Get** /v1/products/{productId}/intent-sessions | Lists a product&#39;s sessions, most recently active first.
[**IntentSessionServiceRecordSessionSettlement**](IntentSessionServiceAPI.md#IntentSessionServiceRecordSessionSettlement) | **Post** /v1/products/{productId}/intent-sessions/{sessionId}:settle | Records or amends a session&#39;s settlement decision.
[**IntentSessionServiceUpsertIntentSession**](IntentSessionServiceAPI.md#IntentSessionServiceUpsertIntentSession) | **Post** /v1/products/{productId}/intent-sessions | Creates or patches a session, keyed by the client-generated session id.



## IntentSessionServiceGetIntentSession

> GetIntentSessionResponse IntentSessionServiceGetIntentSession(ctx, productId, sessionId).Execute()

Fetches one session by id.

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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntentSessionServiceAPI.IntentSessionServiceGetIntentSession(context.Background(), productId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntentSessionServiceAPI.IntentSessionServiceGetIntentSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntentSessionServiceGetIntentSession`: GetIntentSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntentSessionServiceAPI.IntentSessionServiceGetIntentSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntentSessionServiceGetIntentSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIntentSessionResponse**](GetIntentSessionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntentSessionServiceListIntentSessions

> ListIntentSessionsResponse IntentSessionServiceListIntentSessions(ctx, productId).Status(status).GitBranch(gitBranch).BranchId(branchId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).EffectiveStatus(effectiveStatus).Execute()

Lists a product's sessions, most recently active first.



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
	status := "status_example" // string | optional filter — the stored column, verbatim (optional)
	gitBranch := "gitBranch_example" // string | optional filter — cross-machine \"find my open session\" (optional)
	branchId := "branchId_example" // string | optional filter (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)
	effectiveStatus := "effectiveStatus_example" // string | effective_status is an optional filter selecting by the DERIVED status instead of the stored column: \"expired\" (provisional sessions idle past the provisional TTL) or \"stale\" (materialized sessions idle past the materialized TTL). Mutually exclusive with status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntentSessionServiceAPI.IntentSessionServiceListIntentSessions(context.Background(), productId).Status(status).GitBranch(gitBranch).BranchId(branchId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).EffectiveStatus(effectiveStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntentSessionServiceAPI.IntentSessionServiceListIntentSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntentSessionServiceListIntentSessions`: ListIntentSessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `IntentSessionServiceAPI.IntentSessionServiceListIntentSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntentSessionServiceListIntentSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | optional filter — the stored column, verbatim | 
 **gitBranch** | **string** | optional filter — cross-machine \&quot;find my open session\&quot; | 
 **branchId** | **string** | optional filter | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **effectiveStatus** | **string** | effective_status is an optional filter selecting by the DERIVED status instead of the stored column: \&quot;expired\&quot; (provisional sessions idle past the provisional TTL) or \&quot;stale\&quot; (materialized sessions idle past the materialized TTL). Mutually exclusive with status. | 

### Return type

[**ListIntentSessionsResponse**](ListIntentSessionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntentSessionServiceRecordSessionSettlement

> RecordSessionSettlementResponse IntentSessionServiceRecordSessionSettlement(ctx, productId, sessionId).RecordSessionSettlementBody(recordSessionSettlementBody).Execute()

Records or amends a session's settlement decision.



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
	sessionId := "sessionId_example" // string | 
	recordSessionSettlementBody := *openapiclient.NewRecordSessionSettlementBody() // RecordSessionSettlementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntentSessionServiceAPI.IntentSessionServiceRecordSessionSettlement(context.Background(), productId, sessionId).RecordSessionSettlementBody(recordSessionSettlementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntentSessionServiceAPI.IntentSessionServiceRecordSessionSettlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntentSessionServiceRecordSessionSettlement`: RecordSessionSettlementResponse
	fmt.Fprintf(os.Stdout, "Response from `IntentSessionServiceAPI.IntentSessionServiceRecordSessionSettlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntentSessionServiceRecordSessionSettlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordSessionSettlementBody** | [**RecordSessionSettlementBody**](RecordSessionSettlementBody.md) |  | 

### Return type

[**RecordSessionSettlementResponse**](RecordSessionSettlementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntentSessionServiceUpsertIntentSession

> UpsertIntentSessionResponse IntentSessionServiceUpsertIntentSession(ctx, productId).UpsertIntentSessionBody(upsertIntentSessionBody).Execute()

Creates or patches a session, keyed by the client-generated session id.



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
	upsertIntentSessionBody := *openapiclient.NewUpsertIntentSessionBody() // UpsertIntentSessionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntentSessionServiceAPI.IntentSessionServiceUpsertIntentSession(context.Background(), productId).UpsertIntentSessionBody(upsertIntentSessionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntentSessionServiceAPI.IntentSessionServiceUpsertIntentSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntentSessionServiceUpsertIntentSession`: UpsertIntentSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntentSessionServiceAPI.IntentSessionServiceUpsertIntentSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntentSessionServiceUpsertIntentSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertIntentSessionBody** | [**UpsertIntentSessionBody**](UpsertIntentSessionBody.md) |  | 

### Return type

[**UpsertIntentSessionResponse**](UpsertIntentSessionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

