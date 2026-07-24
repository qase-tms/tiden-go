# FileAnchorCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** |  | [optional] 
**RequirementSeq** | Pointer to **int32** |  | [optional] 
**RequirementTitle** | Pointer to **string** |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**TestSeq** | Pointer to **int32** |  | [optional] 
**TestTitle** | Pointer to **string** |  | [optional] 
**TestFilePath** | Pointer to **string** |  | [optional] 
**AnchorPath** | Pointer to **string** |  | [optional] 
**Dir** | Pointer to **string** |  | [optional] 
**Exact** | Pointer to **bool** |  | [optional] 
**Confidence** | Pointer to **float64** |  | [optional] 

## Methods

### NewFileAnchorCandidate

`func NewFileAnchorCandidate() *FileAnchorCandidate`

NewFileAnchorCandidate instantiates a new FileAnchorCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAnchorCandidateWithDefaults

`func NewFileAnchorCandidateWithDefaults() *FileAnchorCandidate`

NewFileAnchorCandidateWithDefaults instantiates a new FileAnchorCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *FileAnchorCandidate) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *FileAnchorCandidate) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *FileAnchorCandidate) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *FileAnchorCandidate) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetRequirementSeq

`func (o *FileAnchorCandidate) GetRequirementSeq() int32`

GetRequirementSeq returns the RequirementSeq field if non-nil, zero value otherwise.

### GetRequirementSeqOk

`func (o *FileAnchorCandidate) GetRequirementSeqOk() (*int32, bool)`

GetRequirementSeqOk returns a tuple with the RequirementSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementSeq

`func (o *FileAnchorCandidate) SetRequirementSeq(v int32)`

SetRequirementSeq sets RequirementSeq field to given value.

### HasRequirementSeq

`func (o *FileAnchorCandidate) HasRequirementSeq() bool`

HasRequirementSeq returns a boolean if a field has been set.

### GetRequirementTitle

`func (o *FileAnchorCandidate) GetRequirementTitle() string`

GetRequirementTitle returns the RequirementTitle field if non-nil, zero value otherwise.

### GetRequirementTitleOk

`func (o *FileAnchorCandidate) GetRequirementTitleOk() (*string, bool)`

GetRequirementTitleOk returns a tuple with the RequirementTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementTitle

`func (o *FileAnchorCandidate) SetRequirementTitle(v string)`

SetRequirementTitle sets RequirementTitle field to given value.

### HasRequirementTitle

`func (o *FileAnchorCandidate) HasRequirementTitle() bool`

HasRequirementTitle returns a boolean if a field has been set.

### GetTestId

`func (o *FileAnchorCandidate) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *FileAnchorCandidate) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *FileAnchorCandidate) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *FileAnchorCandidate) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTestSeq

`func (o *FileAnchorCandidate) GetTestSeq() int32`

GetTestSeq returns the TestSeq field if non-nil, zero value otherwise.

### GetTestSeqOk

`func (o *FileAnchorCandidate) GetTestSeqOk() (*int32, bool)`

GetTestSeqOk returns a tuple with the TestSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSeq

`func (o *FileAnchorCandidate) SetTestSeq(v int32)`

SetTestSeq sets TestSeq field to given value.

### HasTestSeq

`func (o *FileAnchorCandidate) HasTestSeq() bool`

HasTestSeq returns a boolean if a field has been set.

### GetTestTitle

`func (o *FileAnchorCandidate) GetTestTitle() string`

GetTestTitle returns the TestTitle field if non-nil, zero value otherwise.

### GetTestTitleOk

`func (o *FileAnchorCandidate) GetTestTitleOk() (*string, bool)`

GetTestTitleOk returns a tuple with the TestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTitle

`func (o *FileAnchorCandidate) SetTestTitle(v string)`

SetTestTitle sets TestTitle field to given value.

### HasTestTitle

`func (o *FileAnchorCandidate) HasTestTitle() bool`

HasTestTitle returns a boolean if a field has been set.

### GetTestFilePath

`func (o *FileAnchorCandidate) GetTestFilePath() string`

GetTestFilePath returns the TestFilePath field if non-nil, zero value otherwise.

### GetTestFilePathOk

`func (o *FileAnchorCandidate) GetTestFilePathOk() (*string, bool)`

GetTestFilePathOk returns a tuple with the TestFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFilePath

`func (o *FileAnchorCandidate) SetTestFilePath(v string)`

SetTestFilePath sets TestFilePath field to given value.

### HasTestFilePath

`func (o *FileAnchorCandidate) HasTestFilePath() bool`

HasTestFilePath returns a boolean if a field has been set.

### GetAnchorPath

`func (o *FileAnchorCandidate) GetAnchorPath() string`

GetAnchorPath returns the AnchorPath field if non-nil, zero value otherwise.

### GetAnchorPathOk

`func (o *FileAnchorCandidate) GetAnchorPathOk() (*string, bool)`

GetAnchorPathOk returns a tuple with the AnchorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorPath

`func (o *FileAnchorCandidate) SetAnchorPath(v string)`

SetAnchorPath sets AnchorPath field to given value.

### HasAnchorPath

`func (o *FileAnchorCandidate) HasAnchorPath() bool`

HasAnchorPath returns a boolean if a field has been set.

### GetDir

`func (o *FileAnchorCandidate) GetDir() string`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *FileAnchorCandidate) GetDirOk() (*string, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *FileAnchorCandidate) SetDir(v string)`

SetDir sets Dir field to given value.

### HasDir

`func (o *FileAnchorCandidate) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetExact

`func (o *FileAnchorCandidate) GetExact() bool`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *FileAnchorCandidate) GetExactOk() (*bool, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *FileAnchorCandidate) SetExact(v bool)`

SetExact sets Exact field to given value.

### HasExact

`func (o *FileAnchorCandidate) HasExact() bool`

HasExact returns a boolean if a field has been set.

### GetConfidence

`func (o *FileAnchorCandidate) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *FileAnchorCandidate) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *FileAnchorCandidate) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *FileAnchorCandidate) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


