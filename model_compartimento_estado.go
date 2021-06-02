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
	"time"
)

// CompartimentoEstado struct for CompartimentoEstado
type CompartimentoEstado struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefProduto NullableInt64 `json:"refProduto,omitempty"`
	Estado *EstadosCompartimentoEstado `json:"estado,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	Ocupado *bool `json:"ocupado,omitempty"`
	RefCompartimento NullableInt64 `json:"refCompartimento,omitempty"`
	Vasilhame *bool `json:"vasilhame,omitempty"`
	Compartimento *Compartimento `json:"compartimento,omitempty"`
}

// NewCompartimentoEstado instantiates a new CompartimentoEstado object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompartimentoEstado() *CompartimentoEstado {
	this := CompartimentoEstado{}
	return &this
}

// NewCompartimentoEstadoWithDefaults instantiates a new CompartimentoEstado object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompartimentoEstadoWithDefaults() *CompartimentoEstado {
	this := CompartimentoEstado{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CompartimentoEstado) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *CompartimentoEstado) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoEstado) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoEstado) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *CompartimentoEstado) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *CompartimentoEstado) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *CompartimentoEstado) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoEstado) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoEstado) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *CompartimentoEstado) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *CompartimentoEstado) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *CompartimentoEstado) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefProduto returns the RefProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoEstado) GetRefProduto() int64 {
	if o == nil || o.RefProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefProduto.Get()
}

// GetRefProdutoOk returns a tuple with the RefProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoEstado) GetRefProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefProduto.Get(), o.RefProduto.IsSet()
}

// HasRefProduto returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasRefProduto() bool {
	if o != nil && o.RefProduto.IsSet() {
		return true
	}

	return false
}

// SetRefProduto gets a reference to the given NullableInt64 and assigns it to the RefProduto field.
func (o *CompartimentoEstado) SetRefProduto(v int64) {
	o.RefProduto.Set(&v)
}
// SetRefProdutoNil sets the value for RefProduto to be an explicit nil
func (o *CompartimentoEstado) SetRefProdutoNil() {
	o.RefProduto.Set(nil)
}

// UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
func (o *CompartimentoEstado) UnsetRefProduto() {
	o.RefProduto.Unset()
}

// GetEstado returns the Estado field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetEstado() EstadosCompartimentoEstado {
	if o == nil || o.Estado == nil {
		var ret EstadosCompartimentoEstado
		return ret
	}
	return *o.Estado
}

// GetEstadoOk returns a tuple with the Estado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetEstadoOk() (*EstadosCompartimentoEstado, bool) {
	if o == nil || o.Estado == nil {
		return nil, false
	}
	return o.Estado, true
}

// HasEstado returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasEstado() bool {
	if o != nil && o.Estado != nil {
		return true
	}

	return false
}

// SetEstado gets a reference to the given EstadosCompartimentoEstado and assigns it to the Estado field.
func (o *CompartimentoEstado) SetEstado(v EstadosCompartimentoEstado) {
	o.Estado = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoEstado) GetData() time.Time {
	if o == nil || o.Data.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoEstado) GetDataOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *CompartimentoEstado) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *CompartimentoEstado) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *CompartimentoEstado) UnsetData() {
	o.Data.Unset()
}

// GetOcupado returns the Ocupado field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetOcupado() bool {
	if o == nil || o.Ocupado == nil {
		var ret bool
		return ret
	}
	return *o.Ocupado
}

// GetOcupadoOk returns a tuple with the Ocupado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetOcupadoOk() (*bool, bool) {
	if o == nil || o.Ocupado == nil {
		return nil, false
	}
	return o.Ocupado, true
}

// HasOcupado returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasOcupado() bool {
	if o != nil && o.Ocupado != nil {
		return true
	}

	return false
}

// SetOcupado gets a reference to the given bool and assigns it to the Ocupado field.
func (o *CompartimentoEstado) SetOcupado(v bool) {
	o.Ocupado = &v
}

// GetRefCompartimento returns the RefCompartimento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompartimentoEstado) GetRefCompartimento() int64 {
	if o == nil || o.RefCompartimento.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefCompartimento.Get()
}

// GetRefCompartimentoOk returns a tuple with the RefCompartimento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompartimentoEstado) GetRefCompartimentoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefCompartimento.Get(), o.RefCompartimento.IsSet()
}

// HasRefCompartimento returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasRefCompartimento() bool {
	if o != nil && o.RefCompartimento.IsSet() {
		return true
	}

	return false
}

// SetRefCompartimento gets a reference to the given NullableInt64 and assigns it to the RefCompartimento field.
func (o *CompartimentoEstado) SetRefCompartimento(v int64) {
	o.RefCompartimento.Set(&v)
}
// SetRefCompartimentoNil sets the value for RefCompartimento to be an explicit nil
func (o *CompartimentoEstado) SetRefCompartimentoNil() {
	o.RefCompartimento.Set(nil)
}

// UnsetRefCompartimento ensures that no value is present for RefCompartimento, not even an explicit nil
func (o *CompartimentoEstado) UnsetRefCompartimento() {
	o.RefCompartimento.Unset()
}

// GetVasilhame returns the Vasilhame field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetVasilhame() bool {
	if o == nil || o.Vasilhame == nil {
		var ret bool
		return ret
	}
	return *o.Vasilhame
}

// GetVasilhameOk returns a tuple with the Vasilhame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetVasilhameOk() (*bool, bool) {
	if o == nil || o.Vasilhame == nil {
		return nil, false
	}
	return o.Vasilhame, true
}

// HasVasilhame returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasVasilhame() bool {
	if o != nil && o.Vasilhame != nil {
		return true
	}

	return false
}

// SetVasilhame gets a reference to the given bool and assigns it to the Vasilhame field.
func (o *CompartimentoEstado) SetVasilhame(v bool) {
	o.Vasilhame = &v
}

// GetCompartimento returns the Compartimento field value if set, zero value otherwise.
func (o *CompartimentoEstado) GetCompartimento() Compartimento {
	if o == nil || o.Compartimento == nil {
		var ret Compartimento
		return ret
	}
	return *o.Compartimento
}

// GetCompartimentoOk returns a tuple with the Compartimento field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartimentoEstado) GetCompartimentoOk() (*Compartimento, bool) {
	if o == nil || o.Compartimento == nil {
		return nil, false
	}
	return o.Compartimento, true
}

// HasCompartimento returns a boolean if a field has been set.
func (o *CompartimentoEstado) HasCompartimento() bool {
	if o != nil && o.Compartimento != nil {
		return true
	}

	return false
}

// SetCompartimento gets a reference to the given Compartimento and assigns it to the Compartimento field.
func (o *CompartimentoEstado) SetCompartimento(v Compartimento) {
	o.Compartimento = &v
}

func (o CompartimentoEstado) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Upd.IsSet() {
		toSerialize["upd"] = o.Upd.Get()
	}
	if o.Usr.IsSet() {
		toSerialize["usr"] = o.Usr.Get()
	}
	if o.RefProduto.IsSet() {
		toSerialize["refProduto"] = o.RefProduto.Get()
	}
	if o.Estado != nil {
		toSerialize["estado"] = o.Estado
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Ocupado != nil {
		toSerialize["ocupado"] = o.Ocupado
	}
	if o.RefCompartimento.IsSet() {
		toSerialize["refCompartimento"] = o.RefCompartimento.Get()
	}
	if o.Vasilhame != nil {
		toSerialize["vasilhame"] = o.Vasilhame
	}
	if o.Compartimento != nil {
		toSerialize["compartimento"] = o.Compartimento
	}
	return json.Marshal(toSerialize)
}

type NullableCompartimentoEstado struct {
	value *CompartimentoEstado
	isSet bool
}

func (v NullableCompartimentoEstado) Get() *CompartimentoEstado {
	return v.value
}

func (v *NullableCompartimentoEstado) Set(val *CompartimentoEstado) {
	v.value = val
	v.isSet = true
}

func (v NullableCompartimentoEstado) IsSet() bool {
	return v.isSet
}

func (v *NullableCompartimentoEstado) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompartimentoEstado(val *CompartimentoEstado) *NullableCompartimentoEstado {
	return &NullableCompartimentoEstado{value: val, isSet: true}
}

func (v NullableCompartimentoEstado) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompartimentoEstado) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


