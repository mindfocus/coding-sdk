/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the UserData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserData{}

// UserData 用户成员信息
type UserData struct {
	// 头像
	Avatar *string `json:"Avatar,omitempty"`
	// 邮箱
	Email *string `json:"Email,omitempty"`
	// 邮箱是否验证 0 否 /1 是
	EmailValidation *int32 `json:"EmailValidation,omitempty"`
	// 用户 GK
	GlobalKey *string `json:"GlobalKey,omitempty"`
	// 用户Id
	Id *int32 `json:"Id,omitempty"`
	// 用户名
	Name *string `json:"Name,omitempty"`
	// 用户名拼音
	NamePinYin *string `json:"NamePinYin,omitempty"`
	// 手机号
	Phone *string `json:"Phone,omitempty"`
	// 手机是否验证 0 否 /1 是
	PhoneValidation *int32 `json:"PhoneValidation,omitempty"`
	// 用户组
	Roles []Role `json:"Roles,omitempty"`
	// 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队
	Status *int64 `json:"Status,omitempty"`
	// 团队Id
	TeamId *int32 `json:"TeamId,omitempty"`
	DepartmentMember *ProjectMemberDepartmentMember `json:"DepartmentMember,omitempty"`
	// 团队用户扩展字段、唯一。非必填
	UniqueExtField *string `json:"UniqueExtField,omitempty"`
}

// NewUserData instantiates a new UserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserData() *UserData {
	this := UserData{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var phone string = ""
	this.Phone = &phone
	var teamId int32 = 0
	this.TeamId = &teamId
	var uniqueExtField string = ""
	this.UniqueExtField = &uniqueExtField
	return &this
}

// NewUserDataWithDefaults instantiates a new UserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataWithDefaults() *UserData {
	this := UserData{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var phone string = ""
	this.Phone = &phone
	var teamId int32 = 0
	this.TeamId = &teamId
	var uniqueExtField string = ""
	this.UniqueExtField = &uniqueExtField
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UserData) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserData) HasAvatar() bool {
	if o != nil && !utils.IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UserData) SetAvatar(v string) {
	o.Avatar = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserData) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserData) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserData) SetEmail(v string) {
	o.Email = &v
}

// GetEmailValidation returns the EmailValidation field value if set, zero value otherwise.
func (o *UserData) GetEmailValidation() int32 {
	if o == nil || utils.IsNil(o.EmailValidation) {
		var ret int32
		return ret
	}
	return *o.EmailValidation
}

// GetEmailValidationOk returns a tuple with the EmailValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetEmailValidationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EmailValidation) {
		return nil, false
	}
	return o.EmailValidation, true
}

// HasEmailValidation returns a boolean if a field has been set.
func (o *UserData) HasEmailValidation() bool {
	if o != nil && !utils.IsNil(o.EmailValidation) {
		return true
	}

	return false
}

// SetEmailValidation gets a reference to the given int32 and assigns it to the EmailValidation field.
func (o *UserData) SetEmailValidation(v int32) {
	o.EmailValidation = &v
}

// GetGlobalKey returns the GlobalKey field value if set, zero value otherwise.
func (o *UserData) GetGlobalKey() string {
	if o == nil || utils.IsNil(o.GlobalKey) {
		var ret string
		return ret
	}
	return *o.GlobalKey
}

// GetGlobalKeyOk returns a tuple with the GlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GlobalKey) {
		return nil, false
	}
	return o.GlobalKey, true
}

// HasGlobalKey returns a boolean if a field has been set.
func (o *UserData) HasGlobalKey() bool {
	if o != nil && !utils.IsNil(o.GlobalKey) {
		return true
	}

	return false
}

// SetGlobalKey gets a reference to the given string and assigns it to the GlobalKey field.
func (o *UserData) SetGlobalKey(v string) {
	o.GlobalKey = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserData) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserData) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserData) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserData) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserData) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserData) SetName(v string) {
	o.Name = &v
}

// GetNamePinYin returns the NamePinYin field value if set, zero value otherwise.
func (o *UserData) GetNamePinYin() string {
	if o == nil || utils.IsNil(o.NamePinYin) {
		var ret string
		return ret
	}
	return *o.NamePinYin
}

// GetNamePinYinOk returns a tuple with the NamePinYin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetNamePinYinOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NamePinYin) {
		return nil, false
	}
	return o.NamePinYin, true
}

// HasNamePinYin returns a boolean if a field has been set.
func (o *UserData) HasNamePinYin() bool {
	if o != nil && !utils.IsNil(o.NamePinYin) {
		return true
	}

	return false
}

// SetNamePinYin gets a reference to the given string and assigns it to the NamePinYin field.
func (o *UserData) SetNamePinYin(v string) {
	o.NamePinYin = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UserData) GetPhone() string {
	if o == nil || utils.IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetPhoneOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UserData) HasPhone() bool {
	if o != nil && !utils.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UserData) SetPhone(v string) {
	o.Phone = &v
}

