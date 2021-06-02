# Imagem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Tamanho** | Pointer to **NullableInt64** |  | [optional] 
**Tipo** | Pointer to **NullableString** |  | [optional] 
**Conteudo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImagem

`func NewImagem() *Imagem`

NewImagem instantiates a new Imagem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagemWithDefaults

`func NewImagemWithDefaults() *Imagem`

NewImagemWithDefaults instantiates a new Imagem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Imagem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Imagem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Imagem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Imagem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Imagem) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Imagem) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Imagem) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Imagem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Imagem) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Imagem) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Imagem) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Imagem) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Imagem) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Imagem) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Imagem) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Imagem) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Imagem) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Imagem) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Imagem) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Imagem) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Imagem) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Imagem) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Imagem) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Imagem) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Imagem) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Imagem) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetTamanho

`func (o *Imagem) GetTamanho() int64`

GetTamanho returns the Tamanho field if non-nil, zero value otherwise.

### GetTamanhoOk

`func (o *Imagem) GetTamanhoOk() (*int64, bool)`

GetTamanhoOk returns a tuple with the Tamanho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTamanho

`func (o *Imagem) SetTamanho(v int64)`

SetTamanho sets Tamanho field to given value.

### HasTamanho

`func (o *Imagem) HasTamanho() bool`

HasTamanho returns a boolean if a field has been set.

### SetTamanhoNil

`func (o *Imagem) SetTamanhoNil(b bool)`

 SetTamanhoNil sets the value for Tamanho to be an explicit nil

### UnsetTamanho
`func (o *Imagem) UnsetTamanho()`

UnsetTamanho ensures that no value is present for Tamanho, not even an explicit nil
### GetTipo

`func (o *Imagem) GetTipo() string`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Imagem) GetTipoOk() (*string, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Imagem) SetTipo(v string)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Imagem) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### SetTipoNil

`func (o *Imagem) SetTipoNil(b bool)`

 SetTipoNil sets the value for Tipo to be an explicit nil

### UnsetTipo
`func (o *Imagem) UnsetTipo()`

UnsetTipo ensures that no value is present for Tipo, not even an explicit nil
### GetConteudo

`func (o *Imagem) GetConteudo() string`

GetConteudo returns the Conteudo field if non-nil, zero value otherwise.

### GetConteudoOk

`func (o *Imagem) GetConteudoOk() (*string, bool)`

GetConteudoOk returns a tuple with the Conteudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConteudo

`func (o *Imagem) SetConteudo(v string)`

SetConteudo sets Conteudo field to given value.

### HasConteudo

`func (o *Imagem) HasConteudo() bool`

HasConteudo returns a boolean if a field has been set.

### SetConteudoNil

`func (o *Imagem) SetConteudoNil(b bool)`

 SetConteudoNil sets the value for Conteudo to be an explicit nil

### UnsetConteudo
`func (o *Imagem) UnsetConteudo()`

UnsetConteudo ensures that no value is present for Conteudo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


