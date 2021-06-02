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

// Venda struct for Venda
type Venda struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Data *time.Time `json:"data,omitempty"`
	Validade NullableTime `json:"validade,omitempty"`
	DataEstado NullableTime `json:"dataEstado,omitempty"`
	DataLevantamento NullableTime `json:"dataLevantamento,omitempty"`
	Origem *TipoOrigem `json:"origem,omitempty"`
	Estado *TipoVendaEstado `json:"estado,omitempty"`
	RefCartaoProduto NullableInt64 `json:"refCartaoProduto,omitempty"`
	RefGvm NullableInt64 `json:"refGvm,omitempty"`
	RefProduto NullableInt64 `json:"refProduto,omitempty"`
	Vasilhame NullableBool `json:"vasilhame,omitempty"`
	Token NullableString `json:"token,omitempty"`
	Gvm *Gvm `json:"gvm,omitempty"`
	Produto *Produto `json:"produto,omitempty"`
	CartaoProduto *CartaoProduto `json:"cartaoProduto,omitempty"`
}

// NewVenda instantiates a new Venda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVenda() *Venda {
	this := Venda{}
	return &this
}

// NewVendaWithDefaults instantiates a new Venda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendaWithDefaults() *Venda {
	this := Venda{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Venda) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Venda) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Venda) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Venda) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Venda) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Venda) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Venda) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Venda) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Venda) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Venda) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Venda) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Venda) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Venda) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Venda) UnsetUsr() {
	o.Usr.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Venda) GetData() time.Time {
	if o == nil || o.Data == nil {
		var ret time.Time
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetDataOk() (*time.Time, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Venda) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given time.Time and assigns it to the Data field.
func (o *Venda) SetData(v time.Time) {
	o.Data = &v
}

// GetValidade returns the Validade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetValidade() time.Time {
	if o == nil || o.Validade.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Validade.Get()
}

// GetValidadeOk returns a tuple with the Validade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetValidadeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Validade.Get(), o.Validade.IsSet()
}

// HasValidade returns a boolean if a field has been set.
func (o *Venda) HasValidade() bool {
	if o != nil && o.Validade.IsSet() {
		return true
	}

	return false
}

// SetValidade gets a reference to the given NullableTime and assigns it to the Validade field.
func (o *Venda) SetValidade(v time.Time) {
	o.Validade.Set(&v)
}
// SetValidadeNil sets the value for Validade to be an explicit nil
func (o *Venda) SetValidadeNil() {
	o.Validade.Set(nil)
}

// UnsetValidade ensures that no value is present for Validade, not even an explicit nil
func (o *Venda) UnsetValidade() {
	o.Validade.Unset()
}

// GetDataEstado returns the DataEstado field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetDataEstado() time.Time {
	if o == nil || o.DataEstado.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataEstado.Get()
}

// GetDataEstadoOk returns a tuple with the DataEstado field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetDataEstadoOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataEstado.Get(), o.DataEstado.IsSet()
}

// HasDataEstado returns a boolean if a field has been set.
func (o *Venda) HasDataEstado() bool {
	if o != nil && o.DataEstado.IsSet() {
		return true
	}

	return false
}

// SetDataEstado gets a reference to the given NullableTime and assigns it to the DataEstado field.
func (o *Venda) SetDataEstado(v time.Time) {
	o.DataEstado.Set(&v)
}
// SetDataEstadoNil sets the value for DataEstado to be an explicit nil
func (o *Venda) SetDataEstadoNil() {
	o.DataEstado.Set(nil)
}

// UnsetDataEstado ensures that no value is present for DataEstado, not even an explicit nil
func (o *Venda) UnsetDataEstado() {
	o.DataEstado.Unset()
}

// GetDataLevantamento returns the DataLevantamento field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetDataLevantamento() time.Time {
	if o == nil || o.DataLevantamento.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DataLevantamento.Get()
}

// GetDataLevantamentoOk returns a tuple with the DataLevantamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetDataLevantamentoOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DataLevantamento.Get(), o.DataLevantamento.IsSet()
}

// HasDataLevantamento returns a boolean if a field has been set.
func (o *Venda) HasDataLevantamento() bool {
	if o != nil && o.DataLevantamento.IsSet() {
		return true
	}

	return false
}

// SetDataLevantamento gets a reference to the given NullableTime and assigns it to the DataLevantamento field.
func (o *Venda) SetDataLevantamento(v time.Time) {
	o.DataLevantamento.Set(&v)
}
// SetDataLevantamentoNil sets the value for DataLevantamento to be an explicit nil
func (o *Venda) SetDataLevantamentoNil() {
	o.DataLevantamento.Set(nil)
}

// UnsetDataLevantamento ensures that no value is present for DataLevantamento, not even an explicit nil
func (o *Venda) UnsetDataLevantamento() {
	o.DataLevantamento.Unset()
}

// GetOrigem returns the Origem field value if set, zero value otherwise.
func (o *Venda) GetOrigem() TipoOrigem {
	if o == nil || o.Origem == nil {
		var ret TipoOrigem
		return ret
	}
	return *o.Origem
}

