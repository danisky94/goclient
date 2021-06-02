# NotificacaoUtilizador

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefNotificacao** | Pointer to **NullableInt64** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**DataEnvio** | Pointer to **NullableTime** |  | [optional] 
**Estado** | Pointer to [**EstadoNotificacaoUtilizador**](EstadoNotificacaoUtilizador.md) |  | [optional] 
**Erro** | Pointer to **NullableString** |  | [optional] 
**Notificacao** | Pointer to [**Notificacao**](Notificacao.md) |  | [optional] 
**Utilizador** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 

## Methods

### NewNotificacaoUtilizador

`func NewNotificacaoUtilizador() *NotificacaoUtilizador`

NewNotificacaoUtilizador instantiates a new NotificacaoUtilizador object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificacaoUtilizadorWithDefaults

`func NewNotificacaoUtilizadorWithDefaults() *NotificacaoUtilizador`

NewNotificacaoUtilizadorWithDefaults instantiates a new NotificacaoUtilizador object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificacaoUtilizador) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificacaoUtilizador) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificacaoUtilizador) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NotificacaoUtilizador) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *NotificacaoUtilizador) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *NotificacaoUtilizador) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *NotificacaoUtilizador) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *NotificacaoUtilizador) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *NotificacaoUtilizador) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *NotificacaoUtilizador) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *NotificacaoUtilizador) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *NotificacaoUtilizador) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *NotificacaoUtilizador) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *NotificacaoUtilizador) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *NotificacaoUtilizador) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *NotificacaoUtilizador) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *NotificacaoUtilizador) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *NotificacaoUtilizador) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *NotificacaoUtilizador) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *NotificacaoUtilizador) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefNotificacao

`func (o *NotificacaoUtilizador) GetRefNotificacao() int64`

GetRefNotificacao returns the RefNotificacao field if non-nil, zero value otherwise.

### GetRefNotificacaoOk

`func (o *NotificacaoUtilizador) GetRefNotificacaoOk() (*int64, bool)`

GetRefNotificacaoOk returns a tuple with the RefNotificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNotificacao

`func (o *NotificacaoUtilizador) SetRefNotificacao(v int64)`

SetRefNotificacao sets RefNotificacao field to given value.

### HasRefNotificacao

`func (o *NotificacaoUtilizador) HasRefNotificacao() bool`

HasRefNotificacao returns a boolean if a field has been set.

### SetRefNotificacaoNil

`func (o *NotificacaoUtilizador) SetRefNotificacaoNil(b bool)`

 SetRefNotificacaoNil sets the value for RefNotificacao to be an explicit nil

### UnsetRefNotificacao
`func (o *NotificacaoUtilizador) UnsetRefNotificacao()`

UnsetRefNotificacao ensures that no value is present for RefNotificacao, not even an explicit nil
### GetRefUtilizador

`func (o *NotificacaoUtilizador) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *NotificacaoUtilizador) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *NotificacaoUtilizador) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *NotificacaoUtilizador) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *NotificacaoUtilizador) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *NotificacaoUtilizador) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetDataEnvio

`func (o *NotificacaoUtilizador) GetDataEnvio() time.Time`

GetDataEnvio returns the DataEnvio field if non-nil, zero value otherwise.

### GetDataEnvioOk

`func (o *NotificacaoUtilizador) GetDataEnvioOk() (*time.Time, bool)`

GetDataEnvioOk returns a tuple with the DataEnvio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEnvio

`func (o *NotificacaoUtilizador) SetDataEnvio(v time.Time)`

SetDataEnvio sets DataEnvio field to given value.

### HasDataEnvio

`func (o *NotificacaoUtilizador) HasDataEnvio() bool`

HasDataEnvio returns a boolean if a field has been set.

### SetDataEnvioNil

`func (o *NotificacaoUtilizador) SetDataEnvioNil(b bool)`

 SetDataEnvioNil sets the value for DataEnvio to be an explicit nil

### UnsetDataEnvio
`func (o *NotificacaoUtilizador) UnsetDataEnvio()`

UnsetDataEnvio ensures that no value is present for DataEnvio, not even an explicit nil
### GetEstado

`func (o *NotificacaoUtilizador) GetEstado() EstadoNotificacaoUtilizador`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *NotificacaoUtilizador) GetEstadoOk() (*EstadoNotificacaoUtilizador, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *NotificacaoUtilizador) SetEstado(v EstadoNotificacaoUtilizador)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *NotificacaoUtilizador) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetErro

`func (o *NotificacaoUtilizador) GetErro() string`

GetErro returns the Erro field if non-nil, zero value otherwise.

### GetErroOk

`func (o *NotificacaoUtilizador) GetErroOk() (*string, bool)`

GetErroOk returns a tuple with the Erro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErro

`func (o *NotificacaoUtilizador) SetErro(v string)`

SetErro sets Erro field to given value.

### HasErro

`func (o *NotificacaoUtilizador) HasErro() bool`

HasErro returns a boolean if a field has been set.

### SetErroNil

`func (o *NotificacaoUtilizador) SetErroNil(b bool)`

 SetErroNil sets the value for Erro to be an explicit nil

### UnsetErro
`func (o *NotificacaoUtilizador) UnsetErro()`

UnsetErro ensures that no value is present for Erro, not even an explicit nil
### GetNotificacao

`func (o *NotificacaoUtilizador) GetNotificacao() Notificacao`

GetNotificacao returns the Notificacao field if non-nil, zero value otherwise.

### GetNotificacaoOk

`func (o *NotificacaoUtilizador) GetNotificacaoOk() (*Notificacao, bool)`

GetNotificacaoOk returns a tuple with the Notificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificacao

`func (o *NotificacaoUtilizador) SetNotificacao(v Notificacao)`

SetNotificacao sets Notificacao field to given value.

### HasNotificacao

`func (o *NotificacaoUtilizador) HasNotificacao() bool`

HasNotificacao returns a boolean if a field has been set.

### GetUtilizador

`func (o *NotificacaoUtilizador) GetUtilizador() Utilizador`

GetUtilizador returns the Utilizador field if non-nil, zero value otherwise.

### GetUtilizadorOk

`func (o *NotificacaoUtilizador) GetUtilizadorOk() (*Utilizador, bool)`

GetUtilizadorOk returns a tuple with the Utilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizador

`func (o *NotificacaoUtilizador) SetUtilizador(v Utilizador)`

SetUtilizador sets Utilizador field to given value.

### HasUtilizador

`func (o *NotificacaoUtilizador) HasUtilizador() bool`

HasUtilizador returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


