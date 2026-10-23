# GetIssueEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**IssueEvent**](IssueEvent.md) |  | [optional] 

## Methods

### NewGetIssueEventResponse

`func NewGetIssueEventResponse() *GetIssueEventResponse`

NewGetIssueEventResponse instantiates a new GetIssueEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueEventResponseWithDefaults

`func NewGetIssueEventResponseWithDefaults() *GetIssueEventResponse`

NewGetIssueEventResponseWithDefaults instantiates a new GetIssueEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GetIssueEventResponse) GetEvent() IssueEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetIssueEventResponse) GetEventOk() (*IssueEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetIssueEventResponse) SetEvent(v IssueEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GetIssueEventResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


