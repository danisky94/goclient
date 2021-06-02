# CompartimentoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Numero** | Pointer to **NullableString** |  | [optional] 
**RefGVM** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoCompartimento**](TipoCompartimento.md) |  | [optional] 
**Ativo** | Pointer to **bool** |  | [optional] 
**Fechaduras** | Pointer to [**[]FechaduraDTO**](FechaduraDTO.md) |  | [optional] 

## Methods

### NewCompartimentoDTO

`func NewCompartimentoDTO() *CompartimentoDTO`

NewCompartimentoDTO instantiates a new CompartimentoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartimentoDTOWithDefaults

`func NewCompartimentoDTOWithDefaults() *CompartimentoDTO`

NewCompartimentoDTOWithDefaults instantiates a new CompartimentoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompartimentoDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompartimentoDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompartimentoDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CompartimentoDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumero

`func (o *CompartimentoDTO) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *CompartimentoDTO) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *CompartimentoDTO) SetNumero(v string)`

SetNumero sets Numero field to given value.

### HasNumero

`func (o *CompartimentoDTO) HasNumero() bool`

HasNumero returns a boolean if a field has been set.

### SetNumeroNil

`func (o *CompartimentoDTO) SetNumeroNil(b bool)`

 SetNumeroNil sets the value for Numero to be an explicit nil

### UnsetNumero
`func (o *CompartimentoDTO) UnsetNumero()`

UnsetNumero ensures that no value is present for Numero, not even an explicit nil
### GetRefGVM

`func (o *CompartimentoDTO) GetRefGVM() int64`

GetRefGVM returns the RefGVM field if non-nil, zero value otherwise.

### GetRefGVMOk

`func (o *CompartimentoDTO) GetRefGVMOk() (*int64, bool)`

GetRefGVMOk returns a tuple with the RefGVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGVM

`func (o *CompartimentoDTO) SetRefGVM(v int64)`

SetRefGVM sets RefGVM field to given value.

### HasRefGVM

`func (o *CompartimentoDTO) HasRefGVM() bool`

HasRefGVM returns a boolean if a field has been set.

### SetRefGVMNil

`func (o *CompartimentoDTO) SetRefGVMNil(b bool)`

 SetRefGVMNil sets the value for RefGVM to be an explicit nil

### UnsetRefGVM
`func (o *CompartimentoDTO) UnsetRefGVM()`

UnsetRefGVM ensures that no value is present for RefGVM, not even an explicit nil
### GetRefProduto

`func (o *CompartimentoDTO) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *CompartimentoDTO) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *CompartimentoDTO) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *CompartimentoDTO) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *CompartimentoDTO) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *CompartimentoDTO) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetTipo

`func (o *CompartimentoDTO) GetTipo() TipoCompartimento`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *CompartimentoDTO) GetTipoOk() (*TipoCompartimento, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *CompartimentoDTO) SetTipo(v TipoCompartimento)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *CompartimentoDTO) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetAtivo

`func (o *CompartimentoDTO) GetAtivo() bool`

GetAtivo returns the Ativo field if non-nil, zero value otherwise.

### GetAtivoOk

`func (o *CompartimentoDTO) GetAtivoOk() (*bool, bool)`

GetAtivoOk returns a tuple with the Ativo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtivo

`func (o *CompartimentoDTO) SetAtivo(v bool)`

SetAtivo sets Ativo field to given value.

### HasAtivo

`func (o *CompartimentoDTO) HasAtivo() bool`

HasAtivo returns a boolean if a field has been set.

### GetFechaduras

`func (o *CompartimentoDTO) GetFechaduras() []FechaduraDTO`

GetFechaduras returns the Fechaduras field if non-nil, zero value otherwise.

### GetFechadurasOk

`func (o *CompartimentoDTO) GetFechadurasOk() (*[]FechaduraDTO, bool)`

GetFechadurasOk returns a tuple with the Fechaduras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaduras

`func (o *CompartimentoDTO) SetFechaduras(v []FechaduraDTO)`

SetFechaduras sets Fechaduras field to given value.

### HasFechaduras

`func (o *CompartimentoDTO) HasFechaduras() bool`

HasFechaduras returns a boolean if a field has been set.

### SetFechadurasNil

`func (o *CompartimentoDTO) SetFechadurasNil(b bool)`

 SetFechadurasNil sets the value for Fechaduras to be an explicit nil

### UnsetFechaduras
`func (o *CompartimentoDTO) UnsetFechaduras()`

UnsetFechaduras ensures that no value is present for Fechaduras, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


