/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the TestData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestData{}

// TestData 测试任务详情
type TestData struct {
	Test *TestFull `json:"Test,omitempty"`
}

// NewTestData instantiates a new TestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestData() *TestData {
	this := TestData{}
	return &this
}

// NewTestDataWithDefaults instantiates a new TestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestDataWithDefaults() *TestData {
	this := TestData{}
	return &this
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *TestData) GetTest() TestFull {
	if o == nil || utils.IsNil(o.Test) {
		var ret TestFull
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestData) GetTestOk() (*TestFull, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *TestData) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given TestFull and assigns it to the Test field.
func (o *TestData) SetTest(v TestFull) {
	o.Test = &v
}

func (o TestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Test) {
		toSerialize["Test"] = o.Test
	}
	return toSerialize, nil
}

type NullableTestData struct {
	value *TestData
	isSet bool
}

func (v NullableTestData) Get() *TestData {
	return v.value
}

func (v *NullableTestData) Set(val *TestData) {
	v.value = val
	v.isSet = true
}

func (v NullableTestData) IsSet() bool {
	return v.isSet
}

func (v *NullableTestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestData(val *TestData) *NullableTestData {
	return &NullableTestData{value: val, isSet: true}
}

func (v NullableTestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


