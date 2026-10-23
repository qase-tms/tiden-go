# GetSessionProgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | Pointer to [**[]SessionProgressRequirement**](SessionProgressRequirement.md) |  | [optional] 
**Summary** | Pointer to [**SessionProgressSummary**](SessionProgressSummary.md) |  | [optional] 
**Ready** | Pointer to **bool** |  | [optional] 
**NextActions** | Pointer to **[]string** |  | [optional] 
**IntentBranchStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSessionProgressResponse

`func NewGetSessionProgressResponse() *GetSessionProgressResponse`

NewGetSessionProgressResponse instantiates a new GetSessionProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionProgressResponseWithDefaults

`func NewGetSessionProgressResponseWithDefaults() *GetSessionProgressResponse`

NewGetSessionProgressResponseWithDefaults instantiates a new GetSessionProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *GetSessionProgressResponse) GetRequirements() []SessionProgressRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *GetSessionProgressResponse) GetRequirementsOk() (*[]SessionProgressRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *GetSessionProgressResponse) SetRequirements(v []SessionProgressRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *GetSessionProgressResponse) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetSummary

`func (o *GetSessionProgressResponse) GetSummary() SessionProgressSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetSessionProgressResponse) GetSummaryOk() (*SessionProgressSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetSessionProgressResponse) SetSummary(v SessionProgressSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetSessionProgressResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetReady

`func (o *GetSessionProgressResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *GetSessionProgressResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *GetSessionProgressResponse) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *GetSessionProgressResponse) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetNextActions

`func (o *GetSessionProgressResponse) GetNextActions() []string`

GetNextActions returns the NextActions field if non-nil, zero value otherwise.

### GetNextActionsOk

`func (o *GetSessionProgressResponse) GetNextActionsOk() (*[]string, bool)`

GetNextActionsOk returns a tuple with the NextActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextActions

`func (o *GetSessionProgressResponse) SetNextActions(v []string)`

SetNextActions sets NextActions field to given value.

### HasNextActions

`func (o *GetSessionProgressResponse) HasNextActions() bool`

HasNextActions returns a boolean if a field has been set.

### GetIntentBranchStatus

`func (o *GetSessionProgressResponse) GetIntentBranchStatus() string`

GetIntentBranchStatus returns the IntentBranchStatus field if non-nil, zero value otherwise.

### GetIntentBranchStatusOk

`func (o *GetSessionProgressResponse) GetIntentBranchStatusOk() (*string, bool)`

GetIntentBranchStatusOk returns a tuple with the IntentBranchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentBranchStatus

`func (o *GetSessionProgressResponse) SetIntentBranchStatus(v string)`

SetIntentBranchStatus sets IntentBranchStatus field to given value.

### HasIntentBranchStatus

`func (o *GetSessionProgressResponse) HasIntentBranchStatus() bool`

HasIntentBranchStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


