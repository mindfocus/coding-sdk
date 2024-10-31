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

// checks if the ProjectMemberDepartmentMember type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProjectMemberDepartmentMember{}

// ProjectMemberDepartmentMember struct for ProjectMemberDepartmentMember
type ProjectMemberDepartmentMember struct {
	// 部门成员refId
	RefId utils.NullableInt64 `json:"RefId"`
	// 关联信息
	Refs []ProjectMemberMemberRef `json:"Refs"`
	// 三方头像
	ThirdPartyAvatar utils.NullableString `json:"ThirdPartyAvatar"`
	// 三方名
	ThirdPartyName utils.NullableString `json:"ThirdPartyName"`
	// 三方ID，目前仅支持ldap的用户id信息
	ThirdPartyId *string `json:"ThirdPartyId,omitempty"`
}

type _ProjectMemberDepartmentMember ProjectMemberDepartmentMember

// NewProjectMemberDepartmentMember instantiates a new ProjectMemberDepartmentMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMemberDepartmentMember(refId utils.NullableInt64, refs []ProjectMemberMemberRef, thirdPartyAvatar utils.NullableString, thirdPartyName utils.NullableString) *ProjectMemberDepartmentMember {
	this := ProjectMemberDepartmentMember{}
	this.RefId = refId
	this.Refs = refs
	this.ThirdPartyAvatar = thirdPartyAvatar
	this.ThirdPartyName = thirdPartyName
	var thirdPartyId string = ""
	this.ThirdPartyId = &thirdPartyId
	return &this
}

// NewProjectMemberDepartmentMemberWithDefaults instantiates a new ProjectMemberDepartmentMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMemberDepartmentMemberWithDefaults() *ProjectMemberDepartmentMember {
	this := ProjectMemberDepartmentMember{}
	var thirdPartyAvatar string = ""
	this.ThirdPartyAvatar = *utils.NewNullableString(&thirdPartyAvatar)
	var thirdPartyName string = ""
	this.ThirdPartyName = *utils.NewNullableString(&thirdPartyName)
	var thirdPartyId string = ""
	this.ThirdPartyId = &thirdPartyId
	return &this
}

// GetRefId returns the RefId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ProjectMemberDepartmentMember) GetRefId() int64 {
	if o == nil || o.RefId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMemberDepartmentMember) GetRefIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// SetRefId sets field value
func (o *ProjectMemberDepartmentMember) SetRefId(v int64) {
	o.RefId.Set(&v)
}

// GetRefs returns the Refs field value
// If the value is explicit nil, the zero value for []ProjectMemberMemberRef will be returned
func (o *ProjectMemberDepartmentMember) GetRefs() []ProjectMemberMemberRef {
	if o == nil {
		var ret []ProjectMemberMemberRef
		return ret
	}

	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMemberDepartmentMember) GetRefsOk() ([]ProjectMemberMemberRef, bool) {
	if o == nil || utils.IsNil(o.Refs) {
		return nil, false
	}
	return o.Refs, true
}

// SetRefs sets field value
func (o *ProjectMemberDepartmentMember) SetRefs(v []ProjectMemberMemberRef) {
	o.Refs = v
}

// GetThirdPartyAvatar returns the ThirdPartyAvatar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectMemberDepartmentMember) GetThirdPartyAvatar() string {
	if o == nil || o.ThirdPartyAvatar.Get() == nil {
		var ret string
		return ret
	}

	return *o.ThirdPartyAvatar.Get()
}

// GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMemberDepartmentMember) GetThirdPartyAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyAvatar.Get(), o.ThirdPartyAvatar.IsSet()
}

// SetThirdPartyAvatar sets field value
func (o *ProjectMemberDepartmentMember) SetThirdPartyAvatar(v string) {
	o.ThirdPartyAvatar.Set(&v)
}

// GetThirdPartyName returns the ThirdPartyName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectMemberDepartmentMember) GetThirdPartyName() string {
	if o == nil || o.ThirdPartyName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ThirdPartyName.Get()
}

// GetThirdPartyNameOk returns a tuple with the ThirdPartyName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectMemberDepartmentMember) GetThirdPartyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyName.Get(), o.ThirdPartyName.IsSet()
}

// SetThirdPartyName sets field value
func (o *ProjectMemberDepartmentMember) SetThirdPartyName(v string) {
	o.ThirdPartyName.Set(&v)
}

// GetThirdPartyId returns the ThirdPartyId field value if set, zero value otherwise.
func (o *ProjectMemberDepartmentMember) GetThirdPartyId() string {
	if o == nil || utils.IsNil(o.ThirdPartyId) {
		var ret string
		return ret
	}
	return *o.ThirdPartyId
}

// GetThirdPartyIdOk returns a tuple with the ThirdPartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberDepartmentMember) GetThirdPartyIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThirdPartyId) {
		return nil, false
	}
	return o.ThirdPartyId, true
}

// HasThirdPartyId returns a boolean if a field has been set.
func (o *ProjectMemberDepartmentMember) HasThirdPartyId() bool {
	if o != nil && !utils.IsNil(o.ThirdPartyId) {
		return true
	}

	return false
}

// SetThirdPartyId gets a reference to the given string and assigns it to the ThirdPartyId field.
func (o *ProjectMemberDepartmentMember) SetThirdPartyId(v string) {
	o.ThirdPartyId = &v
}

func (o ProjectMemberDepartmentMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMemberDepartmentMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RefId"] = o.RefId.Get()
	if o.Refs != nil {
		toSerialize["Refs"] = o.Refs
	}
	toSerialize["ThirdPartyAvatar"] = o.ThirdPartyAvatar.Get()
	toSerialize["ThirdPartyName"] = o.ThirdPartyName.Get()
	if !utils.IsNil(o.ThirdPartyId) {
		toSerialize["ThirdPartyId"] = o.ThirdPartyId
	}
	return toSerialize, nil
}

func (o *ProjectMemberDepartmentMember) UnmarshalJSON(data []byte) (err error) {
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

	varProjectMemberDepartmentMember := _ProjectMemberDepartmentMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectMemberDepartmentMember)

	if err != nil {
		return err
	}

	*o = ProjectMemberDepartmentMember(varProjectMemberDepartmentMember)

	return err
}

type NullableProjectMemberDepartmentMember struct {
	value *ProjectMemberDepartmentMember
	isSet bool
}

func (v NullableProjectMemberDepartmentMember) Get() *ProjectMemberDepartmentMember {
	return v.value
}

func (v *NullableProjectMemberDepartmentMember) Set(val *ProjectMemberDepartmentMember) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMemberDepartmentMember) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMemberDepartmentMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMemberDepartmentMember(val *ProjectMemberDepartmentMember) *NullableProjectMemberDepartmentMember {
	return &NullableProjectMemberDepartmentMember{value: val, isSet: true}
}

func (v NullableProjectMemberDepartmentMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMemberDepartmentMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


