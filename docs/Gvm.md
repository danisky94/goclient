# Gvm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Nome** | Pointer to **NullableString** |  | [optional] 
**Modelo** | Pointer to **NullableString** |  | [optional] 
**NumeroSerie** | Pointer to **NullableString** |  | [optional] 
**Tipo** | Pointer to [**TipoGvm**](TipoGvm.md) |  | [optional] 
**Estado** | Pointer to [**EstadoGvm**](EstadoGvm.md) |  | [optional] 
**DataProducao** | Pointer to **NullableTime** |  | [optional] 
**VersaoFirmware** | Pointer to **NullableString** |  | [optional] 
**VersaoSoftware** | Pointer to **NullableString** |  | [optional] 
**UltimaComunicacao** | Pointer to **NullableTime** |  | [optional] 
**RefLocal** | Pointer to **NullableInt64** |  | [optional] 
**RefEntidade** | Pointer to **NullableInt64** |  | [optional] 
**RefNo** | Pointer to **NullableInt64** |  | [optional] 
**Latitude** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableString** |  | [optional] 
**TipoRede** | Pointer to [**TipoRede**](TipoRede.md) |  | [optional] 
**IntensidadeSinal** | Pointer to **NullableFloat64** |  | [optional] 
**Imie** | Pointer to **NullableString** |  | [optional] 
**Imsi** | Pointer to **NullableString** |  | [optional] 
**Iccid** | Pointer to **NullableString** |  | [optional] 
**RaspberryPi** | Pointer to [**RaspberryPi**](RaspberryPi.md) |  | [optional] 
**Local** | Pointer to [**Local**](Local.md) |  | [optional] 
**Compartimentos** | Pointer to [**[]Compartimento**](Compartimento.md) |  | [optional] 
**Vendas** | Pointer to [**[]Venda**](Venda.md) |  | [optional] 
**OrdensServico** | Pointer to [**[]OrdemServico**](OrdemServico.md) |  | [optional] 
**Eventos** | Pointer to [**[]Evento**](Evento.md) |  | [optional] 

## Methods

### NewGvm

`func NewGvm() *Gvm`

NewGvm instantiates a new Gvm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGvmWithDefaults

`func NewGvmWithDefaults() *Gvm`

NewGvmWithDefaults instantiates a new Gvm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gvm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gvm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gvm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Gvm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Gvm) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Gvm) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Gvm) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Gvm) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Gvm) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Gvm) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Gvm) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Gvm) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Gvm) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Gvm) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Gvm) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Gvm) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Gvm) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Gvm) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Gvm) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Gvm) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetNome

`func (o *Gvm) GetNome() string`

GetNome returns the Nome field if non-nil, zero value otherwise.

### GetNomeOk

`func (o *Gvm) GetNomeOk() (*string, bool)`

GetNomeOk returns a tuple with the Nome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNome

`func (o *Gvm) SetNome(v string)`

SetNome sets Nome field to given value.

### HasNome

`func (o *Gvm) HasNome() bool`

HasNome returns a boolean if a field has been set.

### SetNomeNil

`func (o *Gvm) SetNomeNil(b bool)`

 SetNomeNil sets the value for Nome to be an explicit nil

### UnsetNome
`func (o *Gvm) UnsetNome()`

UnsetNome ensures that no value is present for Nome, not even an explicit nil
### GetModelo

`func (o *Gvm) GetModelo() string`

GetModelo returns the Modelo field if non-nil, zero value otherwise.

### GetModeloOk

`func (o *Gvm) GetModeloOk() (*string, bool)`

GetModeloOk returns a tuple with the Modelo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelo

`func (o *Gvm) SetModelo(v string)`

SetModelo sets Modelo field to given value.

### HasModelo

`func (o *Gvm) HasModelo() bool`

HasModelo returns a boolean if a field has been set.

### SetModeloNil

`func (o *Gvm) SetModeloNil(b bool)`

 SetModeloNil sets the value for Modelo to be an explicit nil

### UnsetModelo
`func (o *Gvm) UnsetModelo()`

UnsetModelo ensures that no value is present for Modelo, not even an explicit nil
### GetNumeroSerie

`func (o *Gvm) GetNumeroSerie() string`

