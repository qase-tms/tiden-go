# CoveringTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SeqNum** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LastStatus** | Pointer to **string** | last_status is the most recently reported execution status, or \&quot;\&quot; when the test has never run. | [optional] 

## Methods

### NewCoveringTest

`func NewCoveringTest() *CoveringTest`

NewCoveringTest instantiates a new CoveringTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoveringTestWithDefaults

`func NewCoveringTestWithDefaults() *CoveringTest`

NewCoveringTestWithDefaults instantiates a new CoveringTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoveringTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoveringTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoveringTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CoveringTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeqNum

`func (o *CoveringTest) GetSeqNum() int32`

GetSeqNum returns the SeqNum field if non-nil, zero value otherwise.

### GetSeqNumOk

`func (o *CoveringTest) GetSeqNumOk() (*int32, bool)`

GetSeqNumOk returns a tuple with the SeqNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNum

`func (o *CoveringTest) SetSeqNum(v int32)`

SetSeqNum sets SeqNum field to given value.

### HasSeqNum

`func (o *CoveringTest) HasSeqNum() bool`

HasSeqNum returns a boolean if a field has been set.

### GetTitle

`func (o *CoveringTest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CoveringTest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CoveringTest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CoveringTest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLastStatus

`func (o *CoveringTest) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *CoveringTest) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *CoveringTest) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *CoveringTest) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


