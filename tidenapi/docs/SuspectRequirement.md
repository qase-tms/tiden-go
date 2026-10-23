# SuspectRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**MatchedPaths** | Pointer to **[]string** | matched_paths are the suspect paths anchored to this requirement, so a wrong match is visible rather than silent. | [optional] 
**Coverage** | Pointer to **string** | coverage is verified | not_run | no_test — the same vocabulary the traceability matrix uses. Anything but \&quot;verified\&quot; is a reason this error escaped. | [optional] 
**Tests** | Pointer to [**[]CoveringTest**](CoveringTest.md) |  | [optional] 

## Methods

### NewSuspectRequirement

`func NewSuspectRequirement() *SuspectRequirement`

NewSuspectRequirement instantiates a new SuspectRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspectRequirementWithDefaults

`func NewSuspectRequirementWithDefaults() *SuspectRequirement`

NewSuspectRequirementWithDefaults instantiates a new SuspectRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuspectRequirement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuspectRequirement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuspectRequirement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SuspectRequirement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *SuspectRequirement) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *SuspectRequirement) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *SuspectRequirement) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *SuspectRequirement) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *SuspectRequirement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SuspectRequirement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SuspectRequirement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SuspectRequirement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMatchedPaths

`func (o *SuspectRequirement) GetMatchedPaths() []string`

GetMatchedPaths returns the MatchedPaths field if non-nil, zero value otherwise.

### GetMatchedPathsOk

`func (o *SuspectRequirement) GetMatchedPathsOk() (*[]string, bool)`

GetMatchedPathsOk returns a tuple with the MatchedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPaths

`func (o *SuspectRequirement) SetMatchedPaths(v []string)`

SetMatchedPaths sets MatchedPaths field to given value.

### HasMatchedPaths

`func (o *SuspectRequirement) HasMatchedPaths() bool`

HasMatchedPaths returns a boolean if a field has been set.

### GetCoverage

`func (o *SuspectRequirement) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *SuspectRequirement) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *SuspectRequirement) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *SuspectRequirement) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetTests

`func (o *SuspectRequirement) GetTests() []CoveringTest`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *SuspectRequirement) GetTestsOk() (*[]CoveringTest, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *SuspectRequirement) SetTests(v []CoveringTest)`

SetTests sets Tests field to given value.

### HasTests

`func (o *SuspectRequirement) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


