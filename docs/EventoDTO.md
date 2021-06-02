# EventoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RefTipoEvento** | Pointer to **NullableInt64** |  | [optional] 
**Mensagem** | Pointer to **NullableString** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewEventoDTO

`func NewEventoDTO() *EventoDTO`

NewEventoDTO instantiates a new EventoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventoDTOWithDefaults

`func NewEventoDTOWithDefaults() *EventoDTO`

NewEventoDTOWithDefaults instantiates a new EventoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventoDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventoDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventoDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EventoDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefTipoEvento

`func (o *EventoDTO) GetRefTipoEvento() int64`

GetRefTipoEvento returns the RefTipoEvento field if non-nil, zero value otherwise.

### GetRefTipoEventoOk

`func (o *EventoDTO) GetRefTipoEventoOk() (*int64, bool)`

GetRefTipoEventoOk returns a tuple with the RefTipoEvento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefTipoEvento

`func (o *EventoDTO) SetRefTipoEvento(v int64)`

SetRefTipoEvento sets RefTipoEvento field to given value.

### HasRefTipoEvento

`func (o *EventoDTO) HasRefTipoEvento() bool`

HasRefTipoEvento returns a boolean if a field has been set.

### SetRefTipoEventoNil

`func (o *EventoDTO) SetRefTipoEventoNil(b bool)`

 SetRefTipoEventoNil sets the value for RefTipoEvento to be an explicit nil

### UnsetRefTipoEvento
`func (o *EventoDTO) UnsetRefTipoEvento()`

UnsetRefTipoEvento ensures that no value is present for RefTipoEvento, not even an explicit nil
### GetMensagem

`func (o *EventoDTO) GetMensagem() string`

GetMensagem returns the Mensagem field if non-nil, zero value otherwise.

### GetMensagemOk

`func (o *EventoDTO) GetMensagemOk() (*string, bool)`

GetMensagemOk returns a tuple with the Mensagem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensagem

`func (o *EventoDTO) SetMensagem(v string)`

SetMensagem sets Mensagem field to given value.

### HasMensagem

`func (o *EventoDTO) HasMensagem() bool`

HasMensagem returns a boolean if a field has been set.

### SetMensagemNil

`func (o *EventoDTO) SetMensagemNil(b bool)`

 SetMensagemNil sets the value for Mensagem to be an explicit nil

### UnsetMensagem
`func (o *EventoDTO) UnsetMensagem()`

UnsetMensagem ensures that no value is present for Mensagem, not even an explicit nil
### GetRefGvm

`func (o *EventoDTO) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *EventoDTO) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *EventoDTO) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *EventoDTO) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *EventoDTO) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *EventoDTO) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


