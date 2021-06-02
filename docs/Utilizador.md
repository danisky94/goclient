# Utilizador

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Telemovel** | Pointer to **NullableString** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoUtilizador**](TipoUtilizador.md) |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Salt** | Pointer to **NullableString** |  | [optional] 
**Estado** | Pointer to [**UtilizadorEstados**](UtilizadorEstados.md) |  | [optional] 
**RefEntidadeNavigation** | Pointer to [**Entidade**](Entidade.md) |  | [optional] 
**Intervencaos** | Pointer to [**[]Intervencao**](Intervencao.md) |  | [optional] 
**NoticacaoUtilizadors** | Pointer to [**[]NotificacaoUtilizador**](NotificacaoUtilizador.md) |  | [optional] 
**Tarefas** | Pointer to [**[]Tarefa**](Tarefa.md) |  | [optional] 
**TipoNotificacaoUtilizadors** | Pointer to [**[]TipoNotificacaoUtilizador**](TipoNotificacaoUtilizador.md) |  | [optional] 
**UtilizadorPerfilSegurancas** | Pointer to [**[]UtilizadorPerfilSeguranca**](UtilizadorPerfilSeguranca.md) |  | [optional] 
**OrdensServico** | Pointer to [**[]OrdemServico**](OrdemServico.md) |  | [optional] 

## Methods

### NewUtilizador

`func NewUtilizador() *Utilizador`

NewUtilizador instantiates a new Utilizador object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizadorWithDefaults

`func NewUtilizadorWithDefaults() *Utilizador`

NewUtilizadorWithDefaults instantiates a new Utilizador object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Utilizador) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Utilizador) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Utilizador) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Utilizador) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Utilizador) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Utilizador) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Utilizador) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Utilizador) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Utilizador) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Utilizador) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Utilizador) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Utilizador) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Utilizador) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Utilizador) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Utilizador) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Utilizador) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Utilizador) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Utilizador) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Utilizador) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Utilizador) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Utilizador) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Utilizador) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Utilizador) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Utilizador) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Utilizador) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Utilizador) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetEmail

`func (o *Utilizador) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Utilizador) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Utilizador) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Utilizador) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Utilizador) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Utilizador) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTelemovel

`func (o *Utilizador) GetTelemovel() string`

GetTelemovel returns the Telemovel field if non-nil, zero value otherwise.

### GetTelemovelOk

`func (o *Utilizador) GetTelemovelOk() (*string, bool)`

GetTelemovelOk returns a tuple with the Telemovel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemovel

`func (o *Utilizador) SetTelemovel(v string)`

SetTelemovel sets Telemovel field to given value.

### HasTelemovel

`func (o *Utilizador) HasTelemovel() bool`

HasTelemovel returns a boolean if a field has been set.

### SetTelemovelNil

`func (o *Utilizador) SetTelemovelNil(b bool)`

 SetTelemovelNil sets the value for Telemovel to be an explicit nil

### UnsetTelemovel
`func (o *Utilizador) UnsetTelemovel()`

UnsetTelemovel ensures that no value is present for Telemovel, not even an explicit nil
### GetRefEntidade

`func (o *Utilizador) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *Utilizador) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *Utilizador) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *Utilizador) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *Utilizador) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *Utilizador) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetTipo

`func (o *Utilizador) GetTipo() TipoUtilizador`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Utilizador) GetTipoOk() (*TipoUtilizador, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Utilizador) SetTipo(v TipoUtilizador)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Utilizador) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetPassword

`func (o *Utilizador) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Utilizador) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Utilizador) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Utilizador) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Utilizador) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Utilizador) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSalt

`func (o *Utilizador) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *Utilizador) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *Utilizador) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *Utilizador) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### SetSaltNil

`func (o *Utilizador) SetSaltNil(b bool)`

 SetSaltNil sets the value for Salt to be an explicit nil

### UnsetSalt
`func (o *Utilizador) UnsetSalt()`

UnsetSalt ensures that no value is present for Salt, not even an explicit nil
### GetEstado

`func (o *Utilizador) GetEstado() UtilizadorEstados`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Utilizador) GetEstadoOk() (*UtilizadorEstados, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Utilizador) SetEstado(v UtilizadorEstados)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Utilizador) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetRefEntidadeNavigation

