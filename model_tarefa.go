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

// Tarefa struct for Tarefa
type Tarefa struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	Estado *IntervencaoEstados `json:"estado,omitempty"`
	Tipo *TipoTarefa `json:"tipo,omitempty"`
	Assunto NullableString `json:"assunto,omitempty"`
	Descricao NullableString `json:"descricao,omitempty"`
	RefUtilizador NullableInt64 `json:"refUtilizador,omitempty"`
	CriadoEm NullableTime `json:"criadoEm,omitempty"`
	FechadoEm NullableTime `json:"fechadoEm,omitempty"`
	RefEntidade NullableInt64 `json:"refEntidade,omitempty"`
	RefGvm NullableInt64 `json:"refGvm,omitempty"`
	NomeUtilizador NullableString `json:"nomeUtilizador,omitempty"`
	Utilizador *Utilizador `json:"utilizador,omitempty"`
	Intervencaos []Intervencao `json:"intervencaos,omitempty"`
}

// NewTarefa instantiates a new Tarefa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTarefa() *Tarefa {
	this := Tarefa{}
	return &this
}

// NewTarefaWithDefaults instantiates a new Tarefa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTarefaWithDefaults() *Tarefa {
	this := Tarefa{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tarefa) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tarefa) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tarefa) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Tarefa) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Tarefa) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tarefa) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Tarefa) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Tarefa) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Tarefa) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Tarefa) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Tarefa) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Tarefa) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Tarefa) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Tarefa) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Tarefa) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Tarefa) UnsetUsr() {
	o.Usr.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetData() time.Time {
	if o == nil || o.Data.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetDataOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *Tarefa) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *Tarefa) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *Tarefa) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *Tarefa) UnsetData() {
	o.Data.Unset()
}

// GetEstado returns the Estado field value if set, zero value otherwise.
func (o *Tarefa) GetEstado() IntervencaoEstados {
	if o == nil || o.Estado == nil {
		var ret IntervencaoEstados
		return ret
	}
	return *o.Estado
}

// GetEstadoOk returns a tuple with the Estado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tarefa) GetEstadoOk() (*IntervencaoEstados, bool) {
	if o == nil || o.Estado == nil {
		return nil, false
	}
	return o.Estado, true
}

// HasEstado returns a boolean if a field has been set.
func (o *Tarefa) HasEstado() bool {
	if o != nil && o.Estado != nil {
		return true
	}

	return false
}

// SetEstado gets a reference to the given IntervencaoEstados and assigns it to the Estado field.
func (o *Tarefa) SetEstado(v IntervencaoEstados) {
	o.Estado = &v
}

// GetTipo returns the Tipo field value if set, zero value otherwise.
func (o *Tarefa) GetTipo() TipoTarefa {
	if o == nil || o.Tipo == nil {
		var ret TipoTarefa
		return ret
	}
	return *o.Tipo
}

// GetTipoOk returns a tuple with the Tipo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tarefa) GetTipoOk() (*TipoTarefa, bool) {
	if o == nil || o.Tipo == nil {
		return nil, false
	}
	return o.Tipo, true
}

// HasTipo returns a boolean if a field has been set.
func (o *Tarefa) HasTipo() bool {
	if o != nil && o.Tipo != nil {
		return true
	}

	return false
}

// SetTipo gets a reference to the given TipoTarefa and assigns it to the Tipo field.
func (o *Tarefa) SetTipo(v TipoTarefa) {
	o.Tipo = &v
}

// GetAssunto returns the Assunto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetAssunto() string {
	if o == nil || o.Assunto.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assunto.Get()
}

// GetAssuntoOk returns a tuple with the Assunto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetAssuntoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assunto.Get(), o.Assunto.IsSet()
}

// HasAssunto returns a boolean if a field has been set.
func (o *Tarefa) HasAssunto() bool {
	if o != nil && o.Assunto.IsSet() {
		return true
	}

	return false
}

