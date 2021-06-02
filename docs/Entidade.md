# Entidade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Tipo** | Pointer to [**TipoEntidade**](TipoEntidade.md) |  | [optional] 
**Morada** | Pointer to **NullableString** |  | [optional] 
**Telefone** | Pointer to **NullableInt32** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Nif** | Pointer to **NullableInt32** |  | [optional] 
**RefImagem** | Pointer to **NullableInt64** |  | [optional] 
**CorPrincipal** | Pointer to **NullableString** |  | [optional] 
**CorSecundaria** | Pointer to **NullableString** |  | [optional] 
**RefPlano** | Pointer to **NullableInt64** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**Imagem** | Pointer to [**Imagem**](Imagem.md) |  | [optional] 
**RefPlanoNavigation** | Pointer to [**Plano**](Plano.md) |  | [optional] 
**Utilizadores** | Pointer to [**[]Utilizador**](Utilizador.md) |  | [optional] 

## Methods

### NewEntidade

`func NewEntidade() *Entidade`

NewEntidade instantiates a new Entidade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntidadeWithDefaults

`func NewEntidadeWithDefaults() *Entidade`

NewEntidadeWithDefaults instantiates a new Entidade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entidade) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entidade) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entidade) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Entidade) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Entidade) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Entidade) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Entidade) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Entidade) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Entidade) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Entidade) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Entidade) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Entidade) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Entidade) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Entidade) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Entidade) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Entidade) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Entidade) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Entidade) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Entidade) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Entidade) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Entidade) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Entidade) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Entidade) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Entidade) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Entidade) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Entidade) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetTipo

`func (o *Entidade) GetTipo() TipoEntidade`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Entidade) GetTipoOk() (*TipoEntidade, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Entidade) SetTipo(v TipoEntidade)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Entidade) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetMorada

`func (o *Entidade) GetMorada() string`

GetMorada returns the Morada field if non-nil, zero value otherwise.

### GetMoradaOk

`func (o *Entidade) GetMoradaOk() (*string, bool)`

GetMoradaOk returns a tuple with the Morada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorada

`func (o *Entidade) SetMorada(v string)`

SetMorada sets Morada field to given value.

### HasMorada

`func (o *Entidade) HasMorada() bool`

HasMorada returns a boolean if a field has been set.

### SetMoradaNil

`func (o *Entidade) SetMoradaNil(b bool)`

 SetMoradaNil sets the value for Morada to be an explicit nil

### UnsetMorada
`func (o *Entidade) UnsetMorada()`

UnsetMorada ensures that no value is present for Morada, not even an explicit nil
### GetTelefone

`func (o *Entidade) GetTelefone() int32`

GetTelefone returns the Telefone field if non-nil, zero value otherwise.

### GetTelefoneOk

`func (o *Entidade) GetTelefoneOk() (*int32, bool)`

GetTelefoneOk returns a tuple with the Telefone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelefone

`func (o *Entidade) SetTelefone(v int32)`

SetTelefone sets Telefone field to given value.

### HasTelefone

`func (o *Entidade) HasTelefone() bool`

HasTelefone returns a boolean if a field has been set.

### SetTelefoneNil

`func (o *Entidade) SetTelefoneNil(b bool)`

 SetTelefoneNil sets the value for Telefone to be an explicit nil

### UnsetTelefone
`func (o *Entidade) UnsetTelefone()`

UnsetTelefone ensures that no value is present for Telefone, not even an explicit nil
### GetEmail

