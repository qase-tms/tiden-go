# Test

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**SeqNum** | Pointer to **int32** | unset for suites | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Layer** | Pointer to **string** |  | [optional] 
**Muted** | Pointer to **bool** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**AssigneeId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Steps** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**Origin** | Pointer to **string** | \&quot;manual\&quot; | \&quot;imported\&quot; | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalPath** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**LastSyncedAt** | Pointer to **time.Time** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**BranchStatus** | Pointer to **string** |  | [optional] 
**ChildrenCount** | Pointer to **int32** | direct children, any kind | [optional] 
**DirectCaseCount** | Pointer to **int32** |  | [optional] 
**DirectSuiteCount** | Pointer to **int32** |  | [optional] 
**DescendantCaseCount** | Pointer to **int32** |  | [optional] 
**LinkedRequirementCount** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**TestopsId** | Pointer to **int32** |  | [optional] 
**AuthorType** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **string** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**ParameterGroups** | Pointer to [**[]TestParameterGroup**](TestParameterGroup.md) |  | [optional] 
**LatestExecution** | Pointer to [**TestExecution**](TestExecution.md) |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Relations** | Pointer to [**[]TestRelation**](TestRelation.md) |  | [optional] 

## Methods

### NewTest

`func NewTest() *Test`

NewTest instantiates a new Test object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWithDefaults

`func NewTestWithDefaults() *Test`

NewTestWithDefaults instantiates a new Test object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Test) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Test) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Test) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Test) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Test) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Test) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Test) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Test) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetBranchId

`func (o *Test) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *Test) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *Test) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *Test) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetParentId

`func (o *Test) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Test) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Test) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Test) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetKind

`func (o *Test) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Test) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Test) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Test) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetTitle

`func (o *Test) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Test) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Test) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Test) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Test) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Test) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Test) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Test) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPosition

`func (o *Test) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Test) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Test) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Test) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSeqNum

`func (o *Test) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *Test) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *Test) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *Test) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetStatus

`func (o *Test) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Test) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Test) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Test) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *Test) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Test) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Test) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Test) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *Test) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Test) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Test) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Test) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayer

`func (o *Test) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *Test) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *Test) SetLayer(v string)`

SetLayer sets Layer field to given value.

### HasLayer

`func (o *Test) HasLayer() bool`

HasLayer returns a boolean if a field has been set.

### GetMuted

`func (o *Test) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Test) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Test) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *Test) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetComponentId

`func (o *Test) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *Test) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *Test) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *Test) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetAssigneeId

`func (o *Test) GetAssigneeId() string`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *Test) GetAssigneeIdOk() (*string, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *Test) SetAssigneeId(v string)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *Test) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetTags

`func (o *Test) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Test) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Test) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Test) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Test) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Test) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Test) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Test) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSteps

`func (o *Test) GetSteps() []TestStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Test) GetStepsOk() (*[]TestStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Test) SetSteps(v []TestStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Test) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetOrigin

`func (o *Test) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Test) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Test) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Test) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFramework

`func (o *Test) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *Test) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *Test) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *Test) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetExternalId

`func (o *Test) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Test) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Test) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Test) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalPath

`func (o *Test) GetExternalPath() string`

GetExternalPath returns the ExternalPath field if non-nil, zero value otherwise.

### GetExternalPathOk

`func (o *Test) GetExternalPathOk() (*string, bool)`

GetExternalPathOk returns a tuple with the ExternalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPath

`func (o *Test) SetExternalPath(v string)`

SetExternalPath sets ExternalPath field to given value.

### HasExternalPath

`func (o *Test) HasExternalPath() bool`

HasExternalPath returns a boolean if a field has been set.

### GetFilePath

`func (o *Test) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *Test) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *Test) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *Test) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *Test) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *Test) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *Test) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *Test) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetSourceId

`func (o *Test) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Test) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Test) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Test) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetBranchStatus

`func (o *Test) GetBranchStatus() string`

GetBranchStatus returns the BranchStatus field if non-nil, zero value otherwise.

### GetBranchStatusOk

`func (o *Test) GetBranchStatusOk() (*string, bool)`

GetBranchStatusOk returns a tuple with the BranchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchStatus

`func (o *Test) SetBranchStatus(v string)`

SetBranchStatus sets BranchStatus field to given value.

### HasBranchStatus

`func (o *Test) HasBranchStatus() bool`

HasBranchStatus returns a boolean if a field has been set.

### GetChildrenCount

`func (o *Test) GetChildrenCount() int32`

GetChildrenCount returns the ChildrenCount field if non-nil, zero value otherwise.

### GetChildrenCountOk

`func (o *Test) GetChildrenCountOk() (*int32, bool)`

GetChildrenCountOk returns a tuple with the ChildrenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenCount

`func (o *Test) SetChildrenCount(v int32)`

SetChildrenCount sets ChildrenCount field to given value.

### HasChildrenCount

`func (o *Test) HasChildrenCount() bool`

HasChildrenCount returns a boolean if a field has been set.

### GetDirectCaseCount

`func (o *Test) GetDirectCaseCount() int32`

GetDirectCaseCount returns the DirectCaseCount field if non-nil, zero value otherwise.

### GetDirectCaseCountOk

