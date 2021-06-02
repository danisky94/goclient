# EventoDTOResultado

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**EventoDTO**](EventoDTO.md) |  | [optional] 

## Methods

### NewEventoDTOResultado

`func NewEventoDTOResultado() *EventoDTOResultado`

NewEventoDTOResultado instantiates a new EventoDTOResultado object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventoDTOResultadoWithDefaults

`func NewEventoDTOResultadoWithDefaults() *EventoDTOResultado`

NewEventoDTOResultadoWithDefaults instantiates a new EventoDTOResultado object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *EventoDTOResultado) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *EventoDTOResultado) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *EventoDTOResultado) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *EventoDTOResultado) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetMessage

`func (o *EventoDTOResultado) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventoDTOResultado) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventoDTOResultado) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventoDTOResultado) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *EventoDTOResultado) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *EventoDTOResultado) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValue

`func (o *EventoDTOResultado) GetValue() EventoDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventoDTOResultado) GetValueOk() (*EventoDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventoDTOResultado) SetValue(v EventoDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventoDTOResultado) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


