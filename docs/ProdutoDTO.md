# ProdutoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoProduto**](TipoProduto.md) |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**Preco** | Pointer to **NullableFloat64** |  | [optional] 
**RefImagem** | Pointer to **NullableInt64** |  | [optional] 
**PresoTotal** | Pointer to **NullableString** |  | [optional] 
**PesoLiquido** | Pointer to **NullableString** |  | [optional] 
**Dimensoes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProdutoDTO

`func NewProdutoDTO() *ProdutoDTO`

NewProdutoDTO instantiates a new ProdutoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProdutoDTOWithDefaults

`func NewProdutoDTOWithDefaults() *ProdutoDTO`

NewProdutoDTOWithDefaults instantiates a new ProdutoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProdutoDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProdutoDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProdutoDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProdutoDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNome

`func (o *ProdutoDTO) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *ProdutoDTO) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *ProdutoDTO) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *ProdutoDTO) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *ProdutoDTO) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *ProdutoDTO) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetRefEntidade

`func (o *ProdutoDTO) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *ProdutoDTO) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *ProdutoDTO) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *ProdutoDTO) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *ProdutoDTO) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *ProdutoDTO) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetTipo

`func (o *ProdutoDTO) GetTipo() TipoProduto`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *ProdutoDTO) GetTipoOk() (*TipoProduto, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *ProdutoDTO) SetTipo(v TipoProduto)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *ProdutoDTO) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetDescricao

`func (o *ProdutoDTO) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *ProdutoDTO) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *ProdutoDTO) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *ProdutoDTO) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *ProdutoDTO) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *ProdutoDTO) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetPreco

`func (o *ProdutoDTO) GetPreco() float64`

GetPreco returns the Preco field if non-nil, zero value otherwise.

### GetPrecoOk

`func (o *ProdutoDTO) GetPrecoOk() (*float64, bool)`

GetPrecoOk returns a tuple with the Preco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreco

`func (o *ProdutoDTO) SetPreco(v float64)`

SetPreco sets Preco field to given value.

### HasPreco

`func (o *ProdutoDTO) HasPreco() bool`

HasPreco returns a boolean if a field has been set.

### SetPrecoNil

`func (o *ProdutoDTO) SetPrecoNil(b bool)`

 SetPrecoNil sets the value for Preco to be an explicit nil

### UnsetPreco
`func (o *ProdutoDTO) UnsetPreco()`

UnsetPreco ensures that no value is present for Preco, not even an explicit nil
### GetRefImagem

`func (o *ProdutoDTO) GetRefImagem() int64`

GetRefImagem returns the RefImagem field if non-nil, zero value otherwise.

### GetRefImagemOk

`func (o *ProdutoDTO) GetRefImagemOk() (*int64, bool)`

GetRefImagemOk returns a tuple with the RefImagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefImagem

`func (o *ProdutoDTO) SetRefImagem(v int64)`

SetRefImagem sets RefImagem field to given value.

### HasRefImagem

`func (o *ProdutoDTO) HasRefImagem() bool`

HasRefImagem returns a boolean if a field has been set.

### SetRefImagemNil

`func (o *ProdutoDTO) SetRefImagemNil(b bool)`

 SetRefImagemNil sets the value for RefImagem to be an explicit nil

### UnsetRefImagem
`func (o *ProdutoDTO) UnsetRefImagem()`

UnsetRefImagem ensures that no value is present for RefImagem, not even an explicit nil
### GetPresoTotal

`func (o *ProdutoDTO) GetPresoTotal() string`

GetPresoTotal returns the PresoTotal field if non-nil, zero value otherwise.

### GetPresoTotalOk

`func (o *ProdutoDTO) GetPresoTotalOk() (*string, bool)`

GetPresoTotalOk returns a tuple with the PresoTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresoTotal

`func (o *ProdutoDTO) SetPresoTotal(v string)`

SetPresoTotal sets PresoTotal field to given value.

### HasPresoTotal

`func (o *ProdutoDTO) HasPresoTotal() bool`

HasPresoTotal returns a boolean if a field has been set.

### SetPresoTotalNil

`func (o *ProdutoDTO) SetPresoTotalNil(b bool)`

 SetPresoTotalNil sets the value for PresoTotal to be an explicit nil

### UnsetPresoTotal
`func (o *ProdutoDTO) UnsetPresoTotal()`

UnsetPresoTotal ensures that no value is present for PresoTotal, not even an explicit nil
### GetPesoLiquido

`func (o *ProdutoDTO) GetPesoLiquido() string`

GetPesoLiquido returns the PesoLiquido field if non-nil, zero value otherwise.

### GetPesoLiquidoOk

`func (o *ProdutoDTO) GetPesoLiquidoOk() (*string, bool)`

GetPesoLiquidoOk returns a tuple with the PesoLiquido field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesoLiquido

`func (o *ProdutoDTO) SetPesoLiquido(v string)`

SetPesoLiquido sets PesoLiquido field to given value.

### HasPesoLiquido

`func (o *ProdutoDTO) HasPesoLiquido() bool`

HasPesoLiquido returns a boolean if a field has been set.

### SetPesoLiquidoNil

`func (o *ProdutoDTO) SetPesoLiquidoNil(b bool)`

 SetPesoLiquidoNil sets the value for PesoLiquido to be an explicit nil

### UnsetPesoLiquido
`func (o *ProdutoDTO) UnsetPesoLiquido()`

UnsetPesoLiquido ensures that no value is present for PesoLiquido, not even an explicit nil
### GetDimensoes

`func (o *ProdutoDTO) GetDimensoes() string`

GetDimensoes returns the Dimensoes field if non-nil, zero value otherwise.

### GetDimensoesOk

`func (o *ProdutoDTO) GetDimensoesOk() (*string, bool)`

GetDimensoesOk returns a tuple with the Dimensoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensoes

`func (o *ProdutoDTO) SetDimensoes(v string)`

SetDimensoes sets Dimensoes field to given value.

### HasDimensoes

`func (o *ProdutoDTO) HasDimensoes() bool`

HasDimensoes returns a boolean if a field has been set.

### SetDimensoesNil

`func (o *ProdutoDTO) SetDimensoesNil(b bool)`

 SetDimensoesNil sets the value for Dimensoes to be an explicit nil

### UnsetDimensoes
`func (o *ProdutoDTO) UnsetDimensoes()`

UnsetDimensoes ensures that no value is present for Dimensoes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


