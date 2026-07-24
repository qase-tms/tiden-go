# TestParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**[]TestParameter**](TestParameter.md) |  | [optional] 

## Methods

### NewTestParameterGroup

`func NewTestParameterGroup() *TestParameterGroup`

NewTestParameterGroup instantiates a new TestParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestParameterGroupWithDefaults

`func NewTestParameterGroupWithDefaults() *TestParameterGroup`

NewTestParameterGroupWithDefaults instantiates a new TestParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestParameterGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestParameterGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestParameterGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TestParameterGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TestParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestParameterGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestParameterGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *TestParameterGroup) GetValues() []TestParameter`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TestParameterGroup) GetValuesOk() (*[]TestParameter, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TestParameterGroup) SetValues(v []TestParameter)`

SetValues sets Values field to given value.

### HasValues

`func (o *TestParameterGroup) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


