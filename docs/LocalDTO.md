# LocalDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Morada** | Pointer to **NullableString** |  | [optional] 
**RefAssistencia** | Pointer to **NullableInt64** |  | [optional] 
**RefGestorTecnico** | Pointer to **NullableInt64** |  | [optional] 
**RefReposicao** | Pointer to **NullableInt64** |  | [optional] 
**RefGestorReposicao** | Pointer to **NullableInt64** |  | [optional] 
**Latitude** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLocalDTO

`func NewLocalDTO() *LocalDTO`

NewLocalDTO instantiates a new LocalDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalDTOWithDefaults

`func NewLocalDTOWithDefaults() *LocalDTO`

NewLocalDTOWithDefaults instantiates a new LocalDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LocalDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNome

`func (o *LocalDTO) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *LocalDTO) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *LocalDTO) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *LocalDTO) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *LocalDTO) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *LocalDTO) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetMorada

`func (o *LocalDTO) GetMorada() string`

GetMorada returns the Morada field if non-nil, zero value otherwise.

### GetMoradaOk

`func (o *LocalDTO) GetMoradaOk() (*string, bool)`

GetMoradaOk returns a tuple with the Morada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorada

`func (o *LocalDTO) SetMorada(v string)`

SetMorada sets Morada field to given value.

### HasMorada

`func (o *LocalDTO) HasMorada() bool`

HasMorada returns a boolean if a field has been set.

### SetMoradaNil

`func (o *LocalDTO) SetMoradaNil(b bool)`

 SetMoradaNil sets the value for Morada to be an explicit nil

### UnsetMorada
`func (o *LocalDTO) UnsetMorada()`

UnsetMorada ensures that no value is present for Morada, not even an explicit nil
### GetRefAssistencia

`func (o *LocalDTO) GetRefAssistencia() int64`

GetRefAssistencia returns the RefAssistencia field if non-nil, zero value otherwise.

### GetRefAssistenciaOk

`func (o *LocalDTO) GetRefAssistenciaOk() (*int64, bool)`

GetRefAssistenciaOk returns a tuple with the RefAssistencia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefAssistencia

`func (o *LocalDTO) SetRefAssistencia(v int64)`

SetRefAssistencia sets RefAssistencia field to given value.

### HasRefAssistencia

`func (o *LocalDTO) HasRefAssistencia() bool`

HasRefAssistencia returns a boolean if a field has been set.

### SetRefAssistenciaNil

`func (o *LocalDTO) SetRefAssistenciaNil(b bool)`

 SetRefAssistenciaNil sets the value for RefAssistencia to be an explicit nil

### UnsetRefAssistencia
`func (o *LocalDTO) UnsetRefAssistencia()`

UnsetRefAssistencia ensures that no value is present for RefAssistencia, not even an explicit nil
### GetRefGestorTecnico

`func (o *LocalDTO) GetRefGestorTecnico() int64`

GetRefGestorTecnico returns the RefGestorTecnico field if non-nil, zero value otherwise.

### GetRefGestorTecnicoOk

`func (o *LocalDTO) GetRefGestorTecnicoOk() (*int64, bool)`

GetRefGestorTecnicoOk returns a tuple with the RefGestorTecnico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGestorTecnico

`func (o *LocalDTO) SetRefGestorTecnico(v int64)`

SetRefGestorTecnico sets RefGestorTecnico field to given value.

### HasRefGestorTecnico

`func (o *LocalDTO) HasRefGestorTecnico() bool`

HasRefGestorTecnico returns a boolean if a field has been set.

### SetRefGestorTecnicoNil

`func (o *LocalDTO) SetRefGestorTecnicoNil(b bool)`

 SetRefGestorTecnicoNil sets the value for RefGestorTecnico to be an explicit nil

### UnsetRefGestorTecnico
`func (o *LocalDTO) UnsetRefGestorTecnico()`

UnsetRefGestorTecnico ensures that no value is present for RefGestorTecnico, not even an explicit nil
### GetRefReposicao

`func (o *LocalDTO) GetRefReposicao() int64`

GetRefReposicao returns the RefReposicao field if non-nil, zero value otherwise.

### GetRefReposicaoOk

`func (o *LocalDTO) GetRefReposicaoOk() (*int64, bool)`

GetRefReposicaoOk returns a tuple with the RefReposicao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefReposicao

`func (o *LocalDTO) SetRefReposicao(v int64)`

SetRefReposicao sets RefReposicao field to given value.

### HasRefReposicao

`func (o *LocalDTO) HasRefReposicao() bool`

HasRefReposicao returns a boolean if a field has been set.

### SetRefReposicaoNil

`func (o *LocalDTO) SetRefReposicaoNil(b bool)`

 SetRefReposicaoNil sets the value for RefReposicao to be an explicit nil

### UnsetRefReposicao
`func (o *LocalDTO) UnsetRefReposicao()`

UnsetRefReposicao ensures that no value is present for RefReposicao, not even an explicit nil
### GetRefGestorReposicao

`func (o *LocalDTO) GetRefGestorReposicao() int64`

GetRefGestorReposicao returns the RefGestorReposicao field if non-nil, zero value otherwise.

### GetRefGestorReposicaoOk

`func (o *LocalDTO) GetRefGestorReposicaoOk() (*int64, bool)`

GetRefGestorReposicaoOk returns a tuple with the RefGestorReposicao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGestorReposicao

`func (o *LocalDTO) SetRefGestorReposicao(v int64)`

SetRefGestorReposicao sets RefGestorReposicao field to given value.

### HasRefGestorReposicao

`func (o *LocalDTO) HasRefGestorReposicao() bool`

HasRefGestorReposicao returns a boolean if a field has been set.

### SetRefGestorReposicaoNil

`func (o *LocalDTO) SetRefGestorReposicaoNil(b bool)`

 SetRefGestorReposicaoNil sets the value for RefGestorReposicao to be an explicit nil

### UnsetRefGestorReposicao
`func (o *LocalDTO) UnsetRefGestorReposicao()`

UnsetRefGestorReposicao ensures that no value is present for RefGestorReposicao, not even an explicit nil
### GetLatitude

`func (o *LocalDTO) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LocalDTO) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LocalDTO) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *LocalDTO) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *LocalDTO) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *LocalDTO) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *LocalDTO) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LocalDTO) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LocalDTO) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *LocalDTO) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *LocalDTO) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *LocalDTO) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


