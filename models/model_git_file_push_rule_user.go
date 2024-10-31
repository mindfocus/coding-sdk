/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitFilePushRuleUser type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitFilePushRuleUser{}

// GitFilePushRuleUser git 文件推送权限用户
type GitFilePushRuleUser struct {
	// 头像路径
	Avatar utils.NullableString `json:"Avatar,omitempty"`
	// 用户全局 key
	GlobalKey utils.NullableString `json:"GlobalKey,omitempty"`
	// 用户名
	Name utils.NullableString `json:"Name,omitempty"`
}

// NewGitFilePushRuleUser instantiates a new GitFilePushRuleUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitFilePushRuleUser() *GitFilePushRuleUser {
	this := GitFilePushRuleUser{}
	var avatar string = ""
	this.Avatar = *utils.NewNullableString(&avatar)
	var globalKey string = ""
	this.GlobalKey = *utils.NewNullableString(&globalKey)
	var name string = ""
	this.Name = *utils.NewNullableString(&name)
	return &this
}

// NewGitFilePushRuleUserWithDefaults instantiates a new GitFilePushRuleUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitFilePushRuleUserWithDefaults() *GitFilePushRuleUser {
	this := GitFilePushRuleUser{}
	var avatar string = ""
	this.Avatar = *utils.NewNullableString(&avatar)
	var globalKey string = ""
	this.GlobalKey = *utils.NewNullableString(&globalKey)
	var name string = ""
	this.Name = *utils.NewNullableString(&name)
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitFilePushRuleUser) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitFilePushRuleUser) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *GitFilePushRuleUser) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given utils.NullableString and assigns it to the Avatar field.
func (o *GitFilePushRuleUser) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *GitFilePushRuleUser) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *GitFilePushRuleUser) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetGlobalKey returns the GlobalKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitFilePushRuleUser) GetGlobalKey() string {
	if o == nil || utils.IsNil(o.GlobalKey.Get()) {
		var ret string
		return ret
	}
	return *o.GlobalKey.Get()
}

// GetGlobalKeyOk returns a tuple with the GlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitFilePushRuleUser) GetGlobalKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalKey.Get(), o.GlobalKey.IsSet()
}

// HasGlobalKey returns a boolean if a field has been set.
func (o *GitFilePushRuleUser) HasGlobalKey() bool {
	if o != nil && o.GlobalKey.IsSet() {
		return true
	}

	return false
}

// SetGlobalKey gets a reference to the given utils.NullableString and assigns it to the GlobalKey field.
func (o *GitFilePushRuleUser) SetGlobalKey(v string) {
	o.GlobalKey.Set(&v)
}
// SetGlobalKeyNil sets the value for GlobalKey to be an explicit nil
func (o *GitFilePushRuleUser) SetGlobalKeyNil() {
	o.GlobalKey.Set(nil)
}

// UnsetGlobalKey ensures that no value is present for GlobalKey, not even an explicit nil
func (o *GitFilePushRuleUser) UnsetGlobalKey() {
	o.GlobalKey.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitFilePushRuleUser) GetName() string {
	if o == nil || utils.IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitFilePushRuleUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GitFilePushRuleUser) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given utils.NullableString and assigns it to the Name field.
func (o *GitFilePushRuleUser) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GitFilePushRuleUser) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GitFilePushRuleUser) UnsetName() {
	o.Name.Unset()
}

func (o GitFilePushRuleUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitFilePushRuleUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Avatar.IsSet() {
		toSerialize["Avatar"] = o.Avatar.Get()
	}
	if o.GlobalKey.IsSet() {
		toSerialize["GlobalKey"] = o.GlobalKey.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableGitFilePushRuleUser struct {
	value *GitFilePushRuleUser
	isSet bool
}

func (v NullableGitFilePushRuleUser) Get() *GitFilePushRuleUser {
	return v.value
}

func (v *NullableGitFilePushRuleUser) Set(val *GitFilePushRuleUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGitFilePushRuleUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGitFilePushRuleUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitFilePushRuleUser(val *GitFilePushRuleUser) *NullableGitFilePushRuleUser {
	return &NullableGitFilePushRuleUser{value: val, isSet: true}
}

func (v NullableGitFilePushRuleUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitFilePushRuleUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


