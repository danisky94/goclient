# RaspberryPiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Identificador** | Pointer to **NullableString** |  | [optional] 
**RefGvm** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewRaspberryPiDTO

`func NewRaspberryPiDTO() *RaspberryPiDTO`

NewRaspberryPiDTO instantiates a new RaspberryPiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaspberryPiDTOWithDefaults

`func NewRaspberryPiDTOWithDefaults() *RaspberryPiDTO`

NewRaspberryPiDTOWithDefaults instantiates a new RaspberryPiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RaspberryPiDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RaspberryPiDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RaspberryPiDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RaspberryPiDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentificador

`func (o *RaspberryPiDTO) GetIdentificador() string`

GetIdentificador returns the Identificador field if non-nil, zero value otherwise.

### GetIdentificadorOk

`func (o *RaspberryPiDTO) GetIdentificadorOk() (*string, bool)`

GetIdentificadorOk returns a tuple with the Identificador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificador

`func (o *RaspberryPiDTO) SetIdentificador(v string)`

SetIdentificador sets Identificador field to given value.

### HasIdentificador

`func (o *RaspberryPiDTO) HasIdentificador() bool`

HasIdentificador returns a boolean if a field has been set.

### SetIdentificadorNil

`func (o *RaspberryPiDTO) SetIdentificadorNil(b bool)`

 SetIdentificadorNil sets the value for Identificador to be an explicit nil

### UnsetIdentificador
`func (o *RaspberryPiDTO) UnsetIdentificador()`

UnsetIdentificador ensures that no value is present for Identificador, not even an explicit nil
### GetRefGvm

`func (o *RaspberryPiDTO) GetRefGvm() int64`

GetRefGvm returns the RefGvm field if non-nil, zero value otherwise.

### GetRefGvmOk

`func (o *RaspberryPiDTO) GetRefGvmOk() (*int64, bool)`

GetRefGvmOk returns a tuple with the RefGvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGvm

`func (o *RaspberryPiDTO) SetRefGvm(v int64)`

SetRefGvm sets RefGvm field to given value.

### HasRefGvm

`func (o *RaspberryPiDTO) HasRefGvm() bool`

HasRefGvm returns a boolean if a field has been set.

### SetRefGvmNil

`func (o *RaspberryPiDTO) SetRefGvmNil(b bool)`

 SetRefGvmNil sets the value for RefGvm to be an explicit nil

### UnsetRefGvm
`func (o *RaspberryPiDTO) UnsetRefGvm()`

UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


