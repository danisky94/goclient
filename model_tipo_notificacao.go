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

// TipoNotificacao struct for TipoNotificacao
type TipoNotificacao struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	MensagemEmail NullableString `json:"mensagemEmail,omitempty"`
	MenagemSms NullableString `json:"menagemSms,omitempty"`
	Ativa NullableBool `json:"ativa,omitempty"`
	TipoEventos []TipoEvento `json:"tipoEventos,omitempty"`
	NotificacaoUtilizadores []TipoNotificacaoUtilizador `json:"notificacaoUtilizadores,omitempty"`
}

// NewTipoNotificacao instantiates a new TipoNotificacao object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTipoNotificacao() *TipoNotificacao {
	this := TipoNotificacao{}
	return &this
}

// NewTipoNotificacaoWithDefaults instantiates a new TipoNotificacao object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTipoNotificacaoWithDefaults() *TipoNotificacao {
	this := TipoNotificacao{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TipoNotificacao) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacao) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TipoNotificacao) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TipoNotificacao) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TipoNotificacao) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TipoNotificacao) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TipoNotificacao) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *TipoNotificacao) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *TipoNotificacao) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *TipoNotificacao) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *TipoNotificacao) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *TipoNotificacao) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *TipoNotificacao) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *TipoNotificacao) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *TipoNotificacao) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *TipoNotificacao) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *TipoNotificacao) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *TipoNotificacao) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *TipoNotificacao) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *TipoNotificacao) UnsetNome() {
	o.Nome.Unset()
}

// GetMensagemEmail returns the MensagemEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetMensagemEmail() string {
	if o == nil || o.MensagemEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.MensagemEmail.Get()
}

// GetMensagemEmailOk returns a tuple with the MensagemEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetMensagemEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MensagemEmail.Get(), o.MensagemEmail.IsSet()
}

// HasMensagemEmail returns a boolean if a field has been set.
func (o *TipoNotificacao) HasMensagemEmail() bool {
	if o != nil && o.MensagemEmail.IsSet() {
		return true
	}

	return false
}

// SetMensagemEmail gets a reference to the given NullableString and assigns it to the MensagemEmail field.
func (o *TipoNotificacao) SetMensagemEmail(v string) {
	o.MensagemEmail.Set(&v)
}
// SetMensagemEmailNil sets the value for MensagemEmail to be an explicit nil
func (o *TipoNotificacao) SetMensagemEmailNil() {
	o.MensagemEmail.Set(nil)
}

// UnsetMensagemEmail ensures that no value is present for MensagemEmail, not even an explicit nil
func (o *TipoNotificacao) UnsetMensagemEmail() {
	o.MensagemEmail.Unset()
}

// GetMenagemSms returns the MenagemSms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetMenagemSms() string {
	if o == nil || o.MenagemSms.Get() == nil {
		var ret string
		return ret
	}
	return *o.MenagemSms.Get()
}

// GetMenagemSmsOk returns a tuple with the MenagemSms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetMenagemSmsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MenagemSms.Get(), o.MenagemSms.IsSet()
}

// HasMenagemSms returns a boolean if a field has been set.
func (o *TipoNotificacao) HasMenagemSms() bool {
	if o != nil && o.MenagemSms.IsSet() {
		return true
	}

	return false
}

// SetMenagemSms gets a reference to the given NullableString and assigns it to the MenagemSms field.
func (o *TipoNotificacao) SetMenagemSms(v string) {
	o.MenagemSms.Set(&v)
}
// SetMenagemSmsNil sets the value for MenagemSms to be an explicit nil
func (o *TipoNotificacao) SetMenagemSmsNil() {
	o.MenagemSms.Set(nil)
}

// UnsetMenagemSms ensures that no value is present for MenagemSms, not even an explicit nil
func (o *TipoNotificacao) UnsetMenagemSms() {
	o.MenagemSms.Unset()
}

// GetAtiva returns the Ativa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetAtiva() bool {
	if o == nil || o.Ativa.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Ativa.Get()
}

// GetAtivaOk returns a tuple with the Ativa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetAtivaOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ativa.Get(), o.Ativa.IsSet()
}

