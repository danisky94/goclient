# Venda

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **time.Time** |  | [optional] 
**Validade** | Pointer to **NullableTime** |  | [optional] 
**DataEstado** | Pointer to **NullableTime** |  | [optional] 
**DataLevantamento** | Pointer to **NullableTime** |  | [optional] 
**Origem** | Pointer to [**TipoOrigem**](TipoOrigem.md) |  | [optional] 
**Estado** | Pointer to [**TipoVendaEstado**](TipoVendaEstado.md) |  | [optional] 
**RefCartaoProduto** | Pointer to **NullableInt64** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Vasilhame** | Pointer to **NullableBool** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**Gvm** | Pointer to [**Gvm**](Gvm.md) |  | [optional] 
**Produto** | Pointer to [**Produto**](Produto.md) |  | [optional] 
**CartaoProduto** | Pointer to [**CartaoProduto**](CartaoProduto.md) |  | [optional] 

## Methods

### NewVenda

`func NewVenda() *Venda`

NewVenda instantiates a new Venda object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendaWithDefaults

`func NewVendaWithDefaults() *Venda`

NewVendaWithDefaults instantiates a new Venda object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Venda) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Venda) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Venda) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Venda) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Venda) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Venda) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Venda) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Venda) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Venda) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Venda) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Venda) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Venda) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Venda) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Venda) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Venda) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Venda) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Venda) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Venda) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Venda) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Venda) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetData

`func (o *Venda) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Venda) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Venda) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *Venda) HasData() bool`

HasData returns a boolean if a field has been set.

### GetValidade

`func (o *Venda) GetValidade() time.Time`

GetValidade returns the Validade field if non-nil, zero value otherwise.

### GetValidadeOk

`func (o *Venda) GetValidadeOk() (*time.Time, bool)`

GetValidadeOk returns a tuple with the Validade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidade

`func (o *Venda) SetValidade(v time.Time)`

SetValidade sets Validade field to given value.

### HasValidade

`func (o *Venda) HasValidade() bool`

HasValidade returns a boolean if a field has been set.

### SetValidadeNil

`func (o *Venda) SetValidadeNil(b bool)`

 SetValidadeNil sets the value for Validade to be an explicit nil

### UnsetValidade
`func (o *Venda) UnsetValidade()`

UnsetValidade ensures that no value is present for Validade, not even an explicit nil
### GetDataEstado

`func (o *Venda) GetDataEstado() time.Time`

GetDataEstado returns the DataEstado field if non-nil, zero value otherwise.

### GetDataEstadoOk

`func (o *Venda) GetDataEstadoOk() (*time.Time, bool)`

GetDataEstadoOk returns a tuple with the DataEstado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEstado

`func (o *Venda) SetDataEstado(v time.Time)`

SetDataEstado sets DataEstado field to given value.

### HasDataEstado

`func (o *Venda) HasDataEstado() bool`

HasDataEstado returns a boolean if a field has been set.

### SetDataEstadoNil

`func (o *Venda) SetDataEstadoNil(b bool)`

 SetDataEstadoNil sets the value for DataEstado to be an explicit nil

### UnsetDataEstado
`func (o *Venda) UnsetDataEstado()`

UnsetDataEstado ensures that no value is present for DataEstado, not even an explicit nil
### GetDataLevantamento

`func (o *Venda) GetDataLevantamento() time.Time`

GetDataLevantamento returns the DataLevantamento field if non-nil, zero value otherwise.

### GetDataLevantamentoOk

`func (o *Venda) GetDataLevantamentoOk() (*time.Time, bool)`

GetDataLevantamentoOk returns a tuple with the DataLevantamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLevantamento

`func (o *Venda) SetDataLevantamento(v time.Time)`

SetDataLevantamento sets DataLevantamento field to given value.

### HasDataLevantamento

`func (o *Venda) HasDataLevantamento() bool`

HasDataLevantamento returns a boolean if a field has been set.

### SetDataLevantamentoNil

`func (o *Venda) SetDataLevantamentoNil(b bool)`

 SetDataLevantamentoNil sets the value for DataLevantamento to be an explicit nil

