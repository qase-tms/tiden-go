# Verdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**VerdictScope**](VerdictScope.md) |  | [optional] [default to VERDICT_SCOPE_UNSPECIFIED]
**ReleaseId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**BuildSha** | Pointer to **string** |  | [optional] 
**CommitSha** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**VerdictStatus**](VerdictStatus.md) |  | [optional] [default to VERDICT_STATUS_UNSPECIFIED]
**StateWatermark** | Pointer to **string** |  | [optional] 
**ComputedAt** | Pointer to **time.Time** |  | [optional] 
**InvalidatedAt** | Pointer to **time.Time** |  | [optional] 
**InvalidatedReason** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**[]ComponentResult**](ComponentResult.md) |  | [optional] 
**FixHints** | Pointer to [**[]FixHint**](FixHint.md) |  | [optional] 
**AcceptanceRequired** | Pointer to **bool** |  | [optional] 
**Subjects** | Pointer to [**[]SubjectResult**](SubjectResult.md) |  | [optional] 

## Methods

### NewVerdict

`func NewVerdict() *Verdict`

NewVerdict instantiates a new Verdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerdictWithDefaults

`func NewVerdictWithDefaults() *Verdict`

NewVerdictWithDefaults instantiates a new Verdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Verdict) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Verdict) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Verdict) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Verdict) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Verdict) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Verdict) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Verdict) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Verdict) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetScope

`func (o *Verdict) GetScope() VerdictScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Verdict) GetScopeOk() (*VerdictScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Verdict) SetScope(v VerdictScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Verdict) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetReleaseId

`func (o *Verdict) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *Verdict) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *Verdict) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *Verdict) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetBranchId

`func (o *Verdict) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *Verdict) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *Verdict) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *Verdict) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetBuildSha

`func (o *Verdict) GetBuildSha() string`

GetBuildSha returns the BuildSha field if non-nil, zero value otherwise.

### GetBuildShaOk

`func (o *Verdict) GetBuildShaOk() (*string, bool)`

GetBuildShaOk returns a tuple with the BuildSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSha

`func (o *Verdict) SetBuildSha(v string)`

SetBuildSha sets BuildSha field to given value.

### HasBuildSha

`func (o *Verdict) HasBuildSha() bool`

HasBuildSha returns a boolean if a field has been set.

### GetCommitSha

`func (o *Verdict) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *Verdict) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *Verdict) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *Verdict) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetStatus

`func (o *Verdict) GetStatus() VerdictStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Verdict) GetStatusOk() (*VerdictStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Verdict) SetStatus(v VerdictStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Verdict) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStateWatermark

`func (o *Verdict) GetStateWatermark() string`

GetStateWatermark returns the StateWatermark field if non-nil, zero value otherwise.

### GetStateWatermarkOk

`func (o *Verdict) GetStateWatermarkOk() (*string, bool)`

GetStateWatermarkOk returns a tuple with the StateWatermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateWatermark

`func (o *Verdict) SetStateWatermark(v string)`

SetStateWatermark sets StateWatermark field to given value.

### HasStateWatermark

`func (o *Verdict) HasStateWatermark() bool`

HasStateWatermark returns a boolean if a field has been set.

### GetComputedAt

`func (o *Verdict) GetComputedAt() time.Time`

GetComputedAt returns the ComputedAt field if non-nil, zero value otherwise.

### GetComputedAtOk

`func (o *Verdict) GetComputedAtOk() (*time.Time, bool)`

GetComputedAtOk returns a tuple with the ComputedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedAt

`func (o *Verdict) SetComputedAt(v time.Time)`

SetComputedAt sets ComputedAt field to given value.

### HasComputedAt

`func (o *Verdict) HasComputedAt() bool`

HasComputedAt returns a boolean if a field has been set.

### GetInvalidatedAt

`func (o *Verdict) GetInvalidatedAt() time.Time`

GetInvalidatedAt returns the InvalidatedAt field if non-nil, zero value otherwise.

### GetInvalidatedAtOk

`func (o *Verdict) GetInvalidatedAtOk() (*time.Time, bool)`

GetInvalidatedAtOk returns a tuple with the InvalidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedAt

`func (o *Verdict) SetInvalidatedAt(v time.Time)`

SetInvalidatedAt sets InvalidatedAt field to given value.

### HasInvalidatedAt

`func (o *Verdict) HasInvalidatedAt() bool`

HasInvalidatedAt returns a boolean if a field has been set.

### GetInvalidatedReason

`func (o *Verdict) GetInvalidatedReason() string`

GetInvalidatedReason returns the InvalidatedReason field if non-nil, zero value otherwise.

### GetInvalidatedReasonOk

`func (o *Verdict) GetInvalidatedReasonOk() (*string, bool)`

GetInvalidatedReasonOk returns a tuple with the InvalidatedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidatedReason

`func (o *Verdict) SetInvalidatedReason(v string)`

SetInvalidatedReason sets InvalidatedReason field to given value.

### HasInvalidatedReason

`func (o *Verdict) HasInvalidatedReason() bool`

HasInvalidatedReason returns a boolean if a field has been set.

### GetComponents

`func (o *Verdict) GetComponents() []ComponentResult`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Verdict) GetComponentsOk() (*[]ComponentResult, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Verdict) SetComponents(v []ComponentResult)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Verdict) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetFixHints

`func (o *Verdict) GetFixHints() []FixHint`

GetFixHints returns the FixHints field if non-nil, zero value otherwise.

### GetFixHintsOk

`func (o *Verdict) GetFixHintsOk() (*[]FixHint, bool)`

GetFixHintsOk returns a tuple with the FixHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixHints

`func (o *Verdict) SetFixHints(v []FixHint)`

SetFixHints sets FixHints field to given value.

### HasFixHints

`func (o *Verdict) HasFixHints() bool`

HasFixHints returns a boolean if a field has been set.

### GetAcceptanceRequired

`func (o *Verdict) GetAcceptanceRequired() bool`

GetAcceptanceRequired returns the AcceptanceRequired field if non-nil, zero value otherwise.

### GetAcceptanceRequiredOk

`func (o *Verdict) GetAcceptanceRequiredOk() (*bool, bool)`

GetAcceptanceRequiredOk returns a tuple with the AcceptanceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceRequired

`func (o *Verdict) SetAcceptanceRequired(v bool)`

SetAcceptanceRequired sets AcceptanceRequired field to given value.

### HasAcceptanceRequired

`func (o *Verdict) HasAcceptanceRequired() bool`

HasAcceptanceRequired returns a boolean if a field has been set.

### GetSubjects

`func (o *Verdict) GetSubjects() []SubjectResult`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *Verdict) GetSubjectsOk() (*[]SubjectResult, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *Verdict) SetSubjects(v []SubjectResult)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *Verdict) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


