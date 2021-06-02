# Plano

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**TaxaUtilizacao** | Pointer to **NullableInt32** |  | [optional] 
**Periodicidade** | Pointer to **NullableInt32** |  | [optional] 
**Entidades** | Pointer to [**[]Entidade**](Entidade.md) |  | [optional] 

## Methods

### NewPlano

`func NewPlano() *Plano`

NewPlano instantiates a new Plano object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanoWithDefaults

`func NewPlanoWithDefaults() *Plano`

NewPlanoWithDefaults instantiates a new Plano object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plano) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plano) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plano) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Plano) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Plano) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Plano) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Plano) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Plano) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Plano) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Plano) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Plano) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Plano) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Plano) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Plano) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Plano) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Plano) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Plano) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Plano) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Plano) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Plano) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Plano) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Plano) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Plano) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Plano) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Plano) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Plano) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetTaxaUtilizacao

`func (o *Plano) GetTaxaUtilizacao() int32`

GetTaxaUtilizacao returns the TaxaUtilizacao field if non-nil, zero value otherwise.

### GetTaxaUtilizacaoOk

`func (o *Plano) GetTaxaUtilizacaoOk() (*int32, bool)`

GetTaxaUtilizacaoOk returns a tuple with the TaxaUtilizacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxaUtilizacao

`func (o *Plano) SetTaxaUtilizacao(v int32)`

SetTaxaUtilizacao sets TaxaUtilizacao field to given value.

### HasTaxaUtilizacao

`func (o *Plano) HasTaxaUtilizacao() bool`

HasTaxaUtilizacao returns a boolean if a field has been set.

### SetTaxaUtilizacaoNil

`func (o *Plano) SetTaxaUtilizacaoNil(b bool)`

 SetTaxaUtilizacaoNil sets the value for TaxaUtilizacao to be an explicit nil

### UnsetTaxaUtilizacao
`func (o *Plano) UnsetTaxaUtilizacao()`

UnsetTaxaUtilizacao ensures that no value is present for TaxaUtilizacao, not even an explicit nil
### GetPeriodicidade

`func (o *Plano) GetPeriodicidade() int32`

GetPeriodicidade returns the Periodicidade field if non-nil, zero value otherwise.

### GetPeriodicidadeOk

`func (o *Plano) GetPeriodicidadeOk() (*int32, bool)`

GetPeriodicidadeOk returns a tuple with the Periodicidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicidade

`func (o *Plano) SetPeriodicidade(v int32)`

SetPeriodicidade sets Periodicidade field to given value.

### HasPeriodicidade

`func (o *Plano) HasPeriodicidade() bool`

HasPeriodicidade returns a boolean if a field has been set.

### SetPeriodicidadeNil

`func (o *Plano) SetPeriodicidadeNil(b bool)`

 SetPeriodicidadeNil sets the value for Periodicidade to be an explicit nil

### UnsetPeriodicidade
`func (o *Plano) UnsetPeriodicidade()`

UnsetPeriodicidade ensures that no value is present for Periodicidade, not even an explicit nil
### GetEntidades

`func (o *Plano) GetEntidades() []Entidade`

GetEntidades returns the Entidades field if non-nil, zero value otherwise.

### GetEntidadesOk

`func (o *Plano) GetEntidadesOk() (*[]Entidade, bool)`

GetEntidadesOk returns a tuple with the Entidades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntidades

`func (o *Plano) SetEntidades(v []Entidade)`

SetEntidades sets Entidades field to given value.

### HasEntidades

`func (o *Plano) HasEntidades() bool`

HasEntidades returns a boolean if a field has been set.

### SetEntidadesNil

`func (o *Plano) SetEntidadesNil(b bool)`

 SetEntidadesNil sets the value for Entidades to be an explicit nil

### UnsetEntidades
`func (o *Plano) UnsetEntidades()`

UnsetEntidades ensures that no value is present for Entidades, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


