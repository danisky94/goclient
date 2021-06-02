# GvmDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
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
**Compartimentos** | Pointer to [**[]CompartimentoDTO**](CompartimentoDTO.md) |  | [optional] 
**CompartimentosProduto** | Pointer to [**[]ProdutoDTO**](ProdutoDTO.md) |  | [optional] 

## Methods

### NewGvmDTO

`func NewGvmDTO() *GvmDTO`

NewGvmDTO instantiates a new GvmDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGvmDTOWithDefaults

`func NewGvmDTOWithDefaults() *GvmDTO`

NewGvmDTOWithDefaults instantiates a new GvmDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GvmDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GvmDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GvmDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GvmDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNome

`func (o *GvmDTO) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *GvmDTO) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *GvmDTO) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *GvmDTO) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *GvmDTO) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *GvmDTO) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetModelo

`func (o *GvmDTO) GetModelo() string`

GetModelo returns the Modelo field if non-nil, zero value otherwise.

### GetModeloOk

`func (o *GvmDTO) GetModeloOk() (*string, bool)`

GetModeloOk returns a tuple with the Modelo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelo

`func (o *GvmDTO) SetModelo(v string)`

SetModelo sets Modelo field to given value.

### HasModelo

`func (o *GvmDTO) HasModelo() bool`

HasModelo returns a boolean if a field has been set.

### SetModeloNil

`func (o *GvmDTO) SetModeloNil(b bool)`

 SetModeloNil sets the value for Modelo to be an explicit nil

### UnsetModelo
`func (o *GvmDTO) UnsetModelo()`

UnsetModelo ensures that no value is present for Modelo, not even an explicit nil
### GetNumeroSerie

`func (o *GvmDTO) GetNumeroSerie() string`

GetNumeroSerie returns the NumeroSerie field if non-nil, zero value otherwise.

### GetNumeroSerieOk

`func (o *GvmDTO) GetNumeroSerieOk() (*string, bool)`

GetNumeroSerieOk returns a tuple with the NumeroSerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroSerie

`func (o *GvmDTO) SetNumeroSerie(v string)`

SetNumeroSerie sets NumeroSerie field to given value.

### HasNumeroSerie

`func (o *GvmDTO) HasNumeroSerie() bool`

HasNumeroSerie returns a boolean if a field has been set.

### SetNumeroSerieNil

`func (o *GvmDTO) SetNumeroSerieNil(b bool)`

 SetNumeroSerieNil sets the value for NumeroSerie to be an explicit nil

### UnsetNumeroSerie
`func (o *GvmDTO) UnsetNumeroSerie()`

UnsetNumeroSerie ensures that no value is present for NumeroSerie, not even an explicit nil
### GetDataProducao

`func (o *GvmDTO) GetDataProducao() time.Time`

GetDataProducao returns the DataProducao field if non-nil, zero value otherwise.

### GetDataProducaoOk

`func (o *GvmDTO) GetDataProducaoOk() (*time.Time, bool)`

GetDataProducaoOk returns a tuple with the DataProducao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProducao

`func (o *GvmDTO) SetDataProducao(v time.Time)`

SetDataProducao sets DataProducao field to given value.

### HasDataProducao

`func (o *GvmDTO) HasDataProducao() bool`

HasDataProducao returns a boolean if a field has been set.

### SetDataProducaoNil

`func (o *GvmDTO) SetDataProducaoNil(b bool)`

 SetDataProducaoNil sets the value for DataProducao to be an explicit nil

### UnsetDataProducao
`func (o *GvmDTO) UnsetDataProducao()`

UnsetDataProducao ensures that no value is present for DataProducao, not even an explicit nil
### GetVersaoFirmware

`func (o *GvmDTO) GetVersaoFirmware() string`

GetVersaoFirmware returns the VersaoFirmware field if non-nil, zero value otherwise.

### GetVersaoFirmwareOk

`func (o *GvmDTO) GetVersaoFirmwareOk() (*string, bool)`

GetVersaoFirmwareOk returns a tuple with the VersaoFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoFirmware

`func (o *GvmDTO) SetVersaoFirmware(v string)`

SetVersaoFirmware sets VersaoFirmware field to given value.

### HasVersaoFirmware

`func (o *GvmDTO) HasVersaoFirmware() bool`

HasVersaoFirmware returns a boolean if a field has been set.

### SetVersaoFirmwareNil

`func (o *GvmDTO) SetVersaoFirmwareNil(b bool)`

 SetVersaoFirmwareNil sets the value for VersaoFirmware to be an explicit nil

### UnsetVersaoFirmware
`func (o *GvmDTO) UnsetVersaoFirmware()`

UnsetVersaoFirmware ensures that no value is present for VersaoFirmware, not even an explicit nil
### GetVersaoSoftware

`func (o *GvmDTO) GetVersaoSoftware() string`

GetVersaoSoftware returns the VersaoSoftware field if non-nil, zero value otherwise.

### GetVersaoSoftwareOk

`func (o *GvmDTO) GetVersaoSoftwareOk() (*string, bool)`

GetVersaoSoftwareOk returns a tuple with the VersaoSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoSoftware

`func (o *GvmDTO) SetVersaoSoftware(v string)`

SetVersaoSoftware sets VersaoSoftware field to given value.

### HasVersaoSoftware

`func (o *GvmDTO) HasVersaoSoftware() bool`

HasVersaoSoftware returns a boolean if a field has been set.

### SetVersaoSoftwareNil

`func (o *GvmDTO) SetVersaoSoftwareNil(b bool)`

 SetVersaoSoftwareNil sets the value for VersaoSoftware to be an explicit nil

### UnsetVersaoSoftware
`func (o *GvmDTO) UnsetVersaoSoftware()`

UnsetVersaoSoftware ensures that no value is present for VersaoSoftware, not even an explicit nil
### GetUltimaComunicacao

`func (o *GvmDTO) GetUltimaComunicacao() time.Time`

GetUltimaComunicacao returns the UltimaComunicacao field if non-nil, zero value otherwise.

### GetUltimaComunicacaoOk

`func (o *GvmDTO) GetUltimaComunicacaoOk() (*time.Time, bool)`

GetUltimaComunicacaoOk returns a tuple with the UltimaComunicacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimaComunicacao

`func (o *GvmDTO) SetUltimaComunicacao(v time.Time)`

SetUltimaComunicacao sets UltimaComunicacao field to given value.

### HasUltimaComunicacao

`func (o *GvmDTO) HasUltimaComunicacao() bool`

HasUltimaComunicacao returns a boolean if a field has been set.

### SetUltimaComunicacaoNil

`func (o *GvmDTO) SetUltimaComunicacaoNil(b bool)`

 SetUltimaComunicacaoNil sets the value for UltimaComunicacao to be an explicit nil

### UnsetUltimaComunicacao
`func (o *GvmDTO) UnsetUltimaComunicacao()`

UnsetUltimaComunicacao ensures that no value is present for UltimaComunicacao, not even an explicit nil
### GetRefLocal

`func (o *GvmDTO) GetRefLocal() int64`

GetRefLocal returns the RefLocal field if non-nil, zero value otherwise.

### GetRefLocalOk

`func (o *GvmDTO) GetRefLocalOk() (*int64, bool)`

GetRefLocalOk returns a tuple with the RefLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefLocal

`func (o *GvmDTO) SetRefLocal(v int64)`

SetRefLocal sets RefLocal field to given value.

### HasRefLocal

`func (o *GvmDTO) HasRefLocal() bool`

HasRefLocal returns a boolean if a field has been set.

### SetRefLocalNil

`func (o *GvmDTO) SetRefLocalNil(b bool)`

 SetRefLocalNil sets the value for RefLocal to be an explicit nil

### UnsetRefLocal
`func (o *GvmDTO) UnsetRefLocal()`

UnsetRefLocal ensures that no value is present for RefLocal, not even an explicit nil
### GetRefEntidade

`func (o *GvmDTO) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *GvmDTO) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *GvmDTO) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *GvmDTO) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *GvmDTO) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *GvmDTO) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetRefNo

