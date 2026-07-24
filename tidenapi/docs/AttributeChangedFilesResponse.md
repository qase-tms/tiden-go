# AttributeChangedFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributed** | Pointer to [**[]AttributedChangedFile**](AttributedChangedFile.md) |  | [optional] 
**Unmatched** | Pointer to [**[]ChangedFile**](ChangedFile.md) |  | [optional] 
**ComponentId** | Pointer to **string** | component_id is set when the files resolve to exactly ONE component (requirement.component_id was set to it); empty when they span multiple components (cleared) or none. | [optional] 

## Methods

### NewAttributeChangedFilesResponse

`func NewAttributeChangedFilesResponse() *AttributeChangedFilesResponse`

NewAttributeChangedFilesResponse instantiates a new AttributeChangedFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeChangedFilesResponseWithDefaults

`func NewAttributeChangedFilesResponseWithDefaults() *AttributeChangedFilesResponse`

NewAttributeChangedFilesResponseWithDefaults instantiates a new AttributeChangedFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributed

`func (o *AttributeChangedFilesResponse) GetAttributed() []AttributedChangedFile`

GetAttributed returns the Attributed field if non-nil, zero value otherwise.

### GetAttributedOk

`func (o *AttributeChangedFilesResponse) GetAttributedOk() (*[]AttributedChangedFile, bool)`

GetAttributedOk returns a tuple with the Attributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributed

`func (o *AttributeChangedFilesResponse) SetAttributed(v []AttributedChangedFile)`

SetAttributed sets Attributed field to given value.

### HasAttributed

`func (o *AttributeChangedFilesResponse) HasAttributed() bool`

HasAttributed returns a boolean if a field has been set.

### GetUnmatched

`func (o *AttributeChangedFilesResponse) GetUnmatched() []ChangedFile`

GetUnmatched returns the Unmatched field if non-nil, zero value otherwise.

### GetUnmatchedOk

`func (o *AttributeChangedFilesResponse) GetUnmatchedOk() (*[]ChangedFile, bool)`

GetUnmatchedOk returns a tuple with the Unmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatched

`func (o *AttributeChangedFilesResponse) SetUnmatched(v []ChangedFile)`

SetUnmatched sets Unmatched field to given value.

### HasUnmatched

`func (o *AttributeChangedFilesResponse) HasUnmatched() bool`

HasUnmatched returns a boolean if a field has been set.

### GetComponentId

`func (o *AttributeChangedFilesResponse) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *AttributeChangedFilesResponse) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *AttributeChangedFilesResponse) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *AttributeChangedFilesResponse) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


