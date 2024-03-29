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

// TipoEvento struct for TipoEvento
type TipoEvento struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	Tipo *TiposTipoEvento `json:"tipo,omitempty"`
	RefTipoNotificacao NullableInt64 `json:"refTipoNotificacao,omitempty"`
	Notificacao *TipoNotificacao `json:"notificacao,omitempty"`
	Eventos []Evento `json:"eventos,omitempty"`
}

// NewTipoEvento instantiates a new TipoEvento object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTipoEvento() *TipoEvento {
	this := TipoEvento{}
	return &this
}

// NewTipoEventoWithDefaults instantiates a new TipoEvento object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTipoEventoWithDefaults() *TipoEvento {
	this := TipoEvento{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TipoEvento) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoEvento) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TipoEvento) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TipoEvento) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TipoEvento) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoEvento) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TipoEvento) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *TipoEvento) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoEvento) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoEvento) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *TipoEvento) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *TipoEvento) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *TipoEvento) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *TipoEvento) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoEvento) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoEvento) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *TipoEvento) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *TipoEvento) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *TipoEvento) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *TipoEvento) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoEvento) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoEvento) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *TipoEvento) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *TipoEvento) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *TipoEvento) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *TipoEvento) UnsetNome() {
	o.Nome.Unset()
}

// GetTipo returns the Tipo field value if set, zero value otherwise.
func (o *TipoEvento) GetTipo() TiposTipoEvento {
	if o == nil || o.Tipo == nil {
		var ret TiposTipoEvento
		return ret
	}
	return *o.Tipo
}

// GetTipoOk returns a tuple with the Tipo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoEvento) GetTipoOk() (*TiposTipoEvento, bool) {
	if o == nil || o.Tipo == nil {
		return nil, false
	}
	return o.Tipo, true
}

// HasTipo returns a boolean if a field has been set.
func (o *TipoEvento) HasTipo() bool {
	if o != nil && o.Tipo != nil {
		return true
	}

	return false
}

// SetTipo gets a reference to the given TiposTipoEvento and assigns it to the Tipo field.
func (o *TipoEvento) SetTipo(v TiposTipoEvento) {
	o.Tipo = &v
}

// GetRefTipoNotificacao returns the RefTipoNotificacao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoEvento) GetRefTipoNotificacao() int64 {
	if o == nil || o.RefTipoNotificacao.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefTipoNotificacao.Get()
}

// GetRefTipoNotificacaoOk returns a tuple with the RefTipoNotificacao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoEvento) GetRefTipoNotificacaoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefTipoNotificacao.Get(), o.RefTipoNotificacao.IsSet()
}

// HasRefTipoNotificacao returns a boolean if a field has been set.
func (o *TipoEvento) HasRefTipoNotificacao() bool {
	if o != nil && o.RefTipoNotificacao.IsSet() {
		return true
	}

	return false
}

// SetRefTipoNotificacao gets a reference to the given NullableInt64 and assigns it to the RefTipoNotificacao field.
func (o *TipoEvento) SetRefTipoNotificacao(v int64) {
	o.RefTipoNotificacao.Set(&v)
}
// SetRefTipoNotificacaoNil sets the value for RefTipoNotificacao to be an explicit nil
func (o *TipoEvento) SetRefTipoNotificacaoNil() {
	o.RefTipoNotificacao.Set(nil)
}

// UnsetRefTipoNotificacao ensures that no value is present for RefTipoNotificacao, not even an explicit nil
func (o *TipoEvento) UnsetRefTipoNotificacao() {
	o.RefTipoNotificacao.Unset()
}

// GetNotificacao returns the Notificacao field value if set, zero value otherwise.
func (o *TipoEvento) GetNotificacao() TipoNotificacao {
	if o == nil || o.Notificacao == nil {
		var ret TipoNotificacao
		return ret
	}
	return *o.Notificacao
}

// GetNotificacaoOk returns a tuple with the Notificacao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoEvento) GetNotificacaoOk() (*TipoNotificacao, bool) {
	if o == nil || o.Notificacao == nil {
		return nil, false
	}
	return o.Notificacao, true
}

// HasNotificacao returns a boolean if a field has been set.
func (o *TipoEvento) HasNotificacao() bool {
	if o != nil && o.Notificacao != nil {
		return true
	}

	return false
}

// SetNotificacao gets a reference to the given TipoNotificacao and assigns it to the Notificacao field.
func (o *TipoEvento) SetNotificacao(v TipoNotificacao) {
	o.Notificacao = &v
}

// GetEventos returns the Eventos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoEvento) GetEventos() []Evento {
	if o == nil  {
		var ret []Evento
		return ret
	}
	return o.Eventos
}

// GetEventosOk returns a tuple with the Eventos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoEvento) GetEventosOk() (*[]Evento, bool) {
	if o == nil || o.Eventos == nil {
		return nil, false
	}
	return &o.Eventos, true
}

// HasEventos returns a boolean if a field has been set.
func (o *TipoEvento) HasEventos() bool {
	if o != nil && o.Eventos != nil {
		return true
	}

	return false
}

// SetEventos gets a reference to the given []Evento and assigns it to the Eventos field.
func (o *TipoEvento) SetEventos(v []Evento) {
	o.Eventos = v
}

func (o TipoEvento) MarshalJSON() ([]byte, error) {
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
	if o.Nome.IsSet() {
		toSerialize["nome"] = o.Nome.Get()
	}
	if o.Tipo != nil {
		toSerialize["tipo"] = o.Tipo
	}
	if o.RefTipoNotificacao.IsSet() {
		toSerialize["refTipoNotificacao"] = o.RefTipoNotificacao.Get()
	}
	if o.Notificacao != nil {
		toSerialize["notificacao"] = o.Notificacao
	}
	if o.Eventos != nil {
		toSerialize["eventos"] = o.Eventos
	}
	return json.Marshal(toSerialize)
}

type NullableTipoEvento struct {
	value *TipoEvento
	isSet bool
}

func (v NullableTipoEvento) Get() *TipoEvento {
	return v.value
}

func (v *NullableTipoEvento) Set(val *TipoEvento) {
	v.value = val
	v.isSet = true
}

func (v NullableTipoEvento) IsSet() bool {
	return v.isSet
}

func (v *NullableTipoEvento) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTipoEvento(val *TipoEvento) *NullableTipoEvento {
	return &NullableTipoEvento{value: val, isSet: true}
}

func (v NullableTipoEvento) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTipoEvento) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


