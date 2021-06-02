# Produto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to [**TipoProduto**](TipoProduto.md) |  | [optional] 
**Descricao** | Pointer to **NullableString** |  | [optional] 
**Preco** | Pointer to **NullableFloat64** |  | [optional] 
**PresoTotal** | Pointer to **NullableString** |  | [optional] 
**PesoLiquido** | Pointer to **NullableString** |  | [optional] 
**Dimensoes** | Pointer to **NullableString** |  | [optional] 
**RefImagem** | Pointer to **NullableInt64** |  | [optional] 
**Imagem** | Pointer to [**Imagem**](Imagem.md) |  | [optional] 
**CompartimentoHistoricos** | Pointer to [**[]CompartimentoHistorico**](CompartimentoHistorico.md) |  | [optional] 
**Compartimentos** | Pointer to [**[]Compartimento**](Compartimento.md) |  | [optional] 
**Vendas** | Pointer to [**[]Venda**](Venda.md) |  | [optional] 
**OrdemServicoProdutos** | Pointer to [**[]OrdemServicoProduto**](OrdemServicoProduto.md) |  | [optional] 
**OrdemServicoCompartimentos** | Pointer to [**[]OrdemServicoCompartimento**](OrdemServicoCompartimento.md) |  | [optional] 

## Methods

### NewProduto

`func NewProduto() *Produto`

NewProduto instantiates a new Produto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProdutoWithDefaults

`func NewProdutoWithDefaults() *Produto`

NewProdutoWithDefaults instantiates a new Produto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Produto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Produto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Produto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Produto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Produto) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Produto) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Produto) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Produto) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Produto) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Produto) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Produto) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Produto) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Produto) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Produto) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Produto) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Produto) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Produto) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Produto) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Produto) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Produto) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Produto) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Produto) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Produto) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Produto) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Produto) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Produto) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetRefEntidade

