# OrdemServico

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**RefUtilizador** | Pointer to **NullableInt64** |  | [optional] 
**DataCriacao** | Pointer to **NullableTime** |  | [optional] 
**DataEstado** | Pointer to **NullableTime** |  | [optional] 
**RefOperador** | Pointer to **NullableInt64** |  | [optional] 
**Estado** | Pointer to [**EstadosOrdemServico**](EstadosOrdemServico.md) |  | [optional] 
**Gvm** | Pointer to [**Gvm**](Gvm.md) |  | [optional] 
**Utilizador** | Pointer to [**Utilizador**](Utilizador.md) |  | [optional] 
**OrdemServicoCompartimentos** | Pointer to [**[]OrdemServicoCompartimento**](OrdemServicoCompartimento.md) |  | [optional] 
**OrdemServicoProdutos** | Pointer to [**[]OrdemServicoProduto**](OrdemServicoProduto.md) |  | [optional] 

## Methods

### NewOrdemServico

`func NewOrdemServico() *OrdemServico`

NewOrdemServico instantiates a new OrdemServico object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdemServicoWithDefaults

`func NewOrdemServicoWithDefaults() *OrdemServico`

NewOrdemServicoWithDefaults instantiates a new OrdemServico object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrdemServico) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdemServico) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdemServico) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrdemServico) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *OrdemServico) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrdemServico) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrdemServico) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OrdemServico) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *OrdemServico) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *OrdemServico) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *OrdemServico) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *OrdemServico) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *OrdemServico) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *OrdemServico) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *OrdemServico) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *OrdemServico) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *OrdemServico) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *OrdemServico) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *OrdemServico) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *OrdemServico) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetToken

`func (o *OrdemServico) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrdemServico) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrdemServico) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrdemServico) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *OrdemServico) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *OrdemServico) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetRefGvm

`func (o *OrdemServico) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *OrdemServico) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *OrdemServico) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *OrdemServico) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *OrdemServico) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *OrdemServico) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetRefUtilizador

`func (o *OrdemServico) GetRefUtilizador() int64`

GetRefUtilizador returns the RefUtilizador field if non-nil, zero value otherwise.

### GetRefUtilizadorOk

`func (o *OrdemServico) GetRefUtilizadorOk() (*int64, bool)`

GetRefUtilizadorOk returns a tuple with the RefUtilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefUtilizador

`func (o *OrdemServico) SetRefUtilizador(v int64)`

SetRefUtilizador sets RefUtilizador field to given value.

### HasRefUtilizador

`func (o *OrdemServico) HasRefUtilizador() bool`

HasRefUtilizador returns a boolean if a field has been set.

### SetRefUtilizadorNil

`func (o *OrdemServico) SetRefUtilizadorNil(b bool)`

 SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil

### UnsetRefUtilizador
`func (o *OrdemServico) UnsetRefUtilizador()`

UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
### GetDataCriacao

`func (o *OrdemServico) GetDataCriacao() time.Time`

GetDataCriacao returns the DataCriacao field if non-nil, zero value otherwise.

### GetDataCriacaoOk

`func (o *OrdemServico) GetDataCriacaoOk() (*time.Time, bool)`

GetDataCriacaoOk returns a tuple with the DataCriacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCriacao

`func (o *OrdemServico) SetDataCriacao(v time.Time)`

SetDataCriacao sets DataCriacao field to given value.

### HasDataCriacao

`func (o *OrdemServico) HasDataCriacao() bool`

HasDataCriacao returns a boolean if a field has been set.

### SetDataCriacaoNil

`func (o *OrdemServico) SetDataCriacaoNil(b bool)`

 SetDataCriacaoNil sets the value for DataCriacao to be an explicit nil

### UnsetDataCriacao
`func (o *OrdemServico) UnsetDataCriacao()`

UnsetDataCriacao ensures that no value is present for DataCriacao, not even an explicit nil
### GetDataEstado

`func (o *OrdemServico) GetDataEstado() time.Time`

GetDataEstado returns the DataEstado field if non-nil, zero value otherwise.

### GetDataEstadoOk

`func (o *OrdemServico) GetDataEstadoOk() (*time.Time, bool)`

GetDataEstadoOk returns a tuple with the DataEstado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEstado

`func (o *OrdemServico) SetDataEstado(v time.Time)`

SetDataEstado sets DataEstado field to given value.

### HasDataEstado

`func (o *OrdemServico) HasDataEstado() bool`

HasDataEstado returns a boolean if a field has been set.

### SetDataEstadoNil

`func (o *OrdemServico) SetDataEstadoNil(b bool)`

 SetDataEstadoNil sets the value for DataEstado to be an explicit nil

### UnsetDataEstado
`func (o *OrdemServico) UnsetDataEstado()`

UnsetDataEstado ensures that no value is present for DataEstado, not even an explicit nil
### GetRefOperador