`func (o *Utilizador) GetRefEntidadeNavigation() Entidade`

GetRefEntidadeNavigation returns the RefEntidadeNavigation field if non-nil, zero value otherwise.

### GetRefEntidadeNavigationOk

`func (o *Utilizador) GetRefEntidadeNavigationOk() (*Entidade, bool)`

GetRefEntidadeNavigationOk returns a tuple with the RefEntidadeNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidadeNavigation

`func (o *Utilizador) SetRefEntidadeNavigation(v Entidade)`

SetRefEntidadeNavigation sets RefEntidadeNavigation field to given value.

### HasRefEntidadeNavigation

`func (o *Utilizador) HasRefEntidadeNavigation() bool`

HasRefEntidadeNavigation returns a boolean if a field has been set.

### GetIntervencaos

`func (o *Utilizador) GetIntervencaos() []Intervencao`

GetIntervencaos returns the Intervencaos field if non-nil, zero value otherwise.

### GetIntervencaosOk

`func (o *Utilizador) GetIntervencaosOk() (*[]Intervencao, bool)`

GetIntervencaosOk returns a tuple with the Intervencaos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervencaos

`func (o *Utilizador) SetIntervencaos(v []Intervencao)`

SetIntervencaos sets Intervencaos field to given value.

### HasIntervencaos

`func (o *Utilizador) HasIntervencaos() bool`

HasIntervencaos returns a boolean if a field has been set.

### SetIntervencaosNil

`func (o *Utilizador) SetIntervencaosNil(b bool)`

 SetIntervencaosNil sets the value for Intervencaos to be an explicit nil

### UnsetIntervencaos
`func (o *Utilizador) UnsetIntervencaos()`

UnsetIntervencaos ensures that no value is present for Intervencaos, not even an explicit nil
### GetNoticacaoUtilizadors

`func (o *Utilizador) GetNoticacaoUtilizadors() []NotificacaoUtilizador`

GetNoticacaoUtilizadors returns the NoticacaoUtilizadors field if non-nil, zero value otherwise.

### GetNoticacaoUtilizadorsOk

`func (o *Utilizador) GetNoticacaoUtilizadorsOk() (*[]NotificacaoUtilizador, bool)`

GetNoticacaoUtilizadorsOk returns a tuple with the NoticacaoUtilizadors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticacaoUtilizadors

`func (o *Utilizador) SetNoticacaoUtilizadors(v []NotificacaoUtilizador)`

SetNoticacaoUtilizadors sets NoticacaoUtilizadors field to given value.

### HasNoticacaoUtilizadors

`func (o *Utilizador) HasNoticacaoUtilizadors() bool`

HasNoticacaoUtilizadors returns a boolean if a field has been set.

### SetNoticacaoUtilizadorsNil

`func (o *Utilizador) SetNoticacaoUtilizadorsNil(b bool)`

 SetNoticacaoUtilizadorsNil sets the value for NoticacaoUtilizadors to be an explicit nil

### UnsetNoticacaoUtilizadors
`func (o *Utilizador) UnsetNoticacaoUtilizadors()`

UnsetNoticacaoUtilizadors ensures that no value is present for NoticacaoUtilizadors, not even an explicit nil
### GetTarefas

`func (o *Utilizador) GetTarefas() []Tarefa`

GetTarefas returns the Tarefas field if non-nil, zero value otherwise.

### GetTarefasOk

`func (o *Utilizador) GetTarefasOk() (*[]Tarefa, bool)`

GetTarefasOk returns a tuple with the Tarefas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarefas

`func (o *Utilizador) SetTarefas(v []Tarefa)`

SetTarefas sets Tarefas field to given value.

### HasTarefas

`func (o *Utilizador) HasTarefas() bool`

HasTarefas returns a boolean if a field has been set.

### SetTarefasNil

`func (o *Utilizador) SetTarefasNil(b bool)`

 SetTarefasNil sets the value for Tarefas to be an explicit nil

### UnsetTarefas
`func (o *Utilizador) UnsetTarefas()`

UnsetTarefas ensures that no value is present for Tarefas, not even an explicit nil
### GetTipoNotificacaoUtilizadors

