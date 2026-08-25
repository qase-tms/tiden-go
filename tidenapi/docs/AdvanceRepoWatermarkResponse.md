# AdvanceRepoWatermarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advanced** | Pointer to **bool** | False when a baseline found an existing row or an empty_sync CAS missed. | [optional] 
**Watermark** | Pointer to [**RepoWatermark**](RepoWatermark.md) |  | [optional] 

## Methods

### NewAdvanceRepoWatermarkResponse

`func NewAdvanceRepoWatermarkResponse() *AdvanceRepoWatermarkResponse`

NewAdvanceRepoWatermarkResponse instantiates a new AdvanceRepoWatermarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvanceRepoWatermarkResponseWithDefaults

`func NewAdvanceRepoWatermarkResponseWithDefaults() *AdvanceRepoWatermarkResponse`

NewAdvanceRepoWatermarkResponseWithDefaults instantiates a new AdvanceRepoWatermarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanced

`func (o *AdvanceRepoWatermarkResponse) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *AdvanceRepoWatermarkResponse) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *AdvanceRepoWatermarkResponse) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *AdvanceRepoWatermarkResponse) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetWatermark

`func (o *AdvanceRepoWatermarkResponse) GetWatermark() RepoWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *AdvanceRepoWatermarkResponse) GetWatermarkOk() (*RepoWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *AdvanceRepoWatermarkResponse) SetWatermark(v RepoWatermark)`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *AdvanceRepoWatermarkResponse) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


