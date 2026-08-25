# \AuthServiceAPI

All URIs are relative to *https://api.tiden.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthServiceGetCurrentUser**](AuthServiceAPI.md#AuthServiceGetCurrentUser) | **Get** /v1/auth/me | Returns the authenticated user (whoami).
[**AuthServiceUpdateUserOnboarding**](AuthServiceAPI.md#AuthServiceUpdateUserOnboarding) | **Put** /v1/auth/onboarding | Updates the caller&#39;s onboarding progress flags.



## AuthServiceGetCurrentUser

> GetCurrentUserResponse AuthServiceGetCurrentUser(ctx).Execute()

Returns the authenticated user (whoami).



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceGetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceGetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceGetCurrentUser`: GetCurrentUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceGetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceGetCurrentUserRequest struct via the builder pattern


### Return type

[**GetCurrentUserResponse**](GetCurrentUserResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceUpdateUserOnboarding

> UpdateUserOnboardingResponse AuthServiceUpdateUserOnboarding(ctx).UpdateUserOnboardingRequest(updateUserOnboardingRequest).Execute()

Updates the caller's onboarding progress flags.



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
	updateUserOnboardingRequest := *openapiclient.NewUpdateUserOnboardingRequest() // UpdateUserOnboardingRequest | UpdateUserOnboardingRequest latches onboarding flags: each true field stamps its timestamp; false leaves the flag unchanged. wizard_step and answers are ordinary mutable state: an unset field is left untouched, a set one replaces the stored value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceUpdateUserOnboarding(context.Background()).UpdateUserOnboardingRequest(updateUserOnboardingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceUpdateUserOnboarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceUpdateUserOnboarding`: UpdateUserOnboardingResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceUpdateUserOnboarding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceUpdateUserOnboardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserOnboardingRequest** | [**UpdateUserOnboardingRequest**](UpdateUserOnboardingRequest.md) | UpdateUserOnboardingRequest latches onboarding flags: each true field stamps its timestamp; false leaves the flag unchanged. wizard_step and answers are ordinary mutable state: an unset field is left untouched, a set one replaces the stored value. | 

### Return type

[**UpdateUserOnboardingResponse**](UpdateUserOnboardingResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

