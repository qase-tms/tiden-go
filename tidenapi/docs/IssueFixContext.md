# IssueFixContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issue** | Pointer to [**Issue**](Issue.md) |  | [optional] 
**LatestEvent** | Pointer to [**IssueEvent**](IssueEvent.md) |  | [optional] 
**SuspectPaths** | Pointer to **[]string** | suspect_paths are the file paths of the in-app stack frames, in stack order and deduplicated. These are the files to open. | [optional] 
**Environments** | Pointer to [**[]IssueEnvironmentCount**](IssueEnvironmentCount.md) | environments is the per-environment occurrence split — how you tell a production outage from dev noise. | [optional] 
**FirstReleaseName** | Pointer to **string** |  | [optional] 
**LastReleaseName** | Pointer to **string** |  | [optional] 
**Regressed** | Pointer to **bool** | regressed is true when this issue was resolved and came back, reopening automatically. Treat these as highest priority. | [optional] 
**Component** | Pointer to [**Component**](Component.md) |  | [optional] 
**SuspectRequirements** | Pointer to [**[]SuspectRequirement**](SuspectRequirement.md) |  | [optional] 
**TruncationSignals** | Pointer to **[]string** | truncation_signals names anything omitted to keep the response bounded. | [optional] 

## Methods

### NewIssueFixContext

`func NewIssueFixContext() *IssueFixContext`

NewIssueFixContext instantiates a new IssueFixContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFixContextWithDefaults

`func NewIssueFixContextWithDefaults() *IssueFixContext`

NewIssueFixContextWithDefaults instantiates a new IssueFixContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssue

`func (o *IssueFixContext) GetIssue() Issue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *IssueFixContext) GetIssueOk() (*Issue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *IssueFixContext) SetIssue(v Issue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *IssueFixContext) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetLatestEvent

`func (o *IssueFixContext) GetLatestEvent() IssueEvent`

GetLatestEvent returns the LatestEvent field if non-nil, zero value otherwise.

### GetLatestEventOk

`func (o *IssueFixContext) GetLatestEventOk() (*IssueEvent, bool)`

GetLatestEventOk returns a tuple with the LatestEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestEvent

`func (o *IssueFixContext) SetLatestEvent(v IssueEvent)`

SetLatestEvent sets LatestEvent field to given value.

### HasLatestEvent

`func (o *IssueFixContext) HasLatestEvent() bool`

HasLatestEvent returns a boolean if a field has been set.

### GetSuspectPaths

`func (o *IssueFixContext) GetSuspectPaths() []string`

GetSuspectPaths returns the SuspectPaths field if non-nil, zero value otherwise.

### GetSuspectPathsOk

`func (o *IssueFixContext) GetSuspectPathsOk() (*[]string, bool)`

GetSuspectPathsOk returns a tuple with the SuspectPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspectPaths

`func (o *IssueFixContext) SetSuspectPaths(v []string)`

SetSuspectPaths sets SuspectPaths field to given value.

### HasSuspectPaths

`func (o *IssueFixContext) HasSuspectPaths() bool`

HasSuspectPaths returns a boolean if a field has been set.

### GetEnvironments

`func (o *IssueFixContext) GetEnvironments() []IssueEnvironmentCount`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *IssueFixContext) GetEnvironmentsOk() (*[]IssueEnvironmentCount, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *IssueFixContext) SetEnvironments(v []IssueEnvironmentCount)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *IssueFixContext) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetFirstReleaseName

`func (o *IssueFixContext) GetFirstReleaseName() string`

GetFirstReleaseName returns the FirstReleaseName field if non-nil, zero value otherwise.

### GetFirstReleaseNameOk

`func (o *IssueFixContext) GetFirstReleaseNameOk() (*string, bool)`

GetFirstReleaseNameOk returns a tuple with the FirstReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseName

`func (o *IssueFixContext) SetFirstReleaseName(v string)`

SetFirstReleaseName sets FirstReleaseName field to given value.

### HasFirstReleaseName

`func (o *IssueFixContext) HasFirstReleaseName() bool`

HasFirstReleaseName returns a boolean if a field has been set.

### GetLastReleaseName

`func (o *IssueFixContext) GetLastReleaseName() string`

GetLastReleaseName returns the LastReleaseName field if non-nil, zero value otherwise.

### GetLastReleaseNameOk

`func (o *IssueFixContext) GetLastReleaseNameOk() (*string, bool)`

GetLastReleaseNameOk returns a tuple with the LastReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseName

`func (o *IssueFixContext) SetLastReleaseName(v string)`

SetLastReleaseName sets LastReleaseName field to given value.

### HasLastReleaseName

`func (o *IssueFixContext) HasLastReleaseName() bool`

HasLastReleaseName returns a boolean if a field has been set.

### GetRegressed

`func (o *IssueFixContext) GetRegressed() bool`

GetRegressed returns the Regressed field if non-nil, zero value otherwise.

### GetRegressedOk

`func (o *IssueFixContext) GetRegressedOk() (*bool, bool)`

GetRegressedOk returns a tuple with the Regressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegressed

`func (o *IssueFixContext) SetRegressed(v bool)`

SetRegressed sets Regressed field to given value.

### HasRegressed

`func (o *IssueFixContext) HasRegressed() bool`

HasRegressed returns a boolean if a field has been set.

### GetComponent

`func (o *IssueFixContext) GetComponent() Component`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *IssueFixContext) GetComponentOk() (*Component, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *IssueFixContext) SetComponent(v Component)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *IssueFixContext) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetSuspectRequirements

`func (o *IssueFixContext) GetSuspectRequirements() []SuspectRequirement`

GetSuspectRequirements returns the SuspectRequirements field if non-nil, zero value otherwise.

### GetSuspectRequirementsOk

`func (o *IssueFixContext) GetSuspectRequirementsOk() (*[]SuspectRequirement, bool)`

GetSuspectRequirementsOk returns a tuple with the SuspectRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspectRequirements

`func (o *IssueFixContext) SetSuspectRequirements(v []SuspectRequirement)`

SetSuspectRequirements sets SuspectRequirements field to given value.

### HasSuspectRequirements

`func (o *IssueFixContext) HasSuspectRequirements() bool`

HasSuspectRequirements returns a boolean if a field has been set.

### GetTruncationSignals

`func (o *IssueFixContext) GetTruncationSignals() []string`

GetTruncationSignals returns the TruncationSignals field if non-nil, zero value otherwise.

### GetTruncationSignalsOk

`func (o *IssueFixContext) GetTruncationSignalsOk() (*[]string, bool)`

GetTruncationSignalsOk returns a tuple with the TruncationSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationSignals

`func (o *IssueFixContext) SetTruncationSignals(v []string)`

SetTruncationSignals sets TruncationSignals field to given value.

### HasTruncationSignals

`func (o *IssueFixContext) HasTruncationSignals() bool`

HasTruncationSignals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