`func (o *Entidade) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Entidade) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Entidade) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Entidade) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Entidade) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Entidade) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNif

`func (o *Entidade) GetNif() int32`

GetNif returns the Nif field if non-nil, zero value otherwise.

### GetNifOk

`func (o *Entidade) GetNifOk() (*int32, bool)`

GetNifOk returns a tuple with the Nif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNif

`func (o *Entidade) SetNif(v int32)`

SetNif sets Nif field to given value.

### HasNif

`func (o *Entidade) HasNif() bool`

HasNif returns a boolean if a field has been set.

### SetNifNil

`func (o *Entidade) SetNifNil(b bool)`

 SetNifNil sets the value for Nif to be an explicit nil

### UnsetNif
`func (o *Entidade) UnsetNif()`

UnsetNif ensures that no value is present for Nif, not even an explicit nil
### GetRefImagem

`func (o *Entidade) GetRefImagem() int64`

GetRefImagem returns the RefImagem field if non-nil, zero value otherwise.

### GetRefImagemOk

`func (o *Entidade) GetRefImagemOk() (*int64, bool)`

GetRefImagemOk returns a tuple with the RefImagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefImagem

`func (o *Entidade) SetRefImagem(v int64)`

SetRefImagem sets RefImagem field to given value.

### HasRefImagem

`func (o *Entidade) HasRefImagem() bool`

HasRefImagem returns a boolean if a field has been set.

### SetRefImagemNil

`func (o *Entidade) SetRefImagemNil(b bool)`

 SetRefImagemNil sets the value for RefImagem to be an explicit nil

### UnsetRefImagem
`func (o *Entidade) UnsetRefImagem()`

UnsetRefImagem ensures that no value is present for RefImagem, not even an explicit nil
### GetCorPrincipal

`func (o *Entidade) GetCorPrincipal() string`

GetCorPrincipal returns the CorPrincipal field if non-nil, zero value otherwise.

### GetCorPrincipalOk

`func (o *Entidade) GetCorPrincipalOk() (*string, bool)`

GetCorPrincipalOk returns a tuple with the CorPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorPrincipal

`func (o *Entidade) SetCorPrincipal(v string)`

SetCorPrincipal sets CorPrincipal field to given value.

### HasCorPrincipal

`func (o *Entidade) HasCorPrincipal() bool`

HasCorPrincipal returns a boolean if a field has been set.

### SetCorPrincipalNil

`func (o *Entidade) SetCorPrincipalNil(b bool)`

 SetCorPrincipalNil sets the value for CorPrincipal to be an explicit nil

### UnsetCorPrincipal
`func (o *Entidade) UnsetCorPrincipal()`

UnsetCorPrincipal ensures that no value is present for CorPrincipal, not even an explicit nil
### GetCorSecundaria

`func (o *Entidade) GetCorSecundaria() string`

GetCorSecundaria returns the CorSecundaria field if non-nil, zero value otherwise.

### GetCorSecundariaOk

`func (o *Entidade) GetCorSecundariaOk() (*string, bool)`

GetCorSecundariaOk returns a tuple with the CorSecundaria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorSecundaria

`func (o *Entidade) SetCorSecundaria(v string)`

SetCorSecundaria sets CorSecundaria field to given value.

### HasCorSecundaria

`func (o *Entidade) HasCorSecundaria() bool`

HasCorSecundaria returns a boolean if a field has been set.

### SetCorSecundariaNil

`func (o *Entidade) SetCorSecundariaNil(b bool)`

 SetCorSecundariaNil sets the value for CorSecundaria to be an explicit nil

### UnsetCorSecundaria
`func (o *Entidade) UnsetCorSecundaria()`

UnsetCorSecundaria ensures that no value is present for CorSecundaria, not even an explicit nil
### GetRefPlano

`func (o *Entidade) GetRefPlano() int64`

GetRefPlano returns the RefPlano field if non-nil, zero value otherwise.

### GetRefPlanoOk

`func (o *Entidade) GetRefPlanoOk() (*int64, bool)`

GetRefPlanoOk returns a tuple with the RefPlano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPlano

`func (o *Entidade) SetRefPlano(v int64)`

SetRefPlano sets RefPlano field to given value.

### HasRefPlano

`func (o *Entidade) HasRefPlano() bool`

HasRefPlano returns a boolean if a field has been set.

### SetRefPlanoNil

`func (o *Entidade) SetRefPlanoNil(b bool)`

 SetRefPlanoNil sets the value for RefPlano to be an explicit nil

### UnsetRefPlano
`func (o *Entidade) UnsetRefPlano()`

UnsetRefPlano ensures that no value is present for RefPlano, not even an explicit nil
### GetRefEntidade

`func (o *Entidade) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *Entidade) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *Entidade) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *Entidade) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *Entidade) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *Entidade) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetImagem

`func (o *Entidade) GetImagem() Imagem`

GetImagem returns the Imagem field if non-nil, zero value otherwise.

### GetImagemOk

`func (o *Entidade) GetImagemOk() (*Imagem, bool)`

GetImagemOk returns a tuple with the Imagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagem

`func (o *Entidade) SetImagem(v Imagem)`

SetImagem sets Imagem field to given value.

### HasImagem

`func (o *Entidade) HasImagem() bool`

HasImagem returns a boolean if a field has been set.

### GetRefPlanoNavigation

`func (o *Entidade) GetRefPlanoNavigation() Plano`

GetRefPlanoNavigation returns the RefPlanoNavigation field if non-nil, zero value otherwise.

### GetRefPlanoNavigationOk

`func (o *Entidade) GetRefPlanoNavigationOk() (*Plano, bool)`

GetRefPlanoNavigationOk returns a tuple with the RefPlanoNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefPlanoNavigation

`func (o *Entidade) SetRefPlanoNavigation(v Plano)`

SetRefPlanoNavigation sets RefPlanoNavigation field to given value.

### HasRefPlanoNavigation

`func (o *Entidade) HasRefPlanoNavigation() bool`

HasRefPlanoNavigation returns a boolean if a field has been set.

### GetUtilizadores

`func (o *Entidade) GetUtilizadores() []Utilizador`

GetUtilizadores returns the Utilizadores field if non-nil, zero value otherwise.

### GetUtilizadoresOk

`func (o *Entidade) GetUtilizadoresOk() (*[]Utilizador, bool)`

GetUtilizadoresOk returns a tuple with the Utilizadores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizadores

`func (o *Entidade) SetUtilizadores(v []Utilizador)`

SetUtilizadores sets Utilizadores field to given value.

### HasUtilizadores

`func (o *Entidade) HasUtilizadores() bool`

HasUtilizadores returns a boolean if a field has been set.

### SetUtilizadoresNil

`func (o *Entidade) SetUtilizadoresNil(b bool)`

 SetUtilizadoresNil sets the value for Utilizadores to be an explicit nil

### UnsetUtilizadores
`func (o *Entidade) UnsetUtilizadores()`

UnsetUtilizadores ensures that no value is present for Utilizadores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


