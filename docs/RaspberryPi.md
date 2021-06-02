# RaspberryPi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Identificador** | Pointer to **NullableString** |  | [optional] 
**Gvm** | Pointer to [**Gvm**](Gvm.md) |  | [optional] 

## Methods

### NewRaspberryPi

`func NewRaspberryPi() *RaspberryPi`

NewRaspberryPi instantiates a new RaspberryPi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaspberryPiWithDefaults

`func NewRaspberryPiWithDefaults() *RaspberryPi`

NewRaspberryPiWithDefaults instantiates a new RaspberryPi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RaspberryPi) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RaspberryPi) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RaspberryPi) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RaspberryPi) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *RaspberryPi) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *RaspberryPi) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *RaspberryPi) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *RaspberryPi) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *RaspberryPi) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *RaspberryPi) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *RaspberryPi) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *RaspberryPi) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *RaspberryPi) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *RaspberryPi) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *RaspberryPi) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *RaspberryPi) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *RaspberryPi) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *RaspberryPi) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *RaspberryPi) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *RaspberryPi) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetIdentificador

`func (o *RaspberryPi) GetIdentificador() string`

GetIdentificador returns the Identificador field if non-nil, zero value otherwise.

### GetIdentificadorOk

`func (o *RaspberryPi) GetIdentificadorOk() (*string, bool)`

GetIdentificadorOk returns a tuple with the Identificador field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificador

`func (o *RaspberryPi) SetIdentificador(v string)`

SetIdentificador sets Identificador field to given value.

### HasIdentificador

`func (o *RaspberryPi) HasIdentificador() bool`

HasIdentificador returns a boolean if a field has been set.

### SetIdentificadorNil

`func (o *RaspberryPi) SetIdentificadorNil(b bool)`

 SetIdentificadorNil sets the value for Identificador to be an explicit nil

### UnsetIdentificador
`func (o *RaspberryPi) UnsetIdentificador()`

UnsetIdentificador ensures that no value is present for Identificador, not even an explicit nil
### GetGvm

`func (o *RaspberryPi) GetGvm() Gvm`

GetGvm returns the Gvm field if non-nil, zero value otherwise.

### GetGvmOk

`func (o *RaspberryPi) GetGvmOk() (*Gvm, bool)`

GetGvmOk returns a tuple with the Gvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGvm

`func (o *RaspberryPi) SetGvm(v Gvm)`

SetGvm sets Gvm field to given value.

### HasGvm

`func (o *RaspberryPi) HasGvm() bool`

HasGvm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