// HasAtiva returns a boolean if a field has been set.
func (o *TipoNotificacao) HasAtiva() bool {
	if o != nil && o.Ativa.IsSet() {
		return true
	}

	return false
}

// SetAtiva gets a reference to the given NullableBool and assigns it to the Ativa field.
func (o *TipoNotificacao) SetAtiva(v bool) {
	o.Ativa.Set(&v)
}
// SetAtivaNil sets the value for Ativa to be an explicit nil
func (o *TipoNotificacao) SetAtivaNil() {
	o.Ativa.Set(nil)
}

// UnsetAtiva ensures that no value is present for Ativa, not even an explicit nil
func (o *TipoNotificacao) UnsetAtiva() {
	o.Ativa.Unset()
}

// GetTipoEventos returns the TipoEventos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetTipoEventos() []TipoEvento {
	if o == nil  {
		var ret []TipoEvento
		return ret
	}
	return o.TipoEventos
}

// GetTipoEventosOk returns a tuple with the TipoEventos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetTipoEventosOk() (*[]TipoEvento, bool) {
	if o == nil || o.TipoEventos == nil {
		return nil, false
	}
	return &o.TipoEventos, true
}

// HasTipoEventos returns a boolean if a field has been set.
func (o *TipoNotificacao) HasTipoEventos() bool {
	if o != nil && o.TipoEventos != nil {
		return true
	}

	return false
}

// SetTipoEventos gets a reference to the given []TipoEvento and assigns it to the TipoEventos field.
func (o *TipoNotificacao) SetTipoEventos(v []TipoEvento) {
	o.TipoEventos = v
}

// GetNotificacaoUtilizadores returns the NotificacaoUtilizadores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TipoNotificacao) GetNotificacaoUtilizadores() []TipoNotificacaoUtilizador {
	if o == nil  {
		var ret []TipoNotificacaoUtilizador
		return ret
	}
	return o.NotificacaoUtilizadores
}

// GetNotificacaoUtilizadoresOk returns a tuple with the NotificacaoUtilizadores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TipoNotificacao) GetNotificacaoUtilizadoresOk() (*[]TipoNotificacaoUtilizador, bool) {
	if o == nil || o.NotificacaoUtilizadores == nil {
		return nil, false
	}
	return &o.NotificacaoUtilizadores, true
}

// HasNotificacaoUtilizadores returns a boolean if a field has been set.
func (o *TipoNotificacao) HasNotificacaoUtilizadores() bool {
	if o != nil && o.NotificacaoUtilizadores != nil {
		return true
	}

	return false
}

// SetNotificacaoUtilizadores gets a reference to the given []TipoNotificacaoUtilizador and assigns it to the NotificacaoUtilizadores field.
func (o *TipoNotificacao) SetNotificacaoUtilizadores(v []TipoNotificacaoUtilizador) {
	o.NotificacaoUtilizadores = v
}

func (o TipoNotificacao) MarshalJSON() ([]byte, error) {
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
	if o.MensagemEmail.IsSet() {
		toSerialize["mensagemEmail"] = o.MensagemEmail.Get()
	}
	if o.MenagemSms.IsSet() {
		toSerialize["menagemSms"] = o.MenagemSms.Get()
	}
	if o.Ativa.IsSet() {
		toSerialize["ativa"] = o.Ativa.Get()
	}
	if o.TipoEventos != nil {
		toSerialize["tipoEventos"] = o.TipoEventos
	}
	if o.NotificacaoUtilizadores != nil {
		toSerialize["notificacaoUtilizadores"] = o.NotificacaoUtilizadores
	}
	return json.Marshal(toSerialize)
}

type NullableTipoNotificacao struct {
	value *TipoNotificacao
	isSet bool
}

func (v NullableTipoNotificacao) Get() *TipoNotificacao {
	return v.value
}

func (v *NullableTipoNotificacao) Set(val *TipoNotificacao) {
	v.value = val
	v.isSet = true
}

func (v NullableTipoNotificacao) IsSet() bool {
	return v.isSet
}

func (v *NullableTipoNotificacao) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTipoNotificacao(val *TipoNotificacao) *NullableTipoNotificacao {
	return &NullableTipoNotificacao{value: val, isSet: true}
}

func (v NullableTipoNotificacao) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTipoNotificacao) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


