# \AgentRunServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRunServiceCancelAgentRun**](AgentRunServiceAPI.md#AgentRunServiceCancelAgentRun) | **Post** /v1/agent-runs/{id}:cancel | Cancels a pending or running agent run.
[**AgentRunServiceGetAgentRun**](AgentRunServiceAPI.md#AgentRunServiceGetAgentRun) | **Get** /v1/agent-runs/{id} | Fetches one agent run by id.
[**AgentRunServiceListAgentRunEvents**](AgentRunServiceAPI.md#AgentRunServiceListAgentRunEvents) | **Get** /v1/agent-runs/{runId}/events | Lists the event log of an agent run.
[**AgentRunServiceListAgentRuns**](AgentRunServiceAPI.md#AgentRunServiceListAgentRuns) | **Get** /v1/agent-configs/{agentConfigId}/runs | Lists the runs of an agent configuration.
[**AgentRunServiceStartAgentRun**](AgentRunServiceAPI.md#AgentRunServiceStartAgentRun) | **Post** /v1/agent-configs/{agentConfigId}/runs | Starts a run of an agent configuration.
[**AgentRunServiceStreamAgentRun**](AgentRunServiceAPI.md#AgentRunServiceStreamAgentRun) | **Get** /v1/agent-runs/{runId}/events:stream | Streams an agent run&#39;s events as they happen.



## AgentRunServiceCancelAgentRun

> CancelAgentRunResponse AgentRunServiceCancelAgentRun(ctx, id).CancelAgentRunBody(cancelAgentRunBody).Execute()

Cancels a pending or running agent run.



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
	cancelAgentRunBody := *openapiclient.NewCancelAgentRunBody() // CancelAgentRunBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceCancelAgentRun(context.Background(), id).CancelAgentRunBody(cancelAgentRunBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceCancelAgentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceCancelAgentRun`: CancelAgentRunResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceCancelAgentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceCancelAgentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelAgentRunBody** | [**CancelAgentRunBody**](CancelAgentRunBody.md) |  | 

### Return type

[**CancelAgentRunResponse**](CancelAgentRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRunServiceGetAgentRun

> GetAgentRunResponse AgentRunServiceGetAgentRun(ctx, id).Execute()

Fetches one agent run by id.



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
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceGetAgentRun(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceGetAgentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceGetAgentRun`: GetAgentRunResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceGetAgentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceGetAgentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAgentRunResponse**](GetAgentRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRunServiceListAgentRunEvents

> ListAgentRunEventsResponse AgentRunServiceListAgentRunEvents(ctx, runId).PageSize(pageSize).PageToken(pageToken).Execute()

Lists the event log of an agent run.



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
	runId := "runId_example" // string | 
	pageSize := int32(56) // int32 |  (optional)
	pageToken := "pageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceListAgentRunEvents(context.Background(), runId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceListAgentRunEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceListAgentRunEvents`: ListAgentRunEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceListAgentRunEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceListAgentRunEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | 
 **pageToken** | **string** |  | 

### Return type

[**ListAgentRunEventsResponse**](ListAgentRunEventsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRunServiceListAgentRuns

> ListAgentRunsResponse AgentRunServiceListAgentRuns(ctx, agentConfigId).PageSize(pageSize).PageToken(pageToken).Execute()

Lists the runs of an agent configuration.



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
	agentConfigId := "agentConfigId_example" // string | 
	pageSize := int32(56) // int32 |  (optional)
	pageToken := "pageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceListAgentRuns(context.Background(), agentConfigId).PageSize(pageSize).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceListAgentRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceListAgentRuns`: ListAgentRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceListAgentRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentConfigId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceListAgentRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** |  | 
 **pageToken** | **string** |  | 

### Return type

[**ListAgentRunsResponse**](ListAgentRunsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRunServiceStartAgentRun

> StartAgentRunResponse AgentRunServiceStartAgentRun(ctx, agentConfigId).StartAgentRunBody(startAgentRunBody).Execute()

Starts a run of an agent configuration.



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
	agentConfigId := "agentConfigId_example" // string | 
	startAgentRunBody := *openapiclient.NewStartAgentRunBody() // StartAgentRunBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceStartAgentRun(context.Background(), agentConfigId).StartAgentRunBody(startAgentRunBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceStartAgentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceStartAgentRun`: StartAgentRunResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceStartAgentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentConfigId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceStartAgentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAgentRunBody** | [**StartAgentRunBody**](StartAgentRunBody.md) |  | 

### Return type

[**StartAgentRunResponse**](StartAgentRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRunServiceStreamAgentRun

> StreamResultOfAgentRunEvent AgentRunServiceStreamAgentRun(ctx, runId).AfterEventId(afterEventId).Execute()

Streams an agent run's events as they happen.



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
	runId := "runId_example" // string | 
	afterEventId := "afterEventId_example" // string | resume from this event id (cursor); empty starts from the beginning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRunServiceAPI.AgentRunServiceStreamAgentRun(context.Background(), runId).AfterEventId(afterEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRunServiceAPI.AgentRunServiceStreamAgentRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRunServiceStreamAgentRun`: StreamResultOfAgentRunEvent
	fmt.Fprintf(os.Stdout, "Response from `AgentRunServiceAPI.AgentRunServiceStreamAgentRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRunServiceStreamAgentRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterEventId** | **string** | resume from this event id (cursor); empty starts from the beginning. | 

### Return type

[**StreamResultOfAgentRunEvent**](StreamResultOfAgentRunEvent.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

