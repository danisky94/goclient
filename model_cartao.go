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

// Cartao struct for Cartao
type Cartao struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	NumeroSerie NullableString `json:"numeroSerie,omitempty"`
	DataProducao NullableTime `json:"dataProducao,omitempty"`
	Estado NullableInt32 `json:"estado,omitempty"`
	DataEstado NullableTime `json:"dataEstado,omitempty"`
	Pin NullableString `json:"pin,omitempty"`
	CartaoProdutos []CartaoProduto `json:"cartaoProdutos,omitempty"`
}

// NewCartao instantiates a new Cartao object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartao() *Cartao {
	this := Cartao{}
	return &this
}

// NewCartaoWithDefaults instantiates a new Cartao object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartaoWithDefaults() *Cartao {
	this := Cartao{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cartao) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cartao) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cartao) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Cartao) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Cartao) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cartao) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Cartao) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Cartao) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Cartao) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Cartao) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Cartao) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Cartao) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Cartao) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Cartao) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Cartao) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Cartao) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *Cartao) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *Cartao) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *Cartao) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *Cartao) UnsetNome() {
	o.Nome.Unset()
}

// GetNumeroSerie returns the NumeroSerie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetNumeroSerie() string {
	if o == nil || o.NumeroSerie.Get() == nil {
		var ret string
		return ret
	}
	return *o.NumeroSerie.Get()
}

// GetNumeroSerieOk returns a tuple with the NumeroSerie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetNumeroSerieOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NumeroSerie.Get(), o.NumeroSerie.IsSet()
}

// HasNumeroSerie returns a boolean if a field has been set.
func (o *Cartao) HasNumeroSerie() bool {
	if o != nil && o.NumeroSerie.IsSet() {
		return true
	}

	return false
}

// SetNumeroSerie gets a reference to the given NullableString and assigns it to the NumeroSerie field.
func (o *Cartao) SetNumeroSerie(v string) {
	o.NumeroSerie.Set(&v)
}
// SetNumeroSerieNil sets the value for NumeroSerie to be an explicit nil
func (o *Cartao) SetNumeroSerieNil() {
	o.NumeroSerie.Set(nil)
}

// UnsetNumeroSerie ensures that no value is present for NumeroSerie, not even an explicit nil
func (o *Cartao) UnsetNumeroSerie() {
	o.NumeroSerie.Unset()
}

// GetDataProducao returns the DataProducao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetDataProducao() time.Time {
	if o == nil || o.DataProducao.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataProducao.Get()
}

// GetDataProducaoOk returns a tuple with the DataProducao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetDataProducaoOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataProducao.Get(), o.DataProducao.IsSet()
}

// HasDataProducao returns a boolean if a field has been set.
func (o *Cartao) HasDataProducao() bool {
	if o != nil && o.DataProducao.IsSet() {
		return true
	}

	return false
}

// SetDataProducao gets a reference to the given NullableTime and assigns it to the DataProducao field.
func (o *Cartao) SetDataProducao(v time.Time) {
	o.DataProducao.Set(&v)
}
// SetDataProducaoNil sets the value for DataProducao to be an explicit nil
func (o *Cartao) SetDataProducaoNil() {
	o.DataProducao.Set(nil)
}

// UnsetDataProducao ensures that no value is present for DataProducao, not even an explicit nil
func (o *Cartao) UnsetDataProducao() {
	o.DataProducao.Unset()
}

// GetEstado returns the Estado field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetEstado() int32 {
	if o == nil || o.Estado.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Estado.Get()
}

// GetEstadoOk returns a tuple with the Estado field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetEstadoOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Estado.Get(), o.Estado.IsSet()
}

// HasEstado returns a boolean if a field has been set.
func (o *Cartao) HasEstado() bool {
	if o != nil && o.Estado.IsSet() {
		return true
	}

	return false
}

// SetEstado gets a reference to the given NullableInt32 and assigns it to the Estado field.
func (o *Cartao) SetEstado(v int32) {
	o.Estado.Set(&v)
}
// SetEstadoNil sets the value for Estado to be an explicit nil
func (o *Cartao) SetEstadoNil() {
	o.Estado.Set(nil)
}

