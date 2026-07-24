# \TestRunServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestRunServiceAbortTestRun**](TestRunServiceAPI.md#TestRunServiceAbortTestRun) | **Post** /v1/products/{productId}/runs/{runSeq}:abort | 
[**TestRunServiceCompleteTestRun**](TestRunServiceAPI.md#TestRunServiceCompleteTestRun) | **Post** /v1/products/{productId}/runs/{runSeq}:complete | 
[**TestRunServiceCreateTestRun**](TestRunServiceAPI.md#TestRunServiceCreateTestRun) | **Post** /v1/products/{productId}/runs | 
[**TestRunServiceDeleteTestRun**](TestRunServiceAPI.md#TestRunServiceDeleteTestRun) | **Delete** /v1/products/{productId}/runs/{runSeq} | 
[**TestRunServiceGetRunAttachment**](TestRunServiceAPI.md#TestRunServiceGetRunAttachment) | **Get** /v1/products/{productId}/attachments/{hash} | Resolves a content-hash (uploaded via the reporter multipart route POST /v1/products/{product_id}/attachments:upload) to a presigned download URL. Public so reporter/CLI clients and the SPA (JWT) can both fetch; ATTACHMENT_NOT_FOUND (→ 404) for an unknown hash — the drawer renders \&quot;attachment unavailable\&quot; on that.
[**TestRunServiceGetRunResult**](TestRunServiceAPI.md#TestRunServiceGetRunResult) | **Get** /v1/products/{productId}/runs/{runSeq}/results/{resultId} | 
[**TestRunServiceGetRunSummary**](TestRunServiceAPI.md#TestRunServiceGetRunSummary) | **Get** /v1/products/{productId}/runs/{runSeq}/summary | 
[**TestRunServiceGetTestRun**](TestRunServiceAPI.md#TestRunServiceGetTestRun) | **Get** /v1/products/{productId}/runs/{runSeq} | 
[**TestRunServiceListRunResults**](TestRunServiceAPI.md#TestRunServiceListRunResults) | **Get** /v1/products/{productId}/runs/{runSeq}/results | 
[**TestRunServiceListTestRuns**](TestRunServiceAPI.md#TestRunServiceListTestRuns) | **Get** /v1/products/{productId}/runs | 
[**TestRunServiceReportResults**](TestRunServiceAPI.md#TestRunServiceReportResults) | **Post** /v1/products/{productId}/runs/{runSeq}/results:report | 



## TestRunServiceAbortTestRun

> AbortTestRunResponse TestRunServiceAbortTestRun(ctx, productId, runSeq).Body(body).Execute()



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
	runSeq := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceAbortTestRun(context.Background(), productId, runSeq).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceAbortTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceAbortTestRun`: AbortTestRunResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceAbortTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceAbortTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**AbortTestRunResponse**](AbortTestRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceCompleteTestRun

> CompleteTestRunResponse TestRunServiceCompleteTestRun(ctx, productId, runSeq).Body(body).Execute()



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
	runSeq := int32(56) // int32 | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceCompleteTestRun(context.Background(), productId, runSeq).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceCompleteTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceCompleteTestRun`: CompleteTestRunResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceCompleteTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceCompleteTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**CompleteTestRunResponse**](CompleteTestRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceCreateTestRun

> CreateTestRunResponse TestRunServiceCreateTestRun(ctx, productId).CreateTestRunBody(createTestRunBody).Execute()



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
	createTestRunBody := *openapiclient.NewCreateTestRunBody() // CreateTestRunBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceCreateTestRun(context.Background(), productId).CreateTestRunBody(createTestRunBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceCreateTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceCreateTestRun`: CreateTestRunResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceCreateTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceCreateTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTestRunBody** | [**CreateTestRunBody**](CreateTestRunBody.md) |  | 

### Return type

[**CreateTestRunResponse**](CreateTestRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceDeleteTestRun

> map[string]interface{} TestRunServiceDeleteTestRun(ctx, productId, runSeq).Execute()



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
	runSeq := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceDeleteTestRun(context.Background(), productId, runSeq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceDeleteTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceDeleteTestRun`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceDeleteTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceDeleteTestRunRequest struct via the builder pattern


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


## TestRunServiceGetRunAttachment

> GetRunAttachmentResponse TestRunServiceGetRunAttachment(ctx, productId, hash).Execute()

Resolves a content-hash (uploaded via the reporter multipart route POST /v1/products/{product_id}/attachments:upload) to a presigned download URL. Public so reporter/CLI clients and the SPA (JWT) can both fetch; ATTACHMENT_NOT_FOUND (→ 404) for an unknown hash — the drawer renders \"attachment unavailable\" on that.

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
	hash := "hash_example" // string | sha256 hex content hash

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceGetRunAttachment(context.Background(), productId, hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceGetRunAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceGetRunAttachment`: GetRunAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceGetRunAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**hash** | **string** | sha256 hex content hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceGetRunAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRunAttachmentResponse**](GetRunAttachmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceGetRunResult

> GetRunResultResponse TestRunServiceGetRunResult(ctx, productId, runSeq, resultId).Execute()



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
	runSeq := int32(56) // int32 | 
	resultId := "resultId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceGetRunResult(context.Background(), productId, runSeq, resultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceGetRunResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceGetRunResult`: GetRunResultResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceGetRunResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 
**resultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceGetRunResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetRunResultResponse**](GetRunResultResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceGetRunSummary

> GetRunSummaryResponse TestRunServiceGetRunSummary(ctx, productId, runSeq).Execute()



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
	runSeq := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceGetRunSummary(context.Background(), productId, runSeq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceGetRunSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceGetRunSummary`: GetRunSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceGetRunSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceGetRunSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRunSummaryResponse**](GetRunSummaryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceGetTestRun

> GetTestRunResponse TestRunServiceGetTestRun(ctx, productId, runSeq).Execute()



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
	runSeq := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceGetTestRun(context.Background(), productId, runSeq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceGetTestRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceGetTestRun`: GetTestRunResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceGetTestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceGetTestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTestRunResponse**](GetTestRunResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceListRunResults

> ListRunResultsResponse TestRunServiceListRunResults(ctx, productId, runSeq).Status(status).Search(search).IdentityKey(identityKey).LatestOnly(latestOnly).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()



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
	runSeq := int32(56) // int32 | 
	status := "status_example" // string | optional filter (latest attempt status) (optional)
	search := "search_example" // string | optional title substring (optional)
	identityKey := "identityKey_example" // string | optional: all attempts of one case identity (optional)
	latestOnly := true // bool | collapse retries (default false = all rows) (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceListRunResults(context.Background(), productId, runSeq).Status(status).Search(search).IdentityKey(identityKey).LatestOnly(latestOnly).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceListRunResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceListRunResults`: ListRunResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceListRunResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceListRunResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | optional filter (latest attempt status) | 
 **search** | **string** | optional title substring | 
 **identityKey** | **string** | optional: all attempts of one case identity | 
 **latestOnly** | **bool** | collapse retries (default false &#x3D; all rows) | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListRunResultsResponse**](ListRunResultsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceListTestRuns

> ListTestRunsResponse TestRunServiceListTestRuns(ctx, productId).Status(status).Environment(environment).Branch(branch).Search(search).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()



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
	status := "status_example" // string | optional filter (optional)
	environment := "environment_example" // string | optional environment slug filter (optional)
	branch := "branch_example" // string | optional branch_name filter (optional)
	search := "search_example" // string | optional title substring (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceListTestRuns(context.Background(), productId).Status(status).Environment(environment).Branch(branch).Search(search).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceListTestRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceListTestRuns`: ListTestRunsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceListTestRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceListTestRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | optional filter | 
 **environment** | **string** | optional environment slug filter | 
 **branch** | **string** | optional branch_name filter | 
 **search** | **string** | optional title substring | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 

### Return type

[**ListTestRunsResponse**](ListTestRunsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRunServiceReportResults

> ReportResultsResponse TestRunServiceReportResults(ctx, productId, runSeq).ReportResultsBody(reportResultsBody).Execute()



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
	runSeq := int32(56) // int32 | 
	reportResultsBody := *openapiclient.NewReportResultsBody() // ReportResultsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestRunServiceAPI.TestRunServiceReportResults(context.Background(), productId, runSeq).ReportResultsBody(reportResultsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestRunServiceAPI.TestRunServiceReportResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestRunServiceReportResults`: ReportResultsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestRunServiceAPI.TestRunServiceReportResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**runSeq** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRunServiceReportResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reportResultsBody** | [**ReportResultsBody**](ReportResultsBody.md) |  | 

### Return type

[**ReportResultsResponse**](ReportResultsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

