# UpdateUserOnboardingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Onboarding** | Pointer to [**UserOnboardingState**](UserOnboardingState.md) |  | [optional] 

## Methods

### NewUpdateUserOnboardingResponse

`func NewUpdateUserOnboardingResponse() *UpdateUserOnboardingResponse`

NewUpdateUserOnboardingResponse instantiates a new UpdateUserOnboardingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserOnboardingResponseWithDefaults

`func NewUpdateUserOnboardingResponseWithDefaults() *UpdateUserOnboardingResponse`

NewUpdateUserOnboardingResponseWithDefaults instantiates a new UpdateUserOnboardingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnboarding

`func (o *UpdateUserOnboardingResponse) GetOnboarding() UserOnboardingState`

GetOnboarding returns the Onboarding field if non-nil, zero value otherwise.

### GetOnboardingOk

`func (o *UpdateUserOnboardingResponse) GetOnboardingOk() (*UserOnboardingState, bool)`

GetOnboardingOk returns a tuple with the Onboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboarding

`func (o *UpdateUserOnboardingResponse) SetOnboarding(v UserOnboardingState)`

SetOnboarding sets Onboarding field to given value.

### HasOnboarding

`func (o *UpdateUserOnboardingResponse) HasOnboarding() bool`

HasOnboarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


