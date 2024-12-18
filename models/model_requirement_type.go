/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the RequirementType type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequirementType{}

// RequirementType 需求类型
type RequirementType struct {
	// 需求类型 Id
	Id *int64 `json:"Id,omitempty"`
	// 需求类型名称
	Name *string `json:"Name,omitempty"`
}

// NewRequirementType instantiates a new RequirementType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequirementType() *RequirementType {
	this := RequirementType{}
	var name string = ""
	this.Name = &name
	return &this
}

// NewRequirementTypeWithDefaults instantiates a new RequirementType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequirementTypeWithDefaults() *RequirementType {
	this := RequirementType{}
	var name string = ""
	this.Name = &name
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequirementType) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequirementType) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequirementType) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RequirementType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequirementType) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequirementType) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequirementType) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequirementType) SetName(v string) {
	o.Name = &v
}

func (o RequirementType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequirementType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRequirementType struct {
	value *RequirementType
	isSet bool
}

func (v NullableRequirementType) Get() *RequirementType {
	return v.value
}

func (v *NullableRequirementType) Set(val *RequirementType) {
	v.value = val
	v.isSet = true
}

func (v NullableRequirementType) IsSet() bool {
	return v.isSet
}

func (v *NullableRequirementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequirementType(val *RequirementType) *NullableRequirementType {
	return &NullableRequirementType{value: val, isSet: true}
}

func (v NullableRequirementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequirementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


