# GetSessionProgressBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**RequirementIds** | Pointer to **[]string** |  | [optional] 
**IntentBranch** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSessionProgressBody

`func NewGetSessionProgressBody() *GetSessionProgressBody`

NewGetSessionProgressBody instantiates a new GetSessionProgressBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionProgressBodyWithDefaults

`func NewGetSessionProgressBodyWithDefaults() *GetSessionProgressBody`

NewGetSessionProgressBodyWithDefaults instantiates a new GetSessionProgressBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *GetSessionProgressBody) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GetSessionProgressBody) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GetSessionProgressBody) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GetSessionProgressBody) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetRequirementIds

`func (o *GetSessionProgressBody) GetRequirementIds() []string`

GetRequirementIds returns the RequirementIds field if non-nil, zero value otherwise.

### GetRequirementIdsOk

`func (o *GetSessionProgressBody) GetRequirementIdsOk() (*[]string, bool)`

GetRequirementIdsOk returns a tuple with the RequirementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementIds

`func (o *GetSessionProgressBody) SetRequirementIds(v []string)`

SetRequirementIds sets RequirementIds field to given value.

### HasRequirementIds

`func (o *GetSessionProgressBody) HasRequirementIds() bool`

HasRequirementIds returns a boolean if a field has been set.

### GetIntentBranch

`func (o *GetSessionProgressBody) GetIntentBranch() string`

GetIntentBranch returns the IntentBranch field if non-nil, zero value otherwise.

### GetIntentBranchOk

`func (o *GetSessionProgressBody) GetIntentBranchOk() (*string, bool)`

GetIntentBranchOk returns a tuple with the IntentBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentBranch

`func (o *GetSessionProgressBody) SetIntentBranch(v string)`

SetIntentBranch sets IntentBranch field to given value.

### HasIntentBranch

`func (o *GetSessionProgressBody) HasIntentBranch() bool`

HasIntentBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


