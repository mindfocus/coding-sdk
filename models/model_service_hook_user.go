/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ServiceHookUser type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceHookUser{}

// ServiceHookUser ServiceHook User 对象
type ServiceHookUser struct {
	// 用户头像
	Avatar *string `json:"Avatar,omitempty"`
	// 用户编号
	Id *int64 `json:"Id,omitempty"`
	// 用户名
	Name *string `json:"Name,omitempty"`
	// 用户名拼音
	NamePinyin *string `json:"NamePinyin,omitempty"`
}

// NewServiceHookUser instantiates a new ServiceHookUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHookUser() *ServiceHookUser {
	this := ServiceHookUser{}
	var avatar string = ""
	this.Avatar = &avatar
	var name string = ""
	this.Name = &name
	var namePinyin string = ""
	this.NamePinyin = &namePinyin
	return &this
}

// NewServiceHookUserWithDefaults instantiates a new ServiceHookUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHookUserWithDefaults() *ServiceHookUser {
	this := ServiceHookUser{}
	var avatar string = ""
	this.Avatar = &avatar
	var name string = ""
	this.Name = &name
	var namePinyin string = ""
	this.NamePinyin = &namePinyin
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ServiceHookUser) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookUser) GetAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ServiceHookUser) HasAvatar() bool {
	if o != nil && !utils.IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *ServiceHookUser) SetAvatar(v string) {
	o.Avatar = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceHookUser) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookUser) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceHookUser) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ServiceHookUser) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceHookUser) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookUser) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceHookUser) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceHookUser) SetName(v string) {
	o.Name = &v
}

// GetNamePinyin returns the NamePinyin field value if set, zero value otherwise.
func (o *ServiceHookUser) GetNamePinyin() string {
	if o == nil || utils.IsNil(o.NamePinyin) {
		var ret string
		return ret
	}
	return *o.NamePinyin
}

// GetNamePinyinOk returns a tuple with the NamePinyin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookUser) GetNamePinyinOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NamePinyin) {
		return nil, false
	}
	return o.NamePinyin, true
}

// HasNamePinyin returns a boolean if a field has been set.
func (o *ServiceHookUser) HasNamePinyin() bool {
	if o != nil && !utils.IsNil(o.NamePinyin) {
		return true
	}

	return false
}

// SetNamePinyin gets a reference to the given string and assigns it to the NamePinyin field.
func (o *ServiceHookUser) SetNamePinyin(v string) {
	o.NamePinyin = &v
}

func (o ServiceHookUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHookUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.NamePinyin) {
		toSerialize["NamePinyin"] = o.NamePinyin
	}
	return toSerialize, nil
}

type NullableServiceHookUser struct {
	value *ServiceHookUser
	isSet bool
}

func (v NullableServiceHookUser) Get() *ServiceHookUser {
	return v.value
}

func (v *NullableServiceHookUser) Set(val *ServiceHookUser) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHookUser) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHookUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHookUser(val *ServiceHookUser) *NullableServiceHookUser {
	return &NullableServiceHookUser{value: val, isSet: true}
}

func (v NullableServiceHookUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHookUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