`func (o *GvmDTO) GetRefNo() int64`

GetRefNo returns the RefNo field if non-nil, zero value otherwise.

### GetRefNoOk

`func (o *GvmDTO) GetRefNoOk() (*int64, bool)`

GetRefNoOk returns a tuple with the RefNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNo

`func (o *GvmDTO) SetRefNo(v int64)`

SetRefNo sets RefNo field to given value.

### HasRefNo

`func (o *GvmDTO) HasRefNo() bool`

HasRefNo returns a boolean if a field has been set.

### SetRefNoNil

`func (o *GvmDTO) SetRefNoNil(b bool)`

 SetRefNoNil sets the value for RefNo to be an explicit nil

### UnsetRefNo
`func (o *GvmDTO) UnsetRefNo()`

UnsetRefNo ensures that no value is present for RefNo, not even an explicit nil
### GetLatitude

`func (o *GvmDTO) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *GvmDTO) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *GvmDTO) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *GvmDTO) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *GvmDTO) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *GvmDTO) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *GvmDTO) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *GvmDTO) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *GvmDTO) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *GvmDTO) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *GvmDTO) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *GvmDTO) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetCompartimentos

`func (o *GvmDTO) GetCompartimentos() []CompartimentoDTO`

GetCompartimentos returns the Compartimentos field if non-nil, zero value otherwise.

### GetCompartimentosOk

`func (o *GvmDTO) GetCompartimentosOk() (*[]CompartimentoDTO, bool)`

GetCompartimentosOk returns a tuple with the Compartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentos

`func (o *GvmDTO) SetCompartimentos(v []CompartimentoDTO)`

SetCompartimentos sets Compartimentos field to given value.

### HasCompartimentos

`func (o *GvmDTO) HasCompartimentos() bool`

HasCompartimentos returns a boolean if a field has been set.

### SetCompartimentosNil

`func (o *GvmDTO) SetCompartimentosNil(b bool)`

 SetCompartimentosNil sets the value for Compartimentos to be an explicit nil

### UnsetCompartimentos
`func (o *GvmDTO) UnsetCompartimentos()`

UnsetCompartimentos ensures that no value is present for Compartimentos, not even an explicit nil
### GetCompartimentosProduto

`func (o *GvmDTO) GetCompartimentosProduto() []ProdutoDTO`

GetCompartimentosProduto returns the CompartimentosProduto field if non-nil, zero value otherwise.

### GetCompartimentosProdutoOk

`func (o *GvmDTO) GetCompartimentosProdutoOk() (*[]ProdutoDTO, bool)`

GetCompartimentosProdutoOk returns a tuple with the CompartimentosProduto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentosProduto

`func (o *GvmDTO) SetCompartimentosProduto(v []ProdutoDTO)`

SetCompartimentosProduto sets CompartimentosProduto field to given value.

### HasCompartimentosProduto

`func (o *GvmDTO) HasCompartimentosProduto() bool`

HasCompartimentosProduto returns a boolean if a field has been set.

### SetCompartimentosProdutoNil

`func (o *GvmDTO) SetCompartimentosProdutoNil(b bool)`

 SetCompartimentosProdutoNil sets the value for CompartimentosProduto to be an explicit nil

### UnsetCompartimentosProduto
`func (o *GvmDTO) UnsetCompartimentosProduto()`

UnsetCompartimentosProduto ensures that no value is present for CompartimentosProduto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