GetNumeroSerie returns the NumeroSerie field if non-nil, zero value otherwise.

### GetNumeroSerieOk

`func (o *Gvm) GetNumeroSerieOk() (*string, bool)`

GetNumeroSerieOk returns a tuple with the NumeroSerie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeroSerie

`func (o *Gvm) SetNumeroSerie(v string)`

SetNumeroSerie sets NumeroSerie field to given value.

### HasNumeroSerie

`func (o *Gvm) HasNumeroSerie() bool`

HasNumeroSerie returns a boolean if a field has been set.

### SetNumeroSerieNil

`func (o *Gvm) SetNumeroSerieNil(b bool)`

 SetNumeroSerieNil sets the value for NumeroSerie to be an explicit nil

### UnsetNumeroSerie
`func (o *Gvm) UnsetNumeroSerie()`

UnsetNumeroSerie ensures that no value is present for NumeroSerie, not even an explicit nil
### GetTipo

`func (o *Gvm) GetTipo() TipoGvm`

GetTipo returns the Tipo field if non-nil, zero value otherwise.

### GetTipoOk

`func (o *Gvm) GetTipoOk() (*TipoGvm, bool)`

GetTipoOk returns a tuple with the Tipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipo

`func (o *Gvm) SetTipo(v TipoGvm)`

SetTipo sets Tipo field to given value.

### HasTipo

`func (o *Gvm) HasTipo() bool`

HasTipo returns a boolean if a field has been set.

### GetEstado

`func (o *Gvm) GetEstado() EstadoGvm`

GetEstado returns the Estado field if non-nil, zero value otherwise.

### GetEstadoOk

`func (o *Gvm) GetEstadoOk() (*EstadoGvm, bool)`

GetEstadoOk returns a tuple with the Estado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstado

`func (o *Gvm) SetEstado(v EstadoGvm)`

SetEstado sets Estado field to given value.

### HasEstado

`func (o *Gvm) HasEstado() bool`

HasEstado returns a boolean if a field has been set.

### GetDataProducao

`func (o *Gvm) GetDataProducao() time.Time`

GetDataProducao returns the DataProducao field if non-nil, zero value otherwise.

### GetDataProducaoOk

`func (o *Gvm) GetDataProducaoOk() (*time.Time, bool)`

GetDataProducaoOk returns a tuple with the DataProducao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProducao

`func (o *Gvm) SetDataProducao(v time.Time)`

SetDataProducao sets DataProducao field to given value.

### HasDataProducao

`func (o *Gvm) HasDataProducao() bool`

HasDataProducao returns a boolean if a field has been set.

### SetDataProducaoNil

`func (o *Gvm) SetDataProducaoNil(b bool)`

 SetDataProducaoNil sets the value for DataProducao to be an explicit nil

### UnsetDataProducao
`func (o *Gvm) UnsetDataProducao()`

UnsetDataProducao ensures that no value is present for DataProducao, not even an explicit nil
### GetVersaoFirmware

`func (o *Gvm) GetVersaoFirmware() string`

GetVersaoFirmware returns the VersaoFirmware field if non-nil, zero value otherwise.

### GetVersaoFirmwareOk

`func (o *Gvm) GetVersaoFirmwareOk() (*string, bool)`

GetVersaoFirmwareOk returns a tuple with the VersaoFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoFirmware

`func (o *Gvm) SetVersaoFirmware(v string)`

SetVersaoFirmware sets VersaoFirmware field to given value.

### HasVersaoFirmware

`func (o *Gvm) HasVersaoFirmware() bool`

HasVersaoFirmware returns a boolean if a field has been set.

### SetVersaoFirmwareNil

`func (o *Gvm) SetVersaoFirmwareNil(b bool)`

 SetVersaoFirmwareNil sets the value for VersaoFirmware to be an explicit nil

### UnsetVersaoFirmware
`func (o *Gvm) UnsetVersaoFirmware()`

UnsetVersaoFirmware ensures that no value is present for VersaoFirmware, not even an explicit nil
### GetVersaoSoftware

`func (o *Gvm) GetVersaoSoftware() string`

GetVersaoSoftware returns the VersaoSoftware field if non-nil, zero value otherwise.

