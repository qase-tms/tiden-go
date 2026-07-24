# ChangedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | repository is the canonical repo id OR a local checkout alias (resolved via component repository_aliases). | [optional] 
**Path** | Pointer to **string** | path is repo-relative. | [optional] 
**Status** | Pointer to **string** | status is \&quot;modified\&quot;|\&quot;added\&quot;|\&quot;deleted\&quot; (informational — deleted files are still attributed to their component; the caller decides not to anchor them). | [optional] 

## Methods

### NewChangedFile

`func NewChangedFile() *ChangedFile`

NewChangedFile instantiates a new ChangedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangedFileWithDefaults

`func NewChangedFileWithDefaults() *ChangedFile`

NewChangedFileWithDefaults instantiates a new ChangedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *ChangedFile) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ChangedFile) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ChangedFile) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ChangedFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPath

`func (o *ChangedFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChangedFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChangedFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ChangedFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatus

`func (o *ChangedFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangedFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangedFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChangedFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


