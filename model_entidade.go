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

// Entidade struct for Entidade
type Entidade struct {
	Id *int64 `json:"id,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Upd NullableTime `json:"upd,omitempty"`
	Usr NullableInt64 `json:"usr,omitempty"`
	Nome NullableString `json:"nome,omitempty"`
	Tipo *TipoEntidade `json:"tipo,omitempty"`
	Morada NullableString `json:"morada,omitempty"`
	Telefone NullableInt32 `json:"telefone,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Nif NullableInt32 `json:"nif,omitempty"`
	RefImagem NullableInt64 `json:"refImagem,omitempty"`
	CorPrincipal NullableString `json:"corPrincipal,omitempty"`
	CorSecundaria NullableString `json:"corSecundaria,omitempty"`
	RefPlano NullableInt64 `json:"refPlano,omitempty"`
	RefEntidade NullableInt64 `json:"refEntidade,omitempty"`
	Imagem *Imagem `json:"imagem,omitempty"`
	RefPlanoNavigation *Plano `json:"refPlanoNavigation,omitempty"`
	Utilizadores []Utilizador `json:"utilizadores,omitempty"`
}

// NewEntidade instantiates a new Entidade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntidade() *Entidade {
	this := Entidade{}
	return &this
}

// NewEntidadeWithDefaults instantiates a new Entidade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntidadeWithDefaults() *Entidade {
	this := Entidade{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Entidade) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entidade) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Entidade) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Entidade) SetId(v int64) {
	o.Id = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Entidade) GetDeleted() int32 {
	if o == nil || o.Deleted == nil {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entidade) GetDeletedOk() (*int32, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Entidade) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *Entidade) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetUpd returns the Upd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetUpd() time.Time {
	if o == nil || o.Upd.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Upd.Get()
}

// GetUpdOk returns a tuple with the Upd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetUpdOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Upd.Get(), o.Upd.IsSet()
}

// HasUpd returns a boolean if a field has been set.
func (o *Entidade) HasUpd() bool {
	if o != nil && o.Upd.IsSet() {
		return true
	}

	return false
}

// SetUpd gets a reference to the given NullableTime and assigns it to the Upd field.
func (o *Entidade) SetUpd(v time.Time) {
	o.Upd.Set(&v)
}
// SetUpdNil sets the value for Upd to be an explicit nil
func (o *Entidade) SetUpdNil() {
	o.Upd.Set(nil)
}

// UnsetUpd ensures that no value is present for Upd, not even an explicit nil
func (o *Entidade) UnsetUpd() {
	o.Upd.Unset()
}

// GetUsr returns the Usr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetUsr() int64 {
	if o == nil || o.Usr.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Usr.Get()
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetUsrOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Usr.Get(), o.Usr.IsSet()
}

// HasUsr returns a boolean if a field has been set.
func (o *Entidade) HasUsr() bool {
	if o != nil && o.Usr.IsSet() {
		return true
	}

	return false
}

// SetUsr gets a reference to the given NullableInt64 and assigns it to the Usr field.
func (o *Entidade) SetUsr(v int64) {
	o.Usr.Set(&v)
}
// SetUsrNil sets the value for Usr to be an explicit nil
func (o *Entidade) SetUsrNil() {
	o.Usr.Set(nil)
}

// UnsetUsr ensures that no value is present for Usr, not even an explicit nil
func (o *Entidade) UnsetUsr() {
	o.Usr.Unset()
}

// GetNome returns the Nome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetNome() string {
	if o == nil || o.Nome.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nome.Get()
}

// GetNomeOk returns a tuple with the Nome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetNomeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nome.Get(), o.Nome.IsSet()
}

// HasNome returns a boolean if a field has been set.
func (o *Entidade) HasNome() bool {
	if o != nil && o.Nome.IsSet() {
		return true
	}

	return false
}

// SetNome gets a reference to the given NullableString and assigns it to the Nome field.
func (o *Entidade) SetNome(v string) {
	o.Nome.Set(&v)
}
// SetNomeNil sets the value for Nome to be an explicit nil
func (o *Entidade) SetNomeNil() {
	o.Nome.Set(nil)
}

// UnsetNome ensures that no value is present for Nome, not even an explicit nil
func (o *Entidade) UnsetNome() {
	o.Nome.Unset()
}

// GetTipo returns the Tipo field value if set, zero value otherwise.
func (o *Entidade) GetTipo() TipoEntidade {
	if o == nil || o.Tipo == nil {
		var ret TipoEntidade
		return ret
	}
	return *o.Tipo
}

