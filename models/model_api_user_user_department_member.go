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

// checks if the ApiUserUserDepartmentMember type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiUserUserDepartmentMember{}

// ApiUserUserDepartmentMember struct for ApiUserUserDepartmentMember
type ApiUserUserDepartmentMember struct {
	// 部门成员refId
	RefId int64 `json:"RefId"`
	// 关联信息
	Refs []ApiUserMemberRef `json:"Refs"`
	// 三方头像
	ThirdPartyAvatar string `json:"ThirdPartyAvatar"`
	// 三方名
	ThirdPartyName string `json:"ThirdPartyName"`
	// 三方ID，目前仅支持ldap的用户id信息
	ThirdPartyId *string `json:"ThirdPartyId,omitempty"`
}

type _ApiUserUserDepartmentMember ApiUserUserDepartmentMember

// NewApiUserUserDepartmentMember instantiates a new ApiUserUserDepartmentMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserUserDepartmentMember(refId int64, refs []ApiUserMemberRef, thirdPartyAvatar string, thirdPartyName string) *ApiUserUserDepartmentMember {
	this := ApiUserUserDepartmentMember{}
	this.RefId = refId
	this.Refs = refs
	this.ThirdPartyAvatar = thirdPartyAvatar
	this.ThirdPartyName = thirdPartyName
	var thirdPartyId string = ""
	this.ThirdPartyId = &thirdPartyId
	return &this
}

// NewApiUserUserDepartmentMemberWithDefaults instantiates a new ApiUserUserDepartmentMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserUserDepartmentMemberWithDefaults() *ApiUserUserDepartmentMember {
	this := ApiUserUserDepartmentMember{}
	var refId int64 = 0
	this.RefId = refId
	var thirdPartyAvatar string = ""
	this.ThirdPartyAvatar = thirdPartyAvatar
	var thirdPartyName string = ""
	this.ThirdPartyName = thirdPartyName
	var thirdPartyId string = ""
	this.ThirdPartyId = &thirdPartyId
	return &this
}

// GetRefId returns the RefId field value
func (o *ApiUserUserDepartmentMember) GetRefId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value
// and a boolean to check if the value has been set.
func (o *ApiUserUserDepartmentMember) GetRefIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefId, true
}

// SetRefId sets field value
func (o *ApiUserUserDepartmentMember) SetRefId(v int64) {
	o.RefId = v
}

// GetRefs returns the Refs field value
func (o *ApiUserUserDepartmentMember) GetRefs() []ApiUserMemberRef {
	if o == nil {
		var ret []ApiUserMemberRef
		return ret
	}

	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value
// and a boolean to check if the value has been set.
func (o *ApiUserUserDepartmentMember) GetRefsOk() ([]ApiUserMemberRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refs, true
}

// SetRefs sets field value
func (o *ApiUserUserDepartmentMember) SetRefs(v []ApiUserMemberRef) {
	o.Refs = v
}

// GetThirdPartyAvatar returns the ThirdPartyAvatar field value
func (o *ApiUserUserDepartmentMember) GetThirdPartyAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyAvatar
}

// GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field value
// and a boolean to check if the value has been set.
func (o *ApiUserUserDepartmentMember) GetThirdPartyAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyAvatar, true
}

// SetThirdPartyAvatar sets field value
func (o *ApiUserUserDepartmentMember) SetThirdPartyAvatar(v string) {
	o.ThirdPartyAvatar = v
}

// GetThirdPartyName returns the ThirdPartyName field value
func (o *ApiUserUserDepartmentMember) GetThirdPartyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyName
}

// GetThirdPartyNameOk returns a tuple with the ThirdPartyName field value
// and a boolean to check if the value has been set.
func (o *ApiUserUserDepartmentMember) GetThirdPartyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyName, true
}

// SetThirdPartyName sets field value
func (o *ApiUserUserDepartmentMember) SetThirdPartyName(v string) {
	o.ThirdPartyName = v
}

// GetThirdPartyId returns the ThirdPartyId field value if set, zero value otherwise.
func (o *ApiUserUserDepartmentMember) GetThirdPartyId() string {
	if o == nil || utils.IsNil(o.ThirdPartyId) {
		var ret string
		return ret
	}
	return *o.ThirdPartyId
}

// GetThirdPartyIdOk returns a tuple with the ThirdPartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserUserDepartmentMember) GetThirdPartyIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThirdPartyId) {
		return nil, false
	}
	return o.ThirdPartyId, true
}

// HasThirdPartyId returns a boolean if a field has been set.
func (o *ApiUserUserDepartmentMember) HasThirdPartyId() bool {
	if o != nil && !utils.IsNil(o.ThirdPartyId) {
		return true
	}

	return false
}

// SetThirdPartyId gets a reference to the given string and assigns it to the ThirdPartyId field.
func (o *ApiUserUserDepartmentMember) SetThirdPartyId(v string) {
	o.ThirdPartyId = &v
}

func (o ApiUserUserDepartmentMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserUserDepartmentMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RefId"] = o.RefId
	toSerialize["Refs"] = o.Refs
	toSerialize["ThirdPartyAvatar"] = o.ThirdPartyAvatar
	toSerialize["ThirdPartyName"] = o.ThirdPartyName
	if !utils.IsNil(o.ThirdPartyId) {
		toSerialize["ThirdPartyId"] = o.ThirdPartyId
	}
	return toSerialize, nil
}

func (o *ApiUserUserDepartmentMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"RefId",
		"Refs",
		"ThirdPartyAvatar",
		"ThirdPartyName",
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

	varApiUserUserDepartmentMember := _ApiUserUserDepartmentMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUserUserDepartmentMember)

	if err != nil {
		return err
	}

	*o = ApiUserUserDepartmentMember(varApiUserUserDepartmentMember)

	return err
}

type NullableApiUserUserDepartmentMember struct {
	value *ApiUserUserDepartmentMember
	isSet bool
}

func (v NullableApiUserUserDepartmentMember) Get() *ApiUserUserDepartmentMember {
	return v.value
}

func (v *NullableApiUserUserDepartmentMember) Set(val *ApiUserUserDepartmentMember) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserUserDepartmentMember) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserUserDepartmentMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserUserDepartmentMember(val *ApiUserUserDepartmentMember) *NullableApiUserUserDepartmentMember {
	return &NullableApiUserUserDepartmentMember{value: val, isSet: true}
}

func (v NullableApiUserUserDepartmentMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserUserDepartmentMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


