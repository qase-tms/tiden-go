# CodebaseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceFiles** | Pointer to [**[]CodebaseFile**](CodebaseFile.md) |  | [optional] 
**TestFiles** | Pointer to [**[]CodebaseFile**](CodebaseFile.md) |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**TestCommand** | Pointer to **string** |  | [optional] 
**ImportGraphHints** | Pointer to **[]string** |  | [optional] 
**StyleExamples** | Pointer to [**[]CodebaseFile**](CodebaseFile.md) |  | [optional] 
**FixturesAndMocks** | Pointer to [**[]CodebaseFile**](CodebaseFile.md) |  | [optional] 
**CandidateFilePaths** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCodebaseContext

`func NewCodebaseContext() *CodebaseContext`

NewCodebaseContext instantiates a new CodebaseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodebaseContextWithDefaults

`func NewCodebaseContextWithDefaults() *CodebaseContext`

NewCodebaseContextWithDefaults instantiates a new CodebaseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceFiles

`func (o *CodebaseContext) GetSourceFiles() []CodebaseFile`

GetSourceFiles returns the SourceFiles field if non-nil, zero value otherwise.

### GetSourceFilesOk

`func (o *CodebaseContext) GetSourceFilesOk() (*[]CodebaseFile, bool)`

GetSourceFilesOk returns a tuple with the SourceFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFiles

`func (o *CodebaseContext) SetSourceFiles(v []CodebaseFile)`

SetSourceFiles sets SourceFiles field to given value.

### HasSourceFiles

`func (o *CodebaseContext) HasSourceFiles() bool`

HasSourceFiles returns a boolean if a field has been set.

### GetTestFiles

`func (o *CodebaseContext) GetTestFiles() []CodebaseFile`

GetTestFiles returns the TestFiles field if non-nil, zero value otherwise.

### GetTestFilesOk

`func (o *CodebaseContext) GetTestFilesOk() (*[]CodebaseFile, bool)`

GetTestFilesOk returns a tuple with the TestFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFiles

`func (o *CodebaseContext) SetTestFiles(v []CodebaseFile)`

SetTestFiles sets TestFiles field to given value.

### HasTestFiles

`func (o *CodebaseContext) HasTestFiles() bool`

HasTestFiles returns a boolean if a field has been set.

### GetFramework

`func (o *CodebaseContext) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *CodebaseContext) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *CodebaseContext) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *CodebaseContext) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetTestCommand

`func (o *CodebaseContext) GetTestCommand() string`

GetTestCommand returns the TestCommand field if non-nil, zero value otherwise.

### GetTestCommandOk

`func (o *CodebaseContext) GetTestCommandOk() (*string, bool)`

GetTestCommandOk returns a tuple with the TestCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCommand

`func (o *CodebaseContext) SetTestCommand(v string)`

SetTestCommand sets TestCommand field to given value.

### HasTestCommand

`func (o *CodebaseContext) HasTestCommand() bool`

HasTestCommand returns a boolean if a field has been set.

### GetImportGraphHints

`func (o *CodebaseContext) GetImportGraphHints() []string`

GetImportGraphHints returns the ImportGraphHints field if non-nil, zero value otherwise.

### GetImportGraphHintsOk

`func (o *CodebaseContext) GetImportGraphHintsOk() (*[]string, bool)`

GetImportGraphHintsOk returns a tuple with the ImportGraphHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportGraphHints

`func (o *CodebaseContext) SetImportGraphHints(v []string)`

SetImportGraphHints sets ImportGraphHints field to given value.

### HasImportGraphHints

`func (o *CodebaseContext) HasImportGraphHints() bool`

HasImportGraphHints returns a boolean if a field has been set.

### GetStyleExamples

`func (o *CodebaseContext) GetStyleExamples() []CodebaseFile`

GetStyleExamples returns the StyleExamples field if non-nil, zero value otherwise.

### GetStyleExamplesOk

`func (o *CodebaseContext) GetStyleExamplesOk() (*[]CodebaseFile, bool)`

GetStyleExamplesOk returns a tuple with the StyleExamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleExamples

`func (o *CodebaseContext) SetStyleExamples(v []CodebaseFile)`

SetStyleExamples sets StyleExamples field to given value.

### HasStyleExamples

`func (o *CodebaseContext) HasStyleExamples() bool`

HasStyleExamples returns a boolean if a field has been set.

### GetFixturesAndMocks

`func (o *CodebaseContext) GetFixturesAndMocks() []CodebaseFile`

GetFixturesAndMocks returns the FixturesAndMocks field if non-nil, zero value otherwise.

### GetFixturesAndMocksOk

`func (o *CodebaseContext) GetFixturesAndMocksOk() (*[]CodebaseFile, bool)`

GetFixturesAndMocksOk returns a tuple with the FixturesAndMocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixturesAndMocks

`func (o *CodebaseContext) SetFixturesAndMocks(v []CodebaseFile)`

SetFixturesAndMocks sets FixturesAndMocks field to given value.

### HasFixturesAndMocks

`func (o *CodebaseContext) HasFixturesAndMocks() bool`

HasFixturesAndMocks returns a boolean if a field has been set.

### GetCandidateFilePaths

`func (o *CodebaseContext) GetCandidateFilePaths() []string`

GetCandidateFilePaths returns the CandidateFilePaths field if non-nil, zero value otherwise.

### GetCandidateFilePathsOk

`func (o *CodebaseContext) GetCandidateFilePathsOk() (*[]string, bool)`

GetCandidateFilePathsOk returns a tuple with the CandidateFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateFilePaths

`func (o *CodebaseContext) SetCandidateFilePaths(v []string)`

SetCandidateFilePaths sets CandidateFilePaths field to given value.

### HasCandidateFilePaths

`func (o *CodebaseContext) HasCandidateFilePaths() bool`

HasCandidateFilePaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


