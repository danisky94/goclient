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

// CartaoProduto struct for CartaoProduto
type CartaoProduto struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefCartao NullableInt64 `json:"refCartao,omitempty"`
	RefProduto NullableInt64 `json:"refProduto,omitempty"`
	Data NullableTime `json:"data,omitempty"`
	Quantidade NullableInt32 `json:"quantidade,omitempty"`
	Vasilhame NullableBool `json:"vasilhame,omitempty"`
	RefCartaoNavigation *Cartao `json:"refCartaoNavigation,omitempty"`
	Venda []Venda `json:"venda,omitempty"`
}

// NewCartaoProduto instantiates a new CartaoProduto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartaoProduto() *CartaoProduto {
	this := CartaoProduto{}
	return &this
}

// NewCartaoProdutoWithDefaults instantiates a new CartaoProduto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartaoProdutoWithDefaults() *CartaoProduto {
	this := CartaoProduto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CartaoProduto) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartaoProduto) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CartaoProduto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CartaoProduto) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CartaoProduto) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartaoProduto) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CartaoProduto) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *CartaoProduto) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *CartaoProduto) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *CartaoProduto) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *CartaoProduto) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *CartaoProduto) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *CartaoProduto) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *CartaoProduto) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *CartaoProduto) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *CartaoProduto) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefCartao returns the RefCartao field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetRefCartao() int64 {
	if o == nil || o.RefCartao.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefCartao.Get()
}

// GetRefCartaoOk returns a tuple with the RefCartao field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetRefCartaoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefCartao.Get(), o.RefCartao.IsSet()
}

// HasRefCartao returns a boolean if a field has been set.
func (o *CartaoProduto) HasRefCartao() bool {
	if o != nil && o.RefCartao.IsSet() {
		return true
	}

	return false
}

// SetRefCartao gets a reference to the given NullableInt64 and assigns it to the RefCartao field.
func (o *CartaoProduto) SetRefCartao(v int64) {
	o.RefCartao.Set(&v)
}
// SetRefCartaoNil sets the value for RefCartao to be an explicit nil
func (o *CartaoProduto) SetRefCartaoNil() {
	o.RefCartao.Set(nil)
}

// UnsetRefCartao ensures that no value is present for RefCartao, not even an explicit nil
func (o *CartaoProduto) UnsetRefCartao() {
	o.RefCartao.Unset()
}

// GetRefProduto returns the RefProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetRefProduto() int64 {
	if o == nil || o.RefProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefProduto.Get()
}

// GetRefProdutoOk returns a tuple with the RefProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetRefProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefProduto.Get(), o.RefProduto.IsSet()
}

// HasRefProduto returns a boolean if a field has been set.
func (o *CartaoProduto) HasRefProduto() bool {
	if o != nil && o.RefProduto.IsSet() {
		return true
	}

	return false
}

// SetRefProduto gets a reference to the given NullableInt64 and assigns it to the RefProduto field.
func (o *CartaoProduto) SetRefProduto(v int64) {
	o.RefProduto.Set(&v)
}
// SetRefProdutoNil sets the value for RefProduto to be an explicit nil
func (o *CartaoProduto) SetRefProdutoNil() {
	o.RefProduto.Set(nil)
}

// UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
func (o *CartaoProduto) UnsetRefProduto() {
	o.RefProduto.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetData() time.Time {
	if o == nil || o.Data.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetDataOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *CartaoProduto) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTime and assigns it to the Data field.
func (o *CartaoProduto) SetData(v time.Time) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *CartaoProduto) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *CartaoProduto) UnsetData() {
	o.Data.Unset()
}

// GetQuantidade returns the Quantidade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetQuantidade() int32 {
	if o == nil || o.Quantidade.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Quantidade.Get()
}

// GetQuantidadeOk returns a tuple with the Quantidade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetQuantidadeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quantidade.Get(), o.Quantidade.IsSet()
}

// HasQuantidade returns a boolean if a field has been set.
func (o *CartaoProduto) HasQuantidade() bool {
	if o != nil && o.Quantidade.IsSet() {
		return true
	}

	return false
}

// SetQuantidade gets a reference to the given NullableInt32 and assigns it to the Quantidade field.
func (o *CartaoProduto) SetQuantidade(v int32) {
	o.Quantidade.Set(&v)
}
// SetQuantidadeNil sets the value for Quantidade to be an explicit nil
func (o *CartaoProduto) SetQuantidadeNil() {
	o.Quantidade.Set(nil)
}

