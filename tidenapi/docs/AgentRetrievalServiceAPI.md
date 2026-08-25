# \AgentRetrievalServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRetrievalServiceAdvanceRepoWatermark**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceAdvanceRepoWatermark) | **Post** /v1/products/{productId}/repo-watermark:advance | Advances the drift watermark of one repository.
[**AgentRetrievalServiceAttributeChangedFiles**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceAttributeChangedFiles) | **Post** /v1/products/{productId}/requirements/{requirementId}:attribute-changed-files | Attributes a requirement&#39;s changed files to owning components.
[**AgentRetrievalServiceDeclareRequirementEdgeIntent**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceDeclareRequirementEdgeIntent) | **Post** /v1/products/{productId}/requirement-edge-intents | Records a deferred graph edge for endpoints not yet on main.
[**AgentRetrievalServiceGetIssueFixContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetIssueFixContext) | **Get** /v1/products/{productId}/issues/{issueId}/fix-context | Returns everything needed to fix one error, in a single call: the issue, its latest occurrence with symbolicated stack frames, the repository files those frames implicate, where the error is happening by environment, and — for each requirement those files implement — whether a test already covers it.
[**AgentRetrievalServiceGetRepoWatermark**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetRepoWatermark) | **Get** /v1/products/{productId}/repo-watermark | Returns the drift watermark of one repository.
[**AgentRetrievalServiceGetRequirementGraph**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetRequirementGraph) | **Get** /v1/products/{productId}/requirement-graph | Returns the product&#39;s full requirement graph.
[**AgentRetrievalServiceGetRequirementTestContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetRequirementTestContext) | **Get** /v1/products/{productId}/requirements/{requirementId}/test-context | Builds the full test-authoring context pack for one requirement.
[**AgentRetrievalServiceGraphCoverageGaps**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGraphCoverageGaps) | **Get** /v1/products/{productId}/requirements/graph-coverage-gaps | Filters a requirement set down to those without test coverage.
[**AgentRetrievalServiceListCoverageGaps**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceListCoverageGaps) | **Get** /v1/products/{productId}/coverage-gaps | Lists requirements with insufficient test coverage.
[**AgentRetrievalServiceListRequirementAnchors**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceListRequirementAnchors) | **Get** /v1/products/{productId}/requirement-anchors | Lists the branch-effective code anchors of all requirements.
[**AgentRetrievalServicePrepareTestGenerationContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServicePrepareTestGenerationContext) | **Post** /v1/products/{productId}/test-generation-context:prepare | Prepares a batched test-generation context for several requirements.
[**AgentRetrievalServiceRequirementImpact**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceRequirementImpact) | **Get** /v1/products/{productId}/requirements/impact | Computes the requirement blast radius of a set of changed files.
[**AgentRetrievalServiceRequirementNeighbors**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceRequirementNeighbors) | **Get** /v1/products/{productId}/requirements/{requirementId}/neighbors | Lists the graph neighbors of one requirement.
[**AgentRetrievalServiceResolveFeatureContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceResolveFeatureContext) | **Get** /v1/products/{productId}/feature-context | Resolves a coding objective into feature-rooted requirement context.
[**AgentRetrievalServiceWriteRequirementEdge**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceWriteRequirementEdge) | **Post** /v1/products/{productId}/requirement-edges | Writes one semantic edge into the requirement graph.



## AgentRetrievalServiceAdvanceRepoWatermark

> AdvanceRepoWatermarkResponse AgentRetrievalServiceAdvanceRepoWatermark(ctx, productId).AdvanceRepoWatermarkBody(advanceRepoWatermarkBody).Execute()

Advances the drift watermark of one repository.



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
	advanceRepoWatermarkBody := *openapiclient.NewAdvanceRepoWatermarkBody() // AdvanceRepoWatermarkBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceAdvanceRepoWatermark(context.Background(), productId).AdvanceRepoWatermarkBody(advanceRepoWatermarkBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceAdvanceRepoWatermark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceAdvanceRepoWatermark`: AdvanceRepoWatermarkResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceAdvanceRepoWatermark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceAdvanceRepoWatermarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **advanceRepoWatermarkBody** | [**AdvanceRepoWatermarkBody**](AdvanceRepoWatermarkBody.md) |  | 

