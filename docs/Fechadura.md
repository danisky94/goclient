# Fechadura

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Endereco** | Pointer to **NullableString** |  | [optional] 
**Canal** | Pointer to **NullableInt32** |  | [optional] 
**Compartimento** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 

## Methods

### NewFechadura

`func NewFechadura() *Fechadura`

NewFechadura instantiates a new Fechadura object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFechaduraWithDefaults

`func NewFechaduraWithDefaults() *Fechadura`

NewFechaduraWithDefaults instantiates a new Fechadura object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fechadura) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fechadura) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fechadura) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Fechadura) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Fechadura) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Fechadura) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Fechadura) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Fechadura) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Fechadura) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Fechadura) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Fechadura) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Fechadura) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Fechadura) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Fechadura) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Fechadura) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Fechadura) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Fechadura) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Fechadura) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Fechadura) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Fechadura) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefCompartimento

`func (o *Fechadura) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *Fechadura) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *Fechadura) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *Fechadura) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *Fechadura) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *Fechadura) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetEndereco

`func (o *Fechadura) GetEndereco() string`

GetEndereco returns the Endereco field if non-nil, zero value otherwise.

### GetEnderecoOk

`func (o *Fechadura) GetEnderecoOk() (*string, bool)`

GetEnderecoOk returns a tuple with the Endereco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndereco

`func (o *Fechadura) SetEndereco(v string)`

SetEndereco sets Endereco field to given value.

### HasEndereco

`func (o *Fechadura) HasEndereco() bool`

HasEndereco returns a boolean if a field has been set.

### SetEnderecoNil

`func (o *Fechadura) SetEnderecoNil(b bool)`

 SetEnderecoNil sets the value for Endereco to be an explicit nil

### UnsetEndereco
`func (o *Fechadura) UnsetEndereco()`

UnsetEndereco ensures that no value is present for Endereco, not even an explicit nil
### GetCanal

`func (o *Fechadura) GetCanal() int32`

GetCanal returns the Canal field if non-nil, zero value otherwise.

### GetCanalOk

`func (o *Fechadura) GetCanalOk() (*int32, bool)`

GetCanalOk returns a tuple with the Canal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanal

`func (o *Fechadura) SetCanal(v int32)`

SetCanal sets Canal field to given value.

### HasCanal

`func (o *Fechadura) HasCanal() bool`

HasCanal returns a boolean if a field has been set.

### SetCanalNil

`func (o *Fechadura) SetCanalNil(b bool)`

 SetCanalNil sets the value for Canal to be an explicit nil

### UnsetCanal
`func (o *Fechadura) UnsetCanal()`

UnsetCanal ensures that no value is present for Canal, not even an explicit nil
### GetCompartimento

`func (o *Fechadura) GetCompartimento() Compartimento`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *Fechadura) GetCompartimentoOk() (*Compartimento, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *Fechadura) SetCompartimento(v Compartimento)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *Fechadura) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


