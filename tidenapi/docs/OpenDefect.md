# OpenDefect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewOpenDefect

`func NewOpenDefect() *OpenDefect`

NewOpenDefect instantiates a new OpenDefect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenDefectWithDefaults

`func NewOpenDefectWithDefaults() *OpenDefect`

NewOpenDefectWithDefaults instantiates a new OpenDefect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *OpenDefect) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *OpenDefect) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *OpenDefect) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *OpenDefect) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCount

`func (o *OpenDefect) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OpenDefect) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OpenDefect) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OpenDefect) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


