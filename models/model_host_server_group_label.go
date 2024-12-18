/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the HostServerGroupLabel type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HostServerGroupLabel{}

// HostServerGroupLabel HostServerGroupLabel 结构
type HostServerGroupLabel struct {
	// 主机组标签键
	Key string `json:"Key"`
	// 主机组标签值
	Value string `json:"Value"`
}

type _HostServerGroupLabel HostServerGroupLabel

// NewHostServerGroupLabel instantiates a new HostServerGroupLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostServerGroupLabel(key string, value string) *HostServerGroupLabel {
	this := HostServerGroupLabel{}
	this.Key = key
	this.Value = value
	return &this
}

// NewHostServerGroupLabelWithDefaults instantiates a new HostServerGroupLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostServerGroupLabelWithDefaults() *HostServerGroupLabel {
	this := HostServerGroupLabel{}
	var key string = ""
	this.Key = key
	var value string = ""
	this.Value = value
	return &this
}

// GetKey returns the Key field value
func (o *HostServerGroupLabel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *HostServerGroupLabel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *HostServerGroupLabel) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *HostServerGroupLabel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *HostServerGroupLabel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *HostServerGroupLabel) SetValue(v string) {
	o.Value = v
}

func (o HostServerGroupLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostServerGroupLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Key"] = o.Key
	toSerialize["Value"] = o.Value
	return toSerialize, nil
}

func (o *HostServerGroupLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Key",
		"Value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHostServerGroupLabel := _HostServerGroupLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHostServerGroupLabel)

	if err != nil {
		return err
	}

	*o = HostServerGroupLabel(varHostServerGroupLabel)

	return err
}

type NullableHostServerGroupLabel struct {
	value *HostServerGroupLabel
	isSet bool
}

func (v NullableHostServerGroupLabel) Get() *HostServerGroupLabel {
	return v.value
}

func (v *NullableHostServerGroupLabel) Set(val *HostServerGroupLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableHostServerGroupLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableHostServerGroupLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostServerGroupLabel(val *HostServerGroupLabel) *NullableHostServerGroupLabel {
	return &NullableHostServerGroupLabel{value: val, isSet: true}
}

func (v NullableHostServerGroupLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostServerGroupLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