`func (o *Produto) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *Produto) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *Produto) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *Produto) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *Produto) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *Produto) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetTipo

`func (o *Produto) GetTipo() TipoProduto`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Produto) GetTipoOk() (*TipoProduto, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Produto) SetTipo(v TipoProduto)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Produto) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetDescricao

`func (o *Produto) GetDescricao() string`

GetDescricao returns the Descricao field if non-nil, zero value otherwise.

### GetDescricaoOk

`func (o *Produto) GetDescricaoOk() (*string, bool)`

GetDescricaoOk returns a tuple with the Descricao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescricao

`func (o *Produto) SetDescricao(v string)`

SetDescricao sets Descricao field to given value.

### HasDescricao

`func (o *Produto) HasDescricao() bool`

HasDescricao returns a boolean if a field has been set.

### SetDescricaoNil

`func (o *Produto) SetDescricaoNil(b bool)`

 SetDescricaoNil sets the value for Descricao to be an explicit nil

### UnsetDescricao
`func (o *Produto) UnsetDescricao()`

UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
### GetPreco

`func (o *Produto) GetPreco() float64`

GetPreco returns the Preco field if non-nil, zero value otherwise.

### GetPrecoOk

`func (o *Produto) GetPrecoOk() (*float64, bool)`

GetPrecoOk returns a tuple with the Preco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreco

`func (o *Produto) SetPreco(v float64)`

SetPreco sets Preco field to given value.

### HasPreco

`func (o *Produto) HasPreco() bool`

HasPreco returns a boolean if a field has been set.

### SetPrecoNil

`func (o *Produto) SetPrecoNil(b bool)`

 SetPrecoNil sets the value for Preco to be an explicit nil

### UnsetPreco
`func (o *Produto) UnsetPreco()`

UnsetPreco ensures that no value is present for Preco, not even an explicit nil
### GetPresoTotal

`func (o *Produto) GetPresoTotal() string`

GetPresoTotal returns the PresoTotal field if non-nil, zero value otherwise.

### GetPresoTotalOk

`func (o *Produto) GetPresoTotalOk() (*string, bool)`

GetPresoTotalOk returns a tuple with the PresoTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresoTotal

`func (o *Produto) SetPresoTotal(v string)`

SetPresoTotal sets PresoTotal field to given value.

### HasPresoTotal

`func (o *Produto) HasPresoTotal() bool`

HasPresoTotal returns a boolean if a field has been set.

### SetPresoTotalNil

`func (o *Produto) SetPresoTotalNil(b bool)`

 SetPresoTotalNil sets the value for PresoTotal to be an explicit nil

### UnsetPresoTotal
`func (o *Produto) UnsetPresoTotal()`

UnsetPresoTotal ensures that no value is present for PresoTotal, not even an explicit nil
### GetPesoLiquido

`func (o *Produto) GetPesoLiquido() string`

GetPesoLiquido returns the PesoLiquido field if non-nil, zero value otherwise.

### GetPesoLiquidoOk

`func (o *Produto) GetPesoLiquidoOk() (*string, bool)`

GetPesoLiquidoOk returns a tuple with the PesoLiquido field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPesoLiquido

`func (o *Produto) SetPesoLiquido(v string)`

SetPesoLiquido sets PesoLiquido field to given value.

### HasPesoLiquido

`func (o *Produto) HasPesoLiquido() bool`

HasPesoLiquido returns a boolean if a field has been set.

### SetPesoLiquidoNil

`func (o *Produto) SetPesoLiquidoNil(b bool)`

 SetPesoLiquidoNil sets the value for PesoLiquido to be an explicit nil

### UnsetPesoLiquido
`func (o *Produto) UnsetPesoLiquido()`

UnsetPesoLiquido ensures that no value is present for PesoLiquido, not even an explicit nil
### GetDimensoes

`func (o *Produto) GetDimensoes() string`

GetDimensoes returns the Dimensoes field if non-nil, zero value otherwise.

### GetDimensoesOk

`func (o *Produto) GetDimensoesOk() (*string, bool)`

GetDimensoesOk returns a tuple with the Dimensoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensoes

`func (o *Produto) SetDimensoes(v string)`

SetDimensoes sets Dimensoes field to given value.

### HasDimensoes

`func (o *Produto) HasDimensoes() bool`

HasDimensoes returns a boolean if a field has been set.

### SetDimensoesNil

`func (o *Produto) SetDimensoesNil(b bool)`

 SetDimensoesNil sets the value for Dimensoes to be an explicit nil

### UnsetDimensoes
`func (o *Produto) UnsetDimensoes()`

UnsetDimensoes ensures that no value is present for Dimensoes, not even an explicit nil
### GetRefImagem

`func (o *Produto) GetRefImagem() int64`

GetRefImagem returns the RefImagem field if non-nil, zero value otherwise.

### GetRefImagemOk

`func (o *Produto) GetRefImagemOk() (*int64, bool)`

GetRefImagemOk returns a tuple with the RefImagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefImagem

`func (o *Produto) SetRefImagem(v int64)`

SetRefImagem sets RefImagem field to given value.

### HasRefImagem

`func (o *Produto) HasRefImagem() bool`

HasRefImagem returns a boolean if a field has been set.

### SetRefImagemNil

`func (o *Produto) SetRefImagemNil(b bool)`

 SetRefImagemNil sets the value for RefImagem to be an explicit nil

### UnsetRefImagem
`func (o *Produto) UnsetRefImagem()`

UnsetRefImagem ensures that no value is present for RefImagem, not even an explicit nil
### GetImagem

`func (o *Produto) GetImagem() Imagem`

GetImagem returns the Imagem field if non-nil, zero value otherwise.

### GetImagemOk

`func (o *Produto) GetImagemOk() (*Imagem, bool)`

GetImagemOk returns a tuple with the Imagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagem

`func (o *Produto) SetImagem(v Imagem)`

SetImagem sets Imagem field to given value.

### HasImagem

`func (o *Produto) HasImagem() bool`

HasImagem returns a boolean if a field has been set.

### GetCompartimentoHistoricos

`func (o *Produto) GetCompartimentoHistoricos() []CompartimentoHistorico`

GetCompartimentoHistoricos returns the CompartimentoHistoricos field if non-nil, zero value otherwise.

### GetCompartimentoHistoricosOk

`func (o *Produto) GetCompartimentoHistoricosOk() (*[]CompartimentoHistorico, bool)`

GetCompartimentoHistoricosOk returns a tuple with the CompartimentoHistoricos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentoHistoricos

`func (o *Produto) SetCompartimentoHistoricos(v []CompartimentoHistorico)`

SetCompartimentoHistoricos sets CompartimentoHistoricos field to given value.

### HasCompartimentoHistoricos

`func (o *Produto) HasCompartimentoHistoricos() bool`

HasCompartimentoHistoricos returns a boolean if a field has been set.

### SetCompartimentoHistoricosNil

`func (o *Produto) SetCompartimentoHistoricosNil(b bool)`

 SetCompartimentoHistoricosNil sets the value for CompartimentoHistoricos to be an explicit nil

### UnsetCompartimentoHistoricos
`func (o *Produto) UnsetCompartimentoHistoricos()`

UnsetCompartimentoHistoricos ensures that no value is present for CompartimentoHistoricos, not even an explicit nil
### GetCompartimentos

`func (o *Produto) GetCompartimentos() []Compartimento`

GetCompartimentos returns the Compartimentos field if non-nil, zero value otherwise.

### GetCompartimentosOk

`func (o *Produto) GetCompartimentosOk() (*[]Compartimento, bool)`

GetCompartimentosOk returns a tuple with the Compartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentos

`func (o *Produto) SetCompartimentos(v []Compartimento)`

SetCompartimentos sets Compartimentos field to given value.

### HasCompartimentos

`func (o *Produto) HasCompartimentos() bool`

HasCompartimentos returns a boolean if a field has been set.

### SetCompartimentosNil

`func (o *Produto) SetCompartimentosNil(b bool)`

 SetCompartimentosNil sets the value for Compartimentos to be an explicit nil

### UnsetCompartimentos
`func (o *Produto) UnsetCompartimentos()`

UnsetCompartimentos ensures that no value is present for Compartimentos, not even an explicit nil
### GetVendas

`func (o *Produto) GetVendas() []Venda`

GetVendas returns the Vendas field if non-nil, zero value otherwise.

### GetVendasOk

`func (o *Produto) GetVendasOk() (*[]Venda, bool)`

GetVendasOk returns a tuple with the Vendas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendas

`func (o *Produto) SetVendas(v []Venda)`

SetVendas sets Vendas field to given value.

### HasVendas

`func (o *Produto) HasVendas() bool`

HasVendas returns a boolean if a field has been set.

### SetVendasNil

`func (o *Produto) SetVendasNil(b bool)`

 SetVendasNil sets the value for Vendas to be an explicit nil

### UnsetVendas
`func (o *Produto) UnsetVendas()`

UnsetVendas ensures that no value is present for Vendas, not even an explicit nil
### GetOrdemServicoProdutos

`func (o *Produto) GetOrdemServicoProdutos() []OrdemServicoProduto`

GetOrdemServicoProdutos returns the OrdemServicoProdutos field if non-nil, zero value otherwise.

### GetOrdemServicoProdutosOk

`func (o *Produto) GetOrdemServicoProdutosOk() (*[]OrdemServicoProduto, bool)`

GetOrdemServicoProdutosOk returns a tuple with the OrdemServicoProdutos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServicoProdutos

`func (o *Produto) SetOrdemServicoProdutos(v []OrdemServicoProduto)`

SetOrdemServicoProdutos sets OrdemServicoProdutos field to given value.

### HasOrdemServicoProdutos

`func (o *Produto) HasOrdemServicoProdutos() bool`

HasOrdemServicoProdutos returns a boolean if a field has been set.

### SetOrdemServicoProdutosNil

`func (o *Produto) SetOrdemServicoProdutosNil(b bool)`

 SetOrdemServicoProdutosNil sets the value for OrdemServicoProdutos to be an explicit nil

### UnsetOrdemServicoProdutos
`func (o *Produto) UnsetOrdemServicoProdutos()`

UnsetOrdemServicoProdutos ensures that no value is present for OrdemServicoProdutos, not even an explicit nil
### GetOrdemServicoCompartimentos

`func (o *Produto) GetOrdemServicoCompartimentos() []OrdemServicoCompartimento`

GetOrdemServicoCompartimentos returns the OrdemServicoCompartimentos field if non-nil, zero value otherwise.

### GetOrdemServicoCompartimentosOk

`func (o *Produto) GetOrdemServicoCompartimentosOk() (*[]OrdemServicoCompartimento, bool)`

GetOrdemServicoCompartimentosOk returns a tuple with the OrdemServicoCompartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServicoCompartimentos

`func (o *Produto) SetOrdemServicoCompartimentos(v []OrdemServicoCompartimento)`

SetOrdemServicoCompartimentos sets OrdemServicoCompartimentos field to given value.

### HasOrdemServicoCompartimentos

`func (o *Produto) HasOrdemServicoCompartimentos() bool`

HasOrdemServicoCompartimentos returns a boolean if a field has been set.

### SetOrdemServicoCompartimentosNil

`func (o *Produto) SetOrdemServicoCompartimentosNil(b bool)`

 SetOrdemServicoCompartimentosNil sets the value for OrdemServicoCompartimentos to be an explicit nil

### UnsetOrdemServicoCompartimentos
`func (o *Produto) UnsetOrdemServicoCompartimentos()`

UnsetOrdemServicoCompartimentos ensures that no value is present for OrdemServicoCompartimentos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


