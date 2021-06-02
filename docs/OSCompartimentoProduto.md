# OSCompartimentoProduto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compartimento** | Pointer to [**CompartimentoVM**](CompartimentoVM.md) |  | [optional] 
**Produto** | Pointer to [**ProdutoVM**](ProdutoVM.md) |  | [optional] 
**Ocupado** | Pointer to **bool** |  | [optional] 

## Methods

### NewOSCompartimentoProduto

`func NewOSCompartimentoProduto() *OSCompartimentoProduto`

NewOSCompartimentoProduto instantiates a new OSCompartimentoProduto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSCompartimentoProdutoWithDefaults

`func NewOSCompartimentoProdutoWithDefaults() *OSCompartimentoProduto`

NewOSCompartimentoProdutoWithDefaults instantiates a new OSCompartimentoProduto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompartimento

`func (o *OSCompartimentoProduto) GetCompartimento() CompartimentoVM`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *OSCompartimentoProduto) GetCompartimentoOk() (*CompartimentoVM, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *OSCompartimentoProduto) SetCompartimento(v CompartimentoVM)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *OSCompartimentoProduto) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.

### GetProduto

`func (o *OSCompartimentoProduto) GetProduto() ProdutoVM`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *OSCompartimentoProduto) GetProdutoOk() (*ProdutoVM, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *OSCompartimentoProduto) SetProduto(v ProdutoVM)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *OSCompartimentoProduto) HasProduto() bool`

HasProduto returns a boolean if a field has been set.

### GetOcupado

`func (o *OSCompartimentoProduto) GetOcupado() bool`

GetOcupado returns the Ocupado field if non-nil, zero value otherwise.

### GetOcupadoOk

`func (o *OSCompartimentoProduto) GetOcupadoOk() (*bool, bool)`

GetOcupadoOk returns a tuple with the Ocupado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcupado

`func (o *OSCompartimentoProduto) SetOcupado(v bool)`

SetOcupado sets Ocupado field to given value.

### HasOcupado

`func (o *OSCompartimentoProduto) HasOcupado() bool`

HasOcupado returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


