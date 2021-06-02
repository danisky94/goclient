# GvmDTOListResultado

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**[]GvmDTO**](GvmDTO.md) |  | [optional] 

## Methods

### NewGvmDTOListResultado

`func NewGvmDTOListResultado() *GvmDTOListResultado`

NewGvmDTOListResultado instantiates a new GvmDTOListResultado object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGvmDTOListResultadoWithDefaults

`func NewGvmDTOListResultadoWithDefaults() *GvmDTOListResultado`

NewGvmDTOListResultadoWithDefaults instantiates a new GvmDTOListResultado object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *GvmDTOListResultado) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GvmDTOListResultado) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GvmDTOListResultado) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *GvmDTOListResultado) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetMessage

`func (o *GvmDTOListResultado) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GvmDTOListResultado) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GvmDTOListResultado) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GvmDTOListResultado) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GvmDTOListResultado) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GvmDTOListResultado) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValue

`func (o *GvmDTOListResultado) GetValue() []GvmDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GvmDTOListResultado) GetValueOk() (*[]GvmDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GvmDTOListResultado) SetValue(v []GvmDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *GvmDTOListResultado) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *GvmDTOListResultado) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GvmDTOListResultado) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


