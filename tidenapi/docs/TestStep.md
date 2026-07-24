# TestStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Expected** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]TestStep**](TestStep.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InputData** | Pointer to **map[string]interface{}** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**Execution** | Pointer to [**TestExecution**](TestExecution.md) |  | [optional] 

## Methods

### NewTestStep

`func NewTestStep() *TestStep`

NewTestStep instantiates a new TestStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStepWithDefaults

`func NewTestStepWithDefaults() *TestStep`

NewTestStepWithDefaults instantiates a new TestStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *TestStep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestStep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestStep) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestStep) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExpected

`func (o *TestStep) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *TestStep) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *TestStep) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *TestStep) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetData

`func (o *TestStep) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TestStep) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TestStep) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *TestStep) HasData() bool`

HasData returns a boolean if a field has been set.

### GetChildren

`func (o *TestStep) GetChildren() []TestStep`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TestStep) GetChildrenOk() (*[]TestStep, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TestStep) SetChildren(v []TestStep)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TestStep) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetType

`func (o *TestStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestStep) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestStep) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeyword

`func (o *TestStep) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *TestStep) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *TestStep) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *TestStep) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetName

`func (o *TestStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInputData

`func (o *TestStep) GetInputData() map[string]interface{}`

GetInputData returns the InputData field if non-nil, zero value otherwise.

### GetInputDataOk

`func (o *TestStep) GetInputDataOk() (*map[string]interface{}, bool)`

GetInputDataOk returns a tuple with the InputData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputData

`func (o *TestStep) SetInputData(v map[string]interface{})`

SetInputData sets InputData field to given value.

### HasInputData

`func (o *TestStep) HasInputData() bool`

HasInputData returns a boolean if a field has been set.

### GetAttachments

`func (o *TestStep) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *TestStep) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *TestStep) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *TestStep) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetExecution

`func (o *TestStep) GetExecution() TestExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *TestStep) GetExecutionOk() (*TestExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *TestStep) SetExecution(v TestExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *TestStep) HasExecution() bool`

HasExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


