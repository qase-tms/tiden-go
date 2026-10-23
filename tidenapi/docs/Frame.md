# Frame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | Pointer to **string** |  | [optional] 
**AbsPath** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Lineno** | Pointer to **int32** |  | [optional] 
**Colno** | Pointer to **int32** |  | [optional] 
**InApp** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ContextLine** | Pointer to **string** |  | [optional] 
**PreContext** | Pointer to **[]string** |  | [optional] 
**PostContext** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFrame

`func NewFrame() *Frame`

NewFrame instantiates a new Frame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameWithDefaults

`func NewFrameWithDefaults() *Frame`

NewFrameWithDefaults instantiates a new Frame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *Frame) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Frame) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Frame) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *Frame) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetAbsPath

`func (o *Frame) GetAbsPath() string`

GetAbsPath returns the AbsPath field if non-nil, zero value otherwise.

### GetAbsPathOk

`func (o *Frame) GetAbsPathOk() (*string, bool)`

GetAbsPathOk returns a tuple with the AbsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsPath

`func (o *Frame) SetAbsPath(v string)`

SetAbsPath sets AbsPath field to given value.

### HasAbsPath

`func (o *Frame) HasAbsPath() bool`

HasAbsPath returns a boolean if a field has been set.

### GetFilename

`func (o *Frame) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Frame) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Frame) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Frame) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetLineno

`func (o *Frame) GetLineno() int32`

GetLineno returns the Lineno field if non-nil, zero value otherwise.

### GetLinenoOk

`func (o *Frame) GetLinenoOk() (*int32, bool)`

GetLinenoOk returns a tuple with the Lineno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineno

`func (o *Frame) SetLineno(v int32)`

SetLineno sets Lineno field to given value.

### HasLineno

`func (o *Frame) HasLineno() bool`

HasLineno returns a boolean if a field has been set.

### GetColno

`func (o *Frame) GetColno() int32`

GetColno returns the Colno field if non-nil, zero value otherwise.

### GetColnoOk

`func (o *Frame) GetColnoOk() (*int32, bool)`

GetColnoOk returns a tuple with the Colno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColno

`func (o *Frame) SetColno(v int32)`

SetColno sets Colno field to given value.

### HasColno

`func (o *Frame) HasColno() bool`

HasColno returns a boolean if a field has been set.

### GetInApp

`func (o *Frame) GetInApp() bool`

GetInApp returns the InApp field if non-nil, zero value otherwise.

### GetInAppOk

`func (o *Frame) GetInAppOk() (*bool, bool)`

GetInAppOk returns a tuple with the InApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInApp

`func (o *Frame) SetInApp(v bool)`

SetInApp sets InApp field to given value.

### HasInApp

`func (o *Frame) HasInApp() bool`

HasInApp returns a boolean if a field has been set.

### GetState

`func (o *Frame) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Frame) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Frame) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Frame) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContextLine

`func (o *Frame) GetContextLine() string`

GetContextLine returns the ContextLine field if non-nil, zero value otherwise.

### GetContextLineOk

`func (o *Frame) GetContextLineOk() (*string, bool)`

GetContextLineOk returns a tuple with the ContextLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLine

`func (o *Frame) SetContextLine(v string)`

SetContextLine sets ContextLine field to given value.

### HasContextLine

`func (o *Frame) HasContextLine() bool`

HasContextLine returns a boolean if a field has been set.

### GetPreContext

`func (o *Frame) GetPreContext() []string`

GetPreContext returns the PreContext field if non-nil, zero value otherwise.

### GetPreContextOk

`func (o *Frame) GetPreContextOk() (*[]string, bool)`

GetPreContextOk returns a tuple with the PreContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreContext

`func (o *Frame) SetPreContext(v []string)`

SetPreContext sets PreContext field to given value.

### HasPreContext

`func (o *Frame) HasPreContext() bool`

HasPreContext returns a boolean if a field has been set.

### GetPostContext

`func (o *Frame) GetPostContext() []string`

GetPostContext returns the PostContext field if non-nil, zero value otherwise.

### GetPostContextOk

`func (o *Frame) GetPostContextOk() (*[]string, bool)`

GetPostContextOk returns a tuple with the PostContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostContext

`func (o *Frame) SetPostContext(v []string)`

SetPostContext sets PostContext field to given value.

### HasPostContext

`func (o *Frame) HasPostContext() bool`

HasPostContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