### UnsetDataLevantamento
`func (o *Venda) UnsetDataLevantamento()`

UnsetDataLevantamento ensures that no value is present for DataLevantamento, not even an explicit nil
### GetOrigem

`func (o *Venda) GetOrigem() TipoOrigem`

GetOrigem returns the Origem field if non-nil, zero value otherwise.

### GetOrigemOk

`func (o *Venda) GetOrigemOk() (*TipoOrigem, bool)`

GetOrigemOk returns a tuple with the Origem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigem

`func (o *Venda) SetOrigem(v TipoOrigem)`

SetOrigem sets Origem field to given value.

### HasOrigem

`func (o *Venda) HasOrigem() bool`

HasOrigem returns a boolean if a field has been set.

### GetEstado

`func (o *Venda) GetEstado() TipoVendaEstado`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Venda) GetEstadoOk() (*TipoVendaEstado, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Venda) SetEstado(v TipoVendaEstado)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Venda) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetRefCartaoProduto

`func (o *Venda) GetRefCartaoProduto() int64`

GetRefCartaoProduto returns the RefCartaoProduto field if non-nil, zero value otherwise.

### GetRefCartaoProdutoOk

`func (o *Venda) GetRefCartaoProdutoOk() (*int64, bool)`

GetRefCartaoProdutoOk returns a tuple with the RefCartaoProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCartaoProduto

`func (o *Venda) SetRefCartaoProduto(v int64)`

SetRefCartaoProduto sets RefCartaoProduto field to given value.

### HasRefCartaoProduto

`func (o *Venda) HasRefCartaoProduto() bool`

HasRefCartaoProduto returns a boolean if a field has been set.

### SetRefCartaoProdutoNil

`func (o *Venda) SetRefCartaoProdutoNil(b bool)`

 SetRefCartaoProdutoNil sets the value for RefCartaoProduto to be an explicit nil

### UnsetRefCartaoProduto
`func (o *Venda) UnsetRefCartaoProduto()`

UnsetRefCartaoProduto ensures that no value is present for RefCartaoProduto, not even an explicit nil
### GetRefGvm

`func (o *Venda) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *Venda) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *Venda) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *Venda) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *Venda) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *Venda) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetRefProduto

`func (o *Venda) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *Venda) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *Venda) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *Venda) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *Venda) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *Venda) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetVasilhame

`func (o *Venda) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *Venda) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *Venda) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *Venda) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### SetVasilhameNil

`func (o *Venda) SetVasilhameNil(b bool)`

 SetVasilhameNil sets the value for Vasilhame to be an explicit nil

### UnsetVasilhame
`func (o *Venda) UnsetVasilhame()`

UnsetVasilhame ensures that no value is present for Vasilhame, not even an explicit nil
### GetToken

`func (o *Venda) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Venda) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Venda) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Venda) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *Venda) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *Venda) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetGvm

`func (o *Venda) GetGvm() Gvm`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *Venda) GetGvmOk() (*Gvm, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *Venda) SetGvm(v Gvm)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *Venda) HasGvm() bool`

HasGvm returns a boolean if a field has been set.

### GetProduto

`func (o *Venda) GetProduto() Produto`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *Venda) GetProdutoOk() (*Produto, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *Venda) SetProduto(v Produto)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *Venda) HasProduto() bool`

HasProduto returns a boolean if a field has been set.

### GetCartaoProduto

`func (o *Venda) GetCartaoProduto() CartaoProduto`

GetCartaoProduto returns the CartaoProduto field if non-nil, zero value otherwise.

### GetCartaoProdutoOk

`func (o *Venda) GetCartaoProdutoOk() (*CartaoProduto, bool)`

GetCartaoProdutoOk returns a tuple with the CartaoProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartaoProduto

`func (o *Venda) SetCartaoProduto(v CartaoProduto)`

SetCartaoProduto sets CartaoProduto field to given value.

### HasCartaoProduto

`func (o *Venda) HasCartaoProduto() bool`

HasCartaoProduto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