### GetVersaoSoftwareOk

`func (o *Gvm) GetVersaoSoftwareOk() (*string, bool)`

GetVersaoSoftwareOk returns a tuple with the VersaoSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersaoSoftware

`func (o *Gvm) SetVersaoSoftware(v string)`

SetVersaoSoftware sets VersaoSoftware field to given value.

### HasVersaoSoftware

`func (o *Gvm) HasVersaoSoftware() bool`

HasVersaoSoftware returns a boolean if a field has been set.

### SetVersaoSoftwareNil

`func (o *Gvm) SetVersaoSoftwareNil(b bool)`

 SetVersaoSoftwareNil sets the value for VersaoSoftware to be an explicit nil

### UnsetVersaoSoftware
`func (o *Gvm) UnsetVersaoSoftware()`

UnsetVersaoSoftware ensures that no value is present for VersaoSoftware, not even an explicit nil
### GetUltimaComunicacao

`func (o *Gvm) GetUltimaComunicacao() time.Time`

GetUltimaComunicacao returns the UltimaComunicacao field if non-nil, zero value otherwise.

### GetUltimaComunicacaoOk

`func (o *Gvm) GetUltimaComunicacaoOk() (*time.Time, bool)`

GetUltimaComunicacaoOk returns a tuple with the UltimaComunicacao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimaComunicacao

`func (o *Gvm) SetUltimaComunicacao(v time.Time)`

SetUltimaComunicacao sets UltimaComunicacao field to given value.

### HasUltimaComunicacao

`func (o *Gvm) HasUltimaComunicacao() bool`

HasUltimaComunicacao returns a boolean if a field has been set.

### SetUltimaComunicacaoNil

`func (o *Gvm) SetUltimaComunicacaoNil(b bool)`

 SetUltimaComunicacaoNil sets the value for UltimaComunicacao to be an explicit nil

### UnsetUltimaComunicacao
`func (o *Gvm) UnsetUltimaComunicacao()`

UnsetUltimaComunicacao ensures that no value is present for UltimaComunicacao, not even an explicit nil
### GetRefLocal

`func (o *Gvm) GetRefLocal() int64`

GetRefLocal returns the RefLocal field if non-nil, zero value otherwise.

### GetRefLocalOk

`func (o *Gvm) GetRefLocalOk() (*int64, bool)`

GetRefLocalOk returns a tuple with the RefLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefLocal

`func (o *Gvm) SetRefLocal(v int64)`

SetRefLocal sets RefLocal field to given value.

### HasRefLocal

`func (o *Gvm) HasRefLocal() bool`

HasRefLocal returns a boolean if a field has been set.

### SetRefLocalNil

`func (o *Gvm) SetRefLocalNil(b bool)`

 SetRefLocalNil sets the value for RefLocal to be an explicit nil

### UnsetRefLocal
`func (o *Gvm) UnsetRefLocal()`

UnsetRefLocal ensures that no value is present for RefLocal, not even an explicit nil
### GetRefEntidade

`func (o *Gvm) GetRefEntidade() int64`

GetRefEntidade returns the RefEntidade field if non-nil, zero value otherwise.

### GetRefEntidadeOk

`func (o *Gvm) GetRefEntidadeOk() (*int64, bool)`

GetRefEntidadeOk returns a tuple with the RefEntidade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefEntidade

`func (o *Gvm) SetRefEntidade(v int64)`

SetRefEntidade sets RefEntidade field to given value.

### HasRefEntidade

`func (o *Gvm) HasRefEntidade() bool`

HasRefEntidade returns a boolean if a field has been set.

### SetRefEntidadeNil

`func (o *Gvm) SetRefEntidadeNil(b bool)`

 SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil

### UnsetRefEntidade
`func (o *Gvm) UnsetRefEntidade()`

UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
### GetRefNo

`func (o *Gvm) GetRefNo() int64`

GetRefNo returns the RefNo field if non-nil, zero value otherwise.

### GetRefNoOk

`func (o *Gvm) GetRefNoOk() (*int64, bool)`

GetRefNoOk returns a tuple with the RefNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNo

`func (o *Gvm) SetRefNo(v int64)`

SetRefNo sets RefNo field to given value.

