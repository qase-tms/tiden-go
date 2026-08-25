# RecordSessionRiskAcceptancesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** | requirement_id (v2, optional): the session&#39;s draft requirement (branch-local on intent_branch) — the legacy write path, kept for older CLIs. ABSENT means the session record (intent_sessions) must already exist; acceptances/deferrals are then written to intent_session_judgements instead of onto a draft&#39;s provenance. | [optional] 
**IntentBranch** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Acceptances** | Pointer to [**[]SessionRiskAcceptance**](SessionRiskAcceptance.md) |  | [optional] 
**ProposedTestRequirementRefs** | Pointer to **[]string** | The deferral half of the ledger: requirements whose missing test is handed to a next session rather than risk-accepted. Same ref grammar as above; they collapse into one \&quot;test_deferral\&quot; record. | [optional] 

## Methods

### NewRecordSessionRiskAcceptancesBody

`func NewRecordSessionRiskAcceptancesBody() *RecordSessionRiskAcceptancesBody`

NewRecordSessionRiskAcceptancesBody instantiates a new RecordSessionRiskAcceptancesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSessionRiskAcceptancesBodyWithDefaults

`func NewRecordSessionRiskAcceptancesBodyWithDefaults() *RecordSessionRiskAcceptancesBody`

NewRecordSessionRiskAcceptancesBodyWithDefaults instantiates a new RecordSessionRiskAcceptancesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *RecordSessionRiskAcceptancesBody) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *RecordSessionRiskAcceptancesBody) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *RecordSessionRiskAcceptancesBody) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *RecordSessionRiskAcceptancesBody) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetIntentBranch

`func (o *RecordSessionRiskAcceptancesBody) GetIntentBranch() string`

GetIntentBranch returns the IntentBranch field if non-nil, zero value otherwise.

### GetIntentBranchOk

`func (o *RecordSessionRiskAcceptancesBody) GetIntentBranchOk() (*string, bool)`

GetIntentBranchOk returns a tuple with the IntentBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentBranch

`func (o *RecordSessionRiskAcceptancesBody) SetIntentBranch(v string)`

SetIntentBranch sets IntentBranch field to given value.

### HasIntentBranch

`func (o *RecordSessionRiskAcceptancesBody) HasIntentBranch() bool`

HasIntentBranch returns a boolean if a field has been set.

### GetSessionId

`func (o *RecordSessionRiskAcceptancesBody) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *RecordSessionRiskAcceptancesBody) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *RecordSessionRiskAcceptancesBody) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *RecordSessionRiskAcceptancesBody) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetAcceptances

`func (o *RecordSessionRiskAcceptancesBody) GetAcceptances() []SessionRiskAcceptance`

GetAcceptances returns the Acceptances field if non-nil, zero value otherwise.

### GetAcceptancesOk

`func (o *RecordSessionRiskAcceptancesBody) GetAcceptancesOk() (*[]SessionRiskAcceptance, bool)`

GetAcceptancesOk returns a tuple with the Acceptances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptances

`func (o *RecordSessionRiskAcceptancesBody) SetAcceptances(v []SessionRiskAcceptance)`

SetAcceptances sets Acceptances field to given value.

### HasAcceptances

`func (o *RecordSessionRiskAcceptancesBody) HasAcceptances() bool`

HasAcceptances returns a boolean if a field has been set.

### GetProposedTestRequirementRefs

`func (o *RecordSessionRiskAcceptancesBody) GetProposedTestRequirementRefs() []string`

GetProposedTestRequirementRefs returns the ProposedTestRequirementRefs field if non-nil, zero value otherwise.

### GetProposedTestRequirementRefsOk

`func (o *RecordSessionRiskAcceptancesBody) GetProposedTestRequirementRefsOk() (*[]string, bool)`

GetProposedTestRequirementRefsOk returns a tuple with the ProposedTestRequirementRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTestRequirementRefs

`func (o *RecordSessionRiskAcceptancesBody) SetProposedTestRequirementRefs(v []string)`

SetProposedTestRequirementRefs sets ProposedTestRequirementRefs field to given value.

### HasProposedTestRequirementRefs

`func (o *RecordSessionRiskAcceptancesBody) HasProposedTestRequirementRefs() bool`

HasProposedTestRequirementRefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


