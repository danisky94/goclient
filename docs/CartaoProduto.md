# CartaoProduto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefCartao** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**Quantidade** | Pointer to **NullableInt32** |  | [optional] 
**Vasilhame** | Pointer to **NullableBool** |  | [optional] 
**RefCartaoNavigation** | Pointer to [**Cartao**](Cartao.md) |  | [optional] 
**Venda** | Pointer to [**[]Venda**](Venda.md) |  | [optional] 

## Methods

### NewCartaoProduto

`func NewCartaoProduto() *CartaoProduto`

NewCartaoProduto instantiates a new CartaoProduto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartaoProdutoWithDefaults

`func NewCartaoProdutoWithDefaults() *CartaoProduto`

NewCartaoProdutoWithDefaults instantiates a new CartaoProduto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartaoProduto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartaoProduto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartaoProduto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CartaoProduto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *CartaoProduto) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CartaoProduto) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CartaoProduto) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CartaoProduto) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *CartaoProduto) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *CartaoProduto) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *CartaoProduto) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *CartaoProduto) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *CartaoProduto) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *CartaoProduto) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *CartaoProduto) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *CartaoProduto) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *CartaoProduto) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *CartaoProduto) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *CartaoProduto) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *CartaoProduto) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefCartao

`func (o *CartaoProduto) GetRefCartao() int64`

GetRefCartao returns the RefCartao field if non-nil, zero value otherwise.

### GetRefCartaoOk

`func (o *CartaoProduto) GetRefCartaoOk() (*int64, bool)`

GetRefCartaoOk returns a tuple with the RefCartao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCartao

`func (o *CartaoProduto) SetRefCartao(v int64)`

SetRefCartao sets RefCartao field to given value.

### HasRefCartao

`func (o *CartaoProduto) HasRefCartao() bool`

HasRefCartao returns a boolean if a field has been set.

### SetRefCartaoNil

`func (o *CartaoProduto) SetRefCartaoNil(b bool)`

 SetRefCartaoNil sets the value for RefCartao to be an explicit nil

### UnsetRefCartao
`func (o *CartaoProduto) UnsetRefCartao()`

UnsetRefCartao ensures that no value is present for RefCartao, not even an explicit nil
### GetRefProduto

`func (o *CartaoProduto) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *CartaoProduto) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *CartaoProduto) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *CartaoProduto) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *CartaoProduto) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *CartaoProduto) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetData

`func (o *CartaoProduto) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CartaoProduto) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CartaoProduto) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *CartaoProduto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CartaoProduto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CartaoProduto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetQuantidade

`func (o *CartaoProduto) GetQuantidade() int32`

GetQuantidade returns the Quantidade field if non-nil, zero value otherwise.

### GetQuantidadeOk

`func (o *CartaoProduto) GetQuantidadeOk() (*int32, bool)`

GetQuantidadeOk returns a tuple with the Quantidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantidade

`func (o *CartaoProduto) SetQuantidade(v int32)`

SetQuantidade sets Quantidade field to given value.

### HasQuantidade

`func (o *CartaoProduto) HasQuantidade() bool`

HasQuantidade returns a boolean if a field has been set.

### SetQuantidadeNil

`func (o *CartaoProduto) SetQuantidadeNil(b bool)`

 SetQuantidadeNil sets the value for Quantidade to be an explicit nil

### UnsetQuantidade
`func (o *CartaoProduto) UnsetQuantidade()`

UnsetQuantidade ensures that no value is present for Quantidade, not even an explicit nil
### GetVasilhame

`func (o *CartaoProduto) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *CartaoProduto) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *CartaoProduto) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *CartaoProduto) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### SetVasilhameNil

`func (o *CartaoProduto) SetVasilhameNil(b bool)`

 SetVasilhameNil sets the value for Vasilhame to be an explicit nil

### UnsetVasilhame
`func (o *CartaoProduto) UnsetVasilhame()`

UnsetVasilhame ensures that no value is present for Vasilhame, not even an explicit nil
### GetRefCartaoNavigation

`func (o *CartaoProduto) GetRefCartaoNavigation() Cartao`

GetRefCartaoNavigation returns the RefCartaoNavigation field if non-nil, zero value otherwise.

### GetRefCartaoNavigationOk

`func (o *CartaoProduto) GetRefCartaoNavigationOk() (*Cartao, bool)`

GetRefCartaoNavigationOk returns a tuple with the RefCartaoNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCartaoNavigation

`func (o *CartaoProduto) SetRefCartaoNavigation(v Cartao)`

SetRefCartaoNavigation sets RefCartaoNavigation field to given value.

### HasRefCartaoNavigation

`func (o *CartaoProduto) HasRefCartaoNavigation() bool`

HasRefCartaoNavigation returns a boolean if a field has been set.

### GetVenda

`func (o *CartaoProduto) GetVenda() []Venda`

GetVenda returns the Venda field if non-nil, zero value otherwise.

### GetVendaOk

`func (o *CartaoProduto) GetVendaOk() (*[]Venda, bool)`

GetVendaOk returns a tuple with the Venda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenda

`func (o *CartaoProduto) SetVenda(v []Venda)`

SetVenda sets Venda field to given value.

### HasVenda

`func (o *CartaoProduto) HasVenda() bool`

HasVenda returns a boolean if a field has been set.

### SetVendaNil

`func (o *CartaoProduto) SetVendaNil(b bool)`

 SetVendaNil sets the value for Venda to be an explicit nil

### UnsetVenda
`func (o *CartaoProduto) UnsetVenda()`

UnsetVenda ensures that no value is present for Venda, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


