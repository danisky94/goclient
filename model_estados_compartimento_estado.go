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

// EstadosCompartimentoEstado the model 'EstadosCompartimentoEstado'
type EstadosCompartimentoEstado int32

// List of EstadosCompartimentoEstado
const (
	_0 EstadosCompartimentoEstado = 0
	_1 EstadosCompartimentoEstado = 1
	_2 EstadosCompartimentoEstado = 2
)

var allowedEstadosCompartimentoEstadoEnumValues = []EstadosCompartimentoEstado{
	0,
	1,
	2,
}

func (v *EstadosCompartimentoEstado) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EstadosCompartimentoEstado(value)
	for _, existing := range allowedEstadosCompartimentoEstadoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EstadosCompartimentoEstado", value)
}

// NewEstadosCompartimentoEstadoFromValue returns a pointer to a valid EstadosCompartimentoEstado
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEstadosCompartimentoEstadoFromValue(v int32) (*EstadosCompartimentoEstado, error) {
	ev := EstadosCompartimentoEstado(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EstadosCompartimentoEstado: valid values are %v", v, allowedEstadosCompartimentoEstadoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EstadosCompartimentoEstado) IsValid() bool {
	for _, existing := range allowedEstadosCompartimentoEstadoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EstadosCompartimentoEstado value
func (v EstadosCompartimentoEstado) Ptr() *EstadosCompartimentoEstado {
	return &v
}

type NullableEstadosCompartimentoEstado struct {
	value *EstadosCompartimentoEstado
	isSet bool
}

func (v NullableEstadosCompartimentoEstado) Get() *EstadosCompartimentoEstado {
	return v.value
}

func (v *NullableEstadosCompartimentoEstado) Set(val *EstadosCompartimentoEstado) {
	v.value = val
	v.isSet = true
}

func (v NullableEstadosCompartimentoEstado) IsSet() bool {
	return v.isSet
}

func (v *NullableEstadosCompartimentoEstado) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstadosCompartimentoEstado(val *EstadosCompartimentoEstado) *NullableEstadosCompartimentoEstado {
	return &NullableEstadosCompartimentoEstado{value: val, isSet: true}
}

func (v NullableEstadosCompartimentoEstado) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstadosCompartimentoEstado) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

