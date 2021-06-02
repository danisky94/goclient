# CustomVenda

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Produto** | Pointer to [**ProdutoVM**](ProdutoVM.md) |  | [optional] 
**Gvm** | Pointer to [**GvmVM**](GvmVM.md) |  | [optional] 
**Vasilhame** | Pointer to **NullableBool** |  | [optional] 
**RefCompartimentoProduto** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimentoVasilhame** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewCustomVenda

`func NewCustomVenda() *CustomVenda`

NewCustomVenda instantiates a new CustomVenda object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomVendaWithDefaults

`func NewCustomVendaWithDefaults() *CustomVenda`

NewCustomVendaWithDefaults instantiates a new CustomVenda object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduto

`func (o *CustomVenda) GetProduto() ProdutoVM`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *CustomVenda) GetProdutoOk() (*ProdutoVM, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *CustomVenda) SetProduto(v ProdutoVM)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *CustomVenda) HasProduto() bool`

HasProduto returns a boolean if a field has been set.

### GetGvm

`func (o *CustomVenda) GetGvm() GvmVM`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *CustomVenda) GetGvmOk() (*GvmVM, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *CustomVenda) SetGvm(v GvmVM)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *CustomVenda) HasGvm() bool`

HasGvm returns a boolean if a field has been set.

### GetVasilhame

`func (o *CustomVenda) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *CustomVenda) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *CustomVenda) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *CustomVenda) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### SetVasilhameNil

`func (o *CustomVenda) SetVasilhameNil(b bool)`

 SetVasilhameNil sets the value for Vasilhame to be an explicit nil

### UnsetVasilhame
`func (o *CustomVenda) UnsetVasilhame()`

UnsetVasilhame ensures that no value is present for Vasilhame, not even an explicit nil
### GetRefCompartimentoProduto

`func (o *CustomVenda) GetRefCompartimentoProduto() int64`

GetRefCompartimentoProduto returns the RefCompartimentoProduto field if non-nil, zero value otherwise.

### GetRefCompartimentoProdutoOk

`func (o *CustomVenda) GetRefCompartimentoProdutoOk() (*int64, bool)`

GetRefCompartimentoProdutoOk returns a tuple with the RefCompartimentoProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimentoProduto

`func (o *CustomVenda) SetRefCompartimentoProduto(v int64)`

SetRefCompartimentoProduto sets RefCompartimentoProduto field to given value.

### HasRefCompartimentoProduto

`func (o *CustomVenda) HasRefCompartimentoProduto() bool`

HasRefCompartimentoProduto returns a boolean if a field has been set.

### SetRefCompartimentoProdutoNil

`func (o *CustomVenda) SetRefCompartimentoProdutoNil(b bool)`

 SetRefCompartimentoProdutoNil sets the value for RefCompartimentoProduto to be an explicit nil

### UnsetRefCompartimentoProduto
`func (o *CustomVenda) UnsetRefCompartimentoProduto()`

UnsetRefCompartimentoProduto ensures that no value is present for RefCompartimentoProduto, not even an explicit nil
### GetRefCompartimentoVasilhame

`func (o *CustomVenda) GetRefCompartimentoVasilhame() int64`

GetRefCompartimentoVasilhame returns the RefCompartimentoVasilhame field if non-nil, zero value otherwise.

### GetRefCompartimentoVasilhameOk

`func (o *CustomVenda) GetRefCompartimentoVasilhameOk() (*int64, bool)`

GetRefCompartimentoVasilhameOk returns a tuple with the RefCompartimentoVasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimentoVasilhame

`func (o *CustomVenda) SetRefCompartimentoVasilhame(v int64)`

SetRefCompartimentoVasilhame sets RefCompartimentoVasilhame field to given value.

### HasRefCompartimentoVasilhame

`func (o *CustomVenda) HasRefCompartimentoVasilhame() bool`

HasRefCompartimentoVasilhame returns a boolean if a field has been set.

### SetRefCompartimentoVasilhameNil

`func (o *CustomVenda) SetRefCompartimentoVasilhameNil(b bool)`

 SetRefCompartimentoVasilhameNil sets the value for RefCompartimentoVasilhame to be an explicit nil

### UnsetRefCompartimentoVasilhame
`func (o *CustomVenda) UnsetRefCompartimentoVasilhame()`

UnsetRefCompartimentoVasilhame ensures that no value is present for RefCompartimentoVasilhame, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


