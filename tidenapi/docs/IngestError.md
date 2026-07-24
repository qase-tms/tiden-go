# IngestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | INVALID_*, AMBIGUOUS_SEQ, SUITE_REQUIRED, ... | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewIngestError

`func NewIngestError() *IngestError`

NewIngestError instantiates a new IngestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestErrorWithDefaults

`func NewIngestErrorWithDefaults() *IngestError`

NewIngestErrorWithDefaults instantiates a new IngestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *IngestError) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IngestError) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IngestError) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IngestError) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCode

`func (o *IngestError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IngestError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IngestError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IngestError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *IngestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IngestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IngestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IngestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


