# Tarefa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**Estado** | Pointer to [**IntervencaoEstados**](IntervencaoEstados.md) |  | [optional] 
**Tipo** | Pointer to [**TipoTarefa**](TipoTarefa.md) |  | [optional] 
**Assunto** | Pointer to **NullableString** |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**CriadoEm** | Pointer to **NullableTime** |  | [optional] 
**FechadoEm** | Pointer to **NullableTime** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**NomeUtilizador** | Pointer to **NullableString** |  | [optional] [readonly] 
**Utilizador** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 
**Intervencaos** | Pointer to [**[]Intervencao**](Intervencao.md) |  | [optional] 

## Methods

### NewTarefa

`func NewTarefa() *Tarefa`

NewTarefa instantiates a new Tarefa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTarefaWithDefaults

`func NewTarefaWithDefaults() *Tarefa`

NewTarefaWithDefaults instantiates a new Tarefa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tarefa) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tarefa) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tarefa) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Tarefa) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Tarefa) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Tarefa) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Tarefa) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Tarefa) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Tarefa) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Tarefa) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Tarefa) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Tarefa) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Tarefa) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Tarefa) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Tarefa) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Tarefa) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Tarefa) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Tarefa) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Tarefa) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Tarefa) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetData

`func (o *Tarefa) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Tarefa) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Tarefa) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *Tarefa) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Tarefa) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Tarefa) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEstado

`func (o *Tarefa) GetEstado() IntervencaoEstados`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Tarefa) GetEstadoOk() (*IntervencaoEstados, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Tarefa) SetEstado(v IntervencaoEstados)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Tarefa) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetTipo

`func (o *Tarefa) GetTipo() TipoTarefa`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Tarefa) GetTipoOk() (*TipoTarefa, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Tarefa) SetTipo(v TipoTarefa)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Tarefa) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetAssunto

`func (o *Tarefa) GetAssunto() string`

GetAssunto returns the Assunto field if non-nil, zero value otherwise.

### GetAssuntoOk

`func (o *Tarefa) GetAssuntoOk() (*string, bool)`

GetAssuntoOk returns a tuple with the Assunto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssunto

`func (o *Tarefa) SetAssunto(v string)`

SetAssunto sets Assunto field to given value.

### HasAssunto

`func (o *Tarefa) HasAssunto() bool`

HasAssunto returns a boolean if a field has been set.

### SetAssuntoNil

`func (o *Tarefa) SetAssuntoNil(b bool)`

 SetAssuntoNil sets the value for Assunto to be an explicit nil

### UnsetAssunto
`func (o *Tarefa) UnsetAssunto()`

UnsetAssunto ensures that no value is present for Assunto, not even an explicit nil
### GetDescricao

`func (o *Tarefa) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *Tarefa) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *Tarefa) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *Tarefa) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *Tarefa) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *Tarefa) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetRefUtilizador

