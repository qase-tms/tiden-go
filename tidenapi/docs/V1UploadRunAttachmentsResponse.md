# V1UploadRunAttachmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** | Always true on success (reporter-client compatibility). | [optional] 
**Result** | Pointer to [**[]V1UploadRunAttachmentResult**](V1UploadRunAttachmentResult.md) |  | [optional] 

## Methods

### NewV1UploadRunAttachmentsResponse

`func NewV1UploadRunAttachmentsResponse() *V1UploadRunAttachmentsResponse`

NewV1UploadRunAttachmentsResponse instantiates a new V1UploadRunAttachmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UploadRunAttachmentsResponseWithDefaults

`func NewV1UploadRunAttachmentsResponseWithDefaults() *V1UploadRunAttachmentsResponse`

NewV1UploadRunAttachmentsResponseWithDefaults instantiates a new V1UploadRunAttachmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V1UploadRunAttachmentsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1UploadRunAttachmentsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1UploadRunAttachmentsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1UploadRunAttachmentsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *V1UploadRunAttachmentsResponse) GetResult() []V1UploadRunAttachmentResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1UploadRunAttachmentsResponse) GetResultOk() (*[]V1UploadRunAttachmentResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1UploadRunAttachmentsResponse) SetResult(v []V1UploadRunAttachmentResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1UploadRunAttachmentsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


