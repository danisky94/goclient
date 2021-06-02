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

// TipoUtilizador the model 'TipoUtilizador'
type TipoUtilizador int32


var allowedTipoUtilizadorEnumValues = []TipoUtilizador{
	0,
	1,
	2,
	3,
}

func (v *TipoUtilizador) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TipoUtilizador(value)
	for _, existing := range allowedTipoUtilizadorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TipoUtilizador", value)
}

// NewTipoUtilizadorFromValue returns a pointer to a valid TipoUtilizador
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTipoUtilizadorFromValue(v int32) (*TipoUtilizador, error) {
	ev := TipoUtilizador(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TipoUtilizador: valid values are %v", v, allowedTipoUtilizadorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TipoUtilizador) IsValid() bool {
	for _, existing := range allowedTipoUtilizadorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TipoUtilizador value
func (v TipoUtilizador) Ptr() *TipoUtilizador {
	return &v
}

type NullableTipoUtilizador struct {
	value *TipoUtilizador
	isSet bool
}

func (v NullableTipoUtilizador) Get() *TipoUtilizador {
	return v.value
}

func (v *NullableTipoUtilizador) Set(val *TipoUtilizador) {
	v.value = val
	v.isSet = true
}

func (v NullableTipoUtilizador) IsSet() bool {
	return v.isSet
}

func (v *NullableTipoUtilizador) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTipoUtilizador(val *TipoUtilizador) *NullableTipoUtilizador {
	return &NullableTipoUtilizador{value: val, isSet: true}
}

func (v NullableTipoUtilizador) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTipoUtilizador) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
