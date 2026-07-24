# CreateProductBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** | Optional team owner. Empty &#x3D; workspace-owned. | [optional] 

## Methods

### NewCreateProductBody

`func NewCreateProductBody() *CreateProductBody`

NewCreateProductBody instantiates a new CreateProductBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductBodyWithDefaults

`func NewCreateProductBodyWithDefaults() *CreateProductBody`

NewCreateProductBodyWithDefaults instantiates a new CreateProductBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProductBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProductBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateProductBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *CreateProductBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateProductBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateProductBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateProductBody) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTeamId

`func (o *CreateProductBody) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateProductBody) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateProductBody) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CreateProductBody) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


