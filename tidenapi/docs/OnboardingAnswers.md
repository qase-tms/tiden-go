# OnboardingAnswers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | How the user described themselves: engineer | leader | qa | pm | other. | [optional] 
**ReqgenChoice** | Pointer to **string** | How requirements get seeded: empty | own | tiden. | [optional] 
**ReqgenSource** | Pointer to **string** | Where Tiden should read requirements from: github | docs | tracker. | [optional] 
**DocsUrl** | Pointer to **string** | Docs URL supplied when reqgen_source is \&quot;docs\&quot; (max 2000 chars). | [optional] 
**ProductId** | Pointer to **string** | The product created during the wizard, if any. | [optional] 

## Methods

### NewOnboardingAnswers

`func NewOnboardingAnswers() *OnboardingAnswers`

NewOnboardingAnswers instantiates a new OnboardingAnswers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingAnswersWithDefaults

`func NewOnboardingAnswersWithDefaults() *OnboardingAnswers`

NewOnboardingAnswersWithDefaults instantiates a new OnboardingAnswers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OnboardingAnswers) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OnboardingAnswers) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OnboardingAnswers) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OnboardingAnswers) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetReqgenChoice

`func (o *OnboardingAnswers) GetReqgenChoice() string`

GetReqgenChoice returns the ReqgenChoice field if non-nil, zero value otherwise.

### GetReqgenChoiceOk

`func (o *OnboardingAnswers) GetReqgenChoiceOk() (*string, bool)`

GetReqgenChoiceOk returns a tuple with the ReqgenChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqgenChoice

`func (o *OnboardingAnswers) SetReqgenChoice(v string)`

SetReqgenChoice sets ReqgenChoice field to given value.

### HasReqgenChoice

`func (o *OnboardingAnswers) HasReqgenChoice() bool`

HasReqgenChoice returns a boolean if a field has been set.

### GetReqgenSource

`func (o *OnboardingAnswers) GetReqgenSource() string`

GetReqgenSource returns the ReqgenSource field if non-nil, zero value otherwise.

### GetReqgenSourceOk

`func (o *OnboardingAnswers) GetReqgenSourceOk() (*string, bool)`

GetReqgenSourceOk returns a tuple with the ReqgenSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqgenSource

`func (o *OnboardingAnswers) SetReqgenSource(v string)`

SetReqgenSource sets ReqgenSource field to given value.

### HasReqgenSource

`func (o *OnboardingAnswers) HasReqgenSource() bool`

HasReqgenSource returns a boolean if a field has been set.

### GetDocsUrl

`func (o *OnboardingAnswers) GetDocsUrl() string`

GetDocsUrl returns the DocsUrl field if non-nil, zero value otherwise.

### GetDocsUrlOk

`func (o *OnboardingAnswers) GetDocsUrlOk() (*string, bool)`

GetDocsUrlOk returns a tuple with the DocsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsUrl

`func (o *OnboardingAnswers) SetDocsUrl(v string)`

SetDocsUrl sets DocsUrl field to given value.

### HasDocsUrl

`func (o *OnboardingAnswers) HasDocsUrl() bool`

HasDocsUrl returns a boolean if a field has been set.

### GetProductId

`func (o *OnboardingAnswers) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OnboardingAnswers) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OnboardingAnswers) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OnboardingAnswers) HasProductId() bool`

HasProductId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


