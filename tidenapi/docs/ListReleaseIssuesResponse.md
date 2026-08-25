# ListReleaseIssuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewIssues** | Pointer to [**[]Issue**](Issue.md) |  | [optional] 
**SeenCount** | Pointer to **string** |  | [optional] 

## Methods

### NewListReleaseIssuesResponse

`func NewListReleaseIssuesResponse() *ListReleaseIssuesResponse`

NewListReleaseIssuesResponse instantiates a new ListReleaseIssuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReleaseIssuesResponseWithDefaults

`func NewListReleaseIssuesResponseWithDefaults() *ListReleaseIssuesResponse`

NewListReleaseIssuesResponseWithDefaults instantiates a new ListReleaseIssuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewIssues

`func (o *ListReleaseIssuesResponse) GetNewIssues() []Issue`

GetNewIssues returns the NewIssues field if non-nil, zero value otherwise.

### GetNewIssuesOk

`func (o *ListReleaseIssuesResponse) GetNewIssuesOk() (*[]Issue, bool)`

GetNewIssuesOk returns a tuple with the NewIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIssues

`func (o *ListReleaseIssuesResponse) SetNewIssues(v []Issue)`

SetNewIssues sets NewIssues field to given value.

### HasNewIssues

`func (o *ListReleaseIssuesResponse) HasNewIssues() bool`

HasNewIssues returns a boolean if a field has been set.

### GetSeenCount

`func (o *ListReleaseIssuesResponse) GetSeenCount() string`

GetSeenCount returns the SeenCount field if non-nil, zero value otherwise.

### GetSeenCountOk

`func (o *ListReleaseIssuesResponse) GetSeenCountOk() (*string, bool)`

GetSeenCountOk returns a tuple with the SeenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenCount

`func (o *ListReleaseIssuesResponse) SetSeenCount(v string)`

SetSeenCount sets SeenCount field to given value.

### HasSeenCount

`func (o *ListReleaseIssuesResponse) HasSeenCount() bool`

HasSeenCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


