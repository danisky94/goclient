# Compartimento

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Numero** | Pointer to **NullableString** |  | [optional] 
**Tipo** | Pointer to [**TipoCompartimento**](TipoCompartimento.md) |  | [optional] 
**Ativo** | Pointer to **NullableBool** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 
**RefProduto** | Pointer to **NullableInt64** |  | [optional] 
**CompartimentoEstado** | Pointer to [**CompartimentoEstado**](CompartimentoEstado.md) |  | [optional] 
**Gvm** | Pointer to [**Gvm**](Gvm.md) |  | [optional] 
**Produto** | Pointer to [**Produto**](Produto.md) |  | [optional] 
**CompartimentoHistoricos** | Pointer to [**[]CompartimentoHistorico**](CompartimentoHistorico.md) |  | [optional] 
**Fechaduras** | Pointer to [**[]Fechadura**](Fechadura.md) |  | [optional] 
**Eventos** | Pointer to [**[]Evento**](Evento.md) |  | [optional] 
**OrdemServicoCompartimentos** | Pointer to [**[]OrdemServicoCompartimento**](OrdemServicoCompartimento.md) |  | [optional] 
**Leds** | Pointer to [**[]Led**](Led.md) |  | [optional] 

## Methods

### NewCompartimento

`func NewCompartimento() *Compartimento`

NewCompartimento instantiates a new Compartimento object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartimentoWithDefaults

`func NewCompartimentoWithDefaults() *Compartimento`

NewCompartimentoWithDefaults instantiates a new Compartimento object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Compartimento) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Compartimento) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Compartimento) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Compartimento) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Compartimento) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Compartimento) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Compartimento) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Compartimento) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Compartimento) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Compartimento) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Compartimento) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Compartimento) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Compartimento) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Compartimento) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Compartimento) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Compartimento) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Compartimento) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Compartimento) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Compartimento) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Compartimento) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNumero

`func (o *Compartimento) GetNumero() string`

GetNumero returns the Numero field if non-nil, zero value otherwise.

### GetNumeroOk

`func (o *Compartimento) GetNumeroOk() (*string, bool)`

GetNumeroOk returns a tuple with the Numero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumero

`func (o *Compartimento) SetNumero(v string)`

SetNumero sets Numero field to given value.

### HasNumero

`func (o *Compartimento) HasNumero() bool`

HasNumero returns a boolean if a field has been set.

### SetNumeroNil

`func (o *Compartimento) SetNumeroNil(b bool)`

 SetNumeroNil sets the value for Numero to be an explicit nil

### UnsetNumero
`func (o *Compartimento) UnsetNumero()`

UnsetNumero ensures that no value is present for Numero, not even an explicit nil
### GetTipo

`func (o *Compartimento) GetTipo() TipoCompartimento`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Compartimento) GetTipoOk() (*TipoCompartimento, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Compartimento) SetTipo(v TipoCompartimento)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Compartimento) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetAtivo

`func (o *Compartimento) GetAtivo() bool`

GetAtivo returns the Ativo field if non-nil, zero value otherwise.

### GetAtivoOk

`func (o *Compartimento) GetAtivoOk() (*bool, bool)`

GetAtivoOk returns a tuple with the Ativo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtivo

`func (o *Compartimento) SetAtivo(v bool)`

SetAtivo sets Ativo field to given value.

### HasAtivo

`func (o *Compartimento) HasAtivo() bool`

HasAtivo returns a boolean if a field has been set.

### SetAtivoNil

`func (o *Compartimento) SetAtivoNil(b bool)`

 SetAtivoNil sets the value for Ativo to be an explicit nil

### UnsetAtivo
`func (o *Compartimento) UnsetAtivo()`

UnsetAtivo ensures that no value is present for Ativo, not even an explicit nil
### GetRefGvm

