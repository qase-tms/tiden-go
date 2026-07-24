# RequirementTestContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirement** | Pointer to [**Requirement**](Requirement.md) |  | [optional] 
**Parent** | Pointer to [**Requirement**](Requirement.md) |  | [optional] 
**Children** | Pointer to [**[]Requirement**](Requirement.md) |  | [optional] 
**Siblings** | Pointer to [**[]Requirement**](Requirement.md) |  | [optional] 
**Component** | Pointer to [**Component**](Component.md) |  | [optional] 
**LinkedTests** | Pointer to [**[]ContextTest**](ContextTest.md) |  | [optional] 
**ProposedTests** | Pointer to [**[]ContextTest**](ContextTest.md) |  | [optional] 
**RelevantTests** | Pointer to [**[]ContextTest**](ContextTest.md) |  | [optional] 
**StaleSignals** | Pointer to [**[]StaleCoverageSignal**](StaleCoverageSignal.md) |  | [optional] 
**ExtractedFields** | Pointer to [**RequirementTestFields**](RequirementTestFields.md) |  | [optional] 
**MemoryEntries** | Pointer to [**[]AgentMemoryContext**](AgentMemoryContext.md) |  | [optional] 
**Citations** | Pointer to [**[]ContextCitation**](ContextCitation.md) |  | [optional] 
**Sources** | Pointer to [**[]RequirementSource**](RequirementSource.md) |  | [optional] 
**TruncationSignals** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRequirementTestContext

`func NewRequirementTestContext() *RequirementTestContext`

NewRequirementTestContext instantiates a new RequirementTestContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementTestContextWithDefaults

`func NewRequirementTestContextWithDefaults() *RequirementTestContext`

NewRequirementTestContextWithDefaults instantiates a new RequirementTestContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirement

`func (o *RequirementTestContext) GetRequirement() Requirement`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *RequirementTestContext) GetRequirementOk() (*Requirement, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *RequirementTestContext) SetRequirement(v Requirement)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *RequirementTestContext) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### GetParent

`func (o *RequirementTestContext) GetParent() Requirement`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RequirementTestContext) GetParentOk() (*Requirement, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RequirementTestContext) SetParent(v Requirement)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RequirementTestContext) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetChildren

