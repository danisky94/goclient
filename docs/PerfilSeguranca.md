# PerfilSeguranca

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**PerfilSegurancaPermissaos** | Pointer to [**[]PerfilSegurancaPermissao**](PerfilSegurancaPermissao.md) |  | [optional] 
**UtilizadorPerfilSegurancas** | Pointer to [**[]UtilizadorPerfilSeguranca**](UtilizadorPerfilSeguranca.md) |  | [optional] 

## Methods

### NewPerfilSeguranca

`func NewPerfilSeguranca() *PerfilSeguranca`

NewPerfilSeguranca instantiates a new PerfilSeguranca object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfilSegurancaWithDefaults

`func NewPerfilSegurancaWithDefaults() *PerfilSeguranca`

NewPerfilSegurancaWithDefaults instantiates a new PerfilSeguranca object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PerfilSeguranca) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerfilSeguranca) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerfilSeguranca) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PerfilSeguranca) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *PerfilSeguranca) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PerfilSeguranca) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PerfilSeguranca) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PerfilSeguranca) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *PerfilSeguranca) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *PerfilSeguranca) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *PerfilSeguranca) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *PerfilSeguranca) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *PerfilSeguranca) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *PerfilSeguranca) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *PerfilSeguranca) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *PerfilSeguranca) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *PerfilSeguranca) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *PerfilSeguranca) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *PerfilSeguranca) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *PerfilSeguranca) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *PerfilSeguranca) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *PerfilSeguranca) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *PerfilSeguranca) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *PerfilSeguranca) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *PerfilSeguranca) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *PerfilSeguranca) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetPerfilSegurancaPermissaos

`func (o *PerfilSeguranca) GetPerfilSegurancaPermissaos() []PerfilSegurancaPermissao`

GetPerfilSegurancaPermissaos returns the PerfilSegurancaPermissaos field if non-nil, zero value otherwise.

### GetPerfilSegurancaPermissaosOk

`func (o *PerfilSeguranca) GetPerfilSegurancaPermissaosOk() (*[]PerfilSegurancaPermissao, bool)`

GetPerfilSegurancaPermissaosOk returns a tuple with the PerfilSegurancaPermissaos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfilSegurancaPermissaos

`func (o *PerfilSeguranca) SetPerfilSegurancaPermissaos(v []PerfilSegurancaPermissao)`

SetPerfilSegurancaPermissaos sets PerfilSegurancaPermissaos field to given value.

### HasPerfilSegurancaPermissaos

`func (o *PerfilSeguranca) HasPerfilSegurancaPermissaos() bool`

HasPerfilSegurancaPermissaos returns a boolean if a field has been set.

### SetPerfilSegurancaPermissaosNil

`func (o *PerfilSeguranca) SetPerfilSegurancaPermissaosNil(b bool)`

 SetPerfilSegurancaPermissaosNil sets the value for PerfilSegurancaPermissaos to be an explicit nil

### UnsetPerfilSegurancaPermissaos
`func (o *PerfilSeguranca) UnsetPerfilSegurancaPermissaos()`

UnsetPerfilSegurancaPermissaos ensures that no value is present for PerfilSegurancaPermissaos, not even an explicit nil
### GetUtilizadorPerfilSegurancas

`func (o *PerfilSeguranca) GetUtilizadorPerfilSegurancas() []UtilizadorPerfilSeguranca`

GetUtilizadorPerfilSegurancas returns the UtilizadorPerfilSegurancas field if non-nil, zero value otherwise.

### GetUtilizadorPerfilSegurancasOk

`func (o *PerfilSeguranca) GetUtilizadorPerfilSegurancasOk() (*[]UtilizadorPerfilSeguranca, bool)`

GetUtilizadorPerfilSegurancasOk returns a tuple with the UtilizadorPerfilSegurancas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizadorPerfilSegurancas

`func (o *PerfilSeguranca) SetUtilizadorPerfilSegurancas(v []UtilizadorPerfilSeguranca)`

SetUtilizadorPerfilSegurancas sets UtilizadorPerfilSegurancas field to given value.

### HasUtilizadorPerfilSegurancas

`func (o *PerfilSeguranca) HasUtilizadorPerfilSegurancas() bool`

HasUtilizadorPerfilSegurancas returns a boolean if a field has been set.

### SetUtilizadorPerfilSegurancasNil

`func (o *PerfilSeguranca) SetUtilizadorPerfilSegurancasNil(b bool)`

 SetUtilizadorPerfilSegurancasNil sets the value for UtilizadorPerfilSegurancas to be an explicit nil

### UnsetUtilizadorPerfilSegurancas
`func (o *PerfilSeguranca) UnsetUtilizadorPerfilSegurancas()`

UnsetUtilizadorPerfilSegurancas ensures that no value is present for UtilizadorPerfilSegurancas, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