`func (o *Tarefa) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *Tarefa) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *Tarefa) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *Tarefa) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *Tarefa) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *Tarefa) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetCriadoEm

`func (o *Tarefa) GetCriadoEm() time.Time`

GetCriadoEm returns the CriadoEm field if non-nil, zero value otherwise.

### GetCriadoEmOk

`func (o *Tarefa) GetCriadoEmOk() (*time.Time, bool)`

GetCriadoEmOk returns a tuple with the CriadoEm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriadoEm

`func (o *Tarefa) SetCriadoEm(v time.Time)`

SetCriadoEm sets CriadoEm field to given value.

### HasCriadoEm

`func (o *Tarefa) HasCriadoEm() bool`

HasCriadoEm returns a boolean if a field has been set.

### SetCriadoEmNil

`func (o *Tarefa) SetCriadoEmNil(b bool)`

 SetCriadoEmNil sets the value for CriadoEm to be an explicit nil

### UnsetCriadoEm
`func (o *Tarefa) UnsetCriadoEm()`

UnsetCriadoEm ensures that no value is present for CriadoEm, not even an explicit nil
### GetFechadoEm

`func (o *Tarefa) GetFechadoEm() time.Time`

GetFechadoEm returns the FechadoEm field if non-nil, zero value otherwise.

### GetFechadoEmOk

`func (o *Tarefa) GetFechadoEmOk() (*time.Time, bool)`

GetFechadoEmOk returns a tuple with the FechadoEm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechadoEm

`func (o *Tarefa) SetFechadoEm(v time.Time)`

SetFechadoEm sets FechadoEm field to given value.

### HasFechadoEm

`func (o *Tarefa) HasFechadoEm() bool`

HasFechadoEm returns a boolean if a field has been set.

### SetFechadoEmNil

`func (o *Tarefa) SetFechadoEmNil(b bool)`

 SetFechadoEmNil sets the value for FechadoEm to be an explicit nil

### UnsetFechadoEm
`func (o *Tarefa) UnsetFechadoEm()`

UnsetFechadoEm ensures that no value is present for FechadoEm, not even an explicit nil
### GetRefEntidade

`func (o *Tarefa) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *Tarefa) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *Tarefa) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *Tarefa) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *Tarefa) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *Tarefa) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetRefGvm

`func (o *Tarefa) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *Tarefa) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *Tarefa) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *Tarefa) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *Tarefa) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *Tarefa) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetNomeUtilizador

`func (o *Tarefa) GetNomeUtilizador() string`

GetNomeUtilizador returns the NomeUtilizador field if non-nil, zero value otherwise.

### GetNomeUtilizadorOk

`func (o *Tarefa) GetNomeUtilizadorOk() (*string, bool)`

GetNomeUtilizadorOk returns a tuple with the NomeUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomeUtilizador

`func (o *Tarefa) SetNomeUtilizador(v string)`

SetNomeUtilizador sets NomeUtilizador field to given value.

### HasNomeUtilizador

`func (o *Tarefa) HasNomeUtilizador() bool`

HasNomeUtilizador returns a boolean if a field has been set.

### SetNomeUtilizadorNil

`func (o *Tarefa) SetNomeUtilizadorNil(b bool)`

 SetNomeUtilizadorNil sets the value for NomeUtilizador to be an explicit nil

### UnsetNomeUtilizador
`func (o *Tarefa) UnsetNomeUtilizador()`

UnsetNomeUtilizador ensures that no value is present for NomeUtilizador, not even an explicit nil
### GetUtilizador

`func (o *Tarefa) GetUtilizador() Utilizador`

GetUtilizador returns the Utilizador field if non-nil, zero value otherwise.

### GetUtilizadorOk

`func (o *Tarefa) GetUtilizadorOk() (*Utilizador, bool)`

GetUtilizadorOk returns a tuple with the Utilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizador

`func (o *Tarefa) SetUtilizador(v Utilizador)`

SetUtilizador sets Utilizador field to given value.

### HasUtilizador

`func (o *Tarefa) HasUtilizador() bool`

HasUtilizador returns a boolean if a field has been set.

### GetIntervencaos

`func (o *Tarefa) GetIntervencaos() []Intervencao`

GetIntervencaos returns the Intervencaos field if non-nil, zero value otherwise.

### GetIntervencaosOk

`func (o *Tarefa) GetIntervencaosOk() (*[]Intervencao, bool)`

GetIntervencaosOk returns a tuple with the Intervencaos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervencaos

`func (o *Tarefa) SetIntervencaos(v []Intervencao)`

SetIntervencaos sets Intervencaos field to given value.

### HasIntervencaos

`func (o *Tarefa) HasIntervencaos() bool`

HasIntervencaos returns a boolean if a field has been set.

### SetIntervencaosNil

`func (o *Tarefa) SetIntervencaosNil(b bool)`

 SetIntervencaosNil sets the value for Intervencaos to be an explicit nil

### UnsetIntervencaos
`func (o *Tarefa) UnsetIntervencaos()`

UnsetIntervencaos ensures that no value is present for Intervencaos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


