# \AgentRetrievalServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentRetrievalServiceAttributeChangedFiles**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceAttributeChangedFiles) | **Post** /v1/products/{productId}/requirements/{requirementId}:attribute-changed-files | 
[**AgentRetrievalServiceDeclareRequirementEdgeIntent**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceDeclareRequirementEdgeIntent) | **Post** /v1/products/{productId}/requirement-edge-intents | 
[**AgentRetrievalServiceGetRequirementGraph**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetRequirementGraph) | **Get** /v1/products/{productId}/requirement-graph | 
[**AgentRetrievalServiceGetRequirementTestContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGetRequirementTestContext) | **Get** /v1/products/{productId}/requirements/{requirementId}/test-context | 
[**AgentRetrievalServiceGraphCoverageGaps**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceGraphCoverageGaps) | **Get** /v1/products/{productId}/requirements/graph-coverage-gaps | 
[**AgentRetrievalServiceListCoverageGaps**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceListCoverageGaps) | **Get** /v1/products/{productId}/coverage-gaps | 
[**AgentRetrievalServiceListRequirementAnchors**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceListRequirementAnchors) | **Get** /v1/products/{productId}/requirement-anchors | 
[**AgentRetrievalServicePrepareTestGenerationContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServicePrepareTestGenerationContext) | **Post** /v1/products/{productId}/test-generation-context:prepare | 
[**AgentRetrievalServiceRequirementImpact**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceRequirementImpact) | **Get** /v1/products/{productId}/requirements/impact | 
[**AgentRetrievalServiceRequirementNeighbors**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceRequirementNeighbors) | **Get** /v1/products/{productId}/requirements/{requirementId}/neighbors | 
[**AgentRetrievalServiceResolveFeatureContext**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceResolveFeatureContext) | **Get** /v1/products/{productId}/feature-context | 
[**AgentRetrievalServiceWriteRequirementEdge**](AgentRetrievalServiceAPI.md#AgentRetrievalServiceWriteRequirementEdge) | **Post** /v1/products/{productId}/requirement-edges | 



## AgentRetrievalServiceAttributeChangedFiles

> AttributeChangedFilesResponse AgentRetrievalServiceAttributeChangedFiles(ctx, productId, requirementId).AttributeChangedFilesBody(attributeChangedFilesBody).Execute()



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


## AgentRetrievalServiceGetRequirementGraph

> GetRequirementGraphResponse AgentRetrievalServiceGetRequirementGraph(ctx, productId).Execute()



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
	budget := int32(56) // int32 |  (optional)

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
 **budget** | **int32** |  | 

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

> RequirementImpactResponse AgentRetrievalServiceRequirementImpact(ctx, productId).RepoPaths(repoPaths).Depth(depth).EdgeTypes(edgeTypes).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentRetrievalServiceAPI.AgentRetrievalServiceRequirementImpact(context.Background(), productId).RepoPaths(repoPaths).Depth(depth).EdgeTypes(edgeTypes).Execute()
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

