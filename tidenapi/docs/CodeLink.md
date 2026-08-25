# CodeLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**BaseSha** | Pointer to **string** |  | [optional] 
**HeadSha** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCodeLink

`func NewCodeLink() *CodeLink`

NewCodeLink instantiates a new CodeLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLinkWithDefaults

`func NewCodeLinkWithDefaults() *CodeLink`

NewCodeLinkWithDefaults instantiates a new CodeLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodeLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBranchId

`func (o *CodeLink) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *CodeLink) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *CodeLink) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *CodeLink) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetProductId

`func (o *CodeLink) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CodeLink) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CodeLink) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CodeLink) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetKind

`func (o *CodeLink) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CodeLink) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CodeLink) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CodeLink) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRepository

`func (o *CodeLink) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CodeLink) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CodeLink) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *CodeLink) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRef

`func (o *CodeLink) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodeLink) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodeLink) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CodeLink) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUrl

`func (o *CodeLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CodeLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *CodeLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CodeLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CodeLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CodeLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetState

`func (o *CodeLink) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodeLink) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodeLink) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CodeLink) HasState() bool`

HasState returns a boolean if a field has been set.

### GetBaseSha

`func (o *CodeLink) GetBaseSha() string`

GetBaseSha returns the BaseSha field if non-nil, zero value otherwise.

### GetBaseShaOk

`func (o *CodeLink) GetBaseShaOk() (*string, bool)`

GetBaseShaOk returns a tuple with the BaseSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSha

`func (o *CodeLink) SetBaseSha(v string)`

SetBaseSha sets BaseSha field to given value.

### HasBaseSha

`func (o *CodeLink) HasBaseSha() bool`

HasBaseSha returns a boolean if a field has been set.

### GetHeadSha

`func (o *CodeLink) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *CodeLink) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *CodeLink) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.

### HasHeadSha

`func (o *CodeLink) HasHeadSha() bool`

HasHeadSha returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CodeLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CodeLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CodeLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodeLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodeLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CodeLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


