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

// PerfilSegurancaPermissao struct for PerfilSegurancaPermissao
type PerfilSegurancaPermissao struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefPerfilSeguranca NullableInt64 `json:"refPerfilSeguranca,omitempty"`
	RefPermissao NullableInt64 `json:"refPermissao,omitempty"`
	PerfilSeguranca *PerfilSeguranca `json:"perfilSeguranca,omitempty"`
	Permissao *Permissao `json:"permissao,omitempty"`
}

// NewPerfilSegurancaPermissao instantiates a new PerfilSegurancaPermissao object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfilSegurancaPermissao() *PerfilSegurancaPermissao {
	this := PerfilSegurancaPermissao{}
	return &this
}

// NewPerfilSegurancaPermissaoWithDefaults instantiates a new PerfilSegurancaPermissao object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfilSegurancaPermissaoWithDefaults() *PerfilSegurancaPermissao {
	this := PerfilSegurancaPermissao{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PerfilSegurancaPermissao) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfilSegurancaPermissao) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PerfilSegurancaPermissao) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *PerfilSegurancaPermissao) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfilSegurancaPermissao) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *PerfilSegurancaPermissao) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PerfilSegurancaPermissao) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PerfilSegurancaPermissao) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *PerfilSegurancaPermissao) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *PerfilSegurancaPermissao) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *PerfilSegurancaPermissao) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PerfilSegurancaPermissao) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PerfilSegurancaPermissao) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *PerfilSegurancaPermissao) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *PerfilSegurancaPermissao) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *PerfilSegurancaPermissao) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefPerfilSeguranca returns the RefPerfilSeguranca field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PerfilSegurancaPermissao) GetRefPerfilSeguranca() int64 {
	if o == nil || o.RefPerfilSeguranca.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefPerfilSeguranca.Get()
}

// GetRefPerfilSegurancaOk returns a tuple with the RefPerfilSeguranca field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PerfilSegurancaPermissao) GetRefPerfilSegurancaOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefPerfilSeguranca.Get(), o.RefPerfilSeguranca.IsSet()
}

// HasRefPerfilSeguranca returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasRefPerfilSeguranca() bool {
	if o != nil && o.RefPerfilSeguranca.IsSet() {
		return true
	}

	return false
}

// SetRefPerfilSeguranca gets a reference to the given NullableInt64 and assigns it to the RefPerfilSeguranca field.
func (o *PerfilSegurancaPermissao) SetRefPerfilSeguranca(v int64) {
	o.RefPerfilSeguranca.Set(&v)
}
// SetRefPerfilSegurancaNil sets the value for RefPerfilSeguranca to be an explicit nil
func (o *PerfilSegurancaPermissao) SetRefPerfilSegurancaNil() {
	o.RefPerfilSeguranca.Set(nil)
}

// UnsetRefPerfilSeguranca ensures that no value is present for RefPerfilSeguranca, not even an explicit nil
func (o *PerfilSegurancaPermissao) UnsetRefPerfilSeguranca() {
	o.RefPerfilSeguranca.Unset()
}

// GetRefPermissao returns the RefPermissao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PerfilSegurancaPermissao) GetRefPermissao() int64 {
	if o == nil || o.RefPermissao.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefPermissao.Get()
}

// GetRefPermissaoOk returns a tuple with the RefPermissao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PerfilSegurancaPermissao) GetRefPermissaoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefPermissao.Get(), o.RefPermissao.IsSet()
}

// HasRefPermissao returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasRefPermissao() bool {
	if o != nil && o.RefPermissao.IsSet() {
		return true
	}

	return false
}

// SetRefPermissao gets a reference to the given NullableInt64 and assigns it to the RefPermissao field.
func (o *PerfilSegurancaPermissao) SetRefPermissao(v int64) {
	o.RefPermissao.Set(&v)
}
// SetRefPermissaoNil sets the value for RefPermissao to be an explicit nil
func (o *PerfilSegurancaPermissao) SetRefPermissaoNil() {
	o.RefPermissao.Set(nil)
}

// UnsetRefPermissao ensures that no value is present for RefPermissao, not even an explicit nil
func (o *PerfilSegurancaPermissao) UnsetRefPermissao() {
	o.RefPermissao.Unset()
}

// GetPerfilSeguranca returns the PerfilSeguranca field value if set, zero value otherwise.
func (o *PerfilSegurancaPermissao) GetPerfilSeguranca() PerfilSeguranca {
	if o == nil || o.PerfilSeguranca == nil {
		var ret PerfilSeguranca
		return ret
	}
	return *o.PerfilSeguranca
}

// GetPerfilSegurancaOk returns a tuple with the PerfilSeguranca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfilSegurancaPermissao) GetPerfilSegurancaOk() (*PerfilSeguranca, bool) {
	if o == nil || o.PerfilSeguranca == nil {
		return nil, false
	}
	return o.PerfilSeguranca, true
}

// HasPerfilSeguranca returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasPerfilSeguranca() bool {
	if o != nil && o.PerfilSeguranca != nil {
		return true
	}

	return false
}

// SetPerfilSeguranca gets a reference to the given PerfilSeguranca and assigns it to the PerfilSeguranca field.
func (o *PerfilSegurancaPermissao) SetPerfilSeguranca(v PerfilSeguranca) {
	o.PerfilSeguranca = &v
}

// GetPermissao returns the Permissao field value if set, zero value otherwise.
func (o *PerfilSegurancaPermissao) GetPermissao() Permissao {
	if o == nil || o.Permissao == nil {
		var ret Permissao
		return ret
	}
	return *o.Permissao
}

// GetPermissaoOk returns a tuple with the Permissao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfilSegurancaPermissao) GetPermissaoOk() (*Permissao, bool) {
	if o == nil || o.Permissao == nil {
		return nil, false
	}
	return o.Permissao, true
}

// HasPermissao returns a boolean if a field has been set.
func (o *PerfilSegurancaPermissao) HasPermissao() bool {
	if o != nil && o.Permissao != nil {
		return true
	}

	return false
}

// SetPermissao gets a reference to the given Permissao and assigns it to the Permissao field.
func (o *PerfilSegurancaPermissao) SetPermissao(v Permissao) {
	o.Permissao = &v
}

func (o PerfilSegurancaPermissao) MarshalJSON() ([]byte, error) {
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
	if o.RefPerfilSeguranca.IsSet() {
		toSerialize["refPerfilSeguranca"] = o.RefPerfilSeguranca.Get()
	}
	if o.RefPermissao.IsSet() {
		toSerialize["refPermissao"] = o.RefPermissao.Get()
	}
	if o.PerfilSeguranca != nil {
		toSerialize["perfilSeguranca"] = o.PerfilSeguranca
	}
	if o.Permissao != nil {
		toSerialize["permissao"] = o.Permissao
	}
	return json.Marshal(toSerialize)
}

type NullablePerfilSegurancaPermissao struct {
	value *PerfilSegurancaPermissao
	isSet bool
}

func (v NullablePerfilSegurancaPermissao) Get() *PerfilSegurancaPermissao {
	return v.value
}

func (v *NullablePerfilSegurancaPermissao) Set(val *PerfilSegurancaPermissao) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfilSegurancaPermissao) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfilSegurancaPermissao) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfilSegurancaPermissao(val *PerfilSegurancaPermissao) *NullablePerfilSegurancaPermissao {
	return &NullablePerfilSegurancaPermissao{value: val, isSet: true}
}

func (v NullablePerfilSegurancaPermissao) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfilSegurancaPermissao) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