### HasRefNo

`func (o *Gvm) HasRefNo() bool`

HasRefNo returns a boolean if a field has been set.

### SetRefNoNil

`func (o *Gvm) SetRefNoNil(b bool)`

 SetRefNoNil sets the value for RefNo to be an explicit nil

### UnsetRefNo
`func (o *Gvm) UnsetRefNo()`

UnsetRefNo ensures that no value is present for RefNo, not even an explicit nil
### GetLatitude

`func (o *Gvm) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Gvm) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Gvm) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Gvm) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Gvm) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Gvm) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Gvm) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Gvm) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Gvm) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Gvm) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Gvm) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Gvm) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetTipoRede

`func (o *Gvm) GetTipoRede() TipoRede`

GetTipoRede returns the TipoRede field if non-nil, zero value otherwise.

### GetTipoRedeOk

`func (o *Gvm) GetTipoRedeOk() (*TipoRede, bool)`

GetTipoRedeOk returns a tuple with the TipoRede field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoRede

`func (o *Gvm) SetTipoRede(v TipoRede)`

SetTipoRede sets TipoRede field to given value.

### HasTipoRede

`func (o *Gvm) HasTipoRede() bool`

HasTipoRede returns a boolean if a field has been set.

### GetIntensidadeSinal

`func (o *Gvm) GetIntensidadeSinal() float64`

GetIntensidadeSinal returns the IntensidadeSinal field if non-nil, zero value otherwise.

### GetIntensidadeSinalOk

`func (o *Gvm) GetIntensidadeSinalOk() (*float64, bool)`

GetIntensidadeSinalOk returns a tuple with the IntensidadeSinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensidadeSinal

`func (o *Gvm) SetIntensidadeSinal(v float64)`

SetIntensidadeSinal sets IntensidadeSinal field to given value.

### HasIntensidadeSinal

`func (o *Gvm) HasIntensidadeSinal() bool`

HasIntensidadeSinal returns a boolean if a field has been set.

### SetIntensidadeSinalNil

`func (o *Gvm) SetIntensidadeSinalNil(b bool)`

 SetIntensidadeSinalNil sets the value for IntensidadeSinal to be an explicit nil

### UnsetIntensidadeSinal
`func (o *Gvm) UnsetIntensidadeSinal()`

UnsetIntensidadeSinal ensures that no value is present for IntensidadeSinal, not even an explicit nil
### GetImie

`func (o *Gvm) GetImie() string`

GetImie returns the Imie field if non-nil, zero value otherwise.

### GetImieOk

`func (o *Gvm) GetImieOk() (*string, bool)`

GetImieOk returns a tuple with the Imie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImie

`func (o *Gvm) SetImie(v string)`

SetImie sets Imie field to given value.

### HasImie

`func (o *Gvm) HasImie() bool`

HasImie returns a boolean if a field has been set.

### SetImieNil

`func (o *Gvm) SetImieNil(b bool)`

 SetImieNil sets the value for Imie to be an explicit nil

### UnsetImie
`func (o *Gvm) UnsetImie()`

UnsetImie ensures that no value is present for Imie, not even an explicit nil
### GetImsi

`func (o *Gvm) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *Gvm) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *Gvm) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *Gvm) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### SetImsiNil

`func (o *Gvm) SetImsiNil(b bool)`

 SetImsiNil sets the value for Imsi to be an explicit nil

### UnsetImsi
`func (o *Gvm) UnsetImsi()`

UnsetImsi ensures that no value is present for Imsi, not even an explicit nil
### GetIccid

