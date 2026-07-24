# IngestTestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | Pointer to [**IngestStats**](IngestStats.md) |  | [optional] 
**Errors** | Pointer to [**[]IngestError**](IngestError.md) | Per-entry validation errors. Populated when the response carries a 422 status — the gRPC error also carries the same payload via google.rpc.Status details so non-gateway gRPC clients can deserialize them without parsing the gateway response body. | [optional] 

## Methods

### NewIngestTestsResponse

`func NewIngestTestsResponse() *IngestTestsResponse`

NewIngestTestsResponse instantiates a new IngestTestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestTestsResponseWithDefaults

`func NewIngestTestsResponseWithDefaults() *IngestTestsResponse`

NewIngestTestsResponseWithDefaults instantiates a new IngestTestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStats

`func (o *IngestTestsResponse) GetStats() IngestStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *IngestTestsResponse) GetStatsOk() (*IngestStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *IngestTestsResponse) SetStats(v IngestStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *IngestTestsResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetErrors

`func (o *IngestTestsResponse) GetErrors() []IngestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IngestTestsResponse) GetErrorsOk() (*[]IngestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IngestTestsResponse) SetErrors(v []IngestError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IngestTestsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


