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

// TipoNotificacaoUtilizador struct for TipoNotificacaoUtilizador
type TipoNotificacaoUtilizador struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefTipoNotificacao NullableInt64 `json:"refTipoNotificacao,omitempty"`
	RefUtilizador NullableInt64 `json:"refUtilizador,omitempty"`
	PorEmail NullableBool `json:"porEmail,omitempty"`
	PorSms NullableBool `json:"porSms,omitempty"`
	TipoNotificacao *TipoNotificacao `json:"tipoNotificacao,omitempty"`
	Utilizador *Utilizador `json:"utilizador,omitempty"`
}

// NewTipoNotificacaoUtilizador instantiates a new TipoNotificacaoUtilizador object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTipoNotificacaoUtilizador() *TipoNotificacaoUtilizador {
	this := TipoNotificacaoUtilizador{}
	return &this
}

// NewTipoNotificacaoUtilizadorWithDefaults instantiates a new TipoNotificacaoUtilizador object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTipoNotificacaoUtilizadorWithDefaults() *TipoNotificacaoUtilizador {
	this := TipoNotificacaoUtilizador{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TipoNotificacaoUtilizador) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacaoUtilizador) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TipoNotificacaoUtilizador) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TipoNotificacaoUtilizador) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacaoUtilizador) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *TipoNotificacaoUtilizador) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *TipoNotificacaoUtilizador) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *TipoNotificacaoUtilizador) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefTipoNotificacao returns the RefTipoNotificacao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetRefTipoNotificacao() int64 {
	if o == nil || o.RefTipoNotificacao.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefTipoNotificacao.Get()
}

// GetRefTipoNotificacaoOk returns a tuple with the RefTipoNotificacao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetRefTipoNotificacaoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefTipoNotificacao.Get(), o.RefTipoNotificacao.IsSet()
}

// HasRefTipoNotificacao returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasRefTipoNotificacao() bool {
	if o != nil && o.RefTipoNotificacao.IsSet() {
		return true
	}

	return false
}

// SetRefTipoNotificacao gets a reference to the given NullableInt64 and assigns it to the RefTipoNotificacao field.
func (o *TipoNotificacaoUtilizador) SetRefTipoNotificacao(v int64) {
	o.RefTipoNotificacao.Set(&v)
}
// SetRefTipoNotificacaoNil sets the value for RefTipoNotificacao to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetRefTipoNotificacaoNil() {
	o.RefTipoNotificacao.Set(nil)
}

// UnsetRefTipoNotificacao ensures that no value is present for RefTipoNotificacao, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetRefTipoNotificacao() {
	o.RefTipoNotificacao.Unset()
}

// GetRefUtilizador returns the RefUtilizador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetRefUtilizador() int64 {
	if o == nil || o.RefUtilizador.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefUtilizador.Get()
}

// GetRefUtilizadorOk returns a tuple with the RefUtilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetRefUtilizadorOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefUtilizador.Get(), o.RefUtilizador.IsSet()
}

// HasRefUtilizador returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasRefUtilizador() bool {
	if o != nil && o.RefUtilizador.IsSet() {
		return true
	}

	return false
}

// SetRefUtilizador gets a reference to the given NullableInt64 and assigns it to the RefUtilizador field.
func (o *TipoNotificacaoUtilizador) SetRefUtilizador(v int64) {
	o.RefUtilizador.Set(&v)
}
// SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetRefUtilizadorNil() {
	o.RefUtilizador.Set(nil)
}

// UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetRefUtilizador() {
	o.RefUtilizador.Unset()
}

// GetPorEmail returns the PorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetPorEmail() bool {
	if o == nil || o.PorEmail.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PorEmail.Get()
}

// GetPorEmailOk returns a tuple with the PorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetPorEmailOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PorEmail.Get(), o.PorEmail.IsSet()
}

// HasPorEmail returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasPorEmail() bool {
	if o != nil && o.PorEmail.IsSet() {
		return true
	}

	return false
}

// SetPorEmail gets a reference to the given NullableBool and assigns it to the PorEmail field.
func (o *TipoNotificacaoUtilizador) SetPorEmail(v bool) {
	o.PorEmail.Set(&v)
}
// SetPorEmailNil sets the value for PorEmail to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetPorEmailNil() {
	o.PorEmail.Set(nil)
}

// UnsetPorEmail ensures that no value is present for PorEmail, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetPorEmail() {
	o.PorEmail.Unset()
}

// GetPorSms returns the PorSms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacaoUtilizador) GetPorSms() bool {
	if o == nil || o.PorSms.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PorSms.Get()
}

