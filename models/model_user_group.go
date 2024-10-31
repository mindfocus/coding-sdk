/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the UserGroup type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserGroup{}

// UserGroup 用户组
type UserGroup struct {
	// 用户组创建时间
	CreatedAt utils.NullableInt64 `json:"CreatedAt,omitempty"`
	// 用户组描述
	Description utils.NullableString `json:"Description,omitempty"`
	// 用户组 ID
	Id *int64 `json:"Id,omitempty"`
	// 用户组名称
	Name *string `json:"Name,omitempty"`
}

// NewUserGroup instantiates a new UserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup() *UserGroup {
	this := UserGroup{}
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var name string = ""
	this.Name = &name
	return &this
}

// NewUserGroupWithDefaults instantiates a new UserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupWithDefaults() *UserGroup {
	this := UserGroup{}
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var name string = ""
	this.Name = &name
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroup) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroup) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserGroup) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given utils.NullableInt64 and assigns it to the CreatedAt field.
func (o *UserGroup) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *UserGroup) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *UserGroup) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroup) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroup) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UserGroup) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given utils.NullableString and assigns it to the Description field.
func (o *UserGroup) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UserGroup) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UserGroup) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserGroup) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserGroup) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserGroup) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserGroup) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserGroup) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserGroup) SetName(v string) {
	o.Name = &v
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUserGroup struct {
	value *UserGroup
	isSet bool
}

func (v NullableUserGroup) Get() *UserGroup {
	return v.value
}

func (v *NullableUserGroup) Set(val *UserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup(val *UserGroup) *NullableUserGroup {
	return &NullableUserGroup{value: val, isSet: true}
}

func (v NullableUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


