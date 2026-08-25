# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Culprit** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimesSeen** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **time.Time** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**FirstReleaseId** | Pointer to **string** |  | [optional] 
**LastReleaseId** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**ResolvedAt** | Pointer to **time.Time** | Most recent transition to resolved. KEPT through an automatic regression reopen (it is the \&quot;was resolved at\&quot; the UI renders); cleared on any MANUAL transition away from resolved. | [optional] 
**RegressedAt** | Pointer to **time.Time** | Set when the grouping worker auto-reopens a resolved issue (regression); cleared on any manual status change. status is &#39;unresolved&#39; whenever set. | [optional] 

## Methods

### NewIssue

`func NewIssue() *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Issue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *Issue) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Issue) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Issue) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Issue) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetTitle

`func (o *Issue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Issue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Issue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Issue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCulprit

`func (o *Issue) GetCulprit() string`

GetCulprit returns the Culprit field if non-nil, zero value otherwise.

### GetCulpritOk

`func (o *Issue) GetCulpritOk() (*string, bool)`

GetCulpritOk returns a tuple with the Culprit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulprit

`func (o *Issue) SetCulprit(v string)`

SetCulprit sets Culprit field to given value.

### HasCulprit

`func (o *Issue) HasCulprit() bool`

HasCulprit returns a boolean if a field has been set.

### GetLevel

`func (o *Issue) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Issue) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Issue) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Issue) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPlatform

`func (o *Issue) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Issue) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Issue) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Issue) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *Issue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Issue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Issue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Issue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimesSeen

`func (o *Issue) GetTimesSeen() string`

GetTimesSeen returns the TimesSeen field if non-nil, zero value otherwise.

### GetTimesSeenOk

`func (o *Issue) GetTimesSeenOk() (*string, bool)`

GetTimesSeenOk returns a tuple with the TimesSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesSeen

`func (o *Issue) SetTimesSeen(v string)`

SetTimesSeen sets TimesSeen field to given value.

### HasTimesSeen

`func (o *Issue) HasTimesSeen() bool`

HasTimesSeen returns a boolean if a field has been set.

### GetFirstSeen

`func (o *Issue) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Issue) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *Issue) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *Issue) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *Issue) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Issue) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Issue) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Issue) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetFirstReleaseId

`func (o *Issue) GetFirstReleaseId() string`

GetFirstReleaseId returns the FirstReleaseId field if non-nil, zero value otherwise.

### GetFirstReleaseIdOk

`func (o *Issue) GetFirstReleaseIdOk() (*string, bool)`

GetFirstReleaseIdOk returns a tuple with the FirstReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseId

`func (o *Issue) SetFirstReleaseId(v string)`

SetFirstReleaseId sets FirstReleaseId field to given value.

### HasFirstReleaseId

`func (o *Issue) HasFirstReleaseId() bool`

HasFirstReleaseId returns a boolean if a field has been set.

### GetLastReleaseId

`func (o *Issue) GetLastReleaseId() string`

GetLastReleaseId returns the LastReleaseId field if non-nil, zero value otherwise.

### GetLastReleaseIdOk

`func (o *Issue) GetLastReleaseIdOk() (*string, bool)`

GetLastReleaseIdOk returns a tuple with the LastReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseId

`func (o *Issue) SetLastReleaseId(v string)`

SetLastReleaseId sets LastReleaseId field to given value.

### HasLastReleaseId

`func (o *Issue) HasLastReleaseId() bool`

HasLastReleaseId returns a boolean if a field has been set.

### GetComponentId

`func (o *Issue) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *Issue) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *Issue) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *Issue) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetResolvedAt

`func (o *Issue) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Issue) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Issue) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *Issue) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetRegressedAt

`func (o *Issue) GetRegressedAt() time.Time`

GetRegressedAt returns the RegressedAt field if non-nil, zero value otherwise.

### GetRegressedAtOk

`func (o *Issue) GetRegressedAtOk() (*time.Time, bool)`

GetRegressedAtOk returns a tuple with the RegressedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegressedAt

`func (o *Issue) SetRegressedAt(v time.Time)`

SetRegressedAt sets RegressedAt field to given value.

### HasRegressedAt

`func (o *Issue) HasRegressedAt() bool`

HasRegressedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


