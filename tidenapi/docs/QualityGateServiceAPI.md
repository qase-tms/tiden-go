# \QualityGateServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QualityGateServiceAcceptRisk**](QualityGateServiceAPI.md#QualityGateServiceAcceptRisk) | **Post** /v1/products/{productId}/quality-gate:accept-risk | Record a sign-off on a 🟡 (soft-signal) verdict so it becomes shippable. Hard-blocked (🔴) verdicts can&#39;t be accepted. High-severity components need a distinct 2nd approver (ApproveRisk); low-severity self-approve.
[**QualityGateServiceApproveRisk**](QualityGateServiceAPI.md#QualityGateServiceApproveRisk) | **Post** /v1/products/{productId}/quality-gate:approve-risk | Second-approver sign-off for a pending acceptance (must differ from the recorder).
[**QualityGateServiceComputeVerdict**](QualityGateServiceAPI.md#QualityGateServiceComputeVerdict) | **Post** /v1/products/{productId}/quality-gate:compute | Compute (or recompute) the verdict for a release or branch scope and persist an immutable snapshot. Side-effecting; the engine is idempotent on the current data state (CAS on publish).
[**QualityGateServiceGetTraceability**](QualityGateServiceAPI.md#QualityGateServiceGetTraceability) | **Get** /v1/products/{productId}/quality-gate/traceability | The traceability-matrix slice the verdict was computed over (req x case by component), for the matrix page and audit.
[**QualityGateServiceGetVerdict**](QualityGateServiceAPI.md#QualityGateServiceGetVerdict) | **Get** /v1/products/{productId}/quality-gate | Latest non-invalidated verdict for a (scope, ref). On no-go the agent reads the structured component/criterion breakdown + fix hints from here.



## QualityGateServiceAcceptRisk

> AcceptRiskResponse QualityGateServiceAcceptRisk(ctx, productId).AcceptRiskBody(acceptRiskBody).Execute()

Record a sign-off on a 🟡 (soft-signal) verdict so it becomes shippable. Hard-blocked (🔴) verdicts can't be accepted. High-severity components need a distinct 2nd approver (ApproveRisk); low-severity self-approve.

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

Second-approver sign-off for a pending acceptance (must differ from the recorder).

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

Compute (or recompute) the verdict for a release or branch scope and persist an immutable snapshot. Side-effecting; the engine is idempotent on the current data state (CAS on publish).

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


## QualityGateServiceGetTraceability

> GetTraceabilityResponse QualityGateServiceGetTraceability(ctx, productId).Scope(scope).ReleaseId(releaseId).Branch(branch).SubjectType(subjectType).SubjectId(subjectId).Execute()

The traceability-matrix slice the verdict was computed over (req x case by component), for the matrix page and audit.

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

Latest non-invalidated verdict for a (scope, ref). On no-go the agent reads the structured component/criterion breakdown + fix hints from here.

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

