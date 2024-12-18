/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Role{}

// Role 用户组
type Role struct {
	// 用户组Id
	RoleId *int32 `json:"RoleId,omitempty"`
	// 用户组类型： UserDefined 用户自定义的角色， EnterpriseOwner 企业所有者，EnterpriseAdmin 企业管理员， EnterpriseMember 企业普通成员， ProjectAdmin 项目管理员， ProjectMember 项目成员-> 新的权限系统里面叫\"开发\"，ProjectGuest 项目受限成员 -> 新的权限系统里面叫\"测试\"，ProjectManager 项目经理，ProductManager 产品经理，ProjectOperation 运维 ProgramOwner 项目集负责人，ProgramAdmin 项目集管理员，ProgramMember 项目集成员， ProgramProjectMember 项目集-项目成员
	RoleType *string `json:"RoleType,omitempty"`
	// 用户组类型名称
	RoleTypeName *string `json:"RoleTypeName,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	var roleType string = ""
	this.RoleType = &roleType
	var roleTypeName string = ""
	this.RoleTypeName = &roleTypeName
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	var roleType string = ""
	this.RoleType = &roleType
	var roleTypeName string = ""
	this.RoleTypeName = &roleTypeName
	return &this
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *Role) GetRoleId() int32 {
	if o == nil || utils.IsNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *Role) HasRoleId() bool {
	if o != nil && !utils.IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *Role) SetRoleId(v int32) {
	o.RoleId = &v
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *Role) GetRoleType() string {
	if o == nil || utils.IsNil(o.RoleType) {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RoleType) {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *Role) HasRoleType() bool {
	if o != nil && !utils.IsNil(o.RoleType) {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *Role) SetRoleType(v string) {
	o.RoleType = &v
}

// GetRoleTypeName returns the RoleTypeName field value if set, zero value otherwise.
func (o *Role) GetRoleTypeName() string {
	if o == nil || utils.IsNil(o.RoleTypeName) {
		var ret string
		return ret
	}
	return *o.RoleTypeName
}

// GetRoleTypeNameOk returns a tuple with the RoleTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleTypeNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RoleTypeName) {
		return nil, false
	}
	return o.RoleTypeName, true
}

// HasRoleTypeName returns a boolean if a field has been set.
func (o *Role) HasRoleTypeName() bool {
	if o != nil && !utils.IsNil(o.RoleTypeName) {
		return true
	}

	return false
}

// SetRoleTypeName gets a reference to the given string and assigns it to the RoleTypeName field.
func (o *Role) SetRoleTypeName(v string) {
	o.RoleTypeName = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.RoleId) {
		toSerialize["RoleId"] = o.RoleId
	}
	if !utils.IsNil(o.RoleType) {
		toSerialize["RoleType"] = o.RoleType
	}
	if !utils.IsNil(o.RoleTypeName) {
		toSerialize["RoleTypeName"] = o.RoleTypeName
	}
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