### Return type

[**AdvanceRepoWatermarkResponse**](AdvanceRepoWatermarkResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceAttributeChangedFiles

> AttributeChangedFilesResponse AgentRetrievalServiceAttributeChangedFiles(ctx, productId, requirementId).AttributeChangedFilesBody(attributeChangedFilesBody).Execute()

Attributes a requirement's changed files to owning components.



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
	requirementId := "requirementId_example" // string | 
	attributeChangedFilesBody := *openapiclient.NewAttributeChangedFilesBody() // AttributeChangedFilesBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceAttributeChangedFiles(context.Background(), productId, requirementId).AttributeChangedFilesBody(attributeChangedFilesBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceAttributeChangedFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceAttributeChangedFiles`: AttributeChangedFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceAttributeChangedFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**requirementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceAttributeChangedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attributeChangedFilesBody** | [**AttributeChangedFilesBody**](AttributeChangedFilesBody.md) |  | 

### Return type

[**AttributeChangedFilesResponse**](AttributeChangedFilesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceDeclareRequirementEdgeIntent

> DeclareRequirementEdgeIntentResponse AgentRetrievalServiceDeclareRequirementEdgeIntent(ctx, productId).DeclareRequirementEdgeIntentBody(declareRequirementEdgeIntentBody).Execute()

Records a deferred graph edge for endpoints not yet on main.



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
	declareRequirementEdgeIntentBody := *openapiclient.NewDeclareRequirementEdgeIntentBody() // DeclareRequirementEdgeIntentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceDeclareRequirementEdgeIntent(context.Background(), productId).DeclareRequirementEdgeIntentBody(declareRequirementEdgeIntentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceDeclareRequirementEdgeIntent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceDeclareRequirementEdgeIntent`: DeclareRequirementEdgeIntentResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceDeclareRequirementEdgeIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceDeclareRequirementEdgeIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **declareRequirementEdgeIntentBody** | [**DeclareRequirementEdgeIntentBody**](DeclareRequirementEdgeIntentBody.md) |  | 

### Return type

[**DeclareRequirementEdgeIntentResponse**](DeclareRequirementEdgeIntentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceGetIssueFixContext

> GetIssueFixContextResponse AgentRetrievalServiceGetIssueFixContext(ctx, productId, issueId).Branch(branch).MaxFrames(maxFrames).Execute()

Returns everything needed to fix one error, in a single call: the issue, its latest occurrence with symbolicated stack frames, the repository files those frames implicate, where the error is happening by environment, and — for each requirement those files implement — whether a test already covers it.



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
	issueId := "issueId_example" // string | 
	branch := "branch_example" // string | branch scopes the requirement lookup to a branch's effective view. \"\" = main. (optional)
	maxFrames := int32(56) // int32 | max_frames bounds how many stack frames come back. <= 0 uses the server default (10); the cap is 50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceGetIssueFixContext(context.Background(), productId, issueId).Branch(branch).MaxFrames(maxFrames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceGetIssueFixContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceGetIssueFixContext`: GetIssueFixContextResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceGetIssueFixContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**issueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceGetIssueFixContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **string** | branch scopes the requirement lookup to a branch&#39;s effective view. \&quot;\&quot; &#x3D; main. | 
 **maxFrames** | **int32** | max_frames bounds how many stack frames come back. &lt;&#x3D; 0 uses the server default (10); the cap is 50. | 

### Return type

[**GetIssueFixContextResponse**](GetIssueFixContextResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceGetRepoWatermark

> GetRepoWatermarkResponse AgentRetrievalServiceGetRepoWatermark(ctx, productId).Repository(repository).Execute()

Returns the drift watermark of one repository.



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
	repository := "repository_example" // string | Canonical repo id (\"github.com/org/repo\" — the components.repository format), never a local path and never a clone URL. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceGetRepoWatermark(context.Background(), productId).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRepoWatermark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceGetRepoWatermark`: GetRepoWatermarkResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRepoWatermark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceGetRepoWatermarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | Canonical repo id (\&quot;github.com/org/repo\&quot; — the components.repository format), never a local path and never a clone URL. | 

### Return type

[**GetRepoWatermarkResponse**](GetRepoWatermarkResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceGetRequirementGraph

> GetRequirementGraphResponse AgentRetrievalServiceGetRequirementGraph(ctx, productId).Execute()

Returns the product's full requirement graph.



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
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementGraph(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceGetRequirementGraph`: GetRequirementGraphResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceGetRequirementGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRequirementGraphResponse**](GetRequirementGraphResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceGetRequirementTestContext

> GetRequirementTestContextResponse AgentRetrievalServiceGetRequirementTestContext(ctx, productId, requirementId).Branch(branch).Budget(budget).Execute()

Builds the full test-authoring context pack for one requirement.



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
	requirementId := "requirementId_example" // string | 
	branch := "branch_example" // string |  (optional)
	budget := int32(56) // int32 | budget bounds the pack's approximate token size: smaller budgets shrink per-list limits and trim long excerpts; <= 0 uses server defaults. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementTestContext(context.Background(), productId, requirementId).Branch(branch).Budget(budget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementTestContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceGetRequirementTestContext`: GetRequirementTestContextResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceGetRequirementTestContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**requirementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceGetRequirementTestContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **string** |  | 
 **budget** | **int32** | budget bounds the pack&#39;s approximate token size: smaller budgets shrink per-list limits and trim long excerpts; &lt;&#x3D; 0 uses server defaults. | 

### Return type

[**GetRequirementTestContextResponse**](GetRequirementTestContextResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceGraphCoverageGaps

> GraphCoverageGapsResponse AgentRetrievalServiceGraphCoverageGaps(ctx, productId).RequirementIds(requirementIds).Execute()

Filters a requirement set down to those without test coverage.



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
	requirementIds := []string{"Inner_example"} // []string | requirement_ids is the set to check for coverage. Typically the result of RequirementImpact or RequirementNeighbors. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceGraphCoverageGaps(context.Background(), productId).RequirementIds(requirementIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceGraphCoverageGaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceGraphCoverageGaps`: GraphCoverageGapsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceGraphCoverageGaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceGraphCoverageGapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requirementIds** | **[]string** | requirement_ids is the set to check for coverage. Typically the result of RequirementImpact or RequirementNeighbors. | 

### Return type

[**GraphCoverageGapsResponse**](GraphCoverageGapsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceListCoverageGaps

> ListCoverageGapsResponse AgentRetrievalServiceListCoverageGaps(ctx, productId).Branch(branch).CoverageStatuses(coverageStatuses).ComponentId(componentId).Query(query).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).RootRequirementId(rootRequirementId).Execute()

Lists requirements with insufficient test coverage.



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
	branch := "branch_example" // string |  (optional)
	coverageStatuses := []string{"Inner_example"} // []string |  (optional)
	componentId := "componentId_example" // string |  (optional)
	query := "query_example" // string |  (optional)
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)
	rootRequirementId := "rootRequirementId_example" // string | scope gaps to a feature subtree (root + descendants) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceListCoverageGaps(context.Background(), productId).Branch(branch).CoverageStatuses(coverageStatuses).ComponentId(componentId).Query(query).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).RootRequirementId(rootRequirementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceListCoverageGaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceListCoverageGaps`: ListCoverageGapsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceListCoverageGaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceListCoverageGapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** |  | 
 **coverageStatuses** | **[]string** |  | 
 **componentId** | **string** |  | 
 **query** | **string** |  | 
 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **rootRequirementId** | **string** | scope gaps to a feature subtree (root + descendants) | 

### Return type

[**ListCoverageGapsResponse**](ListCoverageGapsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceListRequirementAnchors

> ListRequirementAnchorsResponse AgentRetrievalServiceListRequirementAnchors(ctx, productId).Branch(branch).Execute()

Lists the branch-effective code anchors of all requirements.



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
	branch := "branch_example" // string | empty = main (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceListRequirementAnchors(context.Background(), productId).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceListRequirementAnchors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceListRequirementAnchors`: ListRequirementAnchorsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceListRequirementAnchors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceListRequirementAnchorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | empty &#x3D; main | 

### Return type

[**ListRequirementAnchorsResponse**](ListRequirementAnchorsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServicePrepareTestGenerationContext

> PrepareTestGenerationContextResponse AgentRetrievalServicePrepareTestGenerationContext(ctx, productId).PrepareTestGenerationContextBody(prepareTestGenerationContextBody).Execute()

Prepares a batched test-generation context for several requirements.



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
	prepareTestGenerationContextBody := *openapiclient.NewPrepareTestGenerationContextBody() // PrepareTestGenerationContextBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServicePrepareTestGenerationContext(context.Background(), productId).PrepareTestGenerationContextBody(prepareTestGenerationContextBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServicePrepareTestGenerationContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServicePrepareTestGenerationContext`: PrepareTestGenerationContextResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServicePrepareTestGenerationContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServicePrepareTestGenerationContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prepareTestGenerationContextBody** | [**PrepareTestGenerationContextBody**](PrepareTestGenerationContextBody.md) |  | 

### Return type

[**PrepareTestGenerationContextResponse**](PrepareTestGenerationContextResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceRequirementImpact

> RequirementImpactResponse AgentRetrievalServiceRequirementImpact(ctx, productId).RepoPaths(repoPaths).Depth(depth).EdgeTypes(edgeTypes).Repository(repository).MinConfidence(minConfidence).Execute()

Computes the requirement blast radius of a set of changed files.



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
	repoPaths := []string{"Inner_example"} // []string | repo_paths is the set of changed file paths (e.g. from a merged PR). The backend resolves these to seeded requirement IDs via requirement_sources. (optional)
	depth := int32(56) // int32 | depth controls how many hops the graph traversal expands beyond the seeds. Defaults to 3 on the server if <= 0. (optional)
	edgeTypes := []string{"Inner_example"} // []string | edge_types filters which edge types to traverse. Empty = all canonical types. (optional)
	repository := "repository_example" // string | repository scopes repo_paths to one repository: the canonical repo id (e.g. \"github.com/acme/backend\") OR a local checkout alias resolved via component repository_aliases — same semantics as ChangedFile.repository.  Anchors carry only a repo-relative path, so identical paths in different repositories (\".github/workflows/ci.yml\", \"Makefile\", \"CLAUDE.md\") collide. When set, a seed is kept only if its requirement's component resolves to this repository; requirements with no component still seed (fail-open) and are counted in ImpactCoverage.unverified_repository_seeds.  Empty = no repository filtering (pre-existing behaviour). (optional)
	minConfidence := float64(1.2) // float64 | min_confidence bounds which edges the traversal may step onto: a NULL confidence always passes (parent edges carry none, so the requirement tree is never pruned), and a derived edge (co_anchored/covers, confidence = 1/fan-out) below the floor is not admitted. Default 0 = no floor, the pre-existing unbounded behaviour — every caller that omits this field sees byte-identical results to before it existed. A caller that wants to bound a hub-file's fan-out (e.g. the intent-loop close gate) sets it explicitly; impact-analysis callers that want the deliberately broad radius leave it at 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementImpact(context.Background(), productId).RepoPaths(repoPaths).Depth(depth).EdgeTypes(edgeTypes).Repository(repository).MinConfidence(minConfidence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementImpact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceRequirementImpact`: RequirementImpactResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceRequirementImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repoPaths** | **[]string** | repo_paths is the set of changed file paths (e.g. from a merged PR). The backend resolves these to seeded requirement IDs via requirement_sources. | 
 **depth** | **int32** | depth controls how many hops the graph traversal expands beyond the seeds. Defaults to 3 on the server if &lt;&#x3D; 0. | 
 **edgeTypes** | **[]string** | edge_types filters which edge types to traverse. Empty &#x3D; all canonical types. | 
 **repository** | **string** | repository scopes repo_paths to one repository: the canonical repo id (e.g. \&quot;github.com/acme/backend\&quot;) OR a local checkout alias resolved via component repository_aliases — same semantics as ChangedFile.repository.  Anchors carry only a repo-relative path, so identical paths in different repositories (\&quot;.github/workflows/ci.yml\&quot;, \&quot;Makefile\&quot;, \&quot;CLAUDE.md\&quot;) collide. When set, a seed is kept only if its requirement&#39;s component resolves to this repository; requirements with no component still seed (fail-open) and are counted in ImpactCoverage.unverified_repository_seeds.  Empty &#x3D; no repository filtering (pre-existing behaviour). | 
 **minConfidence** | **float64** | min_confidence bounds which edges the traversal may step onto: a NULL confidence always passes (parent edges carry none, so the requirement tree is never pruned), and a derived edge (co_anchored/covers, confidence &#x3D; 1/fan-out) below the floor is not admitted. Default 0 &#x3D; no floor, the pre-existing unbounded behaviour — every caller that omits this field sees byte-identical results to before it existed. A caller that wants to bound a hub-file&#39;s fan-out (e.g. the intent-loop close gate) sets it explicitly; impact-analysis callers that want the deliberately broad radius leave it at 0. | 

### Return type

[**RequirementImpactResponse**](RequirementImpactResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceRequirementNeighbors

> RequirementNeighborsResponse AgentRetrievalServiceRequirementNeighbors(ctx, productId, requirementId).Depth(depth).EdgeTypes(edgeTypes).Execute()

Lists the graph neighbors of one requirement.



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
	requirementId := "requirementId_example" // string | 
	depth := int32(56) // int32 | depth controls how many hops to traverse (default 1 if <= 0). (optional)
	edgeTypes := []string{"Inner_example"} // []string | edge_types filters which edge types to traverse. Empty = all canonical types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementNeighbors(context.Background(), productId, requirementId).Depth(depth).EdgeTypes(edgeTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceRequirementNeighbors`: RequirementNeighborsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 
**requirementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceRequirementNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **depth** | **int32** | depth controls how many hops to traverse (default 1 if &lt;&#x3D; 0). | 
 **edgeTypes** | **[]string** | edge_types filters which edge types to traverse. Empty &#x3D; all canonical types. | 

### Return type

[**RequirementNeighborsResponse**](RequirementNeighborsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceResolveFeatureContext

> ResolveFeatureContextResponse AgentRetrievalServiceResolveFeatureContext(ctx, productId).Branch(branch).Text(text).RepoPaths(repoPaths).K(k).Execute()

Resolves a coding objective into feature-rooted requirement context.



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
	branch := "branch_example" // string | empty = main (optional)
	text := "text_example" // string | text is the free-text coding objective to resolve into feature context. (optional)
	repoPaths := []string{"Inner_example"} // []string | repo_paths are optional changed file paths that seed anchor retrieval. (optional)
	k := int32(56) // int32 | k bounds each retrieval signal's breadth (server default when <= 0). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceResolveFeatureContext(context.Background(), productId).Branch(branch).Text(text).RepoPaths(repoPaths).K(k).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceResolveFeatureContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceResolveFeatureContext`: ResolveFeatureContextResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceResolveFeatureContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceResolveFeatureContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | empty &#x3D; main | 
 **text** | **string** | text is the free-text coding objective to resolve into feature context. | 
 **repoPaths** | **[]string** | repo_paths are optional changed file paths that seed anchor retrieval. | 
 **k** | **int32** | k bounds each retrieval signal&#39;s breadth (server default when &lt;&#x3D; 0). | 

### Return type

[**ResolveFeatureContextResponse**](ResolveFeatureContextResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentRetrievalServiceWriteRequirementEdge

> WriteRequirementEdgeResponse AgentRetrievalServiceWriteRequirementEdge(ctx, productId).WriteRequirementEdgeBody(writeRequirementEdgeBody).Execute()

Writes one semantic edge into the requirement graph.



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
	writeRequirementEdgeBody := *openapiclient.NewWriteRequirementEdgeBody() // WriteRequirementEdgeBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceWriteRequirementEdge(context.Background(), productId).WriteRequirementEdgeBody(writeRequirementEdgeBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentRetrievalServiceAPI.AgentRetrievalServiceWriteRequirementEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentRetrievalServiceWriteRequirementEdge`: WriteRequirementEdgeResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentRetrievalServiceAPI.AgentRetrievalServiceWriteRequirementEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentRetrievalServiceWriteRequirementEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writeRequirementEdgeBody** | [**WriteRequirementEdgeBody**](WriteRequirementEdgeBody.md) |  | 

### Return type

[**WriteRequirementEdgeResponse**](WriteRequirementEdgeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

