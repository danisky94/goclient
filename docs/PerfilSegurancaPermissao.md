# PerfilSegurancaPermissao

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefPerfilSeguranca** | Pointer to **NullableInt64** |  | [optional] 
**RefPermissao** | Pointer to **NullableInt64** |  | [optional] 
**PerfilSeguranca** | Pointer to [**PerfilSeguranca**](PerfilSeguranca.md) |  | [optional] 
**Permissao** | Pointer to [**Permissao**](Permissao.md) |  | [optional] 

## Methods

### NewPerfilSegurancaPermissao

`func NewPerfilSegurancaPermissao() *PerfilSegurancaPermissao`

NewPerfilSegurancaPermissao instantiates a new PerfilSegurancaPermissao object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfilSegurancaPermissaoWithDefaults

`func NewPerfilSegurancaPermissaoWithDefaults() *PerfilSegurancaPermissao`

NewPerfilSegurancaPermissaoWithDefaults instantiates a new PerfilSegurancaPermissao object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PerfilSegurancaPermissao) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerfilSegurancaPermissao) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerfilSegurancaPermissao) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PerfilSegurancaPermissao) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *PerfilSegurancaPermissao) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PerfilSegurancaPermissao) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PerfilSegurancaPermissao) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PerfilSegurancaPermissao) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *PerfilSegurancaPermissao) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *PerfilSegurancaPermissao) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *PerfilSegurancaPermissao) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *PerfilSegurancaPermissao) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *PerfilSegurancaPermissao) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *PerfilSegurancaPermissao) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *PerfilSegurancaPermissao) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *PerfilSegurancaPermissao) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *PerfilSegurancaPermissao) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *PerfilSegurancaPermissao) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *PerfilSegurancaPermissao) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *PerfilSegurancaPermissao) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefPerfilSeguranca

`func (o *PerfilSegurancaPermissao) GetRefPerfilSeguranca() int64`

GetRefPerfilSeguranca returns the RefPerfilSeguranca field if non-nil, zero value otherwise.

### GetRefPerfilSegurancaOk

`func (o *PerfilSegurancaPermissao) GetRefPerfilSegurancaOk() (*int64, bool)`

GetRefPerfilSegurancaOk returns a tuple with the RefPerfilSeguranca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPerfilSeguranca

`func (o *PerfilSegurancaPermissao) SetRefPerfilSeguranca(v int64)`

SetRefPerfilSeguranca sets RefPerfilSeguranca field to given value.

### HasRefPerfilSeguranca

`func (o *PerfilSegurancaPermissao) HasRefPerfilSeguranca() bool`

HasRefPerfilSeguranca returns a boolean if a field has been set.

### SetRefPerfilSegurancaNil

`func (o *PerfilSegurancaPermissao) SetRefPerfilSegurancaNil(b bool)`

 SetRefPerfilSegurancaNil sets the value for RefPerfilSeguranca to be an explicit nil

### UnsetRefPerfilSeguranca
`func (o *PerfilSegurancaPermissao) UnsetRefPerfilSeguranca()`

UnsetRefPerfilSeguranca ensures that no value is present for RefPerfilSeguranca, not even an explicit nil
### GetRefPermissao

`func (o *PerfilSegurancaPermissao) GetRefPermissao() int64`

GetRefPermissao returns the RefPermissao field if non-nil, zero value otherwise.

### GetRefPermissaoOk

`func (o *PerfilSegurancaPermissao) GetRefPermissaoOk() (*int64, bool)`

GetRefPermissaoOk returns a tuple with the RefPermissao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPermissao

`func (o *PerfilSegurancaPermissao) SetRefPermissao(v int64)`

SetRefPermissao sets RefPermissao field to given value.

### HasRefPermissao

`func (o *PerfilSegurancaPermissao) HasRefPermissao() bool`

HasRefPermissao returns a boolean if a field has been set.

### SetRefPermissaoNil

`func (o *PerfilSegurancaPermissao) SetRefPermissaoNil(b bool)`

 SetRefPermissaoNil sets the value for RefPermissao to be an explicit nil

### UnsetRefPermissao
`func (o *PerfilSegurancaPermissao) UnsetRefPermissao()`

UnsetRefPermissao ensures that no value is present for RefPermissao, not even an explicit nil
### GetPerfilSeguranca

`func (o *PerfilSegurancaPermissao) GetPerfilSeguranca() PerfilSeguranca`

GetPerfilSeguranca returns the PerfilSeguranca field if non-nil, zero value otherwise.

### GetPerfilSegurancaOk

`func (o *PerfilSegurancaPermissao) GetPerfilSegurancaOk() (*PerfilSeguranca, bool)`

GetPerfilSegurancaOk returns a tuple with the PerfilSeguranca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfilSeguranca

`func (o *PerfilSegurancaPermissao) SetPerfilSeguranca(v PerfilSeguranca)`

SetPerfilSeguranca sets PerfilSeguranca field to given value.

### HasPerfilSeguranca

`func (o *PerfilSegurancaPermissao) HasPerfilSeguranca() bool`

HasPerfilSeguranca returns a boolean if a field has been set.

### GetPermissao

`func (o *PerfilSegurancaPermissao) GetPermissao() Permissao`

GetPermissao returns the Permissao field if non-nil, zero value otherwise.

### GetPermissaoOk

`func (o *PerfilSegurancaPermissao) GetPermissaoOk() (*Permissao, bool)`

GetPermissaoOk returns a tuple with the Permissao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissao

`func (o *PerfilSegurancaPermissao) SetPermissao(v Permissao)`

SetPermissao sets Permissao field to given value.

### HasPermissao

`func (o *PerfilSegurancaPermissao) HasPermissao() bool`

HasPermissao returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


