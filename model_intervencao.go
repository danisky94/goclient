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

// Intervencao struct for Intervencao
type Intervencao struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefTarefa NullableInt64 `json:"refTarefa,omitempty"`
	Descricao NullableString `json:"descricao,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	RefUtilizador NullableInt64 `json:"refUtilizador,omitempty"`
	Estado *IntervencaoEstados `json:"estado,omitempty"`
	NomeUtilizador NullableString `json:"nomeUtilizador,omitempty"`
	RefTarefaNavigation *Tarefa `json:"refTarefaNavigation,omitempty"`
	RefUtilizadorNavigation *Utilizador `json:"refUtilizadorNavigation,omitempty"`
}

// NewIntervencao instantiates a new Intervencao object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntervencao() *Intervencao {
	this := Intervencao{}
	return &this
}

// NewIntervencaoWithDefaults instantiates a new Intervencao object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntervencaoWithDefaults() *Intervencao {
	this := Intervencao{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Intervencao) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intervencao) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Intervencao) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Intervencao) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Intervencao) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intervencao) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Intervencao) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Intervencao) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Intervencao) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Intervencao) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Intervencao) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Intervencao) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Intervencao) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Intervencao) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Intervencao) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Intervencao) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefTarefa returns the RefTarefa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetRefTarefa() int64 {
	if o == nil || o.RefTarefa.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefTarefa.Get()
}

// GetRefTarefaOk returns a tuple with the RefTarefa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetRefTarefaOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefTarefa.Get(), o.RefTarefa.IsSet()
}

// HasRefTarefa returns a boolean if a field has been set.
func (o *Intervencao) HasRefTarefa() bool {
	if o != nil && o.RefTarefa.IsSet() {
		return true
	}

	return false
}

// SetRefTarefa gets a reference to the given NullableInt64 and assigns it to the RefTarefa field.
func (o *Intervencao) SetRefTarefa(v int64) {
	o.RefTarefa.Set(&v)
}
// SetRefTarefaNil sets the value for RefTarefa to be an explicit nil
func (o *Intervencao) SetRefTarefaNil() {
	o.RefTarefa.Set(nil)
}

// UnsetRefTarefa ensures that no value is present for RefTarefa, not even an explicit nil
func (o *Intervencao) UnsetRefTarefa() {
	o.RefTarefa.Unset()
}

// GetDescricao returns the Descricao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetDescricao() string {
	if o == nil || o.Descricao.Get() == nil {
		var ret string
		return ret
	}
	return *o.Descricao.Get()
}

// GetDescricaoOk returns a tuple with the Descricao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetDescricaoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Descricao.Get(), o.Descricao.IsSet()
}

// HasDescricao returns a boolean if a field has been set.
func (o *Intervencao) HasDescricao() bool {
	if o != nil && o.Descricao.IsSet() {
		return true
	}

	return false
}

// SetDescricao gets a reference to the given NullableString and assigns it to the Descricao field.
func (o *Intervencao) SetDescricao(v string) {
	o.Descricao.Set(&v)
}
// SetDescricaoNil sets the value for Descricao to be an explicit nil
func (o *Intervencao) SetDescricaoNil() {
	o.Descricao.Set(nil)
}

// UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
func (o *Intervencao) UnsetDescricao() {
	o.Descricao.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetData() time.Time {
	if o == nil || o.Data.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetDataOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *Intervencao) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *Intervencao) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *Intervencao) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *Intervencao) UnsetData() {
	o.Data.Unset()
}

// GetRefUtilizador returns the RefUtilizador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetRefUtilizador() int64 {
	if o == nil || o.RefUtilizador.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefUtilizador.Get()
}

// GetRefUtilizadorOk returns a tuple with the RefUtilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetRefUtilizadorOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefUtilizador.Get(), o.RefUtilizador.IsSet()
}

// HasRefUtilizador returns a boolean if a field has been set.
func (o *Intervencao) HasRefUtilizador() bool {
	if o != nil && o.RefUtilizador.IsSet() {
		return true
	}

	return false
}

// SetRefUtilizador gets a reference to the given NullableInt64 and assigns it to the RefUtilizador field.
func (o *Intervencao) SetRefUtilizador(v int64) {
	o.RefUtilizador.Set(&v)
}
// SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil
func (o *Intervencao) SetRefUtilizadorNil() {
	o.RefUtilizador.Set(nil)
}

// UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
func (o *Intervencao) UnsetRefUtilizador() {
	o.RefUtilizador.Unset()
}

// GetEstado returns the Estado field value if set, zero value otherwise.
func (o *Intervencao) GetEstado() IntervencaoEstados {
	if o == nil || o.Estado == nil {
		var ret IntervencaoEstados
		return ret
	}
	return *o.Estado
}

// GetEstadoOk returns a tuple with the Estado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intervencao) GetEstadoOk() (*IntervencaoEstados, bool) {
	if o == nil || o.Estado == nil {
		return nil, false
	}
	return o.Estado, true
}

// HasEstado returns a boolean if a field has been set.
func (o *Intervencao) HasEstado() bool {
	if o != nil && o.Estado != nil {
		return true
	}

	return false
}

// SetEstado gets a reference to the given IntervencaoEstados and assigns it to the Estado field.
func (o *Intervencao) SetEstado(v IntervencaoEstados) {
	o.Estado = &v
}

// GetNomeUtilizador returns the NomeUtilizador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Intervencao) GetNomeUtilizador() string {
	if o == nil || o.NomeUtilizador.Get() == nil {
		var ret string
		return ret
	}
	return *o.NomeUtilizador.Get()
}

// GetNomeUtilizadorOk returns a tuple with the NomeUtilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Intervencao) GetNomeUtilizadorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NomeUtilizador.Get(), o.NomeUtilizador.IsSet()
}

// HasNomeUtilizador returns a boolean if a field has been set.
func (o *Intervencao) HasNomeUtilizador() bool {
	if o != nil && o.NomeUtilizador.IsSet() {
		return true
	}

	return false
}

// SetNomeUtilizador gets a reference to the given NullableString and assigns it to the NomeUtilizador field.
func (o *Intervencao) SetNomeUtilizador(v string) {
	o.NomeUtilizador.Set(&v)
}
// SetNomeUtilizadorNil sets the value for NomeUtilizador to be an explicit nil
func (o *Intervencao) SetNomeUtilizadorNil() {
	o.NomeUtilizador.Set(nil)
}

// UnsetNomeUtilizador ensures that no value is present for NomeUtilizador, not even an explicit nil
func (o *Intervencao) UnsetNomeUtilizador() {
	o.NomeUtilizador.Unset()
}

// GetRefTarefaNavigation returns the RefTarefaNavigation field value if set, zero value otherwise.
func (o *Intervencao) GetRefTarefaNavigation() Tarefa {
	if o == nil || o.RefTarefaNavigation == nil {
		var ret Tarefa
		return ret
	}
	return *o.RefTarefaNavigation
}

// GetRefTarefaNavigationOk returns a tuple with the RefTarefaNavigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intervencao) GetRefTarefaNavigationOk() (*Tarefa, bool) {
	if o == nil || o.RefTarefaNavigation == nil {
		return nil, false
	}
	return o.RefTarefaNavigation, true
}

// HasRefTarefaNavigation returns a boolean if a field has been set.
func (o *Intervencao) HasRefTarefaNavigation() bool {
	if o != nil && o.RefTarefaNavigation != nil {
		return true
	}

	return false
}

// SetRefTarefaNavigation gets a reference to the given Tarefa and assigns it to the RefTarefaNavigation field.
func (o *Intervencao) SetRefTarefaNavigation(v Tarefa) {
	o.RefTarefaNavigation = &v
}

// GetRefUtilizadorNavigation returns the RefUtilizadorNavigation field value if set, zero value otherwise.
func (o *Intervencao) GetRefUtilizadorNavigation() Utilizador {
	if o == nil || o.RefUtilizadorNavigation == nil {
		var ret Utilizador
		return ret
	}
	return *o.RefUtilizadorNavigation
}

// GetRefUtilizadorNavigationOk returns a tuple with the RefUtilizadorNavigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Intervencao) GetRefUtilizadorNavigationOk() (*Utilizador, bool) {
	if o == nil || o.RefUtilizadorNavigation == nil {
		return nil, false
	}
	return o.RefUtilizadorNavigation, true
}

// HasRefUtilizadorNavigation returns a boolean if a field has been set.
func (o *Intervencao) HasRefUtilizadorNavigation() bool {
	if o != nil && o.RefUtilizadorNavigation != nil {
		return true
	}

	return false
}

// SetRefUtilizadorNavigation gets a reference to the given Utilizador and assigns it to the RefUtilizadorNavigation field.
func (o *Intervencao) SetRefUtilizadorNavigation(v Utilizador) {
	o.RefUtilizadorNavigation = &v
}

func (o Intervencao) MarshalJSON() ([]byte, error) {
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
	if o.RefTarefa.IsSet() {
		toSerialize["refTarefa"] = o.RefTarefa.Get()
	}
	if o.Descricao.IsSet() {
		toSerialize["descricao"] = o.Descricao.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.RefUtilizador.IsSet() {
		toSerialize["refUtilizador"] = o.RefUtilizador.Get()
	}
	if o.Estado != nil {
		toSerialize["estado"] = o.Estado
	}
	if o.NomeUtilizador.IsSet() {
		toSerialize["nomeUtilizador"] = o.NomeUtilizador.Get()
	}
	if o.RefTarefaNavigation != nil {
		toSerialize["refTarefaNavigation"] = o.RefTarefaNavigation
	}
	if o.RefUtilizadorNavigation != nil {
		toSerialize["refUtilizadorNavigation"] = o.RefUtilizadorNavigation
	}
	return json.Marshal(toSerialize)
}

type NullableIntervencao struct {
	value *Intervencao
	isSet bool
}

func (v NullableIntervencao) Get() *Intervencao {
	return v.value
}

func (v *NullableIntervencao) Set(val *Intervencao) {
	v.value = val
	v.isSet = true
}

func (v NullableIntervencao) IsSet() bool {
	return v.isSet
}

func (v *NullableIntervencao) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntervencao(val *Intervencao) *NullableIntervencao {
	return &NullableIntervencao{value: val, isSet: true}
}

func (v NullableIntervencao) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntervencao) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

