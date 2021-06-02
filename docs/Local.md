# Local

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Morada** | Pointer to **NullableString** |  | [optional] 
**RefAssistencia** | Pointer to **NullableInt64** |  | [optional] 
**RefGestorTecnico** | Pointer to **NullableInt64** |  | [optional] 
**RefReposicao** | Pointer to **NullableInt64** |  | [optional] 
**RefGestorReposicao** | Pointer to **NullableInt64** |  | [optional] 
**Latitude** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableString** |  | [optional] 
**Gvms** | Pointer to [**[]Gvm**](Gvm.md) |  | [optional] 

## Methods

### NewLocal

`func NewLocal() *Local`

NewLocal instantiates a new Local object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalWithDefaults

`func NewLocalWithDefaults() *Local`

NewLocalWithDefaults instantiates a new Local object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Local) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Local) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Local) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Local) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Local) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Local) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Local) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Local) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Local) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Local) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Local) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Local) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Local) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Local) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Local) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Local) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Local) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Local) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Local) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Local) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Local) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Local) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Local) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Local) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Local) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Local) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetMorada

`func (o *Local) GetMorada() string`

GetMorada returns the Morada field if non-nil, zero value otherwise.

### GetMoradaOk

`func (o *Local) GetMoradaOk() (*string, bool)`

GetMoradaOk returns a tuple with the Morada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorada

`func (o *Local) SetMorada(v string)`

SetMorada sets Morada field to given value.

### HasMorada

`func (o *Local) HasMorada() bool`

HasMorada returns a boolean if a field has been set.

### SetMoradaNil

`func (o *Local) SetMoradaNil(b bool)`

 SetMoradaNil sets the value for Morada to be an explicit nil

### UnsetMorada
`func (o *Local) UnsetMorada()`

UnsetMorada ensures that no value is present for Morada, not even an explicit nil
### GetRefAssistencia

`func (o *Local) GetRefAssistencia() int64`

GetRefAssistencia returns the RefAssistencia field if non-nil, zero value otherwise.

### GetRefAssistenciaOk

`func (o *Local) GetRefAssistenciaOk() (*int64, bool)`

GetRefAssistenciaOk returns a tuple with the RefAssistencia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefAssistencia

`func (o *Local) SetRefAssistencia(v int64)`

SetRefAssistencia sets RefAssistencia field to given value.

### HasRefAssistencia

`func (o *Local) HasRefAssistencia() bool`

HasRefAssistencia returns a boolean if a field has been set.

### SetRefAssistenciaNil

`func (o *Local) SetRefAssistenciaNil(b bool)`

 SetRefAssistenciaNil sets the value for RefAssistencia to be an explicit nil

### UnsetRefAssistencia
`func (o *Local) UnsetRefAssistencia()`

UnsetRefAssistencia ensures that no value is present for RefAssistencia, not even an explicit nil
### GetRefGestorTecnico

`func (o *Local) GetRefGestorTecnico() int64`

GetRefGestorTecnico returns the RefGestorTecnico field if non-nil, zero value otherwise.

### GetRefGestorTecnicoOk

`func (o *Local) GetRefGestorTecnicoOk() (*int64, bool)`

GetRefGestorTecnicoOk returns a tuple with the RefGestorTecnico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGestorTecnico

`func (o *Local) SetRefGestorTecnico(v int64)`

SetRefGestorTecnico sets RefGestorTecnico field to given value.

### HasRefGestorTecnico

`func (o *Local) HasRefGestorTecnico() bool`

HasRefGestorTecnico returns a boolean if a field has been set.

### SetRefGestorTecnicoNil

`func (o *Local) SetRefGestorTecnicoNil(b bool)`

 SetRefGestorTecnicoNil sets the value for RefGestorTecnico to be an explicit nil

### UnsetRefGestorTecnico
`func (o *Local) UnsetRefGestorTecnico()`

UnsetRefGestorTecnico ensures that no value is present for RefGestorTecnico, not even an explicit nil
### GetRefReposicao

`func (o *Local) GetRefReposicao() int64`

GetRefReposicao returns the RefReposicao field if non-nil, zero value otherwise.

### GetRefReposicaoOk

`func (o *Local) GetRefReposicaoOk() (*int64, bool)`

GetRefReposicaoOk returns a tuple with the RefReposicao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefReposicao

`func (o *Local) SetRefReposicao(v int64)`

SetRefReposicao sets RefReposicao field to given value.

### HasRefReposicao

`func (o *Local) HasRefReposicao() bool`

HasRefReposicao returns a boolean if a field has been set.

### SetRefReposicaoNil

`func (o *Local) SetRefReposicaoNil(b bool)`

 SetRefReposicaoNil sets the value for RefReposicao to be an explicit nil

### UnsetRefReposicao
`func (o *Local) UnsetRefReposicao()`

UnsetRefReposicao ensures that no value is present for RefReposicao, not even an explicit nil
### GetRefGestorReposicao

`func (o *Local) GetRefGestorReposicao() int64`

GetRefGestorReposicao returns the RefGestorReposicao field if non-nil, zero value otherwise.

### GetRefGestorReposicaoOk

`func (o *Local) GetRefGestorReposicaoOk() (*int64, bool)`

GetRefGestorReposicaoOk returns a tuple with the RefGestorReposicao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefGestorReposicao

`func (o *Local) SetRefGestorReposicao(v int64)`

SetRefGestorReposicao sets RefGestorReposicao field to given value.

### HasRefGestorReposicao

`func (o *Local) HasRefGestorReposicao() bool`

HasRefGestorReposicao returns a boolean if a field has been set.

### SetRefGestorReposicaoNil

`func (o *Local) SetRefGestorReposicaoNil(b bool)`

 SetRefGestorReposicaoNil sets the value for RefGestorReposicao to be an explicit nil

### UnsetRefGestorReposicao
`func (o *Local) UnsetRefGestorReposicao()`

UnsetRefGestorReposicao ensures that no value is present for RefGestorReposicao, not even an explicit nil
### GetLatitude

`func (o *Local) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Local) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Local) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Local) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Local) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Local) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Local) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Local) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Local) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Local) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Local) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Local) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetGvms

`func (o *Local) GetGvms() []Gvm`

GetGvms returns the Gvms field if non-nil, zero value otherwise.

### GetGvmsOk

`func (o *Local) GetGvmsOk() (*[]Gvm, bool)`

GetGvmsOk returns a tuple with the Gvms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvms

`func (o *Local) SetGvms(v []Gvm)`

SetGvms sets Gvms field to given value.

### HasGvms

`func (o *Local) HasGvms() bool`

HasGvms returns a boolean if a field has been set.

### SetGvmsNil

`func (o *Local) SetGvmsNil(b bool)`

 SetGvmsNil sets the value for Gvms to be an explicit nil

### UnsetGvms
`func (o *Local) UnsetGvms()`

UnsetGvms ensures that no value is present for Gvms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


