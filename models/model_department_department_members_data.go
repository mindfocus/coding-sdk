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

// checks if the DepartmentDepartmentMembersData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepartmentDepartmentMembersData{}

// DepartmentDepartmentMembersData 部门成员信息
type DepartmentDepartmentMembersData struct {
	// 是否授权 实际表示已加入或已被授权，如果要表示 已授权未加入(allowed && userId>0)
	Allowed bool `json:"Allowed"`
	// 头像
	Avatar string `json:"Avatar"`
	// 是否被锁定
	Locked bool `json:"Locked"`
	// 成员名
	Name string `json:"Name"`
	// 是否部门所有者
	Owner bool `json:"Owner"`
	// 部门用户的refId
	RefId int64 `json:"RefId"`
	// 关联的 部门等信息
	Refs []DepartmentDepartmentMemberRef `json:"Refs"`
	// 三方头像
	ThirdPartyAvatar string `json:"ThirdPartyAvatar"`
	// 三方名
	ThirdPartyName string `json:"ThirdPartyName"`
	// 用户GK
	UserGlobalKey string `json:"UserGlobalKey"`
	// 用户ID
	UserId int64 `json:"UserId"`
	// 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队
	UserStatus int64 `json:"UserStatus"`
}

type _DepartmentDepartmentMembersData DepartmentDepartmentMembersData

// NewDepartmentDepartmentMembersData instantiates a new DepartmentDepartmentMembersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentDepartmentMembersData(allowed bool, avatar string, locked bool, name string, owner bool, refId int64, refs []DepartmentDepartmentMemberRef, thirdPartyAvatar string, thirdPartyName string, userGlobalKey string, userId int64, userStatus int64) *DepartmentDepartmentMembersData {
	this := DepartmentDepartmentMembersData{}
	this.Allowed = allowed
	this.Avatar = avatar
	this.Locked = locked
	this.Name = name
	this.Owner = owner
	this.RefId = refId
	this.Refs = refs
	this.ThirdPartyAvatar = thirdPartyAvatar
	this.ThirdPartyName = thirdPartyName
	this.UserGlobalKey = userGlobalKey
	this.UserId = userId
	this.UserStatus = userStatus
	return &this
}

// NewDepartmentDepartmentMembersDataWithDefaults instantiates a new DepartmentDepartmentMembersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentDepartmentMembersDataWithDefaults() *DepartmentDepartmentMembersData {
	this := DepartmentDepartmentMembersData{}
	var allowed bool = false
	this.Allowed = allowed
	var avatar string = ""
	this.Avatar = avatar
	var locked bool = false
	this.Locked = locked
	var name string = ""
	this.Name = name
	var owner bool = false
	this.Owner = owner
	var refId int64 = 0
	this.RefId = refId
	var thirdPartyAvatar string = ""
	this.ThirdPartyAvatar = thirdPartyAvatar
	var thirdPartyName string = ""
	this.ThirdPartyName = thirdPartyName
	var userGlobalKey string = ""
	this.UserGlobalKey = userGlobalKey
	var userId int64 = 0
	this.UserId = userId
	var userStatus int64 = 0
	this.UserStatus = userStatus
	return &this
}

// GetAllowed returns the Allowed field value
func (o *DepartmentDepartmentMembersData) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *DepartmentDepartmentMembersData) SetAllowed(v bool) {
	o.Allowed = v
}

// GetAvatar returns the Avatar field value
func (o *DepartmentDepartmentMembersData) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *DepartmentDepartmentMembersData) SetAvatar(v string) {
	o.Avatar = v
}

// GetLocked returns the Locked field value
func (o *DepartmentDepartmentMembersData) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *DepartmentDepartmentMembersData) SetLocked(v bool) {
	o.Locked = v
}

// GetName returns the Name field value
func (o *DepartmentDepartmentMembersData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DepartmentDepartmentMembersData) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value
func (o *DepartmentDepartmentMembersData) GetOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *DepartmentDepartmentMembersData) SetOwner(v bool) {
	o.Owner = v
}

// GetRefId returns the RefId field value
func (o *DepartmentDepartmentMembersData) GetRefId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetRefIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefId, true
}

// SetRefId sets field value
func (o *DepartmentDepartmentMembersData) SetRefId(v int64) {
	o.RefId = v
}

