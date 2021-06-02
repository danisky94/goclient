/*
 * Petrolifera API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VendaPetroliferaListResultado struct for VendaPetroliferaListResultado
type VendaPetroliferaListResultado struct {
	Ok *bool `json:"ok,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Value []VendaPetrolifera `json:"value,omitempty"`
}

// NewVendaPetroliferaListResultado instantiates a new VendaPetroliferaListResultado object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendaPetroliferaListResultado() *VendaPetroliferaListResultado {
	this := VendaPetroliferaListResultado{}
	return &this
}

// NewVendaPetroliferaListResultadoWithDefaults instantiates a new VendaPetroliferaListResultado object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendaPetroliferaListResultadoWithDefaults() *VendaPetroliferaListResultado {
	this := VendaPetroliferaListResultado{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *VendaPetroliferaListResultado) GetOk() bool {
	if o == nil || o.Ok == nil {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendaPetroliferaListResultado) GetOkOk() (*bool, bool) {
	if o == nil || o.Ok == nil {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *VendaPetroliferaListResultado) HasOk() bool {
	if o != nil && o.Ok != nil {
		return true
	}

	return false
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *VendaPetroliferaListResultado) SetOk(v bool) {
	o.Ok = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendaPetroliferaListResultado) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendaPetroliferaListResultado) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *VendaPetroliferaListResultado) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *VendaPetroliferaListResultado) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *VendaPetroliferaListResultado) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *VendaPetroliferaListResultado) UnsetMessage() {
	o.Message.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendaPetroliferaListResultado) GetValue() []VendaPetrolifera {
	if o == nil  {
		var ret []VendaPetrolifera
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendaPetroliferaListResultado) GetValueOk() (*[]VendaPetrolifera, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VendaPetroliferaListResultado) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []VendaPetrolifera and assigns it to the Value field.
func (o *VendaPetroliferaListResultado) SetValue(v []VendaPetrolifera) {
	o.Value = v
}

func (o VendaPetroliferaListResultado) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ok != nil {
		toSerialize["ok"] = o.Ok
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVendaPetroliferaListResultado struct {
	value *VendaPetroliferaListResultado
	isSet bool
}

func (v NullableVendaPetroliferaListResultado) Get() *VendaPetroliferaListResultado {
	return v.value
}

func (v *NullableVendaPetroliferaListResultado) Set(val *VendaPetroliferaListResultado) {
	v.value = val
	v.isSet = true
}

func (v NullableVendaPetroliferaListResultado) IsSet() bool {
	return v.isSet
}

func (v *NullableVendaPetroliferaListResultado) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendaPetroliferaListResultado(val *VendaPetroliferaListResultado) *NullableVendaPetroliferaListResultado {
	return &NullableVendaPetroliferaListResultado{value: val, isSet: true}
}

func (v NullableVendaPetroliferaListResultado) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendaPetroliferaListResultado) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

