# IntentSessionSettlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]IntentSessionSettlementItem**](IntentSessionSettlementItem.md) |  | [optional] 
**NothingToDistill** | Pointer to [**IntentSessionNothingToDistill**](IntentSessionNothingToDistill.md) |  | [optional] 
**SettledAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIntentSessionSettlement

`func NewIntentSessionSettlement() *IntentSessionSettlement`

NewIntentSessionSettlement instantiates a new IntentSessionSettlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntentSessionSettlementWithDefaults

`func NewIntentSessionSettlementWithDefaults() *IntentSessionSettlement`

NewIntentSessionSettlementWithDefaults instantiates a new IntentSessionSettlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IntentSessionSettlement) GetItems() []IntentSessionSettlementItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntentSessionSettlement) GetItemsOk() (*[]IntentSessionSettlementItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntentSessionSettlement) SetItems(v []IntentSessionSettlementItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *IntentSessionSettlement) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNothingToDistill

`func (o *IntentSessionSettlement) GetNothingToDistill() IntentSessionNothingToDistill`

GetNothingToDistill returns the NothingToDistill field if non-nil, zero value otherwise.

### GetNothingToDistillOk

`func (o *IntentSessionSettlement) GetNothingToDistillOk() (*IntentSessionNothingToDistill, bool)`

GetNothingToDistillOk returns a tuple with the NothingToDistill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNothingToDistill

`func (o *IntentSessionSettlement) SetNothingToDistill(v IntentSessionNothingToDistill)`

SetNothingToDistill sets NothingToDistill field to given value.

### HasNothingToDistill

`func (o *IntentSessionSettlement) HasNothingToDistill() bool`

HasNothingToDistill returns a boolean if a field has been set.

### GetSettledAt

`func (o *IntentSessionSettlement) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *IntentSessionSettlement) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *IntentSessionSettlement) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *IntentSessionSettlement) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


