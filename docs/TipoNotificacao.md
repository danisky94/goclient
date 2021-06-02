# TipoNotificacao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**MensagemEmail** | Pointer to **NullableString** |  | [optional] 
**MenagemSms** | Pointer to **NullableString** |  | [optional] 
**Ativa** | Pointer to **NullableBool** |  | [optional] 
**TipoEventos** | Pointer to [**[]TipoEvento**](TipoEvento.md) |  | [optional] 
**NotificacaoUtilizadores** | Pointer to [**[]TipoNotificacaoUtilizador**](TipoNotificacaoUtilizador.md) |  | [optional] 

## Methods

### NewTipoNotificacao

`func NewTipoNotificacao() *TipoNotificacao`

NewTipoNotificacao instantiates a new TipoNotificacao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTipoNotificacaoWithDefaults

`func NewTipoNotificacaoWithDefaults() *TipoNotificacao`

NewTipoNotificacaoWithDefaults instantiates a new TipoNotificacao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TipoNotificacao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TipoNotificacao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TipoNotificacao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TipoNotificacao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *TipoNotificacao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TipoNotificacao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TipoNotificacao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TipoNotificacao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *TipoNotificacao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *TipoNotificacao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *TipoNotificacao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *TipoNotificacao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *TipoNotificacao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *TipoNotificacao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *TipoNotificacao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *TipoNotificacao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *TipoNotificacao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *TipoNotificacao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *TipoNotificacao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *TipoNotificacao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *TipoNotificacao) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *TipoNotificacao) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *TipoNotificacao) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *TipoNotificacao) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *TipoNotificacao) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *TipoNotificacao) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetMensagemEmail

`func (o *TipoNotificacao) GetMensagemEmail() string`

GetMensagemEmail returns the MensagemEmail field if non-nil, zero value otherwise.

### GetMensagemEmailOk

`func (o *TipoNotificacao) GetMensagemEmailOk() (*string, bool)`

GetMensagemEmailOk returns a tuple with the MensagemEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensagemEmail

`func (o *TipoNotificacao) SetMensagemEmail(v string)`

SetMensagemEmail sets MensagemEmail field to given value.

### HasMensagemEmail

`func (o *TipoNotificacao) HasMensagemEmail() bool`

HasMensagemEmail returns a boolean if a field has been set.

### SetMensagemEmailNil

`func (o *TipoNotificacao) SetMensagemEmailNil(b bool)`

 SetMensagemEmailNil sets the value for MensagemEmail to be an explicit nil

### UnsetMensagemEmail
`func (o *TipoNotificacao) UnsetMensagemEmail()`

UnsetMensagemEmail ensures that no value is present for MensagemEmail, not even an explicit nil
### GetMenagemSms

`func (o *TipoNotificacao) GetMenagemSms() string`

GetMenagemSms returns the MenagemSms field if non-nil, zero value otherwise.

### GetMenagemSmsOk

`func (o *TipoNotificacao) GetMenagemSmsOk() (*string, bool)`

GetMenagemSmsOk returns a tuple with the MenagemSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenagemSms

`func (o *TipoNotificacao) SetMenagemSms(v string)`

SetMenagemSms sets MenagemSms field to given value.

### HasMenagemSms

`func (o *TipoNotificacao) HasMenagemSms() bool`

HasMenagemSms returns a boolean if a field has been set.

### SetMenagemSmsNil

`func (o *TipoNotificacao) SetMenagemSmsNil(b bool)`

 SetMenagemSmsNil sets the value for MenagemSms to be an explicit nil

### UnsetMenagemSms
`func (o *TipoNotificacao) UnsetMenagemSms()`

UnsetMenagemSms ensures that no value is present for MenagemSms, not even an explicit nil
### GetAtiva

`func (o *TipoNotificacao) GetAtiva() bool`

GetAtiva returns the Ativa field if non-nil, zero value otherwise.

### GetAtivaOk

`func (o *TipoNotificacao) GetAtivaOk() (*bool, bool)`

GetAtivaOk returns a tuple with the Ativa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtiva

`func (o *TipoNotificacao) SetAtiva(v bool)`

SetAtiva sets Ativa field to given value.

### HasAtiva

`func (o *TipoNotificacao) HasAtiva() bool`

HasAtiva returns a boolean if a field has been set.

### SetAtivaNil

`func (o *TipoNotificacao) SetAtivaNil(b bool)`

 SetAtivaNil sets the value for Ativa to be an explicit nil

### UnsetAtiva
`func (o *TipoNotificacao) UnsetAtiva()`

UnsetAtiva ensures that no value is present for Ativa, not even an explicit nil
### GetTipoEventos

`func (o *TipoNotificacao) GetTipoEventos() []TipoEvento`

GetTipoEventos returns the TipoEventos field if non-nil, zero value otherwise.

### GetTipoEventosOk

`func (o *TipoNotificacao) GetTipoEventosOk() (*[]TipoEvento, bool)`

GetTipoEventosOk returns a tuple with the TipoEventos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoEventos

`func (o *TipoNotificacao) SetTipoEventos(v []TipoEvento)`

SetTipoEventos sets TipoEventos field to given value.

### HasTipoEventos

`func (o *TipoNotificacao) HasTipoEventos() bool`

HasTipoEventos returns a boolean if a field has been set.

### SetTipoEventosNil

`func (o *TipoNotificacao) SetTipoEventosNil(b bool)`

 SetTipoEventosNil sets the value for TipoEventos to be an explicit nil

### UnsetTipoEventos
`func (o *TipoNotificacao) UnsetTipoEventos()`

UnsetTipoEventos ensures that no value is present for TipoEventos, not even an explicit nil
### GetNotificacaoUtilizadores

`func (o *TipoNotificacao) GetNotificacaoUtilizadores() []TipoNotificacaoUtilizador`

GetNotificacaoUtilizadores returns the NotificacaoUtilizadores field if non-nil, zero value otherwise.

### GetNotificacaoUtilizadoresOk

`func (o *TipoNotificacao) GetNotificacaoUtilizadoresOk() (*[]TipoNotificacaoUtilizador, bool)`

GetNotificacaoUtilizadoresOk returns a tuple with the NotificacaoUtilizadores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificacaoUtilizadores

`func (o *TipoNotificacao) SetNotificacaoUtilizadores(v []TipoNotificacaoUtilizador)`

SetNotificacaoUtilizadores sets NotificacaoUtilizadores field to given value.

### HasNotificacaoUtilizadores

`func (o *TipoNotificacao) HasNotificacaoUtilizadores() bool`

HasNotificacaoUtilizadores returns a boolean if a field has been set.

### SetNotificacaoUtilizadoresNil

`func (o *TipoNotificacao) SetNotificacaoUtilizadoresNil(b bool)`

 SetNotificacaoUtilizadoresNil sets the value for NotificacaoUtilizadores to be an explicit nil

### UnsetNotificacaoUtilizadores
`func (o *TipoNotificacao) UnsetNotificacaoUtilizadores()`

UnsetNotificacaoUtilizadores ensures that no value is present for NotificacaoUtilizadores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


