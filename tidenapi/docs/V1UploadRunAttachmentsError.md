# V1UploadRunAttachmentsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Always false. | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewV1UploadRunAttachmentsError

`func NewV1UploadRunAttachmentsError() *V1UploadRunAttachmentsError`

NewV1UploadRunAttachmentsError instantiates a new V1UploadRunAttachmentsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UploadRunAttachmentsErrorWithDefaults

`func NewV1UploadRunAttachmentsErrorWithDefaults() *V1UploadRunAttachmentsError`

NewV1UploadRunAttachmentsErrorWithDefaults instantiates a new V1UploadRunAttachmentsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V1UploadRunAttachmentsError) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1UploadRunAttachmentsError) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1UploadRunAttachmentsError) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1UploadRunAttachmentsError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *V1UploadRunAttachmentsError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1UploadRunAttachmentsError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1UploadRunAttachmentsError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1UploadRunAttachmentsError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


