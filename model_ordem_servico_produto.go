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

// OrdemServicoProduto struct for OrdemServicoProduto
type OrdemServicoProduto struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	RefOrdemServico NullableInt64 `json:"refOrdemServico,omitempty"`
	RefProduto NullableInt64 `json:"refProduto,omitempty"`
	Quant NullableInt32 `json:"quant,omitempty"`
	Vasilhame *bool `json:"vasilhame,omitempty"`
	OrdemServico *OrdemServico `json:"ordemServico,omitempty"`
	Produto *Produto `json:"produto,omitempty"`
}

// NewOrdemServicoProduto instantiates a new OrdemServicoProduto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdemServicoProduto() *OrdemServicoProduto {
	this := OrdemServicoProduto{}
	return &this
}

// NewOrdemServicoProdutoWithDefaults instantiates a new OrdemServicoProduto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdemServicoProdutoWithDefaults() *OrdemServicoProduto {
	this := OrdemServicoProduto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrdemServicoProduto) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdemServicoProduto) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *OrdemServicoProduto) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *OrdemServicoProduto) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdemServicoProduto) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *OrdemServicoProduto) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrdemServicoProduto) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdemServicoProduto) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *OrdemServicoProduto) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *OrdemServicoProduto) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *OrdemServicoProduto) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrdemServicoProduto) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdemServicoProduto) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *OrdemServicoProduto) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *OrdemServicoProduto) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *OrdemServicoProduto) UnsetUsr() {
	o.Usr.Unset()
}

// GetRefOrdemServico returns the RefOrdemServico field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrdemServicoProduto) GetRefOrdemServico() int64 {
	if o == nil || o.RefOrdemServico.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefOrdemServico.Get()
}

// GetRefOrdemServicoOk returns a tuple with the RefOrdemServico field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdemServicoProduto) GetRefOrdemServicoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefOrdemServico.Get(), o.RefOrdemServico.IsSet()
}

// HasRefOrdemServico returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasRefOrdemServico() bool {
	if o != nil && o.RefOrdemServico.IsSet() {
		return true
	}

	return false
}

// SetRefOrdemServico gets a reference to the given NullableInt64 and assigns it to the RefOrdemServico field.
func (o *OrdemServicoProduto) SetRefOrdemServico(v int64) {
	o.RefOrdemServico.Set(&v)
}
// SetRefOrdemServicoNil sets the value for RefOrdemServico to be an explicit nil
func (o *OrdemServicoProduto) SetRefOrdemServicoNil() {
	o.RefOrdemServico.Set(nil)
}

// UnsetRefOrdemServico ensures that no value is present for RefOrdemServico, not even an explicit nil
func (o *OrdemServicoProduto) UnsetRefOrdemServico() {
	o.RefOrdemServico.Unset()
}

// GetRefProduto returns the RefProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrdemServicoProduto) GetRefProduto() int64 {
	if o == nil || o.RefProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefProduto.Get()
}

// GetRefProdutoOk returns a tuple with the RefProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdemServicoProduto) GetRefProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefProduto.Get(), o.RefProduto.IsSet()
}

// HasRefProduto returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasRefProduto() bool {
	if o != nil && o.RefProduto.IsSet() {
		return true
	}

	return false
}

// SetRefProduto gets a reference to the given NullableInt64 and assigns it to the RefProduto field.
func (o *OrdemServicoProduto) SetRefProduto(v int64) {
	o.RefProduto.Set(&v)
}
// SetRefProdutoNil sets the value for RefProduto to be an explicit nil
func (o *OrdemServicoProduto) SetRefProdutoNil() {
	o.RefProduto.Set(nil)
}

// UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
func (o *OrdemServicoProduto) UnsetRefProduto() {
	o.RefProduto.Unset()
}

// GetQuant returns the Quant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrdemServicoProduto) GetQuant() int32 {
	if o == nil || o.Quant.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Quant.Get()
}

// GetQuantOk returns a tuple with the Quant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdemServicoProduto) GetQuantOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quant.Get(), o.Quant.IsSet()
}

// HasQuant returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasQuant() bool {
	if o != nil && o.Quant.IsSet() {
		return true
	}

	return false
}

