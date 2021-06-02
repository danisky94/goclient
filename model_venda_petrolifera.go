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

// VendaPetrolifera struct for VendaPetrolifera
type VendaPetrolifera struct {
	RefGvm NullableInt64 `json:"refGvm,omitempty"`
	RefProduto NullableInt64 `json:"refProduto,omitempty"`
	Vasilhame *bool `json:"vasilhame,omitempty"`
	Guid NullableString `json:"guid,omitempty"`
	Origem *TipoOrigem `json:"origem,omitempty"`
}

// NewVendaPetrolifera instantiates a new VendaPetrolifera object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendaPetrolifera() *VendaPetrolifera {
	this := VendaPetrolifera{}
	return &this
}

// NewVendaPetroliferaWithDefaults instantiates a new VendaPetrolifera object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendaPetroliferaWithDefaults() *VendaPetrolifera {
	this := VendaPetrolifera{}
	return &this
}

// GetRefGvm returns the RefGvm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendaPetrolifera) GetRefGvm() int64 {
	if o == nil || o.RefGvm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefGvm.Get()
}

// GetRefGvmOk returns a tuple with the RefGvm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendaPetrolifera) GetRefGvmOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefGvm.Get(), o.RefGvm.IsSet()
}

// HasRefGvm returns a boolean if a field has been set.
func (o *VendaPetrolifera) HasRefGvm() bool {
	if o != nil && o.RefGvm.IsSet() {
		return true
	}

	return false
}

// SetRefGvm gets a reference to the given NullableInt64 and assigns it to the RefGvm field.
func (o *VendaPetrolifera) SetRefGvm(v int64) {
	o.RefGvm.Set(&v)
}
// SetRefGvmNil sets the value for RefGvm to be an explicit nil
func (o *VendaPetrolifera) SetRefGvmNil() {
	o.RefGvm.Set(nil)
}

// UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
func (o *VendaPetrolifera) UnsetRefGvm() {
	o.RefGvm.Unset()
}

// GetRefProduto returns the RefProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendaPetrolifera) GetRefProduto() int64 {
	if o == nil || o.RefProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefProduto.Get()
}

// GetRefProdutoOk returns a tuple with the RefProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendaPetrolifera) GetRefProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefProduto.Get(), o.RefProduto.IsSet()
}

// HasRefProduto returns a boolean if a field has been set.
func (o *VendaPetrolifera) HasRefProduto() bool {
	if o != nil && o.RefProduto.IsSet() {
		return true
	}

	return false
}

// SetRefProduto gets a reference to the given NullableInt64 and assigns it to the RefProduto field.
func (o *VendaPetrolifera) SetRefProduto(v int64) {
	o.RefProduto.Set(&v)
}
// SetRefProdutoNil sets the value for RefProduto to be an explicit nil
func (o *VendaPetrolifera) SetRefProdutoNil() {
	o.RefProduto.Set(nil)
}

// UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
func (o *VendaPetrolifera) UnsetRefProduto() {
	o.RefProduto.Unset()
}

// GetVasilhame returns the Vasilhame field value if set, zero value otherwise.
func (o *VendaPetrolifera) GetVasilhame() bool {
	if o == nil || o.Vasilhame == nil {
		var ret bool
		return ret
	}
	return *o.Vasilhame
}

// GetVasilhameOk returns a tuple with the Vasilhame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendaPetrolifera) GetVasilhameOk() (*bool, bool) {
	if o == nil || o.Vasilhame == nil {
		return nil, false
	}
	return o.Vasilhame, true
}

// HasVasilhame returns a boolean if a field has been set.
func (o *VendaPetrolifera) HasVasilhame() bool {
	if o != nil && o.Vasilhame != nil {
		return true
	}

	return false
}

// SetVasilhame gets a reference to the given bool and assigns it to the Vasilhame field.
func (o *VendaPetrolifera) SetVasilhame(v bool) {
	o.Vasilhame = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendaPetrolifera) GetGuid() string {
	if o == nil || o.Guid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendaPetrolifera) GetGuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *VendaPetrolifera) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *VendaPetrolifera) SetGuid(v string) {
	o.Guid.Set(&v)
}
// SetGuidNil sets the value for Guid to be an explicit nil
func (o *VendaPetrolifera) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *VendaPetrolifera) UnsetGuid() {
	o.Guid.Unset()
}

// GetOrigem returns the Origem field value if set, zero value otherwise.
func (o *VendaPetrolifera) GetOrigem() TipoOrigem {
	if o == nil || o.Origem == nil {
		var ret TipoOrigem
		return ret
	}
	return *o.Origem
}

// GetOrigemOk returns a tuple with the Origem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendaPetrolifera) GetOrigemOk() (*TipoOrigem, bool) {
	if o == nil || o.Origem == nil {
		return nil, false
	}
	return o.Origem, true
}

// HasOrigem returns a boolean if a field has been set.
func (o *VendaPetrolifera) HasOrigem() bool {
	if o != nil && o.Origem != nil {
		return true
	}

	return false
}

// SetOrigem gets a reference to the given TipoOrigem and assigns it to the Origem field.
func (o *VendaPetrolifera) SetOrigem(v TipoOrigem) {
	o.Origem = &v
}

func (o VendaPetrolifera) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RefGvm.IsSet() {
		toSerialize["refGvm"] = o.RefGvm.Get()
	}
	if o.RefProduto.IsSet() {
		toSerialize["refProduto"] = o.RefProduto.Get()
	}
	if o.Vasilhame != nil {
		toSerialize["vasilhame"] = o.Vasilhame
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.Origem != nil {
		toSerialize["origem"] = o.Origem
	}
	return json.Marshal(toSerialize)
}

type NullableVendaPetrolifera struct {
	value *VendaPetrolifera
	isSet bool
}

func (v NullableVendaPetrolifera) Get() *VendaPetrolifera {
	return v.value
}

func (v *NullableVendaPetrolifera) Set(val *VendaPetrolifera) {
	v.value = val
	v.isSet = true
}

func (v NullableVendaPetrolifera) IsSet() bool {
	return v.isSet
}

func (v *NullableVendaPetrolifera) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendaPetrolifera(val *VendaPetrolifera) *NullableVendaPetrolifera {
	return &NullableVendaPetrolifera{value: val, isSet: true}
}

func (v NullableVendaPetrolifera) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendaPetrolifera) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

