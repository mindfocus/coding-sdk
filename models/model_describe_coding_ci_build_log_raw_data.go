/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuildLogRawData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildLogRawData{}

// DescribeCodingCIBuildLogRawData 构建完整的日志
type DescribeCodingCIBuildLogRawData struct {
	// 日志
	Raw *string `json:"Raw,omitempty"`
}

// NewDescribeCodingCIBuildLogRawData instantiates a new DescribeCodingCIBuildLogRawData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildLogRawData() *DescribeCodingCIBuildLogRawData {
	this := DescribeCodingCIBuildLogRawData{}
	var raw string = ""
	this.Raw = &raw
	return &this
}

// NewDescribeCodingCIBuildLogRawDataWithDefaults instantiates a new DescribeCodingCIBuildLogRawData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildLogRawDataWithDefaults() *DescribeCodingCIBuildLogRawData {
	this := DescribeCodingCIBuildLogRawData{}
	var raw string = ""
	this.Raw = &raw
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildLogRawData) GetRaw() string {
	if o == nil || utils.IsNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildLogRawData) GetRawOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildLogRawData) HasRaw() bool {
	if o != nil && !utils.IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *DescribeCodingCIBuildLogRawData) SetRaw(v string) {
	o.Raw = &v
}

func (o DescribeCodingCIBuildLogRawData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildLogRawData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Raw) {
		toSerialize["Raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIBuildLogRawData struct {
	value *DescribeCodingCIBuildLogRawData
	isSet bool
}

func (v NullableDescribeCodingCIBuildLogRawData) Get() *DescribeCodingCIBuildLogRawData {
	return v.value
}

func (v *NullableDescribeCodingCIBuildLogRawData) Set(val *DescribeCodingCIBuildLogRawData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildLogRawData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildLogRawData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildLogRawData(val *DescribeCodingCIBuildLogRawData) *NullableDescribeCodingCIBuildLogRawData {
	return &NullableDescribeCodingCIBuildLogRawData{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildLogRawData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildLogRawData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