`func (o *Compartimento) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *Compartimento) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *Compartimento) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *Compartimento) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *Compartimento) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *Compartimento) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
### GetRefProduto

`func (o *Compartimento) GetRefProduto() int64`

GetRefProduto returns the RefProduto field if non-nil, zero value otherwise.

### GetRefProdutoOk

`func (o *Compartimento) GetRefProdutoOk() (*int64, bool)`

GetRefProdutoOk returns a tuple with the RefProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefProduto

`func (o *Compartimento) SetRefProduto(v int64)`

SetRefProduto sets RefProduto field to given value.

### HasRefProduto

`func (o *Compartimento) HasRefProduto() bool`

HasRefProduto returns a boolean if a field has been set.

### SetRefProdutoNil

`func (o *Compartimento) SetRefProdutoNil(b bool)`

 SetRefProdutoNil sets the value for RefProduto to be an explicit nil

### UnsetRefProduto
`func (o *Compartimento) UnsetRefProduto()`

UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
### GetCompartimentoEstado

`func (o *Compartimento) GetCompartimentoEstado() CompartimentoEstado`

GetCompartimentoEstado returns the CompartimentoEstado field if non-nil, zero value otherwise.

### GetCompartimentoEstadoOk

`func (o *Compartimento) GetCompartimentoEstadoOk() (*CompartimentoEstado, bool)`

GetCompartimentoEstadoOk returns a tuple with the CompartimentoEstado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentoEstado

`func (o *Compartimento) SetCompartimentoEstado(v CompartimentoEstado)`

SetCompartimentoEstado sets CompartimentoEstado field to given value.

### HasCompartimentoEstado

`func (o *Compartimento) HasCompartimentoEstado() bool`

HasCompartimentoEstado returns a boolean if a field has been set.

### GetGvm

`func (o *Compartimento) GetGvm() Gvm`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *Compartimento) GetGvmOk() (*Gvm, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *Compartimento) SetGvm(v Gvm)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *Compartimento) HasGvm() bool`

HasGvm returns a boolean if a field has been set.

### GetProduto

`func (o *Compartimento) GetProduto() Produto`

GetProduto returns the Produto field if non-nil, zero value otherwise.

### GetProdutoOk

`func (o *Compartimento) GetProdutoOk() (*Produto, bool)`

GetProdutoOk returns a tuple with the Produto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduto

`func (o *Compartimento) SetProduto(v Produto)`

SetProduto sets Produto field to given value.

### HasProduto

`func (o *Compartimento) HasProduto() bool`

HasProduto returns a boolean if a field has been set.

### GetCompartimentoHistoricos

`func (o *Compartimento) GetCompartimentoHistoricos() []CompartimentoHistorico`

GetCompartimentoHistoricos returns the CompartimentoHistoricos field if non-nil, zero value otherwise.

### GetCompartimentoHistoricosOk

`func (o *Compartimento) GetCompartimentoHistoricosOk() (*[]CompartimentoHistorico, bool)`

GetCompartimentoHistoricosOk returns a tuple with the CompartimentoHistoricos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentoHistoricos

`func (o *Compartimento) SetCompartimentoHistoricos(v []CompartimentoHistorico)`

SetCompartimentoHistoricos sets CompartimentoHistoricos field to given value.

### HasCompartimentoHistoricos

`func (o *Compartimento) HasCompartimentoHistoricos() bool`

HasCompartimentoHistoricos returns a boolean if a field has been set.

### SetCompartimentoHistoricosNil

`func (o *Compartimento) SetCompartimentoHistoricosNil(b bool)`

 SetCompartimentoHistoricosNil sets the value for CompartimentoHistoricos to be an explicit nil

### UnsetCompartimentoHistoricos
`func (o *Compartimento) UnsetCompartimentoHistoricos()`

UnsetCompartimentoHistoricos ensures that no value is present for CompartimentoHistoricos, not even an explicit nil
### GetFechaduras

`func (o *Compartimento) GetFechaduras() []Fechadura`

GetFechaduras returns the Fechaduras field if non-nil, zero value otherwise.