// UnsetEstado ensures that no value is present for Estado, not even an explicit nil
func (o *Cartao) UnsetEstado() {
	o.Estado.Unset()
}

// GetDataEstado returns the DataEstado field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetDataEstado() time.Time {
	if o == nil || o.DataEstado.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataEstado.Get()
}

// GetDataEstadoOk returns a tuple with the DataEstado field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetDataEstadoOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataEstado.Get(), o.DataEstado.IsSet()
}

// HasDataEstado returns a boolean if a field has been set.
func (o *Cartao) HasDataEstado() bool {
	if o != nil && o.DataEstado.IsSet() {
		return true
	}

	return false
}

// SetDataEstado gets a reference to the given NullableTime and assigns it to the DataEstado field.
func (o *Cartao) SetDataEstado(v time.Time) {
	o.DataEstado.Set(&v)
}
// SetDataEstadoNil sets the value for DataEstado to be an explicit nil
func (o *Cartao) SetDataEstadoNil() {
	o.DataEstado.Set(nil)
}

// UnsetDataEstado ensures that no value is present for DataEstado, not even an explicit nil
func (o *Cartao) UnsetDataEstado() {
	o.DataEstado.Unset()
}

// GetPin returns the Pin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetPin() string {
	if o == nil || o.Pin.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pin.Get()
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetPinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Pin.Get(), o.Pin.IsSet()
}

// HasPin returns a boolean if a field has been set.
func (o *Cartao) HasPin() bool {
	if o != nil && o.Pin.IsSet() {
		return true
	}

	return false
}

// SetPin gets a reference to the given NullableString and assigns it to the Pin field.
func (o *Cartao) SetPin(v string) {
	o.Pin.Set(&v)
}
// SetPinNil sets the value for Pin to be an explicit nil
func (o *Cartao) SetPinNil() {
	o.Pin.Set(nil)
}

// UnsetPin ensures that no value is present for Pin, not even an explicit nil
func (o *Cartao) UnsetPin() {
	o.Pin.Unset()
}

// GetCartaoProdutos returns the CartaoProdutos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cartao) GetCartaoProdutos() []CartaoProduto {
	if o == nil  {
		var ret []CartaoProduto
		return ret
	}
	return o.CartaoProdutos
}

// GetCartaoProdutosOk returns a tuple with the CartaoProdutos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cartao) GetCartaoProdutosOk() (*[]CartaoProduto, bool) {
	if o == nil || o.CartaoProdutos == nil {
		return nil, false
	}
	return &o.CartaoProdutos, true
}

// HasCartaoProdutos returns a boolean if a field has been set.
func (o *Cartao) HasCartaoProdutos() bool {
	if o != nil && o.CartaoProdutos != nil {
		return true
	}

	return false
}

// SetCartaoProdutos gets a reference to the given []CartaoProduto and assigns it to the CartaoProdutos field.
func (o *Cartao) SetCartaoProdutos(v []CartaoProduto) {
	o.CartaoProdutos = v
}

func (o Cartao) MarshalJSON() ([]byte, error) {
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
	if o.NumeroSerie.IsSet() {
		toSerialize["numeroSerie"] = o.NumeroSerie.Get()
	}
	if o.DataProducao.IsSet() {
		toSerialize["dataProducao"] = o.DataProducao.Get()
	}
	if o.Estado.IsSet() {
		toSerialize["estado"] = o.Estado.Get()
	}
	if o.DataEstado.IsSet() {
		toSerialize["dataEstado"] = o.DataEstado.Get()
	}
	if o.Pin.IsSet() {
		toSerialize["pin"] = o.Pin.Get()
	}
	if o.CartaoProdutos != nil {
		toSerialize["cartaoProdutos"] = o.CartaoProdutos
	}
	return json.Marshal(toSerialize)
}

type NullableCartao struct {
	value *Cartao
	isSet bool
}

func (v NullableCartao) Get() *Cartao {
	return v.value
}

func (v *NullableCartao) Set(val *Cartao) {
	v.value = val
	v.isSet = true
}

func (v NullableCartao) IsSet() bool {
	return v.isSet
}

func (v *NullableCartao) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartao(val *Cartao) *NullableCartao {
	return &NullableCartao{value: val, isSet: true}
}

func (v NullableCartao) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartao) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


