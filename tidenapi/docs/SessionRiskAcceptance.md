# SessionRiskAcceptance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementRefs** | Pointer to **[]string** | Requirement refs exactly as the agent typed them: \&quot;&lt;CODE&gt;-&lt;N&gt;\&quot; (e.g. \&quot;FEED-49\&quot;) or a canonical lowercase UUID. The server resolves each against the intent branch&#39;s view to a canonical (main-twin) id and records BOTH — ids for machines, refs for the human reading merge-preview. | [optional] 
**Criterion** | Pointer to **string** |  | [optional] 
**Evidence** | Pointer to **string** |  | [optional] 
**FollowUp** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionRiskAcceptance

`func NewSessionRiskAcceptance() *SessionRiskAcceptance`

NewSessionRiskAcceptance instantiates a new SessionRiskAcceptance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRiskAcceptanceWithDefaults

`func NewSessionRiskAcceptanceWithDefaults() *SessionRiskAcceptance`

NewSessionRiskAcceptanceWithDefaults instantiates a new SessionRiskAcceptance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementRefs

`func (o *SessionRiskAcceptance) GetRequirementRefs() []string`

GetRequirementRefs returns the RequirementRefs field if non-nil, zero value otherwise.

### GetRequirementRefsOk

`func (o *SessionRiskAcceptance) GetRequirementRefsOk() (*[]string, bool)`

GetRequirementRefsOk returns a tuple with the RequirementRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementRefs

`func (o *SessionRiskAcceptance) SetRequirementRefs(v []string)`

SetRequirementRefs sets RequirementRefs field to given value.

### HasRequirementRefs

`func (o *SessionRiskAcceptance) HasRequirementRefs() bool`

HasRequirementRefs returns a boolean if a field has been set.

### GetCriterion

`func (o *SessionRiskAcceptance) GetCriterion() string`

GetCriterion returns the Criterion field if non-nil, zero value otherwise.

### GetCriterionOk

`func (o *SessionRiskAcceptance) GetCriterionOk() (*string, bool)`

GetCriterionOk returns a tuple with the Criterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterion

`func (o *SessionRiskAcceptance) SetCriterion(v string)`

SetCriterion sets Criterion field to given value.

### HasCriterion

`func (o *SessionRiskAcceptance) HasCriterion() bool`

HasCriterion returns a boolean if a field has been set.

### GetEvidence

`func (o *SessionRiskAcceptance) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *SessionRiskAcceptance) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *SessionRiskAcceptance) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *SessionRiskAcceptance) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetFollowUp

`func (o *SessionRiskAcceptance) GetFollowUp() string`

GetFollowUp returns the FollowUp field if non-nil, zero value otherwise.

### GetFollowUpOk

`func (o *SessionRiskAcceptance) GetFollowUpOk() (*string, bool)`

GetFollowUpOk returns a tuple with the FollowUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUp

`func (o *SessionRiskAcceptance) SetFollowUp(v string)`

SetFollowUp sets FollowUp field to given value.

### HasFollowUp

`func (o *SessionRiskAcceptance) HasFollowUp() bool`

HasFollowUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


