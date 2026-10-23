# SessionProgressRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Coverage** | Pointer to **string** |  | [optional] 
**ProposedOnly** | Pointer to **bool** |  | [optional] 
**MovedThisSession** | Pointer to **bool** |  | [optional] 
**Tests** | Pointer to [**[]SessionProgressTest**](SessionProgressTest.md) |  | [optional] 
**Adopted** | Pointer to **bool** | Not in the caller&#39;s requested slice — added because a test this session executed links to it. Informational: adopted rows are excluded from summary/ready/next_actions (a broad validation run must not hold the session&#39;s readiness hostage to requirements it never touched). | [optional] 

## Methods

### NewSessionProgressRequirement

`func NewSessionProgressRequirement() *SessionProgressRequirement`

NewSessionProgressRequirement instantiates a new SessionProgressRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionProgressRequirementWithDefaults

`func NewSessionProgressRequirementWithDefaults() *SessionProgressRequirement`

NewSessionProgressRequirementWithDefaults instantiates a new SessionProgressRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *SessionProgressRequirement) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *SessionProgressRequirement) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *SessionProgressRequirement) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *SessionProgressRequirement) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetDisplay

`func (o *SessionProgressRequirement) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SessionProgressRequirement) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SessionProgressRequirement) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *SessionProgressRequirement) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTitle

`func (o *SessionProgressRequirement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SessionProgressRequirement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SessionProgressRequirement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SessionProgressRequirement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCoverage

`func (o *SessionProgressRequirement) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *SessionProgressRequirement) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *SessionProgressRequirement) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *SessionProgressRequirement) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetProposedOnly

`func (o *SessionProgressRequirement) GetProposedOnly() bool`

GetProposedOnly returns the ProposedOnly field if non-nil, zero value otherwise.

### GetProposedOnlyOk

`func (o *SessionProgressRequirement) GetProposedOnlyOk() (*bool, bool)`

GetProposedOnlyOk returns a tuple with the ProposedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedOnly

`func (o *SessionProgressRequirement) SetProposedOnly(v bool)`

SetProposedOnly sets ProposedOnly field to given value.

### HasProposedOnly

`func (o *SessionProgressRequirement) HasProposedOnly() bool`

HasProposedOnly returns a boolean if a field has been set.

### GetMovedThisSession

`func (o *SessionProgressRequirement) GetMovedThisSession() bool`

GetMovedThisSession returns the MovedThisSession field if non-nil, zero value otherwise.

### GetMovedThisSessionOk

`func (o *SessionProgressRequirement) GetMovedThisSessionOk() (*bool, bool)`

GetMovedThisSessionOk returns a tuple with the MovedThisSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedThisSession

`func (o *SessionProgressRequirement) SetMovedThisSession(v bool)`

SetMovedThisSession sets MovedThisSession field to given value.

### HasMovedThisSession

`func (o *SessionProgressRequirement) HasMovedThisSession() bool`

HasMovedThisSession returns a boolean if a field has been set.

### GetTests

`func (o *SessionProgressRequirement) GetTests() []SessionProgressTest`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *SessionProgressRequirement) GetTestsOk() (*[]SessionProgressTest, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *SessionProgressRequirement) SetTests(v []SessionProgressTest)`

SetTests sets Tests field to given value.

### HasTests

`func (o *SessionProgressRequirement) HasTests() bool`

HasTests returns a boolean if a field has been set.

### GetAdopted

`func (o *SessionProgressRequirement) GetAdopted() bool`

GetAdopted returns the Adopted field if non-nil, zero value otherwise.

### GetAdoptedOk

`func (o *SessionProgressRequirement) GetAdoptedOk() (*bool, bool)`

GetAdoptedOk returns a tuple with the Adopted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdopted

`func (o *SessionProgressRequirement) SetAdopted(v bool)`

SetAdopted sets Adopted field to given value.

### HasAdopted

`func (o *SessionProgressRequirement) HasAdopted() bool`

HasAdopted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


