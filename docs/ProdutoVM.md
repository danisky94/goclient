# ProdutoVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoProduto**](TipoProduto.md) |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**Preco** | Pointer to **NullableFloat64** |  | [optional] 
**PresoTotal** | Pointer to **NullableString** |  | [optional] 
**PesoLiquido** | Pointer to **NullableString** |  | [optional] 
**Dimensoes** | Pointer to **NullableString** |  | [optional] 
**RefImagem** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewProdutoVM

`func NewProdutoVM() *ProdutoVM`

NewProdutoVM instantiates a new ProdutoVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProdutoVMWithDefaults

`func NewProdutoVMWithDefaults() *ProdutoVM`

NewProdutoVMWithDefaults instantiates a new ProdutoVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProdutoVM) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProdutoVM) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProdutoVM) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProdutoVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProdutoVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProdutoVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNome

`func (o *ProdutoVM) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *ProdutoVM) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *ProdutoVM) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *ProdutoVM) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *ProdutoVM) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *ProdutoVM) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetRefEntidade

`func (o *ProdutoVM) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *ProdutoVM) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *ProdutoVM) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *ProdutoVM) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *ProdutoVM) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *ProdutoVM) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetTipo

`func (o *ProdutoVM) GetTipo() TipoProduto`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *ProdutoVM) GetTipoOk() (*TipoProduto, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *ProdutoVM) SetTipo(v TipoProduto)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *ProdutoVM) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetDescricao

`func (o *ProdutoVM) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *ProdutoVM) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *ProdutoVM) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *ProdutoVM) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *ProdutoVM) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *ProdutoVM) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetPreco

`func (o *ProdutoVM) GetPreco() float64`

GetPreco returns the Preco field if non-nil, zero value otherwise.

### GetPrecoOk

`func (o *ProdutoVM) GetPrecoOk() (*float64, bool)`

GetPrecoOk returns a tuple with the Preco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreco

`func (o *ProdutoVM) SetPreco(v float64)`

SetPreco sets Preco field to given value.

### HasPreco

`func (o *ProdutoVM) HasPreco() bool`

HasPreco returns a boolean if a field has been set.

### SetPrecoNil

`func (o *ProdutoVM) SetPrecoNil(b bool)`

 SetPrecoNil sets the value for Preco to be an explicit nil

### UnsetPreco
`func (o *ProdutoVM) UnsetPreco()`

UnsetPreco ensures that no value is present for Preco, not even an explicit nil
### GetPresoTotal

`func (o *ProdutoVM) GetPresoTotal() string`

GetPresoTotal returns the PresoTotal field if non-nil, zero value otherwise.

### GetPresoTotalOk

`func (o *ProdutoVM) GetPresoTotalOk() (*string, bool)`

GetPresoTotalOk returns a tuple with the PresoTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresoTotal

`func (o *ProdutoVM) SetPresoTotal(v string)`

SetPresoTotal sets PresoTotal field to given value.

### HasPresoTotal

`func (o *ProdutoVM) HasPresoTotal() bool`

HasPresoTotal returns a boolean if a field has been set.

### SetPresoTotalNil

`func (o *ProdutoVM) SetPresoTotalNil(b bool)`

 SetPresoTotalNil sets the value for PresoTotal to be an explicit nil

### UnsetPresoTotal
`func (o *ProdutoVM) UnsetPresoTotal()`

UnsetPresoTotal ensures that no value is present for PresoTotal, not even an explicit nil
### GetPesoLiquido

`func (o *ProdutoVM) GetPesoLiquido() string`

GetPesoLiquido returns the PesoLiquido field if non-nil, zero value otherwise.

### GetPesoLiquidoOk

`func (o *ProdutoVM) GetPesoLiquidoOk() (*string, bool)`

GetPesoLiquidoOk returns a tuple with the PesoLiquido field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesoLiquido

`func (o *ProdutoVM) SetPesoLiquido(v string)`

SetPesoLiquido sets PesoLiquido field to given value.

### HasPesoLiquido

`func (o *ProdutoVM) HasPesoLiquido() bool`

HasPesoLiquido returns a boolean if a field has been set.

### SetPesoLiquidoNil

`func (o *ProdutoVM) SetPesoLiquidoNil(b bool)`

 SetPesoLiquidoNil sets the value for PesoLiquido to be an explicit nil

### UnsetPesoLiquido
`func (o *ProdutoVM) UnsetPesoLiquido()`

UnsetPesoLiquido ensures that no value is present for PesoLiquido, not even an explicit nil
### GetDimensoes

`func (o *ProdutoVM) GetDimensoes() string`

GetDimensoes returns the Dimensoes field if non-nil, zero value otherwise.

### GetDimensoesOk

`func (o *ProdutoVM) GetDimensoesOk() (*string, bool)`

GetDimensoesOk returns a tuple with the Dimensoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensoes

`func (o *ProdutoVM) SetDimensoes(v string)`

SetDimensoes sets Dimensoes field to given value.

### HasDimensoes

`func (o *ProdutoVM) HasDimensoes() bool`

HasDimensoes returns a boolean if a field has been set.

### SetDimensoesNil

`func (o *ProdutoVM) SetDimensoesNil(b bool)`

 SetDimensoesNil sets the value for Dimensoes to be an explicit nil

### UnsetDimensoes
`func (o *ProdutoVM) UnsetDimensoes()`

UnsetDimensoes ensures that no value is present for Dimensoes, not even an explicit nil
### GetRefImagem

`func (o *ProdutoVM) GetRefImagem() int64`

GetRefImagem returns the RefImagem field if non-nil, zero value otherwise.

### GetRefImagemOk

`func (o *ProdutoVM) GetRefImagemOk() (*int64, bool)`

GetRefImagemOk returns a tuple with the RefImagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefImagem

`func (o *ProdutoVM) SetRefImagem(v int64)`

SetRefImagem sets RefImagem field to given value.

### HasRefImagem

`func (o *ProdutoVM) HasRefImagem() bool`

HasRefImagem returns a boolean if a field has been set.

### SetRefImagemNil

`func (o *ProdutoVM) SetRefImagemNil(b bool)`

 SetRefImagemNil sets the value for RefImagem to be an explicit nil

### UnsetRefImagem
`func (o *ProdutoVM) UnsetRefImagem()`

UnsetRefImagem ensures that no value is present for RefImagem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


