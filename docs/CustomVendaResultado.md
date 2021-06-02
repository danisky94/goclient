# CustomVendaResultado

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**CustomVenda**](CustomVenda.md) |  | [optional] 

## Methods

### NewCustomVendaResultado

`func NewCustomVendaResultado() *CustomVendaResultado`

NewCustomVendaResultado instantiates a new CustomVendaResultado object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomVendaResultadoWithDefaults

`func NewCustomVendaResultadoWithDefaults() *CustomVendaResultado`

NewCustomVendaResultadoWithDefaults instantiates a new CustomVendaResultado object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *CustomVendaResultado) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CustomVendaResultado) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CustomVendaResultado) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CustomVendaResultado) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetMessage

`func (o *CustomVendaResultado) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomVendaResultado) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomVendaResultado) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CustomVendaResultado) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CustomVendaResultado) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CustomVendaResultado) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetValue

`func (o *CustomVendaResultado) GetValue() CustomVenda`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomVendaResultado) GetValueOk() (*CustomVenda, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomVendaResultado) SetValue(v CustomVenda)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomVendaResultado) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


