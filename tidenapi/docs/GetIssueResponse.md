# GetIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issue** | Pointer to [**Issue**](Issue.md) |  | [optional] 
**LatestEvent** | Pointer to [**IssueEvent**](IssueEvent.md) |  | [optional] 

## Methods

### NewGetIssueResponse

`func NewGetIssueResponse() *GetIssueResponse`

NewGetIssueResponse instantiates a new GetIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueResponseWithDefaults

`func NewGetIssueResponseWithDefaults() *GetIssueResponse`

NewGetIssueResponseWithDefaults instantiates a new GetIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssue

`func (o *GetIssueResponse) GetIssue() Issue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *GetIssueResponse) GetIssueOk() (*Issue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *GetIssueResponse) SetIssue(v Issue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *GetIssueResponse) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetLatestEvent

`func (o *GetIssueResponse) GetLatestEvent() IssueEvent`

GetLatestEvent returns the LatestEvent field if non-nil, zero value otherwise.

### GetLatestEventOk

`func (o *GetIssueResponse) GetLatestEventOk() (*IssueEvent, bool)`

GetLatestEventOk returns a tuple with the LatestEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEvent

`func (o *GetIssueResponse) SetLatestEvent(v IssueEvent)`

SetLatestEvent sets LatestEvent field to given value.

### HasLatestEvent

`func (o *GetIssueResponse) HasLatestEvent() bool`

HasLatestEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