// GetOrigemOk returns a tuple with the Origem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetOrigemOk() (*TipoOrigem, bool) {
	if o == nil || o.Origem == nil {
		return nil, false
	}
	return o.Origem, true
}

// HasOrigem returns a boolean if a field has been set.
func (o *Venda) HasOrigem() bool {
	if o != nil && o.Origem != nil {
		return true
	}

	return false
}

// SetOrigem gets a reference to the given TipoOrigem and assigns it to the Origem field.
func (o *Venda) SetOrigem(v TipoOrigem) {
	o.Origem = &v
}

// GetEstado returns the Estado field value if set, zero value otherwise.
func (o *Venda) GetEstado() TipoVendaEstado {
	if o == nil || o.Estado == nil {
		var ret TipoVendaEstado
		return ret
	}
	return *o.Estado
}

// GetEstadoOk returns a tuple with the Estado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetEstadoOk() (*TipoVendaEstado, bool) {
	if o == nil || o.Estado == nil {
		return nil, false
	}
	return o.Estado, true
}

// HasEstado returns a boolean if a field has been set.
func (o *Venda) HasEstado() bool {
	if o != nil && o.Estado != nil {
		return true
	}

	return false
}

// SetEstado gets a reference to the given TipoVendaEstado and assigns it to the Estado field.
func (o *Venda) SetEstado(v TipoVendaEstado) {
	o.Estado = &v
}

// GetRefCartaoProduto returns the RefCartaoProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetRefCartaoProduto() int64 {
	if o == nil || o.RefCartaoProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefCartaoProduto.Get()
}

// GetRefCartaoProdutoOk returns a tuple with the RefCartaoProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetRefCartaoProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefCartaoProduto.Get(), o.RefCartaoProduto.IsSet()
}

// HasRefCartaoProduto returns a boolean if a field has been set.
func (o *Venda) HasRefCartaoProduto() bool {
	if o != nil && o.RefCartaoProduto.IsSet() {
		return true
	}

	return false
}

// SetRefCartaoProduto gets a reference to the given NullableInt64 and assigns it to the RefCartaoProduto field.
func (o *Venda) SetRefCartaoProduto(v int64) {
	o.RefCartaoProduto.Set(&v)
}
// SetRefCartaoProdutoNil sets the value for RefCartaoProduto to be an explicit nil
func (o *Venda) SetRefCartaoProdutoNil() {
	o.RefCartaoProduto.Set(nil)
}

// UnsetRefCartaoProduto ensures that no value is present for RefCartaoProduto, not even an explicit nil
func (o *Venda) UnsetRefCartaoProduto() {
	o.RefCartaoProduto.Unset()
}

// GetRefGvm returns the RefGvm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetRefGvm() int64 {
	if o == nil || o.RefGvm.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefGvm.Get()
}

// GetRefGvmOk returns a tuple with the RefGvm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetRefGvmOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefGvm.Get(), o.RefGvm.IsSet()
}

// HasRefGvm returns a boolean if a field has been set.
func (o *Venda) HasRefGvm() bool {
	if o != nil && o.RefGvm.IsSet() {
		return true
	}

	return false
}

// SetRefGvm gets a reference to the given NullableInt64 and assigns it to the RefGvm field.
func (o *Venda) SetRefGvm(v int64) {
	o.RefGvm.Set(&v)
}
// SetRefGvmNil sets the value for RefGvm to be an explicit nil
func (o *Venda) SetRefGvmNil() {
	o.RefGvm.Set(nil)
}

// UnsetRefGvm ensures that no value is present for RefGvm, not even an explicit nil
func (o *Venda) UnsetRefGvm() {
	o.RefGvm.Unset()
}

// GetRefProduto returns the RefProduto field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetRefProduto() int64 {
	if o == nil || o.RefProduto.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefProduto.Get()
}

// GetRefProdutoOk returns a tuple with the RefProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetRefProdutoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefProduto.Get(), o.RefProduto.IsSet()
}

// HasRefProduto returns a boolean if a field has been set.
func (o *Venda) HasRefProduto() bool {
	if o != nil && o.RefProduto.IsSet() {
		return true
	}

	return false
}

// SetRefProduto gets a reference to the given NullableInt64 and assigns it to the RefProduto field.
func (o *Venda) SetRefProduto(v int64) {
	o.RefProduto.Set(&v)
}
// SetRefProdutoNil sets the value for RefProduto to be an explicit nil
func (o *Venda) SetRefProdutoNil() {
	o.RefProduto.Set(nil)
}

// UnsetRefProduto ensures that no value is present for RefProduto, not even an explicit nil
func (o *Venda) UnsetRefProduto() {
	o.RefProduto.Unset()
}

// GetVasilhame returns the Vasilhame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetVasilhame() bool {
	if o == nil || o.Vasilhame.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Vasilhame.Get()
}

// GetVasilhameOk returns a tuple with the Vasilhame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetVasilhameOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vasilhame.Get(), o.Vasilhame.IsSet()
}

// HasVasilhame returns a boolean if a field has been set.
func (o *Venda) HasVasilhame() bool {
	if o != nil && o.Vasilhame.IsSet() {
		return true
	}

	return false
}

