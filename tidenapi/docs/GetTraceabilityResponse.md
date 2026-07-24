# GetTraceabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matrix** | Pointer to [**TraceabilityMatrix**](TraceabilityMatrix.md) |  | [optional] 

## Methods

### NewGetTraceabilityResponse

`func NewGetTraceabilityResponse() *GetTraceabilityResponse`

NewGetTraceabilityResponse instantiates a new GetTraceabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTraceabilityResponseWithDefaults

`func NewGetTraceabilityResponseWithDefaults() *GetTraceabilityResponse`

NewGetTraceabilityResponseWithDefaults instantiates a new GetTraceabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatrix

`func (o *GetTraceabilityResponse) GetMatrix() TraceabilityMatrix`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *GetTraceabilityResponse) GetMatrixOk() (*TraceabilityMatrix, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *GetTraceabilityResponse) SetMatrix(v TraceabilityMatrix)`

SetMatrix sets Matrix field to given value.

### HasMatrix

`func (o *GetTraceabilityResponse) HasMatrix() bool`

HasMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


