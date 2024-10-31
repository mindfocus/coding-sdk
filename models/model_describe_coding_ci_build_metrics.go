/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuildMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildMetrics{}

// DescribeCodingCIBuildMetrics DescribeCodingCIBuildMetrics 结构
type DescribeCodingCIBuildMetrics struct {
	// 日期
	Date utils.NullableString `json:"Date,omitempty"`
	// 构建成功总数
	SuccessBuildCount *int32 `json:"SuccessBuildCount,omitempty"`
	// 构建总数
	TotalBuildCount *int32 `json:"TotalBuildCount,omitempty"`
	// 构建总耗时，单位毫秒
	TotalDuration *int32 `json:"TotalDuration,omitempty"`
}

// NewDescribeCodingCIBuildMetrics instantiates a new DescribeCodingCIBuildMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildMetrics() *DescribeCodingCIBuildMetrics {
	this := DescribeCodingCIBuildMetrics{}
	var date string = ""
	this.Date = *utils.NewNullableString(&date)
	return &this
}

// NewDescribeCodingCIBuildMetricsWithDefaults instantiates a new DescribeCodingCIBuildMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildMetricsWithDefaults() *DescribeCodingCIBuildMetrics {
	this := DescribeCodingCIBuildMetrics{}
	var date string = ""
	this.Date = *utils.NewNullableString(&date)
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeCodingCIBuildMetrics) GetDate() string {
	if o == nil || utils.IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeCodingCIBuildMetrics) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildMetrics) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given utils.NullableString and assigns it to the Date field.
func (o *DescribeCodingCIBuildMetrics) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *DescribeCodingCIBuildMetrics) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *DescribeCodingCIBuildMetrics) UnsetDate() {
	o.Date.Unset()
}

// GetSuccessBuildCount returns the SuccessBuildCount field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildMetrics) GetSuccessBuildCount() int32 {
	if o == nil || utils.IsNil(o.SuccessBuildCount) {
		var ret int32
		return ret
	}
	return *o.SuccessBuildCount
}

// GetSuccessBuildCountOk returns a tuple with the SuccessBuildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildMetrics) GetSuccessBuildCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SuccessBuildCount) {
		return nil, false
	}
	return o.SuccessBuildCount, true
}

// HasSuccessBuildCount returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildMetrics) HasSuccessBuildCount() bool {
	if o != nil && !utils.IsNil(o.SuccessBuildCount) {
		return true
	}

	return false
}

// SetSuccessBuildCount gets a reference to the given int32 and assigns it to the SuccessBuildCount field.
func (o *DescribeCodingCIBuildMetrics) SetSuccessBuildCount(v int32) {
	o.SuccessBuildCount = &v
}

// GetTotalBuildCount returns the TotalBuildCount field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildMetrics) GetTotalBuildCount() int32 {
	if o == nil || utils.IsNil(o.TotalBuildCount) {
		var ret int32
		return ret
	}
	return *o.TotalBuildCount
}

// GetTotalBuildCountOk returns a tuple with the TotalBuildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildMetrics) GetTotalBuildCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalBuildCount) {
		return nil, false
	}
	return o.TotalBuildCount, true
}

// HasTotalBuildCount returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildMetrics) HasTotalBuildCount() bool {
	if o != nil && !utils.IsNil(o.TotalBuildCount) {
		return true
	}

	return false
}

// SetTotalBuildCount gets a reference to the given int32 and assigns it to the TotalBuildCount field.
func (o *DescribeCodingCIBuildMetrics) SetTotalBuildCount(v int32) {
	o.TotalBuildCount = &v
}

// GetTotalDuration returns the TotalDuration field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildMetrics) GetTotalDuration() int32 {
	if o == nil || utils.IsNil(o.TotalDuration) {
		var ret int32
		return ret
	}
	return *o.TotalDuration
}

// GetTotalDurationOk returns a tuple with the TotalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildMetrics) GetTotalDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalDuration) {
		return nil, false
	}
	return o.TotalDuration, true
}

// HasTotalDuration returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildMetrics) HasTotalDuration() bool {
	if o != nil && !utils.IsNil(o.TotalDuration) {
		return true
	}

	return false
}

// SetTotalDuration gets a reference to the given int32 and assigns it to the TotalDuration field.
func (o *DescribeCodingCIBuildMetrics) SetTotalDuration(v int32) {
	o.TotalDuration = &v
}

func (o DescribeCodingCIBuildMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Date.IsSet() {
		toSerialize["Date"] = o.Date.Get()
	}
	if !utils.IsNil(o.SuccessBuildCount) {
		toSerialize["SuccessBuildCount"] = o.SuccessBuildCount
	}
	if !utils.IsNil(o.TotalBuildCount) {
		toSerialize["TotalBuildCount"] = o.TotalBuildCount
	}
	if !utils.IsNil(o.TotalDuration) {
		toSerialize["TotalDuration"] = o.TotalDuration
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIBuildMetrics struct {
	value *DescribeCodingCIBuildMetrics
	isSet bool
}

func (v NullableDescribeCodingCIBuildMetrics) Get() *DescribeCodingCIBuildMetrics {
	return v.value
}

func (v *NullableDescribeCodingCIBuildMetrics) Set(val *DescribeCodingCIBuildMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildMetrics(val *DescribeCodingCIBuildMetrics) *NullableDescribeCodingCIBuildMetrics {
	return &NullableDescribeCodingCIBuildMetrics{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