// GetPorSmsOk returns a tuple with the PorSms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacaoUtilizador) GetPorSmsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PorSms.Get(), o.PorSms.IsSet()
}

// HasPorSms returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasPorSms() bool {
	if o != nil && o.PorSms.IsSet() {
		return true
	}

	return false
}

// SetPorSms gets a reference to the given NullableBool and assigns it to the PorSms field.
func (o *TipoNotificacaoUtilizador) SetPorSms(v bool) {
	o.PorSms.Set(&v)
}
// SetPorSmsNil sets the value for PorSms to be an explicit nil
func (o *TipoNotificacaoUtilizador) SetPorSmsNil() {
	o.PorSms.Set(nil)
}

// UnsetPorSms ensures that no value is present for PorSms, not even an explicit nil
func (o *TipoNotificacaoUtilizador) UnsetPorSms() {
	o.PorSms.Unset()
}

// GetTipoNotificacao returns the TipoNotificacao field value if set, zero value otherwise.
func (o *TipoNotificacaoUtilizador) GetTipoNotificacao() TipoNotificacao {
	if o == nil || o.TipoNotificacao == nil {
		var ret TipoNotificacao
		return ret
	}
	return *o.TipoNotificacao
}

// GetTipoNotificacaoOk returns a tuple with the TipoNotificacao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacaoUtilizador) GetTipoNotificacaoOk() (*TipoNotificacao, bool) {
	if o == nil || o.TipoNotificacao == nil {
		return nil, false
	}
	return o.TipoNotificacao, true
}

// HasTipoNotificacao returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasTipoNotificacao() bool {
	if o != nil && o.TipoNotificacao != nil {
		return true
	}

	return false
}

// SetTipoNotificacao gets a reference to the given TipoNotificacao and assigns it to the TipoNotificacao field.
func (o *TipoNotificacaoUtilizador) SetTipoNotificacao(v TipoNotificacao) {
	o.TipoNotificacao = &v
}

// GetUtilizador returns the Utilizador field value if set, zero value otherwise.
func (o *TipoNotificacaoUtilizador) GetUtilizador() Utilizador {
	if o == nil || o.Utilizador == nil {
		var ret Utilizador
		return ret
	}
	return *o.Utilizador
}

// GetUtilizadorOk returns a tuple with the Utilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacaoUtilizador) GetUtilizadorOk() (*Utilizador, bool) {
	if o == nil || o.Utilizador == nil {
		return nil, false
	}
	return o.Utilizador, true
}

// HasUtilizador returns a boolean if a field has been set.
func (o *TipoNotificacaoUtilizador) HasUtilizador() bool {
	if o != nil && o.Utilizador != nil {
		return true
	}

	return false
}

// SetUtilizador gets a reference to the given Utilizador and assigns it to the Utilizador field.
func (o *TipoNotificacaoUtilizador) SetUtilizador(v Utilizador) {
	o.Utilizador = &v
}

func (o TipoNotificacaoUtilizador) MarshalJSON() ([]byte, error) {
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
	if o.RefTipoNotificacao.IsSet() {
		toSerialize["refTipoNotificacao"] = o.RefTipoNotificacao.Get()
	}
	if o.RefUtilizador.IsSet() {
		toSerialize["refUtilizador"] = o.RefUtilizador.Get()
	}
	if o.PorEmail.IsSet() {
		toSerialize["porEmail"] = o.PorEmail.Get()
	}
	if o.PorSms.IsSet() {
		toSerialize["porSms"] = o.PorSms.Get()
	}
	if o.TipoNotificacao != nil {
		toSerialize["tipoNotificacao"] = o.TipoNotificacao
	}
	if o.Utilizador != nil {
		toSerialize["utilizador"] = o.Utilizador
	}
	return json.Marshal(toSerialize)
}

type NullableTipoNotificacaoUtilizador struct {
	value *TipoNotificacaoUtilizador
	isSet bool
}

func (v NullableTipoNotificacaoUtilizador) Get() *TipoNotificacaoUtilizador {
	return v.value
}

func (v *NullableTipoNotificacaoUtilizador) Set(val *TipoNotificacaoUtilizador) {
	v.value = val
	v.isSet = true
}

func (v NullableTipoNotificacaoUtilizador) IsSet() bool {
	return v.isSet
}

func (v *NullableTipoNotificacaoUtilizador) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTipoNotificacaoUtilizador(val *TipoNotificacaoUtilizador) *NullableTipoNotificacaoUtilizador {
	return &NullableTipoNotificacaoUtilizador{value: val, isSet: true}
}

func (v NullableTipoNotificacaoUtilizador) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTipoNotificacaoUtilizador) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