// SetVasilhame gets a reference to the given NullableBool and assigns it to the Vasilhame field.
func (o *Venda) SetVasilhame(v bool) {
	o.Vasilhame.Set(&v)
}
// SetVasilhameNil sets the value for Vasilhame to be an explicit nil
func (o *Venda) SetVasilhameNil() {
	o.Vasilhame.Set(nil)
}

// UnsetVasilhame ensures that no value is present for Vasilhame, not even an explicit nil
func (o *Venda) UnsetVasilhame() {
	o.Vasilhame.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Venda) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Venda) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *Venda) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *Venda) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *Venda) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *Venda) UnsetToken() {
	o.Token.Unset()
}

// GetGvm returns the Gvm field value if set, zero value otherwise.
func (o *Venda) GetGvm() Gvm {
	if o == nil || o.Gvm == nil {
		var ret Gvm
		return ret
	}
	return *o.Gvm
}

// GetGvmOk returns a tuple with the Gvm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetGvmOk() (*Gvm, bool) {
	if o == nil || o.Gvm == nil {
		return nil, false
	}
	return o.Gvm, true
}

// HasGvm returns a boolean if a field has been set.
func (o *Venda) HasGvm() bool {
	if o != nil && o.Gvm != nil {
		return true
	}

	return false
}

// SetGvm gets a reference to the given Gvm and assigns it to the Gvm field.
func (o *Venda) SetGvm(v Gvm) {
	o.Gvm = &v
}

// GetProduto returns the Produto field value if set, zero value otherwise.
func (o *Venda) GetProduto() Produto {
	if o == nil || o.Produto == nil {
		var ret Produto
		return ret
	}
	return *o.Produto
}

// GetProdutoOk returns a tuple with the Produto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetProdutoOk() (*Produto, bool) {
	if o == nil || o.Produto == nil {
		return nil, false
	}
	return o.Produto, true
}

// HasProduto returns a boolean if a field has been set.
func (o *Venda) HasProduto() bool {
	if o != nil && o.Produto != nil {
		return true
	}

	return false
}

// SetProduto gets a reference to the given Produto and assigns it to the Produto field.
func (o *Venda) SetProduto(v Produto) {
	o.Produto = &v
}

// GetCartaoProduto returns the CartaoProduto field value if set, zero value otherwise.
func (o *Venda) GetCartaoProduto() CartaoProduto {
	if o == nil || o.CartaoProduto == nil {
		var ret CartaoProduto
		return ret
	}
	return *o.CartaoProduto
}

// GetCartaoProdutoOk returns a tuple with the CartaoProduto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Venda) GetCartaoProdutoOk() (*CartaoProduto, bool) {
	if o == nil || o.CartaoProduto == nil {
		return nil, false
	}
	return o.CartaoProduto, true
}

// HasCartaoProduto returns a boolean if a field has been set.
func (o *Venda) HasCartaoProduto() bool {
	if o != nil && o.CartaoProduto != nil {
		return true
	}

	return false
}

// SetCartaoProduto gets a reference to the given CartaoProduto and assigns it to the CartaoProduto field.
func (o *Venda) SetCartaoProduto(v CartaoProduto) {
	o.CartaoProduto = &v
}

func (o Venda) MarshalJSON() ([]byte, error) {
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
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Validade.IsSet() {
		toSerialize["validade"] = o.Validade.Get()
	}
	if o.DataEstado.IsSet() {
		toSerialize["dataEstado"] = o.DataEstado.Get()
	}
	if o.DataLevantamento.IsSet() {
		toSerialize["dataLevantamento"] = o.DataLevantamento.Get()
	}
	if o.Origem != nil {
		toSerialize["origem"] = o.Origem
	}
	if o.Estado != nil {
		toSerialize["estado"] = o.Estado
	}
	if o.RefCartaoProduto.IsSet() {
		toSerialize["refCartaoProduto"] = o.RefCartaoProduto.Get()
	}
	if o.RefGvm.IsSet() {
		toSerialize["refGvm"] = o.RefGvm.Get()
	}
	if o.RefProduto.IsSet() {
		toSerialize["refProduto"] = o.RefProduto.Get()
	}
	if o.Vasilhame.IsSet() {
		toSerialize["vasilhame"] = o.Vasilhame.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.Gvm != nil {
		toSerialize["gvm"] = o.Gvm
	}
	if o.Produto != nil {
		toSerialize["produto"] = o.Produto
	}
	if o.CartaoProduto != nil {
		toSerialize["cartaoProduto"] = o.CartaoProduto
	}
	return json.Marshal(toSerialize)
}

type NullableVenda struct {
	value *Venda
	isSet bool
}

func (v NullableVenda) Get() *Venda {
	return v.value
}

func (v *NullableVenda) Set(val *Venda) {
	v.value = val
	v.isSet = true
}

func (v NullableVenda) IsSet() bool {
	return v.isSet
}

func (v *NullableVenda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVenda(val *Venda) *NullableVenda {
	return &NullableVenda{value: val, isSet: true}
}

func (v NullableVenda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVenda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

