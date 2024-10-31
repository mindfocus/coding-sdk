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

// checks if the GrantObjectInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GrantObjectInfo{}

// GrantObjectInfo 授权对象信息
type GrantObjectInfo struct {
	// 创建时间
	CreatedAt int32 `json:"CreatedAt"`
	// 授权对象 ID
	GrantObjectId string `json:"GrantObjectId"`
	// 授权对象名称
	GrantObjectName string `json:"GrantObjectName"`
	// 授权对象类型：USER,USER_GROUP,DEPARTMENT
	GrantScope string `json:"GrantScope"`
}

type _GrantObjectInfo GrantObjectInfo

// NewGrantObjectInfo instantiates a new GrantObjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantObjectInfo(createdAt int32, grantObjectId string, grantObjectName string, grantScope string) *GrantObjectInfo {
	this := GrantObjectInfo{}
	this.CreatedAt = createdAt
	this.GrantObjectId = grantObjectId
	this.GrantObjectName = grantObjectName
	this.GrantScope = grantScope
	return &this
}

// NewGrantObjectInfoWithDefaults instantiates a new GrantObjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantObjectInfoWithDefaults() *GrantObjectInfo {
	this := GrantObjectInfo{}
	var grantObjectId string = ""
	this.GrantObjectId = grantObjectId
	var grantObjectName string = ""
	this.GrantObjectName = grantObjectName
	var grantScope string = ""
	this.GrantScope = grantScope
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *GrantObjectInfo) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GrantObjectInfo) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GrantObjectInfo) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetGrantObjectId returns the GrantObjectId field value
func (o *GrantObjectInfo) GetGrantObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantObjectId
}

// GetGrantObjectIdOk returns a tuple with the GrantObjectId field value
// and a boolean to check if the value has been set.
func (o *GrantObjectInfo) GetGrantObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantObjectId, true
}

// SetGrantObjectId sets field value
func (o *GrantObjectInfo) SetGrantObjectId(v string) {
	o.GrantObjectId = v
}

// GetGrantObjectName returns the GrantObjectName field value
func (o *GrantObjectInfo) GetGrantObjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantObjectName
}

// GetGrantObjectNameOk returns a tuple with the GrantObjectName field value
// and a boolean to check if the value has been set.
func (o *GrantObjectInfo) GetGrantObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantObjectName, true
}

// SetGrantObjectName sets field value
func (o *GrantObjectInfo) SetGrantObjectName(v string) {
	o.GrantObjectName = v
}

// GetGrantScope returns the GrantScope field value
func (o *GrantObjectInfo) GetGrantScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantScope
}

// GetGrantScopeOk returns a tuple with the GrantScope field value
// and a boolean to check if the value has been set.
func (o *GrantObjectInfo) GetGrantScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantScope, true
}

// SetGrantScope sets field value
func (o *GrantObjectInfo) SetGrantScope(v string) {
	o.GrantScope = v
}

func (o GrantObjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantObjectInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CreatedAt"] = o.CreatedAt
	toSerialize["GrantObjectId"] = o.GrantObjectId
	toSerialize["GrantObjectName"] = o.GrantObjectName
	toSerialize["GrantScope"] = o.GrantScope
	return toSerialize, nil
}

func (o *GrantObjectInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CreatedAt",
		"GrantObjectId",
		"GrantObjectName",
		"GrantScope",
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

	varGrantObjectInfo := _GrantObjectInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGrantObjectInfo)

	if err != nil {
		return err
	}

	*o = GrantObjectInfo(varGrantObjectInfo)

	return err
}

type NullableGrantObjectInfo struct {
	value *GrantObjectInfo
	isSet bool
}

func (v NullableGrantObjectInfo) Get() *GrantObjectInfo {
	return v.value
}

func (v *NullableGrantObjectInfo) Set(val *GrantObjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantObjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantObjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantObjectInfo(val *GrantObjectInfo) *NullableGrantObjectInfo {
	return &NullableGrantObjectInfo{value: val, isSet: true}
}

func (v NullableGrantObjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantObjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