`func (o *RequirementTestContext) GetChildren() []Requirement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *RequirementTestContext) GetChildrenOk() (*[]Requirement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *RequirementTestContext) SetChildren(v []Requirement)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *RequirementTestContext) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSiblings

`func (o *RequirementTestContext) GetSiblings() []Requirement`

GetSiblings returns the Siblings field if non-nil, zero value otherwise.

### GetSiblingsOk

`func (o *RequirementTestContext) GetSiblingsOk() (*[]Requirement, bool)`

GetSiblingsOk returns a tuple with the Siblings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiblings

`func (o *RequirementTestContext) SetSiblings(v []Requirement)`

SetSiblings sets Siblings field to given value.

### HasSiblings

`func (o *RequirementTestContext) HasSiblings() bool`

HasSiblings returns a boolean if a field has been set.

### GetComponent

`func (o *RequirementTestContext) GetComponent() Component`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RequirementTestContext) GetComponentOk() (*Component, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RequirementTestContext) SetComponent(v Component)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RequirementTestContext) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetLinkedTests

`func (o *RequirementTestContext) GetLinkedTests() []ContextTest`

GetLinkedTests returns the LinkedTests field if non-nil, zero value otherwise.

### GetLinkedTestsOk

`func (o *RequirementTestContext) GetLinkedTestsOk() (*[]ContextTest, bool)`

GetLinkedTestsOk returns a tuple with the LinkedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTests

`func (o *RequirementTestContext) SetLinkedTests(v []ContextTest)`

SetLinkedTests sets LinkedTests field to given value.

### HasLinkedTests

`func (o *RequirementTestContext) HasLinkedTests() bool`

HasLinkedTests returns a boolean if a field has been set.

### GetProposedTests

`func (o *RequirementTestContext) GetProposedTests() []ContextTest`

GetProposedTests returns the ProposedTests field if non-nil, zero value otherwise.

### GetProposedTestsOk

`func (o *RequirementTestContext) GetProposedTestsOk() (*[]ContextTest, bool)`

GetProposedTestsOk returns a tuple with the ProposedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTests

`func (o *RequirementTestContext) SetProposedTests(v []ContextTest)`

SetProposedTests sets ProposedTests field to given value.

### HasProposedTests

`func (o *RequirementTestContext) HasProposedTests() bool`

HasProposedTests returns a boolean if a field has been set.

### GetRelevantTests

`func (o *RequirementTestContext) GetRelevantTests() []ContextTest`

GetRelevantTests returns the RelevantTests field if non-nil, zero value otherwise.

### GetRelevantTestsOk

`func (o *RequirementTestContext) GetRelevantTestsOk() (*[]ContextTest, bool)`

GetRelevantTestsOk returns a tuple with the RelevantTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevantTests

`func (o *RequirementTestContext) SetRelevantTests(v []ContextTest)`

SetRelevantTests sets RelevantTests field to given value.

### HasRelevantTests

`func (o *RequirementTestContext) HasRelevantTests() bool`

HasRelevantTests returns a boolean if a field has been set.

### GetStaleSignals

`func (o *RequirementTestContext) GetStaleSignals() []StaleCoverageSignal`

GetStaleSignals returns the StaleSignals field if non-nil, zero value otherwise.

### GetStaleSignalsOk

`func (o *RequirementTestContext) GetStaleSignalsOk() (*[]StaleCoverageSignal, bool)`

GetStaleSignalsOk returns a tuple with the StaleSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleSignals

`func (o *RequirementTestContext) SetStaleSignals(v []StaleCoverageSignal)`

SetStaleSignals sets StaleSignals field to given value.

### HasStaleSignals

`func (o *RequirementTestContext) HasStaleSignals() bool`

HasStaleSignals returns a boolean if a field has been set.

### GetExtractedFields

`func (o *RequirementTestContext) GetExtractedFields() RequirementTestFields`

GetExtractedFields returns the ExtractedFields field if non-nil, zero value otherwise.

### GetExtractedFieldsOk

`func (o *RequirementTestContext) GetExtractedFieldsOk() (*RequirementTestFields, bool)`

GetExtractedFieldsOk returns a tuple with the ExtractedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedFields

`func (o *RequirementTestContext) SetExtractedFields(v RequirementTestFields)`

SetExtractedFields sets ExtractedFields field to given value.

### HasExtractedFields

`func (o *RequirementTestContext) HasExtractedFields() bool`

HasExtractedFields returns a boolean if a field has been set.

### GetMemoryEntries

`func (o *RequirementTestContext) GetMemoryEntries() []AgentMemoryContext`

GetMemoryEntries returns the MemoryEntries field if non-nil, zero value otherwise.

### GetMemoryEntriesOk

`func (o *RequirementTestContext) GetMemoryEntriesOk() (*[]AgentMemoryContext, bool)`

GetMemoryEntriesOk returns a tuple with the MemoryEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEntries

`func (o *RequirementTestContext) SetMemoryEntries(v []AgentMemoryContext)`

SetMemoryEntries sets MemoryEntries field to given value.

### HasMemoryEntries

`func (o *RequirementTestContext) HasMemoryEntries() bool`

HasMemoryEntries returns a boolean if a field has been set.

### GetCitations

`func (o *RequirementTestContext) GetCitations() []ContextCitation`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *RequirementTestContext) GetCitationsOk() (*[]ContextCitation, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *RequirementTestContext) SetCitations(v []ContextCitation)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *RequirementTestContext) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetSources

`func (o *RequirementTestContext) GetSources() []RequirementSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RequirementTestContext) GetSourcesOk() (*[]RequirementSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RequirementTestContext) SetSources(v []RequirementSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RequirementTestContext) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTruncationSignals

`func (o *RequirementTestContext) GetTruncationSignals() []string`

GetTruncationSignals returns the TruncationSignals field if non-nil, zero value otherwise.

### GetTruncationSignalsOk

`func (o *RequirementTestContext) GetTruncationSignalsOk() (*[]string, bool)`

GetTruncationSignalsOk returns a tuple with the TruncationSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationSignals

`func (o *RequirementTestContext) SetTruncationSignals(v []string)`

SetTruncationSignals sets TruncationSignals field to given value.

### HasTruncationSignals

`func (o *RequirementTestContext) HasTruncationSignals() bool`

HasTruncationSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