// SetAssunto gets a reference to the given NullableString and assigns it to the Assunto field.
func (o *Tarefa) SetAssunto(v string) {
	o.Assunto.Set(&v)
}
// SetAssuntoNil sets the value for Assunto to be an explicit nil
func (o *Tarefa) SetAssuntoNil() {
	o.Assunto.Set(nil)
}

// UnsetAssunto ensures that no value is present for Assunto, not even an explicit nil
func (o *Tarefa) UnsetAssunto() {
	o.Assunto.Unset()
}

// GetDescricao returns the Descricao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetDescricao() string {
	if o == nil || o.Descricao.Get() == nil {
		var ret string
		return ret
	}
	return *o.Descricao.Get()
}

// GetDescricaoOk returns a tuple with the Descricao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetDescricaoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Descricao.Get(), o.Descricao.IsSet()
}

// HasDescricao returns a boolean if a field has been set.
func (o *Tarefa) HasDescricao() bool {
	if o != nil && o.Descricao.IsSet() {
		return true
	}

	return false
}

// SetDescricao gets a reference to the given NullableString and assigns it to the Descricao field.
func (o *Tarefa) SetDescricao(v string) {
	o.Descricao.Set(&v)
}
// SetDescricaoNil sets the value for Descricao to be an explicit nil
func (o *Tarefa) SetDescricaoNil() {
	o.Descricao.Set(nil)
}

// UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
func (o *Tarefa) UnsetDescricao() {
	o.Descricao.Unset()
}

// GetRefUtilizador returns the RefUtilizador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetRefUtilizador() int64 {
	if o == nil || o.RefUtilizador.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefUtilizador.Get()
}

// GetRefUtilizadorOk returns a tuple with the RefUtilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetRefUtilizadorOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefUtilizador.Get(), o.RefUtilizador.IsSet()
}

// HasRefUtilizador returns a boolean if a field has been set.
func (o *Tarefa) HasRefUtilizador() bool {
	if o != nil && o.RefUtilizador.IsSet() {
		return true
	}

	return false
}

// SetRefUtilizador gets a reference to the given NullableInt64 and assigns it to the RefUtilizador field.
func (o *Tarefa) SetRefUtilizador(v int64) {
	o.RefUtilizador.Set(&v)
}
// SetRefUtilizadorNil sets the value for RefUtilizador to be an explicit nil
func (o *Tarefa) SetRefUtilizadorNil() {
	o.RefUtilizador.Set(nil)
}

// UnsetRefUtilizador ensures that no value is present for RefUtilizador, not even an explicit nil
func (o *Tarefa) UnsetRefUtilizador() {
	o.RefUtilizador.Unset()
}

// GetCriadoEm returns the CriadoEm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetCriadoEm() time.Time {
	if o == nil || o.CriadoEm.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CriadoEm.Get()
}

// GetCriadoEmOk returns a tuple with the CriadoEm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetCriadoEmOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CriadoEm.Get(), o.CriadoEm.IsSet()
}

// HasCriadoEm returns a boolean if a field has been set.
func (o *Tarefa) HasCriadoEm() bool {
	if o != nil && o.CriadoEm.IsSet() {
		return true
	}

	return false
}

// SetCriadoEm gets a reference to the given NullableTime and assigns it to the CriadoEm field.
func (o *Tarefa) SetCriadoEm(v time.Time) {
	o.CriadoEm.Set(&v)
}
// SetCriadoEmNil sets the value for CriadoEm to be an explicit nil
func (o *Tarefa) SetCriadoEmNil() {
	o.CriadoEm.Set(nil)
}

// UnsetCriadoEm ensures that no value is present for CriadoEm, not even an explicit nil
func (o *Tarefa) UnsetCriadoEm() {
	o.CriadoEm.Unset()
}

// GetFechadoEm returns the FechadoEm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetFechadoEm() time.Time {
	if o == nil || o.FechadoEm.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FechadoEm.Get()
}

