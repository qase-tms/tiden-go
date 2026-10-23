# \IssueServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueServiceBulkUpdateIssueStatus**](IssueServiceAPI.md#IssueServiceBulkUpdateIssueStatus) | **Post** /v1/products/{productId}/issues:bulkSetStatus | Sets the same status on many issues at once. Every id must belong to the given product.
[**IssueServiceConfirmSourceMapUpload**](IssueServiceAPI.md#IssueServiceConfirmSourceMapUpload) | **Post** /v1/sourcemaps/{id}:confirm | Finalizes a source-map upload (phase 2 of 2).
[**IssueServiceCreateSourceMapUpload**](IssueServiceAPI.md#IssueServiceCreateSourceMapUpload) | **Post** /v1/products/{productId}/sourcemaps | Starts a source-map upload (phase 1 of 2).
[**IssueServiceGetIssue**](IssueServiceAPI.md#IssueServiceGetIssue) | **Get** /v1/issues/{id} | Gets one issue together with its most recent occurrence.
[**IssueServiceGetIssueEvent**](IssueServiceAPI.md#IssueServiceGetIssueEvent) | **Get** /v1/issues/{id}/events/{eventId} | Gets one occurrence of an issue with symbolicated stack frames.
[**IssueServiceGetIssueEventStats**](IssueServiceAPI.md#IssueServiceGetIssueEventStats) | **Get** /v1/issues/{id}/stats | Returns occurrence statistics for one issue: counts bucketed over time plus a per-environment split.
[**IssueServiceListIssueEvents**](IssueServiceAPI.md#IssueServiceListIssueEvents) | **Get** /v1/issues/{id}/events | Lists an issue&#39;s individual occurrences, most recent first.
[**IssueServiceListIssues**](IssueServiceAPI.md#IssueServiceListIssues) | **Get** /v1/products/{productId}/issues | Lists a product&#39;s issues, most recently active first.
[**IssueServiceListReleaseIssues**](IssueServiceAPI.md#IssueServiceListReleaseIssues) | **Get** /v1/releases/{releaseId}/issues | Returns the issues attributable to a release: those first seen in it, plus a count of every issue seen during it. The post-deploy regression check.
[**IssueServiceUpdateIssueStatus**](IssueServiceAPI.md#IssueServiceUpdateIssueStatus) | **Post** /v1/issues/{id}:setStatus | Sets one issue&#39;s status to unresolved, resolved, or ignored.



## IssueServiceBulkUpdateIssueStatus

> BulkUpdateIssueStatusResponse IssueServiceBulkUpdateIssueStatus(ctx, productId).BulkUpdateIssueStatusBody(bulkUpdateIssueStatusBody).Execute()

Sets the same status on many issues at once. Every id must belong to the given product.

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
	bulkUpdateIssueStatusBody := *openapiclient.NewBulkUpdateIssueStatusBody() // BulkUpdateIssueStatusBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceBulkUpdateIssueStatus(context.Background(), productId).BulkUpdateIssueStatusBody(bulkUpdateIssueStatusBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceBulkUpdateIssueStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceBulkUpdateIssueStatus`: BulkUpdateIssueStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceBulkUpdateIssueStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceBulkUpdateIssueStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateIssueStatusBody** | [**BulkUpdateIssueStatusBody**](BulkUpdateIssueStatusBody.md) |  | 

### Return type

[**BulkUpdateIssueStatusResponse**](BulkUpdateIssueStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceConfirmSourceMapUpload

> ConfirmSourceMapUploadResponse IssueServiceConfirmSourceMapUpload(ctx, id).Body(body).Execute()

Finalizes a source-map upload (phase 2 of 2).



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

Starts a source-map upload (phase 1 of 2).



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


## IssueServiceGetIssue

> GetIssueResponse IssueServiceGetIssue(ctx, id).Execute()

Gets one issue together with its most recent occurrence.



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
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceGetIssue(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceGetIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceGetIssue`: GetIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceGetIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIssueResponse**](GetIssueResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceGetIssueEvent

> GetIssueEventResponse IssueServiceGetIssueEvent(ctx, id, eventId).Execute()

Gets one occurrence of an issue with symbolicated stack frames.



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
	id := "id_example" // string | issue id
	eventId := "eventId_example" // string | issue_events.id (row id, not the SDK event_id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceGetIssueEvent(context.Background(), id, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceGetIssueEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceGetIssueEvent`: GetIssueEventResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceGetIssueEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | issue id | 
**eventId** | **string** | issue_events.id (row id, not the SDK event_id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceGetIssueEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIssueEventResponse**](GetIssueEventResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceGetIssueEventStats

> GetIssueEventStatsResponse IssueServiceGetIssueEventStats(ctx, id).Execute()

Returns occurrence statistics for one issue: counts bucketed over time plus a per-environment split.



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
	id := "id_example" // string | issue id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceGetIssueEventStats(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceGetIssueEventStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceGetIssueEventStats`: GetIssueEventStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceGetIssueEventStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | issue id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceGetIssueEventStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIssueEventStatsResponse**](GetIssueEventStatsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceListIssueEvents

> ListIssueEventsResponse IssueServiceListIssueEvents(ctx, id).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()

Lists an issue's individual occurrences, most recent first.



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
	id := "id_example" // string | issue id
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceListIssueEvents(context.Background(), id).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceListIssueEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceListIssueEvents`: ListIssueEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceListIssueEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | issue id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceListIssueEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListIssueEventsResponse**](ListIssueEventsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceListIssues

> ListIssuesResponse IssueServiceListIssues(ctx, productId).Status(status).EnvironmentId(environmentId).ReleaseId(releaseId).Sort(sort).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).ComponentId(componentId).Platforms(platforms).Levels(levels).Period(period).Search(search).Signal(signal).Execute()

Lists a product's issues, most recently active first.



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
	status := "status_example" // string | unresolved|resolved|ignored, \"\" = all (optional)
	environmentId := "environmentId_example" // string | optional filter (optional)
	releaseId := "releaseId_example" // string | optional filter (optional)
	sort := "sort_example" // string | last_seen|first_seen|times_seen (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)
	componentId := "componentId_example" // string | optional: filter by attributed component (optional)
	platforms := []string{"Inner_example"} // []string | optional: filter by platform (platform dropdown, OR within) (optional)
	levels := []string{"Inner_example"} // []string | optional: filter by level (token bar, OR within) (optional)
	period := "period_example" // string | optional last-seen window: 3m|1h|12h|1d|7d|30d (\"\" = all time) (optional)
	search := "search_example" // string | optional: case-insensitive match on title/culprit (optional)
	signal := "signal_example" // string | optional saved-view predicate: blocking|spiking|new|regressed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceListIssues(context.Background(), productId).Status(status).EnvironmentId(environmentId).ReleaseId(releaseId).Sort(sort).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).ComponentId(componentId).Platforms(platforms).Levels(levels).Period(period).Search(search).Signal(signal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceListIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceListIssues`: ListIssuesResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceListIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceListIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | unresolved|resolved|ignored, \&quot;\&quot; &#x3D; all | 
 **environmentId** | **string** | optional filter | 
 **releaseId** | **string** | optional filter | 
 **sort** | **string** | last_seen|first_seen|times_seen | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **componentId** | **string** | optional: filter by attributed component | 
 **platforms** | **[]string** | optional: filter by platform (platform dropdown, OR within) | 
 **levels** | **[]string** | optional: filter by level (token bar, OR within) | 
 **period** | **string** | optional last-seen window: 3m|1h|12h|1d|7d|30d (\&quot;\&quot; &#x3D; all time) | 
 **search** | **string** | optional: case-insensitive match on title/culprit | 
 **signal** | **string** | optional saved-view predicate: blocking|spiking|new|regressed | 

### Return type

[**ListIssuesResponse**](ListIssuesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceListReleaseIssues

> ListReleaseIssuesResponse IssueServiceListReleaseIssues(ctx, releaseId).Execute()

Returns the issues attributable to a release: those first seen in it, plus a count of every issue seen during it. The post-deploy regression check.

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
	releaseId := "releaseId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceListReleaseIssues(context.Background(), releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceListReleaseIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceListReleaseIssues`: ListReleaseIssuesResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceListReleaseIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceListReleaseIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListReleaseIssuesResponse**](ListReleaseIssuesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueServiceUpdateIssueStatus

> UpdateIssueStatusResponse IssueServiceUpdateIssueStatus(ctx, id).UpdateIssueStatusBody(updateIssueStatusBody).Execute()

Sets one issue's status to unresolved, resolved, or ignored.



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
	updateIssueStatusBody := *openapiclient.NewUpdateIssueStatusBody() // UpdateIssueStatusBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueServiceAPI.IssueServiceUpdateIssueStatus(context.Background(), id).UpdateIssueStatusBody(updateIssueStatusBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueServiceAPI.IssueServiceUpdateIssueStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueServiceUpdateIssueStatus`: UpdateIssueStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueServiceAPI.IssueServiceUpdateIssueStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueServiceUpdateIssueStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIssueStatusBody** | [**UpdateIssueStatusBody**](UpdateIssueStatusBody.md) |  | 

### Return type

[**UpdateIssueStatusResponse**](UpdateIssueStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

