# OrdemServicoProduto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**RefOrdemServico** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**Quant** | Pointer to **NullableInt32** |  | [optional] 
**Vasilhame** | Pointer to **bool** |  | [optional] 
**OrdemServico** | Pointer to [**OrdemServico**](OrdemServico.md) |  | [optional] 
**Produto** | Pointer to [**Produto**](Produto.md) |  | [optional] 

## Methods

### NewOrdemServicoProduto

`func NewOrdemServicoProduto() *OrdemServicoProduto`

NewOrdemServicoProduto instantiates a new OrdemServicoProduto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdemServicoProdutoWithDefaults

`func NewOrdemServicoProdutoWithDefaults() *OrdemServicoProduto`

NewOrdemServicoProdutoWithDefaults instantiates a new OrdemServicoProduto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrdemServicoProduto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrdemServicoProduto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrdemServicoProduto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OrdemServicoProduto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *OrdemServicoProduto) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *OrdemServicoProduto) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *OrdemServicoProduto) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *OrdemServicoProduto) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *OrdemServicoProduto) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *OrdemServicoProduto) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *OrdemServicoProduto) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *OrdemServicoProduto) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *OrdemServicoProduto) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *OrdemServicoProduto) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *OrdemServicoProduto) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *OrdemServicoProduto) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *OrdemServicoProduto) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *OrdemServicoProduto) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *OrdemServicoProduto) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *OrdemServicoProduto) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetRefOrdemServico

`func (o *OrdemServicoProduto) GetRefOrdemServico() int64`

GetRefOrdemServico returns the RefOrdemServico field if non-nil, zero value otherwise.

### GetRefOrdemServicoOk

`func (o *OrdemServicoProduto) GetRefOrdemServicoOk() (*int64, bool)`

GetRefOrdemServicoOk returns a tuple with the RefOrdemServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefOrdemServico

`func (o *OrdemServicoProduto) SetRefOrdemServico(v int64)`

SetRefOrdemServico sets RefOrdemServico field to given value.

### HasRefOrdemServico

`func (o *OrdemServicoProduto) HasRefOrdemServico() bool`

HasRefOrdemServico returns a boolean if a field has been set.

### SetRefOrdemServicoNil

`func (o *OrdemServicoProduto) SetRefOrdemServicoNil(b bool)`

 SetRefOrdemServicoNil sets the value for RefOrdemServico to be an explicit nil

### UnsetRefOrdemServico
`func (o *OrdemServicoProduto) UnsetRefOrdemServico()`

UnsetRefOrdemServico ensures that no value is present for RefOrdemServico, not even an explicit nil
### GetRefProduto

`func (o *OrdemServicoProduto) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *OrdemServicoProduto) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *OrdemServicoProduto) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *OrdemServicoProduto) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *OrdemServicoProduto) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *OrdemServicoProduto) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetQuant

`func (o *OrdemServicoProduto) GetQuant() int32`

GetQuant returns the Quant field if non-nil, zero value otherwise.

### GetQuantOk

`func (o *OrdemServicoProduto) GetQuantOk() (*int32, bool)`

GetQuantOk returns a tuple with the Quant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuant

`func (o *OrdemServicoProduto) SetQuant(v int32)`

SetQuant sets Quant field to given value.

### HasQuant

`func (o *OrdemServicoProduto) HasQuant() bool`

HasQuant returns a boolean if a field has been set.

### SetQuantNil

`func (o *OrdemServicoProduto) SetQuantNil(b bool)`

 SetQuantNil sets the value for Quant to be an explicit nil

### UnsetQuant
`func (o *OrdemServicoProduto) UnsetQuant()`

UnsetQuant ensures that no value is present for Quant, not even an explicit nil
### GetVasilhame

`func (o *OrdemServicoProduto) GetVasilhame() bool`

GetVasilhame returns the Vasilhame field if non-nil, zero value otherwise.

### GetVasilhameOk

`func (o *OrdemServicoProduto) GetVasilhameOk() (*bool, bool)`

GetVasilhameOk returns a tuple with the Vasilhame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVasilhame

`func (o *OrdemServicoProduto) SetVasilhame(v bool)`

SetVasilhame sets Vasilhame field to given value.

### HasVasilhame

`func (o *OrdemServicoProduto) HasVasilhame() bool`

HasVasilhame returns a boolean if a field has been set.

### GetOrdemServico

`func (o *OrdemServicoProduto) GetOrdemServico() OrdemServico`

GetOrdemServico returns the OrdemServico field if non-nil, zero value otherwise.

### GetOrdemServicoOk

`func (o *OrdemServicoProduto) GetOrdemServicoOk() (*OrdemServico, bool)`

GetOrdemServicoOk returns a tuple with the OrdemServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServico

`func (o *OrdemServicoProduto) SetOrdemServico(v OrdemServico)`

SetOrdemServico sets OrdemServico field to given value.

### HasOrdemServico

`func (o *OrdemServicoProduto) HasOrdemServico() bool`

HasOrdemServico returns a boolean if a field has been set.

### GetProduto

`func (o *OrdemServicoProduto) GetProduto() Produto`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *OrdemServicoProduto) GetProdutoOk() (*Produto, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *OrdemServicoProduto) SetProduto(v Produto)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *OrdemServicoProduto) HasProduto() bool`

HasProduto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


