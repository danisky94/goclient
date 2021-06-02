# Cartao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**NumeroSerie** | Pointer to **NullableString** |  | [optional] 
**DataProducao** | Pointer to **NullableTime** |  | [optional] 
**Estado** | Pointer to **NullableInt32** |  | [optional] 
**DataEstado** | Pointer to **NullableTime** |  | [optional] 
**Pin** | Pointer to **NullableString** |  | [optional] 
**CartaoProdutos** | Pointer to [**[]CartaoProduto**](CartaoProduto.md) |  | [optional] 

## Methods

### NewCartao

`func NewCartao() *Cartao`

NewCartao instantiates a new Cartao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartaoWithDefaults

`func NewCartaoWithDefaults() *Cartao`

NewCartaoWithDefaults instantiates a new Cartao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cartao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cartao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cartao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Cartao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Cartao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Cartao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Cartao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Cartao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Cartao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Cartao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Cartao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Cartao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Cartao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Cartao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Cartao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Cartao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Cartao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Cartao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Cartao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Cartao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Cartao) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Cartao) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Cartao) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Cartao) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Cartao) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Cartao) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetNumeroSerie

`func (o *Cartao) GetNumeroSerie() string`

GetNumeroSerie returns the NumeroSerie field if non-nil, zero value otherwise.

### GetNumeroSerieOk

`func (o *Cartao) GetNumeroSerieOk() (*string, bool)`

GetNumeroSerieOk returns a tuple with the NumeroSerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroSerie

`func (o *Cartao) SetNumeroSerie(v string)`

SetNumeroSerie sets NumeroSerie field to given value.

### HasNumeroSerie

`func (o *Cartao) HasNumeroSerie() bool`

HasNumeroSerie returns a boolean if a field has been set.

### SetNumeroSerieNil

`func (o *Cartao) SetNumeroSerieNil(b bool)`

 SetNumeroSerieNil sets the value for NumeroSerie to be an explicit nil

### UnsetNumeroSerie
`func (o *Cartao) UnsetNumeroSerie()`

UnsetNumeroSerie ensures that no value is present for NumeroSerie, not even an explicit nil
### GetDataProducao

`func (o *Cartao) GetDataProducao() time.Time`

GetDataProducao returns the DataProducao field if non-nil, zero value otherwise.

### GetDataProducaoOk

`func (o *Cartao) GetDataProducaoOk() (*time.Time, bool)`

GetDataProducaoOk returns a tuple with the DataProducao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProducao

`func (o *Cartao) SetDataProducao(v time.Time)`

SetDataProducao sets DataProducao field to given value.

### HasDataProducao

`func (o *Cartao) HasDataProducao() bool`

HasDataProducao returns a boolean if a field has been set.

### SetDataProducaoNil

`func (o *Cartao) SetDataProducaoNil(b bool)`

 SetDataProducaoNil sets the value for DataProducao to be an explicit nil

### UnsetDataProducao
`func (o *Cartao) UnsetDataProducao()`

UnsetDataProducao ensures that no value is present for DataProducao, not even an explicit nil
### GetEstado

`func (o *Cartao) GetEstado() int32`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Cartao) GetEstadoOk() (*int32, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Cartao) SetEstado(v int32)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Cartao) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### SetEstadoNil

`func (o *Cartao) SetEstadoNil(b bool)`

 SetEstadoNil sets the value for Estado to be an explicit nil

### UnsetEstado
`func (o *Cartao) UnsetEstado()`

UnsetEstado ensures that no value is present for Estado, not even an explicit nil
### GetDataEstado

`func (o *Cartao) GetDataEstado() time.Time`

GetDataEstado returns the DataEstado field if non-nil, zero value otherwise.

### GetDataEstadoOk

`func (o *Cartao) GetDataEstadoOk() (*time.Time, bool)`

GetDataEstadoOk returns a tuple with the DataEstado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEstado

`func (o *Cartao) SetDataEstado(v time.Time)`

SetDataEstado sets DataEstado field to given value.

### HasDataEstado

`func (o *Cartao) HasDataEstado() bool`

HasDataEstado returns a boolean if a field has been set.

### SetDataEstadoNil

`func (o *Cartao) SetDataEstadoNil(b bool)`

 SetDataEstadoNil sets the value for DataEstado to be an explicit nil

### UnsetDataEstado
`func (o *Cartao) UnsetDataEstado()`

UnsetDataEstado ensures that no value is present for DataEstado, not even an explicit nil
### GetPin

`func (o *Cartao) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *Cartao) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *Cartao) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *Cartao) HasPin() bool`

HasPin returns a boolean if a field has been set.

### SetPinNil

`func (o *Cartao) SetPinNil(b bool)`

 SetPinNil sets the value for Pin to be an explicit nil

### UnsetPin
`func (o *Cartao) UnsetPin()`

UnsetPin ensures that no value is present for Pin, not even an explicit nil
### GetCartaoProdutos

`func (o *Cartao) GetCartaoProdutos() []CartaoProduto`

GetCartaoProdutos returns the CartaoProdutos field if non-nil, zero value otherwise.

### GetCartaoProdutosOk

`func (o *Cartao) GetCartaoProdutosOk() (*[]CartaoProduto, bool)`

GetCartaoProdutosOk returns a tuple with the CartaoProdutos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartaoProdutos

`func (o *Cartao) SetCartaoProdutos(v []CartaoProduto)`

SetCartaoProdutos sets CartaoProdutos field to given value.

### HasCartaoProdutos

`func (o *Cartao) HasCartaoProdutos() bool`

HasCartaoProdutos returns a boolean if a field has been set.

### SetCartaoProdutosNil

`func (o *Cartao) SetCartaoProdutosNil(b bool)`

 SetCartaoProdutosNil sets the value for CartaoProdutos to be an explicit nil

### UnsetCartaoProdutos
`func (o *Cartao) UnsetCartaoProdutos()`

UnsetCartaoProdutos ensures that no value is present for CartaoProdutos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


