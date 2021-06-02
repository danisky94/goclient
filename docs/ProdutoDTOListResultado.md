# ProdutoDTOListResultado

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**[]ProdutoDTO**](ProdutoDTO.md) |  | [optional] 

## Methods

### NewProdutoDTOListResultado

`func NewProdutoDTOListResultado() *ProdutoDTOListResultado`

NewProdutoDTOListResultado instantiates a new ProdutoDTOListResultado object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProdutoDTOListResultadoWithDefaults

`func NewProdutoDTOListResultadoWithDefaults() *ProdutoDTOListResultado`

NewProdutoDTOListResultadoWithDefaults instantiates a new ProdutoDTOListResultado object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *ProdutoDTOListResultado) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ProdutoDTOListResultado) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ProdutoDTOListResultado) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *ProdutoDTOListResultado) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetMessage

`func (o *ProdutoDTOListResultado) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProdutoDTOListResultado) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProdutoDTOListResultado) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProdutoDTOListResultado) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ProdutoDTOListResultado) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProdutoDTOListResultado) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValue

`func (o *ProdutoDTOListResultado) GetValue() []ProdutoDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProdutoDTOListResultado) GetValueOk() (*[]ProdutoDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProdutoDTOListResultado) SetValue(v []ProdutoDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProdutoDTOListResultado) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProdutoDTOListResultado) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProdutoDTOListResultado) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


