/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CdDeployTime type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CdDeployTime{}

// CdDeployTime CdDeployTime 结构
type CdDeployTime struct {
	// 发布时长（单位秒）
	DeployTime *int64 `json:"DeployTime,omitempty"`
	// 发布总次数
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewCdDeployTime instantiates a new CdDeployTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdDeployTime() *CdDeployTime {
	this := CdDeployTime{}
	return &this
}

// NewCdDeployTimeWithDefaults instantiates a new CdDeployTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdDeployTimeWithDefaults() *CdDeployTime {
	this := CdDeployTime{}
	return &this
}

// GetDeployTime returns the DeployTime field value if set, zero value otherwise.
func (o *CdDeployTime) GetDeployTime() int64 {
	if o == nil || utils.IsNil(o.DeployTime) {
		var ret int64
		return ret
	}
	return *o.DeployTime
}

// GetDeployTimeOk returns a tuple with the DeployTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTime) GetDeployTimeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DeployTime) {
		return nil, false
	}
	return o.DeployTime, true
}

// HasDeployTime returns a boolean if a field has been set.
func (o *CdDeployTime) HasDeployTime() bool {
	if o != nil && !utils.IsNil(o.DeployTime) {
		return true
	}

	return false
}

// SetDeployTime gets a reference to the given int64 and assigns it to the DeployTime field.
func (o *CdDeployTime) SetDeployTime(v int64) {
	o.DeployTime = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CdDeployTime) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTime) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CdDeployTime) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *CdDeployTime) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o CdDeployTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdDeployTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DeployTime) {
		toSerialize["DeployTime"] = o.DeployTime
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableCdDeployTime struct {
	value *CdDeployTime
	isSet bool
}

func (v NullableCdDeployTime) Get() *CdDeployTime {
	return v.value
}

func (v *NullableCdDeployTime) Set(val *CdDeployTime) {
	v.value = val
	v.isSet = true
}

func (v NullableCdDeployTime) IsSet() bool {
	return v.isSet
}

func (v *NullableCdDeployTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdDeployTime(val *CdDeployTime) *NullableCdDeployTime {
	return &NullableCdDeployTime{value: val, isSet: true}
}

func (v NullableCdDeployTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdDeployTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

