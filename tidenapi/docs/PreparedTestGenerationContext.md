# PreparedTestGenerationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**TokenBudget** | Pointer to **int32** |  | [optional] 
**Contexts** | Pointer to [**[]RequirementTestContext**](RequirementTestContext.md) |  | [optional] 
**CodebaseContext** | Pointer to [**CodebaseContext**](CodebaseContext.md) |  | [optional] 
**Citations** | Pointer to [**[]ContextCitation**](ContextCitation.md) |  | [optional] 
**TruncationSignals** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPreparedTestGenerationContext

`func NewPreparedTestGenerationContext() *PreparedTestGenerationContext`

NewPreparedTestGenerationContext instantiates a new PreparedTestGenerationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparedTestGenerationContextWithDefaults

`func NewPreparedTestGenerationContextWithDefaults() *PreparedTestGenerationContext`

NewPreparedTestGenerationContextWithDefaults instantiates a new PreparedTestGenerationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PreparedTestGenerationContext) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PreparedTestGenerationContext) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PreparedTestGenerationContext) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PreparedTestGenerationContext) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetBranch

`func (o *PreparedTestGenerationContext) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PreparedTestGenerationContext) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PreparedTestGenerationContext) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PreparedTestGenerationContext) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFramework

`func (o *PreparedTestGenerationContext) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *PreparedTestGenerationContext) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *PreparedTestGenerationContext) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *PreparedTestGenerationContext) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetTokenBudget

`func (o *PreparedTestGenerationContext) GetTokenBudget() int32`

GetTokenBudget returns the TokenBudget field if non-nil, zero value otherwise.

### GetTokenBudgetOk

`func (o *PreparedTestGenerationContext) GetTokenBudgetOk() (*int32, bool)`

GetTokenBudgetOk returns a tuple with the TokenBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBudget

`func (o *PreparedTestGenerationContext) SetTokenBudget(v int32)`

SetTokenBudget sets TokenBudget field to given value.

### HasTokenBudget

`func (o *PreparedTestGenerationContext) HasTokenBudget() bool`

HasTokenBudget returns a boolean if a field has been set.

### GetContexts

`func (o *PreparedTestGenerationContext) GetContexts() []RequirementTestContext`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *PreparedTestGenerationContext) GetContextsOk() (*[]RequirementTestContext, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *PreparedTestGenerationContext) SetContexts(v []RequirementTestContext)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *PreparedTestGenerationContext) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetCodebaseContext

`func (o *PreparedTestGenerationContext) GetCodebaseContext() CodebaseContext`

GetCodebaseContext returns the CodebaseContext field if non-nil, zero value otherwise.

### GetCodebaseContextOk

`func (o *PreparedTestGenerationContext) GetCodebaseContextOk() (*CodebaseContext, bool)`

GetCodebaseContextOk returns a tuple with the CodebaseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodebaseContext

`func (o *PreparedTestGenerationContext) SetCodebaseContext(v CodebaseContext)`

SetCodebaseContext sets CodebaseContext field to given value.

### HasCodebaseContext

`func (o *PreparedTestGenerationContext) HasCodebaseContext() bool`

HasCodebaseContext returns a boolean if a field has been set.

### GetCitations

`func (o *PreparedTestGenerationContext) GetCitations() []ContextCitation`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *PreparedTestGenerationContext) GetCitationsOk() (*[]ContextCitation, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *PreparedTestGenerationContext) SetCitations(v []ContextCitation)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *PreparedTestGenerationContext) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetTruncationSignals

`func (o *PreparedTestGenerationContext) GetTruncationSignals() []string`

GetTruncationSignals returns the TruncationSignals field if non-nil, zero value otherwise.

### GetTruncationSignalsOk

`func (o *PreparedTestGenerationContext) GetTruncationSignalsOk() (*[]string, bool)`

GetTruncationSignalsOk returns a tuple with the TruncationSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationSignals

`func (o *PreparedTestGenerationContext) SetTruncationSignals(v []string)`

SetTruncationSignals sets TruncationSignals field to given value.

### HasTruncationSignals

`func (o *PreparedTestGenerationContext) HasTruncationSignals() bool`

HasTruncationSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


