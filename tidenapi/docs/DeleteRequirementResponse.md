# DeleteRequirementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoryId** | Pointer to **string** | Id of the recorded deletion history event. Pass it to RestoreRequirement to undo the delete. Empty for deletes that record no undoable event. | [optional] 

## Methods

### NewDeleteRequirementResponse

`func NewDeleteRequirementResponse() *DeleteRequirementResponse`

NewDeleteRequirementResponse instantiates a new DeleteRequirementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRequirementResponseWithDefaults

`func NewDeleteRequirementResponseWithDefaults() *DeleteRequirementResponse`

NewDeleteRequirementResponseWithDefaults instantiates a new DeleteRequirementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoryId

`func (o *DeleteRequirementResponse) GetHistoryId() string`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *DeleteRequirementResponse) GetHistoryIdOk() (*string, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *DeleteRequirementResponse) SetHistoryId(v string)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *DeleteRequirementResponse) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


