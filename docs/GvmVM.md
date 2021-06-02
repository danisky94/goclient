# GvmVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nome** | Pointer to **NullableString** |  | [optional] 
**Modelo** | Pointer to **NullableString** |  | [optional] 
**NumeroSerie** | Pointer to **NullableString** |  | [optional] 
**DataProducao** | Pointer to **NullableTime** |  | [optional] 
**VersaoFirmware** | Pointer to **NullableString** |  | [optional] 
**VersaoSoftware** | Pointer to **NullableString** |  | [optional] 
**UltimaComunicacao** | Pointer to **NullableTime** |  | [optional] 
**RefLocal** | Pointer to **NullableInt64** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**RefNo** | Pointer to **NullableInt64** |  | [optional] 
**Latitude** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGvmVM

`func NewGvmVM() *GvmVM`

NewGvmVM instantiates a new GvmVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGvmVMWithDefaults

`func NewGvmVMWithDefaults() *GvmVM`

NewGvmVMWithDefaults instantiates a new GvmVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNome

`func (o *GvmVM) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *GvmVM) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *GvmVM) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *GvmVM) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *GvmVM) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *GvmVM) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetModelo

`func (o *GvmVM) GetModelo() string`

GetModelo returns the Modelo field if non-nil, zero value otherwise.

### GetModeloOk

`func (o *GvmVM) GetModeloOk() (*string, bool)`

GetModeloOk returns a tuple with the Modelo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelo

`func (o *GvmVM) SetModelo(v string)`

SetModelo sets Modelo field to given value.

### HasModelo

`func (o *GvmVM) HasModelo() bool`

HasModelo returns a boolean if a field has been set.

### SetModeloNil

`func (o *GvmVM) SetModeloNil(b bool)`

 SetModeloNil sets the value for Modelo to be an explicit nil

### UnsetModelo
`func (o *GvmVM) UnsetModelo()`

UnsetModelo ensures that no value is present for Modelo, not even an explicit nil
### GetNumeroSerie

`func (o *GvmVM) GetNumeroSerie() string`

GetNumeroSerie returns the NumeroSerie field if non-nil, zero value otherwise.

### GetNumeroSerieOk

`func (o *GvmVM) GetNumeroSerieOk() (*string, bool)`

GetNumeroSerieOk returns a tuple with the NumeroSerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroSerie

`func (o *GvmVM) SetNumeroSerie(v string)`

SetNumeroSerie sets NumeroSerie field to given value.

### HasNumeroSerie

`func (o *GvmVM) HasNumeroSerie() bool`

HasNumeroSerie returns a boolean if a field has been set.

### SetNumeroSerieNil

`func (o *GvmVM) SetNumeroSerieNil(b bool)`

 SetNumeroSerieNil sets the value for NumeroSerie to be an explicit nil

### UnsetNumeroSerie
`func (o *GvmVM) UnsetNumeroSerie()`

UnsetNumeroSerie ensures that no value is present for NumeroSerie, not even an explicit nil
### GetDataProducao

`func (o *GvmVM) GetDataProducao() time.Time`

GetDataProducao returns the DataProducao field if non-nil, zero value otherwise.

### GetDataProducaoOk

`func (o *GvmVM) GetDataProducaoOk() (*time.Time, bool)`

GetDataProducaoOk returns a tuple with the DataProducao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProducao

`func (o *GvmVM) SetDataProducao(v time.Time)`

SetDataProducao sets DataProducao field to given value.

### HasDataProducao

`func (o *GvmVM) HasDataProducao() bool`

HasDataProducao returns a boolean if a field has been set.

### SetDataProducaoNil

`func (o *GvmVM) SetDataProducaoNil(b bool)`

 SetDataProducaoNil sets the value for DataProducao to be an explicit nil

### UnsetDataProducao
`func (o *GvmVM) UnsetDataProducao()`

UnsetDataProducao ensures that no value is present for DataProducao, not even an explicit nil
### GetVersaoFirmware

`func (o *GvmVM) GetVersaoFirmware() string`

GetVersaoFirmware returns the VersaoFirmware field if non-nil, zero value otherwise.

### GetVersaoFirmwareOk

`func (o *GvmVM) GetVersaoFirmwareOk() (*string, bool)`

GetVersaoFirmwareOk returns a tuple with the VersaoFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoFirmware

`func (o *GvmVM) SetVersaoFirmware(v string)`

SetVersaoFirmware sets VersaoFirmware field to given value.

### HasVersaoFirmware

`func (o *GvmVM) HasVersaoFirmware() bool`

HasVersaoFirmware returns a boolean if a field has been set.

### SetVersaoFirmwareNil

`func (o *GvmVM) SetVersaoFirmwareNil(b bool)`

 SetVersaoFirmwareNil sets the value for VersaoFirmware to be an explicit nil

### UnsetVersaoFirmware
`func (o *GvmVM) UnsetVersaoFirmware()`

UnsetVersaoFirmware ensures that no value is present for VersaoFirmware, not even an explicit nil
### GetVersaoSoftware

`func (o *GvmVM) GetVersaoSoftware() string`

GetVersaoSoftware returns the VersaoSoftware field if non-nil, zero value otherwise.

### GetVersaoSoftwareOk

`func (o *GvmVM) GetVersaoSoftwareOk() (*string, bool)`

GetVersaoSoftwareOk returns a tuple with the VersaoSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoSoftware

`func (o *GvmVM) SetVersaoSoftware(v string)`

SetVersaoSoftware sets VersaoSoftware field to given value.

### HasVersaoSoftware

`func (o *GvmVM) HasVersaoSoftware() bool`

HasVersaoSoftware returns a boolean if a field has been set.

### SetVersaoSoftwareNil

`func (o *GvmVM) SetVersaoSoftwareNil(b bool)`

 SetVersaoSoftwareNil sets the value for VersaoSoftware to be an explicit nil

### UnsetVersaoSoftware
`func (o *GvmVM) UnsetVersaoSoftware()`

UnsetVersaoSoftware ensures that no value is present for VersaoSoftware, not even an explicit nil
### GetUltimaComunicacao

`func (o *GvmVM) GetUltimaComunicacao() time.Time`

GetUltimaComunicacao returns the UltimaComunicacao field if non-nil, zero value otherwise.

### GetUltimaComunicacaoOk

`func (o *GvmVM) GetUltimaComunicacaoOk() (*time.Time, bool)`

GetUltimaComunicacaoOk returns a tuple with the UltimaComunicacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimaComunicacao

`func (o *GvmVM) SetUltimaComunicacao(v time.Time)`

SetUltimaComunicacao sets UltimaComunicacao field to given value.

### HasUltimaComunicacao

`func (o *GvmVM) HasUltimaComunicacao() bool`

HasUltimaComunicacao returns a boolean if a field has been set.

### SetUltimaComunicacaoNil

`func (o *GvmVM) SetUltimaComunicacaoNil(b bool)`

 SetUltimaComunicacaoNil sets the value for UltimaComunicacao to be an explicit nil

### UnsetUltimaComunicacao
`func (o *GvmVM) UnsetUltimaComunicacao()`

UnsetUltimaComunicacao ensures that no value is present for UltimaComunicacao, not even an explicit nil
### GetRefLocal

`func (o *GvmVM) GetRefLocal() int64`

GetRefLocal returns the RefLocal field if non-nil, zero value otherwise.

### GetRefLocalOk

`func (o *GvmVM) GetRefLocalOk() (*int64, bool)`

GetRefLocalOk returns a tuple with the RefLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefLocal

`func (o *GvmVM) SetRefLocal(v int64)`

SetRefLocal sets RefLocal field to given value.

### HasRefLocal

`func (o *GvmVM) HasRefLocal() bool`

HasRefLocal returns a boolean if a field has been set.

### SetRefLocalNil

`func (o *GvmVM) SetRefLocalNil(b bool)`

 SetRefLocalNil sets the value for RefLocal to be an explicit nil

### UnsetRefLocal
`func (o *GvmVM) UnsetRefLocal()`

UnsetRefLocal ensures that no value is present for RefLocal, not even an explicit nil
### GetRefEntidade

`func (o *GvmVM) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *GvmVM) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *GvmVM) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *GvmVM) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *GvmVM) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *GvmVM) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetRefNo

`func (o *GvmVM) GetRefNo() int64`

GetRefNo returns the RefNo field if non-nil, zero value otherwise.

### GetRefNoOk

`func (o *GvmVM) GetRefNoOk() (*int64, bool)`

GetRefNoOk returns a tuple with the RefNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNo

`func (o *GvmVM) SetRefNo(v int64)`

SetRefNo sets RefNo field to given value.

### HasRefNo

`func (o *GvmVM) HasRefNo() bool`

HasRefNo returns a boolean if a field has been set.

### SetRefNoNil

`func (o *GvmVM) SetRefNoNil(b bool)`

 SetRefNoNil sets the value for RefNo to be an explicit nil

### UnsetRefNo
`func (o *GvmVM) UnsetRefNo()`

UnsetRefNo ensures that no value is present for RefNo, not even an explicit nil
### GetLatitude

`func (o *GvmVM) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GvmVM) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GvmVM) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GvmVM) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *GvmVM) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *GvmVM) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *GvmVM) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GvmVM) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GvmVM) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GvmVM) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *GvmVM) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *GvmVM) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


