# \QualityGateServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QualityGateServiceAcceptRisk**](QualityGateServiceAPI.md#QualityGateServiceAcceptRisk) | **Post** /v1/products/{productId}/quality-gate:accept-risk | Signs off the residual risk on a soft-signal verdict.
[**QualityGateServiceApproveRisk**](QualityGateServiceAPI.md#QualityGateServiceApproveRisk) | **Post** /v1/products/{productId}/quality-gate:approve-risk | Second-approver sign-off for a pending risk acceptance.
[**QualityGateServiceComputeVerdict**](QualityGateServiceAPI.md#QualityGateServiceComputeVerdict) | **Post** /v1/products/{productId}/quality-gate:compute | Computes and persists a quality-gate verdict.
[**QualityGateServiceGetSessionProgress**](QualityGateServiceAPI.md#QualityGateServiceGetSessionProgress) | **Post** /v1/products/{productId}/quality-gate:session-progress | Returns one intent session&#39;s per-requirement progress slice.
[**QualityGateServiceGetTraceability**](QualityGateServiceAPI.md#QualityGateServiceGetTraceability) | **Get** /v1/products/{productId}/quality-gate/traceability | Returns the traceability matrix behind a verdict.
[**QualityGateServiceGetVerdict**](QualityGateServiceAPI.md#QualityGateServiceGetVerdict) | **Get** /v1/products/{productId}/quality-gate | Fetches the latest verdict for a scope.
[**QualityGateServiceRecordSessionRiskAcceptances**](QualityGateServiceAPI.md#QualityGateServiceRecordSessionRiskAcceptances) | **Post** /v1/products/{productId}/quality-gate:session-acceptances | Records one intent session&#39;s risk acceptances and test deferrals.



## QualityGateServiceAcceptRisk

> AcceptRiskResponse QualityGateServiceAcceptRisk(ctx, productId).AcceptRiskBody(acceptRiskBody).Execute()

Signs off the residual risk on a soft-signal verdict.



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
	acceptRiskBody := *openapiclient.NewAcceptRiskBody() // AcceptRiskBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceAcceptRisk(context.Background(), productId).AcceptRiskBody(acceptRiskBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceAcceptRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceAcceptRisk`: AcceptRiskResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceAcceptRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceAcceptRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptRiskBody** | [**AcceptRiskBody**](AcceptRiskBody.md) |  | 

### Return type

[**AcceptRiskResponse**](AcceptRiskResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceApproveRisk

> ApproveRiskResponse QualityGateServiceApproveRisk(ctx, productId).ApproveRiskBody(approveRiskBody).Execute()

Second-approver sign-off for a pending risk acceptance.



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
	approveRiskBody := *openapiclient.NewApproveRiskBody() // ApproveRiskBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceApproveRisk(context.Background(), productId).ApproveRiskBody(approveRiskBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceApproveRisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceApproveRisk`: ApproveRiskResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceApproveRisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceApproveRiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approveRiskBody** | [**ApproveRiskBody**](ApproveRiskBody.md) |  | 

### Return type

[**ApproveRiskResponse**](ApproveRiskResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceComputeVerdict

> ComputeVerdictResponse QualityGateServiceComputeVerdict(ctx, productId).ComputeVerdictBody(computeVerdictBody).Execute()

Computes and persists a quality-gate verdict.



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
	computeVerdictBody := *openapiclient.NewComputeVerdictBody() // ComputeVerdictBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceComputeVerdict(context.Background(), productId).ComputeVerdictBody(computeVerdictBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceComputeVerdict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceComputeVerdict`: ComputeVerdictResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceComputeVerdict`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceComputeVerdictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computeVerdictBody** | [**ComputeVerdictBody**](ComputeVerdictBody.md) |  | 

### Return type

[**ComputeVerdictResponse**](ComputeVerdictResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceGetSessionProgress

> GetSessionProgressResponse QualityGateServiceGetSessionProgress(ctx, productId).GetSessionProgressBody(getSessionProgressBody).Execute()

Returns one intent session's per-requirement progress slice.



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
	getSessionProgressBody := *openapiclient.NewGetSessionProgressBody() // GetSessionProgressBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceGetSessionProgress(context.Background(), productId).GetSessionProgressBody(getSessionProgressBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceGetSessionProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceGetSessionProgress`: GetSessionProgressResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceGetSessionProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceGetSessionProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getSessionProgressBody** | [**GetSessionProgressBody**](GetSessionProgressBody.md) |  | 

### Return type

[**GetSessionProgressResponse**](GetSessionProgressResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceGetTraceability

> GetTraceabilityResponse QualityGateServiceGetTraceability(ctx, productId).Scope(scope).ReleaseId(releaseId).Branch(branch).SubjectType(subjectType).SubjectId(subjectId).Execute()

Returns the traceability matrix behind a verdict.



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
	scope := "scope_example" // string |  - VERDICT_SCOPE_RELEASE: canonical, against main-live entities of a release build  - VERDICT_SCOPE_BRANCH: pre-merge preview, against the merge-preview projection  - VERDICT_SCOPE_MAIN: current main, not tied to a release (latest exec per test) (optional) (default to "VERDICT_SCOPE_UNSPECIFIED")
	releaseId := "releaseId_example" // string |  (optional)
	branch := "branch_example" // string |  (optional)
	subjectType := "subjectType_example" // string | filter the matrix to one subject (\"component\"|\"feature\") (optional)
	subjectId := "subjectId_example" // string | paired with subject_type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceGetTraceability(context.Background(), productId).Scope(scope).ReleaseId(releaseId).Branch(branch).SubjectType(subjectType).SubjectId(subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceGetTraceability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceGetTraceability`: GetTraceabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceGetTraceability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceGetTraceabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** |  - VERDICT_SCOPE_RELEASE: canonical, against main-live entities of a release build  - VERDICT_SCOPE_BRANCH: pre-merge preview, against the merge-preview projection  - VERDICT_SCOPE_MAIN: current main, not tied to a release (latest exec per test) | [default to &quot;VERDICT_SCOPE_UNSPECIFIED&quot;]
 **releaseId** | **string** |  | 
 **branch** | **string** |  | 
 **subjectType** | **string** | filter the matrix to one subject (\&quot;component\&quot;|\&quot;feature\&quot;) | 
 **subjectId** | **string** | paired with subject_type | 

### Return type

[**GetTraceabilityResponse**](GetTraceabilityResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceGetVerdict

> GetVerdictResponse QualityGateServiceGetVerdict(ctx, productId).Scope(scope).ReleaseId(releaseId).Branch(branch).SubjectType(subjectType).SubjectId(subjectId).Execute()

Fetches the latest verdict for a scope.



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
	scope := "scope_example" // string |  - VERDICT_SCOPE_RELEASE: canonical, against main-live entities of a release build  - VERDICT_SCOPE_BRANCH: pre-merge preview, against the merge-preview projection  - VERDICT_SCOPE_MAIN: current main, not tied to a release (latest exec per test) (optional) (default to "VERDICT_SCOPE_UNSPECIFIED")
	releaseId := "releaseId_example" // string |  (optional)
	branch := "branch_example" // string |  (optional)
	subjectType := "subjectType_example" // string | filter the result to one subject (\"component\"|\"feature\"|\"product\") (optional)
	subjectId := "subjectId_example" // string | paired with subject_type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceGetVerdict(context.Background(), productId).Scope(scope).ReleaseId(releaseId).Branch(branch).SubjectType(subjectType).SubjectId(subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceGetVerdict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceGetVerdict`: GetVerdictResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceGetVerdict`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceGetVerdictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** |  - VERDICT_SCOPE_RELEASE: canonical, against main-live entities of a release build  - VERDICT_SCOPE_BRANCH: pre-merge preview, against the merge-preview projection  - VERDICT_SCOPE_MAIN: current main, not tied to a release (latest exec per test) | [default to &quot;VERDICT_SCOPE_UNSPECIFIED&quot;]
 **releaseId** | **string** |  | 
 **branch** | **string** |  | 
 **subjectType** | **string** | filter the result to one subject (\&quot;component\&quot;|\&quot;feature\&quot;|\&quot;product\&quot;) | 
 **subjectId** | **string** | paired with subject_type | 

### Return type

[**GetVerdictResponse**](GetVerdictResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QualityGateServiceRecordSessionRiskAcceptances

> RecordSessionRiskAcceptancesResponse QualityGateServiceRecordSessionRiskAcceptances(ctx, productId).RecordSessionRiskAcceptancesBody(recordSessionRiskAcceptancesBody).Execute()

Records one intent session's risk acceptances and test deferrals.



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
	recordSessionRiskAcceptancesBody := *openapiclient.NewRecordSessionRiskAcceptancesBody() // RecordSessionRiskAcceptancesBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityGateServiceAPI.QualityGateServiceRecordSessionRiskAcceptances(context.Background(), productId).RecordSessionRiskAcceptancesBody(recordSessionRiskAcceptancesBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityGateServiceAPI.QualityGateServiceRecordSessionRiskAcceptances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QualityGateServiceRecordSessionRiskAcceptances`: RecordSessionRiskAcceptancesResponse
	fmt.Fprintf(os.Stdout, "Response from `QualityGateServiceAPI.QualityGateServiceRecordSessionRiskAcceptances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQualityGateServiceRecordSessionRiskAcceptancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordSessionRiskAcceptancesBody** | [**RecordSessionRiskAcceptancesBody**](RecordSessionRiskAcceptancesBody.md) |  | 

### Return type

[**RecordSessionRiskAcceptancesResponse**](RecordSessionRiskAcceptancesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

