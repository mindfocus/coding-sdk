/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyCdCloudAccountResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyCdCloudAccountResponseData{}

// ModifyCdCloudAccountResponseData ModifyCdCloudAccountResponseData 结构
type ModifyCdCloudAccountResponseData struct {
	CloudAccount *CloudAccount `json:"CloudAccount,omitempty"`
}

// NewModifyCdCloudAccountResponseData instantiates a new ModifyCdCloudAccountResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyCdCloudAccountResponseData() *ModifyCdCloudAccountResponseData {
	this := ModifyCdCloudAccountResponseData{}
	return &this
}

// NewModifyCdCloudAccountResponseDataWithDefaults instantiates a new ModifyCdCloudAccountResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyCdCloudAccountResponseDataWithDefaults() *ModifyCdCloudAccountResponseData {
	this := ModifyCdCloudAccountResponseData{}
	return &this
}

// GetCloudAccount returns the CloudAccount field value if set, zero value otherwise.
func (o *ModifyCdCloudAccountResponseData) GetCloudAccount() CloudAccount {
	if o == nil || utils.IsNil(o.CloudAccount) {
		var ret CloudAccount
		return ret
	}
	return *o.CloudAccount
}

// GetCloudAccountOk returns a tuple with the CloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyCdCloudAccountResponseData) GetCloudAccountOk() (*CloudAccount, bool) {
	if o == nil || utils.IsNil(o.CloudAccount) {
		return nil, false
	}
	return o.CloudAccount, true
}

// HasCloudAccount returns a boolean if a field has been set.
func (o *ModifyCdCloudAccountResponseData) HasCloudAccount() bool {
	if o != nil && !utils.IsNil(o.CloudAccount) {
		return true
	}

	return false
}

// SetCloudAccount gets a reference to the given CloudAccount and assigns it to the CloudAccount field.
func (o *ModifyCdCloudAccountResponseData) SetCloudAccount(v CloudAccount) {
	o.CloudAccount = &v
}

func (o ModifyCdCloudAccountResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyCdCloudAccountResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CloudAccount) {
		toSerialize["CloudAccount"] = o.CloudAccount
	}
	return toSerialize, nil
}

type NullableModifyCdCloudAccountResponseData struct {
	value *ModifyCdCloudAccountResponseData
	isSet bool
}

func (v NullableModifyCdCloudAccountResponseData) Get() *ModifyCdCloudAccountResponseData {
	return v.value
}

func (v *NullableModifyCdCloudAccountResponseData) Set(val *ModifyCdCloudAccountResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCdCloudAccountResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCdCloudAccountResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCdCloudAccountResponseData(val *ModifyCdCloudAccountResponseData) *NullableModifyCdCloudAccountResponseData {
	return &NullableModifyCdCloudAccountResponseData{value: val, isSet: true}
}

func (v NullableModifyCdCloudAccountResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCdCloudAccountResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

