/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DepotSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepotSpec{}

// DepotSpec 仓库规范
type DepotSpec struct {
	// 仓库规范描述信息
	Description *string `json:"Description,omitempty"`
	// 仓库规范名字
	Name *string `json:"Name,omitempty"`
	// 仓库规范类型 system:系统级别，team：团队级别
	Type *string `json:"Type,omitempty"`
}

// NewDepotSpec instantiates a new DepotSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotSpec() *DepotSpec {
	this := DepotSpec{}
	var description string = ""
	this.Description = &description
	var name string = ""
	this.Name = &name
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewDepotSpecWithDefaults instantiates a new DepotSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotSpecWithDefaults() *DepotSpec {
	this := DepotSpec{}
	var description string = ""
	this.Description = &description
	var name string = ""
	this.Name = &name
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DepotSpec) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DepotSpec) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DepotSpec) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DepotSpec) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpec) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DepotSpec) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DepotSpec) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DepotSpec) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpec) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DepotSpec) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DepotSpec) SetType(v string) {
	o.Type = &v
}

func (o DepotSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDepotSpec struct {
	value *DepotSpec
	isSet bool
}

func (v NullableDepotSpec) Get() *DepotSpec {
	return v.value
}

func (v *NullableDepotSpec) Set(val *DepotSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotSpec(val *DepotSpec) *NullableDepotSpec {
	return &NullableDepotSpec{value: val, isSet: true}
}

func (v NullableDepotSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


