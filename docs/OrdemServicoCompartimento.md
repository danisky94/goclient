# OrdemServicoCompartimento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefOrdemServico** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoOrdemServicoCompartimento**](TipoOrdemServicoCompartimento.md) |  | [optional] 
**Terminado** | Pointer to **bool** |  | [optional] 
**OrdemServico** | Pointer to [**OrdemServico**](OrdemServico.md) |  | [optional] 
**Produto** | Pointer to [**Produto**](Produto.md) |  | [optional] 
**Compartimento** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 

## Methods

### NewOrdemServicoCompartimento

`func NewOrdemServicoCompartimento() *OrdemServicoCompartimento`

NewOrdemServicoCompartimento instantiates a new OrdemServicoCompartimento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdemServicoCompartimentoWithDefaults

`func NewOrdemServicoCompartimentoWithDefaults() *OrdemServicoCompartimento`

NewOrdemServicoCompartimentoWithDefaults instantiates a new OrdemServicoCompartimento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrdemServicoCompartimento) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdemServicoCompartimento) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdemServicoCompartimento) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrdemServicoCompartimento) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *OrdemServicoCompartimento) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrdemServicoCompartimento) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrdemServicoCompartimento) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OrdemServicoCompartimento) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *OrdemServicoCompartimento) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *OrdemServicoCompartimento) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *OrdemServicoCompartimento) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *OrdemServicoCompartimento) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *OrdemServicoCompartimento) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *OrdemServicoCompartimento) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *OrdemServicoCompartimento) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *OrdemServicoCompartimento) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *OrdemServicoCompartimento) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *OrdemServicoCompartimento) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *OrdemServicoCompartimento) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *OrdemServicoCompartimento) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefOrdemServico

`func (o *OrdemServicoCompartimento) GetRefOrdemServico() int64`

GetRefOrdemServico returns the RefOrdemServico field if non-nil, zero value otherwise.

### GetRefOrdemServicoOk

`func (o *OrdemServicoCompartimento) GetRefOrdemServicoOk() (*int64, bool)`

GetRefOrdemServicoOk returns a tuple with the RefOrdemServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefOrdemServico

`func (o *OrdemServicoCompartimento) SetRefOrdemServico(v int64)`

SetRefOrdemServico sets RefOrdemServico field to given value.

### HasRefOrdemServico

`func (o *OrdemServicoCompartimento) HasRefOrdemServico() bool`

HasRefOrdemServico returns a boolean if a field has been set.

### SetRefOrdemServicoNil

`func (o *OrdemServicoCompartimento) SetRefOrdemServicoNil(b bool)`

 SetRefOrdemServicoNil sets the value for RefOrdemServico to be an explicit nil

### UnsetRefOrdemServico
`func (o *OrdemServicoCompartimento) UnsetRefOrdemServico()`

UnsetRefOrdemServico ensures that no value is present for RefOrdemServico, not even an explicit nil
### GetRefProduto

`func (o *OrdemServicoCompartimento) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *OrdemServicoCompartimento) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *OrdemServicoCompartimento) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *OrdemServicoCompartimento) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *OrdemServicoCompartimento) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *OrdemServicoCompartimento) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetRefCompartimento

`func (o *OrdemServicoCompartimento) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *OrdemServicoCompartimento) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *OrdemServicoCompartimento) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *OrdemServicoCompartimento) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *OrdemServicoCompartimento) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *OrdemServicoCompartimento) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetTipo

`func (o *OrdemServicoCompartimento) GetTipo() TipoOrdemServicoCompartimento`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *OrdemServicoCompartimento) GetTipoOk() (*TipoOrdemServicoCompartimento, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *OrdemServicoCompartimento) SetTipo(v TipoOrdemServicoCompartimento)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *OrdemServicoCompartimento) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetTerminado

`func (o *OrdemServicoCompartimento) GetTerminado() bool`

GetTerminado returns the Terminado field if non-nil, zero value otherwise.

### GetTerminadoOk

`func (o *OrdemServicoCompartimento) GetTerminadoOk() (*bool, bool)`

GetTerminadoOk returns a tuple with the Terminado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminado

`func (o *OrdemServicoCompartimento) SetTerminado(v bool)`

SetTerminado sets Terminado field to given value.

### HasTerminado

`func (o *OrdemServicoCompartimento) HasTerminado() bool`

HasTerminado returns a boolean if a field has been set.

### GetOrdemServico

`func (o *OrdemServicoCompartimento) GetOrdemServico() OrdemServico`

GetOrdemServico returns the OrdemServico field if non-nil, zero value otherwise.

### GetOrdemServicoOk

`func (o *OrdemServicoCompartimento) GetOrdemServicoOk() (*OrdemServico, bool)`

GetOrdemServicoOk returns a tuple with the OrdemServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServico

`func (o *OrdemServicoCompartimento) SetOrdemServico(v OrdemServico)`

SetOrdemServico sets OrdemServico field to given value.

### HasOrdemServico

`func (o *OrdemServicoCompartimento) HasOrdemServico() bool`

HasOrdemServico returns a boolean if a field has been set.

### GetProduto

`func (o *OrdemServicoCompartimento) GetProduto() Produto`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *OrdemServicoCompartimento) GetProdutoOk() (*Produto, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *OrdemServicoCompartimento) SetProduto(v Produto)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *OrdemServicoCompartimento) HasProduto() bool`

HasProduto returns a boolean if a field has been set.

### GetCompartimento

`func (o *OrdemServicoCompartimento) GetCompartimento() Compartimento`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *OrdemServicoCompartimento) GetCompartimentoOk() (*Compartimento, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *OrdemServicoCompartimento) SetCompartimento(v Compartimento)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *OrdemServicoCompartimento) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


