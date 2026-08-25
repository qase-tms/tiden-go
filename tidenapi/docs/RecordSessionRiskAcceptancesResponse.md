# RecordSessionRiskAcceptancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptancesRecorded** | Pointer to **int32** |  | [optional] 
**DeferredRequirements** | Pointer to **int32** |  | [optional] 
**ReplacedRows** | Pointer to **int32** | Rows this call superseded — this session&#39;s equivalent records from an earlier close attempt, replaced rather than stacked. | [optional] 

## Methods

### NewRecordSessionRiskAcceptancesResponse

`func NewRecordSessionRiskAcceptancesResponse() *RecordSessionRiskAcceptancesResponse`

NewRecordSessionRiskAcceptancesResponse instantiates a new RecordSessionRiskAcceptancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSessionRiskAcceptancesResponseWithDefaults

`func NewRecordSessionRiskAcceptancesResponseWithDefaults() *RecordSessionRiskAcceptancesResponse`

NewRecordSessionRiskAcceptancesResponseWithDefaults instantiates a new RecordSessionRiskAcceptancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptancesRecorded

`func (o *RecordSessionRiskAcceptancesResponse) GetAcceptancesRecorded() int32`

GetAcceptancesRecorded returns the AcceptancesRecorded field if non-nil, zero value otherwise.

### GetAcceptancesRecordedOk

`func (o *RecordSessionRiskAcceptancesResponse) GetAcceptancesRecordedOk() (*int32, bool)`

GetAcceptancesRecordedOk returns a tuple with the AcceptancesRecorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptancesRecorded

`func (o *RecordSessionRiskAcceptancesResponse) SetAcceptancesRecorded(v int32)`

SetAcceptancesRecorded sets AcceptancesRecorded field to given value.

### HasAcceptancesRecorded

`func (o *RecordSessionRiskAcceptancesResponse) HasAcceptancesRecorded() bool`

HasAcceptancesRecorded returns a boolean if a field has been set.

### GetDeferredRequirements

`func (o *RecordSessionRiskAcceptancesResponse) GetDeferredRequirements() int32`

GetDeferredRequirements returns the DeferredRequirements field if non-nil, zero value otherwise.

### GetDeferredRequirementsOk

`func (o *RecordSessionRiskAcceptancesResponse) GetDeferredRequirementsOk() (*int32, bool)`

GetDeferredRequirementsOk returns a tuple with the DeferredRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredRequirements

`func (o *RecordSessionRiskAcceptancesResponse) SetDeferredRequirements(v int32)`

SetDeferredRequirements sets DeferredRequirements field to given value.

### HasDeferredRequirements

`func (o *RecordSessionRiskAcceptancesResponse) HasDeferredRequirements() bool`

HasDeferredRequirements returns a boolean if a field has been set.

### GetReplacedRows

`func (o *RecordSessionRiskAcceptancesResponse) GetReplacedRows() int32`

GetReplacedRows returns the ReplacedRows field if non-nil, zero value otherwise.

### GetReplacedRowsOk

`func (o *RecordSessionRiskAcceptancesResponse) GetReplacedRowsOk() (*int32, bool)`

GetReplacedRowsOk returns a tuple with the ReplacedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedRows

`func (o *RecordSessionRiskAcceptancesResponse) SetReplacedRows(v int32)`

SetReplacedRows sets ReplacedRows field to given value.

### HasReplacedRows

`func (o *RecordSessionRiskAcceptancesResponse) HasReplacedRows() bool`

HasReplacedRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