// GetPhoneValidation returns the PhoneValidation field value if set, zero value otherwise.
func (o *UserData) GetPhoneValidation() int32 {
	if o == nil || utils.IsNil(o.PhoneValidation) {
		var ret int32
		return ret
	}
	return *o.PhoneValidation
}

// GetPhoneValidationOk returns a tuple with the PhoneValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetPhoneValidationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PhoneValidation) {
		return nil, false
	}
	return o.PhoneValidation, true
}

// HasPhoneValidation returns a boolean if a field has been set.
func (o *UserData) HasPhoneValidation() bool {
	if o != nil && !utils.IsNil(o.PhoneValidation) {
		return true
	}

	return false
}

// SetPhoneValidation gets a reference to the given int32 and assigns it to the PhoneValidation field.
func (o *UserData) SetPhoneValidation(v int32) {
	o.PhoneValidation = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserData) GetRoles() []Role {
	if o == nil || utils.IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetRolesOk() ([]Role, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserData) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *UserData) SetRoles(v []Role) {
	o.Roles = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserData) GetStatus() int64 {
	if o == nil || utils.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetStatusOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserData) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *UserData) SetStatus(v int64) {
	o.Status = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *UserData) GetTeamId() int32 {
	if o == nil || utils.IsNil(o.TeamId) {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetTeamIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *UserData) HasTeamId() bool {
	if o != nil && !utils.IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *UserData) SetTeamId(v int32) {
	o.TeamId = &v
}

// GetDepartmentMember returns the DepartmentMember field value if set, zero value otherwise.
func (o *UserData) GetDepartmentMember() ProjectMemberDepartmentMember {
	if o == nil || utils.IsNil(o.DepartmentMember) {
		var ret ProjectMemberDepartmentMember
		return ret
	}
	return *o.DepartmentMember
}

// GetDepartmentMemberOk returns a tuple with the DepartmentMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetDepartmentMemberOk() (*ProjectMemberDepartmentMember, bool) {
	if o == nil || utils.IsNil(o.DepartmentMember) {
		return nil, false
	}
	return o.DepartmentMember, true
}

// HasDepartmentMember returns a boolean if a field has been set.
func (o *UserData) HasDepartmentMember() bool {
	if o != nil && !utils.IsNil(o.DepartmentMember) {
		return true
	}

	return false
}

// SetDepartmentMember gets a reference to the given ProjectMemberDepartmentMember and assigns it to the DepartmentMember field.
func (o *UserData) SetDepartmentMember(v ProjectMemberDepartmentMember) {
	o.DepartmentMember = &v
}

// GetUniqueExtField returns the UniqueExtField field value if set, zero value otherwise.
func (o *UserData) GetUniqueExtField() string {
	if o == nil || utils.IsNil(o.UniqueExtField) {
		var ret string
		return ret
	}
	return *o.UniqueExtField
}

// GetUniqueExtFieldOk returns a tuple with the UniqueExtField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserData) GetUniqueExtFieldOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UniqueExtField) {
		return nil, false
	}
	return o.UniqueExtField, true
}

// HasUniqueExtField returns a boolean if a field has been set.
func (o *UserData) HasUniqueExtField() bool {
	if o != nil && !utils.IsNil(o.UniqueExtField) {
		return true
	}

	return false
}

// SetUniqueExtField gets a reference to the given string and assigns it to the UniqueExtField field.
func (o *UserData) SetUniqueExtField(v string) {
	o.UniqueExtField = &v
}

func (o UserData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
	}
	if !utils.IsNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	if !utils.IsNil(o.EmailValidation) {
		toSerialize["EmailValidation"] = o.EmailValidation
	}
	if !utils.IsNil(o.GlobalKey) {
		toSerialize["GlobalKey"] = o.GlobalKey
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.NamePinYin) {
		toSerialize["NamePinYin"] = o.NamePinYin
	}
	if !utils.IsNil(o.Phone) {
		toSerialize["Phone"] = o.Phone
	}
	if !utils.IsNil(o.PhoneValidation) {
		toSerialize["PhoneValidation"] = o.PhoneValidation
	}
	if !utils.IsNil(o.Roles) {
		toSerialize["Roles"] = o.Roles
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !utils.IsNil(o.TeamId) {
		toSerialize["TeamId"] = o.TeamId
	}
	if !utils.IsNil(o.DepartmentMember) {
		toSerialize["DepartmentMember"] = o.DepartmentMember
	}
	if !utils.IsNil(o.UniqueExtField) {
		toSerialize["UniqueExtField"] = o.UniqueExtField
	}
	return toSerialize, nil
}

type NullableUserData struct {
	value *UserData
	isSet bool
}

func (v NullableUserData) Get() *UserData {
	return v.value
}

func (v *NullableUserData) Set(val *UserData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserData(val *UserData) *NullableUserData {
	return &NullableUserData{value: val, isSet: true}
}

func (v NullableUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

