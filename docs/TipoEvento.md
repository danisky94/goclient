# TipoEvento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Tipo** | Pointer to [**TiposTipoEvento**](TiposTipoEvento.md) |  | [optional] 
**RefTipoNotificacao** | Pointer to **NullableInt64** |  | [optional] 
**Notificacao** | Pointer to [**TipoNotificacao**](TipoNotificacao.md) |  | [optional] 
**Eventos** | Pointer to [**[]Evento**](Evento.md) |  | [optional] 

## Methods

### NewTipoEvento

`func NewTipoEvento() *TipoEvento`

NewTipoEvento instantiates a new TipoEvento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTipoEventoWithDefaults

`func NewTipoEventoWithDefaults() *TipoEvento`

NewTipoEventoWithDefaults instantiates a new TipoEvento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TipoEvento) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TipoEvento) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TipoEvento) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TipoEvento) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *TipoEvento) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TipoEvento) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TipoEvento) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TipoEvento) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *TipoEvento) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *TipoEvento) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *TipoEvento) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *TipoEvento) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *TipoEvento) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *TipoEvento) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *TipoEvento) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *TipoEvento) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *TipoEvento) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *TipoEvento) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *TipoEvento) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *TipoEvento) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *TipoEvento) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *TipoEvento) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *TipoEvento) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *TipoEvento) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *TipoEvento) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *TipoEvento) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetTipo

`func (o *TipoEvento) GetTipo() TiposTipoEvento`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *TipoEvento) GetTipoOk() (*TiposTipoEvento, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *TipoEvento) SetTipo(v TiposTipoEvento)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *TipoEvento) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetRefTipoNotificacao

`func (o *TipoEvento) GetRefTipoNotificacao() int64`

GetRefTipoNotificacao returns the RefTipoNotificacao field if non-nil, zero value otherwise.

### GetRefTipoNotificacaoOk

`func (o *TipoEvento) GetRefTipoNotificacaoOk() (*int64, bool)`

GetRefTipoNotificacaoOk returns a tuple with the RefTipoNotificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTipoNotificacao

`func (o *TipoEvento) SetRefTipoNotificacao(v int64)`

SetRefTipoNotificacao sets RefTipoNotificacao field to given value.

### HasRefTipoNotificacao

`func (o *TipoEvento) HasRefTipoNotificacao() bool`

HasRefTipoNotificacao returns a boolean if a field has been set.

### SetRefTipoNotificacaoNil

`func (o *TipoEvento) SetRefTipoNotificacaoNil(b bool)`

 SetRefTipoNotificacaoNil sets the value for RefTipoNotificacao to be an explicit nil

### UnsetRefTipoNotificacao
`func (o *TipoEvento) UnsetRefTipoNotificacao()`

UnsetRefTipoNotificacao ensures that no value is present for RefTipoNotificacao, not even an explicit nil
### GetNotificacao

`func (o *TipoEvento) GetNotificacao() TipoNotificacao`

GetNotificacao returns the Notificacao field if non-nil, zero value otherwise.

### GetNotificacaoOk

`func (o *TipoEvento) GetNotificacaoOk() (*TipoNotificacao, bool)`

GetNotificacaoOk returns a tuple with the Notificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificacao

`func (o *TipoEvento) SetNotificacao(v TipoNotificacao)`

SetNotificacao sets Notificacao field to given value.

### HasNotificacao

`func (o *TipoEvento) HasNotificacao() bool`

HasNotificacao returns a boolean if a field has been set.

### GetEventos

`func (o *TipoEvento) GetEventos() []Evento`

GetEventos returns the Eventos field if non-nil, zero value otherwise.

### GetEventosOk

`func (o *TipoEvento) GetEventosOk() (*[]Evento, bool)`

GetEventosOk returns a tuple with the Eventos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventos

`func (o *TipoEvento) SetEventos(v []Evento)`

SetEventos sets Eventos field to given value.

### HasEventos

`func (o *TipoEvento) HasEventos() bool`

HasEventos returns a boolean if a field has been set.

### SetEventosNil

`func (o *TipoEvento) SetEventosNil(b bool)`

 SetEventosNil sets the value for Eventos to be an explicit nil

### UnsetEventos
`func (o *TipoEvento) UnsetEventos()`

UnsetEventos ensures that no value is present for Eventos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


