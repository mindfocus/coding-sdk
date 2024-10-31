/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the TestsData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestsData{}

// TestsData 测试任务列表
type TestsData struct {
	// 测试任务信息
	Tests []Test `json:"Tests,omitempty"`
}

// NewTestsData instantiates a new TestsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestsData() *TestsData {
	this := TestsData{}
	return &this
}

// NewTestsDataWithDefaults instantiates a new TestsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsDataWithDefaults() *TestsData {
	this := TestsData{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestsData) GetTests() []Test {
	if o == nil {
		var ret []Test
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestsData) GetTestsOk() ([]Test, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *TestsData) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []Test and assigns it to the Tests field.
func (o *TestsData) SetTests(v []Test) {
	o.Tests = v
}

func (o TestsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tests != nil {
		toSerialize["Tests"] = o.Tests
	}
	return toSerialize, nil
}

type NullableTestsData struct {
	value *TestsData
	isSet bool
}

func (v NullableTestsData) Get() *TestsData {
	return v.value
}

func (v *NullableTestsData) Set(val *TestsData) {
	v.value = val
	v.isSet = true
}

func (v NullableTestsData) IsSet() bool {
	return v.isSet
}

func (v *NullableTestsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestsData(val *TestsData) *NullableTestsData {
	return &NullableTestsData{value: val, isSet: true}
}

func (v NullableTestsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