// GetTipoOk returns a tuple with the Tipo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entidade) GetTipoOk() (*TipoEntidade, bool) {
	if o == nil || o.Tipo == nil {
		return nil, false
	}
	return o.Tipo, true
}

// HasTipo returns a boolean if a field has been set.
func (o *Entidade) HasTipo() bool {
	if o != nil && o.Tipo != nil {
		return true
	}

	return false
}

// SetTipo gets a reference to the given TipoEntidade and assigns it to the Tipo field.
func (o *Entidade) SetTipo(v TipoEntidade) {
	o.Tipo = &v
}

// GetMorada returns the Morada field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetMorada() string {
	if o == nil || o.Morada.Get() == nil {
		var ret string
		return ret
	}
	return *o.Morada.Get()
}

// GetMoradaOk returns a tuple with the Morada field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetMoradaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Morada.Get(), o.Morada.IsSet()
}

// HasMorada returns a boolean if a field has been set.
func (o *Entidade) HasMorada() bool {
	if o != nil && o.Morada.IsSet() {
		return true
	}

	return false
}

// SetMorada gets a reference to the given NullableString and assigns it to the Morada field.
func (o *Entidade) SetMorada(v string) {
	o.Morada.Set(&v)
}
// SetMoradaNil sets the value for Morada to be an explicit nil
func (o *Entidade) SetMoradaNil() {
	o.Morada.Set(nil)
}

// UnsetMorada ensures that no value is present for Morada, not even an explicit nil
func (o *Entidade) UnsetMorada() {
	o.Morada.Unset()
}

// GetTelefone returns the Telefone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetTelefone() int32 {
	if o == nil || o.Telefone.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Telefone.Get()
}

// GetTelefoneOk returns a tuple with the Telefone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetTelefoneOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Telefone.Get(), o.Telefone.IsSet()
}

// HasTelefone returns a boolean if a field has been set.
func (o *Entidade) HasTelefone() bool {
	if o != nil && o.Telefone.IsSet() {
		return true
	}

	return false
}

// SetTelefone gets a reference to the given NullableInt32 and assigns it to the Telefone field.
func (o *Entidade) SetTelefone(v int32) {
	o.Telefone.Set(&v)
}
// SetTelefoneNil sets the value for Telefone to be an explicit nil
func (o *Entidade) SetTelefoneNil() {
	o.Telefone.Set(nil)
}

// UnsetTelefone ensures that no value is present for Telefone, not even an explicit nil
func (o *Entidade) UnsetTelefone() {
	o.Telefone.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Entidade) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Entidade) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Entidade) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Entidade) UnsetEmail() {
	o.Email.Unset()
}

// GetNif returns the Nif field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetNif() int32 {
	if o == nil || o.Nif.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Nif.Get()
}

// GetNifOk returns a tuple with the Nif field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetNifOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nif.Get(), o.Nif.IsSet()
}

// HasNif returns a boolean if a field has been set.
func (o *Entidade) HasNif() bool {
	if o != nil && o.Nif.IsSet() {
		return true
	}

	return false
}

// SetNif gets a reference to the given NullableInt32 and assigns it to the Nif field.
func (o *Entidade) SetNif(v int32) {
	o.Nif.Set(&v)
}
// SetNifNil sets the value for Nif to be an explicit nil
func (o *Entidade) SetNifNil() {
	o.Nif.Set(nil)
}

// UnsetNif ensures that no value is present for Nif, not even an explicit nil
func (o *Entidade) UnsetNif() {
	o.Nif.Unset()
}

// GetRefImagem returns the RefImagem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetRefImagem() int64 {
	if o == nil || o.RefImagem.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefImagem.Get()
}

// GetRefImagemOk returns a tuple with the RefImagem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetRefImagemOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefImagem.Get(), o.RefImagem.IsSet()
}

// HasRefImagem returns a boolean if a field has been set.
func (o *Entidade) HasRefImagem() bool {
	if o != nil && o.RefImagem.IsSet() {
		return true
	}

	return false
}

// SetRefImagem gets a reference to the given NullableInt64 and assigns it to the RefImagem field.
func (o *Entidade) SetRefImagem(v int64) {
	o.RefImagem.Set(&v)
}
// SetRefImagemNil sets the value for RefImagem to be an explicit nil
func (o *Entidade) SetRefImagemNil() {
	o.RefImagem.Set(nil)
}

// UnsetRefImagem ensures that no value is present for RefImagem, not even an explicit nil
func (o *Entidade) UnsetRefImagem() {
	o.RefImagem.Unset()
}