// GetFechadoEmOk returns a tuple with the FechadoEm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetFechadoEmOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FechadoEm.Get(), o.FechadoEm.IsSet()
}

// HasFechadoEm returns a boolean if a field has been set.
func (o *Tarefa) HasFechadoEm() bool {
	if o != nil && o.FechadoEm.IsSet() {
		return true
	}

	return false
}

// SetFechadoEm gets a reference to the given NullableTime and assigns it to the FechadoEm field.
func (o *Tarefa) SetFechadoEm(v time.Time) {
	o.FechadoEm.Set(&v)
}
// SetFechadoEmNil sets the value for FechadoEm to be an explicit nil
func (o *Tarefa) SetFechadoEmNil() {
	o.FechadoEm.Set(nil)
}

// UnsetFechadoEm ensures that no value is present for FechadoEm, not even an explicit nil
func (o *Tarefa) UnsetFechadoEm() {
	o.FechadoEm.Unset()
}

// GetRefEntidade returns the RefEntidade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetRefEntidade() int64 {
	if o == nil || o.RefEntidade.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefEntidade.Get()
}

// GetRefEntidadeOk returns a tuple with the RefEntidade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetRefEntidadeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefEntidade.Get(), o.RefEntidade.IsSet()
}

// HasRefEntidade returns a boolean if a field has been set.
func (o *Tarefa) HasRefEntidade() bool {
	if o != nil && o.RefEntidade.IsSet() {
		return true
	}

	return false
}

// SetRefEntidade gets a reference to the given NullableInt64 and assigns it to the RefEntidade field.
func (o *Tarefa) SetRefEntidade(v int64) {
	o.RefEntidade.Set(&v)
}
// SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil
func (o *Tarefa) SetRefEntidadeNil() {
	o.RefEntidade.Set(nil)
}

// UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
func (o *Tarefa) UnsetRefEntidade() {
	o.RefEntidade.Unset()
}

// GetRefGvm returns the RefGvm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetRefGvm() int64 {
	if o == nil || o.RefGvm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefGvm.Get()
}

// GetRefGvmOk returns a tuple with the RefGvm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetRefGvmOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefGvm.Get(), o.RefGvm.IsSet()
}

// HasRefGvm returns a boolean if a field has been set.
func (o *Tarefa) HasRefGvm() bool {
	if o != nil && o.RefGvm.IsSet() {
		return true
	}

	return false
}

// SetRefGvm gets a reference to the given NullableInt64 and assigns it to the RefGvm field.
func (o *Tarefa) SetRefGvm(v int64) {
	o.RefGvm.Set(&v)
}
// SetRefGvmNil sets the value for RefGvm to be an explicit nil
func (o *Tarefa) SetRefGvmNil() {
	o.RefGvm.Set(nil)
}

// UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
func (o *Tarefa) UnsetRefGvm() {
	o.RefGvm.Unset()
}

// GetNomeUtilizador returns the NomeUtilizador field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetNomeUtilizador() string {
	if o == nil || o.NomeUtilizador.Get() == nil {
		var ret string
		return ret
	}
	return *o.NomeUtilizador.Get()
}

// GetNomeUtilizadorOk returns a tuple with the NomeUtilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetNomeUtilizadorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NomeUtilizador.Get(), o.NomeUtilizador.IsSet()
}

// HasNomeUtilizador returns a boolean if a field has been set.
func (o *Tarefa) HasNomeUtilizador() bool {
	if o != nil && o.NomeUtilizador.IsSet() {
		return true
	}

	return false
}

// SetNomeUtilizador gets a reference to the given NullableString and assigns it to the NomeUtilizador field.
func (o *Tarefa) SetNomeUtilizador(v string) {
	o.NomeUtilizador.Set(&v)
}
// SetNomeUtilizadorNil sets the value for NomeUtilizador to be an explicit nil
func (o *Tarefa) SetNomeUtilizadorNil() {
	o.NomeUtilizador.Set(nil)
}

