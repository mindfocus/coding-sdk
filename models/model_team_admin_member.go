/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the TeamAdminMember type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TeamAdminMember{}

// TeamAdminMember 团队管理员信息
type TeamAdminMember struct {
	// 头像
	Avatar *string `json:"Avatar,omitempty"`
	// 用户 GK
	GlobalKey *string `json:"GlobalKey,omitempty"`
	// 用户Id
	Id *int32 `json:"Id,omitempty"`
	// 用户名
	Name *string `json:"Name,omitempty"`
	// 用户名拼音
	NamePinYin *string `json:"NamePinYin,omitempty"`
	// 团队Id
	TeamId *int32 `json:"TeamId,omitempty"`
}

// NewTeamAdminMember instantiates a new TeamAdminMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamAdminMember() *TeamAdminMember {
	this := TeamAdminMember{}
	var avatar string = ""
	this.Avatar = &avatar
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	return &this
}

// NewTeamAdminMemberWithDefaults instantiates a new TeamAdminMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamAdminMemberWithDefaults() *TeamAdminMember {
	this := TeamAdminMember{}
	var avatar string = ""
	this.Avatar = &avatar
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *TeamAdminMember) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamAdminMember) HasAvatar() bool {
	if o != nil && !utils.IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *TeamAdminMember) SetAvatar(v string) {
	o.Avatar = &v
}

// GetGlobalKey returns the GlobalKey field value if set, zero value otherwise.
func (o *TeamAdminMember) GetGlobalKey() string {
	if o == nil || utils.IsNil(o.GlobalKey) {
		var ret string
		return ret
	}
	return *o.GlobalKey
}

// GetGlobalKeyOk returns a tuple with the GlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GlobalKey) {
		return nil, false
	}
	return o.GlobalKey, true
}

// HasGlobalKey returns a boolean if a field has been set.
func (o *TeamAdminMember) HasGlobalKey() bool {
	if o != nil && !utils.IsNil(o.GlobalKey) {
		return true
	}

	return false
}

// SetGlobalKey gets a reference to the given string and assigns it to the GlobalKey field.
func (o *TeamAdminMember) SetGlobalKey(v string) {
	o.GlobalKey = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamAdminMember) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamAdminMember) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TeamAdminMember) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamAdminMember) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamAdminMember) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamAdminMember) SetName(v string) {
	o.Name = &v
}

// GetNamePinYin returns the NamePinYin field value if set, zero value otherwise.
func (o *TeamAdminMember) GetNamePinYin() string {
	if o == nil || utils.IsNil(o.NamePinYin) {
		var ret string
		return ret
	}
	return *o.NamePinYin
}

// GetNamePinYinOk returns a tuple with the NamePinYin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetNamePinYinOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NamePinYin) {
		return nil, false
	}
	return o.NamePinYin, true
}

// HasNamePinYin returns a boolean if a field has been set.
func (o *TeamAdminMember) HasNamePinYin() bool {
	if o != nil && !utils.IsNil(o.NamePinYin) {
		return true
	}

	return false
}

// SetNamePinYin gets a reference to the given string and assigns it to the NamePinYin field.
func (o *TeamAdminMember) SetNamePinYin(v string) {
	o.NamePinYin = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *TeamAdminMember) GetTeamId() int32 {
	if o == nil || utils.IsNil(o.TeamId) {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamAdminMember) GetTeamIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamAdminMember) HasTeamId() bool {
	if o != nil && !utils.IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *TeamAdminMember) SetTeamId(v int32) {
	o.TeamId = &v
}

func (o TeamAdminMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamAdminMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
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
	if !utils.IsNil(o.TeamId) {
		toSerialize["TeamId"] = o.TeamId
	}
	return toSerialize, nil
}

type NullableTeamAdminMember struct {
	value *TeamAdminMember
	isSet bool
}

func (v NullableTeamAdminMember) Get() *TeamAdminMember {
	return v.value
}

func (v *NullableTeamAdminMember) Set(val *TeamAdminMember) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamAdminMember) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamAdminMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamAdminMember(val *TeamAdminMember) *NullableTeamAdminMember {
	return &NullableTeamAdminMember{value: val, isSet: true}
}

func (v NullableTeamAdminMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamAdminMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


