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

// checks if the PrincipalDel type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalDel{}

// PrincipalDel 删除项目成员主体 成员对象
type PrincipalDel struct {
	// 主体ID
	PrincipalId string `json:"PrincipalId"`
	// 主体类型，支持USER_GROUP、DEPARTMENT、USER
	PrincipalType string `json:"PrincipalType"`
}

type _PrincipalDel PrincipalDel

// NewPrincipalDel instantiates a new PrincipalDel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalDel(principalId string, principalType string) *PrincipalDel {
	this := PrincipalDel{}
	this.PrincipalId = principalId
	this.PrincipalType = principalType
	return &this
}

// NewPrincipalDelWithDefaults instantiates a new PrincipalDel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalDelWithDefaults() *PrincipalDel {
	this := PrincipalDel{}
	var principalId string = ""
	this.PrincipalId = principalId
	var principalType string = ""
	this.PrincipalType = principalType
	return &this
}

// GetPrincipalId returns the PrincipalId field value
func (o *PrincipalDel) GetPrincipalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value
// and a boolean to check if the value has been set.
func (o *PrincipalDel) GetPrincipalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalId, true
}

// SetPrincipalId sets field value
func (o *PrincipalDel) SetPrincipalId(v string) {
	o.PrincipalId = v
}

// GetPrincipalType returns the PrincipalType field value
func (o *PrincipalDel) GetPrincipalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalType
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value
// and a boolean to check if the value has been set.
func (o *PrincipalDel) GetPrincipalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalType, true
}

// SetPrincipalType sets field value
func (o *PrincipalDel) SetPrincipalType(v string) {
	o.PrincipalType = v
}

func (o PrincipalDel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalDel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PrincipalId"] = o.PrincipalId
	toSerialize["PrincipalType"] = o.PrincipalType
	return toSerialize, nil
}

func (o *PrincipalDel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PrincipalId",
		"PrincipalType",
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

	varPrincipalDel := _PrincipalDel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrincipalDel)

	if err != nil {
		return err
	}

	*o = PrincipalDel(varPrincipalDel)

	return err
}

type NullablePrincipalDel struct {
	value *PrincipalDel
	isSet bool
}

func (v NullablePrincipalDel) Get() *PrincipalDel {
	return v.value
}

func (v *NullablePrincipalDel) Set(val *PrincipalDel) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalDel) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalDel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalDel(val *PrincipalDel) *NullablePrincipalDel {
	return &NullablePrincipalDel{value: val, isSet: true}
}

func (v NullablePrincipalDel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalDel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