### GetFechadurasOk

`func (o *Compartimento) GetFechadurasOk() (*[]Fechadura, bool)`

GetFechadurasOk returns a tuple with the Fechaduras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaduras

`func (o *Compartimento) SetFechaduras(v []Fechadura)`

SetFechaduras sets Fechaduras field to given value.

### HasFechaduras

`func (o *Compartimento) HasFechaduras() bool`

HasFechaduras returns a boolean if a field has been set.

### SetFechadurasNil

`func (o *Compartimento) SetFechadurasNil(b bool)`

 SetFechadurasNil sets the value for Fechaduras to be an explicit nil

### UnsetFechaduras
`func (o *Compartimento) UnsetFechaduras()`

UnsetFechaduras ensures that no value is present for Fechaduras, not even an explicit nil
### GetEventos

`func (o *Compartimento) GetEventos() []Evento`

GetEventos returns the Eventos field if non-nil, zero value otherwise.

### GetEventosOk

`func (o *Compartimento) GetEventosOk() (*[]Evento, bool)`

GetEventosOk returns a tuple with the Eventos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventos

`func (o *Compartimento) SetEventos(v []Evento)`

SetEventos sets Eventos field to given value.

### HasEventos

`func (o *Compartimento) HasEventos() bool`

HasEventos returns a boolean if a field has been set.

### SetEventosNil

`func (o *Compartimento) SetEventosNil(b bool)`

 SetEventosNil sets the value for Eventos to be an explicit nil

### UnsetEventos
`func (o *Compartimento) UnsetEventos()`

UnsetEventos ensures that no value is present for Eventos, not even an explicit nil
### GetOrdemServicoCompartimentos

`func (o *Compartimento) GetOrdemServicoCompartimentos() []OrdemServicoCompartimento`

GetOrdemServicoCompartimentos returns the OrdemServicoCompartimentos field if non-nil, zero value otherwise.

### GetOrdemServicoCompartimentosOk

`func (o *Compartimento) GetOrdemServicoCompartimentosOk() (*[]OrdemServicoCompartimento, bool)`

GetOrdemServicoCompartimentosOk returns a tuple with the OrdemServicoCompartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdemServicoCompartimentos

`func (o *Compartimento) SetOrdemServicoCompartimentos(v []OrdemServicoCompartimento)`

SetOrdemServicoCompartimentos sets OrdemServicoCompartimentos field to given value.

### HasOrdemServicoCompartimentos

`func (o *Compartimento) HasOrdemServicoCompartimentos() bool`

HasOrdemServicoCompartimentos returns a boolean if a field has been set.

### SetOrdemServicoCompartimentosNil

`func (o *Compartimento) SetOrdemServicoCompartimentosNil(b bool)`

 SetOrdemServicoCompartimentosNil sets the value for OrdemServicoCompartimentos to be an explicit nil

### UnsetOrdemServicoCompartimentos
`func (o *Compartimento) UnsetOrdemServicoCompartimentos()`

UnsetOrdemServicoCompartimentos ensures that no value is present for OrdemServicoCompartimentos, not even an explicit nil
### GetLeds

`func (o *Compartimento) GetLeds() []Led`

GetLeds returns the Leds field if non-nil, zero value otherwise.

### GetLedsOk

`func (o *Compartimento) GetLedsOk() (*[]Led, bool)`

GetLedsOk returns a tuple with the Leds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeds

`func (o *Compartimento) SetLeds(v []Led)`

SetLeds sets Leds field to given value.

### HasLeds

`func (o *Compartimento) HasLeds() bool`

HasLeds returns a boolean if a field has been set.

### SetLedsNil

`func (o *Compartimento) SetLedsNil(b bool)`

 SetLedsNil sets the value for Leds to be an explicit nil

### UnsetLeds
`func (o *Compartimento) UnsetLeds()`

UnsetLeds ensures that no value is present for Leds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