// GetCorPrincipal returns the CorPrincipal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetCorPrincipal() string {
	if o == nil || o.CorPrincipal.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorPrincipal.Get()
}

// GetCorPrincipalOk returns a tuple with the CorPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetCorPrincipalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CorPrincipal.Get(), o.CorPrincipal.IsSet()
}

// HasCorPrincipal returns a boolean if a field has been set.
func (o *Entidade) HasCorPrincipal() bool {
	if o != nil && o.CorPrincipal.IsSet() {
		return true
	}

	return false
}

// SetCorPrincipal gets a reference to the given NullableString and assigns it to the CorPrincipal field.
func (o *Entidade) SetCorPrincipal(v string) {
	o.CorPrincipal.Set(&v)
}
// SetCorPrincipalNil sets the value for CorPrincipal to be an explicit nil
func (o *Entidade) SetCorPrincipalNil() {
	o.CorPrincipal.Set(nil)
}

// UnsetCorPrincipal ensures that no value is present for CorPrincipal, not even an explicit nil
func (o *Entidade) UnsetCorPrincipal() {
	o.CorPrincipal.Unset()
}

// GetCorSecundaria returns the CorSecundaria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetCorSecundaria() string {
	if o == nil || o.CorSecundaria.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorSecundaria.Get()
}

// GetCorSecundariaOk returns a tuple with the CorSecundaria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetCorSecundariaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CorSecundaria.Get(), o.CorSecundaria.IsSet()
}

// HasCorSecundaria returns a boolean if a field has been set.
func (o *Entidade) HasCorSecundaria() bool {
	if o != nil && o.CorSecundaria.IsSet() {
		return true
	}

	return false
}

// SetCorSecundaria gets a reference to the given NullableString and assigns it to the CorSecundaria field.
func (o *Entidade) SetCorSecundaria(v string) {
	o.CorSecundaria.Set(&v)
}
// SetCorSecundariaNil sets the value for CorSecundaria to be an explicit nil
func (o *Entidade) SetCorSecundariaNil() {
	o.CorSecundaria.Set(nil)
}

// UnsetCorSecundaria ensures that no value is present for CorSecundaria, not even an explicit nil
func (o *Entidade) UnsetCorSecundaria() {
	o.CorSecundaria.Unset()
}

// GetRefPlano returns the RefPlano field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetRefPlano() int64 {
	if o == nil || o.RefPlano.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefPlano.Get()
}

// GetRefPlanoOk returns a tuple with the RefPlano field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetRefPlanoOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefPlano.Get(), o.RefPlano.IsSet()
}

// HasRefPlano returns a boolean if a field has been set.
func (o *Entidade) HasRefPlano() bool {
	if o != nil && o.RefPlano.IsSet() {
		return true
	}

	return false
}

// SetRefPlano gets a reference to the given NullableInt64 and assigns it to the RefPlano field.
func (o *Entidade) SetRefPlano(v int64) {
	o.RefPlano.Set(&v)
}
// SetRefPlanoNil sets the value for RefPlano to be an explicit nil
func (o *Entidade) SetRefPlanoNil() {
	o.RefPlano.Set(nil)
}

// UnsetRefPlano ensures that no value is present for RefPlano, not even an explicit nil
func (o *Entidade) UnsetRefPlano() {
	o.RefPlano.Unset()
}

// GetRefEntidade returns the RefEntidade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetRefEntidade() int64 {
	if o == nil || o.RefEntidade.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RefEntidade.Get()
}

// GetRefEntidadeOk returns a tuple with the RefEntidade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetRefEntidadeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefEntidade.Get(), o.RefEntidade.IsSet()
}

// HasRefEntidade returns a boolean if a field has been set.
func (o *Entidade) HasRefEntidade() bool {
	if o != nil && o.RefEntidade.IsSet() {
		return true
	}

	return false
}

// SetRefEntidade gets a reference to the given NullableInt64 and assigns it to the RefEntidade field.
func (o *Entidade) SetRefEntidade(v int64) {
	o.RefEntidade.Set(&v)
}
// SetRefEntidadeNil sets the value for RefEntidade to be an explicit nil
func (o *Entidade) SetRefEntidadeNil() {
	o.RefEntidade.Set(nil)
}

// UnsetRefEntidade ensures that no value is present for RefEntidade, not even an explicit nil
func (o *Entidade) UnsetRefEntidade() {
	o.RefEntidade.Unset()
}

