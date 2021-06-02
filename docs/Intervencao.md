# Intervencao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefTarefa** | Pointer to **NullableInt64** |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**Estado** | Pointer to [**IntervencaoEstados**](IntervencaoEstados.md) |  | [optional] 
**NomeUtilizador** | Pointer to **NullableString** |  | [optional] [readonly] 
**RefTarefaNavigation** | Pointer to [**Tarefa**](Tarefa.md) |  | [optional] 
**RefUtilizadorNavigation** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 

## Methods

### NewIntervencao

`func NewIntervencao() *Intervencao`

NewIntervencao instantiates a new Intervencao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntervencaoWithDefaults

`func NewIntervencaoWithDefaults() *Intervencao`

NewIntervencaoWithDefaults instantiates a new Intervencao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intervencao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intervencao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intervencao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Intervencao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Intervencao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Intervencao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Intervencao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Intervencao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Intervencao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Intervencao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Intervencao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Intervencao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Intervencao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Intervencao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Intervencao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Intervencao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Intervencao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Intervencao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Intervencao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Intervencao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefTarefa

`func (o *Intervencao) GetRefTarefa() int64`

GetRefTarefa returns the RefTarefa field if non-nil, zero value otherwise.

### GetRefTarefaOk

`func (o *Intervencao) GetRefTarefaOk() (*int64, bool)`

GetRefTarefaOk returns a tuple with the RefTarefa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTarefa

`func (o *Intervencao) SetRefTarefa(v int64)`

SetRefTarefa sets RefTarefa field to given value.

### HasRefTarefa

`func (o *Intervencao) HasRefTarefa() bool`

HasRefTarefa returns a boolean if a field has been set.

### SetRefTarefaNil

`func (o *Intervencao) SetRefTarefaNil(b bool)`

 SetRefTarefaNil sets the value for RefTarefa to be an explicit nil

### UnsetRefTarefa
`func (o *Intervencao) UnsetRefTarefa()`

UnsetRefTarefa ensures that no value is present for RefTarefa, not even an explicit nil
### GetDescricao

`func (o *Intervencao) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *Intervencao) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *Intervencao) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *Intervencao) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *Intervencao) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *Intervencao) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetData

`func (o *Intervencao) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Intervencao) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Intervencao) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *Intervencao) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Intervencao) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Intervencao) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetRefUtilizador

`func (o *Intervencao) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *Intervencao) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *Intervencao) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *Intervencao) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *Intervencao) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *Intervencao) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetEstado

`func (o *Intervencao) GetEstado() IntervencaoEstados`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Intervencao) GetEstadoOk() (*IntervencaoEstados, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Intervencao) SetEstado(v IntervencaoEstados)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Intervencao) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetNomeUtilizador

`func (o *Intervencao) GetNomeUtilizador() string`

GetNomeUtilizador returns the NomeUtilizador field if non-nil, zero value otherwise.

### GetNomeUtilizadorOk

`func (o *Intervencao) GetNomeUtilizadorOk() (*string, bool)`

GetNomeUtilizadorOk returns a tuple with the NomeUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomeUtilizador

`func (o *Intervencao) SetNomeUtilizador(v string)`

SetNomeUtilizador sets NomeUtilizador field to given value.

### HasNomeUtilizador

`func (o *Intervencao) HasNomeUtilizador() bool`

HasNomeUtilizador returns a boolean if a field has been set.

### SetNomeUtilizadorNil

`func (o *Intervencao) SetNomeUtilizadorNil(b bool)`

 SetNomeUtilizadorNil sets the value for NomeUtilizador to be an explicit nil

### UnsetNomeUtilizador
`func (o *Intervencao) UnsetNomeUtilizador()`

UnsetNomeUtilizador ensures that no value is present for NomeUtilizador, not even an explicit nil
### GetRefTarefaNavigation

`func (o *Intervencao) GetRefTarefaNavigation() Tarefa`

GetRefTarefaNavigation returns the RefTarefaNavigation field if non-nil, zero value otherwise.

### GetRefTarefaNavigationOk

`func (o *Intervencao) GetRefTarefaNavigationOk() (*Tarefa, bool)`

GetRefTarefaNavigationOk returns a tuple with the RefTarefaNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTarefaNavigation

`func (o *Intervencao) SetRefTarefaNavigation(v Tarefa)`

SetRefTarefaNavigation sets RefTarefaNavigation field to given value.

### HasRefTarefaNavigation

`func (o *Intervencao) HasRefTarefaNavigation() bool`

HasRefTarefaNavigation returns a boolean if a field has been set.

### GetRefUtilizadorNavigation

`func (o *Intervencao) GetRefUtilizadorNavigation() Utilizador`

GetRefUtilizadorNavigation returns the RefUtilizadorNavigation field if non-nil, zero value otherwise.

### GetRefUtilizadorNavigationOk

`func (o *Intervencao) GetRefUtilizadorNavigationOk() (*Utilizador, bool)`

GetRefUtilizadorNavigationOk returns a tuple with the RefUtilizadorNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizadorNavigation

`func (o *Intervencao) SetRefUtilizadorNavigation(v Utilizador)`

SetRefUtilizadorNavigation sets RefUtilizadorNavigation field to given value.

### HasRefUtilizadorNavigation

`func (o *Intervencao) HasRefUtilizadorNavigation() bool`

HasRefUtilizadorNavigation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


