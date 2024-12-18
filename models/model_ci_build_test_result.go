/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CIBuildTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CIBuildTestResult{}

// CIBuildTestResult 构建的测试结果
type CIBuildTestResult struct {
	// 时长
	Duration int32 `json:"Duration"`
	// 是否空
	Empty bool `json:"Empty"`
	// 失败次数
	FailCount int32 `json:"FailCount"`
	// 通过次数
	PassCount int32 `json:"PassCount"`
	// 跳过次数
	SkipCount int32 `json:"SkipCount"`
	// 总次数
	TotalCount int32 `json:"TotalCount"`
}

type _CIBuildTestResult CIBuildTestResult

// NewCIBuildTestResult instantiates a new CIBuildTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCIBuildTestResult(duration int32, empty bool, failCount int32, passCount int32, skipCount int32, totalCount int32) *CIBuildTestResult {
	this := CIBuildTestResult{}
	this.Duration = duration
	this.Empty = empty
	this.FailCount = failCount
	this.PassCount = passCount
	this.SkipCount = skipCount
	this.TotalCount = totalCount
	return &this
}

// NewCIBuildTestResultWithDefaults instantiates a new CIBuildTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCIBuildTestResultWithDefaults() *CIBuildTestResult {
	this := CIBuildTestResult{}
	var empty bool = false
	this.Empty = empty
	return &this
}

// GetDuration returns the Duration field value
func (o *CIBuildTestResult) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CIBuildTestResult) SetDuration(v int32) {
	o.Duration = v
}

// GetEmpty returns the Empty field value
func (o *CIBuildTestResult) GetEmpty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetEmptyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Empty, true
}

// SetEmpty sets field value
func (o *CIBuildTestResult) SetEmpty(v bool) {
	o.Empty = v
}

// GetFailCount returns the FailCount field value
func (o *CIBuildTestResult) GetFailCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailCount
}

// GetFailCountOk returns a tuple with the FailCount field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetFailCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailCount, true
}

// SetFailCount sets field value
func (o *CIBuildTestResult) SetFailCount(v int32) {
	o.FailCount = v
}

// GetPassCount returns the PassCount field value
func (o *CIBuildTestResult) GetPassCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PassCount
}

// GetPassCountOk returns a tuple with the PassCount field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetPassCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassCount, true
}

// SetPassCount sets field value
func (o *CIBuildTestResult) SetPassCount(v int32) {
	o.PassCount = v
}

// GetSkipCount returns the SkipCount field value
func (o *CIBuildTestResult) GetSkipCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetSkipCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipCount, true
}

// SetSkipCount sets field value
func (o *CIBuildTestResult) SetSkipCount(v int32) {
	o.SkipCount = v
}

// GetTotalCount returns the TotalCount field value
func (o *CIBuildTestResult) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *CIBuildTestResult) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *CIBuildTestResult) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o CIBuildTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CIBuildTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Duration"] = o.Duration
	toSerialize["Empty"] = o.Empty
	toSerialize["FailCount"] = o.FailCount
	toSerialize["PassCount"] = o.PassCount
	toSerialize["SkipCount"] = o.SkipCount
	toSerialize["TotalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *CIBuildTestResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Duration",
		"Empty",
		"FailCount",
		"PassCount",
		"SkipCount",
		"TotalCount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCIBuildTestResult := _CIBuildTestResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCIBuildTestResult)

	if err != nil {
		return err
	}

	*o = CIBuildTestResult(varCIBuildTestResult)

	return err
}

type NullableCIBuildTestResult struct {
	value *CIBuildTestResult
	isSet bool
}

func (v NullableCIBuildTestResult) Get() *CIBuildTestResult {
	return v.value
}

func (v *NullableCIBuildTestResult) Set(val *CIBuildTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCIBuildTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCIBuildTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCIBuildTestResult(val *CIBuildTestResult) *NullableCIBuildTestResult {
	return &NullableCIBuildTestResult{value: val, isSet: true}
}

func (v NullableCIBuildTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCIBuildTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