// UnsetNomeUtilizador ensures that no value is present for NomeUtilizador, not even an explicit nil
func (o *Tarefa) UnsetNomeUtilizador() {
	o.NomeUtilizador.Unset()
}

// GetUtilizador returns the Utilizador field value if set, zero value otherwise.
func (o *Tarefa) GetUtilizador() Utilizador {
	if o == nil || o.Utilizador == nil {
		var ret Utilizador
		return ret
	}
	return *o.Utilizador
}

// GetUtilizadorOk returns a tuple with the Utilizador field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tarefa) GetUtilizadorOk() (*Utilizador, bool) {
	if o == nil || o.Utilizador == nil {
		return nil, false
	}
	return o.Utilizador, true
}

// HasUtilizador returns a boolean if a field has been set.
func (o *Tarefa) HasUtilizador() bool {
	if o != nil && o.Utilizador != nil {
		return true
	}

	return false
}

// SetUtilizador gets a reference to the given Utilizador and assigns it to the Utilizador field.
func (o *Tarefa) SetUtilizador(v Utilizador) {
	o.Utilizador = &v
}

// GetIntervencaos returns the Intervencaos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tarefa) GetIntervencaos() []Intervencao {
	if o == nil  {
		var ret []Intervencao
		return ret
	}
	return o.Intervencaos
}

// GetIntervencaosOk returns a tuple with the Intervencaos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tarefa) GetIntervencaosOk() (*[]Intervencao, bool) {
	if o == nil || o.Intervencaos == nil {
		return nil, false
	}
	return &o.Intervencaos, true
}

// HasIntervencaos returns a boolean if a field has been set.
func (o *Tarefa) HasIntervencaos() bool {
	if o != nil && o.Intervencaos != nil {
		return true
	}

	return false
}

// SetIntervencaos gets a reference to the given []Intervencao and assigns it to the Intervencaos field.
func (o *Tarefa) SetIntervencaos(v []Intervencao) {
	o.Intervencaos = v
}

func (o Tarefa) MarshalJSON() ([]byte, error) {
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
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Estado != nil {
		toSerialize["estado"] = o.Estado
	}
	if o.Tipo != nil {
		toSerialize["tipo"] = o.Tipo
	}
	if o.Assunto.IsSet() {
		toSerialize["assunto"] = o.Assunto.Get()
	}
	if o.Descricao.IsSet() {
		toSerialize["descricao"] = o.Descricao.Get()
	}
	if o.RefUtilizador.IsSet() {
		toSerialize["refUtilizador"] = o.RefUtilizador.Get()
	}
	if o.CriadoEm.IsSet() {
		toSerialize["criadoEm"] = o.CriadoEm.Get()
	}
	if o.FechadoEm.IsSet() {
		toSerialize["fechadoEm"] = o.FechadoEm.Get()
	}
	if o.RefEntidade.IsSet() {
		toSerialize["refEntidade"] = o.RefEntidade.Get()
	}
	if o.RefGvm.IsSet() {
		toSerialize["refGvm"] = o.RefGvm.Get()
	}
	if o.NomeUtilizador.IsSet() {
		toSerialize["nomeUtilizador"] = o.NomeUtilizador.Get()
	}
	if o.Utilizador != nil {
		toSerialize["utilizador"] = o.Utilizador
	}
	if o.Intervencaos != nil {
		toSerialize["intervencaos"] = o.Intervencaos
	}
	return json.Marshal(toSerialize)
}

type NullableTarefa struct {
	value *Tarefa
	isSet bool
}

func (v NullableTarefa) Get() *Tarefa {
	return v.value
}

func (v *NullableTarefa) Set(val *Tarefa) {
	v.value = val
	v.isSet = true
}

func (v NullableTarefa) IsSet() bool {
	return v.isSet
}

func (v *NullableTarefa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTarefa(val *Tarefa) *NullableTarefa {
	return &NullableTarefa{value: val, isSet: true}
}

func (v NullableTarefa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTarefa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


