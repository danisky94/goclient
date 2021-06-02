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

// CompartimentoDTO struct for CompartimentoDTO
type CompartimentoDTO struct {
	Id *int64 `json:"id,omitempty"`
	Numero NullableString `json:"numero,omitempty"`
	Tipo *TipoCompartimento `json:"tipo,omitempty"`
	Ativo *bool `json:"ativo,omitempty"`
	Fechaduras []FechaduraDTO `json:"fechaduras,omitempty"`
	Leds []LedDTO `json:"leds,omitempty"`
}

// NewCompartimentoDTO instantiates a new CompartimentoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompartimentoDTO() *CompartimentoDTO {
	this := CompartimentoDTO{}
	return &this
}

// NewCompartimentoDTOWithDefaults instantiates a new CompartimentoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompartimentoDTOWithDefaults() *CompartimentoDTO {
	this := CompartimentoDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompartimentoDTO) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoDTO) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CompartimentoDTO) SetId(v int64) {
	o.Id = &v
}

// GetNumero returns the Numero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoDTO) GetNumero() string {
	if o == nil || o.Numero.Get() == nil {
		var ret string
		return ret
	}
	return *o.Numero.Get()
}

// GetNumeroOk returns a tuple with the Numero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoDTO) GetNumeroOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Numero.Get(), o.Numero.IsSet()
}

// HasNumero returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasNumero() bool {
	if o != nil && o.Numero.IsSet() {
		return true
	}

	return false
}

// SetNumero gets a reference to the given NullableString and assigns it to the Numero field.
func (o *CompartimentoDTO) SetNumero(v string) {
	o.Numero.Set(&v)
}
// SetNumeroNil sets the value for Numero to be an explicit nil
func (o *CompartimentoDTO) SetNumeroNil() {
	o.Numero.Set(nil)
}

// UnsetNumero ensures that no value is present for Numero, not even an explicit nil
func (o *CompartimentoDTO) UnsetNumero() {
	o.Numero.Unset()
}

// GetTipo returns the Tipo field value if set, zero value otherwise.
func (o *CompartimentoDTO) GetTipo() TipoCompartimento {
	if o == nil || o.Tipo == nil {
		var ret TipoCompartimento
		return ret
	}
	return *o.Tipo
}

// GetTipoOk returns a tuple with the Tipo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoDTO) GetTipoOk() (*TipoCompartimento, bool) {
	if o == nil || o.Tipo == nil {
		return nil, false
	}
	return o.Tipo, true
}

// HasTipo returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasTipo() bool {
	if o != nil && o.Tipo != nil {
		return true
	}

	return false
}

// SetTipo gets a reference to the given TipoCompartimento and assigns it to the Tipo field.
func (o *CompartimentoDTO) SetTipo(v TipoCompartimento) {
	o.Tipo = &v
}

// GetAtivo returns the Ativo field value if set, zero value otherwise.
func (o *CompartimentoDTO) GetAtivo() bool {
	if o == nil || o.Ativo == nil {
		var ret bool
		return ret
	}
	return *o.Ativo
}

// GetAtivoOk returns a tuple with the Ativo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoDTO) GetAtivoOk() (*bool, bool) {
	if o == nil || o.Ativo == nil {
		return nil, false
	}
	return o.Ativo, true
}

// HasAtivo returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasAtivo() bool {
	if o != nil && o.Ativo != nil {
		return true
	}

	return false
}

// SetAtivo gets a reference to the given bool and assigns it to the Ativo field.
func (o *CompartimentoDTO) SetAtivo(v bool) {
	o.Ativo = &v
}

// GetFechaduras returns the Fechaduras field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoDTO) GetFechaduras() []FechaduraDTO {
	if o == nil  {
		var ret []FechaduraDTO
		return ret
	}
	return o.Fechaduras
}

// GetFechadurasOk returns a tuple with the Fechaduras field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoDTO) GetFechadurasOk() (*[]FechaduraDTO, bool) {
	if o == nil || o.Fechaduras == nil {
		return nil, false
	}
	return &o.Fechaduras, true
}

// HasFechaduras returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasFechaduras() bool {
	if o != nil && o.Fechaduras != nil {
		return true
	}

	return false
}

// SetFechaduras gets a reference to the given []FechaduraDTO and assigns it to the Fechaduras field.
func (o *CompartimentoDTO) SetFechaduras(v []FechaduraDTO) {
	o.Fechaduras = v
}

// GetLeds returns the Leds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoDTO) GetLeds() []LedDTO {
	if o == nil  {
		var ret []LedDTO
		return ret
	}
	return o.Leds
}

// GetLedsOk returns a tuple with the Leds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoDTO) GetLedsOk() (*[]LedDTO, bool) {
	if o == nil || o.Leds == nil {
		return nil, false
	}
	return &o.Leds, true
}

// HasLeds returns a boolean if a field has been set.
func (o *CompartimentoDTO) HasLeds() bool {
	if o != nil && o.Leds != nil {
		return true
	}

	return false
}

// SetLeds gets a reference to the given []LedDTO and assigns it to the Leds field.
func (o *CompartimentoDTO) SetLeds(v []LedDTO) {
	o.Leds = v
}

func (o CompartimentoDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Numero.IsSet() {
		toSerialize["numero"] = o.Numero.Get()
	}
	if o.Tipo != nil {
		toSerialize["tipo"] = o.Tipo
	}
	if o.Ativo != nil {
		toSerialize["ativo"] = o.Ativo
	}
	if o.Fechaduras != nil {
		toSerialize["fechaduras"] = o.Fechaduras
	}
	if o.Leds != nil {
		toSerialize["leds"] = o.Leds
	}
	return json.Marshal(toSerialize)
}

type NullableCompartimentoDTO struct {
	value *CompartimentoDTO
	isSet bool
}

func (v NullableCompartimentoDTO) Get() *CompartimentoDTO {
	return v.value
}

func (v *NullableCompartimentoDTO) Set(val *CompartimentoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCompartimentoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCompartimentoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompartimentoDTO(val *CompartimentoDTO) *NullableCompartimentoDTO {
	return &NullableCompartimentoDTO{value: val, isSet: true}
}

func (v NullableCompartimentoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompartimentoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


