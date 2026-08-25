# UpdateUserOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliVerified** | Pointer to **bool** |  | [optional] 
**Dismissed** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**WizardStep** | Pointer to **string** | Wizard screen to resume at; one of \&quot;\&quot;, you-org, start, product, reqgen, repos, agents, plan, done. | [optional] 
**Answers** | Pointer to [**OnboardingAnswers**](OnboardingAnswers.md) |  | [optional] 

## Methods

### NewUpdateUserOnboardingRequest

`func NewUpdateUserOnboardingRequest() *UpdateUserOnboardingRequest`

NewUpdateUserOnboardingRequest instantiates a new UpdateUserOnboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserOnboardingRequestWithDefaults

`func NewUpdateUserOnboardingRequestWithDefaults() *UpdateUserOnboardingRequest`

NewUpdateUserOnboardingRequestWithDefaults instantiates a new UpdateUserOnboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliVerified

`func (o *UpdateUserOnboardingRequest) GetCliVerified() bool`

GetCliVerified returns the CliVerified field if non-nil, zero value otherwise.

### GetCliVerifiedOk

`func (o *UpdateUserOnboardingRequest) GetCliVerifiedOk() (*bool, bool)`

GetCliVerifiedOk returns a tuple with the CliVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVerified

`func (o *UpdateUserOnboardingRequest) SetCliVerified(v bool)`

SetCliVerified sets CliVerified field to given value.

### HasCliVerified

`func (o *UpdateUserOnboardingRequest) HasCliVerified() bool`

HasCliVerified returns a boolean if a field has been set.

### GetDismissed

`func (o *UpdateUserOnboardingRequest) GetDismissed() bool`

GetDismissed returns the Dismissed field if non-nil, zero value otherwise.

### GetDismissedOk

`func (o *UpdateUserOnboardingRequest) GetDismissedOk() (*bool, bool)`

GetDismissedOk returns a tuple with the Dismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissed

`func (o *UpdateUserOnboardingRequest) SetDismissed(v bool)`

SetDismissed sets Dismissed field to given value.

### HasDismissed

`func (o *UpdateUserOnboardingRequest) HasDismissed() bool`

HasDismissed returns a boolean if a field has been set.

### GetCompleted

`func (o *UpdateUserOnboardingRequest) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *UpdateUserOnboardingRequest) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *UpdateUserOnboardingRequest) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *UpdateUserOnboardingRequest) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetWizardStep

`func (o *UpdateUserOnboardingRequest) GetWizardStep() string`

GetWizardStep returns the WizardStep field if non-nil, zero value otherwise.

### GetWizardStepOk

`func (o *UpdateUserOnboardingRequest) GetWizardStepOk() (*string, bool)`

GetWizardStepOk returns a tuple with the WizardStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWizardStep

`func (o *UpdateUserOnboardingRequest) SetWizardStep(v string)`

SetWizardStep sets WizardStep field to given value.

### HasWizardStep

`func (o *UpdateUserOnboardingRequest) HasWizardStep() bool`

HasWizardStep returns a boolean if a field has been set.

### GetAnswers

`func (o *UpdateUserOnboardingRequest) GetAnswers() OnboardingAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *UpdateUserOnboardingRequest) GetAnswersOk() (*OnboardingAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *UpdateUserOnboardingRequest) SetAnswers(v OnboardingAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *UpdateUserOnboardingRequest) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


