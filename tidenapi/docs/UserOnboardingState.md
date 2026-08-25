# UserOnboardingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**CliVerifiedAt** | Pointer to **time.Time** |  | [optional] 
**DismissedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WizardStep** | Pointer to **string** | Current wizard screen; \&quot;\&quot; means the wizard was never started. | [optional] 
**Answers** | Pointer to [**OnboardingAnswers**](OnboardingAnswers.md) |  | [optional] 

## Methods

### NewUserOnboardingState

`func NewUserOnboardingState() *UserOnboardingState`

NewUserOnboardingState instantiates a new UserOnboardingState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOnboardingStateWithDefaults

`func NewUserOnboardingStateWithDefaults() *UserOnboardingState`

NewUserOnboardingStateWithDefaults instantiates a new UserOnboardingState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserOnboardingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserOnboardingState) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserOnboardingState) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserOnboardingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCliVerifiedAt

`func (o *UserOnboardingState) GetCliVerifiedAt() time.Time`

GetCliVerifiedAt returns the CliVerifiedAt field if non-nil, zero value otherwise.

### GetCliVerifiedAtOk

`func (o *UserOnboardingState) GetCliVerifiedAtOk() (*time.Time, bool)`

GetCliVerifiedAtOk returns a tuple with the CliVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVerifiedAt

`func (o *UserOnboardingState) SetCliVerifiedAt(v time.Time)`

SetCliVerifiedAt sets CliVerifiedAt field to given value.

### HasCliVerifiedAt

`func (o *UserOnboardingState) HasCliVerifiedAt() bool`

HasCliVerifiedAt returns a boolean if a field has been set.

### GetDismissedAt

`func (o *UserOnboardingState) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *UserOnboardingState) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *UserOnboardingState) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.

### HasDismissedAt

`func (o *UserOnboardingState) HasDismissedAt() bool`

HasDismissedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *UserOnboardingState) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *UserOnboardingState) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *UserOnboardingState) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *UserOnboardingState) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserOnboardingState) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserOnboardingState) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserOnboardingState) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserOnboardingState) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserOnboardingState) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserOnboardingState) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserOnboardingState) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserOnboardingState) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWizardStep

`func (o *UserOnboardingState) GetWizardStep() string`

GetWizardStep returns the WizardStep field if non-nil, zero value otherwise.

### GetWizardStepOk

`func (o *UserOnboardingState) GetWizardStepOk() (*string, bool)`

GetWizardStepOk returns a tuple with the WizardStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWizardStep

`func (o *UserOnboardingState) SetWizardStep(v string)`

SetWizardStep sets WizardStep field to given value.

### HasWizardStep

`func (o *UserOnboardingState) HasWizardStep() bool`

HasWizardStep returns a boolean if a field has been set.

### GetAnswers

`func (o *UserOnboardingState) GetAnswers() OnboardingAnswers`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *UserOnboardingState) GetAnswersOk() (*OnboardingAnswers, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *UserOnboardingState) SetAnswers(v OnboardingAnswers)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *UserOnboardingState) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


