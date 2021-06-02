# Permissao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**Chave** | Pointer to **NullableString** |  | [optional] 
**IsGroup** | Pointer to **bool** |  | [optional] 
**RefPermissao** | Pointer to **NullableInt64** |  | [optional] 
**PerfilSegurancaPermissaos** | Pointer to [**[]PerfilSegurancaPermissao**](PerfilSegurancaPermissao.md) |  | [optional] 

## Methods

### NewPermissao

`func NewPermissao() *Permissao`

NewPermissao instantiates a new Permissao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissaoWithDefaults

`func NewPermissaoWithDefaults() *Permissao`

NewPermissaoWithDefaults instantiates a new Permissao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Permissao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permissao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permissao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Permissao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Permissao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Permissao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Permissao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Permissao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Permissao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Permissao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Permissao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Permissao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Permissao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Permissao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Permissao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Permissao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Permissao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Permissao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Permissao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Permissao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Permissao) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Permissao) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Permissao) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Permissao) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Permissao) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Permissao) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetDescricao

`func (o *Permissao) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *Permissao) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *Permissao) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *Permissao) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *Permissao) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *Permissao) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetChave

`func (o *Permissao) GetChave() string`

GetChave returns the Chave field if non-nil, zero value otherwise.

### GetChaveOk

`func (o *Permissao) GetChaveOk() (*string, bool)`

GetChaveOk returns a tuple with the Chave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChave

`func (o *Permissao) SetChave(v string)`

SetChave sets Chave field to given value.

### HasChave

`func (o *Permissao) HasChave() bool`

HasChave returns a boolean if a field has been set.

### SetChaveNil

`func (o *Permissao) SetChaveNil(b bool)`

 SetChaveNil sets the value for Chave to be an explicit nil

### UnsetChave
`func (o *Permissao) UnsetChave()`

UnsetChave ensures that no value is present for Chave, not even an explicit nil
### GetIsGroup

`func (o *Permissao) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *Permissao) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *Permissao) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *Permissao) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetRefPermissao

`func (o *Permissao) GetRefPermissao() int64`

GetRefPermissao returns the RefPermissao field if non-nil, zero value otherwise.

### GetRefPermissaoOk

`func (o *Permissao) GetRefPermissaoOk() (*int64, bool)`

GetRefPermissaoOk returns a tuple with the RefPermissao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPermissao

`func (o *Permissao) SetRefPermissao(v int64)`

SetRefPermissao sets RefPermissao field to given value.

### HasRefPermissao

`func (o *Permissao) HasRefPermissao() bool`

HasRefPermissao returns a boolean if a field has been set.

### SetRefPermissaoNil

`func (o *Permissao) SetRefPermissaoNil(b bool)`

 SetRefPermissaoNil sets the value for RefPermissao to be an explicit nil

### UnsetRefPermissao
`func (o *Permissao) UnsetRefPermissao()`

UnsetRefPermissao ensures that no value is present for RefPermissao, not even an explicit nil
### GetPerfilSegurancaPermissaos

`func (o *Permissao) GetPerfilSegurancaPermissaos() []PerfilSegurancaPermissao`

GetPerfilSegurancaPermissaos returns the PerfilSegurancaPermissaos field if non-nil, zero value otherwise.

### GetPerfilSegurancaPermissaosOk

`func (o *Permissao) GetPerfilSegurancaPermissaosOk() (*[]PerfilSegurancaPermissao, bool)`

GetPerfilSegurancaPermissaosOk returns a tuple with the PerfilSegurancaPermissaos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfilSegurancaPermissaos

`func (o *Permissao) SetPerfilSegurancaPermissaos(v []PerfilSegurancaPermissao)`

SetPerfilSegurancaPermissaos sets PerfilSegurancaPermissaos field to given value.

### HasPerfilSegurancaPermissaos

`func (o *Permissao) HasPerfilSegurancaPermissaos() bool`

HasPerfilSegurancaPermissaos returns a boolean if a field has been set.

### SetPerfilSegurancaPermissaosNil

`func (o *Permissao) SetPerfilSegurancaPermissaosNil(b bool)`

 SetPerfilSegurancaPermissaosNil sets the value for PerfilSegurancaPermissaos to be an explicit nil

### UnsetPerfilSegurancaPermissaos
`func (o *Permissao) UnsetPerfilSegurancaPermissaos()`

UnsetPerfilSegurancaPermissaos ensures that no value is present for PerfilSegurancaPermissaos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


