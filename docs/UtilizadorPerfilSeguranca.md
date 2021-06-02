# UtilizadorPerfilSeguranca

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**RefPerfilSeguranca** | Pointer to **NullableInt64** |  | [optional] 
**RefPerfilSegurancaNavigation** | Pointer to [**PerfilSeguranca**](PerfilSeguranca.md) |  | [optional] 
**RefUtilizadorNavigation** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 

## Methods

### NewUtilizadorPerfilSeguranca

`func NewUtilizadorPerfilSeguranca() *UtilizadorPerfilSeguranca`

NewUtilizadorPerfilSeguranca instantiates a new UtilizadorPerfilSeguranca object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizadorPerfilSegurancaWithDefaults

`func NewUtilizadorPerfilSegurancaWithDefaults() *UtilizadorPerfilSeguranca`

NewUtilizadorPerfilSegurancaWithDefaults instantiates a new UtilizadorPerfilSeguranca object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UtilizadorPerfilSeguranca) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UtilizadorPerfilSeguranca) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UtilizadorPerfilSeguranca) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UtilizadorPerfilSeguranca) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *UtilizadorPerfilSeguranca) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *UtilizadorPerfilSeguranca) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *UtilizadorPerfilSeguranca) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *UtilizadorPerfilSeguranca) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *UtilizadorPerfilSeguranca) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *UtilizadorPerfilSeguranca) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *UtilizadorPerfilSeguranca) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *UtilizadorPerfilSeguranca) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *UtilizadorPerfilSeguranca) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *UtilizadorPerfilSeguranca) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *UtilizadorPerfilSeguranca) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *UtilizadorPerfilSeguranca) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *UtilizadorPerfilSeguranca) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *UtilizadorPerfilSeguranca) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *UtilizadorPerfilSeguranca) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *UtilizadorPerfilSeguranca) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefUtilizador

`func (o *UtilizadorPerfilSeguranca) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *UtilizadorPerfilSeguranca) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *UtilizadorPerfilSeguranca) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *UtilizadorPerfilSeguranca) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *UtilizadorPerfilSeguranca) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *UtilizadorPerfilSeguranca) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetRefPerfilSeguranca

`func (o *UtilizadorPerfilSeguranca) GetRefPerfilSeguranca() int64`

GetRefPerfilSeguranca returns the RefPerfilSeguranca field if non-nil, zero value otherwise.

### GetRefPerfilSegurancaOk

`func (o *UtilizadorPerfilSeguranca) GetRefPerfilSegurancaOk() (*int64, bool)`

GetRefPerfilSegurancaOk returns a tuple with the RefPerfilSeguranca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPerfilSeguranca

`func (o *UtilizadorPerfilSeguranca) SetRefPerfilSeguranca(v int64)`

SetRefPerfilSeguranca sets RefPerfilSeguranca field to given value.

### HasRefPerfilSeguranca

`func (o *UtilizadorPerfilSeguranca) HasRefPerfilSeguranca() bool`

HasRefPerfilSeguranca returns a boolean if a field has been set.

### SetRefPerfilSegurancaNil

`func (o *UtilizadorPerfilSeguranca) SetRefPerfilSegurancaNil(b bool)`

 SetRefPerfilSegurancaNil sets the value for RefPerfilSeguranca to be an explicit nil

### UnsetRefPerfilSeguranca
`func (o *UtilizadorPerfilSeguranca) UnsetRefPerfilSeguranca()`

UnsetRefPerfilSeguranca ensures that no value is present for RefPerfilSeguranca, not even an explicit nil
### GetRefPerfilSegurancaNavigation

`func (o *UtilizadorPerfilSeguranca) GetRefPerfilSegurancaNavigation() PerfilSeguranca`

GetRefPerfilSegurancaNavigation returns the RefPerfilSegurancaNavigation field if non-nil, zero value otherwise.

### GetRefPerfilSegurancaNavigationOk

`func (o *UtilizadorPerfilSeguranca) GetRefPerfilSegurancaNavigationOk() (*PerfilSeguranca, bool)`

GetRefPerfilSegurancaNavigationOk returns a tuple with the RefPerfilSegurancaNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPerfilSegurancaNavigation

`func (o *UtilizadorPerfilSeguranca) SetRefPerfilSegurancaNavigation(v PerfilSeguranca)`

SetRefPerfilSegurancaNavigation sets RefPerfilSegurancaNavigation field to given value.

### HasRefPerfilSegurancaNavigation

`func (o *UtilizadorPerfilSeguranca) HasRefPerfilSegurancaNavigation() bool`

HasRefPerfilSegurancaNavigation returns a boolean if a field has been set.

### GetRefUtilizadorNavigation

`func (o *UtilizadorPerfilSeguranca) GetRefUtilizadorNavigation() Utilizador`

GetRefUtilizadorNavigation returns the RefUtilizadorNavigation field if non-nil, zero value otherwise.

### GetRefUtilizadorNavigationOk

`func (o *UtilizadorPerfilSeguranca) GetRefUtilizadorNavigationOk() (*Utilizador, bool)`

GetRefUtilizadorNavigationOk returns a tuple with the RefUtilizadorNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizadorNavigation

`func (o *UtilizadorPerfilSeguranca) SetRefUtilizadorNavigation(v Utilizador)`

SetRefUtilizadorNavigation sets RefUtilizadorNavigation field to given value.

### HasRefUtilizadorNavigation

`func (o *UtilizadorPerfilSeguranca) HasRefUtilizadorNavigation() bool`

HasRefUtilizadorNavigation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


