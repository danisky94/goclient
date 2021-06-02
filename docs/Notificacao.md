# Notificacao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**Tipo** | Pointer to [**NotificacaoTipo**](NotificacaoTipo.md) |  | [optional] 
**Assunto** | Pointer to **NullableString** |  | [optional] 
**Mensagem** | Pointer to **NullableString** |  | [optional] 
**Estado** | Pointer to [**EstadoNotificacao**](EstadoNotificacao.md) |  | [optional] 
**NoticacaoUtilizadors** | Pointer to [**[]NotificacaoUtilizador**](NotificacaoUtilizador.md) |  | [optional] 

## Methods

### NewNotificacao

`func NewNotificacao() *Notificacao`

NewNotificacao instantiates a new Notificacao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificacaoWithDefaults

`func NewNotificacaoWithDefaults() *Notificacao`

NewNotificacaoWithDefaults instantiates a new Notificacao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notificacao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notificacao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notificacao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Notificacao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Notificacao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Notificacao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Notificacao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Notificacao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Notificacao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Notificacao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Notificacao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Notificacao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Notificacao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Notificacao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Notificacao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Notificacao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Notificacao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Notificacao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Notificacao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Notificacao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetData

`func (o *Notificacao) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Notificacao) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Notificacao) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *Notificacao) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Notificacao) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Notificacao) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTipo

`func (o *Notificacao) GetTipo() NotificacaoTipo`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Notificacao) GetTipoOk() (*NotificacaoTipo, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Notificacao) SetTipo(v NotificacaoTipo)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Notificacao) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetAssunto

`func (o *Notificacao) GetAssunto() string`

GetAssunto returns the Assunto field if non-nil, zero value otherwise.

### GetAssuntoOk

`func (o *Notificacao) GetAssuntoOk() (*string, bool)`

GetAssuntoOk returns a tuple with the Assunto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssunto

`func (o *Notificacao) SetAssunto(v string)`

SetAssunto sets Assunto field to given value.

### HasAssunto

`func (o *Notificacao) HasAssunto() bool`

HasAssunto returns a boolean if a field has been set.

### SetAssuntoNil

`func (o *Notificacao) SetAssuntoNil(b bool)`

 SetAssuntoNil sets the value for Assunto to be an explicit nil

### UnsetAssunto
`func (o *Notificacao) UnsetAssunto()`

UnsetAssunto ensures that no value is present for Assunto, not even an explicit nil
### GetMensagem

`func (o *Notificacao) GetMensagem() string`

GetMensagem returns the Mensagem field if non-nil, zero value otherwise.

### GetMensagemOk

`func (o *Notificacao) GetMensagemOk() (*string, bool)`

GetMensagemOk returns a tuple with the Mensagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensagem

`func (o *Notificacao) SetMensagem(v string)`

SetMensagem sets Mensagem field to given value.

### HasMensagem

`func (o *Notificacao) HasMensagem() bool`

HasMensagem returns a boolean if a field has been set.

### SetMensagemNil

`func (o *Notificacao) SetMensagemNil(b bool)`

 SetMensagemNil sets the value for Mensagem to be an explicit nil

### UnsetMensagem
`func (o *Notificacao) UnsetMensagem()`

UnsetMensagem ensures that no value is present for Mensagem, not even an explicit nil
### GetEstado

`func (o *Notificacao) GetEstado() EstadoNotificacao`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Notificacao) GetEstadoOk() (*EstadoNotificacao, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Notificacao) SetEstado(v EstadoNotificacao)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Notificacao) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetNoticacaoUtilizadors

`func (o *Notificacao) GetNoticacaoUtilizadors() []NotificacaoUtilizador`

GetNoticacaoUtilizadors returns the NoticacaoUtilizadors field if non-nil, zero value otherwise.

### GetNoticacaoUtilizadorsOk

`func (o *Notificacao) GetNoticacaoUtilizadorsOk() (*[]NotificacaoUtilizador, bool)`

GetNoticacaoUtilizadorsOk returns a tuple with the NoticacaoUtilizadors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticacaoUtilizadors

`func (o *Notificacao) SetNoticacaoUtilizadors(v []NotificacaoUtilizador)`

SetNoticacaoUtilizadors sets NoticacaoUtilizadors field to given value.

### HasNoticacaoUtilizadors

`func (o *Notificacao) HasNoticacaoUtilizadors() bool`

HasNoticacaoUtilizadors returns a boolean if a field has been set.

### SetNoticacaoUtilizadorsNil

`func (o *Notificacao) SetNoticacaoUtilizadorsNil(b bool)`

 SetNoticacaoUtilizadorsNil sets the value for NoticacaoUtilizadors to be an explicit nil

### UnsetNoticacaoUtilizadors
`func (o *Notificacao) UnsetNoticacaoUtilizadors()`

UnsetNoticacaoUtilizadors ensures that no value is present for NoticacaoUtilizadors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


