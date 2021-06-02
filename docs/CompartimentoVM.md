# CompartimentoVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Numero** | Pointer to **NullableString** |  | [optional] 
**Tipo** | Pointer to [**TipoCompartimento**](TipoCompartimento.md) |  | [optional] 
**Ativo** | Pointer to **NullableBool** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**LivreSemProduto** | Pointer to **bool** |  | [optional] 
**CompartimentoBloqueado** | Pointer to **bool** |  | [optional] 
**Vasilhame** | Pointer to **bool** |  | [optional] 
**Troca** | Pointer to **bool** |  | [optional] 

## Methods

### NewCompartimentoVM

`func NewCompartimentoVM() *CompartimentoVM`

NewCompartimentoVM instantiates a new CompartimentoVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartimentoVMWithDefaults

`func NewCompartimentoVMWithDefaults() *CompartimentoVM`

NewCompartimentoVMWithDefaults instantiates a new CompartimentoVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompartimentoVM) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompartimentoVM) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompartimentoVM) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CompartimentoVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CompartimentoVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CompartimentoVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNumero

`func (o *CompartimentoVM) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *CompartimentoVM) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *CompartimentoVM) SetNumero(v string)`

SetNumero sets Numero field to given value.

### HasNumero

`func (o *CompartimentoVM) HasNumero() bool`

HasNumero returns a boolean if a field has been set.

### SetNumeroNil

`func (o *CompartimentoVM) SetNumeroNil(b bool)`

 SetNumeroNil sets the value for Numero to be an explicit nil

### UnsetNumero
`func (o *CompartimentoVM) UnsetNumero()`

UnsetNumero ensures that no value is present for Numero, not even an explicit nil
### GetTipo

`func (o *CompartimentoVM) GetTipo() TipoCompartimento`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *CompartimentoVM) GetTipoOk() (*TipoCompartimento, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *CompartimentoVM) SetTipo(v TipoCompartimento)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *CompartimentoVM) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetAtivo

`func (o *CompartimentoVM) GetAtivo() bool`

GetAtivo returns the Ativo field if non-nil, zero value otherwise.

### GetAtivoOk

`func (o *CompartimentoVM) GetAtivoOk() (*bool, bool)`

GetAtivoOk returns a tuple with the Ativo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtivo

`func (o *CompartimentoVM) SetAtivo(v bool)`

SetAtivo sets Ativo field to given value.

### HasAtivo

`func (o *CompartimentoVM) HasAtivo() bool`

HasAtivo returns a boolean if a field has been set.

### SetAtivoNil

`func (o *CompartimentoVM) SetAtivoNil(b bool)`

 SetAtivoNil sets the value for Ativo to be an explicit nil

### UnsetAtivo
`func (o *CompartimentoVM) UnsetAtivo()`

UnsetAtivo ensures that no value is present for Ativo, not even an explicit nil
### GetRefGvm

`func (o *CompartimentoVM) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *CompartimentoVM) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *CompartimentoVM) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *CompartimentoVM) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *CompartimentoVM) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *CompartimentoVM) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetRefProduto

`func (o *CompartimentoVM) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *CompartimentoVM) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *CompartimentoVM) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *CompartimentoVM) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *CompartimentoVM) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *CompartimentoVM) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetLivreSemProduto

`func (o *CompartimentoVM) GetLivreSemProduto() bool`

GetLivreSemProduto returns the LivreSemProduto field if non-nil, zero value otherwise.

### GetLivreSemProdutoOk

`func (o *CompartimentoVM) GetLivreSemProdutoOk() (*bool, bool)`

GetLivreSemProdutoOk returns a tuple with the LivreSemProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivreSemProduto

`func (o *CompartimentoVM) SetLivreSemProduto(v bool)`

SetLivreSemProduto sets LivreSemProduto field to given value.

### HasLivreSemProduto

`func (o *CompartimentoVM) HasLivreSemProduto() bool`

HasLivreSemProduto returns a boolean if a field has been set.

### GetCompartimentoBloqueado

`func (o *CompartimentoVM) GetCompartimentoBloqueado() bool`

GetCompartimentoBloqueado returns the CompartimentoBloqueado field if non-nil, zero value otherwise.

### GetCompartimentoBloqueadoOk

`func (o *CompartimentoVM) GetCompartimentoBloqueadoOk() (*bool, bool)`

GetCompartimentoBloqueadoOk returns a tuple with the CompartimentoBloqueado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentoBloqueado

`func (o *CompartimentoVM) SetCompartimentoBloqueado(v bool)`

SetCompartimentoBloqueado sets CompartimentoBloqueado field to given value.

### HasCompartimentoBloqueado

`func (o *CompartimentoVM) HasCompartimentoBloqueado() bool`

HasCompartimentoBloqueado returns a boolean if a field has been set.

### GetVasilhame

`func (o *CompartimentoVM) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *CompartimentoVM) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *CompartimentoVM) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *CompartimentoVM) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### GetTroca

`func (o *CompartimentoVM) GetTroca() bool`

GetTroca returns the Troca field if non-nil, zero value otherwise.

### GetTrocaOk

`func (o *CompartimentoVM) GetTrocaOk() (*bool, bool)`

GetTrocaOk returns a tuple with the Troca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroca

`func (o *CompartimentoVM) SetTroca(v bool)`

SetTroca sets Troca field to given value.

### HasTroca

`func (o *CompartimentoVM) HasTroca() bool`

HasTroca returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


