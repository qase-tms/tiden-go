# \TestServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestServiceCreateTest**](TestServiceAPI.md#TestServiceCreateTest) | **Post** /v1/products/{productId}/tests | Creates a test suite or case.
[**TestServiceDeleteTest**](TestServiceAPI.md#TestServiceDeleteTest) | **Delete** /v1/tests/{id} | Deletes a test.
[**TestServiceDeriveTestLinks**](TestServiceAPI.md#TestServiceDeriveTestLinks) | **Post** /v1/products/{productId}/test-links:derive | Derives test-requirement links from shared file anchors.
[**TestServiceGetTest**](TestServiceAPI.md#TestServiceGetTest) | **Get** /v1/tests/{id} | Fetches one test by id.
[**TestServiceIngestTests**](TestServiceAPI.md#TestServiceIngestTests) | **Post** /v1/products/{productId}/tests:ingest | Batch-upserts tests from a reporter (live-documentation ingest).
[**TestServiceLinkRequirement**](TestServiceAPI.md#TestServiceLinkRequirement) | **Post** /v1/tests/{testId}/links | Links a test case to a requirement.
[**TestServiceListBranchLinkProposals**](TestServiceAPI.md#TestServiceListBranchLinkProposals) | **Get** /v1/branches/{branchId}/link-proposals | Lists a branch&#39;s test-requirement link proposals.
[**TestServiceListLinks**](TestServiceAPI.md#TestServiceListLinks) | **Get** /v1/tests/{testId}/links | Lists a test&#39;s requirement links.
[**TestServiceListTests**](TestServiceAPI.md#TestServiceListTests) | **Get** /v1/products/{productId}/tests | Lists a product&#39;s tests.
[**TestServiceReviewBranchLinkProposals**](TestServiceAPI.md#TestServiceReviewBranchLinkProposals) | **Post** /v1/branches/{branchId}/link-proposals:review | Accepts or rejects branch link proposals.
[**TestServiceUnlinkRequirement**](TestServiceAPI.md#TestServiceUnlinkRequirement) | **Delete** /v1/tests/{testId}/links/{requirementId} | Removes a test-requirement link.
[**TestServiceUpdateTest**](TestServiceAPI.md#TestServiceUpdateTest) | **Put** /v1/tests/{id} | Updates a test suite or case.



## TestServiceCreateTest

> CreateTestResponse TestServiceCreateTest(ctx, productId).CreateTestBody(createTestBody).Execute()

Creates a test suite or case.



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
	createTestBody := *openapiclient.NewCreateTestBody() // CreateTestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceCreateTest(context.Background(), productId).CreateTestBody(createTestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceCreateTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceCreateTest`: CreateTestResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceCreateTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceCreateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTestBody** | [**CreateTestBody**](CreateTestBody.md) |  | 

### Return type

[**CreateTestResponse**](CreateTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceDeleteTest

> map[string]interface{} TestServiceDeleteTest(ctx, id).Branch(branch).Execute()

Deletes a test.



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
	branch := "branch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceDeleteTest(context.Background(), id).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceDeleteTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceDeleteTest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceDeleteTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceDeleteTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** |  | 

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


## TestServiceDeriveTestLinks

> DeriveTestLinksResponse TestServiceDeriveTestLinks(ctx, productId).Body(body).Execute()

Derives test-requirement links from shared file anchors.



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceDeriveTestLinks(context.Background(), productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceDeriveTestLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceDeriveTestLinks`: DeriveTestLinksResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceDeriveTestLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceDeriveTestLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**DeriveTestLinksResponse**](DeriveTestLinksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceGetTest

> GetTestResponse TestServiceGetTest(ctx, id).Execute()

Fetches one test by id.



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
	resp, r, err := apiClient.TestServiceAPI.TestServiceGetTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceGetTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceGetTest`: GetTestResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceGetTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceGetTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTestResponse**](GetTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceIngestTests

> IngestTestsResponse TestServiceIngestTests(ctx, productId).IngestTestsBody(ingestTestsBody).Execute()

Batch-upserts tests from a reporter (live-documentation ingest).



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
	ingestTestsBody := *openapiclient.NewIngestTestsBody() // IngestTestsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceIngestTests(context.Background(), productId).IngestTestsBody(ingestTestsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceIngestTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceIngestTests`: IngestTestsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceIngestTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceIngestTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingestTestsBody** | [**IngestTestsBody**](IngestTestsBody.md) |  | 

### Return type

[**IngestTestsResponse**](IngestTestsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceLinkRequirement

> map[string]interface{} TestServiceLinkRequirement(ctx, testId).LinkRequirementBody(linkRequirementBody).Execute()

Links a test case to a requirement.



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
	testId := "testId_example" // string | 
	linkRequirementBody := *openapiclient.NewLinkRequirementBody() // LinkRequirementBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceLinkRequirement(context.Background(), testId).LinkRequirementBody(linkRequirementBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceLinkRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceLinkRequirement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceLinkRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceLinkRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkRequirementBody** | [**LinkRequirementBody**](LinkRequirementBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceListBranchLinkProposals

> ListBranchLinkProposalsResponse TestServiceListBranchLinkProposals(ctx, branchId).Statuses(statuses).Execute()

Lists a branch's test-requirement link proposals.



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
	branchId := "branchId_example" // string | 
	statuses := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceListBranchLinkProposals(context.Background(), branchId).Statuses(statuses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceListBranchLinkProposals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceListBranchLinkProposals`: ListBranchLinkProposalsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceListBranchLinkProposals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceListBranchLinkProposalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statuses** | **[]string** |  | 

### Return type

[**ListBranchLinkProposalsResponse**](ListBranchLinkProposalsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceListLinks

> ListLinksResponse TestServiceListLinks(ctx, testId).Branch(branch).Execute()

Lists a test's requirement links.



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
	testId := "testId_example" // string | 
	branch := "branch_example" // string | when set, links are read-only via COW (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceListLinks(context.Background(), testId).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceListLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceListLinks`: ListLinksResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceListLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceListLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | when set, links are read-only via COW | 

### Return type

[**ListLinksResponse**](ListLinksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceListTests

> ListTestsResponse TestServiceListTests(ctx, productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).Execute()

Lists a product's tests.



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
	paginationPageSize := int32(56) // int32 |  (optional)
	paginationPageToken := "paginationPageToken_example" // string |  (optional)
	branch := "branch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceListTests(context.Background(), productId).PaginationPageSize(paginationPageSize).PaginationPageToken(paginationPageToken).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceListTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceListTests`: ListTestsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceListTests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceListTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPageSize** | **int32** |  | 
 **paginationPageToken** | **string** |  | 
 **branch** | **string** |  | 

### Return type

[**ListTestsResponse**](ListTestsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceReviewBranchLinkProposals

> ReviewBranchLinkProposalsResponse TestServiceReviewBranchLinkProposals(ctx, branchId).ReviewBranchLinkProposalsBody(reviewBranchLinkProposalsBody).Execute()

Accepts or rejects branch link proposals.



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
	branchId := "branchId_example" // string | 
	reviewBranchLinkProposalsBody := *openapiclient.NewReviewBranchLinkProposalsBody() // ReviewBranchLinkProposalsBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceReviewBranchLinkProposals(context.Background(), branchId).ReviewBranchLinkProposalsBody(reviewBranchLinkProposalsBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceReviewBranchLinkProposals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceReviewBranchLinkProposals`: ReviewBranchLinkProposalsResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceReviewBranchLinkProposals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceReviewBranchLinkProposalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewBranchLinkProposalsBody** | [**ReviewBranchLinkProposalsBody**](ReviewBranchLinkProposalsBody.md) |  | 

### Return type

[**ReviewBranchLinkProposalsResponse**](ReviewBranchLinkProposalsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestServiceUnlinkRequirement

> map[string]interface{} TestServiceUnlinkRequirement(ctx, testId, requirementId).Branch(branch).Execute()

Removes a test-requirement link.



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
	testId := "testId_example" // string | 
	requirementId := "requirementId_example" // string | 
	branch := "branch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceUnlinkRequirement(context.Background(), testId, requirementId).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceUnlinkRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceUnlinkRequirement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceUnlinkRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **string** |  | 
**requirementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceUnlinkRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **string** |  | 

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


## TestServiceUpdateTest

> UpdateTestResponse TestServiceUpdateTest(ctx, id).UpdateTestBody(updateTestBody).Execute()

Updates a test suite or case.



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
	updateTestBody := *openapiclient.NewUpdateTestBody() // UpdateTestBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestServiceAPI.TestServiceUpdateTest(context.Background(), id).UpdateTestBody(updateTestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestServiceAPI.TestServiceUpdateTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestServiceUpdateTest`: UpdateTestResponse
	fmt.Fprintf(os.Stdout, "Response from `TestServiceAPI.TestServiceUpdateTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceUpdateTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTestBody** | [**UpdateTestBody**](UpdateTestBody.md) |  | 

### Return type

[**UpdateTestResponse**](UpdateTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

