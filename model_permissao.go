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

// Permissao struct for Permissao
type Permissao struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	Descricao NullableString `json:"descricao,omitempty"`
	Chave NullableString `json:"chave,omitempty"`
	IsGroup *bool `json:"isGroup,omitempty"`
	RefPermissao NullableInt64 `json:"refPermissao,omitempty"`
	PerfilSegurancaPermissaos []PerfilSegurancaPermissao `json:"perfilSegurancaPermissaos,omitempty"`
}

// NewPermissao instantiates a new Permissao object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissao() *Permissao {
	this := Permissao{}
	return &this
}

// NewPermissaoWithDefaults instantiates a new Permissao object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissaoWithDefaults() *Permissao {
	this := Permissao{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Permissao) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissao) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Permissao) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Permissao) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Permissao) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissao) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Permissao) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Permissao) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Permissao) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Permissao) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Permissao) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Permissao) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Permissao) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Permissao) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Permissao) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Permissao) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *Permissao) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *Permissao) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *Permissao) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *Permissao) UnsetNome() {
	o.Nome.Unset()
}

// GetDescricao returns the Descricao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetDescricao() string {
	if o == nil || o.Descricao.Get() == nil {
		var ret string
		return ret
	}
	return *o.Descricao.Get()
}

// GetDescricaoOk returns a tuple with the Descricao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetDescricaoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Descricao.Get(), o.Descricao.IsSet()
}

// HasDescricao returns a boolean if a field has been set.
func (o *Permissao) HasDescricao() bool {
	if o != nil && o.Descricao.IsSet() {
		return true
	}

	return false
}

// SetDescricao gets a reference to the given NullableString and assigns it to the Descricao field.
func (o *Permissao) SetDescricao(v string) {
	o.Descricao.Set(&v)
}
// SetDescricaoNil sets the value for Descricao to be an explicit nil
func (o *Permissao) SetDescricaoNil() {
	o.Descricao.Set(nil)
}

// UnsetDescricao ensures that no value is present for Descricao, not even an explicit nil
func (o *Permissao) UnsetDescricao() {
	o.Descricao.Unset()
}

// GetChave returns the Chave field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetChave() string {
	if o == nil || o.Chave.Get() == nil {
		var ret string
		return ret
	}
	return *o.Chave.Get()
}

// GetChaveOk returns a tuple with the Chave field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetChaveOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Chave.Get(), o.Chave.IsSet()
}

// HasChave returns a boolean if a field has been set.
func (o *Permissao) HasChave() bool {
	if o != nil && o.Chave.IsSet() {
		return true
	}

	return false
}

// SetChave gets a reference to the given NullableString and assigns it to the Chave field.
func (o *Permissao) SetChave(v string) {
	o.Chave.Set(&v)
}
// SetChaveNil sets the value for Chave to be an explicit nil
func (o *Permissao) SetChaveNil() {
	o.Chave.Set(nil)
}

// UnsetChave ensures that no value is present for Chave, not even an explicit nil
func (o *Permissao) UnsetChave() {
	o.Chave.Unset()
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise.
func (o *Permissao) GetIsGroup() bool {
	if o == nil || o.IsGroup == nil {
		var ret bool
		return ret
	}
	return *o.IsGroup
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissao) GetIsGroupOk() (*bool, bool) {
	if o == nil || o.IsGroup == nil {
		return nil, false
	}
	return o.IsGroup, true
}

// HasIsGroup returns a boolean if a field has been set.
func (o *Permissao) HasIsGroup() bool {
	if o != nil && o.IsGroup != nil {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given bool and assigns it to the IsGroup field.
func (o *Permissao) SetIsGroup(v bool) {
	o.IsGroup = &v
}

// GetRefPermissao returns the RefPermissao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetRefPermissao() int64 {
	if o == nil || o.RefPermissao.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefPermissao.Get()
}

// GetRefPermissaoOk returns a tuple with the RefPermissao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetRefPermissaoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefPermissao.Get(), o.RefPermissao.IsSet()
}

// HasRefPermissao returns a boolean if a field has been set.
func (o *Permissao) HasRefPermissao() bool {
	if o != nil && o.RefPermissao.IsSet() {
		return true
	}

	return false
}

// SetRefPermissao gets a reference to the given NullableInt64 and assigns it to the RefPermissao field.
func (o *Permissao) SetRefPermissao(v int64) {
	o.RefPermissao.Set(&v)
}
// SetRefPermissaoNil sets the value for RefPermissao to be an explicit nil
func (o *Permissao) SetRefPermissaoNil() {
	o.RefPermissao.Set(nil)
}

// UnsetRefPermissao ensures that no value is present for RefPermissao, not even an explicit nil
func (o *Permissao) UnsetRefPermissao() {
	o.RefPermissao.Unset()
}

// GetPerfilSegurancaPermissaos returns the PerfilSegurancaPermissaos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permissao) GetPerfilSegurancaPermissaos() []PerfilSegurancaPermissao {
	if o == nil  {
		var ret []PerfilSegurancaPermissao
		return ret
	}
	return o.PerfilSegurancaPermissaos
}

// GetPerfilSegurancaPermissaosOk returns a tuple with the PerfilSegurancaPermissaos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permissao) GetPerfilSegurancaPermissaosOk() (*[]PerfilSegurancaPermissao, bool) {
	if o == nil || o.PerfilSegurancaPermissaos == nil {
		return nil, false
	}
	return &o.PerfilSegurancaPermissaos, true
}

// HasPerfilSegurancaPermissaos returns a boolean if a field has been set.
func (o *Permissao) HasPerfilSegurancaPermissaos() bool {
	if o != nil && o.PerfilSegurancaPermissaos != nil {
		return true
	}

	return false
}

// SetPerfilSegurancaPermissaos gets a reference to the given []PerfilSegurancaPermissao and assigns it to the PerfilSegurancaPermissaos field.
func (o *Permissao) SetPerfilSegurancaPermissaos(v []PerfilSegurancaPermissao) {
	o.PerfilSegurancaPermissaos = v
}

func (o Permissao) MarshalJSON() ([]byte, error) {
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
	if o.Descricao.IsSet() {
		toSerialize["descricao"] = o.Descricao.Get()
	}
	if o.Chave.IsSet() {
		toSerialize["chave"] = o.Chave.Get()
	}
	if o.IsGroup != nil {
		toSerialize["isGroup"] = o.IsGroup
	}
	if o.RefPermissao.IsSet() {
		toSerialize["refPermissao"] = o.RefPermissao.Get()
	}
	if o.PerfilSegurancaPermissaos != nil {
		toSerialize["perfilSegurancaPermissaos"] = o.PerfilSegurancaPermissaos
	}
	return json.Marshal(toSerialize)
}

type NullablePermissao struct {
	value *Permissao
	isSet bool
}

func (v NullablePermissao) Get() *Permissao {
	return v.value
}

func (v *NullablePermissao) Set(val *Permissao) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissao) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissao) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissao(val *Permissao) *NullablePermissao {
	return &NullablePermissao{value: val, isSet: true}
}

func (v NullablePermissao) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissao) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

