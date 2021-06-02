# Led

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Upd** | Pointer to **NullableTime** |  | [optional] 
**Usr** | Pointer to **NullableInt64** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Channel** | Pointer to **NullableInt32** |  | [optional] 
**Start** | Pointer to **NullableInt32** |  | [optional] 
**End** | Pointer to **NullableInt32** |  | [optional] 
**RefCompartimento** | Pointer to **NullableInt64** |  | [optional] 
**Compartimento** | Pointer to [**Compartimento**](Compartimento.md) |  | [optional] 

## Methods

### NewLed

`func NewLed() *Led`

NewLed instantiates a new Led object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedWithDefaults

`func NewLedWithDefaults() *Led`

NewLedWithDefaults instantiates a new Led object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Led) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Led) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Led) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Led) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Led) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Led) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Led) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Led) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpd

`func (o *Led) GetUpd() time.Time`

GetUpd returns the Upd field if non-nil, zero value otherwise.

### GetUpdOk

`func (o *Led) GetUpdOk() (*time.Time, bool)`

GetUpdOk returns a tuple with the Upd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpd

`func (o *Led) SetUpd(v time.Time)`

SetUpd sets Upd field to given value.

### HasUpd

`func (o *Led) HasUpd() bool`

HasUpd returns a boolean if a field has been set.

### SetUpdNil

`func (o *Led) SetUpdNil(b bool)`

 SetUpdNil sets the value for Upd to be an explicit nil

### UnsetUpd
`func (o *Led) UnsetUpd()`

UnsetUpd ensures that no value is present for Upd, not even an explicit nil
### GetUsr

`func (o *Led) GetUsr() int64`

GetUsr returns the Usr field if non-nil, zero value otherwise.

### GetUsrOk

`func (o *Led) GetUsrOk() (*int64, bool)`

GetUsrOk returns a tuple with the Usr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsr

`func (o *Led) SetUsr(v int64)`

SetUsr sets Usr field to given value.

### HasUsr

`func (o *Led) HasUsr() bool`

HasUsr returns a boolean if a field has been set.

### SetUsrNil

`func (o *Led) SetUsrNil(b bool)`

 SetUsrNil sets the value for Usr to be an explicit nil

### UnsetUsr
`func (o *Led) UnsetUsr()`

UnsetUsr ensures that no value is present for Usr, not even an explicit nil
### GetAddress

`func (o *Led) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Led) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Led) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Led) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Led) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Led) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetChannel

`func (o *Led) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Led) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Led) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Led) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *Led) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *Led) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetStart

`func (o *Led) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Led) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Led) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Led) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *Led) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *Led) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *Led) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Led) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Led) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Led) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *Led) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *Led) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetRefCompartimento

`func (o *Led) GetRefCompartimento() int64`

GetRefCompartimento returns the RefCompartimento field if non-nil, zero value otherwise.

### GetRefCompartimentoOk

`func (o *Led) GetRefCompartimentoOk() (*int64, bool)`

GetRefCompartimentoOk returns a tuple with the RefCompartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCompartimento

`func (o *Led) SetRefCompartimento(v int64)`

SetRefCompartimento sets RefCompartimento field to given value.

### HasRefCompartimento

`func (o *Led) HasRefCompartimento() bool`

HasRefCompartimento returns a boolean if a field has been set.

### SetRefCompartimentoNil

`func (o *Led) SetRefCompartimentoNil(b bool)`

 SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil

### UnsetRefCompartimento
`func (o *Led) UnsetRefCompartimento()`

UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
### GetCompartimento

`func (o *Led) GetCompartimento() Compartimento`

GetCompartimento returns the Compartimento field if non-nil, zero value otherwise.

### GetCompartimentoOk

`func (o *Led) GetCompartimentoOk() (*Compartimento, bool)`

GetCompartimentoOk returns a tuple with the Compartimento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartimento

`func (o *Led) SetCompartimento(v Compartimento)`

SetCompartimento sets Compartimento field to given value.

### HasCompartimento

`func (o *Led) HasCompartimento() bool`

HasCompartimento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