// GetImagem returns the Imagem field value if set, zero value otherwise.
func (o *Entidade) GetImagem() Imagem {
	if o == nil || o.Imagem == nil {
		var ret Imagem
		return ret
	}
	return *o.Imagem
}

// GetImagemOk returns a tuple with the Imagem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entidade) GetImagemOk() (*Imagem, bool) {
	if o == nil || o.Imagem == nil {
		return nil, false
	}
	return o.Imagem, true
}

// HasImagem returns a boolean if a field has been set.
func (o *Entidade) HasImagem() bool {
	if o != nil && o.Imagem != nil {
		return true
	}

	return false
}

// SetImagem gets a reference to the given Imagem and assigns it to the Imagem field.
func (o *Entidade) SetImagem(v Imagem) {
	o.Imagem = &v
}

// GetRefPlanoNavigation returns the RefPlanoNavigation field value if set, zero value otherwise.
func (o *Entidade) GetRefPlanoNavigation() Plano {
	if o == nil || o.RefPlanoNavigation == nil {
		var ret Plano
		return ret
	}
	return *o.RefPlanoNavigation
}

// GetRefPlanoNavigationOk returns a tuple with the RefPlanoNavigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entidade) GetRefPlanoNavigationOk() (*Plano, bool) {
	if o == nil || o.RefPlanoNavigation == nil {
		return nil, false
	}
	return o.RefPlanoNavigation, true
}

// HasRefPlanoNavigation returns a boolean if a field has been set.
func (o *Entidade) HasRefPlanoNavigation() bool {
	if o != nil && o.RefPlanoNavigation != nil {
		return true
	}

	return false
}

// SetRefPlanoNavigation gets a reference to the given Plano and assigns it to the RefPlanoNavigation field.
func (o *Entidade) SetRefPlanoNavigation(v Plano) {
	o.RefPlanoNavigation = &v
}

// GetUtilizadores returns the Utilizadores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entidade) GetUtilizadores() []Utilizador {
	if o == nil  {
		var ret []Utilizador
		return ret
	}
	return o.Utilizadores
}

// GetUtilizadoresOk returns a tuple with the Utilizadores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entidade) GetUtilizadoresOk() (*[]Utilizador, bool) {
	if o == nil || o.Utilizadores == nil {
		return nil, false
	}
	return &o.Utilizadores, true
}

// HasUtilizadores returns a boolean if a field has been set.
func (o *Entidade) HasUtilizadores() bool {
	if o != nil && o.Utilizadores != nil {
		return true
	}

	return false
}

// SetUtilizadores gets a reference to the given []Utilizador and assigns it to the Utilizadores field.
func (o *Entidade) SetUtilizadores(v []Utilizador) {
	o.Utilizadores = v
}

func (o Entidade) MarshalJSON() ([]byte, error) {
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
	if o.Morada.IsSet() {
		toSerialize["morada"] = o.Morada.Get()
	}
	if o.Telefone.IsSet() {
		toSerialize["telefone"] = o.Telefone.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Nif.IsSet() {
		toSerialize["nif"] = o.Nif.Get()
	}
	if o.RefImagem.IsSet() {
		toSerialize["refImagem"] = o.RefImagem.Get()
	}
	if o.CorPrincipal.IsSet() {
		toSerialize["corPrincipal"] = o.CorPrincipal.Get()
	}
	if o.CorSecundaria.IsSet() {
		toSerialize["corSecundaria"] = o.CorSecundaria.Get()
	}
	if o.RefPlano.IsSet() {
		toSerialize["refPlano"] = o.RefPlano.Get()
	}
	if o.RefEntidade.IsSet() {
		toSerialize["refEntidade"] = o.RefEntidade.Get()
	}
	if o.Imagem != nil {
		toSerialize["imagem"] = o.Imagem
	}
	if o.RefPlanoNavigation != nil {
		toSerialize["refPlanoNavigation"] = o.RefPlanoNavigation
	}
	if o.Utilizadores != nil {
		toSerialize["utilizadores"] = o.Utilizadores
	}
	return json.Marshal(toSerialize)
}

type NullableEntidade struct {
	value *Entidade
	isSet bool
}

func (v NullableEntidade) Get() *Entidade {
	return v.value
}

func (v *NullableEntidade) Set(val *Entidade) {
	v.value = val
	v.isSet = true
}

func (v NullableEntidade) IsSet() bool {
	return v.isSet
}

func (v *NullableEntidade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntidade(val *Entidade) *NullableEntidade {
	return &NullableEntidade{value: val, isSet: true}
}

func (v NullableEntidade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntidade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