`func (o *Gvm) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *Gvm) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *Gvm) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *Gvm) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### SetIccidNil

`func (o *Gvm) SetIccidNil(b bool)`

 SetIccidNil sets the value for Iccid to be an explicit nil

### UnsetIccid
`func (o *Gvm) UnsetIccid()`

UnsetIccid ensures that no value is present for Iccid, not even an explicit nil
### GetRaspberryPi

`func (o *Gvm) GetRaspberryPi() RaspberryPi`

GetRaspberryPi returns the RaspberryPi field if non-nil, zero value otherwise.

### GetRaspberryPiOk

`func (o *Gvm) GetRaspberryPiOk() (*RaspberryPi, bool)`

GetRaspberryPiOk returns a tuple with the RaspberryPi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaspberryPi

`func (o *Gvm) SetRaspberryPi(v RaspberryPi)`

SetRaspberryPi sets RaspberryPi field to given value.

### HasRaspberryPi

`func (o *Gvm) HasRaspberryPi() bool`

HasRaspberryPi returns a boolean if a field has been set.

### GetLocal

`func (o *Gvm) GetLocal() Local`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *Gvm) GetLocalOk() (*Local, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *Gvm) SetLocal(v Local)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *Gvm) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetCompartimentos

`func (o *Gvm) GetCompartimentos() []Compartimento`

GetCompartimentos returns the Compartimentos field if non-nil, zero value otherwise.

### GetCompartimentosOk

`func (o *Gvm) GetCompartimentosOk() (*[]Compartimento, bool)`

GetCompartimentosOk returns a tuple with the Compartimentos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimentos

`func (o *Gvm) SetCompartimentos(v []Compartimento)`

SetCompartimentos sets Compartimentos field to given value.

### HasCompartimentos

`func (o *Gvm) HasCompartimentos() bool`

HasCompartimentos returns a boolean if a field has been set.

### SetCompartimentosNil

`func (o *Gvm) SetCompartimentosNil(b bool)`

 SetCompartimentosNil sets the value for Compartimentos to be an explicit nil

### UnsetCompartimentos
`func (o *Gvm) UnsetCompartimentos()`

UnsetCompartimentos ensures that no value is present for Compartimentos, not even an explicit nil
### GetVendas

`func (o *Gvm) GetVendas() []Venda`

GetVendas returns the Vendas field if non-nil, zero value otherwise.

### GetVendasOk

`func (o *Gvm) GetVendasOk() (*[]Venda, bool)`

GetVendasOk returns a tuple with the Vendas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendas

`func (o *Gvm) SetVendas(v []Venda)`

SetVendas sets Vendas field to given value.

### HasVendas

`func (o *Gvm) HasVendas() bool`

HasVendas returns a boolean if a field has been set.

### SetVendasNil

`func (o *Gvm) SetVendasNil(b bool)`

 SetVendasNil sets the value for Vendas to be an explicit nil

### UnsetVendas
`func (o *Gvm) UnsetVendas()`

UnsetVendas ensures that no value is present for Vendas, not even an explicit nil
### GetOrdensServico

`func (o *Gvm) GetOrdensServico() []OrdemServico`

GetOrdensServico returns the OrdensServico field if non-nil, zero value otherwise.

### GetOrdensServicoOk

`func (o *Gvm) GetOrdensServicoOk() (*[]OrdemServico, bool)`

GetOrdensServicoOk returns a tuple with the OrdensServico field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdensServico

`func (o *Gvm) SetOrdensServico(v []OrdemServico)`

SetOrdensServico sets OrdensServico field to given value.

### HasOrdensServico

`func (o *Gvm) HasOrdensServico() bool`

HasOrdensServico returns a boolean if a field has been set.

### SetOrdensServicoNil

`func (o *Gvm) SetOrdensServicoNil(b bool)`

 SetOrdensServicoNil sets the value for OrdensServico to be an explicit nil

### UnsetOrdensServico
`func (o *Gvm) UnsetOrdensServico()`

UnsetOrdensServico ensures that no value is present for OrdensServico, not even an explicit nil
### GetEventos

`func (o *Gvm) GetEventos() []Evento`

GetEventos returns the Eventos field if non-nil, zero value otherwise.

### GetEventosOk

`func (o *Gvm) GetEventosOk() (*[]Evento, bool)`

GetEventosOk returns a tuple with the Eventos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventos

`func (o *Gvm) SetEventos(v []Evento)`

SetEventos sets Eventos field to given value.

### HasEventos

`func (o *Gvm) HasEventos() bool`

HasEventos returns a boolean if a field has been set.

### SetEventosNil

`func (o *Gvm) SetEventosNil(b bool)`

 SetEventosNil sets the value for Eventos to be an explicit nil

### UnsetEventos
`func (o *Gvm) UnsetEventos()`

UnsetEventos ensures that no value is present for Eventos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


