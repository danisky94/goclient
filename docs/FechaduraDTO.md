# FechaduraDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Endereco** | Pointer to **NullableString** |  | [optional] 
**Canal** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFechaduraDTO

`func NewFechaduraDTO() *FechaduraDTO`

NewFechaduraDTO instantiates a new FechaduraDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFechaduraDTOWithDefaults

`func NewFechaduraDTOWithDefaults() *FechaduraDTO`

NewFechaduraDTOWithDefaults instantiates a new FechaduraDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FechaduraDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FechaduraDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FechaduraDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FechaduraDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefCompartimento

`func (o *FechaduraDTO) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *FechaduraDTO) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *FechaduraDTO) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *FechaduraDTO) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *FechaduraDTO) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *FechaduraDTO) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetEndereco

`func (o *FechaduraDTO) GetEndereco() string`

GetEndereco returns the Endereco field if non-nil, zero value otherwise.

### GetEnderecoOk

`func (o *FechaduraDTO) GetEnderecoOk() (*string, bool)`

GetEnderecoOk returns a tuple with the Endereco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndereco

`func (o *FechaduraDTO) SetEndereco(v string)`

SetEndereco sets Endereco field to given value.

### HasEndereco

`func (o *FechaduraDTO) HasEndereco() bool`

HasEndereco returns a boolean if a field has been set.

### SetEnderecoNil

`func (o *FechaduraDTO) SetEnderecoNil(b bool)`

 SetEnderecoNil sets the value for Endereco to be an explicit nil

### UnsetEndereco
`func (o *FechaduraDTO) UnsetEndereco()`

UnsetEndereco ensures that no value is present for Endereco, not even an explicit nil
### GetCanal

`func (o *FechaduraDTO) GetCanal() int32`

GetCanal returns the Canal field if non-nil, zero value otherwise.

### GetCanalOk

`func (o *FechaduraDTO) GetCanalOk() (*int32, bool)`

GetCanalOk returns a tuple with the Canal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanal

`func (o *FechaduraDTO) SetCanal(v int32)`

SetCanal sets Canal field to given value.

### HasCanal

`func (o *FechaduraDTO) HasCanal() bool`

HasCanal returns a boolean if a field has been set.

### SetCanalNil

`func (o *FechaduraDTO) SetCanalNil(b bool)`

 SetCanalNil sets the value for Canal to be an explicit nil

### UnsetCanal
`func (o *FechaduraDTO) UnsetCanal()`

UnsetCanal ensures that no value is present for Canal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


