# CompartimentoHistorico

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Evento** | Pointer to **NullableInt32** |  | [optional] 
**Mensagem** | Pointer to **NullableString** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimentoNavigation** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 
**RefProdutoNavigation** | Pointer to [**Produto**](Produto.md) |  | [optional] 

## Methods

### NewCompartimentoHistorico

`func NewCompartimentoHistorico() *CompartimentoHistorico`

NewCompartimentoHistorico instantiates a new CompartimentoHistorico object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartimentoHistoricoWithDefaults

`func NewCompartimentoHistoricoWithDefaults() *CompartimentoHistorico`

NewCompartimentoHistoricoWithDefaults instantiates a new CompartimentoHistorico object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompartimentoHistorico) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompartimentoHistorico) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompartimentoHistorico) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CompartimentoHistorico) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *CompartimentoHistorico) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CompartimentoHistorico) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CompartimentoHistorico) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CompartimentoHistorico) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *CompartimentoHistorico) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *CompartimentoHistorico) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *CompartimentoHistorico) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *CompartimentoHistorico) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *CompartimentoHistorico) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *CompartimentoHistorico) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *CompartimentoHistorico) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *CompartimentoHistorico) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *CompartimentoHistorico) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *CompartimentoHistorico) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *CompartimentoHistorico) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *CompartimentoHistorico) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefCompartimento

`func (o *CompartimentoHistorico) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *CompartimentoHistorico) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *CompartimentoHistorico) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *CompartimentoHistorico) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *CompartimentoHistorico) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *CompartimentoHistorico) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetData

`func (o *CompartimentoHistorico) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompartimentoHistorico) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompartimentoHistorico) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *CompartimentoHistorico) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CompartimentoHistorico) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CompartimentoHistorico) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetRefProduto

`func (o *CompartimentoHistorico) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *CompartimentoHistorico) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *CompartimentoHistorico) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *CompartimentoHistorico) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *CompartimentoHistorico) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *CompartimentoHistorico) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetEvento

`func (o *CompartimentoHistorico) GetEvento() int32`

GetEvento returns the Evento field if non-nil, zero value otherwise.

### GetEventoOk

`func (o *CompartimentoHistorico) GetEventoOk() (*int32, bool)`

GetEventoOk returns a tuple with the Evento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvento

`func (o *CompartimentoHistorico) SetEvento(v int32)`

SetEvento sets Evento field to given value.

### HasEvento

`func (o *CompartimentoHistorico) HasEvento() bool`

HasEvento returns a boolean if a field has been set.

### SetEventoNil

`func (o *CompartimentoHistorico) SetEventoNil(b bool)`

 SetEventoNil sets the value for Evento to be an explicit nil

### UnsetEvento
`func (o *CompartimentoHistorico) UnsetEvento()`

UnsetEvento ensures that no value is present for Evento, not even an explicit nil
### GetMensagem

`func (o *CompartimentoHistorico) GetMensagem() string`

GetMensagem returns the Mensagem field if non-nil, zero value otherwise.

### GetMensagemOk

`func (o *CompartimentoHistorico) GetMensagemOk() (*string, bool)`

GetMensagemOk returns a tuple with the Mensagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensagem

`func (o *CompartimentoHistorico) SetMensagem(v string)`

SetMensagem sets Mensagem field to given value.

### HasMensagem

`func (o *CompartimentoHistorico) HasMensagem() bool`

HasMensagem returns a boolean if a field has been set.

### SetMensagemNil

`func (o *CompartimentoHistorico) SetMensagemNil(b bool)`

 SetMensagemNil sets the value for Mensagem to be an explicit nil

### UnsetMensagem
`func (o *CompartimentoHistorico) UnsetMensagem()`

UnsetMensagem ensures that no value is present for Mensagem, not even an explicit nil
### GetRefUtilizador

`func (o *CompartimentoHistorico) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *CompartimentoHistorico) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *CompartimentoHistorico) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *CompartimentoHistorico) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *CompartimentoHistorico) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *CompartimentoHistorico) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetRefCompartimentoNavigation

`func (o *CompartimentoHistorico) GetRefCompartimentoNavigation() Compartimento`

GetRefCompartimentoNavigation returns the RefCompartimentoNavigation field if non-nil, zero value otherwise.

### GetRefCompartimentoNavigationOk

`func (o *CompartimentoHistorico) GetRefCompartimentoNavigationOk() (*Compartimento, bool)`

GetRefCompartimentoNavigationOk returns a tuple with the RefCompartimentoNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimentoNavigation

`func (o *CompartimentoHistorico) SetRefCompartimentoNavigation(v Compartimento)`

SetRefCompartimentoNavigation sets RefCompartimentoNavigation field to given value.

### HasRefCompartimentoNavigation

`func (o *CompartimentoHistorico) HasRefCompartimentoNavigation() bool`

HasRefCompartimentoNavigation returns a boolean if a field has been set.

### GetRefProdutoNavigation

`func (o *CompartimentoHistorico) GetRefProdutoNavigation() Produto`

GetRefProdutoNavigation returns the RefProdutoNavigation field if non-nil, zero value otherwise.

### GetRefProdutoNavigationOk

`func (o *CompartimentoHistorico) GetRefProdutoNavigationOk() (*Produto, bool)`

GetRefProdutoNavigationOk returns a tuple with the RefProdutoNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProdutoNavigation

`func (o *CompartimentoHistorico) SetRefProdutoNavigation(v Produto)`

SetRefProdutoNavigation sets RefProdutoNavigation field to given value.

### HasRefProdutoNavigation

`func (o *CompartimentoHistorico) HasRefProdutoNavigation() bool`

HasRefProdutoNavigation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


