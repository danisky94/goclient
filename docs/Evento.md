# Evento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefTipoEvento** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**Mensagem** | Pointer to **NullableString** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**Gvm** | Pointer to [**Gvm**](Gvm.md) |  | [optional] 
**TipoEvento** | Pointer to [**TipoEvento**](TipoEvento.md) |  | [optional] 
**Compartimento** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 

## Methods

### NewEvento

`func NewEvento() *Evento`

NewEvento instantiates a new Evento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventoWithDefaults

`func NewEventoWithDefaults() *Evento`

NewEventoWithDefaults instantiates a new Evento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Evento) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Evento) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Evento) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Evento) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Evento) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Evento) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Evento) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Evento) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Evento) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Evento) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Evento) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Evento) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Evento) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Evento) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Evento) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Evento) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Evento) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Evento) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Evento) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Evento) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefTipoEvento

`func (o *Evento) GetRefTipoEvento() int64`

GetRefTipoEvento returns the RefTipoEvento field if non-nil, zero value otherwise.

### GetRefTipoEventoOk

`func (o *Evento) GetRefTipoEventoOk() (*int64, bool)`

GetRefTipoEventoOk returns a tuple with the RefTipoEvento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTipoEvento

`func (o *Evento) SetRefTipoEvento(v int64)`

SetRefTipoEvento sets RefTipoEvento field to given value.

### HasRefTipoEvento

`func (o *Evento) HasRefTipoEvento() bool`

HasRefTipoEvento returns a boolean if a field has been set.

### SetRefTipoEventoNil

`func (o *Evento) SetRefTipoEventoNil(b bool)`

 SetRefTipoEventoNil sets the value for RefTipoEvento to be an explicit nil

### UnsetRefTipoEvento
`func (o *Evento) UnsetRefTipoEvento()`

UnsetRefTipoEvento ensures that no value is present for RefTipoEvento, not even an explicit nil
### GetRefCompartimento

`func (o *Evento) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *Evento) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *Evento) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *Evento) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *Evento) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *Evento) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetData

`func (o *Evento) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Evento) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Evento) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *Evento) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Evento) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Evento) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMensagem

`func (o *Evento) GetMensagem() string`

GetMensagem returns the Mensagem field if non-nil, zero value otherwise.

### GetMensagemOk

`func (o *Evento) GetMensagemOk() (*string, bool)`

GetMensagemOk returns a tuple with the Mensagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensagem

`func (o *Evento) SetMensagem(v string)`

SetMensagem sets Mensagem field to given value.

### HasMensagem

`func (o *Evento) HasMensagem() bool`

HasMensagem returns a boolean if a field has been set.

### SetMensagemNil

`func (o *Evento) SetMensagemNil(b bool)`

 SetMensagemNil sets the value for Mensagem to be an explicit nil

### UnsetMensagem
`func (o *Evento) UnsetMensagem()`

UnsetMensagem ensures that no value is present for Mensagem, not even an explicit nil
### GetRefGvm

`func (o *Evento) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *Evento) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *Evento) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *Evento) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *Evento) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *Evento) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetGvm

`func (o *Evento) GetGvm() Gvm`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *Evento) GetGvmOk() (*Gvm, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *Evento) SetGvm(v Gvm)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *Evento) HasGvm() bool`

HasGvm returns a boolean if a field has been set.

### GetTipoEvento

`func (o *Evento) GetTipoEvento() TipoEvento`

GetTipoEvento returns the TipoEvento field if non-nil, zero value otherwise.

### GetTipoEventoOk

`func (o *Evento) GetTipoEventoOk() (*TipoEvento, bool)`

GetTipoEventoOk returns a tuple with the TipoEvento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoEvento

`func (o *Evento) SetTipoEvento(v TipoEvento)`

SetTipoEvento sets TipoEvento field to given value.

### HasTipoEvento

`func (o *Evento) HasTipoEvento() bool`

HasTipoEvento returns a boolean if a field has been set.

### GetCompartimento

`func (o *Evento) GetCompartimento() Compartimento`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *Evento) GetCompartimentoOk() (*Compartimento, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *Evento) SetCompartimento(v Compartimento)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *Evento) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


