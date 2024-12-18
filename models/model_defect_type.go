/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DefectType type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DefectType{}

// DefectType 缺陷类型
type DefectType struct {
	// 图标地址
	IconUrl *string `json:"IconUrl,omitempty"`
	// 缺陷类型 Id
	Id *int64 `json:"Id,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
}

// NewDefectType instantiates a new DefectType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefectType() *DefectType {
	this := DefectType{}
	var iconUrl string = ""
	this.IconUrl = &iconUrl
	var name string = ""
	this.Name = &name
	return &this
}

// NewDefectTypeWithDefaults instantiates a new DefectType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefectTypeWithDefaults() *DefectType {
	this := DefectType{}
	var iconUrl string = ""
	this.IconUrl = &iconUrl
	var name string = ""
	this.Name = &name
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *DefectType) GetIconUrl() string {
	if o == nil || utils.IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectType) GetIconUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DefectType) HasIconUrl() bool {
	if o != nil && !utils.IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *DefectType) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DefectType) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectType) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DefectType) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DefectType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DefectType) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefectType) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DefectType) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DefectType) SetName(v string) {
	o.Name = &v
}

func (o DefectType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefectType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IconUrl) {
		toSerialize["IconUrl"] = o.IconUrl
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableDefectType struct {
	value *DefectType
	isSet bool
}

func (v NullableDefectType) Get() *DefectType {
	return v.value
}

func (v *NullableDefectType) Set(val *DefectType) {
	v.value = val
	v.isSet = true
}

func (v NullableDefectType) IsSet() bool {
	return v.isSet
}

func (v *NullableDefectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefectType(val *DefectType) *NullableDefectType {
	return &NullableDefectType{value: val, isSet: true}
}

func (v NullableDefectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