// SetQuant gets a reference to the given NullableInt32 and assigns it to the Quant field.
func (o *OrdemServicoProduto) SetQuant(v int32) {
	o.Quant.Set(&v)
}
// SetQuantNil sets the value for Quant to be an explicit nil
func (o *OrdemServicoProduto) SetQuantNil() {
	o.Quant.Set(nil)
}

// UnsetQuant ensures that no value is present for Quant, not even an explicit nil
func (o *OrdemServicoProduto) UnsetQuant() {
	o.Quant.Unset()
}

// GetVasilhame returns the Vasilhame field value if set, zero value otherwise.
func (o *OrdemServicoProduto) GetVasilhame() bool {
	if o == nil || o.Vasilhame == nil {
		var ret bool
		return ret
	}
	return *o.Vasilhame
}

// GetVasilhameOk returns a tuple with the Vasilhame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdemServicoProduto) GetVasilhameOk() (*bool, bool) {
	if o == nil || o.Vasilhame == nil {
		return nil, false
	}
	return o.Vasilhame, true
}

// HasVasilhame returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasVasilhame() bool {
	if o != nil && o.Vasilhame != nil {
		return true
	}

	return false
}

// SetVasilhame gets a reference to the given bool and assigns it to the Vasilhame field.
func (o *OrdemServicoProduto) SetVasilhame(v bool) {
	o.Vasilhame = &v
}

// GetOrdemServico returns the OrdemServico field value if set, zero value otherwise.
func (o *OrdemServicoProduto) GetOrdemServico() OrdemServico {
	if o == nil || o.OrdemServico == nil {
		var ret OrdemServico
		return ret
	}
	return *o.OrdemServico
}

// GetOrdemServicoOk returns a tuple with the OrdemServico field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdemServicoProduto) GetOrdemServicoOk() (*OrdemServico, bool) {
	if o == nil || o.OrdemServico == nil {
		return nil, false
	}
	return o.OrdemServico, true
}

// HasOrdemServico returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasOrdemServico() bool {
	if o != nil && o.OrdemServico != nil {
		return true
	}

	return false
}

// SetOrdemServico gets a reference to the given OrdemServico and assigns it to the OrdemServico field.
func (o *OrdemServicoProduto) SetOrdemServico(v OrdemServico) {
	o.OrdemServico = &v
}

// GetProduto returns the Produto field value if set, zero value otherwise.
func (o *OrdemServicoProduto) GetProduto() Produto {
	if o == nil || o.Produto == nil {
		var ret Produto
		return ret
	}
	return *o.Produto
}

// GetProdutoOk returns a tuple with the Produto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdemServicoProduto) GetProdutoOk() (*Produto, bool) {
	if o == nil || o.Produto == nil {
		return nil, false
	}
	return o.Produto, true
}

// HasProduto returns a boolean if a field has been set.
func (o *OrdemServicoProduto) HasProduto() bool {
	if o != nil && o.Produto != nil {
		return true
	}

	return false
}

// SetProduto gets a reference to the given Produto and assigns it to the Produto field.
func (o *OrdemServicoProduto) SetProduto(v Produto) {
	o.Produto = &v
}

func (o OrdemServicoProduto) MarshalJSON() ([]byte, error) {
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
	if o.RefOrdemServico.IsSet() {
		toSerialize["refOrdemServico"] = o.RefOrdemServico.Get()
	}
	if o.RefProduto.IsSet() {
		toSerialize["refProduto"] = o.RefProduto.Get()
	}
	if o.Quant.IsSet() {
		toSerialize["quant"] = o.Quant.Get()
	}
	if o.Vasilhame != nil {
		toSerialize["vasilhame"] = o.Vasilhame
	}
	if o.OrdemServico != nil {
		toSerialize["ordemServico"] = o.OrdemServico
	}
	if o.Produto != nil {
		toSerialize["produto"] = o.Produto
	}
	return json.Marshal(toSerialize)
}

type NullableOrdemServicoProduto struct {
	value *OrdemServicoProduto
	isSet bool
}

func (v NullableOrdemServicoProduto) Get() *OrdemServicoProduto {
	return v.value
}

func (v *NullableOrdemServicoProduto) Set(val *OrdemServicoProduto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdemServicoProduto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdemServicoProduto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdemServicoProduto(val *OrdemServicoProduto) *NullableOrdemServicoProduto {
	return &NullableOrdemServicoProduto{value: val, isSet: true}
}

func (v NullableOrdemServicoProduto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdemServicoProduto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