// UnsetQuantidade ensures that no value is present for Quantidade, not even an explicit nil
func (o *CartaoProduto) UnsetQuantidade() {
	o.Quantidade.Unset()
}

// GetVasilhame returns the Vasilhame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetVasilhame() bool {
	if o == nil || o.Vasilhame.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Vasilhame.Get()
}

// GetVasilhameOk returns a tuple with the Vasilhame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetVasilhameOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vasilhame.Get(), o.Vasilhame.IsSet()
}

// HasVasilhame returns a boolean if a field has been set.
func (o *CartaoProduto) HasVasilhame() bool {
	if o != nil && o.Vasilhame.IsSet() {
		return true
	}

	return false
}

// SetVasilhame gets a reference to the given NullableBool and assigns it to the Vasilhame field.
func (o *CartaoProduto) SetVasilhame(v bool) {
	o.Vasilhame.Set(&v)
}
// SetVasilhameNil sets the value for Vasilhame to be an explicit nil
func (o *CartaoProduto) SetVasilhameNil() {
	o.Vasilhame.Set(nil)
}

// UnsetVasilhame ensures that no value is present for Vasilhame, not even an explicit nil
func (o *CartaoProduto) UnsetVasilhame() {
	o.Vasilhame.Unset()
}

// GetRefCartaoNavigation returns the RefCartaoNavigation field value if set, zero value otherwise.
func (o *CartaoProduto) GetRefCartaoNavigation() Cartao {
	if o == nil || o.RefCartaoNavigation == nil {
		var ret Cartao
		return ret
	}
	return *o.RefCartaoNavigation
}

// GetRefCartaoNavigationOk returns a tuple with the RefCartaoNavigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartaoProduto) GetRefCartaoNavigationOk() (*Cartao, bool) {
	if o == nil || o.RefCartaoNavigation == nil {
		return nil, false
	}
	return o.RefCartaoNavigation, true
}

// HasRefCartaoNavigation returns a boolean if a field has been set.
func (o *CartaoProduto) HasRefCartaoNavigation() bool {
	if o != nil && o.RefCartaoNavigation != nil {
		return true
	}

	return false
}

// SetRefCartaoNavigation gets a reference to the given Cartao and assigns it to the RefCartaoNavigation field.
func (o *CartaoProduto) SetRefCartaoNavigation(v Cartao) {
	o.RefCartaoNavigation = &v
}

// GetVenda returns the Venda field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartaoProduto) GetVenda() []Venda {
	if o == nil  {
		var ret []Venda
		return ret
	}
	return o.Venda
}

// GetVendaOk returns a tuple with the Venda field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartaoProduto) GetVendaOk() (*[]Venda, bool) {
	if o == nil || o.Venda == nil {
		return nil, false
	}
	return &o.Venda, true
}

// HasVenda returns a boolean if a field has been set.
func (o *CartaoProduto) HasVenda() bool {
	if o != nil && o.Venda != nil {
		return true
	}

	return false
}

// SetVenda gets a reference to the given []Venda and assigns it to the Venda field.
func (o *CartaoProduto) SetVenda(v []Venda) {
	o.Venda = v
}

func (o CartaoProduto) MarshalJSON() ([]byte, error) {
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
	if o.RefCartao.IsSet() {
		toSerialize["refCartao"] = o.RefCartao.Get()
	}
	if o.RefProduto.IsSet() {
		toSerialize["refProduto"] = o.RefProduto.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Quantidade.IsSet() {
		toSerialize["quantidade"] = o.Quantidade.Get()
	}
	if o.Vasilhame.IsSet() {
		toSerialize["vasilhame"] = o.Vasilhame.Get()
	}
	if o.RefCartaoNavigation != nil {
		toSerialize["refCartaoNavigation"] = o.RefCartaoNavigation
	}
	if o.Venda != nil {
		toSerialize["venda"] = o.Venda
	}
	return json.Marshal(toSerialize)
}

type NullableCartaoProduto struct {
	value *CartaoProduto
	isSet bool
}

func (v NullableCartaoProduto) Get() *CartaoProduto {
	return v.value
}

func (v *NullableCartaoProduto) Set(val *CartaoProduto) {
	v.value = val
	v.isSet = true
}

func (v NullableCartaoProduto) IsSet() bool {
	return v.isSet
}

func (v *NullableCartaoProduto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartaoProduto(val *CartaoProduto) *NullableCartaoProduto {
	return &NullableCartaoProduto{value: val, isSet: true}
}

func (v NullableCartaoProduto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartaoProduto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