`func (o *OrdemServico) GetRefOperador() int64`

GetRefOperador returns the RefOperador field if non-nil, zero value otherwise.

### GetRefOperadorOk

`func (o *OrdemServico) GetRefOperadorOk() (*int64, bool)`

GetRefOperadorOk returns a tuple with the RefOperador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefOperador

`func (o *OrdemServico) SetRefOperador(v int64)`

SetRefOperador sets RefOperador field to given value.

### HasRefOperador

`func (o *OrdemServico) HasRefOperador() bool`

HasRefOperador returns a boolean if a field has been set.

### SetRefOperadorNil

`func (o *OrdemServico) SetRefOperadorNil(b bool)`

 SetRefOperadorNil sets the value for RefOperador to be an explicit nil

### UnsetRefOperador
`func (o *OrdemServico) UnsetRefOperador()`

UnsetRefOperador ensures that no value is present for RefOperador, not even an explicit nil
### GetEstado

`func (o *OrdemServico) GetEstado() EstadosOrdemServico`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *OrdemServico) GetEstadoOk() (*EstadosOrdemServico, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *OrdemServico) SetEstado(v EstadosOrdemServico)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *OrdemServico) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetGvm

`func (o *OrdemServico) GetGvm() Gvm`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *OrdemServico) GetGvmOk() (*Gvm, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *OrdemServico) SetGvm(v Gvm)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *OrdemServico) HasGvm() bool`

HasGvm returns a boolean if a field has been set.

### GetUtilizador

`func (o *OrdemServico) GetUtilizador() Utilizador`

GetUtilizador returns the Utilizador field if non-nil, zero value otherwise.

### GetUtilizadorOk

`func (o *OrdemServico) GetUtilizadorOk() (*Utilizador, bool)`

GetUtilizadorOk returns a tuple with the Utilizador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizador

`func (o *OrdemServico) SetUtilizador(v Utilizador)`

SetUtilizador sets Utilizador field to given value.

### HasUtilizador

`func (o *OrdemServico) HasUtilizador() bool`

HasUtilizador returns a boolean if a field has been set.

### GetOrdemServicoCompartimentos

`func (o *OrdemServico) GetOrdemServicoCompartimentos() []OrdemServicoCompartimento`

GetOrdemServicoCompartimentos returns the OrdemServicoCompartimentos field if non-nil, zero value otherwise.

### GetOrdemServicoCompartimentosOk

`func (o *OrdemServico) GetOrdemServicoCompartimentosOk() (*[]OrdemServicoCompartimento, bool)`

GetOrdemServicoCompartimentosOk returns a tuple with the OrdemServicoCompartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServicoCompartimentos

`func (o *OrdemServico) SetOrdemServicoCompartimentos(v []OrdemServicoCompartimento)`

SetOrdemServicoCompartimentos sets OrdemServicoCompartimentos field to given value.

### HasOrdemServicoCompartimentos

`func (o *OrdemServico) HasOrdemServicoCompartimentos() bool`

HasOrdemServicoCompartimentos returns a boolean if a field has been set.

### SetOrdemServicoCompartimentosNil

`func (o *OrdemServico) SetOrdemServicoCompartimentosNil(b bool)`

 SetOrdemServicoCompartimentosNil sets the value for OrdemServicoCompartimentos to be an explicit nil

### UnsetOrdemServicoCompartimentos
`func (o *OrdemServico) UnsetOrdemServicoCompartimentos()`

UnsetOrdemServicoCompartimentos ensures that no value is present for OrdemServicoCompartimentos, not even an explicit nil
### GetOrdemServicoProdutos

`func (o *OrdemServico) GetOrdemServicoProdutos() []OrdemServicoProduto`

GetOrdemServicoProdutos returns the OrdemServicoProdutos field if non-nil, zero value otherwise.

### GetOrdemServicoProdutosOk

`func (o *OrdemServico) GetOrdemServicoProdutosOk() (*[]OrdemServicoProduto, bool)`

GetOrdemServicoProdutosOk returns a tuple with the OrdemServicoProdutos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServicoProdutos

`func (o *OrdemServico) SetOrdemServicoProdutos(v []OrdemServicoProduto)`

SetOrdemServicoProdutos sets OrdemServicoProdutos field to given value.

### HasOrdemServicoProdutos

`func (o *OrdemServico) HasOrdemServicoProdutos() bool`

HasOrdemServicoProdutos returns a boolean if a field has been set.

### SetOrdemServicoProdutosNil

`func (o *OrdemServico) SetOrdemServicoProdutosNil(b bool)`

 SetOrdemServicoProdutosNil sets the value for OrdemServicoProdutos to be an explicit nil

### UnsetOrdemServicoProdutos
`func (o *OrdemServico) UnsetOrdemServicoProdutos()`

UnsetOrdemServicoProdutos ensures that no value is present for OrdemServicoProdutos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


