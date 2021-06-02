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
	"fmt"
)

// TipoVendaEstado the model 'TipoVendaEstado'
type TipoVendaEstado int32


var allowedTipoVendaEstadoEnumValues = []TipoVendaEstado{
	0,
	1,
	2,
	3,
}

func (v *TipoVendaEstado) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TipoVendaEstado(value)
	for _, existing := range allowedTipoVendaEstadoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TipoVendaEstado", value)
}

// NewTipoVendaEstadoFromValue returns a pointer to a valid TipoVendaEstado
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTipoVendaEstadoFromValue(v int32) (*TipoVendaEstado, error) {
	ev := TipoVendaEstado(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TipoVendaEstado: valid values are %v", v, allowedTipoVendaEstadoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TipoVendaEstado) IsValid() bool {
	for _, existing := range allowedTipoVendaEstadoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TipoVendaEstado value
func (v TipoVendaEstado) Ptr() *TipoVendaEstado {
	return &v
}

type NullableTipoVendaEstado struct {
	value *TipoVendaEstado
	isSet bool
}

func (v NullableTipoVendaEstado) Get() *TipoVendaEstado {
	return v.value
}

func (v *NullableTipoVendaEstado) Set(val *TipoVendaEstado) {
	v.value = val
	v.isSet = true
}

func (v NullableTipoVendaEstado) IsSet() bool {
	return v.isSet
}

func (v *NullableTipoVendaEstado) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTipoVendaEstado(val *TipoVendaEstado) *NullableTipoVendaEstado {
	return &NullableTipoVendaEstado{value: val, isSet: true}
}

func (v NullableTipoVendaEstado) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTipoVendaEstado) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

