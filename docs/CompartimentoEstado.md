# CompartimentoEstado

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Estado** | Pointer to [**EstadosCompartimentoEstado**](EstadosCompartimentoEstado.md) |  | [optional] 
**Data** | Pointer to **NullableTime** |  | [optional] 
**Ocupado** | Pointer to **bool** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Vasilhame** | Pointer to **bool** |  | [optional] 
**Compartimento** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 

## Methods

### NewCompartimentoEstado

`func NewCompartimentoEstado() *CompartimentoEstado`

NewCompartimentoEstado instantiates a new CompartimentoEstado object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartimentoEstadoWithDefaults

`func NewCompartimentoEstadoWithDefaults() *CompartimentoEstado`

NewCompartimentoEstadoWithDefaults instantiates a new CompartimentoEstado object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompartimentoEstado) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompartimentoEstado) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompartimentoEstado) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CompartimentoEstado) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *CompartimentoEstado) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CompartimentoEstado) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CompartimentoEstado) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CompartimentoEstado) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *CompartimentoEstado) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *CompartimentoEstado) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *CompartimentoEstado) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *CompartimentoEstado) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *CompartimentoEstado) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *CompartimentoEstado) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *CompartimentoEstado) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *CompartimentoEstado) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *CompartimentoEstado) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *CompartimentoEstado) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *CompartimentoEstado) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *CompartimentoEstado) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefProduto

`func (o *CompartimentoEstado) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *CompartimentoEstado) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *CompartimentoEstado) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *CompartimentoEstado) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *CompartimentoEstado) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *CompartimentoEstado) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetEstado

`func (o *CompartimentoEstado) GetEstado() EstadosCompartimentoEstado`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *CompartimentoEstado) GetEstadoOk() (*EstadosCompartimentoEstado, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *CompartimentoEstado) SetEstado(v EstadosCompartimentoEstado)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *CompartimentoEstado) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetData

`func (o *CompartimentoEstado) GetData() time.Time`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompartimentoEstado) GetDataOk() (*time.Time, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompartimentoEstado) SetData(v time.Time)`

SetData sets Data field to given value.

### HasData

`func (o *CompartimentoEstado) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CompartimentoEstado) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CompartimentoEstado) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetOcupado

`func (o *CompartimentoEstado) GetOcupado() bool`

GetOcupado returns the Ocupado field if non-nil, zero value otherwise.

### GetOcupadoOk

`func (o *CompartimentoEstado) GetOcupadoOk() (*bool, bool)`

GetOcupadoOk returns a tuple with the Ocupado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcupado

`func (o *CompartimentoEstado) SetOcupado(v bool)`

SetOcupado sets Ocupado field to given value.

### HasOcupado

`func (o *CompartimentoEstado) HasOcupado() bool`

HasOcupado returns a boolean if a field has been set.

### GetRefCompartimento

`func (o *CompartimentoEstado) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *CompartimentoEstado) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *CompartimentoEstado) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *CompartimentoEstado) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *CompartimentoEstado) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *CompartimentoEstado) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetVasilhame

`func (o *CompartimentoEstado) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *CompartimentoEstado) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *CompartimentoEstado) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *CompartimentoEstado) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### GetCompartimento

`func (o *CompartimentoEstado) GetCompartimento() Compartimento`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *CompartimentoEstado) GetCompartimentoOk() (*Compartimento, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *CompartimentoEstado) SetCompartimento(v Compartimento)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *CompartimentoEstado) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


