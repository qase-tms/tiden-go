# UpsertBranchCodeLinksBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLinks** | Pointer to [**[]CodeLink**](CodeLink.md) | id/branch_id/product_id/created_at/updated_at on each entry are ignored — the upsert identity is (branch_id from the path, kind, repository, ref). Capped at MaxBranchCodeLinksPerBatch entries per call. | [optional] 

## Methods

### NewUpsertBranchCodeLinksBody

`func NewUpsertBranchCodeLinksBody() *UpsertBranchCodeLinksBody`

NewUpsertBranchCodeLinksBody instantiates a new UpsertBranchCodeLinksBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertBranchCodeLinksBodyWithDefaults

`func NewUpsertBranchCodeLinksBodyWithDefaults() *UpsertBranchCodeLinksBody`

NewUpsertBranchCodeLinksBodyWithDefaults instantiates a new UpsertBranchCodeLinksBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeLinks

`func (o *UpsertBranchCodeLinksBody) GetCodeLinks() []CodeLink`

GetCodeLinks returns the CodeLinks field if non-nil, zero value otherwise.

### GetCodeLinksOk

`func (o *UpsertBranchCodeLinksBody) GetCodeLinksOk() (*[]CodeLink, bool)`

GetCodeLinksOk returns a tuple with the CodeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLinks

`func (o *UpsertBranchCodeLinksBody) SetCodeLinks(v []CodeLink)`

SetCodeLinks sets CodeLinks field to given value.

### HasCodeLinks

`func (o *UpsertBranchCodeLinksBody) HasCodeLinks() bool`

HasCodeLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