`func (o *Utilizador) GetTipoNotificacaoUtilizadors() []TipoNotificacaoUtilizador`

GetTipoNotificacaoUtilizadors returns the TipoNotificacaoUtilizadors field if non-nil, zero value otherwise.

### GetTipoNotificacaoUtilizadorsOk

`func (o *Utilizador) GetTipoNotificacaoUtilizadorsOk() (*[]TipoNotificacaoUtilizador, bool)`

GetTipoNotificacaoUtilizadorsOk returns a tuple with the TipoNotificacaoUtilizadors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoNotificacaoUtilizadors

`func (o *Utilizador) SetTipoNotificacaoUtilizadors(v []TipoNotificacaoUtilizador)`

SetTipoNotificacaoUtilizadors sets TipoNotificacaoUtilizadors field to given value.

### HasTipoNotificacaoUtilizadors

`func (o *Utilizador) HasTipoNotificacaoUtilizadors() bool`

HasTipoNotificacaoUtilizadors returns a boolean if a field has been set.

### SetTipoNotificacaoUtilizadorsNil

`func (o *Utilizador) SetTipoNotificacaoUtilizadorsNil(b bool)`

 SetTipoNotificacaoUtilizadorsNil sets the value for TipoNotificacaoUtilizadors to be an explicit nil

### UnsetTipoNotificacaoUtilizadors
`func (o *Utilizador) UnsetTipoNotificacaoUtilizadors()`

UnsetTipoNotificacaoUtilizadors ensures that no value is present for TipoNotificacaoUtilizadors, not even an explicit nil
### GetUtilizadorPerfilSegurancas

`func (o *Utilizador) GetUtilizadorPerfilSegurancas() []UtilizadorPerfilSeguranca`

GetUtilizadorPerfilSegurancas returns the UtilizadorPerfilSegurancas field if non-nil, zero value otherwise.

### GetUtilizadorPerfilSegurancasOk

`func (o *Utilizador) GetUtilizadorPerfilSegurancasOk() (*[]UtilizadorPerfilSeguranca, bool)`

GetUtilizadorPerfilSegurancasOk returns a tuple with the UtilizadorPerfilSegurancas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizadorPerfilSegurancas

`func (o *Utilizador) SetUtilizadorPerfilSegurancas(v []UtilizadorPerfilSeguranca)`

SetUtilizadorPerfilSegurancas sets UtilizadorPerfilSegurancas field to given value.

### HasUtilizadorPerfilSegurancas

`func (o *Utilizador) HasUtilizadorPerfilSegurancas() bool`

HasUtilizadorPerfilSegurancas returns a boolean if a field has been set.

### SetUtilizadorPerfilSegurancasNil

`func (o *Utilizador) SetUtilizadorPerfilSegurancasNil(b bool)`

 SetUtilizadorPerfilSegurancasNil sets the value for UtilizadorPerfilSegurancas to be an explicit nil

### UnsetUtilizadorPerfilSegurancas
`func (o *Utilizador) UnsetUtilizadorPerfilSegurancas()`

UnsetUtilizadorPerfilSegurancas ensures that no value is present for UtilizadorPerfilSegurancas, not even an explicit nil
### GetOrdensServico

`func (o *Utilizador) GetOrdensServico() []OrdemServico`

GetOrdensServico returns the OrdensServico field if non-nil, zero value otherwise.

### GetOrdensServicoOk

`func (o *Utilizador) GetOrdensServicoOk() (*[]OrdemServico, bool)`

GetOrdensServicoOk returns a tuple with the OrdensServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdensServico

`func (o *Utilizador) SetOrdensServico(v []OrdemServico)`

SetOrdensServico sets OrdensServico field to given value.

### HasOrdensServico

`func (o *Utilizador) HasOrdensServico() bool`

HasOrdensServico returns a boolean if a field has been set.

### SetOrdensServicoNil

`func (o *Utilizador) SetOrdensServicoNil(b bool)`

 SetOrdensServicoNil sets the value for OrdensServico to be an explicit nil

### UnsetOrdensServico
`func (o *Utilizador) UnsetOrdensServico()`

UnsetOrdensServico ensures that no value is present for OrdensServico, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


