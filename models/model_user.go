/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &User{}

// User 用户信息
type User struct {
	// 头像地址
	Avatar *string `json:"Avatar,omitempty"`
	// 邮箱
	Email *string `json:"Email,omitempty"`
	// gk
	GlobalKey *string `json:"GlobalKey,omitempty"`
	// userId
	Id *int32 `json:"Id,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	// 电话
	Phone *string `json:"Phone,omitempty"`
	// 状态
	Status utils.NullableInt32 `json:"Status,omitempty"`
	// 团队gk
	TeamGlobalKey *string `json:"TeamGlobalKey,omitempty"`
	// 团队Id
	TeamId *int32 `json:"TeamId,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var phone string = ""
	this.Phone = &phone
	var teamGlobalKey string = ""
	this.TeamGlobalKey = &teamGlobalKey
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var phone string = ""
	this.Phone = &phone
	var teamGlobalKey string = ""
	this.TeamGlobalKey = &teamGlobalKey
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *User) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *User) HasAvatar() bool {
	if o != nil && !utils.IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *User) SetAvatar(v string) {
	o.Avatar = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetGlobalKey returns the GlobalKey field value if set, zero value otherwise.
func (o *User) GetGlobalKey() string {
	if o == nil || utils.IsNil(o.GlobalKey) {
		var ret string
		return ret
	}
	return *o.GlobalKey
}

// GetGlobalKeyOk returns a tuple with the GlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GlobalKey) {
		return nil, false
	}
	return o.GlobalKey, true
}

// HasGlobalKey returns a boolean if a field has been set.
func (o *User) HasGlobalKey() bool {
	if o != nil && !utils.IsNil(o.GlobalKey) {
		return true
	}

	return false
}

// SetGlobalKey gets a reference to the given string and assigns it to the GlobalKey field.
func (o *User) SetGlobalKey(v string) {
	o.GlobalKey = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *User) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *User) GetPhone() string {
	if o == nil || utils.IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && !utils.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetStatus() int32 {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given utils.NullableInt32 and assigns it to the Status field.
func (o *User) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *User) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *User) UnsetStatus() {
	o.Status.Unset()
}

// GetTeamGlobalKey returns the TeamGlobalKey field value if set, zero value otherwise.
func (o *User) GetTeamGlobalKey() string {
	if o == nil || utils.IsNil(o.TeamGlobalKey) {
		var ret string
		return ret
	}
	return *o.TeamGlobalKey
}

// GetTeamGlobalKeyOk returns a tuple with the TeamGlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTeamGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TeamGlobalKey) {
		return nil, false
	}
	return o.TeamGlobalKey, true
}

// HasTeamGlobalKey returns a boolean if a field has been set.
func (o *User) HasTeamGlobalKey() bool {
	if o != nil && !utils.IsNil(o.TeamGlobalKey) {
		return true
	}

	return false
}

// SetTeamGlobalKey gets a reference to the given string and assigns it to the TeamGlobalKey field.
func (o *User) SetTeamGlobalKey(v string) {
	o.TeamGlobalKey = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *User) GetTeamId() int32 {
	if o == nil || utils.IsNil(o.TeamId) {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTeamIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *User) HasTeamId() bool {
	if o != nil && !utils.IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *User) SetTeamId(v int32) {
	o.TeamId = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
	}
	if !utils.IsNil(o.Email) {
		toSerialize["Email"] = o.Email
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
	if !utils.IsNil(o.Phone) {
		toSerialize["Phone"] = o.Phone
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if !utils.IsNil(o.TeamGlobalKey) {
		toSerialize["TeamGlobalKey"] = o.TeamGlobalKey
	}
	if !utils.IsNil(o.TeamId) {
		toSerialize["TeamId"] = o.TeamId
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


