# LedDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Channel** | Pointer to **NullableInt32** |  | [optional] 
**Start** | Pointer to **NullableInt32** |  | [optional] 
**End** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewLedDTO

`func NewLedDTO() *LedDTO`

NewLedDTO instantiates a new LedDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedDTOWithDefaults

`func NewLedDTOWithDefaults() *LedDTO`

NewLedDTOWithDefaults instantiates a new LedDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LedDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LedDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LedDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LedDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *LedDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LedDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LedDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LedDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *LedDTO) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *LedDTO) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetChannel

`func (o *LedDTO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *LedDTO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *LedDTO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *LedDTO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *LedDTO) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *LedDTO) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetStart

`func (o *LedDTO) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LedDTO) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LedDTO) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *LedDTO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *LedDTO) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LedDTO) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *LedDTO) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LedDTO) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LedDTO) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LedDTO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *LedDTO) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *LedDTO) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


