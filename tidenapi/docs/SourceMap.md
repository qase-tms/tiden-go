# SourceMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**DebugId** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ReleaseName** | Pointer to **string** |  | [optional] 
**ByteSize** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | NOTE: metadata only — never a download URL (source maps stay private). | [optional] 

## Methods

### NewSourceMap

`func NewSourceMap() *SourceMap`

NewSourceMap instantiates a new SourceMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceMapWithDefaults

`func NewSourceMapWithDefaults() *SourceMap`

NewSourceMapWithDefaults instantiates a new SourceMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceMap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceMap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceMap) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceMap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *SourceMap) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SourceMap) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SourceMap) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SourceMap) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetDebugId

`func (o *SourceMap) GetDebugId() string`

GetDebugId returns the DebugId field if non-nil, zero value otherwise.

### GetDebugIdOk

`func (o *SourceMap) GetDebugIdOk() (*string, bool)`

GetDebugIdOk returns a tuple with the DebugId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugId

`func (o *SourceMap) SetDebugId(v string)`

SetDebugId sets DebugId field to given value.

### HasDebugId

`func (o *SourceMap) HasDebugId() bool`

HasDebugId returns a boolean if a field has been set.

### GetFileName

`func (o *SourceMap) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SourceMap) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SourceMap) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SourceMap) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetReleaseName

`func (o *SourceMap) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *SourceMap) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *SourceMap) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *SourceMap) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetByteSize

`func (o *SourceMap) GetByteSize() string`

GetByteSize returns the ByteSize field if non-nil, zero value otherwise.

### GetByteSizeOk

`func (o *SourceMap) GetByteSizeOk() (*string, bool)`

GetByteSizeOk returns a tuple with the ByteSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteSize

`func (o *SourceMap) SetByteSize(v string)`

SetByteSize sets ByteSize field to given value.

### HasByteSize

`func (o *SourceMap) HasByteSize() bool`

HasByteSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SourceMap) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceMap) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceMap) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceMap) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


