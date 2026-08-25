# RecordSessionSettlementBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]IntentSessionSettlementItem**](IntentSessionSettlementItem.md) | items and nothing_to_distill are mutually exclusive; items are merged by item_id onto the stored settlement, nothing_to_distill replaces it wholesale. See IntentSessionService.RecordSessionSettlement for the full semantics. | [optional] 
**NothingToDistill** | Pointer to [**IntentSessionNothingToDistill**](IntentSessionNothingToDistill.md) |  | [optional] 

## Methods

### NewRecordSessionSettlementBody

`func NewRecordSessionSettlementBody() *RecordSessionSettlementBody`

NewRecordSessionSettlementBody instantiates a new RecordSessionSettlementBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSessionSettlementBodyWithDefaults

`func NewRecordSessionSettlementBodyWithDefaults() *RecordSessionSettlementBody`

NewRecordSessionSettlementBodyWithDefaults instantiates a new RecordSessionSettlementBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RecordSessionSettlementBody) GetItems() []IntentSessionSettlementItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecordSessionSettlementBody) GetItemsOk() (*[]IntentSessionSettlementItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecordSessionSettlementBody) SetItems(v []IntentSessionSettlementItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *RecordSessionSettlementBody) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNothingToDistill

`func (o *RecordSessionSettlementBody) GetNothingToDistill() IntentSessionNothingToDistill`

GetNothingToDistill returns the NothingToDistill field if non-nil, zero value otherwise.

### GetNothingToDistillOk

`func (o *RecordSessionSettlementBody) GetNothingToDistillOk() (*IntentSessionNothingToDistill, bool)`

GetNothingToDistillOk returns a tuple with the NothingToDistill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNothingToDistill

`func (o *RecordSessionSettlementBody) SetNothingToDistill(v IntentSessionNothingToDistill)`

SetNothingToDistill sets NothingToDistill field to given value.

### HasNothingToDistill

`func (o *RecordSessionSettlementBody) HasNothingToDistill() bool`

HasNothingToDistill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


