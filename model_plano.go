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

// Plano struct for Plano
type Plano struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	TaxaUtilizacao NullableInt32 `json:"taxaUtilizacao,omitempty"`
	Periodicidade NullableInt32 `json:"periodicidade,omitempty"`
	Entidades []Entidade `json:"entidades,omitempty"`
}

// NewPlano instantiates a new Plano object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlano() *Plano {
	this := Plano{}
	return &this
}

// NewPlanoWithDefaults instantiates a new Plano object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanoWithDefaults() *Plano {
	this := Plano{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Plano) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plano) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Plano) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Plano) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Plano) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plano) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Plano) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Plano) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Plano) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Plano) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Plano) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Plano) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Plano) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Plano) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Plano) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Plano) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *Plano) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *Plano) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *Plano) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *Plano) UnsetNome() {
	o.Nome.Unset()
}

// GetTaxaUtilizacao returns the TaxaUtilizacao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetTaxaUtilizacao() int32 {
	if o == nil || o.TaxaUtilizacao.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TaxaUtilizacao.Get()
}

// GetTaxaUtilizacaoOk returns a tuple with the TaxaUtilizacao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetTaxaUtilizacaoOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxaUtilizacao.Get(), o.TaxaUtilizacao.IsSet()
}

// HasTaxaUtilizacao returns a boolean if a field has been set.
func (o *Plano) HasTaxaUtilizacao() bool {
	if o != nil && o.TaxaUtilizacao.IsSet() {
		return true
	}

	return false
}

// SetTaxaUtilizacao gets a reference to the given NullableInt32 and assigns it to the TaxaUtilizacao field.
func (o *Plano) SetTaxaUtilizacao(v int32) {
	o.TaxaUtilizacao.Set(&v)
}
// SetTaxaUtilizacaoNil sets the value for TaxaUtilizacao to be an explicit nil
func (o *Plano) SetTaxaUtilizacaoNil() {
	o.TaxaUtilizacao.Set(nil)
}

// UnsetTaxaUtilizacao ensures that no value is present for TaxaUtilizacao, not even an explicit nil
func (o *Plano) UnsetTaxaUtilizacao() {
	o.TaxaUtilizacao.Unset()
}

// GetPeriodicidade returns the Periodicidade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetPeriodicidade() int32 {
	if o == nil || o.Periodicidade.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Periodicidade.Get()
}

// GetPeriodicidadeOk returns a tuple with the Periodicidade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetPeriodicidadeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Periodicidade.Get(), o.Periodicidade.IsSet()
}

// HasPeriodicidade returns a boolean if a field has been set.
func (o *Plano) HasPeriodicidade() bool {
	if o != nil && o.Periodicidade.IsSet() {
		return true
	}

	return false
}

// SetPeriodicidade gets a reference to the given NullableInt32 and assigns it to the Periodicidade field.
func (o *Plano) SetPeriodicidade(v int32) {
	o.Periodicidade.Set(&v)
}
// SetPeriodicidadeNil sets the value for Periodicidade to be an explicit nil
func (o *Plano) SetPeriodicidadeNil() {
	o.Periodicidade.Set(nil)
}

// UnsetPeriodicidade ensures that no value is present for Periodicidade, not even an explicit nil
func (o *Plano) UnsetPeriodicidade() {
	o.Periodicidade.Unset()
}

// GetEntidades returns the Entidades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plano) GetEntidades() []Entidade {
	if o == nil  {
		var ret []Entidade
		return ret
	}
	return o.Entidades
}

// GetEntidadesOk returns a tuple with the Entidades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plano) GetEntidadesOk() (*[]Entidade, bool) {
	if o == nil || o.Entidades == nil {
		return nil, false
	}
	return &o.Entidades, true
}

// HasEntidades returns a boolean if a field has been set.
func (o *Plano) HasEntidades() bool {
	if o != nil && o.Entidades != nil {
		return true
	}

	return false
}

// SetEntidades gets a reference to the given []Entidade and assigns it to the Entidades field.
func (o *Plano) SetEntidades(v []Entidade) {
	o.Entidades = v
}

func (o Plano) MarshalJSON() ([]byte, error) {
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
	if o.TaxaUtilizacao.IsSet() {
		toSerialize["taxaUtilizacao"] = o.TaxaUtilizacao.Get()
	}
	if o.Periodicidade.IsSet() {
		toSerialize["periodicidade"] = o.Periodicidade.Get()
	}
	if o.Entidades != nil {
		toSerialize["entidades"] = o.Entidades
	}
	return json.Marshal(toSerialize)
}

type NullablePlano struct {
	value *Plano
	isSet bool
}

func (v NullablePlano) Get() *Plano {
	return v.value
}

func (v *NullablePlano) Set(val *Plano) {
	v.value = val
	v.isSet = true
}

func (v NullablePlano) IsSet() bool {
	return v.isSet
}

func (v *NullablePlano) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlano(val *Plano) *NullablePlano {
	return &NullablePlano{value: val, isSet: true}
}

func (v NullablePlano) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlano) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


