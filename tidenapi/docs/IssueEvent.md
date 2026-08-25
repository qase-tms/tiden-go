# IssueEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ExceptionType** | Pointer to **string** |  | [optional] 
**ExceptionValue** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ReleaseId** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**ReceivedAt** | Pointer to **time.Time** |  | [optional] 
**Frames** | Pointer to [**[]Frame**](Frame.md) | Server-resolved stacktrace overlay (symbolicated frames), populated only on GetIssue.latest_event. payload stays for breadcrumbs/request/etc. | [optional] 
**ReleaseName** | Pointer to **string** | Human-readable release/environment as sent by the SDK. The *_id fields above are the resolved entity FKs; these names are what the UI renders directly. | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 

## Methods

### NewIssueEvent

`func NewIssueEvent() *IssueEvent`

NewIssueEvent instantiates a new IssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueEventWithDefaults

`func NewIssueEventWithDefaults() *IssueEvent`

NewIssueEventWithDefaults instantiates a new IssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventId

`func (o *IssueEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *IssueEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *IssueEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *IssueEvent) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetLevel

`func (o *IssueEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IssueEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IssueEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IssueEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *IssueEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IssueEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IssueEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IssueEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExceptionType

`func (o *IssueEvent) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *IssueEvent) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *IssueEvent) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *IssueEvent) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetExceptionValue

`func (o *IssueEvent) GetExceptionValue() string`

GetExceptionValue returns the ExceptionValue field if non-nil, zero value otherwise.

### GetExceptionValueOk

`func (o *IssueEvent) GetExceptionValueOk() (*string, bool)`

GetExceptionValueOk returns a tuple with the ExceptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionValue

`func (o *IssueEvent) SetExceptionValue(v string)`

SetExceptionValue sets ExceptionValue field to given value.

### HasExceptionValue

`func (o *IssueEvent) HasExceptionValue() bool`

HasExceptionValue returns a boolean if a field has been set.

### GetPlatform

`func (o *IssueEvent) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *IssueEvent) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *IssueEvent) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *IssueEvent) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetReleaseId

`func (o *IssueEvent) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *IssueEvent) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *IssueEvent) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *IssueEvent) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *IssueEvent) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *IssueEvent) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *IssueEvent) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *IssueEvent) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetPayload

`func (o *IssueEvent) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IssueEvent) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IssueEvent) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IssueEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetReceivedAt

`func (o *IssueEvent) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *IssueEvent) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *IssueEvent) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *IssueEvent) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetFrames

`func (o *IssueEvent) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *IssueEvent) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *IssueEvent) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *IssueEvent) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetReleaseName

`func (o *IssueEvent) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *IssueEvent) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *IssueEvent) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *IssueEvent) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *IssueEvent) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *IssueEvent) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *IssueEvent) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *IssueEvent) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


