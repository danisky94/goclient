# TipoNotificacaoUtilizador

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefTipoNotificacao** | Pointer to **NullableInt64** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**PorEmail** | Pointer to **NullableBool** |  | [optional] 
**PorSms** | Pointer to **NullableBool** |  | [optional] 
**TipoNotificacao** | Pointer to [**TipoNotificacao**](TipoNotificacao.md) |  | [optional] 
**Utilizador** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 

## Methods

### NewTipoNotificacaoUtilizador

`func NewTipoNotificacaoUtilizador() *TipoNotificacaoUtilizador`

NewTipoNotificacaoUtilizador instantiates a new TipoNotificacaoUtilizador object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTipoNotificacaoUtilizadorWithDefaults

`func NewTipoNotificacaoUtilizadorWithDefaults() *TipoNotificacaoUtilizador`

NewTipoNotificacaoUtilizadorWithDefaults instantiates a new TipoNotificacaoUtilizador object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TipoNotificacaoUtilizador) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TipoNotificacaoUtilizador) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TipoNotificacaoUtilizador) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TipoNotificacaoUtilizador) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *TipoNotificacaoUtilizador) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *TipoNotificacaoUtilizador) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *TipoNotificacaoUtilizador) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *TipoNotificacaoUtilizador) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *TipoNotificacaoUtilizador) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *TipoNotificacaoUtilizador) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *TipoNotificacaoUtilizador) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *TipoNotificacaoUtilizador) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *TipoNotificacaoUtilizador) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *TipoNotificacaoUtilizador) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *TipoNotificacaoUtilizador) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *TipoNotificacaoUtilizador) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *TipoNotificacaoUtilizador) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *TipoNotificacaoUtilizador) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *TipoNotificacaoUtilizador) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *TipoNotificacaoUtilizador) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefTipoNotificacao

`func (o *TipoNotificacaoUtilizador) GetRefTipoNotificacao() int64`

GetRefTipoNotificacao returns the RefTipoNotificacao field if non-nil, zero value otherwise.

### GetRefTipoNotificacaoOk

`func (o *TipoNotificacaoUtilizador) GetRefTipoNotificacaoOk() (*int64, bool)`

GetRefTipoNotificacaoOk returns a tuple with the RefTipoNotificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTipoNotificacao

`func (o *TipoNotificacaoUtilizador) SetRefTipoNotificacao(v int64)`

SetRefTipoNotificacao sets RefTipoNotificacao field to given value.

### HasRefTipoNotificacao

`func (o *TipoNotificacaoUtilizador) HasRefTipoNotificacao() bool`

HasRefTipoNotificacao returns a boolean if a field has been set.

### SetRefTipoNotificacaoNil

`func (o *TipoNotificacaoUtilizador) SetRefTipoNotificacaoNil(b bool)`

 SetRefTipoNotificacaoNil sets the value for RefTipoNotificacao to be an explicit nil

### UnsetRefTipoNotificacao
`func (o *TipoNotificacaoUtilizador) UnsetRefTipoNotificacao()`

UnsetRefTipoNotificacao ensures that no value is present for RefTipoNotificacao, not even an explicit nil
### GetRefUtilizador

`func (o *TipoNotificacaoUtilizador) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *TipoNotificacaoUtilizador) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *TipoNotificacaoUtilizador) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *TipoNotificacaoUtilizador) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *TipoNotificacaoUtilizador) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *TipoNotificacaoUtilizador) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetPorEmail

`func (o *TipoNotificacaoUtilizador) GetPorEmail() bool`

GetPorEmail returns the PorEmail field if non-nil, zero value otherwise.

### GetPorEmailOk

`func (o *TipoNotificacaoUtilizador) GetPorEmailOk() (*bool, bool)`

GetPorEmailOk returns a tuple with the PorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorEmail

`func (o *TipoNotificacaoUtilizador) SetPorEmail(v bool)`

SetPorEmail sets PorEmail field to given value.

### HasPorEmail

`func (o *TipoNotificacaoUtilizador) HasPorEmail() bool`

HasPorEmail returns a boolean if a field has been set.

### SetPorEmailNil

`func (o *TipoNotificacaoUtilizador) SetPorEmailNil(b bool)`

 SetPorEmailNil sets the value for PorEmail to be an explicit nil

### UnsetPorEmail
`func (o *TipoNotificacaoUtilizador) UnsetPorEmail()`

UnsetPorEmail ensures that no value is present for PorEmail, not even an explicit nil
### GetPorSms

`func (o *TipoNotificacaoUtilizador) GetPorSms() bool`

GetPorSms returns the PorSms field if non-nil, zero value otherwise.

### GetPorSmsOk

`func (o *TipoNotificacaoUtilizador) GetPorSmsOk() (*bool, bool)`

GetPorSmsOk returns a tuple with the PorSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorSms

`func (o *TipoNotificacaoUtilizador) SetPorSms(v bool)`

SetPorSms sets PorSms field to given value.

### HasPorSms

`func (o *TipoNotificacaoUtilizador) HasPorSms() bool`

HasPorSms returns a boolean if a field has been set.

### SetPorSmsNil

`func (o *TipoNotificacaoUtilizador) SetPorSmsNil(b bool)`

 SetPorSmsNil sets the value for PorSms to be an explicit nil

### UnsetPorSms
`func (o *TipoNotificacaoUtilizador) UnsetPorSms()`

UnsetPorSms ensures that no value is present for PorSms, not even an explicit nil
### GetTipoNotificacao

`func (o *TipoNotificacaoUtilizador) GetTipoNotificacao() TipoNotificacao`

GetTipoNotificacao returns the TipoNotificacao field if non-nil, zero value otherwise.

### GetTipoNotificacaoOk

`func (o *TipoNotificacaoUtilizador) GetTipoNotificacaoOk() (*TipoNotificacao, bool)`

GetTipoNotificacaoOk returns a tuple with the TipoNotificacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoNotificacao

`func (o *TipoNotificacaoUtilizador) SetTipoNotificacao(v TipoNotificacao)`

SetTipoNotificacao sets TipoNotificacao field to given value.

### HasTipoNotificacao

`func (o *TipoNotificacaoUtilizador) HasTipoNotificacao() bool`

HasTipoNotificacao returns a boolean if a field has been set.

### GetUtilizador

`func (o *TipoNotificacaoUtilizador) GetUtilizador() Utilizador`

GetUtilizador returns the Utilizador field if non-nil, zero value otherwise.

### GetUtilizadorOk

`func (o *TipoNotificacaoUtilizador) GetUtilizadorOk() (*Utilizador, bool)`

GetUtilizadorOk returns a tuple with the Utilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizador

`func (o *TipoNotificacaoUtilizador) SetUtilizador(v Utilizador)`

SetUtilizador sets Utilizador field to given value.

### HasUtilizador

`func (o *TipoNotificacaoUtilizador) HasUtilizador() bool`

HasUtilizador returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


