/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Credential type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Credential{}

// Credential 凭据结构
type Credential struct {
	// 凭据唯一 ID
	CredentialId *string `json:"CredentialId,omitempty"`
	// 凭据名称
	Name *string `json:"Name,omitempty"`
}

// NewCredential instantiates a new Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredential() *Credential {
	this := Credential{}
	var credentialId string = ""
	this.CredentialId = &credentialId
	var name string = ""
	this.Name = &name
	return &this
}

// NewCredentialWithDefaults instantiates a new Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialWithDefaults() *Credential {
	this := Credential{}
	var credentialId string = ""
	this.CredentialId = &credentialId
	var name string = ""
	this.Name = &name
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *Credential) GetCredentialId() string {
	if o == nil || utils.IsNil(o.CredentialId) {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetCredentialIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CredentialId) {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *Credential) HasCredentialId() bool {
	if o != nil && !utils.IsNil(o.CredentialId) {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *Credential) SetCredentialId(v string) {
	o.CredentialId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Credential) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Credential) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Credential) SetName(v string) {
	o.Name = &v
}

func (o Credential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CredentialId) {
		toSerialize["CredentialId"] = o.CredentialId
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCredential struct {
	value *Credential
	isSet bool
}

func (v NullableCredential) Get() *Credential {
	return v.value
}

func (v *NullableCredential) Set(val *Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredential(val *Credential) *NullableCredential {
	return &NullableCredential{value: val, isSet: true}
}

func (v NullableCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