// GetRefs returns the Refs field value
func (o *DepartmentDepartmentMembersData) GetRefs() []DepartmentDepartmentMemberRef {
	if o == nil {
		var ret []DepartmentDepartmentMemberRef
		return ret
	}

	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetRefsOk() ([]DepartmentDepartmentMemberRef, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refs, true
}

// SetRefs sets field value
func (o *DepartmentDepartmentMembersData) SetRefs(v []DepartmentDepartmentMemberRef) {
	o.Refs = v
}

// GetThirdPartyAvatar returns the ThirdPartyAvatar field value
func (o *DepartmentDepartmentMembersData) GetThirdPartyAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyAvatar
}

// GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetThirdPartyAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyAvatar, true
}

// SetThirdPartyAvatar sets field value
func (o *DepartmentDepartmentMembersData) SetThirdPartyAvatar(v string) {
	o.ThirdPartyAvatar = v
}

// GetThirdPartyName returns the ThirdPartyName field value
func (o *DepartmentDepartmentMembersData) GetThirdPartyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyName
}

// GetThirdPartyNameOk returns a tuple with the ThirdPartyName field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetThirdPartyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyName, true
}

// SetThirdPartyName sets field value
func (o *DepartmentDepartmentMembersData) SetThirdPartyName(v string) {
	o.ThirdPartyName = v
}

// GetUserGlobalKey returns the UserGlobalKey field value
func (o *DepartmentDepartmentMembersData) GetUserGlobalKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserGlobalKey
}

// GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetUserGlobalKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGlobalKey, true
}

// SetUserGlobalKey sets field value
func (o *DepartmentDepartmentMembersData) SetUserGlobalKey(v string) {
	o.UserGlobalKey = v
}

// GetUserId returns the UserId field value
func (o *DepartmentDepartmentMembersData) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DepartmentDepartmentMembersData) SetUserId(v int64) {
	o.UserId = v
}

// GetUserStatus returns the UserStatus field value
func (o *DepartmentDepartmentMembersData) GetUserStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentMembersData) GetUserStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserStatus, true
}

// SetUserStatus sets field value
func (o *DepartmentDepartmentMembersData) SetUserStatus(v int64) {
	o.UserStatus = v
}

func (o DepartmentDepartmentMembersData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentDepartmentMembersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Allowed"] = o.Allowed
	toSerialize["Avatar"] = o.Avatar
	toSerialize["Locked"] = o.Locked
	toSerialize["Name"] = o.Name
	toSerialize["Owner"] = o.Owner
	toSerialize["RefId"] = o.RefId
	toSerialize["Refs"] = o.Refs
	toSerialize["ThirdPartyAvatar"] = o.ThirdPartyAvatar
	toSerialize["ThirdPartyName"] = o.ThirdPartyName
	toSerialize["UserGlobalKey"] = o.UserGlobalKey
	toSerialize["UserId"] = o.UserId
	toSerialize["UserStatus"] = o.UserStatus
	return toSerialize, nil
}

func (o *DepartmentDepartmentMembersData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Allowed",
		"Avatar",
		"Locked",
		"Name",
		"Owner",
		"RefId",
		"Refs",
		"ThirdPartyAvatar",
		"ThirdPartyName",
		"UserGlobalKey",
		"UserId",
		"UserStatus",
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

	varDepartmentDepartmentMembersData := _DepartmentDepartmentMembersData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepartmentDepartmentMembersData)

	if err != nil {
		return err
	}

	*o = DepartmentDepartmentMembersData(varDepartmentDepartmentMembersData)

	return err
}

type NullableDepartmentDepartmentMembersData struct {
	value *DepartmentDepartmentMembersData
	isSet bool
}

func (v NullableDepartmentDepartmentMembersData) Get() *DepartmentDepartmentMembersData {
	return v.value
}

func (v *NullableDepartmentDepartmentMembersData) Set(val *DepartmentDepartmentMembersData) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentDepartmentMembersData) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentDepartmentMembersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentDepartmentMembersData(val *DepartmentDepartmentMembersData) *NullableDepartmentDepartmentMembersData {
	return &NullableDepartmentDepartmentMembersData{value: val, isSet: true}
}

func (v NullableDepartmentDepartmentMembersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentDepartmentMembersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