`func (o *Test) GetDirectCaseCountOk() (*int32, bool)`

GetDirectCaseCountOk returns a tuple with the DirectCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCaseCount

`func (o *Test) SetDirectCaseCount(v int32)`

SetDirectCaseCount sets DirectCaseCount field to given value.

### HasDirectCaseCount

`func (o *Test) HasDirectCaseCount() bool`

HasDirectCaseCount returns a boolean if a field has been set.

### GetDirectSuiteCount

`func (o *Test) GetDirectSuiteCount() int32`

GetDirectSuiteCount returns the DirectSuiteCount field if non-nil, zero value otherwise.

### GetDirectSuiteCountOk

`func (o *Test) GetDirectSuiteCountOk() (*int32, bool)`

GetDirectSuiteCountOk returns a tuple with the DirectSuiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSuiteCount

`func (o *Test) SetDirectSuiteCount(v int32)`

SetDirectSuiteCount sets DirectSuiteCount field to given value.

### HasDirectSuiteCount

`func (o *Test) HasDirectSuiteCount() bool`

HasDirectSuiteCount returns a boolean if a field has been set.

### GetDescendantCaseCount

`func (o *Test) GetDescendantCaseCount() int32`

GetDescendantCaseCount returns the DescendantCaseCount field if non-nil, zero value otherwise.

### GetDescendantCaseCountOk

`func (o *Test) GetDescendantCaseCountOk() (*int32, bool)`

GetDescendantCaseCountOk returns a tuple with the DescendantCaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendantCaseCount

`func (o *Test) SetDescendantCaseCount(v int32)`

SetDescendantCaseCount sets DescendantCaseCount field to given value.

### HasDescendantCaseCount

`func (o *Test) HasDescendantCaseCount() bool`

HasDescendantCaseCount returns a boolean if a field has been set.

### GetLinkedRequirementCount

`func (o *Test) GetLinkedRequirementCount() int32`

GetLinkedRequirementCount returns the LinkedRequirementCount field if non-nil, zero value otherwise.

### GetLinkedRequirementCountOk

`func (o *Test) GetLinkedRequirementCountOk() (*int32, bool)`

GetLinkedRequirementCountOk returns a tuple with the LinkedRequirementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRequirementCount

`func (o *Test) SetLinkedRequirementCount(v int32)`

SetLinkedRequirementCount sets LinkedRequirementCount field to given value.

### HasLinkedRequirementCount

`func (o *Test) HasLinkedRequirementCount() bool`

HasLinkedRequirementCount returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Test) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Test) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Test) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Test) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Test) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Test) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Test) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Test) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Test) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Test) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Test) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Test) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsAutomated

`func (o *Test) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *Test) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *Test) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *Test) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetSignature

`func (o *Test) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Test) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Test) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Test) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTestopsId

`func (o *Test) GetTestopsId() int32`

GetTestopsId returns the TestopsId field if non-nil, zero value otherwise.

### GetTestopsIdOk

`func (o *Test) GetTestopsIdOk() (*int32, bool)`

GetTestopsIdOk returns a tuple with the TestopsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestopsId

`func (o *Test) SetTestopsId(v int32)`

SetTestopsId sets TestopsId field to given value.

### HasTestopsId

`func (o *Test) HasTestopsId() bool`

HasTestopsId returns a boolean if a field has been set.

### GetAuthorType

`func (o *Test) GetAuthorType() string`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *Test) GetAuthorTypeOk() (*string, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *Test) SetAuthorType(v string)`

SetAuthorType sets AuthorType field to given value.

### HasAuthorType

`func (o *Test) HasAuthorType() bool`

HasAuthorType returns a boolean if a field has been set.

### GetAuthorId

`func (o *Test) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Test) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Test) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *Test) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *Test) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Test) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Test) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *Test) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetParameterGroups

`func (o *Test) GetParameterGroups() []TestParameterGroup`

GetParameterGroups returns the ParameterGroups field if non-nil, zero value otherwise.

### GetParameterGroupsOk

`func (o *Test) GetParameterGroupsOk() (*[]TestParameterGroup, bool)`

GetParameterGroupsOk returns a tuple with the ParameterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroups

`func (o *Test) SetParameterGroups(v []TestParameterGroup)`

SetParameterGroups sets ParameterGroups field to given value.

### HasParameterGroups

`func (o *Test) HasParameterGroups() bool`

HasParameterGroups returns a boolean if a field has been set.

### GetLatestExecution

`func (o *Test) GetLatestExecution() TestExecution`

GetLatestExecution returns the LatestExecution field if non-nil, zero value otherwise.

### GetLatestExecutionOk

`func (o *Test) GetLatestExecutionOk() (*TestExecution, bool)`

GetLatestExecutionOk returns a tuple with the LatestExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestExecution

`func (o *Test) SetLatestExecution(v TestExecution)`

SetLatestExecution sets LatestExecution field to given value.

### HasLatestExecution

`func (o *Test) HasLatestExecution() bool`

HasLatestExecution returns a boolean if a field has been set.

### GetAttachments

`func (o *Test) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Test) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Test) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Test) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetRelations

`func (o *Test) GetRelations() []TestRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Test) GetRelationsOk() (*[]TestRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Test) SetRelations(v []TestRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Test) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


