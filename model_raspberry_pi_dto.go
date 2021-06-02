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

// RaspberryPiDTO struct for RaspberryPiDTO
type RaspberryPiDTO struct {
	Id *int64 `json:"id,omitempty"`
	Identificador NullableString `json:"identificador,omitempty"`
	RefGvm NullableInt64 `json:"refGvm,omitempty"`
}

// NewRaspberryPiDTO instantiates a new RaspberryPiDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRaspberryPiDTO() *RaspberryPiDTO {
	this := RaspberryPiDTO{}
	return &this
}

// NewRaspberryPiDTOWithDefaults instantiates a new RaspberryPiDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRaspberryPiDTOWithDefaults() *RaspberryPiDTO {
	this := RaspberryPiDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RaspberryPiDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaspberryPiDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RaspberryPiDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RaspberryPiDTO) SetId(v int64) {
	o.Id = &v
}

// GetIdentificador returns the Identificador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaspberryPiDTO) GetIdentificador() string {
	if o == nil || o.Identificador.Get() == nil {
		var ret string
		return ret
	}
	return *o.Identificador.Get()
}

// GetIdentificadorOk returns a tuple with the Identificador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaspberryPiDTO) GetIdentificadorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identificador.Get(), o.Identificador.IsSet()
}

// HasIdentificador returns a boolean if a field has been set.
func (o *RaspberryPiDTO) HasIdentificador() bool {
	if o != nil && o.Identificador.IsSet() {
		return true
	}

	return false
}

// SetIdentificador gets a reference to the given NullableString and assigns it to the Identificador field.
func (o *RaspberryPiDTO) SetIdentificador(v string) {
	o.Identificador.Set(&v)
}
// SetIdentificadorNil sets the value for Identificador to be an explicit nil
func (o *RaspberryPiDTO) SetIdentificadorNil() {
	o.Identificador.Set(nil)
}

// UnsetIdentificador ensures that no value is present for Identificador, not even an explicit nil
func (o *RaspberryPiDTO) UnsetIdentificador() {
	o.Identificador.Unset()
}

// GetRefGvm returns the RefGvm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaspberryPiDTO) GetRefGvm() int64 {
	if o == nil || o.RefGvm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefGvm.Get()
}

// GetRefGvmOk returns a tuple with the RefGvm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaspberryPiDTO) GetRefGvmOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefGvm.Get(), o.RefGvm.IsSet()
}

// HasRefGvm returns a boolean if a field has been set.
func (o *RaspberryPiDTO) HasRefGvm() bool {
	if o != nil && o.RefGvm.IsSet() {
		return true
	}

	return false
}

// SetRefGvm gets a reference to the given NullableInt64 and assigns it to the RefGvm field.
func (o *RaspberryPiDTO) SetRefGvm(v int64) {
	o.RefGvm.Set(&v)
}
// SetRefGvmNil sets the value for RefGvm to be an explicit nil
func (o *RaspberryPiDTO) SetRefGvmNil() {
	o.RefGvm.Set(nil)
}

// UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
func (o *RaspberryPiDTO) UnsetRefGvm() {
	o.RefGvm.Unset()
}

func (o RaspberryPiDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Identificador.IsSet() {
		toSerialize["identificador"] = o.Identificador.Get()
	}
	if o.RefGvm.IsSet() {
		toSerialize["refGvm"] = o.RefGvm.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRaspberryPiDTO struct {
	value *RaspberryPiDTO
	isSet bool
}

func (v NullableRaspberryPiDTO) Get() *RaspberryPiDTO {
	return v.value
}

func (v *NullableRaspberryPiDTO) Set(val *RaspberryPiDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRaspberryPiDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRaspberryPiDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaspberryPiDTO(val *RaspberryPiDTO) *NullableRaspberryPiDTO {
	return &NullableRaspberryPiDTO{value: val, isSet: true}
}

func (v NullableRaspberryPiDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaspberryPiDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


