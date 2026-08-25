# V1UploadRunAttachmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | sha256 hex content hash — what a result&#39;s &#x60;attachments&#x60; entry carries. | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Mime** | Pointer to **string** | Detected from the first bytes of the part, not from the client. | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Presigned download URL; short-lived. Re-resolve a hash via GetRunAttachment instead of storing this. | [optional] 

## Methods

### NewV1UploadRunAttachmentResult

`func NewV1UploadRunAttachmentResult() *V1UploadRunAttachmentResult`

NewV1UploadRunAttachmentResult instantiates a new V1UploadRunAttachmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UploadRunAttachmentResultWithDefaults

`func NewV1UploadRunAttachmentResultWithDefaults() *V1UploadRunAttachmentResult`

NewV1UploadRunAttachmentResultWithDefaults instantiates a new V1UploadRunAttachmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *V1UploadRunAttachmentResult) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V1UploadRunAttachmentResult) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V1UploadRunAttachmentResult) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *V1UploadRunAttachmentResult) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetFilename

`func (o *V1UploadRunAttachmentResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *V1UploadRunAttachmentResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *V1UploadRunAttachmentResult) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *V1UploadRunAttachmentResult) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMime

`func (o *V1UploadRunAttachmentResult) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *V1UploadRunAttachmentResult) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *V1UploadRunAttachmentResult) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *V1UploadRunAttachmentResult) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetSize

`func (o *V1UploadRunAttachmentResult) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V1UploadRunAttachmentResult) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V1UploadRunAttachmentResult) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *V1UploadRunAttachmentResult) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUrl

`func (o *V1UploadRunAttachmentResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1UploadRunAttachmentResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1UploadRunAttachmentResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1UploadRunAttachmentResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


